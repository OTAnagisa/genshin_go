package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UidInfoJson struct {
	PlayerInfo PlayerInfo `json:"playerInfo"`
	Ttl        int        `json:"ttl"`
	Uid        string     `json:"uid"`
}

type PlayerInfo struct {
	NickName                string           `json:"nickname"`
	Level                   int              `json:"level"`
	WorldLevel              int              `json:"worldLevel"`
	NameCardIcon            string           `json:"nameCardIcon"`
	FinfinishAchievementNum int              `json:"finishAchievementNum"`
	TowerFloorIndex         int              `json:"towerFloorIndex"`
	TowerLevelIndex         int              `json:"towerLevelIndex"`
	ShowAvatarInfoList      []ShowAvatarInfo `json:"showAvatarInfoList"`
	ShowNameCardIdList      []int            `json:"showNameCardIdList"`
	ProfilePicture          ProfilePicture   `json:"profilePicture"`
}

type ShowAvatarInfo struct {
	AvatarId int `json:"avatarId"`
	Level    int `json:"level"`
}

type ProfilePicture struct {
	AvatarId int `json:"avatarId"`
}

func PlayerSearchHandler(c *gin.Context) {
	uid := c.Param("uid")
	url := "https://enka.network/api/uid/" + uid
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 処理が終わったら必ずレスポンスを閉じる
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	body, _ := io.ReadAll(resp.Body)

	// fmt.Printf("%+v\n", resp)

	var uidInfoJson UidInfoJson

	if err := json.Unmarshal(body, &uidInfoJson); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", uidInfoJson)
}
