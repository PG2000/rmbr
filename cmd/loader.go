package cmd

import (
	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Repository interface {
	Load() (*RmbrNotes, error)
	Write(applications *RmbrNotes) error
	New(note RmbrNote) error
}

type FileSystemRepository struct {
}

func (m FileSystemRepository) New(note RmbrNote) error {
	notes, _ := m.Load()

	notes.Notes = append(notes.Notes, note)

	e := m.Write(notes)

	if e != nil {
		log.Printf("Error while creating new note")
	}

	return nil
}

func (m FileSystemRepository) Load() (*RmbrNotes, error) {

	path := getFullFilePath()
	bytes, err := ioutil.ReadFile(filepath.FromSlash(path))

	if err != nil {
		log.Print("rmbr.yml not found")
		return &RmbrNotes{}, err
	}

	notes, err := Unmarshal(string(bytes))

	if err != nil {
		log.Printf("error while parsing rmbr.yml")
		return &RmbrNotes{}, err
	}

	return &notes, nil
}

func NewFileSystemLoader() (Repository, error) {

	path, filename := getFilePath()
	ensureFileExists(path, filename)
	loader := &FileSystemRepository{}
	return loader, nil
}

func ensureFileExists(path string, filename string) {
	_, e := os.Stat(path)

	if e != nil {
		log.Printf("Directory not found %s", e.Error())
		_ = os.MkdirAll(path, os.ModePerm)
	}

	filePath := filepath.FromSlash(path + "/" + filename)
	_, e = os.Stat(filePath)

	if e != nil {
		log.Printf("File not found %s", e.Error())
		os.Create(filePath)
	}

}

func getFilePath() (string, string) {
	home, e := homedir.Dir()
	if e != nil {
		log.Println(e)
	}

	return filepath.FromSlash(home + "/"), ".rmbr.yml"
}

func Unmarshal(yamlString string) (RmbrNotes, error) {
	t := RmbrNotes{}

	err := yaml.Unmarshal([]byte(yamlString), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return t, err
}

func (m FileSystemRepository) Write(applications *RmbrNotes) error {
	yamlBytes, err := Marshal(applications)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	err = ioutil.WriteFile(getFullFilePath(), yamlBytes, 0644)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}

func getFullFilePath() string {
	path, filename := getFilePath()
	return filepath.FromSlash(path + "/" + filename)
}

func Marshal(applications *RmbrNotes) ([]byte, error) {
	out, err := yaml.Marshal(&applications)
	return out, err
}
