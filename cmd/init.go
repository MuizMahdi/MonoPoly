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
	Short: "A brief description of your command",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		// fmt.Errorf("invalid color specified: %s", args[0]) // Prints error

		createStage(args[0])
	},
}

func init() {
	stageCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createStage(stageName string) {

	// Create Stage directory
	os.Mkdir(stageName, 0755)
	os.Mkdir(stageName+"/libs", 0755) // Generate shared libraries folder
	fmt.Println("Stage directory created")

	// Generate Stage map yaml
	stageMap := map[string]helpers.Stage{"stage": {
		stageName,
		"Description. TODO: Use flag for description or default",
		[]helpers.Actor{},
	}}

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
