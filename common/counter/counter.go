package counter

import (
	"fmt"
	"sync/atomic"

	"github.com/inhies/go-bytesize"
)

// totalReadBytes is the total number of bytes read
var _totalReadBytes uint64 = 0

// totalWrittenBytes is the total number of bytes written
var _totalWrittenBytes uint64 = 0

// GetReadBytes returns the number of bytes read
func GetReadBytes() uint64 {
	return _totalReadBytes
}

// GetWrittenBytes returns the number of bytes written
// GetWrittenBytes returns the number of bytes written
func GetWrittenBytes() uint64 {
	return _totalWrittenBytes
}

// PrintBytes returns the bytes info
// PrintBytes returns the bytes info
func PrintBytes() string {
	return fmt.Sprintf("download %v upload %v", bytesize.New(float64(GetWrittenBytes())).String(), bytesize.New(float64(GetReadBytes())).String())
}

//IncrWrittenBytes increments the number of bytes written
//IncrWrittenBytes increments the number of bytes written
func IncrWrittenBytes(n int) {
	atomic.AddUint64(&_totalWrittenBytes, uint64(n))
}

// IncrReadBytes increments the number of bytes read
// IncrReadBytes increments the number of bytes read
func IncrReadBytes(n int) {
	atomic.AddUint64(&_totalReadBytes, uint64(n))
}
