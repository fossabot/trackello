// Copyright © 2016 Nick Klauer <klauer@gmail.com>
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

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure your Trello API keys and preferred Trello Board",
	Long: `To communicate to a Trello API, you will need to configure a
minimum of 3 parameters:
  - TRELLO_APPKEY
  - TRELLO_TOKEN
  - preferredBoard

If you do not know or have any of these, you can review  the documentation
on this site: https://trello.com/app-key
`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("config called")
	},
}

func init() {
	RootCmd.AddCommand(configCmd)
}
