package musiclibrary_test

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/glumpo/highload-2019/golang/gozam/models"
	"github.com/glumpo/highload-2019/golang/gozam/musiclibrary"
)

func trimExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func TestOneTrack(t *testing.T) {
	cfg := models.Config{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASSWORD"),
		DBname:   os.Getenv("DBNAME"),
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
	}

	musicLib, err := musiclibrary.Open(cfg)
	if err != nil {
		t.Error(err)
	}
	defer musicLib.Close()

	originPath := "testdata/origin/kitay brusnika himky les (forest).mp3"
	samplePath := "testdata/sample/forest.mp3"

	originName := trimExtension(filepath.Base(originPath))

	_ = musicLib.Index(originPath)
	result, err := musicLib.Recognize(samplePath)
	if err != nil {
		t.Error(err)
	}

	if result != originName {
		t.Error("Wrong")
	}
}
