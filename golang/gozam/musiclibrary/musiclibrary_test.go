package musiclibrary_test

import (
	"os"
	"testing"

	"github.com/glumpo/highload-2019/golang/gozam/models"
	"github.com/glumpo/highload-2019/golang/gozam/musiclibrary"
)

func TestOneTrack(t *testing.T) {
	cfg := &models.Config{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASSWORD"),
		DBname:   os.Getenv("DBNAME"),
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
	}

	musicLib, _ := musiclibrary.Open(cfg)
	defer musicLib.Close()

	name := "../samples/forest.mp3"
	_ = musicLib.Index(name)
	result, err := musicLib.Recognize(name)
	if err != nil {
		t.Error(err)
	}

	if result != name {
		t.Error("Wrong")
	}
}
