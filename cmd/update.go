/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the write command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "updates fields of a database",
	Long:  `Updates fields values of a database by adding a new field, change values, making fields blank. `,

	Run: func(cmd *cobra.Command, args []string) {
		user := User{Fields: updateUserFields}
		commonDriver.Write("UserTest", updateDatabaseName, user)
	},
}

var nameValue string
var updateUserFields map[string]string
var updateDatabaseName string
var updateCollectionName string

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringToStringVar(&updateUserFields, "field", updateUserFields, "Dynamic fields for the user")
	updateCmd.MarkFlagRequired("field")
	updateCmd.Flags().StringVarP(&updateDatabaseName, "databaseName", "d", "", "used to name database names(file name) for reference")
	updateCmd.MarkFlagRequired("databaseName")
	updateCmd.Flags().StringVarP(&updateCollectionName, "collectionName", "c", "User", "Used to refer to the collection in which data is stored")

	//updateCmd.Flags().StringVarP(&nameValue, "name", "n", "", "used to store names") //working

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
