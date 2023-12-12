package repository

import (
	"testing"
)

func TestPetsFromFile(t *testing.T) {
	repository := &File{}
	var pets []Pet
	var err error
	if pets, err = repository.PetsFromFile(); err != nil {
		t.Errorf(err.Error())
	}
	t.Log(pets)
}

func TestPetsToFile(t *testing.T) {
	repository := &File{}
	pets := []Pet{{
		Name:    "Sharik",
		Surname: "Zagumennikov",
	},
	}
	if err := repository.PetsToFile(pets); err != nil {
		t.Errorf(err.Error())
	}
	t.Log(pets)
}
