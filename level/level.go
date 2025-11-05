package level

import (
	"encoding/json"
	"image/color"
	"main/camera"
	"main/utils"
	"math"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var BlockTextures map[uint8]rl.Texture2D
var BlockModel rl.Model

func InitBlockRendering() {
	BlockModel = rl.LoadModel("./art/Tiles/Grass.obj")

	BlockTextures = map[uint8]rl.Texture2D{
		0: rl.LoadTexture("./art/Tiles/Grass.png"),
		1: rl.LoadTexture("./art/Tiles/Grass.png"),
		2: rl.LoadTexture("./art/Tiles/Grass.png"),
		3: rl.LoadTexture("./art/Tiles/Grass.png"),
		4: rl.LoadTexture("./art/Tiles/Grass.png"),
		5: rl.LoadTexture("./art/Tiles/Grass.png"),
	}
}

type BlockStruct struct {
	Pos   rl.Vector3
	Size  rl.Vector3
	Color uint8
}

func (block *BlockStruct) CheckCollision(object_pos, object_size rl.Vector3) bool {
	return utils.Collision(block.Pos, block.Size, object_pos, object_size)
}

type LevelStruct struct {
	Blocks         []BlockStruct
	CameraPos      rl.Vector3
	CameraRot      rl.Vector3
	CameraRef      *camera.CameraStruct
	Selected_Block *BlockStruct
}

func (level *LevelStruct) Edit() {
	level.CameraRef.Pos = level.CameraPos

	level.CameraRef.Camera.Target = rl.NewVector3(
		level.CameraPos.X+float32(math.Cos(utils.DegToRad(float64(level.CameraRot.Y))))*float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X)))),
		level.CameraPos.Y+float32(math.Sin(utils.DegToRad(float64(level.CameraRot.X)))),
		level.CameraPos.Z+float32(math.Sin(utils.DegToRad(float64(level.CameraRot.Y))))*float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X)))),
	)

	level.CameraRot.X -= rl.GetMouseDelta().Y / 10

	level.CameraRot.Y += rl.GetMouseDelta().X / 10

	if rl.IsKeyDown(rl.KeyW) {
		level.CameraPos.X += float32(math.Cos(utils.DegToRad(float64(level.CameraRot.Y)))) * float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X)))) * rl.GetFrameTime() * 60
		level.CameraPos.Y += float32(math.Sin(utils.DegToRad(float64(level.CameraRot.X)))) * rl.GetFrameTime() * 60
		level.CameraPos.Z += float32(math.Sin(utils.DegToRad(float64(level.CameraRot.Y)))) * float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X)))) * rl.GetFrameTime() * 60
	} else if rl.IsKeyDown(rl.KeyS) && !rl.IsKeyDown(rl.KeyLeftControl) {
		level.CameraPos.X -= float32(math.Cos(utils.DegToRad(float64(level.CameraRot.Y)))) * float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X)))) * rl.GetFrameTime() * 60
		level.CameraPos.Y -= float32(math.Sin(utils.DegToRad(float64(level.CameraRot.X)))) * rl.GetFrameTime() * 60
		level.CameraPos.Z -= float32(math.Sin(utils.DegToRad(float64(level.CameraRot.Y)))) * float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X)))) * rl.GetFrameTime() * 60
	}
	if rl.IsKeyDown(rl.KeyD) {
		level.CameraPos.X += float32(math.Cos(utils.DegToRad(float64(level.CameraRot.Y+90)))) * float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X)))) * rl.GetFrameTime() * 60
		level.CameraPos.Z += float32(math.Sin(utils.DegToRad(float64(level.CameraRot.Y+90)))) * float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X)))) * rl.GetFrameTime() * 60
	} else if rl.IsKeyDown(rl.KeyA) {
		level.CameraPos.X -= float32(math.Cos(utils.DegToRad(float64(level.CameraRot.Y+90)))) * float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X)))) * rl.GetFrameTime() * 60
		level.CameraPos.Z -= float32(math.Sin(utils.DegToRad(float64(level.CameraRot.Y+90)))) * float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X)))) * rl.GetFrameTime() * 60
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		level.Blocks = append(level.Blocks, BlockStruct{level.CameraRef.Camera.Target, rl.NewVector3(1, 1, 1), 0})
	}

	if level.Selected_Block != nil {
		if rl.IsKeyDown(rl.KeyI) && !rl.IsKeyDown(rl.KeyLeftShift) {
			level.Selected_Block.Size.X += 1 * rl.GetFrameTime()
		} else if rl.IsKeyDown(rl.KeyI) && rl.IsKeyDown(rl.KeyLeftShift) {
			level.Selected_Block.Size.X -= 1 * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyO) && !rl.IsKeyDown(rl.KeyLeftShift) {
			level.Selected_Block.Size.Y += 1 * rl.GetFrameTime()
		} else if rl.IsKeyDown(rl.KeyO) && rl.IsKeyDown(rl.KeyLeftShift) {
			level.Selected_Block.Size.Y -= 1 * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyP) && !rl.IsKeyDown(rl.KeyLeftShift) {
			level.Selected_Block.Size.Z += 1 * rl.GetFrameTime()
		} else if rl.IsKeyDown(rl.KeyP) && rl.IsKeyDown(rl.KeyLeftShift) {
			level.Selected_Block.Size.Z -= 1 * rl.GetFrameTime()
		}

		if rl.IsKeyDown(rl.KeyJ) && !rl.IsKeyDown(rl.KeyLeftShift) {
			level.Selected_Block.Pos.X += 1 * rl.GetFrameTime()
		} else if rl.IsKeyDown(rl.KeyJ) && rl.IsKeyDown(rl.KeyLeftShift) {
			level.Selected_Block.Pos.X -= 1 * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyK) && !rl.IsKeyDown(rl.KeyLeftShift) {
			level.Selected_Block.Pos.Y += 1 * rl.GetFrameTime()
		} else if rl.IsKeyDown(rl.KeyK) && rl.IsKeyDown(rl.KeyLeftShift) {
			level.Selected_Block.Pos.Y -= 1 * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyL) && !rl.IsKeyDown(rl.KeyLeftShift) && !rl.IsKeyDown(rl.KeyLeftControl) {
			level.Selected_Block.Pos.Z += 1 * rl.GetFrameTime()
		} else if rl.IsKeyDown(rl.KeyL) && rl.IsKeyDown(rl.KeyLeftShift) && !rl.IsKeyDown(rl.KeyLeftControl) {
			level.Selected_Block.Pos.Z -= 1 * rl.GetFrameTime()
		}
	}

	if rl.IsKeyPressed(rl.KeyS) && rl.IsKeyDown(rl.KeyLeftControl) {
		map_save, err := json.Marshal(level.Blocks)
		if err != nil {
			panic(err)
		}

		f, err := os.Create("./level_save")
		if err != nil {
			panic(err)
		}
		f.Write(map_save)
	}

	if rl.IsKeyPressed(rl.KeyL) && rl.IsKeyDown(rl.KeyLeftControl) {
		level_save_data, err := os.ReadFile("./level_save")
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal(level_save_data, &level.Blocks); err != nil {
			panic(err)
		}
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		go func() {
			for length := range 1000 {
				l := float64(length / 10)
				Ray_Dir := rl.NewVector3(
					float32(math.Cos(utils.DegToRad(float64(level.CameraRot.Y))))*float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X))))*float32(l),
					float32(math.Sin(utils.DegToRad(float64(level.CameraRot.X))))*float32(l),
					float32(math.Sin(utils.DegToRad(float64(level.CameraRot.Y))))*float32(math.Cos(utils.DegToRad(float64(level.CameraRot.X))))*float32(l),
				)

				for i := range level.Blocks {
					block := &level.Blocks[i]
					if block.CheckCollision(rl.Vector3Add(level.CameraPos, Ray_Dir), rl.NewVector3(0.1, 0.1, 0.1)) {
						level.Selected_Block = block
						return
					}
				}
			}
		}()
	}

	if level.Selected_Block != nil {
		if rl.IsKeyPressed(rl.KeyDelete) {
			for i := range level.Blocks {
				if &level.Blocks[i] == level.Selected_Block {
					utils.RemoveArrayElement(i, &level.Blocks)
				}
			}
		}
	}
}

func (level *LevelStruct) Draw() {
	for _, block := range level.Blocks {
		rl.SetMaterialTexture(&BlockModel.GetMaterials()[0], rl.MapDiffuse, BlockTextures[block.Color])
		rl.DrawModelEx(BlockModel, block.Pos, rl.Vector3Zero(), 0, rl.Vector3Divide(block.Size, rl.NewVector3(2, 2, 2)), rl.White)
	}

	if level.Selected_Block != nil {
		rl.DrawCubeWiresV(level.Selected_Block.Pos, level.Selected_Block.Size, color.RGBA{0, 0, 0, 255})
	}
}

func (level *LevelStruct) CheckCollision(object_pos, object_size rl.Vector3) bool {
	for _, block := range level.Blocks {
		if block.CheckCollision(object_pos, object_size) {
			return true
		}
	}
	return false
}

func NewLevel() (level LevelStruct) {
	level.Blocks = []BlockStruct{
		{Pos: rl.NewVector3(0, -10, 0), Size: rl.NewVector3(100, 1, 100), Color: 0},
	}

	level.CameraRef = &camera.Camera

	return level
}

var Level = NewLevel()
