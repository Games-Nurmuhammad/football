package methods

import (
	"encoding/json"
	"os"
	"viscabarca/structs"
)

func TeamBarca() {
	a := []structs.Team{}
	barca := structs.Team{}

	barca.ClubName = "BARCELONA"
	barca.ClubShortName = "FCB"
	barca.ClubLogo = "🔴🔵"

	barca.Players[0].PlayerName = "Messi"
	barca.Players[1].PlayerName = "Neymar"
	barca.Players[2].PlayerName = "Suarez"
	barca.Players[3].PlayerName = "Lewandowski"
	barca.Players[4].PlayerName = "Iniesta"
	barca.Players[5].PlayerName = "Ronaldinho"
	barca.Players[6].PlayerName = "Xavi"
	barca.Players[7].PlayerName = "De Jong"
	barca.Players[8].PlayerName = "Pique"
	barca.Players[9].PlayerName = "Alves"
	barca.Players[10].PlayerName = "Ter Stegen"
	a = append(a, barca)
	Writejson(&a)

}

func TeamReal() {
	real := structs.Team{}

	real.ClubName = "REAL MADRID"
	real.ClubShortName = "RMA"
	real.ClubLogo = "⚪⚪"

	real.Players[0].PlayerName = "Ronaldo7"
	real.Players[1].PlayerName = "Ronaldo9"
	real.Players[2].PlayerName = "Benzema"
	real.Players[3].PlayerName = "Raul"
	real.Players[4].PlayerName = "Modric"
	real.Players[5].PlayerName = "Kross"
	real.Players[6].PlayerName = "Zidane"
	real.Players[7].PlayerName = "Ramos"
	real.Players[8].PlayerName = "Camavinga"
	real.Players[9].PlayerName = "Carvahal"
	real.Players[10].PlayerName = "Casilas"
	team1 := ReadJson()
	team1 = append(team1, real)
	Writejson(&team1)

}

func Writejson(team *[]structs.Team) {
	file, err := os.OpenFile("./jsons/team.json", os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(team)
	if err != nil {
		panic(err)
	}
}
