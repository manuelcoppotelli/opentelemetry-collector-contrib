// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ciliumhubblereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
)

type ciliumhubbleReceiver struct {
}

func newCiliumHubbleTraceReceiver(_ *Config, _ consumer.Traces, _ *zap.Logger) (*ciliumhubbleReceiver, error) {
	return &ciliumhubbleReceiver{}, nil
}

func (r *ciliumhubbleReceiver) Start(_ context.Context, _ component.Host) error {
	return nil
}

func (r *ciliumhubbleReceiver) Shutdown(_ context.Context) error {
	return nil
}
