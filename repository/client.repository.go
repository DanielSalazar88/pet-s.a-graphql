package repository

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errInsertClient = errors.New("error insertando el cliente")
	errDeleteClient = errors.New("error eliminando el cliente")
	errUpdateClient = errors.New("error eliminando el cliente")
)

func (r *MyRepository) InsertClient(infoReplace map[string]string) (string, error) {
	sqlFile, err := content.ReadFile("queries/clients/insertClient.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoReplace {
		query = strings.ReplaceAll(query, marker, valor)
	}

	_, err = DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errInsertClient, err)
	}

	return "Ok", nil
}

func (r *MyRepository) DeleteClient(infoClientDelete map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/clients/deleteClient.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoClientDelete {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	sqlStatements := strings.Split(string(query), ";")

	for _, query := range sqlStatements {
		query = strings.TrimSpace(query)

		if query == "" {
			continue
		}

		_, err = DB.ExecMysqlQuery(query)
		if err != nil {
			return "", fmt.Errorf("%s : %s", errDeleteClient, err)
		}
	}

	return "Ok", nil
}

func (r *MyRepository) UpdateClient(infoReplace map[string]string) (string, error) {
	sqlFile, err := content.ReadFile("queries/clients/updateClient.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoReplace {
		query = strings.ReplaceAll(query, marker, valor)
	}

	_, err = DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errUpdateClient, err)
	}

	return "Ok", nil
}

func (r *MyRepository) GetClients() (string, error) {
	sqlFile, err := content.ReadFile("queries/clients/getClients.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	result, err := DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errInsertClient, err)
	}

	return result, nil
}
