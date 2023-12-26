/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a data file with fields",
	Long:  `use --field to create name of field and put values, Also use --databaseName to put the name on database`,

	//Run: func(cmd *cobra.Command, args []string) {
	//	var testNaming string = "test"
	//	user := User{Fields: createUserFields}
	//	commonDriver.Write("UserTest", testNaming, User{
	//		Fields: user.Fields,
	//	})
	//},
	Run: func(cmd *cobra.Command, args []string) {
		user := User{Fields: createUserFields}
		commonDriver.Write(createCollectionName, createDatabaseName, user)
		fmt.Println("Data stored at collectionName: " + createCollectionName + "Database name: " + createDatabaseName)
	},
}

type User struct {
	Fields map[string]string
}

var createUserFields map[string]string

var createDatabaseName string

var createCollectionName string

func init() {
	rootCmd.AddCommand(createCmd)

	createUserFields = make(map[string]string)

	createCmd.Flags().StringToStringVar(&createUserFields, "field", createUserFields, "Dynamic fields for the user")
	createCmd.MarkFlagRequired("field")
	createCmd.Flags().StringVarP(&createDatabaseName, "databaseName", "d", "", "used to name database names(file name) for reference")
	createCmd.MarkFlagRequired("databaseName")
	createCmd.Flags().StringVarP(&createCollectionName, "collectionName", "c", "User", "Used to name the Collections/folder name (Optional)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
