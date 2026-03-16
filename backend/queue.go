package backend

import (
	"log"
	"sync"
	"time"

	"immich-desktop-sync/backend/db"
	"immich-desktop-sync/backend/immich"
)

const (
	maxConcurrent = 3
	maxRetries    = 5
)

var retryDelays = []time.Duration{
	1 * time.Minute,
	5 * time.Minute,
	15 * time.Minute,
	30 * time.Minute,
	60 * time.Minute,
}

type UploadQueue struct {
	database    *db.DB
	client      *immich.Client
	sem         chan struct{}
	wg          sync.WaitGroup
	stop        chan struct{}
	notify      chan struct{}
	onStart     func()
	onDone      func()
}

func NewUploadQueue(database *db.DB, client *immich.Client, onStart, onDone func()) *UploadQueue {
	return &UploadQueue{
		database: database,
		client:   client,
		sem:      make(chan struct{}, maxConcurrent),
		stop:     make(chan struct{}),
		notify:   make(chan struct{}, 1),
		onStart:  onStart,
		onDone:   onDone,
	}
}

func (q *UploadQueue) Start() {
	go q.loop()
}

func (q *UploadQueue) Stop() {
	close(q.stop)
	q.wg.Wait()
}

func (q *UploadQueue) Notify() {
	select {
	case q.notify <- struct{}{}:
	default:
	}
}

func (q *UploadQueue) loop() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-q.stop:
			return
		case <-q.notify:
			q.processNext()
		case <-ticker.C:
			q.processNext()
		}
	}
}

func (q *UploadQueue) processNext() {
	items, err := q.database.DequeueNextPending(maxConcurrent)
	if err != nil {
		log.Printf("queue: dequeue error: %v", err)
		return
	}
	for _, item := range items {
		item := item
		if item.RetryCount > 0 && item.LastAttempt != "" {
			last, err := time.Parse(time.RFC3339, item.LastAttempt)
			if err == nil && time.Since(last) < retryDelay(item.RetryCount) {
				continue
			}
		}

		q.sem <- struct{}{}
		q.wg.Add(1)
		go func() {
			defer func() {
				<-q.sem
				q.wg.Done()
				q.Notify()
				if q.onDone != nil {
					q.onDone()
				}
			}()

			if err := q.database.MarkUploading(item.ID); err != nil {
				log.Printf("queue: mark uploading %d: %v", item.ID, err)
				return
			}
			if q.onStart != nil {
				q.onStart()
			}

			assetID, err := q.client.UploadFile(item.FilePath)
			if err != nil {
				log.Printf("queue: upload %s failed: %v", item.FilePath, err)
				_ = q.database.MarkFailed(item.ID, err.Error())
				return
			}

			if err := q.database.MarkDone(item.ID, item.FilePath, assetID); err != nil {
				log.Printf("queue: mark done %d: %v", item.ID, err)
			} else {
				log.Printf("queue: uploaded %s → %s", item.FilePath, assetID)
			}
		}()
	}
}

func retryDelay(retryCount int) time.Duration {
	idx := retryCount - 1
	if idx < 0 {
		return 0
	}
	if idx >= len(retryDelays) {
		return retryDelays[len(retryDelays)-1]
	}
	return retryDelays[idx]
}
