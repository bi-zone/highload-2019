package musiclibrary

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/a.karpov/gozam/fingerprint"
	"github.com/a.karpov/gozam/models"
	"github.com/remeh/sizedwaitgroup"
)

// MusicLibrary is the central structure of the package.
// It is the link for fingerprinting and repository interaction.
type MusicLibrary struct {
	db *sql.DB
}

// Open connects to existing audio repository
func Open(cfg *models.Config) (*MusicLibrary, error) {
	fmt.Printf("Initializing library...\n\n")

	db, err := models.NewDB(cfg)
	if err != nil {
		return nil, err
	}

	return &MusicLibrary{db}, err
}

// Close closes library
func (lib *MusicLibrary) Close() error {
	err := lib.db.Close()
	return err
}

// Index inserts song into library
func (lib *MusicLibrary) Index(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("Index: file not found")
	}

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

	return nil
}

// IndexDir indexes whole directory
func (lib *MusicLibrary) IndexDir(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("IndexDir: invalid directory '%s'", path)
	}

	wg := sizedwaitgroup.New(8)
	for _, f := range files {
		filename := path + "/" + f.Name()
		if filepath.Ext(f.Name()) == ".mp3" {
			wg.Add()
			go func() {
				defer wg.Done()
				lib.Index(filename)
			}()
		}
	}
	wg.Wait()

	return nil
}

// Recognize searches library and returns table
func (lib *MusicLibrary) Recognize(filename string) (result string, err error) {
	fmt.Printf("Recognizing '%s'...\n", filename)

	hashArray, err := fingerprint.Fingerprint(filename)
	if err != nil {
		return "", err
	}

	songName, err := models.Recognize(lib.db, hashArray)

	result = fmt.Sprintf("Best match: %s\n", songName)
	return
}

// RecognizeDir recognizes whole directory
func (lib *MusicLibrary) RecognizeDir(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("IndexDir: invalid directory '%s'", path)
	}

	for _, f := range files {
		filename := path + "/" + f.Name()
		if filepath.Ext(f.Name()) == ".mp3" {
			res, err := lib.Recognize(filename)

			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(res)
		}
	}

	return nil
}

// Delete deletes song from library
func (lib *MusicLibrary) Delete(song string) (affected int64, err error) {
	fmt.Printf("Deleting '%s'...\n", song)
	return models.Delete(lib.db, song)
}
