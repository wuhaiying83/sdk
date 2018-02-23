// Copyright 2017 luoji

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package client

import "sync"

type MQClient struct {
	index       int32
	clientId    string
	cfg         Config
	producers   map[string]producerInner // key: group
	producersMu sync.RWMutex             //
	consumers   map[string]consumerInner // key: group
	consumersMu sync.RWMutex             //
	clientAPI   *mqClientAPI
}