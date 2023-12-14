

package opa

import (
	"context"
	"time"

	"github.com/khulnasoft/terrasec/pkg/policy"
)

// EngineStats Contains misc stats
type EngineStats struct {
	ruleCount         int
	regoFileCount     int
	metadataFileCount int
	metadataCount     int
	runTime           time.Duration
}

// Engine Implements the policy engine interface
type Engine struct {
	results     policy.EngineOutput
	context     context.Context
	regoFileMap map[string][]byte
	regoDataMap map[string]*policy.RegoData
	stats       EngineStats
}
