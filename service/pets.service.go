package service

import (
	"encoding/json"
	"fmt"
	"log"

	"pet-s.a-graphql/graph/model"
)

func (s *MyService) InsertPet(pet model.NewPet) (*model.Pet, error) {
	values := map[string]interface{}{
		"nombre_replace": pet.Nombre,
		"raza_replace":   pet.Raza,
		"edad_replace":   pet.Edad,
		"peso_replace":   pet.Peso,
		"cedula_replace": pet.CedulaCliente,
	}

	_, err := repo.InsertPet(values)
	if err != nil {
		return &model.Pet{}, err
	}

	fmt.Println("###############################")
	log.Println("MASCOTA INSERTADA: NOMBRE: ", pet.Nombre)
	fmt.Println("###############################")

	return &model.Pet{
		Nombre:        pet.Nombre,
		Raza:          pet.Raza,
		Edad:          pet.Edad,
		Peso:          pet.Peso,
		CedulaCliente: pet.CedulaCliente,
	}, err
}

func (s *MyService) DeletePet(pet int) (int, error) {
	values := map[string]interface{}{
		"replace_id": pet,
	}

	_, err := repo.DeletePet(values)
	if err != nil {
		return 0, err
	}

	fmt.Println("###############################")
	log.Println("MASCOTA ELIMINADA: ID: ", pet)
	fmt.Println("###############################")

	return pet, nil
}

func (s *MyService) UpdatePet(pet model.UpdatePet) (*model.Pet, error) {
	values := map[string]interface{}{
		"id_replace":     pet.ID,
		"nombre_replace": pet.Nombre,
		"raza_replace":   pet.Raza,
		"edad_replace":   pet.Edad,
		"peso_replace":   pet.Peso,
	}

	_, err := repo.UpdatePet(values)
	if err != nil {
		return &model.Pet{}, err
	}

	fmt.Println("###############################")
	log.Println("MASCOTA ACTUALIZADA: NOMBRE: ", pet.Nombre)
	fmt.Println("###############################")

	return &model.Pet{
		Nombre:        pet.Nombre,
		Raza:          pet.Raza,
		Edad:          pet.Edad,
		Peso:          pet.Peso,
		CedulaCliente: pet.CedulaCliente,
	}, err
}

func (s *MyService) GetPets() ([]*model.Pet, error) {
	result, err := repo.GetPets()
	if err != nil {
		return []*model.Pet{}, err
	}

	var pets []*model.Pet

	json.Unmarshal([]byte(result), &pets)

	fmt.Println("###############################")
	log.Println("MASCOSTAS OBTENIDAS ")
	fmt.Println("###############################")

	return pets, err
}
