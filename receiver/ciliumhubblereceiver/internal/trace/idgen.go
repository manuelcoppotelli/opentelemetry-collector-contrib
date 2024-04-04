// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package trace // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver/internal/trace"

import (
	"time"
)

const (
	DefaultTraceCacheWindow = 20 * time.Minute
)
