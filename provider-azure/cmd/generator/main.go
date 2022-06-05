/*
Copyright 2021 The Crossplane Authors.

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

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/pipeline"

	"github.com/upbound/official-providers/provider-azure/config"
)

func main() {
	// delete API dirs
	deleteGenDirs("apis", map[string]struct{}{
		"v1alpha1": {},
		"rconfig":  {},
	})
	// delete controller dirs
	deleteGenDirs("internal/controller", map[string]struct{}{
		"providerconfig": {},
	})

	pipeline.Run(config.GetProvider(), "")
}

// delete API subdirs for a clean start
func deleteGenDirs(rootDir string, keepMap map[string]struct{}) {
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		panic(errors.Wrapf(err, "cannot list files under %s", rootDir))
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}
		if _, ok := keepMap[f.Name()]; ok {
			continue
		}
		removeDir := filepath.Join(rootDir, f.Name())
		if err := os.RemoveAll(removeDir); err != nil {
			panic(errors.Wrapf(err, "cannot remove API dir: %s", removeDir))
		}
	}
}
