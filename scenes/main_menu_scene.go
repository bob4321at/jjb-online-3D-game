package scenes

import (
	"main/networking"
	"main/utils"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenuScene struct {
	SetedUp bool
}

var ServerName *string

func (scene *MainMenuScene) Setup() {
	ServerName = &networking.ServerName
	scene.SetedUp = true
}

func (scene *MainMenuScene) Draw() {
	rl.DrawText("Enter Server: "+*ServerName, 10, 10, 32, rl.White)
	rl.DrawText("Press Enter To Connect. Press Shift Enter To Enter Serverless", 10, 74, 48, rl.White)
}

func (scene *MainMenuScene) Update() {
	var hit_keys []uint

	hit_keys = append(hit_keys, uint(rl.GetKeyPressed()))

	if !rl.IsKeyDown(rl.KeyLeftControl) {
		if len(hit_keys) != 0 {
			if rl.IsKeyDown(rl.KeyLeftShift) {
				*ServerName += utils.Key_To_String[hit_keys[0]]
			} else {
				*ServerName += strings.ToLower(utils.Key_To_String[hit_keys[0]])
			}
		}
	}

	if len(*ServerName) != 0 {
		if rl.IsKeyPressed(rl.KeyBackspace) {
			*ServerName = (*ServerName)[:len(*ServerName)-1]
		}
	}

	if rl.IsKeyPressed(rl.KeyDelete) {
		*ServerName = ""
	}

	if rl.IsKeyPressed(rl.KeyV) {
		if rl.IsKeyDown(rl.KeyLeftControl) {
			*ServerName = strings.Replace(rl.GetClipboardText(), "https", "https", 1) + ":8080"
		}
	}

	if rl.IsKeyPressed(rl.KeyEnter) && !rl.IsKeyDown(rl.KeyLeftShift) {
		Current_Scene_Id = 1
	} else if rl.IsKeyDown(rl.KeyLeftControl) {
		*ServerName = ""
		Current_Scene_Id = 1
	}

}

func (scene *MainMenuScene) GetSetup() bool {
	return scene.SetedUp
}
