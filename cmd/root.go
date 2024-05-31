/*
Copyright © 2024 The Got'em Maintainers

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
	"os"

	"github.com/spf13/cobra"
	"github.com/thomaszub/gotem/internal"
)

var rootCmd = &cobra.Command{
	Use:   "gotem [flags] templateUri",
	Short: "A template scaffolding tool",
	Long:  "A template scaffolding tool written in Go.",
	Args:  cobra.ExactArgs(1),
	RunE:  main,
}

func init() {
	rootCmd.Flags().StringP("directory", "d", ".", "target directory to scaffold the template into")
	rootCmd.Flags().BoolP("force", "f", false, "force scaffolding, this deletes the target directory if it is not empty")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func main(cmd *cobra.Command, args []string) error {
	cfg, err := parseConfig(cmd, args)
	if err != nil {
		return err
	}
	return internal.Exec(cfg)
}

func parseConfig(cmd *cobra.Command, args []string) (internal.Config, error) {
	dir, err := cmd.Flags().GetString("directory")
	if err != nil {
		return internal.Config{}, err
	}
	force, err := cmd.Flags().GetBool("force")
	if err != nil {
		return internal.Config{}, err
	}
	return internal.Config{
		TemplateUri:     args[0],
		TargetDirectory: dir,
		Force:           force,
	}, nil
}
