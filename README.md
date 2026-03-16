# Immich Desktop Sync

A Windows desktop app for syncing and managing your [Immich](https://immich.app) photo library. Built with [Wails](https://wails.io) (Go + Svelte).

## Features

- **Gallery** - browse all your Immich photos and albums with lazy-loaded thumbnails
- **Upload** - drag & drop images onto the app, or add watched folders that auto-sync
- **Download** - download original files from your Immich server to a local folder
- **Upload queue** - real-time progress panel showing pending, uploading, and failed items with retry support
- **Lightbox** - full-screen image viewer with keyboard navigation and per-image download

## Requirements

- [Go 1.21+](https://go.dev/dl/)
- [Node.js 18+](https://nodejs.org/)
- [Wails CLI](https://wails.io/docs/gettingstarted/installation): `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
- An [Immich](https://immich.app) server

> No C compiler (GCC/CGO) required - SQLite uses the pure-Go `modernc.org/sqlite` driver.

## Development

```bash
wails dev
```

Starts the app with hot-reload. A browser dev server is also available at `http://localhost:34115`.

## Build

```bash
wails build
```

Produces a self-contained executable at `build/bin/ImmichDesktopSync.exe`.

## Configuration

Settings are stored in `~/.config/immich-desktop/config.json`. The SQLite database (upload queue, thumbnail cache, watched folders) lives at `~/.config/immich-desktop/sync.db`.

## Usage

1. **Login** - enter your Immich server URL and credentials
2. **Upload via drag & drop** - drag image/video files onto the app window
3. **Upload via watched folder** - go to Settings → Watched Folders → add a folder path; all media inside is scanned and queued automatically
4. **Download** - open any photo in the lightbox and click the download icon (↓); set your downloads folder in Settings first
5. **Gallery refresh** - after uploads complete, a blue dot appears on the refresh button; click it to reload the gallery with newly uploaded photos
