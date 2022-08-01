// Copyright (c) 2022 EPAM Systems, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package dex

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/agilestacks/dexctl/config"
)

func newGrpcConnection() (*grpc.ClientConn, error) {
	var creds credentials.TransportCredentials

	if config.CaPath != "" {
		cPool := x509.NewCertPool()
		caCert, err := os.ReadFile(config.CaPath)
		if err != nil {
			return nil, fmt.Errorf("invalid CA crt file: %s", config.CaPath)
		}
		if !cPool.AppendCertsFromPEM(caCert) {
			return nil, fmt.Errorf("failed to parse CA crt")
		}

		if config.ClientCrt != "" && config.ClientKey != "" {
			clientCert, err := tls.LoadX509KeyPair(config.ClientCrt, config.ClientKey)
			if err != nil {
				return nil, fmt.Errorf("invalid client crt file: %s", config.ClientCrt)
			}

			clientTLSConfig := &tls.Config{
				RootCAs:      cPool,
				Certificates: []tls.Certificate{clientCert},
			}

			creds = credentials.NewTLS(clientTLSConfig)
		} else {
			creds, err = credentials.NewClientTLSFromFile(config.CaPath, "")
			if err != nil {
				return nil, fmt.Errorf("failed to load CA crt: %s", err)
			}
		}
	} else {
		creds = insecure.NewCredentials()
	}

	target := fmt.Sprintf("%s:%s", config.Host, config.Port)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("dial: %v", err)
	}
	return conn, nil
}
