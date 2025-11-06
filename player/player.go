package player

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/camera"
	"main/level"
	"main/networking"
	"main/utils"
	"math"
	"net/http"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerStruct struct {
	Pos       rl.Vector3
	CameraRot rl.Vector3
	Vel       rl.Vector3
	CameraRef *camera.CameraStruct
	Moves     Moveset
	Health    int
}

func (player *PlayerStruct) Update(level level.LevelStruct) {
	if player.Health <= 0 {
		player.Pos.Y = 0
		player.Vel.Y = 0
	}

	player.CameraRef.Pos = player.Pos

	player.CameraRef.Camera.Target = rl.NewVector3(
		player.Pos.X+float32(math.Cos(utils.DegToRad(float64(player.CameraRot.Y))))*float32(math.Cos(utils.DegToRad(float64(player.CameraRot.X)))),
		player.Pos.Y+float32(math.Sin(utils.DegToRad(float64(player.CameraRot.X)))),
		player.Pos.Z+float32(math.Sin(utils.DegToRad(float64(player.CameraRot.Y))))*float32(math.Cos(utils.DegToRad(float64(player.CameraRot.X)))),
	)

	if rl.IsKeyDown(rl.KeyJ) && player.CameraRot.X > -80 {
		player.CameraRot.X -= 300 * rl.GetFrameTime()
	} else if rl.IsKeyDown(rl.KeyK) && player.CameraRot.X < 80 {
		player.CameraRot.X += 300 * rl.GetFrameTime()
	}
	player.CameraRot.X -= rl.GetMouseDelta().Y / 10

	if rl.IsKeyDown(rl.KeyH) {
		player.CameraRot.Y -= 300 * rl.GetFrameTime()
	} else if rl.IsKeyDown(rl.KeyL) {
		player.CameraRot.Y += 300 * rl.GetFrameTime()
	}
	player.CameraRot.Y += rl.GetMouseDelta().X / 10

	player.Vel.Y -= 50 * rl.GetFrameTime()

	if networking.ServerName != "" {
		ManageMoveset(player, player.Moves, *networking.OtherPlayers)
	}

	if level.CheckCollision(rl.Vector3Add(player.Pos, rl.NewVector3(0, player.Vel.Y*rl.GetFrameTime(), 0)), rl.NewVector3(1, 2, 1)) {
		if player.Vel.Y < 0 {
			player.Vel.X *= 0.9 * rl.GetFrameTime()
			player.Vel.Z *= 0.9 * rl.GetFrameTime()

			if rl.IsKeyDown(rl.KeyW) {
				player.Vel.X += 0.5 * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.Y))))
				player.Vel.Z += 0.5 * float32(math.Sin(utils.DegToRad(float64(player.CameraRot.Y))))
			} else if rl.IsKeyDown(rl.KeyS) {
				player.Vel.X -= 0.5 * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.Y))))
				player.Vel.Z -= 0.5 * float32(math.Sin(utils.DegToRad(float64(player.CameraRot.Y))))
			}

			if rl.IsKeyDown(rl.KeyD) {
				player.Vel.X += 0.5 * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.Y)+90)))
				player.Vel.Z += 0.5 * float32(math.Sin(utils.DegToRad(float64(player.CameraRot.Y)+90)))
			} else if rl.IsKeyDown(rl.KeyA) {
				player.Vel.X -= 0.5 * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.Y)+90)))
				player.Vel.Z -= 0.5 * float32(math.Sin(utils.DegToRad(float64(player.CameraRot.Y)+90)))
			}

			if rl.IsKeyDown(rl.KeySpace) {
				player.Vel.Y = 15
			} else {
				player.Vel.Y = 0
			}
		} else {
			player.Vel.Y = 0
		}
	}

	if level.CheckCollision(rl.Vector3Add(player.Pos, rl.NewVector3(player.Vel.X*rl.GetFrameTime()*20, 0, 0)), rl.NewVector3(1, 2, 1)) {
		player.Vel.X = 0
	}
	if level.CheckCollision(rl.Vector3Add(player.Pos, rl.NewVector3(0, 0, player.Vel.Z*rl.GetFrameTime()*20)), rl.NewVector3(1, 2, 1)) {
		player.Vel.Z = 0
	}

	player.Pos.X += player.Vel.X * rl.GetFrameTime() * 20
	player.Pos.Y += player.Vel.Y * rl.GetFrameTime()
	player.Pos.Z += player.Vel.Z * rl.GetFrameTime() * 20
}

func (player *PlayerStruct) NetworkPlayer(ServerName string) *http.Response {
	if networking.ServerName == "" {
		return nil
	}

	player_bytes, err := json.Marshal(networking.NetworkedPlayer{Pos_X: player.Pos.X, Pos_Y: player.Pos.Y, Pos_Z: player.Pos.Z, Health: player.Health, ID: 0})
	if err != nil {
		panic(err)
	}
	resp, err := http.Post(ServerName+"/AddPlayer", "application/json", bytes.NewBuffer(player_bytes))
	if err != nil {
		panic(err)
	}

	get_player_id_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	get_player_id_temp := networking.NetworkedPlayer{}
	json.Unmarshal(get_player_id_bytes, &get_player_id_temp)

	networking.PlayerId = get_player_id_temp.ID

	// Damage Checks
	go func() {
		for true {
			this_player := networking.NetworkedPlayer{Pos_X: player.Pos.X, Pos_Y: player.Pos.Y, Pos_Z: player.Pos.Z, Health: player.Health, ID: networking.PlayerId}
			if err != nil {
				panic(err)
			}

			for i, projectile := range networking.Projectiles.Projectiles {
				if utils.Collision(rl.NewVector3(player.Pos.X, player.Pos.Y, player.Pos.Z), rl.NewVector3(1, 2, 1), rl.NewVector3(projectile.Pos_X, projectile.Pos_Y, projectile.Pos_Z), rl.NewVector3(1, 1, 1)) {
					damage_bytes, err := json.Marshal(networking.PlayerAndProjectileNetworked{Player: this_player, Projectile: projectile})
					if err != nil {
						panic(err)
					}
					http.Post(ServerName+"/DamagePlayer", "application/json", bytes.NewBuffer(damage_bytes))
					utils.RemoveArrayElement(i, &networking.Projectiles.Projectiles)
					break
				}
			}
		}
	}()

	for true {
		time.Sleep(time.Second / 15)

		this_player := networking.NetworkedPlayer{Pos_X: player.Pos.X, Pos_Y: player.Pos.Y, Pos_Z: player.Pos.Z, Health: player.Health, ID: networking.PlayerId}
		this_player_bytes, err := json.Marshal(this_player)
		if err != nil {
			panic(err)
		}

		resp, err = http.Post(ServerName+"/UpdatePlayerPos", "application/json", bytes.NewBuffer(this_player_bytes))
		if err != nil {
			panic(err)
		}

		resp, err = http.Post(ServerName+"/GetPlayerHealth", "application/json", bytes.NewBuffer(this_player_bytes))
		if err != nil {
			panic(err)
		}

		player_health_bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		player_health := networking.NetworkedPlayer{}
		json.Unmarshal(player_health_bytes, &player_health)

		player.Health = int(player_health.Health)

		fmt.Println(player.Health)
	}

	return resp
}

func NewPlayer(pos rl.Vector3) (player PlayerStruct) {
	player.Pos = pos
	player.Vel = rl.NewVector3(0, 0, 0)
	player.CameraRef = &camera.Camera
	player.Moves = &GregMoveset{}
	player.Health = 100

	return player
}

var Player = NewPlayer(rl.NewVector3(-5, 0, -5))
