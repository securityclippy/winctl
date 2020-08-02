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
	"github.com/securityclippy/winctl/pkg/input"
	"github.com/securityclippy/winctl/pkg/window"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// sendleftclicklocationCmd represents the sendleftclicklocation command
var sendleftclicklocationCmd = &cobra.Command{
	Use:   "sendleftclicklocation",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		a1,_ := strconv.Atoi(args[1])
		a2,_ := strconv.Atoi(args[2])



		hwdn, err := window.GetHandle("", args[0])
		if err != nil {
			log.Fatal(err)
		}


		loc := &input.POINT{
			X: int32(a1),
			Y: int32(a2),
		}

		log.Infoln("sending left click in 10 seconds to loc: %+v\n", loc)

		time.Sleep(time.Second * 10)
		err = input.SendLeftClickLocation(hwdn, loc)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	mouseCmd.AddCommand(sendleftclicklocationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendleftclicklocationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendleftclicklocationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
