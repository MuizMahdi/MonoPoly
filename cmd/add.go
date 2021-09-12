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
	"io/ioutil"
	"log"
	helpers "monopoly/helpers"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new actor to the stage",
	Long:  `Adds a new actor to the stage, an actor could be a module, subproject, or a microservice that would have its own repository`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		actorName := args[0]
		actorDescription, err := cmd.Flags().GetString("description")
		if err != nil {
			fmt.Println(err)
		}
		createActor(actorName, actorDescription)
	},
}

func init() {
	addCmd.Flags().StringP("description", "d", "", "A description of the actor")
	actorCmd.AddCommand(addCmd)
}

func createActor(actorName string, actorDescription string) {

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

	// Go through lines
	for scanner.Scan() {

		// Current line
		line := scanner.Text()

		// If actors start indicator is found
		if line == helpers.ActorsStartIdicator {
			isActorStartIndicatorFound = true
		}

		// Insert new actor if actors end indicator is found
		if isActorStartIndicatorFound && line == helpers.ActorsEndIdicator {

			isActorEndIndicatorFound = true

			// Add the new actor in this line
			err := helpers.InsertStringToFile(".gitignore", actorName+"\n", idx)
			if err != nil {
				log.Fatalf(err.Error())
			}

			break

		}

		idx++
	}

	if !isActorStartIndicatorFound && !isActorEndIndicatorFound {
		helpers.WriteLines([]string{helpers.ActorsStartIdicator, actorName, helpers.ActorsEndIdicator}, ".gitignore")
	}

	// Update the stage's map with the added actor

	// Read stage map file
	mapFile, mapReadErr := ioutil.ReadFile("stage.yaml")
	if mapReadErr != nil {
		log.Fatal(mapReadErr)
	}

	// Unmarshal stage map data
	mapData := make(map[string]helpers.Stage)
	mapUnmarshalErr := yaml.Unmarshal(mapFile, &mapData)
	if mapUnmarshalErr != nil {
		log.Fatal(mapUnmarshalErr)
	}

	// Add actor to actors in map
	if entry, ok := mapData["stage"]; ok {
		entry.Actors = append(entry.Actors, helpers.Actor{actorName, actorDescription})
		mapData["stage"] = entry
	}

	// Marshal updated map data
	data, marshalErr := yaml.Marshal(&mapData)
	if marshalErr != nil {
		log.Fatal(marshalErr.Error())
	}

	// Write updates
	writeErr := ioutil.WriteFile("stage.yaml", data, 0777)
	if writeErr != nil {
		fmt.Println(writeErr.Error())
	}

	fmt.Println("Actor added!")
}
