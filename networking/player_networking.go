package networking

import (
	"bytes"
	"encoding/json"
	"io"
	"main/camera"
	"math"
	"net/http"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type NetworkedPlayer struct {
	Pos_X  float32
	Pos_Y  float32
	Pos_Z  float32
	Health uint8
	ID     uint8
}

type CollectionOfNetworkedPlayers struct {
	Players []NetworkedPlayer
}

var ServerName string = "http://localhost:8080"

var PlayerId uint8

var OtherPlayers *CollectionOfNetworkedPlayers
var OtherPlayersInterpelated map[uint8]*NetworkedPlayer

var CrazyBloxPlayer rl.Texture2D
var GregPlayer rl.Texture2D

func DrawOtherPlayers() {
	for _, player := range OtherPlayersInterpelated {
		rl.DrawBillboard(camera.Camera.Camera, GregPlayer, rl.NewVector3(player.Pos_X, player.Pos_Y, player.Pos_Z), 2, rl.White)
	}
}

func InterpolatePlayers() {
	for _, other_player := range OtherPlayers.Players {
		other_player_old_pos := rl.NewVector3(
			OtherPlayersInterpelated[other_player.ID].Pos_X,
			OtherPlayersInterpelated[other_player.ID].Pos_Y,
			OtherPlayersInterpelated[other_player.ID].Pos_Z,
		)
		Move_Dir := rl.Vector3Subtract(
			rl.NewVector3(other_player.Pos_X, other_player.Pos_Y, other_player.Pos_Z),
			other_player_old_pos,
		)

		Move_Dir.X *= 10
		Move_Dir.X *= rl.GetFrameTime()
		Move_Dir.Y *= 10
		Move_Dir.Y *= rl.GetFrameTime()
		Move_Dir.Z *= 10
		Move_Dir.Z *= rl.GetFrameTime()

		OtherPlayersInterpelated[other_player.ID].Pos_X += Move_Dir.X
		OtherPlayersInterpelated[other_player.ID].Pos_Y += Move_Dir.Y
		OtherPlayersInterpelated[other_player.ID].Pos_Z += Move_Dir.Z
	}
}

func OtherPlayerNetworking() {
	this_player := NetworkedPlayer{0, 0, 0, 0, PlayerId}
	this_player_bytes, err := json.Marshal(this_player)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(ServerName+"/GetOtherPlayers", "application/json", bytes.NewBuffer(this_player_bytes))
	if err != nil {
		panic(err)
	}

	other_player_data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(other_player_data, OtherPlayers)
	for _, player := range OtherPlayers.Players {
		_, exists := OtherPlayersInterpelated[player.ID]

		if !exists {
			OtherPlayersInterpelated[player.ID] = &player
		}

		if math.Abs(float64(player.Pos_X-OtherPlayersInterpelated[player.ID].Pos_X)) > 10 {
			OtherPlayersInterpelated[player.ID] = &player
		} else if math.Abs(float64(player.Pos_Y-OtherPlayersInterpelated[player.ID].Pos_Y)) > 10 {
			OtherPlayersInterpelated[player.ID] = &player

		} else if math.Abs(float64(player.Pos_Z-OtherPlayersInterpelated[player.ID].Pos_Z)) > 10 {
			OtherPlayersInterpelated[player.ID] = &player
		}
	}
}
