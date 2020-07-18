/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete task or day",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		tasks := ReadJsonFile("tasks.json")

		t := time.Now()

		var date string
		var task string

		today, _ := cmd.Flags().GetString("today")
		day, _ := cmd.Flags().GetString("day")


		if day != "" {
			flag := false
			for i, _ := range  tasks{
				if(i == day){
					flag = true
				}
			}

			if flag{
				delete(tasks, day)
			}else {
				fmt.Println("Такого дня нет")
			}

		} else{

			if len(args) == 0 {
				date = t.Format("01-02-2006")
				task = today
			} else {
				date = args[0]
				task = args[1]
			}

			var index int
			flag := false
			for i := 0; i < len(tasks[date]); i++ {
				if tasks[date][i].Name == task {
					index = i;
					flag = true
				}
			}
			if (flag) {
				slice := tasks[date]
				slice = append(slice[:index], slice[index+1:]...)
				tasks[date] = slice
				if len(tasks[date]) == 0 {
					delete(tasks, date)
				}
			} else {
				fmt.Println("Такой задачи нет")
			}
		}


		WriteJsonFile("tasks.json", tasks)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rmCmd.Flags().StringP("today", "t", "", "Delete task in for today")
	rmCmd.Flags().StringP("day", "d", "", "Delete a specific day")

}
