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
	"github.com/spf13/cobra"
	"time"
)

func IsKeyExist(Map map[string][]Task, key string) bool  {
	for i, _ := range Map{
		if i == key{
			return true
		}
	}
	return  false
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to our list",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tasks := ReadJsonFile("tasks.json")

		//Create a new task
		var newTask Task
		newTask.Name = args[0]
		newTask.IsComplete = "Incomplete"

		//Add a task to the map
		t := time.Now()
		var date string
		if len(args) == 1 {
			date = t.Format("01-02-2006")
		} else{
			date = args[1]
		}

		tasks[date] = append(tasks[date], newTask)

		WriteJsonFile("tasks.json", tasks)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}