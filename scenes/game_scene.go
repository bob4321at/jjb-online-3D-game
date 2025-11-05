package scenes

import (
	"main/camera"
	"main/level"
	"main/networking"
	"main/player"

	"github.com/gen2brain/raylib-go/raylib"
)

type GameScene struct {
	SetedUp bool
	Editing bool
}

var CrosshairTexture rl.Texture2D

func (scene *GameScene) Setup() {
	networking.StartNetworking()
	go player.Player.NetworkPlayer(networking.ServerName)
	scene.SetedUp = true
}

func (scene *GameScene) Draw() {
	rl.BeginMode3D(camera.Camera.Camera)

	level.Level.Draw()
	if networking.ServerName != "" {
		networking.DrawOtherPlayers()
		networking.DrawProjectiles()
	}

	rl.EndMode3D()

	rl.DrawTexture(CrosshairTexture, 1920/2, 1080/2, rl.White)
}

func (scene *GameScene) Update() {
	if rl.IsKeyPressed(rl.KeyTab) {
		scene.Editing = !scene.Editing
	}

	if !scene.Editing {
		player.Player.Update(level.Level)
	} else {
		level.Level.Edit()
	}
	camera.Camera.Update()

	if networking.ServerName != "" {
		networking.InterpolatePlayers()
		networking.UpdateProjectiles()
	}
}

func (scene *GameScene) GetSetup() bool {
	return scene.SetedUp
}
