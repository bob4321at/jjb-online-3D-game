package main

import (
	"image/color"
	"main/level"
	"main/networking"
	"main/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw() {
	rl.BeginDrawing()

	rl.ClearBackground(color.RGBA{75, 125, 255, 255})

	scenes.Scenes[scenes.Current_Scene_Id].Draw()

	rl.DrawFPS(10, 10)

	rl.EndDrawing()
}

func update() {
	if rl.IsKeyPressed(rl.KeyT) {
		fps_toggle = !fps_toggle
		if fps_toggle {
			rl.SetTargetFPS(60)
		} else {
			rl.SetTargetFPS(0)
		}
	}

	if !scenes.Scenes[scenes.Current_Scene_Id].GetSetup() {
		scenes.Scenes[scenes.Current_Scene_Id].Setup()
	}

	scenes.Scenes[scenes.Current_Scene_Id].Update()
}

var fps_toggle = false

func main() {
	rl.InitWindow(1920, 1080, "this is gonna go terribly")
	rl.SetConfigFlags(rl.FlagWindowResizable)

	networking.CrazyBloxPlayer = rl.LoadTexture("./art/player/crazyblox.png")
	networking.GregPlayer = rl.LoadTexture("./art/player/greg.png")

	networking.InitProjectileImages()

	scenes.CrosshairTexture = rl.LoadTexture("./art/ui/crosshair.png")

	level.InitBlockRendering()

	rl.DisableCursor()

	for !rl.WindowShouldClose() {
		update()
		draw()
	}

	rl.CloseWindow()
}
