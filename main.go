package main

import (
	"fmt"
	"time"
	"viscabarca/game"
	"viscabarca/methods"
)

func main() {
	drawGate()
}

func drawGate() {

	methods.ClearTerminal()
	methods.Delete()
	fmt.Println("\t\tHELLO EVERYONE")
	time.Sleep(2 * time.Second)
	fmt.Println("\t\tWELCOME TO UCL 2024 FINAL MATCH")
	time.Sleep(2 * time.Second)
	methods.TeamBarca()
	methods.TeamReal()
	teams := methods.ReadJson()
	fmt.Printf("\t\t%s%s VS %s%s", teams[0].ClubName, teams[0].ClubLogo, teams[1].ClubName, teams[1].ClubLogo)
	time.Sleep(3 * time.Second)
	methods.ClearTerminal()

	fmt.Printf("%s %s FUTBOLCHILARI\n", teams[0].ClubName, teams[0].ClubLogo)
	for i := 0; i < 11; i++ {
		fmt.Println(i+1, ":", teams[0].Players[i].PlayerName)
	}

	time.Sleep(3 * time.Second)
	methods.ClearTerminal()
	fmt.Printf("%s %s FUTBOLCHILARI\n", teams[1].ClubName, teams[1].ClubLogo)
	for i := 0; i < 11; i++ {
		fmt.Println(i+1, ":", teams[1].Players[i].PlayerName)
	}

	time.Sleep(3 * time.Second)

	methods.ClearTerminal()

	count1, count2 := game.Start()
	if count1 > count2 {
		fmt.Printf("WINNER IS %s %s", teams[0].ClubName, teams[0].ClubLogo)
	} else {
		fmt.Printf("WINNER IS %s %s", teams[1].ClubName, teams[1].ClubLogo)
	}

}
