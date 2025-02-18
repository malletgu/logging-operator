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

package util

import (
	"sort"

	"github.com/iancoleman/orderedmap"
)

// MergeLabels merge into map[string]string map
func MergeLabels(labelGroups ...map[string]string) map[string]string {
	mergedLabels := make(map[string]string)
	for _, labels := range labelGroups {
		for k, v := range labels {
			mergedLabels[k] = v
		}
	}
	return mergedLabels
}

// IntPointer converts int32 to *int32
func IntPointer(i int32) *int32 {
	return &i
}

// BoolPointer converts bool to *bool
func BoolPointer(b bool) *bool {
	return &b
}

func OrderedStringMap(original map[string]string) *orderedmap.OrderedMap {
	o := orderedmap.New()
	for k, v := range original {
		o.Set(k, v)
	}
	o.SortKeys(sort.Strings)
	return o
}
