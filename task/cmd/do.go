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

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task complete ",
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

		if len(args) > 1{
			date = args[1]
			task = args[0]
		}else{
			date = t.Format("01-02-2006")
			task = args[0]
		}

		if(!TaskExist(tasks, task, date) && !DateExist(tasks, date)) {
			fmt.Println("Такой задачи и дня нет")

		}else{
			if !DateExist(tasks,date) {
				fmt.Println("Такого дня нет")
			}else {
				if !TaskExist(tasks, task, date) {
					fmt.Println("Такой задачи нет")
				} else{
					for i,  v := range tasks{
						if i == date{
							for e := 0; e < len(v); e++ {
								if (v[e].Name == task) {
									v[e].IsComplete = "!!Выполнено!!"
								}
							}
						}
					}


				}
			}
		}
		

		WriteJsonFile("tasks.json", tasks)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	//doCmd.Flags().StringP("do", "d", "", "Marks the tasks completed")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
