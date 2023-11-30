package docker

import (
	"dsync.io/gco/agent/internal/files"
	"dsync.io/gco/agent/pkg/feature"
	"dsync.io/gco/agent/pkg/resource"
	"fmt"
)

// createFluentBitContainer creates an internal container definition based on the feature.FluentBit configuration.
func (p *Provider) createFluentBitContainer(fb *feature.FluentBit) (*internalContainer, error) {
	version := "latest"
	if fb.Version != "" {
		version = fb.Version
	}

	// create volume binding
	config := fb.CreateConfig()
	configLocation, err := files.WriteConfigFileFromString(config, "fluent-bit.conf")
	if err != nil {
		return nil, err
	}

	ic := &internalContainer{
		name:   fmt.Sprintf("%s.%s", labelPrefix, fb.Name()),
		labels: make(map[string]string),
		image:  fmt.Sprintf("cr.fluentbit.io/fluent/fluent-bit:%s", version),
		ports: []containerPort{
			newContainerPort(24224, 24224, "tcp"),
		},
		volumes: []volumeMount{
			{destination: "/fluent-bit/etc/fluent-bit.conf", source: configLocation, readonly: true},
		},
		pullPolicy: whenNotPresentPolicy,
	}

	// add container labels
	ic.addLabel(kindLabel(resource.FeatureKind))
	ic.addLabel(featureLabel(fb.Name()))
	ic.addLabel(configLabel(feature.EncodeFeature(fb)))

	return ic, nil
}