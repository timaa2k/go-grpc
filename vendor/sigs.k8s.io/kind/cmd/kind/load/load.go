/*
Copyright 2019 The Kubernetes Authors.

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

// Package load implements the `load` command
package load

import (
	"github.com/spf13/cobra"

	dockerimage "sigs.k8s.io/kind/cmd/kind/load/docker-image"
	imagearchive "sigs.k8s.io/kind/cmd/kind/load/image-archive"
)

// NewCommand returns a new cobra.Command for get
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "load",
		Short: "Loads images into nodes",
		Long:  "Loads images into node from an archive or image on host",
	}
	// add subcommands
	cmd.AddCommand(dockerimage.NewCommand())
	cmd.AddCommand(imagearchive.NewCommand())
	return cmd
}
