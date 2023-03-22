/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"kubenodes/cmd"
)

var version = "dev"

func main() {
	cmd.SetVersion(version)
	cmd.Execute()
}
