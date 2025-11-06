package networking

import (
	"bytes"
	"encoding/json"
	"io"
	"main/camera"
	"net/http"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type NetworkedProjectile struct {
	Pos_X      float32
	Pos_Y      float32
	Pos_Z      float32
	Vel_X      float32
	Vel_Y      float32
	Vel_Z      float32
	Speed      float32
	Damage     int
	Name       string
	ServerTime uint64
}

type CollectionOfNetworkedProjectiles struct {
	Projectiles []NetworkedProjectile
}

type PlayerAndProjectileNetworked struct {
	Player     NetworkedPlayer
	Projectile NetworkedProjectile
}

var Projectiles *CollectionOfNetworkedProjectiles

var ProjectileImages map[string]rl.Texture2D

func InitProjectileImages() {
	ProjectileImages = map[string]rl.Texture2D{
		"Greg Rock":    rl.LoadTexture("./art/projectiles/Greg Rock.png"),
		"Crazy Button": rl.LoadTexture("./art/projectiles/Crazy Button.png"),
	}
}

func SpawnProjectile(projectile NetworkedProjectile) {
	if ServerName == "" {
		return
	}

	time := GetServerData().Time

	projectile.ServerTime = time

	bytes_to_send, err := json.Marshal(projectile)
	if err != nil {
		panic(err)
	}

	if _, err := http.Post(ServerName+"/SpawnProjectile", "application/json", bytes.NewBuffer(bytes_to_send)); err != nil {
		panic(err)
	}
}

func UpdateProjectiles() {
	for i := range Projectiles.Projectiles {
		projectile := &Projectiles.Projectiles[i]
		projectile.Pos_X += projectile.Vel_X * rl.GetFrameTime() * 60
		projectile.Pos_Y += projectile.Vel_Y * rl.GetFrameTime() * 60
		projectile.Pos_Z += projectile.Vel_Z * rl.GetFrameTime() * 60
	}
}

func DrawProjectiles() {
	if Projectiles != nil {
		for _, projectile := range Projectiles.Projectiles {
			rl.DrawBillboard(camera.Camera.Camera, ProjectileImages[projectile.Name], rl.NewVector3(projectile.Pos_X, projectile.Pos_Y, projectile.Pos_Z), 1, rl.White)
		}
	}
}

func ProjectileNetworking() {
	resp, err := http.Get(ServerName + "/GetProjectiles")
	if err != nil {
		panic(err)
	}

	projectile_data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(projectile_data, Projectiles)
}
