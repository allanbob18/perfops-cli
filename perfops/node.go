// Copyright 2017 The PerfOps-CLI Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package perfops

type (
	// Continent contains information about a continent.
	Continent struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		ISO  string `json:"iso"`
	}

	// Country contains information about a country.
	Country struct {
		ID         int        `json:"id"`
		Name       string     `json:"name"`
		ISO        string     `json:"iso"`
		ISONumeric string     `json:"iso_numeric"`
		Continent  *Continent `json:"continent,omitempty"`
	}

	// Node contains informatin about a test node.
	Node struct {
		ID        int      `json:"id"`
		Latitude  float64  `json:"latitude"`
		Longitude float64  `json:"longitude"`
		City      string   `json:"city"`
		SubRegion string   `json:"sub_region"`
		Country   *Country `json:"country,omitempty"`
	}
)
