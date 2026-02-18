/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"goshop/cmd"
	"goshop/config"
)

func init() {
	config.InitDB("goshop.db")
}

func main() {
	cmd.Execute()
}
