/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

// diskUsageCmd represents the diskUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "Information about disk usage",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("diskUsage called")
		usage := du.NewDiskUsage(".")
		fmt.Printf("used: %v bytes, free: %v bytes\n", usage.Used(), usage.Free())
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
