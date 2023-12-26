/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// readAllCmd represents the readAll command
var readAllCmd = &cobra.Command{
	Use:   "readAll",
	Short: "ability to read all data files of a collection",
	Long:  `we can read entire data of a particular collection by giving the name of collection`,
	Run: func(cmd *cobra.Command, args []string) {
		records, err := commonDriver.ReadAll(collectionName)
		if err != nil {
			fmt.Println("no collection found")
		}
		fmt.Println(records)
	},
}
var collectionName string

func init() {
	rootCmd.AddCommand(readAllCmd)
	readAllCmd.Flags().StringVarP(&collectionName, "collectionName", "c", "", "displays entire value of a collection")
	//readAllCmd.Flags().StringVarP(&hi, "readAll", "ra", "", "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
