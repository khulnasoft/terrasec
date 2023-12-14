

package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/itchyny/gojq"
	"go.uber.org/zap"
)

// JQFilterWithQuery runs jq query on the given input and returns the output
func JQFilterWithQuery(jqQuery string, jsonInput []byte) ([]byte, error) {

	var processed []byte

	// convert read json input into map[string]interface{}
	var input map[string]interface{}
	if err := json.Unmarshal(jsonInput, &input); err != nil {
		return processed, fmt.Errorf("failed to decode input JSON. error: '%v'", err)
	}

	// parse jq query
	query, err := gojq.Parse(jqQuery)
	if err != nil {
		return processed, fmt.Errorf("failed to parse jq query. error: '%v'", err)
	}

	// run jq query on input
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	iter := query.RunWithContext(ctx, input)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			zap.S().Warn("error in processing jq query; error: '%v'", err)
			continue
		}

		jqout, err := json.Marshal(v)
		if err != nil {
			zap.S().Warn("failed to encode jq output into JSON. error: '%v'", err)
			continue
		}
		processed = append(processed, jqout...)
	}

	return processed, nil
}
