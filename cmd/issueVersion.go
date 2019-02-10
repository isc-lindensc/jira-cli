// Copyright © 2019 Robert Sotomski <sotomski@gmail.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/sotomskir/jira-cli/jiraApi"
	"github.com/sotomskir/jira-cli/logger"
	"github.com/spf13/cobra"
)

// issueVersionCmd represents the issueVersion command
var issueVersionCmd = &cobra.Command{
	Use:   "version ISSUE_KEY VERSION",
	Short: "Set issue fix version",
	Aliases: []string{"v"},
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		issueKey := args[0]
		version := args[1]
		result := jiraApi.SetFixVersion(issueKey, version)
		if result {
			logger.SuccessF("Success version %s set for issue %s\n", version, issueKey)
		}
	},
}

func init() {
	issueCmd.AddCommand(issueVersionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// issueVersionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issueVersionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}