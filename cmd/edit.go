/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"goshop/config"
	"goshop/internal/repository"
	"strconv"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing item",
	Long:  `Edit an existing item in the shopping list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Arguments must be [Item Number] [New Item Name]")
			return
		}
		itemNumber, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid item number")
			return
		}
		newItemName := args[1]

		err = repository.EditShoppingList(config.DB, itemNumber, newItemName)
		if err != nil {
			fmt.Println("Error editing item:", err)
			return
		}
		fmt.Printf("Item %d edited to %s\n", itemNumber, newItemName)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
