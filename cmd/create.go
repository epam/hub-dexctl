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

var createCmd = &cobra.Command{
	Use:   "create <password | oidc>",
	Short: "Create dex resource",
}

var createPasswordCmd = &cobra.Command{
	Use:   "password",
	Short: "Create static password in dex",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dex.CreatePassword(email, password)
	},
}

var createOidcCmd = &cobra.Command{
	Use:   "oidc",
	Short: "Create oauth2 static client in dex",
	RunE: func(cmd *cobra.Command, args []string) error {
		return dex.CreateOidc(clientId, clientSecret, redirectUris, trustedPeers, public, name, logoUrl)
	},
}

func init() {
	createPasswordCmd.Flags().StringVarP(&email, "email", "e", "", "Username email")
	createPasswordCmd.MarkFlagRequired("email")
	createPasswordCmd.Flags().StringVarP(&password, "password", "p", "", "Username password")
	createPasswordCmd.MarkFlagRequired("password")
	createPasswordCmd.MarkFlagsRequiredTogether("email", "password")
	createCmd.AddCommand(createPasswordCmd)

	createOidcCmd.Flags().StringVarP(&clientId, "client-id", "c", "", "Client ID used to identify the client")
	createOidcCmd.MarkFlagRequired("client-id")
	createOidcCmd.Flags().StringVarP(&clientSecret, "client-secret", "s", "", "Client secret used to identify the client")
	createOidcCmd.MarkFlagRequired("client-secret")
	createOidcCmd.Flags().StringArrayVarP(&redirectUris, "redirect-uris", "r", nil, "A registered set of redirect URIs")
	createOidcCmd.MarkFlagRequired("redirect-uris")
	createOidcCmd.Flags().StringArrayVarP(&trustedPeers, "trusted-peers", "t", nil, "TrustedPeers are a list of peers which can issue tokens on this client's behalf using the dynamic scope")
	createOidcCmd.Flags().BoolVarP(&public, "public", "p", true, "Public clients are inspired by Google’s “Installed Applications” and are meant to impose restrictions on applications that don’t intend to keep their client secret private")
	createOidcCmd.Flags().StringVarP(&name, "name", "n", "", "Name used when displaying this client to the end user")
	createOidcCmd.MarkFlagRequired("name")
	createOidcCmd.Flags().StringVarP(&logoUrl, "logo-url", "l", "", "LogoURL used when displaying this client to the end user")
	createOidcCmd.MarkFlagsRequiredTogether("client-id", "client-secret", "redirect-uris", "name")
	createCmd.AddCommand(createOidcCmd)

	rootCmd.AddCommand(createCmd)
}
