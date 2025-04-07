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
	        Theme: windows.Dark,
	        CustomTheme: &windows.ThemeSettings{
	             DarkModeTitleBar:   windows.RGB(18, 24, 38),
	             DarkModeTitleText:  windows.RGB(200, 200, 200),
	             DarkModeBorder:     windows.RGB(30, 30, 37),
	             LightModeTitleBar:  windows.RGB(240, 240, 240),
	             LightModeTitleText: windows.RGB(20, 20, 20),
	             LightModeBorder:    windows.RGB(240, 240, 240),
	             DarkModeTitleBarInactive:   windows.RGB(30, 30, 37),
	             DarkModeTitleTextInactive:  windows.RGB(128, 128, 128),
	             DarkModeBorderInactive:     windows.RGB(30, 30, 37),
	        },
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