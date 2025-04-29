package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var CURRENT_PLAYER string = "O"
var board [3][3]string

type Buttons struct {
	Grid [3][3]*widget.Button
}

func resetButtons(buttons *Buttons) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			buttons.Grid[i][j].SetText("")
		}
	}
}

func resetBoard() {
	for i := range board {
		for j := range board[i] {
			board[i][j] = ""
		}
	}
}

func showWinnerAlert(window fyne.Window, winner string, buttons *Buttons) {
	message := fmt.Sprintf("Game over — %s won! Play again?", winner)
	dialog.NewCustomConfirm("Game Over", "Replay", "Close",
		widget.NewLabel(message),
		func(replay bool) {
			if replay {
				CURRENT_PLAYER = "O"
				resetButtons(buttons)
				resetBoard()
			} else {
				window.Close()
			}
		},
		window).Show()
}

func isGameOver() (bool, string) {
	// Check rows and columns
	for i := range 3 {
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true, board[i][0]
		}
		if board[0][i] != "" && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return true, board[0][i]
		}
	}
	// Check diagonals
	if board[0][0] != "" && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true, board[0][0]
	}
	if board[0][2] != "" && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true, board[0][2]
	}
	return false, ""
}

func updateButton(button *widget.Button, row, col int, window fyne.Window, buttons *Buttons) {
	if board[row][col] != "" {
		fmt.Println("Invalid move! Cell already occupied.")
		return
	}
	button.SetText(CURRENT_PLAYER)
	board[row][col] = CURRENT_PLAYER

	if over, winner := isGameOver(); over {
		showWinnerAlert(window, winner, buttons)
		return
	}

	if CURRENT_PLAYER == "X" {
		CURRENT_PLAYER = "O"
	} else {
		CURRENT_PLAYER = "X"
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TikTakToe")

	buttons := &Buttons{}
	gridItems := []fyne.CanvasObject{}

	for i := range 3 {
		for j := 0; j < 3; j++ {
			i, j := i, j // capture loop vars
			btn := widget.NewButton("", func() {
				updateButton(buttons.Grid[i][j], i, j, myWindow, buttons)
			})
			btn.Resize(fyne.NewSize(300, 300))
			buttons.Grid[i][j] = btn
			gridItems = append(gridItems, btn)
		}
	}

	grid := container.NewGridWithColumns(3, gridItems...)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.SetContent(grid)
	myWindow.ShowAndRun()
}
