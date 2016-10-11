package chrysanthemum

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"time"
)

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
	writer  io.Writer
	stop    chan bool
	stopped bool
}

func New(writer io.Writer, text string) *Chrysanthemum {
	fmt.Fprint(writer, "   "+text)
	return &Chrysanthemum{
		writer:  writer,
		stop:    make(chan bool),
		stopped: false,
	}
}

func (c *Chrysanthemum) Start() *Chrysanthemum {
	fmt.Fprint(c.writer, "\033[?25l") // hide cursor

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
			fmt.Fprintf(c.writer, "\r %s ", Frames[i])
			i++
		}
	}()

	return c
}

func (c *Chrysanthemum) end(flag string) {
	if c.stopped {
		return
	}
	c.stop <- true
	c.stopped = true
	fmt.Fprintf(c.writer, "\033[?25h") // show cursor
	fmt.Fprintf(c.writer, "\r %s \n", flag)
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
