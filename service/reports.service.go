package service

import (
	"encoding/json"
	"fmt"
	"log"

	"pet-s.a-graphql/graph/model"
)

func (s *MyService) GeneralClientReport(cedula string) ([]*model.GeneralClientReport, error) {
	values := map[string]string{
		"cedula_replace": cedula,
	}

	result, err := repo.GeneralClientReport(values)
	if err != nil {
		return []*model.GeneralClientReport{}, err
	}

	var report []*model.GeneralClientReport

	json.Unmarshal([]byte(result), &report)

	fmt.Println("###############################")
	log.Println("REPORTE GENERAL DEL CLIENTE")
	fmt.Println("###############################")

	return report, err
}

func (s *MyService) RecipesPetReport(pet int) ([]*model.PetReport, error) {
	values := map[string]interface{}{
		"id_mascota_replace": pet,
	}

	result, err := repo.RecipesPetReport(values)
	if err != nil {
		return []*model.PetReport{}, err
	}

	var report []*model.PetReport

	json.Unmarshal([]byte(result), &report)

	fmt.Println("###############################")
	log.Println("REPORTE GENERAL DE RECETAS POR MASCOTA")
	fmt.Println("###############################")

	return report, nil
}
