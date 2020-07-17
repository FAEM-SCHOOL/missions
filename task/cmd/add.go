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
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

type Task struct {
	Name string
	IsComplete string
}

type ListTask struct {
	Tasks []Task
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
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile("tasks.json")
		if err != nil {
			log.Fatal("Cannot load settings:", err)
			}

			var tasks ListTask
			err = json.Unmarshal(data, &tasks)
			if err != nil {
			log.Fatal("Invalid settings format:", err)
			}

			var newTask Task
			newTask.Name = args[0]
			newTask.IsComplete = "Incomplete"

			tasks.Tasks = append(tasks.Tasks, newTask)

			data, err = json.MarshalIndent(&tasks, "", " ")
			if err != nil{
				log.Fatal("JSON marshaling failed:", err)
			}

			err = ioutil.WriteFile("tasks.json", data, 0)
			if err != nil {
				log.Fatal("Cannot write updated settings file:", err)
			}
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