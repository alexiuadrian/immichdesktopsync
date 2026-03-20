package main

import (
	"embed"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"immich-desktop-sync/backend"
)

//go:embed all:frontend/dist
var assets embed.FS

func setupLogging() {
	logPath := backend.LogPath()
	if err := os.MkdirAll(filepath.Dir(logPath), 0700); err == nil {
		if f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600); err == nil {
			log.SetOutput(io.MultiWriter(os.Stderr, f))
		}
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	setupLogging()

	app := NewApp()

	err := wails.Run(&options.App{
		Title:            "Immich Desktop Sync",
		Width:            1200,
		Height:           800,
		MinWidth:         800,
		MinHeight:        600,
		BackgroundColour: &options.RGBA{R: 17, G: 17, B: 27, A: 1},
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop: true,
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		log.Fatalf("wails: %v", err)
	}
}
