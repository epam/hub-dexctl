// Copyright (c) 2022 EPAM Systems, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package dex

import (
	"context"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/dexidp/dex/api/v2"
)

func CreatePassword(email string, password string) error {
	conn, err := newGrpcConnection()

	if err != nil {
		return fmt.Errorf("failed to initialise connection to dex api: %s", err)
	}
	defer conn.Close()

	client := api.NewDexClient(conn)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to bcrypt password: %s", err)
	}

	userId := base64.StdEncoding.EncodeToString([]byte(email))

	p := &api.Password{
		Email:    email,
		Hash:     hash,
		Username: email,
		UserId:   userId,
	}

	req := &api.CreatePasswordReq{
		Password: p,
	}

	if resp, err := client.CreatePassword(context.TODO(), req); err != nil || (resp != nil && resp.AlreadyExists) {
		if resp != nil && resp.AlreadyExists {
			return fmt.Errorf("password for %s already exists", email)
		}

		return fmt.Errorf("failed to create password %s", err)
	}

	return nil
}

func UpdatePassword(email string, newPassword string, newEmail string) error {
	conn, err := newGrpcConnection()

	if err != nil {
		return fmt.Errorf("failed to initialise connection to dex api: %s", err)
	}
	defer conn.Close()

	client := api.NewDexClient(conn)

	hash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to bcrypt password: %s", err)
	}

	req := &api.UpdatePasswordReq{
		Email:       email,
		NewHash:     hash,
		NewUsername: newEmail,
	}

	if resp, err := client.UpdatePassword(context.TODO(), req); err != nil || (resp != nil && resp.NotFound) {
		if resp != nil && resp.NotFound {
			return fmt.Errorf("password for %s not found", email)
		}

		return fmt.Errorf("failed to update password %s", err)
	}
	return nil
}

func DeletePassword(email string) error {
	conn, err := newGrpcConnection()

	if err != nil {
		return fmt.Errorf("failed to initialise connection to dex api: %s", err)
	}
	defer conn.Close()

	client := api.NewDexClient(conn)

	req := &api.DeletePasswordReq{
		Email: email,
	}
	if resp, err := client.DeletePassword(context.TODO(), req); err != nil || (resp != nil && resp.NotFound) {
		if resp != nil && resp.NotFound {
			return fmt.Errorf("password for %s not found", email)
		}

		return fmt.Errorf("failed to update password %s", err)
	}

	return nil
}
