package decode

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kisasexypantera94/go-mpg123/mpg123"
)

/*
#include <mpg123.h>
#cgo LDFLAGS: -lmpg123
*/
import "C"

const chunkSize = 2048

// Mp3 decodes mp3 files using `libmpg123`
func Mp3(filename string) ([]float64, error) {
	decoder, err := mpg123.NewDecoder("", C.MPG123_MONO_MIX|C.MPG123_FORCE_FLOAT)
	if err != nil {
		return nil, err
	}

	err = decoder.Open(filename)
	if err != nil {
		return nil, err
	}
	defer decoder.Close()

	decoder.GetFormat()
	var pcm64 []float64
	tmp := make([]float32, chunkSize/4)
	for {
		buf := make([]byte, chunkSize)
		_, err := decoder.Read(buf)

		if err != nil {
			break
		}

		err = binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, tmp)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(tmp); i++ {
			mono := tmp[i]
			pcm64 = append(pcm64, (float64)(mono))
		}
	}

	decoder.Delete()
	return pcm64, nil
}

// Decode file
func Decode(filename string) (pcm []float64, err error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("Decode: file not found")
	}

	var pcm64 []float64
	switch filepath.Ext(filename) {
	case ".mp3":
		pcm64, err = Mp3(filename)
	default:
		return nil, fmt.Errorf("Decode: invalid file")
	}
	if err != nil {
		return nil, err
	}
	return pcm64, nil
}
