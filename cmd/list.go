// Copyright © 2017 The Kubicorn Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/kris-nova/kubicorn/cutil/logger"
	"github.com/kris-nova/kubicorn/state"
	"github.com/kris-nova/kubicorn/state/fs"
	"github.com/kris-nova/kubicorn/state/git"
	"github.com/kris-nova/kubicorn/state/jsonfs"
	"github.com/spf13/cobra"
	"github.com/minio/minio-go"
	"github.com/kris-nova/kubicorn/state/s3"
)

type ListOptions struct {
	Options
	Profile string
}

var lo = &ListOptions{}

var noHeaders bool

// ListCmd represents the list command
func ListCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "List available states",
		Long:  `List the states available in the _state directory`,
		Run: func(cmd *cobra.Command, args []string) {
			err := RunList(lo)
			if err != nil {
				logger.Critical(err.Error())
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&lo.StateStore, "state-store", "s", strEnvDef("KUBICORN_STATE_STORE", "fs"), "The state store type to use for the cluster")
	cmd.Flags().StringVarP(&lo.StateStorePath, "state-store-path", "S", strEnvDef("KUBICORN_STATE_STORE_PATH", "./_state"), "The state store path to use")
	cmd.Flags().BoolVarP(&noHeaders, "no-headers", "n", false, "Show the list containing names only")

	// s3 flags
	cmd.Flags().StringVar(&lo.S3AccessKey, "s3-access", strEnvDef("KUBICORN_S3_ACCESS_KEY", ""), "The s3 access key.")
	cmd.Flags().StringVar(&lo.S3SecretKey, "s3-secret", strEnvDef("KUBICORN_S3_SECRET_KEY", ""), "The s3 secret key.")
	cmd.Flags().StringVar(&lo.BucketEndpointURL, "s3-endpoint", strEnvDef("KUBICORN_S3_ENDPOINT", ""), "The s3 endpoint url.")
	cmd.Flags().StringVar(&lo.BucketLocation, "s3-location", strEnvDef("KUBICORN_S3_LOCATION", ""), "The s3 bucket location.")
	cmd.Flags().StringVar(&lo.BucketName, "s3-bucket", strEnvDef("KUBICORN_S3_BUCKET", ""), "The s3 bucket name to be used for saving the s3 state for the cluster.")

	return cmd
}

func RunList(options *ListOptions) error {
	options.StateStorePath = expandPath(options.StateStorePath)

	var stateStore state.ClusterStorer
	switch options.StateStore {
	case "fs":
		if !noHeaders {
			logger.Info("Selected [fs] state store")
		}
		stateStore = fs.NewFileSystemStore(&fs.FileSystemStoreOptions{
			BasePath: options.StateStorePath,
		})

	case "git":
		if !noHeaders {
			logger.Info("Selected [git] state store")
		}
		stateStore = git.NewJSONGitStore(&git.JSONGitStoreOptions{
			BasePath: options.StateStorePath,
		})
	case "jsonfs":
		if !noHeaders {
			logger.Info("Selected [jsonfs] state store")
		}
		stateStore = jsonfs.NewJSONFileSystemStore(&jsonfs.JSONFileSystemStoreOptions{
			BasePath: options.StateStorePath,
		})
	case "s3":
		client, err := minio.New(lo.BucketEndpointURL, lo.S3AccessKey, lo.S3SecretKey, true)
		if err != nil {
			return err
		}

		logger.Info("Selected [s3] state store")
		stateStore = s3.NewJSONFS3Store(&s3.JSONS3StoreOptions{
			Client: client,
			BasePath:    options.StateStorePath,
			BucketOptions: &s3.S3BucketOptions{
				EndpointURL: lo.BucketEndpointURL,
				BucketName: lo.BucketName,
				BucketLocation: lo.BucketLocation,
			},
		})
	}

	clusters, err := stateStore.List()
	if err != nil {
		return fmt.Errorf("Unable to list clusters: %v", err)
	}
	for _, cluster := range clusters {
		if !noHeaders {
			logger.Always(cluster)
		} else {
			fmt.Println(cluster)
		}
	}

	return nil
}
