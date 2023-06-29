/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

// enpassCmd represents the enpass command
var enpassCmd = &cobra.Command{
	Use:   "enpass",
	Short: "Encode Password for user dsn connection Vertica",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if s := strings.Split(url.UserPassword("user", args[0]).String(), ":"); len(s) == 2 {
			fmt.Println(s[1])
		} else {
			fmt.Println("Not set")
		}
	},
}

func init() {
	rootCmd.AddCommand(enpassCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enpassCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enpassCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
