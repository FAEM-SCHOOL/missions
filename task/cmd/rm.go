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


		task, _ := cmd.Flags().GetString("today")
		del_day, _ := cmd.Flags().GetString("day")

		var day string
		var del_task string

		//Check input
		if len(args) > 1 {
			day = args[1]
			del_task = args[0]

			if !DateExist(tasks,day){
				fmt.Println("Такого дня не сущетсвует")
				return
			}else {
				if(DateExist(tasks, day) && !TaskExist(tasks, del_task, day)){
					fmt.Println("В этот день такой задачи не существует")
					return
				}
			}
		}else{
			if del_day != ""{

				if(!DateExist(tasks, del_day)){
					fmt.Println("Такого дня не существует")
					return
				}else {
					delete(tasks, del_day)
					WriteJsonFile("tasks.json", tasks)
					return
				}
			}

			if task != ""{
				day = time.Now().Format("01-02-2006")
				del_task = task
				if !TaskExist(tasks, del_task, day) {
					fmt.Println("Такой задачи нет")
					return
				}
			}
		}

		index := SearchTaskIndex(tasks, del_task, day)
		if len(tasks[day]) - 1 == 0{
			delete(tasks, day)
		}else {
			slice := tasks[day]
			slice = append(slice[:index], slice[index+1:]...)
			tasks[day] = slice
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
