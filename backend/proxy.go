package backend

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"immich-desktop-sync/backend/immich"
)

type StreamProxy struct {
	port   int
	client *immich.Client
	srv    *http.Server
}

func NewStreamProxy(client *immich.Client) (*StreamProxy, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	p := &StreamProxy{
		port:   ln.Addr().(*net.TCPAddr).Port,
		client: client,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/video/", p.handleVideo)
	p.srv = &http.Server{Handler: mux}

	go func() {
		if err := p.srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Printf("stream proxy: %v", err)
		}
	}()

	log.Printf("stream proxy listening on 127.0.0.1:%d", p.port)
	return p, nil
}

func (p *StreamProxy) Port() int { return p.port }

func (p *StreamProxy) Close() { _ = p.srv.Close() }

func (p *StreamProxy) handleVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Range")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Range, Accept-Ranges")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	assetID := strings.TrimPrefix(r.URL.Path, "/video/")
	if assetID == "" {
		http.Error(w, "missing asset id", http.StatusBadRequest)
		return
	}

	immichURL := fmt.Sprintf("%s/api/assets/%s/original", p.client.BaseURL, assetID)
	req, err := http.NewRequestWithContext(r.Context(), "GET", immichURL, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Authorization", "Bearer "+p.client.Token())

	if rng := r.Header.Get("Range"); rng != "" {
		req.Header.Set("Range", rng)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for _, h := range []string{
		"Content-Type", "Content-Length", "Content-Range",
		"Accept-Ranges", "Last-Modified", "ETag",
	} {
		if v := resp.Header.Get(h); v != "" {
			w.Header().Set(h, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}
