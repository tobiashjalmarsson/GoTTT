package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type Player struct {
	Name   string
	Symbol string
}

func main() {
	players := createPlayers()
	var currentPlayer int = 0 // the index of the current player thats playing in the array

	grid := [5][5]string{
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
	}

	var gp *[5][5]string
	gp = &grid

	clearScreen()
	displayGrid(gp)

	for {
		validInput := getUserInput(gp, &players[currentPlayer])

		clearScreen()
		displayGrid(gp)
		// If the game has ended, we exit
		if checkWin(gp, &players[currentPlayer]) {
			fmt.Printf("%s won the game", players[currentPlayer].Name)
			break
		}
		// Update the current Player
		if validInput {
			currentPlayer = 1 - currentPlayer
		} else {
			fmt.Println("Please choose a valid position")
		}

	}

}

func createPlayers() [2]Player {
	var players [2]Player
	players[0] = createPlayer("X")
	players[1] = createPlayer("+")
	return players
}

func createPlayer(symbol string) Player {
	clearScreen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name:")
	name, _ := reader.ReadString('\n')
	return Player{name, symbol}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func displayGrid(grid *[5][5]string) {
	for idx, value := range *grid {
		fmt.Println(idx+1, value)
	}
}

func getUserInput(grid *[5][5]string, player *Player) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Row and Column number 21 for column 2 and row 1:")
	text, _ := reader.ReadString('\n')
	firstIndex, _ := strconv.Atoi(text[0:1])
	secondIndex, _ := strconv.Atoi(text[1:2])

	if firstIndex > 5 && secondIndex > 5 {
		return false
	}

	if grid[firstIndex-1][secondIndex-1] == "0" {
		grid[firstIndex-1][secondIndex-1] = player.Symbol
		return true
	}
	return false
}

func checkWin(grid *[5][5]string, player *Player) bool {
	//count := 0
	// CHECK ROW 	and grid[i][j+1] and grid[i][j+2]
	for idx, value := range *grid {
		for idx2, value2 := range value {
			if value2 == player.Symbol && (idx2 < len(value)-2) {
				// Check the row
				if value[idx2+1] == player.Symbol && value[idx2+2] == player.Symbol {
					return true
				}
			}
			if value2 == player.Symbol && (idx < len(value)-2) {
				if grid[idx+1][idx2] == player.Symbol && grid[idx+2][idx2] == player.Symbol {
					return true
				}
			}
			if value2 == player.Symbol && (idx < len(value)-2) && (idx2 < len(value)-2) {
				if grid[idx+1][idx2+1] == player.Symbol && grid[idx+2][idx2+2] == player.Symbol {
					return true
				}
			}

		}
	}
	return false
}
