package trace

import (
	"bytes"
	"log"
	"strings"
	"testing"
	"time"
)

func ExampleExecuteTime() {
	ExecuteTime("step 1.", func() {
		time.Sleep(200 * time.Millisecond)
	})

	ExecuteTime("step 2.", func() {
		time.Sleep(300 * time.Millisecond)
	})

	ExecuteTime("step 3.", func() {
		time.Sleep(400 * time.Millisecond)
	})

	// Output:
}

func TestExecuteTime(t *testing.T) {
	var buf bytes.Buffer
	orig := log.Writer()
	log.SetOutput(&buf)
	defer log.SetOutput(orig)

	ExecuteTime("unit test", func() {
		time.Sleep(10 * time.Millisecond)
	})

	out := buf.String()
	if !strings.Contains(out, "[unit test] elapsed=") || !strings.Contains(out, "ms") {
		t.Errorf("log output format incorrect: %q", out)
	}
}
