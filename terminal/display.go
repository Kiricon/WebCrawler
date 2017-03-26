package terminal

import "fmt"
import "time"

//Display - Struct of displaying info to user in the terminal
type Display struct {
	SpinnerIndex int
}

//NewDisplay - Returns a new display struct with default values
func NewDisplay() Display {
	return Display{SpinnerIndex: 0}
}

var spinner = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

//Draw - Draw the current display in the terminal.
func (d *Display) Draw() {
	fmt.Print("\r" + spinner[d.SpinnerIndex] + " Crawling...")

	if d.SpinnerIndex == len(spinner)-1 {
		d.SpinnerIndex = 0
	} else {
		d.SpinnerIndex++
	}

	time.Sleep(1000)
	go d.Draw()
	//fmt.Printf("\r Links Found:%d | Running Routines:%d", c.Stats.TotalLinks, c.Stats.TotalRoutines)
}
