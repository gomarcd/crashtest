package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create API service
	apiService := NewAPIService()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Crashtest",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    false,
			Theme:                windows.Dark,
			Frameless:            true,
		},
		Mac: &mac.Options{
		    TitleBar: &mac.TitleBar{
		        TitlebarAppearsTransparent: true,
		        HideTitle:                  false,
		        HideTitleBar:               false,
		        FullSizeContent:            true,
		        UseToolbar:                 false,
		        HideToolbarSeparator:       true,
		    },
		    Appearance: mac.NSAppearanceNameDarkAqua,
		},

		BackgroundColour: &options.RGBA{R: 17, G: 24, B: 39, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			apiService,
		},
	})

	if err != nil {
		log.Fatal("Error starting application:", err)
	}
}