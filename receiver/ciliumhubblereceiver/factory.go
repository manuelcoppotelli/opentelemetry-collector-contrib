// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ciliumhubblereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciliumhubblereceiver/internal/metadata"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithTraces(createTracesReceiver, metadata.TracesStability),
	)
}

func createTracesReceiver(_ context.Context, settings receiver.CreateSettings, cc component.Config, consumer consumer.Traces) (receiver.Traces, error) {
	return newCiliumHubbleTraceReceiver(cc.(*Config), consumer, settings.Logger)
}
