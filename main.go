package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "image/color"
)

var CURRENT_PLAYER string = "O"

var board [3][3]string

func updateButton(button *widget.Button, firstIndex uint8, secondIndex uint8, window *fyne.Window, buttons *Buttons) {
	// Check if the game is over
	isOver, winner := isGameOver()

	if !isOver {
		// Update the button text with the current player's symbol
		if CURRENT_PLAYER == "X" {
			button.SetText(CURRENT_PLAYER)
			CURRENT_PLAYER = "O"
		} else if CURRENT_PLAYER == "O" {
			button.SetText(CURRENT_PLAYER)
			CURRENT_PLAYER = "X"
		}

		// Update the board with the current player's symbol
		board[firstIndex][secondIndex] = CURRENT_PLAYER

		// Check if the game is over after this move
		isOver, winner = isGameOver()

		if isOver {
			// Display the winner (use fmt for console output or display an alert)
			fmt.Printf("Game Over! %s won!\n", winner)

			// Optionally show a dialog in the Fyne window
			showWinnerAlert(window, CURRENT_PLAYER)
			resetButtons(buttons)
		}
	} else {
		// If the game is already over, show the winner again
		fmt.Printf("Game Over! %s won!\n", winner)

		// Optionally show a dialog in the Fyne window
		showWinnerAlert(window, CURRENT_PLAYER)
		resetButtons(buttons)
	}
}

func resetButtons(buttons *Buttons) {
	buttons = &Buttons{
		Button11: widget.NewButton("", nil),
		Button12: widget.NewButton("", nil),
		Button13: widget.NewButton("", nil),
		Button21: widget.NewButton("", nil),
		Button22: widget.NewButton("", nil),
		Button23: widget.NewButton("", nil),
		Button31: widget.NewButton("", nil),
		Button32: widget.NewButton("", nil),
		Button33: widget.NewButton("", nil),
	}
}

// Function to display the winner in a dialog
func showWinnerAlert(window *fyne.Window, winner string) {
	message := fmt.Sprintf("🎮 Game over — %s won! Play again?", winner)

	// Custom Confirm dialog with "Replay" and "Close"
	dialog.NewCustomConfirm("Game Over", "Replay", "Close",
		widget.NewLabel(message),
		func(replay bool) {
			if replay {
				fmt.Println("Restarting game...")
				CURRENT_PLAYER = "O"

				for i := range board {
					for j := range i {
						board[i][j] = ""
					}
				}
			} else {
				fmt.Println("Closing game...") // Replace with exit or cleanup
			}
		},
		*window).Show()
}

func isGameOver() (bool, string) {
	//Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] != "" && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true, board[i][0]
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
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

type Buttons struct {
	Button11 *widget.Button
	Button12 *widget.Button
	Button13 *widget.Button
	Button21 *widget.Button
	Button22 *widget.Button
	Button23 *widget.Button
	Button31 *widget.Button
	Button32 *widget.Button
	Button33 *widget.Button
}

func main() {
	// initialize the app first and window
	app := app.New()
	window := app.NewWindow("TikTakToe")

	buttons := Buttons{
		Button11: widget.NewButton("", nil),
		Button12: widget.NewButton("", nil),
		Button13: widget.NewButton("", nil),
		Button21: widget.NewButton("", nil),
		Button22: widget.NewButton("", nil),
		Button23: widget.NewButton("", nil),
		Button31: widget.NewButton("", nil),
		Button32: widget.NewButton("", nil),
		Button33: widget.NewButton("", nil),
	}

	// Step 2: Now attach functions
	buttons.Button11.OnTapped = func() { updateButton(buttons.Button11, 0, 0, &window, &buttons) }
	buttons.Button12.OnTapped = func() { updateButton(buttons.Button12, 0, 1, &window, &buttons) }
	buttons.Button13.OnTapped = func() { updateButton(buttons.Button13, 0, 2, &window, &buttons) }
	buttons.Button21.OnTapped = func() { updateButton(buttons.Button21, 1, 0, &window, &buttons) }
	buttons.Button22.OnTapped = func() { updateButton(buttons.Button22, 1, 1, &window, &buttons) }
	buttons.Button23.OnTapped = func() { updateButton(buttons.Button23, 1, 2, &window, &buttons) }
	buttons.Button31.OnTapped = func() { updateButton(buttons.Button31, 2, 0, &window, &buttons) }
	buttons.Button32.OnTapped = func() { updateButton(buttons.Button32, 2, 1, &window, &buttons) }
	buttons.Button33.OnTapped = func() { updateButton(buttons.Button33, 2, 2, &window, &buttons) }

	// Arrange the buttons in a grid
	layout := container.NewGridWithColumns(3,
		buttons.Button11, buttons.Button12, buttons.Button13,
		buttons.Button21, buttons.Button22, buttons.Button23,
		buttons.Button31, buttons.Button32, buttons.Button33,
	)

	// resize the windows to a suitable size
	window.Resize(fyne.NewSize(500, 500))
	window.SetContent(layout)

	// Show the window
	window.ShowAndRun()
}
