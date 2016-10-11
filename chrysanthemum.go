package chrysanthemum

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"os"
	"time"
)

var isTerminal = isatty.IsTerminal(os.Stdout.Fd())

var Frames = []string{
	color.MagentaString("⠋"),
	color.MagentaString("⠙"),
	color.MagentaString("⠹"),
	color.MagentaString("⠸"),
	color.MagentaString("⠼"),
	color.MagentaString("⠴"),
	color.MagentaString("⠦"),
	color.MagentaString("⠧"),
	color.MagentaString("⠇"),
	color.MagentaString("⠏"),
}
var Success = color.GreenString("✓")
var Fail = color.RedString("✗")

type Chrysanthemum struct {
	stop    chan bool
	stopped bool
}

func New(text string) *Chrysanthemum {
	if !isTerminal {
		fmt.Print(text)
	} else {
		fmt.Print("   " + text)
	}
	return &Chrysanthemum{
		stop:    make(chan bool),
		stopped: false,
	}
}

func (c *Chrysanthemum) Start() *Chrysanthemum {
	if !isTerminal {
		return c
	}

	fmt.Print("\033[?25l") // hide cursor

	i := 0
	go func() {
		for {
			select {
			case <-c.stop:
				return
			default:
				time.Sleep(100 * time.Millisecond)
			}
			if i == len(Frames) {
				i = 0
			}
			fmt.Printf("\r %s ", Frames[i])
			i++
		}
	}()

	return c
}

func (c *Chrysanthemum) end(flag string) {
	if !isTerminal {
		fmt.Println()
		return
	}

	if c.stopped {
		return
	}
	c.stop <- true
	c.stopped = true
	fmt.Printf("\033[?25h") // show cursor
	fmt.Printf("\r %s \n", flag)
}

func (c *Chrysanthemum) Successed() {
	c.end(Success)
}

func (c *Chrysanthemum) Failed() {
	c.end(Fail)
}

func (c *Chrysanthemum) End() {
	c.Successed()
}
