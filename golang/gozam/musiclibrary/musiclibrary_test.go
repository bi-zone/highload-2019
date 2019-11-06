package musiclibrary_test

import (
	"os"
	"path"
	"log"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bi-zone/highload-2019/golang/gozam/models"
	"github.com/bi-zone/highload-2019/golang/gozam/musiclibrary"
)

func trimExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func TestOneTrack(t *testing.T) {
	user := os.Getenv("DBUSER")
	if user == "" {
		log.Fatal("DBUSER is not provided")
		t.Error("DBUSER env var not found")
	}
	dbName := os.Getenv("DBNAME")
	if dbName == "" {
		log.Fatal("DBNAME is not provided")
		t.Error("DBNAME env var not found")
	}

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
