package backend

import (
	"github.com/getlantern/systray"
)

type TrayManager struct {
	onShow func()
	onQuit func()
}

func NewTrayManager(onShow, onQuit func()) *TrayManager {
	return &TrayManager{onShow: onShow, onQuit: onQuit}
}

func (t *TrayManager) Run() {
	systray.Run(t.onReady, t.onExit)
}

func (t *TrayManager) RunAsync() {
	go t.Run()
}

func (t *TrayManager) onReady() {
	systray.SetTitle("Immich Sync")
	systray.SetTooltip("Immich Desktop Sync")

	mShow := systray.AddMenuItem("Open", "Show the application window")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit Immich Desktop Sync")

	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				if t.onShow != nil {
					t.onShow()
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func (t *TrayManager) onExit() {
	if t.onQuit != nil {
		t.onQuit()
	}
}

func (t *TrayManager) SetSyncing(active bool) {
	if active {
		systray.SetTooltip("Immich Desktop Sync — syncing…")
	} else {
		systray.SetTooltip("Immich Desktop Sync — idle")
	}
}
