package fingerprint

import (
	"errors"
	"fmt"
	"os"

	"github.com/bi-zone/highload-2019/golang/gozam/decode"
)

// Fingerprint constructs fingerprint for song and returns hash
func Fingerprint(filename string) (hashArray []int, err error) {
	// []int type for hash surely may be chanched
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("Fingerprint: file not found")
	}

	// pcm64 is from pulse code modulation, 64bit float
	pcm64, err := decode.Decode(filename)
	_ = pcm64 // Avoid declared and not used

	// TODO: Implement

	return nil, errors.New("Not Implemented")
}
