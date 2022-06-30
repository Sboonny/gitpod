// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/heroku/docker-registry-client/registry"
	"github.com/spf13/cobra"
)

var checkOpts struct {
	Username      string
	Password      string
	ServerAddress string
	InCluster     bool
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks registry connection",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !checkOpts.InCluster {
			serverAddress, err := url.Parse(checkOpts.ServerAddress)
			if err != nil {
				return err
			}
			if serverAddress.Scheme == "" {
				// If no scheme, default to HTTPS
				serverAddress.Scheme = "https"
			}
			fmt.Println(serverAddress)
			if strings.HasPrefix(serverAddress.Path, "gcr.io") && serverAddress.Path != "gcr.io" {
				// GCR must use gcr.io else the validation returns false
				return fmt.Errorf("google container registries must use the address gcr.io, not %s", serverAddress.Path)
			}

			_, err = registry.New(serverAddress.String(), checkOpts.Username, checkOpts.Password)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringVarP(&checkOpts.Username, "username", "u", "", "Registry username")
	checkCmd.Flags().StringVarP(&checkOpts.Password, "password", "p", "", "Registry password")
	checkCmd.Flags().StringVarP(&checkOpts.ServerAddress, "server-address", "s", "", "Registry server address")
	checkCmd.Flags().BoolVar(&checkOpts.InCluster, "in-cluster", false, "Registry in-cluster")
}
