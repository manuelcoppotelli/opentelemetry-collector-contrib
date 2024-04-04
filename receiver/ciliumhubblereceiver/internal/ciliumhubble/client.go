// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ciliumhubble // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver/internal/ciliumhubble"

import (
	"fmt"
)

const (
	DefaultTraceEncoding     = EncodingFlatStringMap
	EncodingJSON             = "JSON"
	EncodingJSONBASE64       = "JSON+base64"
	EncodingFlatStringMap    = "FlatStringMap"
	EncodingSemiFlatTypedMap = "SemiFlatTypedMap"
	EncodingTypedMap         = "TypedMap"
)

func EncodingFormats() []string {
	return []string{
		EncodingJSON,
		EncodingJSONBASE64,
		EncodingFlatStringMap,
		EncodingSemiFlatTypedMap,
	}
}

type EncodingOptions struct {
	Encoding         string `mapstructure:"encoding"`
	TopLevelKeys     bool   `mapstructure:"top_level_keys"`
	LabelsAsMaps     bool   `mapstructure:"labels_as_maps"`
	HeadersAsMaps    bool   `mapstructure:"headers_as_maps"`
	LogPayloadAsBody bool   `mapstructure:"log_payload_as_body"`
}

func (o *EncodingOptions) Validate() error {
	if err := o.validateFormat("trace", EncodingFormats()); err != nil {
		return err
	}
	switch o.Encoding {
	case EncodingJSON, EncodingJSONBASE64:
		if o.TopLevelKeys {
			return fmt.Errorf("option \"TopLevelKeys\" is not compatible with %s encoding", o.Encoding)
		}
	}
	return nil
}

func (o *EncodingOptions) validateFormat(dataType string, formats []string) error {
	if len(o.Encoding) == 0 {
		return fmt.Errorf("encoding format must be set")
	}

	invalidFormat := true
	for _, format := range formats {
		if o.Encoding == format {
			invalidFormat = false
		}
	}
	if invalidFormat {
		return fmt.Errorf("encoding %s is invalid for %s data", o.Encoding, dataType)
	}
	return nil
}
