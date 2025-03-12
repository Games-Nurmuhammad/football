package methods

import "os"

func Delete() {
	file, err := os.OpenFile("./jsons/team.json", os.O_TRUNC, 06666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
