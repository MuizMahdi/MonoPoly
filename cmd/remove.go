/*
Copyright © 2021 Muizz M. Mahdy <muizzmahdy@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package cmd

import (
	"errors"
	"fmt"
	helpers "monopoly/helpers"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		removeActor(args[0])
	},
}

func init() {
	actorCmd.AddCommand(removeCmd)
}

func removeActor(actorName string) {

	fmt.Println("Removing " + actorName + " ...")

	// Check if there're any pending changes inside actor's directory
	gitStatusCmd := exec.Command("git", "status", "-s")
	gitStatusCmd.Dir = actorName
	gitStatusCmdOut, gitStatusCmdErr := gitStatusCmd.Output()
	if gitStatusCmdErr != nil {
		fmt.Println(gitStatusCmdErr.Error())
	}

	if len(gitStatusCmdOut) > 0 { // If there are pending changes

		// Ask user if they still want to delete them
		if helpers.PromptYesNoQuestion("There are uncommited changes in the directory, are you sure that you want to delete it?") {

			// Delete actor's directory
			dirRemoveErr := os.RemoveAll(actorName)
			if dirRemoveErr == nil {
				// Remove actor from the stage map
				clearActorMetadata(actorName)
			}

		} else {
			return
		}
	}

	// Delete actor's directory
	dirRemoveErr := os.RemoveAll(actorName)
	if dirRemoveErr == nil {
		// Remove actor from the stage map
		clearActorMetadata(actorName)
	}

	fmt.Println("Actor removed.")
}

func clearActorMetadata(actorName string) {

}
