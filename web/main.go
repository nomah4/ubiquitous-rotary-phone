package main

import (
	"context"
	"os"

	"github.com/GoSeoTaxi/dh8_msg/cmd"
)

func main() {
	// fmt.Println(color.HiMagentaString("Preparing..."))
	ctx, cancel := context.WithCancel(context.Background())

	err := cmd.RunService(ctx)
	if err != nil {
		os.Exit(1)
	}
	cancel()
	os.Exit(0)
}
