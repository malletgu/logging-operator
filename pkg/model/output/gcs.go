// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package output

import (
	"github.com/banzaicloud/logging-operator/pkg/model/secret"
	"github.com/banzaicloud/logging-operator/pkg/model/types"
)

// +kubebuilder:object:generate=true
type GCSOutput struct {
	// Project identifier for GCS
	Project string `json:"project"`
	// Path of GCS service account credentials JSON file
	Keyfile string `json:"keyfile,omitempty"`
	// GCS service account credentials in JSON format
	// +docLink:"Secret,./secret.md"
	CredentialsJson *secret.Secret `json:"credentials_json,omitempty"`
	// Number of times to retry requests on server error
	ClientRetries int `json:"client_retries,omitempty"`
	// Default timeout to use in requests
	ClientTimeout int `json:"client_timeout,omitempty"`
	// Name of a GCS bucket
	Bucket string `json:"bucket"`
	// Format of GCS object keys (default: %{path}%{time_slice}_%{index}.%{file_extension})
	ObjectKeyFormat string `json:"object_key_format,omitempty"`
	// Path prefix of the files on GCS
	Path string `json:"path,omitempty"`
	// Archive format on GCS: gzip json text (default: gzip)
	StoreAs string `json:"store_as,omitempty"`
	// Enable the decompressive form of transcoding
	Transcoding bool `json:"transcoding,omitempty"`
	// Create GCS bucket if it does not exists (default: true)
	AutoCreateBucket bool `json:"auto_create_bucket,omitempty"`
	// Max length of `%{hex_random}` placeholder(4-16) (default: 4)
	HexRandomLength int `json:"hex_random_length,omitempty"`
	// Overwrite already existing path (default: false)
	Overwrite bool `json:"overwrite,omitempty"`
	// Permission for the object in GCS: auth_read owner_full owner_read private project_private public_read
	// +kubebuilder:validation:enum=auth_read,owner_full,owner_read,private,project_private,public_read
	Acl string `json:"acl,omitempty"`
	// Storage class of the file: dra nearline coldline multi_regional regional standard
	// +kubebuilder:validation:enum=dra,nearline,coldline,multi_regional,regional,standard
	StorageClass string `json:"storage_class,omitempty"`
	// Customer-supplied, AES-256 encryption key
	EncryptionKey string `json:"encryption_key,omitempty"`
	// User provided web-safe keys and arbitrary string values that will returned with requests for the file as "x-goog-meta-" response headers.
	// +docLink:"Object Metadata,#ObjectMetadata"
	ObjectMetadata []ObjectMetadata `json:"object_metadata,omitempty"`
	// +docLink:"Format,./format.md"
	Format *Format `json:"format,omitempty"`
	// +docLink:"Buffer,./buffer.md"
	Buffer *Buffer `json:"buffer,omitempty"`
}

func (g *GCSOutput) ToDirective(secretLoader secret.SecretLoader, id string) (types.Directive, error) {
	pluginType := "gcs"
	gcs := &types.OutputPlugin{
		PluginMeta: types.PluginMeta{
			Type:      pluginType,
			Directive: "match",
			Tag:       "**",
			Id:        id + "-" + pluginType,
		},
	}
	if params, err := types.NewStructToStringMapper(secretLoader).StringsMap(g); err != nil {
		return nil, err
	} else {
		gcs.Params = params
	}
	if g.Buffer != nil {
		if buffer, err := g.Buffer.ToDirective(secretLoader, ""); err != nil {
			return nil, err
		} else {
			gcs.SubDirectives = append(gcs.SubDirectives, buffer)
		}
	}
	if g.Format != nil {
		if format, err := g.Format.ToDirective(secretLoader, ""); err != nil {
			return nil, err
		} else {
			gcs.SubDirectives = append(gcs.SubDirectives, format)
		}
	}
	if len(g.ObjectMetadata) > 0 {
		for _, metadata := range g.ObjectMetadata {
			if meta, err := metadata.ToDirective(secretLoader, ""); err != nil {
				return nil, err
			} else {
				gcs.SubDirectives = append(gcs.SubDirectives, meta)
			}
		}
	}
	return gcs, nil
}

type ObjectMetadata struct {
	// Key
	Key string `json:"key"`
	// Value
	Value string `json:"value"`
}

func (o *ObjectMetadata) ToDirective(secretLoader secret.SecretLoader, id string) (types.Directive, error) {
	return types.NewFlatDirective(types.PluginMeta{
		Directive: "object_metadata",
	}, o, secretLoader)
}
