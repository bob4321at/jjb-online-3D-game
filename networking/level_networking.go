package networking

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type NetworkedBlockStruct struct {
	Pos_X  float32
	Pos_Y  float32
	Pos_Z  float32
	Size_X float32
	Size_Y float32
	Size_Z float32
	Color  uint8
}

type NetworkedLevel struct {
	Blocks []NetworkedBlockStruct
}

func SendNewLevel(new_level NetworkedLevel) {
	level_bytes, err := json.Marshal(new_level)
	if err != nil {
		panic(err)
	}
	http.Post(ServerName+"/SendLevel", "application/json", bytes.NewBuffer(level_bytes))
}

func GetNewLevel() NetworkedLevel {
	resp, err := http.Get(ServerName + "/GetLevel")
	if err != nil {
		panic(err)
	}

	new_level_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var new_level NetworkedLevel
	json.Unmarshal(new_level_bytes, &new_level)

	return new_level
}
