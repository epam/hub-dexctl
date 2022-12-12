// Copyright (c) 2022 EPAM Systems, Inc.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/epam/hub-dexctl/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "dex-grpc-client",
	Short: "gRPC client for Dex A Federated OpenID Connect Provider",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&config.Host, "host", "127.0.0.1", "Host address of Dex API")
	rootCmd.PersistentFlags().StringVar(&config.Port, "port", "5557", "Port of Dex API")
	rootCmd.PersistentFlags().StringVar(&config.CaPath, "ca-path", "", "Path to CA certificate file")
	rootCmd.PersistentFlags().StringVar(&config.ClientCrt, "client-crt", "", "Path to client certificate file")
	rootCmd.PersistentFlags().StringVar(&config.ClientKey, "client-key", "", "Path to client key file")
	rootCmd.PersistentFlags().BoolVar(&config.SkipExitCode, "skip-exit-code", false, "Skip exit code on error")
}

func initConfig() {
	viper.SetEnvPrefix("dex_api")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()

	if host := viper.GetString("host"); host != "" {
		config.Host = host
	}

	if port := viper.GetString("port"); port != "" {
		config.Port = port
	}

	if caPath := viper.GetString("ca-path"); caPath != "" {
		config.CaPath = caPath
	}

	if clientCrt := viper.GetString("client-crt"); clientCrt != "" {
		config.ClientCrt = clientCrt
	}

	if clientKey := viper.GetString("client-key"); clientKey != "" {
		config.ClientKey = clientKey
	}
}
