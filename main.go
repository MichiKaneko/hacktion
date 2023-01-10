/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/MichiKaneko/hacktion/cmd"
	"github.com/MichiKaneko/hacktion/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello!" + cfg.User.Name)
	cmd.Execute()
}
