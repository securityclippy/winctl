// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/securityclippy/winctl/pkg/input"
	"github.com/securityclippy/winctl/pkg/window"
	"time"

	"github.com/spf13/cobra"
)

var cls string
var wName string

// sendleftclickCmd represents the sendleftclick command
var sendleftclickCmd = &cobra.Command{
	Use:   "sendleftclick",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		hwdn, err := window.GetHandle("", args[0])
		if err != nil {
			log.Fatal(err)
		}

		log.Infoln("sending left click in 10 seconds to current mouse loc")


		time.Sleep(time.Second * 10)
		err = input.SendLeftClick(hwdn, true)
		if err != nil {
			log.Fatal(err)
		}


		fmt.Println("sendleftclick called")
	},
}

func init() {
	mouseCmd.AddCommand(sendleftclickCmd)
	//sendleftclickCmd.Flags().StringVarP(&cls, "class", "c", "", "window class")
	//sendleftclickCmd.MarkFlagRequired("class")
	//sendleftclickCmd.Flags().StringVarP(&wName, "window", "w", "", "window name")
	//sendleftclickCmd.MarkFlagRequired("window")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendleftclickCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendleftclickCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
