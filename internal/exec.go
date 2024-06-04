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
	"os/exec"
	"path/filepath"
)

type Config struct {
	TemplateUri     string
	TargetDirectory string
	Force           bool
}

func Exec(cfg Config) error {
	err := createDirectory(cfg.TargetDirectory, cfg.Force)
	if err != nil {
		return err
	}
	err = cloneTemplate(cfg.TargetDirectory, cfg.TemplateUri)
	if err != nil {
		return err
	}
	err = os.RemoveAll(filepath.Join(cfg.TargetDirectory, ".git"))
	if err != nil {
		return err
	}
	return nil
}

func createDirectory(dir string, force bool) error {
	empty, err := isDirectoryEmpty(dir)
	if err != nil {
		return err
	}
	if !empty {
		if !force {
			return fmt.Errorf("directory %s is not empty, scaffolding must be enforced", dir)
		}
		err := os.RemoveAll(dir)
		if err != nil {
			return err
		}
	}
	return os.MkdirAll(dir, 0744)
}

func isDirectoryEmpty(dir string) (bool, error) {
	f, err := os.Open(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return true, nil
		}
		return false, err
	}
	defer f.Close()
	names, err := f.Readdirnames(1)
	if err != nil {
		if err != io.EOF {
			return false, err
		}
		return true, nil
	}
	return len(names) == 0, nil
}

func cloneTemplate(dir string, uri string) error {
	cmd := exec.Command("git", "clone", "--depth", "1", uri, dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
