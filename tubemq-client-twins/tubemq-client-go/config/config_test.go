// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseAddress(t *testing.T) {
	address := "127.0.0.1:9092,127.0.0.1:9093?topics=Topic1@12312323,1212;Topic2@121212,2321323&group=Group&tlsEnable=false&msgNotFoundWait=10000&heartbeatMaxRetryTimes=6"
	topicFilters := make(map[string][]string)
	topicFilters["Topic1"] = []string{"12312323", "1212"}
	topicFilters["Topic2"] = []string{"121212", "2321323"}
	c, err := ParseAddress(address)
	assert.Nil(t, err)
	assert.Equal(t, c.Consumer.Masters, "127.0.0.1:9092,127.0.0.1:9093")
	assert.Equal(t, c.Consumer.Topics, []string{"Topic1", "Topic2"})
	assert.Equal(t, c.Consumer.TopicFilters, topicFilters)
	assert.Equal(t, c.Consumer.Group, "Group")
	assert.Equal(t, c.Consumer.MsgNotFoundWait, 10000*time.Millisecond)

	assert.Equal(t, c.Net.TLS.Enable, false)

	assert.Equal(t, c.Heartbeat.MaxRetryTimes, 6)

	address = ""
	_, err = ParseAddress(address)
	assert.NotNil(t, err)

	address = "127.0.0.1:9092,127.0.0.1:9093?topics=Topic&ttt"
	_, err = ParseAddress(address)
	assert.NotNil(t, err)

	address = "127.0.0.1:9092,127.0.0.1:9093?topics=Topic&ttt=ttt"
	_, err = ParseAddress(address)
	assert.NotNil(t, err)
}