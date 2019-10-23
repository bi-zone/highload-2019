package musiclibrary

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/glumpo/highload-2019/golang/gozam/fingerprint"
	"github.com/glumpo/highload-2019/golang/gozam/models"
)

// MusicLibrary is the central structure of the package.
// It is the link for fingerprinting and repository interaction.
type MusicLibrary struct {
	db *sql.DB
}

// Open connects to existing audio repository
func Open(cfg models.Config) (*MusicLibrary, error) {
	fmt.Printf("Initializing library...\n\n")

	db, err := models.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	return &MusicLibrary{db}, err
}

// Close closes library
func (lib MusicLibrary) Close() error {
	err := lib.db.Close()
	return err
}

// Index inserts song into library
func (lib MusicLibrary) Index(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("Index: file not found")
	}

	// NOTE: name of the song doesnt contain .mp3 nor full path
	fmt.Printf("Indexing '%s'...\n", filename)
	dotIdx := strings.LastIndex(filename, ".")
	slashIdx := strings.LastIndex(filename, "/")
	if dotIdx == -1 {
		return fmt.Errorf("Index: invalid file '%s'", filename)
	}
	songName := filename[slashIdx+1 : dotIdx]

	hashArray, err := fingerprint.Fingerprint(filename)
	if err != nil {
		return err
	}

	err = models.Index(lib.db, songName, hashArray)

	return err
}

// Recognize searches library and returns table
func (lib MusicLibrary) Recognize(filename string) (result string, err error) {
	fmt.Printf("Recognizing '%s'...\n", filename)

	hashArray, err := fingerprint.Fingerprint(filename)
	if err != nil {
		return "", err
	}

	songName, err := models.Recognize(lib.db, hashArray)

	result = fmt.Sprintf("Best match: %s\n", songName)
	return
}

// Delete deletes song from library
func (lib MusicLibrary) Delete(song string) (affected int64, err error) {
	fmt.Printf("Deleting '%s'...\n", song)
	return models.Delete(lib.db, song)
}
