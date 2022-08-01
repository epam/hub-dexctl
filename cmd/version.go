// Copyright (c) 2022 EPAM Systems, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	ref     = "main"
	commit  = "HEAD"
	buildAt = "now"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print dexctl version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("dexctl %s %s build at %s %s\n", ref, commit, buildAt, runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
