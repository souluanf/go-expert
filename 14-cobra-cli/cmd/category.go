package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		exists, _ := cmd.Flags().GetBool("exists")
		id, _ := cmd.Flags().GetInt16("id")
		fmt.Println("category called", category)
		fmt.Println("exists", exists)
		fmt.Println("id", id)
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "Name of category")
	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if category exists")
	categoryCmd.PersistentFlags().Int16P("id", "i", 0, "ID of category")
}
