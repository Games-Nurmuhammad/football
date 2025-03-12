package structs

type Team struct {
	ClubName      string     `json:"clubname"`
	ClubLogo      string     `json:"clublogo"`
	ClubShortName string     `json:"clubshortname"`
	Players       [11]Player `json:"players"`
}
