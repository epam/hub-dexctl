// Copyright (c) 2022 EPAM Systems, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"github.com/epam/hub-dexctl/dex"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <password | oidc>",
	Short: "Update dex resource",
}

var updatePasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "update static password in dex",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dex.UpdatePassword(email, newPassword, newEmail)
	},
}

var updateOidcCmd = &cobra.Command{
	Use:   "oidc",
	Short: "update oauth2 static client in dex",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dex.UpdateOidc(clientId, redirectUris, trustedPeers, name, logoUrl)
	},
}

func init() {
	updatePasswordCmd.Flags().StringVarP(&email, "email", "e", "", "Email of username which to update")
	updatePasswordCmd.MarkFlagRequired("email")
	updatePasswordCmd.Flags().StringVarP(&newEmail, "new-email", "n", "", "New username email")
	updatePasswordCmd.MarkFlagRequired("new-email")
	updatePasswordCmd.Flags().StringVarP(&newPassword, "new-password", "p", "", "New username password")
	updatePasswordCmd.MarkFlagRequired("new-password")
	updatePasswordCmd.MarkFlagsRequiredTogether("email", "new-email", "new-password")
	updateCmd.AddCommand(updatePasswordCmd)

	updateOidcCmd.Flags().StringVarP(&clientId, "client-id", "c", "", "Client ID used to identify the client")
	updateOidcCmd.MarkFlagRequired("client-id")
	updateOidcCmd.Flags().StringArrayVarP(&redirectUris, "redirect-uris", "r", nil, "A registered set of redirect URIs")
	updateOidcCmd.Flags().StringArrayVarP(&trustedPeers, "trusted-peers", "t", nil, "TrustedPeers are a list of peers which can issue tokens on this client's behalf using the dynamic scope")
	updateOidcCmd.Flags().StringVarP(&name, "name", "n", "", "Name used when displaying this client to the end user")
	updateOidcCmd.Flags().StringVarP(&logoUrl, "logo-url", "l", "", "LogoURL used when displaying this client to the end user")
	updateCmd.AddCommand(updateOidcCmd)

	rootCmd.AddCommand(updateCmd)
}
