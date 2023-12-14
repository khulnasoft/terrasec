package dblogs

import (
	"testing"
	"time"
)

func TestLogs(t *testing.T) {

	var logger = WebhookScanLogger{
		Test: true,
	}
	defer logger.ClearDbFilePath()

	// insert a new db record
	var log = WebhookScanLog{
		UID:                "myUID",
		Request:            "MyRequest",
		CreatedAt:          time.Now(),
		Allowed:            true,
		DeniableViolations: "MyViolations",
		ViolationsSummary:  "ViolationsSummary",
	}

	t.Run("initialize db", func(t *testing.T) {

		// no logs exist in db, should return 0 logs
		fetchedLogs, err := logger.FetchLogs()
		if len(fetchedLogs) > 0 {
			t.Errorf("no logs should exist in db; got: '%v' logs", len(fetchedLogs))
		}
		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
		}
	})

	t.Run("insert db record", func(t *testing.T) {
		if err := logger.Log(log); err != nil {
			t.Errorf("unexpected error: '%v'", err)
		}
	})

	t.Run("one db record", func(t *testing.T) {
		fetchedLogs, err := logger.FetchLogs()
		if err != nil {
			t.Errorf("unexpected error: '%v'", err)
		}
		if len(fetchedLogs) != 1 {
			t.Errorf("db has one log, got: '%v' logs", len(fetchedLogs))
		}
	})
}
