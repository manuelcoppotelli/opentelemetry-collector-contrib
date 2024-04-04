// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ciliumhubblereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver"

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver/internal/ciliumhubble"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver/internal/trace"
)

// Config defines the configuration for the receiver.
type Config struct {
	configgrpc.ClientConfig `mapstructure:",squash"`

	BufferSize int `mapstructure:"buffer_size"`

	FlowEncodingOptions FlowEncodingOptions `mapstructure:"flow_encoding_options"`
	IncludeFlowTypes    IncludeFlowTypes    `mapstructure:"include_flow_types"`

	FallbackServiceNamePrefix string        `mapstructure:"fallback_service_name_prefix"`
	TraceCacheWindow          time.Duration `mapstructure:"trace_cache_window"`
	ParseTraceHeaders         bool          `mapstructure:"parse_trace_headers"`
}

type FlowEncodingOptions struct {
	Traces ciliumhubble.EncodingOptions `mapstructure:"traces"`
}

type IncludeFlowTypes struct {
	Traces ciliumhubble.IncludeFlowTypes `mapstructure:"traces"`
}

func createDefaultConfig() component.Config {
	return &Config{
		BufferSize:                2048,
		FallbackServiceNamePrefix: ciliumhubble.OTelAttrServiceNameDefaultPrefix,
		TraceCacheWindow:          trace.DefaultTraceCacheWindow,
		ParseTraceHeaders:         true,
		FlowEncodingOptions: FlowEncodingOptions{
			Traces: ciliumhubble.EncodingOptions{
				Encoding:      ciliumhubble.DefaultTraceEncoding,
				LabelsAsMaps:  true,
				HeadersAsMaps: true,
				TopLevelKeys:  true,
			},
		},
		IncludeFlowTypes: IncludeFlowTypes{
			Traces: ciliumhubble.IncludeFlowTypes{},
		},
	}
}

func (c Config) Validate() error {
	if c.Endpoint == "" {
		return errors.New("Hubble endpoint must be specified")
	}
	if err := c.FlowEncodingOptions.Traces.Validate(); err != nil {
		return err
	}
	if err := c.IncludeFlowTypes.Traces.Validate(); err != nil {
		return err
	}
	return nil
}
