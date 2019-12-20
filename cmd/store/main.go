package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	root := &cobra.Command{
		Use:   "store",
		Short: "cloth store api cli",
	}
	root.AddCommand(
		createUser(),
		server(),
	)
	err := root.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
