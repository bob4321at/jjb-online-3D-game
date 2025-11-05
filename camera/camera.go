package camera

import rl "github.com/gen2brain/raylib-go/raylib"

type CameraStruct struct {
	Pos    rl.Vector3
	Rot    rl.Vector3
	Camera rl.Camera3D
}

func (camera *CameraStruct) Update() {
	camera.Camera.Position = camera.Pos
	if camera.Rot != rl.Vector3Zero() {
		camera.Camera.Target = camera.Rot
	}
}

func NewCamera(pos rl.Vector3, rot rl.Vector3) (camera CameraStruct) {
	camera.Pos = pos
	camera.Rot = rot
	camera.Camera = rl.NewCamera3D(camera.Pos, rl.NewVector3(0, 0, 0), rl.NewVector3(0, 1, 0), 66, rl.CameraPerspective)

	return camera
}

var Camera = NewCamera(rl.NewVector3(0, 1, 0), rl.NewVector3(0, 0, 0))
