package player

import (
	"main/networking"
	"main/utils"
	"math"
)

type GregMoveset struct{}

func (moveset *GregMoveset) Move1(player *PlayerStruct, other_players networking.CollectionOfNetworkedPlayers) {
	player.Vel.Y += 1 * float32(math.Sin(utils.DegToRad(float64(player.CameraRot.X)))) * 50
	player.Vel.X += 1 * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.Y)))) * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.X))))
	player.Vel.Z += 1 * float32(math.Sin(utils.DegToRad(float64(player.CameraRot.Y)))) * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.X))))
}

func (moveset *GregMoveset) Move2(player *PlayerStruct, other_players networking.CollectionOfNetworkedPlayers) {
	projectile := networking.NetworkedProjectile{
		Pos_X:  player.Pos.X,
		Pos_Y:  player.Pos.Y,
		Pos_Z:  player.Pos.Z,
		Vel_X:  float32(math.Cos(utils.DegToRad(float64(player.CameraRot.Y)))) * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.X)))),
		Vel_Y:  float32(math.Sin(utils.DegToRad(float64(player.CameraRot.X)))),
		Vel_Z:  float32(math.Sin(utils.DegToRad(float64(player.CameraRot.Y)))) * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.X)))),
		Speed:  1,
		Damage: 5,
		Name:   "Greg Rock",
	}

	networking.SpawnProjectile(projectile)
}

func (moveset *GregMoveset) Move3(player *PlayerStruct, other_players networking.CollectionOfNetworkedPlayers) {
	projectile := networking.NetworkedProjectile{
		Pos_X:  player.Pos.X,
		Pos_Y:  player.Pos.Y,
		Pos_Z:  player.Pos.Z,
		Vel_X:  float32(math.Cos(utils.DegToRad(float64(player.CameraRot.Y)))) * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.X)))),
		Vel_Y:  float32(math.Sin(utils.DegToRad(float64(player.CameraRot.X)))),
		Vel_Z:  float32(math.Sin(utils.DegToRad(float64(player.CameraRot.Y)))) * float32(math.Cos(utils.DegToRad(float64(player.CameraRot.X)))),
		Speed:  2,
		Damage: 1,
		Name:   "Crazy Button",
	}

	networking.SpawnProjectile(projectile)
}

func (moveset *GregMoveset) GetCooldowns() (float64, float64, float64) {
	return 0, 0, 0
}

func (moveset *GregMoveset) Update() {

}
