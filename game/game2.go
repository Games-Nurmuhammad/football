package game

import (
	"fmt"
	"math/rand"
	"viscabarca/methods"
)

func Game2(countteam1 *int, countteam2 *int, countname int) {
	teams := methods.ReadJson()

	str1, str2 := "", ""

	plchoise := 0

	for i := 0; true; i++ {
		score1, score2 := false, false
		// boshlanishi
		if countname == 11 {
			countname = 0
		}

		fmt.Printf("%s %s\n", teams[0].ClubShortName, str1)
		fmt.Printf("%s %s\n", teams[1].ClubShortName, str2)

		for i := 0; i <= 20; i++ {
			for j := 0; j <= 20; j++ {
				if i == 0 {
					fmt.Print("===")
				} else if (j == 0 || j == 20) && i <= 10 {
					fmt.Print("||  ")
				} else if j == 10 && i == 6 {
					fmt.Print("O  ")
				} else if j == 9 && i == 7 {
					fmt.Print("🧤/|\\🧤")
				} else if (j == 10 || j == 11) && i == 7 {
					fmt.Print(" ")
				} else if j == 9 && i == 8 {
					fmt.Print("  / \\")
				} else if j == 10 && i == 8 {
					fmt.Print(" ")
				} else if i == 15 && j == 10 {
					fmt.Print("⚽️")
				} else if i == 16 && j == 5 {
					fmt.Print("  O")
				} else if i == 17 && j == 5 {
					fmt.Print("-/|\\-")

				} else if i == 18 && j == 5 {

					fmt.Print(" / \\")
				} else {
					fmt.Print("   ")
				}
			}
			fmt.Println()
		}

		// check qilishi

		gkrand := rand.Intn(3) + 1

		fmt.Println("1.chapga\n2.O'rtaga\n3.O'nga\n ")
		fmt.Printf("%s ning navbati to'p oldida %s>>> ", teams[i%2].ClubName, teams[i%2].Players[countname].PlayerName)
		fmt.Scan(&plchoise)
		if plchoise == gkrand {
			fmt.Println("  _   _  ____     _____  ____          _      ")
			fmt.Println(" | \\ | |/ __ \\   / ____|/ __ \\   /\\   | |     ")
			fmt.Println(" |  \\| | |  | | | |  __| |  | | /  \\  | |     ")
			fmt.Println(" | . ` | |  | | | | |_ | |  | |/ /\\ \\ | |     ")
			fmt.Println(" | |\\  | |__| | | |__| | |__| / ____ \\| |____ ")
			fmt.Println(" |_| \\_|\\____/   \\_____|\\____/_/    \\_\\______|")
			fmt.Println("                                               ")
			fmt.Println("                                               ")
			if i%2 == 0 {
				score1 = false
				changeState(score1, &str1)
			} else {
				score2 = false
				changeState(score2, &str2)
			}
		} else {
			fmt.Println(" _______  _______  _______  _        ")
			fmt.Println("(  ____ \\(  ___  )(  ___  )( \\       ")
			fmt.Println("| (    \\/| (   ) || (   ) || (       ")
			fmt.Println("| |      | |   | || (___) || |       ")
			fmt.Println("| | ____ | |   | ||  ___  || |       ")
			fmt.Println("| | \\_  )| |   | || (   ) || |       ")
			fmt.Println("| (___) || (___) || )   ( || (____/\\")
			fmt.Println("(_______)(_______)|/     \\|(_______/")
			if i%2 == 0 {
				*countteam1++
				score1 = true
				changeState(score1, &str1)
			} else {
				*countteam2++
				score2 = true
				changeState(score2, &str2)
			}

		}
		Picture(plchoise, gkrand)
		if i%2 == 1 {
			countname++
			if *countteam1 != *countteam2 {
				break
			}
		}
	}

}

func changeState(score bool, str *string) {
	if score {
		*str += "🟢"
	} else {
		*str += "🔴"
	}

}
