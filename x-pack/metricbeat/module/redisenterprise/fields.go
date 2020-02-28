// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package redisenterprise

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "redisenterprise", asset.ModuleFieldsPri, AssetRedisenterprise); err != nil {
		panic(err)
	}
}

// AssetRedisenterprise returns asset data.
// This is the base64 encoded gzipped contents of module/redisenterprise.
func AssetRedisenterprise() string {
	return "eJykjztOBDEQRHOforT5cgAHZFwALrDGroUWHtvqbhBze2Q+q9VoEAEv7E/p1REvXCOURYzNqUPFGAAXr4w43M8N7i6rQwAKLasMl94ibgMAfJ0tdJVsyL1WZmfBWfuCbQYeqG/UmwAoK5Mx4pGeAnAW1mLxM/OIlhbuyU18HYx40v46vic7WpPT5v+E3JsnaQZ/JqSduy5pfiG1AvPkYj5r/CU/uRa+lm698DLcs51s2//wW+jQ/r7+O/UjAAD//wOWl64="
}
