// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ciliumhubblereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver"

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap/confmaptest"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver/internal/ciliumhubble"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver/internal/metadata"
)

func TestConfig_Validate_Invalid(t *testing.T) {
	cfg := Config{}
	assert.Error(t, cfg.Validate())
}

func TestConfig_Validate_Valid(t *testing.T) {
	cfg := Config{
		ClientConfig: configgrpc.ClientConfig{
			Endpoint: "hubble-relay.cilium.svc.cluster.local:80",
		},
		FlowEncodingOptions: FlowEncodingOptions{
			Traces: ciliumhubble.EncodingOptions{
				Encoding: "JSON",
			},
		},
	}
	assert.NoError(t, cfg.Validate())
}

func TestLoadConfig(t *testing.T) {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)

	tests := []struct {
		id           component.ID
		expected     component.Config
		errorMessage string
	}{
		{
			id:           component.NewIDWithName(metadata.Type, ""),
			errorMessage: "Hubble endpoint must be specified",
		},
		{
			id:           component.NewIDWithName(metadata.Type, "1"),
			errorMessage: "encoding format must be set; encoding format must be set",
		},
		{
			id: component.NewIDWithName(metadata.Type, "2"),
			expected: &Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: "unix:///var/run/cilium/hubble.sock",
					TLSSetting: configtls.ClientConfig{
						TLSSetting: configtls.TLSSetting{
							CAFile:   "../testdata/certs/ca.pem",
							CertFile: "../testdata/certs/test-client.pem",
							KeyFile:  "../testdata/certs/test-client-key.pem",
						},
					},
				},
				BufferSize:                2048,
				FallbackServiceNamePrefix: ciliumhubble.OTelAttrServiceNameDefaultPrefix,
				TraceCacheWindow:          1 * time.Hour,
				ParseTraceHeaders:         false,
				FlowEncodingOptions: FlowEncodingOptions{
					Traces: ciliumhubble.EncodingOptions{
						Encoding:      ciliumhubble.EncodingJSON,
						TopLevelKeys:  false,
						LabelsAsMaps:  true,
						HeadersAsMaps: true,
					},
				},
				IncludeFlowTypes: IncludeFlowTypes{
					Traces: ciliumhubble.IncludeFlowTypes{
						"l7",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.id.String(), func(t *testing.T) {
			factory := NewFactory()
			cfg := factory.CreateDefaultConfig()

			sub, err := cm.Sub(tt.id.String())
			require.NoError(t, err)
			require.NoError(t, component.UnmarshalConfig(sub, cfg))

			if tt.errorMessage != "" {
				assert.EqualError(t, component.ValidateConfig(cfg), tt.errorMessage)
				return
			}

			assert.NoError(t, component.ValidateConfig(cfg))
			assert.Equal(t, tt.expected, cfg)
		})
	}
}
