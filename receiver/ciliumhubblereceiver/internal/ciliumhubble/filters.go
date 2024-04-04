// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ciliumhubble // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver/internal/ciliumhubble"

import (
	"fmt"

	monitorAPI "github.com/cilium/cilium/pkg/monitor/api"
)

type IncludeFlowTypes []string

func (it IncludeFlowTypes) Validate() error {
	for _, t := range it {
		if _, ok := monitorAPI.MessageTypeNames[t]; ok {
			continue
		}
		switch t {
		case "":
			return fmt.Errorf("type filter cannot be an empty string")

		case "*", "all":
			if len(it) != 1 {
				return fmt.Errorf("type filter %q can only be specified on its own", t)
			}
		default:
			return fmt.Errorf("unknown type filter %q", t)
		}
	}
	return nil
}
