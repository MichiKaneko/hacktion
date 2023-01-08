/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			fmt.Println(err)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			fmt.Println(err)
		}

		endpoint := "https://hacktion.app.mode-co.jp/api/v1/login"

		login := Login{
			Email:    email,
			Password: password,
		}

		jsonBytes, err := json.Marshal(login)
		if err != nil {
			fmt.Println(err)
		}

		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBytes))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)

		body, err := io.ReadAll(resp.Body)

		fmt.Println("response Body:", string(body))

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("email", "e", "", "email")
	loginCmd.Flags().StringP("password", "p", "", "password")

}
