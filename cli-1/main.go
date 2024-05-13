package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mycli",
		Short: "MyCLI is a fast and flexible installation tool similar to Vue CLI",
		Long: `A Fast and Flexible installation tool built with love by you in Go.
Complete documentation is available at http://example.com/mycli`,
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.Execute()
}

var initCmd = &cobra.Command{
	Use:   "init [project name]",
	Short: "Initialize a new project",
	Long:  `Initialize a new project with the specified name.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please specify a project name")
			return
		}
		fmt.Println("Initializing project:", args[0])
	},
}
