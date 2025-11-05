package utils

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RemoveArrayElement[T any](index_to_remove int, slice *[]T) {
	*slice = append((*slice)[:index_to_remove], (*slice)[index_to_remove+1:]...)
}

func Distance(p1, p2 rl.Vector3) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dz := p2.Z - p1.Z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func Collision(origonal_pos, origonal_size, other, other_size rl.Vector3) bool {
	origonal_pos.X -= (origonal_size.X / 2)
	origonal_pos.Y -= (origonal_size.Y / 2)
	origonal_pos.Z -= (origonal_size.Z / 2)

	other.X -= (other_size.X / 2)
	other.Y -= (other_size.Y / 2)
	other.Z -= (other_size.Z / 2)

	if origonal_pos.X < other.X+other_size.X && origonal_pos.X+origonal_size.X > other.X {
		if origonal_pos.Y < other.Y+other_size.Y && origonal_pos.Y+origonal_size.Y > other.Y {
			if origonal_pos.Z < other.Z+other_size.Z && origonal_pos.Z+origonal_size.Z > other.Z {
				return true
			}
		}
	}

	origonal_pos.X += (origonal_size.X / 2)
	origonal_pos.Y += (origonal_size.Y / 2)
	origonal_pos.Z += (origonal_size.Z / 2)

	other.X += (other_size.X / 2)
	other.Y += (other_size.Y / 2)
	other.Z += (other_size.Z / 2)

	return false
}

func DegToRad(num float64) float64 {
	return num * (3.14159 / 180)
}

var Key_To_String = map[uint]string{
	rl.KeyQ:         "Q",
	rl.KeyW:         "W",
	rl.KeyE:         "E",
	rl.KeyR:         "R",
	rl.KeyT:         "T",
	rl.KeyY:         "Y",
	rl.KeyU:         "U",
	rl.KeyI:         "I",
	rl.KeyO:         "O",
	rl.KeyP:         "P",
	rl.KeyA:         "A",
	rl.KeyS:         "S",
	rl.KeyD:         "D",
	rl.KeyF:         "F",
	rl.KeyG:         "G",
	rl.KeyH:         "H",
	rl.KeyJ:         "J",
	rl.KeyK:         "K",
	rl.KeyL:         "L",
	rl.KeyZ:         "Z",
	rl.KeyX:         "X",
	rl.KeyC:         "C",
	rl.KeyV:         "V",
	rl.KeyB:         "B",
	rl.KeyN:         "N",
	rl.KeyM:         "M",
	rl.KeyOne:       "1",
	rl.KeyTwo:       "2",
	rl.KeyThree:     "3",
	rl.KeyFour:      "4",
	rl.KeyFive:      "5",
	rl.KeySix:       "6",
	rl.KeySeven:     "7",
	rl.KeyEight:     "8",
	rl.KeyNine:      "9",
	rl.KeyZero:      "0",
	rl.KeySpace:     " ",
	rl.KeyPeriod:    ".",
	rl.KeySemicolon: ":",
	rl.KeySlash:     "/",
}
