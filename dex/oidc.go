// Copyright (c) 2022 EPAM Systems, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package dex

import (
	"context"
	"fmt"

	"github.com/dexidp/dex/api/v2"
)

func CreateOidc(clientId string, clientSecret string, redirectUris []string, trustedPeers []string, public bool, name string, logoUrl string) error {
	conn, err := newGrpcConnection()

	if err != nil {
		return fmt.Errorf("failed to initialise connection to dex api: %s", err)
	}
	defer conn.Close()

	client := api.NewDexClient(conn)

	c := api.Client{
		Id:           clientId,
		Secret:       clientSecret,
		RedirectUris: redirectUris,
		TrustedPeers: trustedPeers,
		Public:       public,
		Name:         name,
		LogoUrl:      logoUrl,
	}

	req := &api.CreateClientReq{
		Client: &c,
	}

	if resp, err := client.CreateClient(context.TODO(), req); err != nil || (resp != nil && resp.AlreadyExists) {
		if resp != nil && resp.AlreadyExists {
			return fmt.Errorf("oauth client %s already exists", clientId)
		}

		return fmt.Errorf("failed to create oauth client %s", err)
	}

	return nil
}

func UpdateOidc(clientId string, redirectUris []string, trustedPeers []string, name string, logoUrl string) error {
	conn, err := newGrpcConnection()

	if err != nil {
		return fmt.Errorf("failed to initialise connection to dex api: %s", err)
	}
	defer conn.Close()

	client := api.NewDexClient(conn)

	req := &api.UpdateClientReq{
		Id:           clientId,
		RedirectUris: redirectUris,
		TrustedPeers: trustedPeers,
		Name:         name,
		LogoUrl:      logoUrl,
	}

	if resp, err := client.UpdateClient(context.TODO(), req); err != nil || (resp != nil && resp.NotFound) {
		if resp != nil && resp.NotFound {
			return fmt.Errorf("oauth client %s not found", clientId)
		}

		return fmt.Errorf("failed to update oauth client %s", err)
	}

	return nil
}

func DeleteOidc(clientId string) error {
	conn, err := newGrpcConnection()

	if err != nil {
		return fmt.Errorf("failed to initialise connection to dex api: %s", err)
	}
	defer conn.Close()

	client := api.NewDexClient(conn)

	req := &api.DeleteClientReq{
		Id: clientId,
	}

	if resp, err := client.DeleteClient(context.TODO(), req); err != nil || (resp != nil && resp.NotFound) {
		if resp != nil && resp.NotFound {
			return fmt.Errorf("oauth client %s not found", clientId)
		}

		return fmt.Errorf("failed to update oauth client %s", err)
	}

	return nil
}
