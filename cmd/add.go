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
	"bufio"
	"errors"
	"fmt"
	"log"
	helpers "monopoly/helpers"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		createActor(args[0])
	},
}

func init() {
	actorCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createActor(actorName string) {

	fmt.Println("Creating Actor...")

	// Create actor's folder
	os.Mkdir(actorName, 0775)

	// Initialize git repository
	gitInitCmd := exec.Command("git", "init")
	gitInitCmd.Dir = actorName
	gitInitCmdErr := gitInitCmd.Run()
	if gitInitCmdErr != nil {
		fmt.Println(gitInitCmdErr.Error())
	}

	isActorStartIndicatorFound := false
	isActorEndIndicatorFound := false

	fmt.Println("Reading gitignore...")

	// Read gitignore
	file, err := os.Open(".gitignore")

	// Handle read errors
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Close the file
	defer file.Close()

	// Read lines
	scanner := bufio.NewScanner(file)

	idx := 0

	fmt.Println("Scanning gitingore lines...")

	// Go through lines
	for scanner.Scan() {

		// Current line
		line := scanner.Text()

		fmt.Println("Scanned Line: " + line)

		// If actors start indicator is found
		if line == helpers.ActorsStartIdicator {
			isActorStartIndicatorFound = true
		}

		// Insert new actor if actors end indicator is found
		if isActorStartIndicatorFound && line == helpers.ActorsEndIdicator {

			fmt.Println("Found the end indicator.")
			fmt.Println("Adding Actor to above the end indicator...")

			isActorEndIndicatorFound = true

			// Add the new actor in this line
			err := helpers.InsertStringToFile(".gitignore", actorName+"\n", idx)
			if err != nil {
				log.Fatalf(err.Error())
			}

			fmt.Println("Done!")

			break

		}

		idx++
	}

	if !isActorStartIndicatorFound && !isActorEndIndicatorFound {
		helpers.WriteLines([]string{helpers.ActorsStartIdicator, actorName, helpers.ActorsEndIdicator}, ".gitignore")
	}

	// Update the stage's map with the added actor

}
