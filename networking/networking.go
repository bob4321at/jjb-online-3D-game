package networking

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type ServerDataStruct struct {
	Time uint64
}

func GetServerData() (server_data ServerDataStruct) {
	resp, err := http.Get(ServerName + "/GetServerData")
	if err != nil {
		panic(err)
	}

	time_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(time_bytes, &server_data)

	return server_data
}

func StartNetworking() {
	if ServerName == "" {
		return
	}

	other_players := CollectionOfNetworkedPlayers{}
	OtherPlayers = &other_players
	OtherPlayersInterpelated = make(map[uint8]*NetworkedPlayer)

	projectiles := CollectionOfNetworkedProjectiles{}
	Projectiles = &projectiles

	go func() {
		for true {
			time.Sleep(time.Second / 10)

			OtherPlayerNetworking()
			ProjectileNetworking()
		}
	}()
}
