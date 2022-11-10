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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/upbound/upjet/pkg/pipeline"

	"github.com/upbound/provider-azure/config"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	p := config.GetProvider()
	pipeline.Run(p, absRootDir)
	if fp := os.Getenv("SKIPPED_RESOURCES_CSV"); len(fp) != 0 {
		if err := os.WriteFile(fp, []byte(strings.Join(p.GetSkippedResourceNames(), ",")), 0o600); err != nil {
			panic(fmt.Sprintf("cannot write skipped resources CSV to file %s: %s", fp, err.Error()))
		}
	}
}
