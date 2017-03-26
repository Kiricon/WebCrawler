package terminal

import (
	"WebCrawler/stats"
	"fmt"
	"time"
)

//Display - Struct of displaying info to user in the terminal
type Display struct {
	spinnerIndex int
	stats        *stats.Stats
}

//NewDisplay - Returns a new display struct with default values
func NewDisplay(s *stats.Stats) Display {
	return Display{spinnerIndex: 0, stats: s}
}

var spinner = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

//Draw - Draw the current display in the terminal.
func (d *Display) Draw() {

	fmt.Printf("\033[0;0H")
	fmt.Print(spinner[d.spinnerIndex] + " Crawling...\n")
	fmt.Printf("Links Found: %d ✔ \n", d.stats.TotalLinks)
	fmt.Printf("Crawlers Running: %d ◉ \n", d.stats.TotalRoutines)
	fmt.Print("Runtime:0\n")

	if d.spinnerIndex == len(spinner)-1 {
		d.spinnerIndex = 0
	} else {
		d.spinnerIndex++
	}

	time.Sleep(1000)
	go d.Draw()
	//fmt.Printf("\r Links Found:%d | Running Routines:%d", c.Stats.TotalLinks, c.Stats.TotalRoutines)
}
