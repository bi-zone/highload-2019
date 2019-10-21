package musiclibrary_test

import (
	"os"
	"testing"

	"github.com/kisasexypantera94/khalzam/musiclibrary"
)

func TestOneTrack(t *testing.T) {
	cfg := &musiclibrary.Config{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASSWORD"),
		DBname:   os.Getenv("DBNAME"),
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
	}

	musicLib, _ := musiclibrary.Open(cfg)
	defer musicLib.Close()

	name := "китай брусника - сломался антигравитационный двигатель, решили посидеть у костра (bonus track).mp3"
	musicLib.Index(name)
	result, err := musicLib.Recognize(name)
	if err != nil {
		t.Error(err)
	}

	if result != name {
		t.Error("Wrong")
	}
}
