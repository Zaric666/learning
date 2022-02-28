package logger

import "testing"

func TestJSONLogger(t *testing.T) {
	logger, err := NewLogger(
		WithInfoLevel(),
	)
	if err != nil {
		t.Fatal(err)
	}
	logger.info("hello")
}
