/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package config

import (
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	viperutil "github.com/hyperledger-labs/fabric-smart-client/platform/view/core/config/viper"
)

func TestLoad(t *testing.T) {
	v := viper.New()
	v.SetConfigName("core")
	v.AddConfigPath("./testdata")
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	require.NoError(t, v.ReadInConfig())

	var networks []Network
	require.NoError(t, viperutil.EnhancedExactUnmarshal(v, "fabric.networks", &networks))

	assert.Equalf(t, 1, len(networks), "expected len to be 1")
	assert.Equalf(t, "default", networks[0].Name, "expected default name")

}
