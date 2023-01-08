/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// hackCmd represents the hack command
var hackCmd = &cobra.Command{
	Use:   "hack",
	Short: "A brief description of your command",
	Long:  `test`,
	Run: func(cmd *cobra.Command, args []string) {
		postcode, err := cmd.Flags().GetString("post")
		if err != nil {
			log.Fatal(err)
		}

		endpoint := "https://zipcloud.ibsnet.co.jp/api/search?zipcode=" + postcode

		res, err := http.Get(endpoint)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			log.Fatal(res.Status)
		}
		body, _ := io.ReadAll(res.Body)
		fmt.Print(string(body))

	},
}

func init() {
	rootCmd.AddCommand(hackCmd)
	hackCmd.Flags().StringP("post", "p", "", "cache post")
}
