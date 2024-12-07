package main

import (
	"context"
	"embed"

	"Wails-train/pkg/testpkg"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	test := testpkg.NewTest("test_value")

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Wails-train",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup: func(ctx context.Context) {
			test.Startup(ctx)
		},

		Bind: []interface{}{
			test,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
