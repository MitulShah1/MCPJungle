package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/mcpjungle/mcpjungle/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		if !errors.Is(err, cmd.ErrSilent) {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		os.Exit(1)
	}
}
