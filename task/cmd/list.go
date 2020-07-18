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
	"strconv"
	"strings"
	"time"
)

//Converting a date to a number for convenience

func DateLess(date1 string, date2 string) bool{
	date1_split := strings.Split(date1, "-")
	day1, _ := strconv.Atoi(date1_split[0])
	mouth1, _ := strconv.Atoi(date1_split[1])
	year1, _ := strconv.Atoi(date1_split[2])

	date2_split := strings.Split(date2, "-")
	day2, _ := strconv.Atoi(date2_split[0])
	mouth2, _ := strconv.Atoi(date2_split[1])
	year2, _ := strconv.Atoi(date2_split[2])

	if year1 < year2{
		return true
	}
	if mouth1 < mouth2 && year2 == year1{
		return  true
	}

	if day1 < day2 && year1 == year2 && mouth2 == mouth1{
		return true
	}
	return false
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show our list with tasks",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		tasks := ReadJsonFile("tasks.json")

		//Slice map's keys for sorting
		keys := make([]string, 0, len(tasks))
		for i, _ := range tasks{
			keys = append(keys, i)
		}

		//Sorting
		for i := 1; i < len(keys); i++{
			temp := keys[i]
			j := i
			for j > 0 && DateLess(keys[j-1], temp){
				keys[j] = keys[j - 1]
				j--
			}
			keys[j] = temp
		}

		date := time.Now()

		show, _ := cmd.Flags().GetString("show")

		//Output of all tasks
		if show == "all" || show == ""{
			for i := 0; i < len(keys); i++ {
				fmt.Println(keys[i])
				for e := 0; e < len(tasks[keys[i]]); e++ {
					fmt.Println("  ", tasks[keys[i]][e].Name, "-->", tasks[keys[i]][e].IsComplete)
				}
			}
		}
		//Output of tasks for today
		if show == "today" {
			for i, v := range tasks {
				if i == date.Format("01-02-2006") {
					fmt.Println(date.Format("01-02-2006"))
					for e := 0; e < len(v); e++ {
						fmt.Println("  ", v[e].Name, "-->", v[e].IsComplete)
					}
				}
			}
		}

		//Output of tasks for the specified day
		for i, v := range tasks {
			if i == show {
				fmt.Println(show)
				for e := 0; e < len(v); e++ {
					fmt.Println("  ", v[e].Name, "-->", v[e].IsComplete)
					}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	listCmd.Flags().StringP("show", "s", "", "Shows all tasks or today's tasks")

	//listCmd.PersistentFlags().String("foo", "", "Hello")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
