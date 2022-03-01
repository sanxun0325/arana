//
// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package config

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalTextForProtocolTypeNil(t *testing.T) {
	var protocolType ProtocolType
	var text = []byte("http")
	err := protocolType.UnmarshalText(text)
	assert.True(t, err == nil)
	assert.Equal(t, Http, protocolType)
}

func TestUnmarshalTextForUnrecognizedProtocolType(t *testing.T) {
	var protocolType = Http
	var text = []byte("PostgreSQL")
	err := protocolType.UnmarshalText(text)
	assert.Error(t, err)
}

func TestUnmarshalText(t *testing.T) {
	var protocolType = Http
	var text = []byte("mysql")
	err := protocolType.UnmarshalText(text)
	assert.True(t, err == nil)
	assert.Equal(t, Mysql, protocolType)
}

func TestLoad(t *testing.T) {
	cfg := Load("../../docker/conf/config.yaml")
	assert.Equal(t, Mysql, cfg.Listeners[0].ProtocolType)
	assert.Equal(t, "redirect", cfg.Executors[0].Name)
	assert.Equal(t, 3, len(cfg.DataSources))
	assert.Equal(t, "employees", cfg.DataSources[0].Name)
	assert.Equal(t, Master, cfg.DataSources[0].Role)
	assert.Equal(t, DBMysql, cfg.DataSources[0].Type)
}