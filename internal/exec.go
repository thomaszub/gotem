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
package internal

import (
	"fmt"
	"io"
	"os"
)

type Config struct {
	TemplateUri     string
	TargetDirectory string
	Force           bool
}

func Exec(cfg Config) error {
	err := createTargetDirectory(cfg.TargetDirectory, cfg.Force)
	if err != nil {
		return err
	}
	return nil
}

func createTargetDirectory(dir string, force bool) error {
	err := os.MkdirAll(dir, 0744)
	if err != nil {
		return err
	}
	if force {
		return nil
	}
	return validateTargetDirectoryNonEmpty(dir)
}

func validateTargetDirectoryNonEmpty(dir string) error {
	f, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	return fmt.Errorf("directory %s is not empty, scaffolding must be enforced", dir)
}
