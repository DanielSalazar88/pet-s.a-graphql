package service

import (
	"encoding/json"
	"fmt"
	"log"

	"pet-s.a-graphql/graph/model"
)

func (s *MyService) InsertMedicine(medicine model.NewMedicine) (*model.Medicine, error) {
	values := map[string]interface{}{
		"nombre_replace":      medicine.Nombre,
		"descripcion_replace": medicine.Descripcion,
		"dosis_replace":       medicine.Dosis,
	}

	_, err := repo.InsertMedicine(values)
	if err != nil {
		return &model.Medicine{}, err
	}

	fmt.Println("###############################")
	log.Println("MEDICAMENTO INSERTADO: NOMBRE: ", medicine.Nombre)
	fmt.Println("###############################")

	return &model.Medicine{
		Nombre:      medicine.Nombre,
		Descripcion: medicine.Descripcion,
		Dosis:       medicine.Dosis,
	}, err

}

func (s *MyService) DeleteMedicine(medicine int) (int, error) {
	values := map[string]interface{}{
		"replace_id": medicine,
	}

	_, err := repo.DeleteMedicine(values)
	if err != nil {
		return 0, err
	}

	fmt.Println("###############################")
	log.Println("MEDICAMENTO ELIMINADO: ID: ", medicine)
	fmt.Println("###############################")

	return medicine, err
}

func (s *MyService) UpdateMedicine(medicine model.UpdateMedicine) (*model.Medicine, error) {
	values := map[string]interface{}{
		"id_replace":          medicine.ID,
		"nombre_replace":      medicine.Nombre,
		"descripcion_replace": medicine.Descripcion,
		"dosis_replace":       medicine.Dosis,
	}

	_, err := repo.UpdateMedicine(values)
	if err != nil {
		return &model.Medicine{}, err
	}

	fmt.Println("###############################")
	log.Println("MEDICAMENTO ACTUALIZARO: NOMBRE: ", medicine.Nombre)
	fmt.Println("###############################")

	return &model.Medicine{
		Nombre:      medicine.Nombre,
		Descripcion: medicine.Descripcion,
		Dosis:       medicine.Dosis,
	}, err
}

func (s *MyService) GetMedicines() ([]*model.Medicine, error) {
	result, err := repo.GetMedicines()
	if err != nil {
		return []*model.Medicine{}, err
	}

	var medicines []*model.Medicine

	json.Unmarshal([]byte(result), &medicines)

	fmt.Println("###############################")
	log.Println("MEDICAMENTOS OBTENIDOS ")
	fmt.Println("###############################")

	return medicines, err
}
