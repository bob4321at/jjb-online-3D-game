package player

import (
	"main/networking"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Moveset interface {
	Move1(player *PlayerStruct, other_players networking.CollectionOfNetworkedPlayers)
	Move2(player *PlayerStruct, other_players networking.CollectionOfNetworkedPlayers)
	Move3(player *PlayerStruct, other_players networking.CollectionOfNetworkedPlayers)
	GetCooldowns() (float64, float64, float64)
	Update()
}

func ManageMoveset(player *PlayerStruct, moveset Moveset, other_players networking.CollectionOfNetworkedPlayers) {
	cooldown_1, cooldown_2, cooldown_3 := moveset.GetCooldowns()

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		if cooldown_1 <= 0 {
			moveset.Move1(player, other_players)
		}
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		if cooldown_2 <= 0 {
			moveset.Move2(player, other_players)
		}
	}
	if rl.IsKeyPressed(rl.KeyE) {
		if cooldown_3 <= 0 {
			moveset.Move3(player, other_players)
		}
	}
}
