package repository

import (
	"fmt"
	"strings"
)

var (
	errInsertPet = "error insertando la mascota"
	errDeletePet = "error eliminando la mascota"
	errUpdatePet = "error actualizando la mascota"
	errGetPets   = "error obteniendo las mascotas"
)

func (r *MyRepository) InsertPet(infoPet map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/pets/insertPet.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoPet {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	_, err = DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errInsertPet, err)
	}

	return "OK", nil
}

func (r *MyRepository) DeletePet(infoPet map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/pets/deletePet.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoPet {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	_, err = DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errDeletePet, err)
	}

	return "OK", nil
}

func (r *MyRepository) UpdatePet(infoPet map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/pets/updatePet.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoPet {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	_, err = DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errUpdatePet, err)
	}

	return "OK", nil
}

func (r *MyRepository) GetPets() (string, error) {
	sqlFile, err := content.ReadFile("queries/pets/getPets.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	result, err := DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errGetPets, err)
	}

	return result, nil
}
