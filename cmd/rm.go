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

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove items from the shopping list",
	Long:  `Remove items from the shopping list`,
	Run: func(cmd *cobra.Command, args []string) {
		argsInt := make([]int, len(args))
		for i, arg := range args {
			itemID, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Invalid item ID: %s\n", arg)
				continue
			}
			argsInt[i] = itemID
		}
		repository.RemoveFromShoppingList(config.DB, argsInt...)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
