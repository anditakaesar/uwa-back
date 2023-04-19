package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/anditakaesar/uwa-back/env"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting application %s", env.AppName())
	}
}

func run() error {

	return nil
}
