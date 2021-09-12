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
	"io/ioutil"
	"log"
	helpers "monopoly/helpers"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new stage with no actors",
	Long:  `Initializes a new stage, a stage is a workspace that contains actors (modules, subprojects, or microservices)`,
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) != 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		stageName := args[0]
		stageDescription, err := cmd.Flags().GetString("description")
		if err != nil {
			fmt.Println(err)
		}

		createStage(stageName, stageDescription)
	},
}

func init() {
	initCmd.Flags().StringP("description", "d", "", "A description of your project")
	stageCmd.AddCommand(initCmd)
}

func createStage(stageName string, stageDescription string) {

	// Create Stage directory
	os.Mkdir(stageName, 0755)
	os.Mkdir(stageName+"/libs", 0755) // Generate shared libraries folder
	fmt.Println("Stage directory created")

	// Generate Stage map yaml
	stageMap := map[string]helpers.Stage{"stage": {stageName, stageDescription, []helpers.Actor{}}}

	data, marshalErr := yaml.Marshal(&stageMap)
	if marshalErr != nil {
		log.Fatal(marshalErr.Error())
	}

	writeErr := ioutil.WriteFile(stageName+"/stage.yaml", data, 0777)
	if writeErr != nil {
		log.Fatal(writeErr.Error())
	}

	fmt.Println("Stage map created")

	// Initialize git repository
	gitInitCmd := exec.Command("git", "init")
	gitInitCmd.Dir = stageName
	gitInitCmdErr := gitInitCmd.Run()
	if gitInitCmdErr != nil {
		fmt.Println(gitInitCmdErr.Error())
	}

	// Create .gitignore
	gitignoreGen := ioutil.WriteFile(stageName+"/.gitignore", []byte{}, 0777)
	if gitignoreGen != nil {
		fmt.Println(gitignoreGen.Error())
	}

}
