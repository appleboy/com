package trace

import (
	"log"
	"time"
)

func ExecuteTime(title string, fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	log.Printf("[%s] elapsed=%dms", title, elapsed.Milliseconds())
}
