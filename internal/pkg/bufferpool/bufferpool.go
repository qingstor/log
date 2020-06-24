// Package bufferpool holds the buffer pool used in log.
package bufferpool

import (
	"github.com/Xuanwo/go-bufferpool"
)

var (
	// Log entry mostly lower than 1K size.
	pool = bufferpool.New(1024)
	// Get retrieves a buffer from the pool, creating one if necessary.
	Get = pool.Get
)
