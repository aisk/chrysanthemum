package main

import (
	"github.com/aisk/chrysanthemum"
	"os"
	"time"
)

func main() {
	c := chrysanthemum.New(os.Stdout, "I'll be ok").Start()
	time.Sleep(5 * time.Second)
	c.End()
	c = chrysanthemum.New(os.Stdout, "I'll be error").Start()
	time.Sleep(5 * time.Second)
	c.Failed()
}
