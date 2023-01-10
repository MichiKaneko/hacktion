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

	"github.com/MichiKaneko/hacktion/config"
	"github.com/spf13/cobra"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long:  `A long description of your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			fmt.Println(err)
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			fmt.Println(err)
		}

		endpoint := "http://localhost:8080/api/token"

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
		if err != nil {
			fmt.Println(err)
		}

		var token Token
		json.Unmarshal(body, &token)

		cfg, err := config.Load()
		if err != nil {
			fmt.Println(err)
		}
		cfg.User.Token = token.Token
		cfg.User.Email = token.User.Email
		cfg.User.Name = token.User.Name
		cfg.User.ID = token.User.ID

		err = config.Save(cfg)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("response Body:", string(body))
		fmt.Println("token:", string(token.Token))
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("email", "e", "", "email")
	loginCmd.Flags().StringP("password", "p", "", "password")

}
