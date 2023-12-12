package repository

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pet struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
type File struct{}

type FileRep interface {
	PetsFromFile() ([]Pet, error)
	PetsToFile() error
}

func (f *File) PetsFromFile() ([]Pet, error) {
	pets := make([]Pet, 0)
	stream, err := os.ReadFile("peoples.json")
	if os.IsNotExist(err) {
		file, er := os.Create("peoples.json")
		if er != nil {
			return pets, fmt.Errorf("ошибка при создании файла err: %v", err)
		}
		file.Write([]byte("[{}]"))
		stream = []byte("[{}]")
	}
	if err := json.Unmarshal(stream, &pets); err != nil {
		return pets, fmt.Errorf("ошибка при десериализации данных err: %v", err)
	}
	return pets, nil
}

func (f *File) PetsToFile(pets []Pet) error {
	data, er := json.Marshal(pets)
	if er != nil {
		return fmt.Errorf("ошибка при сериализации данных err: %v", er)
	}
	if err := os.WriteFile("peoples.json", data, os.ModeAppend); err != nil {
		return fmt.Errorf("ошибка при записи данных в файл err: %v", err)
	}
	return nil
}
