package terminal

import "fmt"

//Display - Struct of displaying info to user in the terminal
type Display struct {
	SpinnerIndex int8
}

var spinner = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

//Draw - Draw the current display in the terminal.
func (d *Display) Draw() {
	fmt.Print("\r" + spinner[d.SpinnerIndex] + " Crawling...")
	//fmt.Printf("\r Links Found:%d | Running Routines:%d", c.Stats.TotalLinks, c.Stats.TotalRoutines)
}
