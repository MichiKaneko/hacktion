/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// hackCmd represents the hack command
var hackCmd = &cobra.Command{
	Use:   "hack",
	Short: "A brief description of your command",
	Long:  `test`,
	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			log.Fatal(err)
		}

		content, err := cmd.Flags().GetString("content")
		if err != nil {
			log.Fatal(err)
		}

		endpoint := "http://localhost:8080/api/users/post"

		post := Post{
			Title:   title,
			Content: content,
		}

		jsonBytes, err := json.Marshal(post)
		if err != nil {
			log.Fatal(err)
		}

		token := viper.GetString("user.token")
		fmt.Println(token)
		if token == "" {
			log.Fatal("Token is empty. Please login first.")
		}

		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBytes))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", token)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(body))
	},
}

func init() {
	rootCmd.AddCommand(hackCmd)
	hackCmd.Flags().StringP("title", "t", "", "cache title")
	hackCmd.Flags().StringP("content", "c", "", "cache content")
}
