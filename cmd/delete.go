// Copyright (c) 2022 EPAM Systems, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"github.com/agilestacks/dexctl/dex"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <password | oidc>",
	Short: "Delete dex resource",
}

var deletePasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "delete static password in dex",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dex.DeletePassword(email)
	},
}

var deleteOidcCmd = &cobra.Command{
	Use:   "oidc",
	Short: "delete oauth2 static client in dex",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dex.DeleteOidc(clientId)
	},
}

func init() {
	deletePasswordCmd.Flags().StringVarP(&email, "email", "e", "", "Username email")
	deletePasswordCmd.MarkFlagRequired("email")
	deleteCmd.AddCommand(deletePasswordCmd)

	deleteOidcCmd.Flags().StringVarP(&clientId, "client-id", "c", "", "Client ID used to identify the client")
	deleteOidcCmd.MarkFlagRequired("client-id")
	deleteCmd.AddCommand(deleteOidcCmd)

	rootCmd.AddCommand(deleteCmd)
}
