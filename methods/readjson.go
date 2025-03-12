package methods

import (
	"encoding/json"
	"os"
	"viscabarca/structs"
)

func ReadJson() []structs.Team {
	var teams []structs.Team
	file, err := os.OpenFile("./jsons/team.json", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decor := json.NewDecoder(file)
	err = decor.Decode(&teams)
	if err != nil {
		panic(err)
	}
	return teams
}
