package header

import (
	bitreader "github.com/luispmenezes/battle-replays/internal"
	"github.com/luispmenezes/battle-replays/pkg/utils"
)

type UserData struct {
	UserIdObj   UserId
	Username    string
	UserId      uint64
	Team        int
	B           bool
	C           bool
	ChampionObj GameObjectType
	Champion    string
}

func deserializeUserData(reader *bitreader.NetBuffer) UserData {
	userData := UserData{
		UserIdObj:   deserializeUserId(reader),
		Username:    reader.ReadString(),
		UserId:      reader.ReadUInt64(),
		Team:        int(reader.ReadByte()),
		B:           reader.ReadBoolean(),
		C:           reader.ReadBoolean(),
		ChampionObj: deserializeGameObjectType(reader),
	}

	userData.Champion = utils.GetChampionById(int(userData.ChampionObj.Id))

	return userData
}
