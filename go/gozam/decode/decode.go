package decode

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/jfreymuth/oggvorbis"
	"github.com/kisasexypantera94/go-mpg123/mpg123"
	"github.com/youpy/go-wav"
)

/*
#include <mpg123.h>
#cgo LDFLAGS: -lmpg123
*/
import "C"

const chunkSize = 2048

// Mp3 decodes mp3 files using `libmpg123`
func Mp3(filename string) []float64 {
	decoder, err := mpg123.NewDecoder("", C.MPG123_MONO_MIX|C.MPG123_FORCE_FLOAT)
	checkErr(err)

	err = decoder.Open(filename)
	checkErr(err)
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

		binary.Read(bytes.NewBuffer(buf), binary.LittleEndian, tmp)
		for i := 0; i < len(tmp); i++ {
			mono := tmp[i]
			pcm64 = append(pcm64, (float64)(mono))
		}
	}

	decoder.Delete()
	return pcm64
}

// Ogg decodes ogg files
func Ogg(filename string) []float64 {
	f, _ := os.Open(filename)
	defer f.Close()
	var r io.Reader
	r = f
	pcm32, _, _ := oggvorbis.ReadAll(r)
	pcm64 := make([]float64, len(pcm32))
	for i := 0; i < len(pcm32); i++ {
		pcm64[i] = (float64)(pcm32[i])
	}

	return pcm64
}

// Wav decodes wav files
func Wav(filename string) []float64 {
	file, _ := os.Open(filename)
	defer file.Close()
	reader := wav.NewReader(file)

	var pcm []float64
	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}

		for _, sample := range samples {
			pcm = append(pcm, (reader.FloatValue(sample, 0)+reader.FloatValue(sample, 1))/2)
		}
	}

	return pcm
}

// Decode mp3, wav or ogg files
func Decode(filename string) (pcm []float64, err error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("Decode: file not found")
	}

	var pcm64 []float64
	switch filepath.Ext(filename) {
	case ".mp3":
		pcm64 = Mp3(filename)
	case ".wav":
		pcm64 = Wav(filename)
	case ".ogg":
		pcm64 = Ogg(filename)
	default:
		return nil, fmt.Errorf("Decode: invalid file")
	}
	return pcm64, nil
}

func checkErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
