package cmd

import (
	"goshop/internal/router"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		r := router.Setup()
		r.Run(":3000")
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
