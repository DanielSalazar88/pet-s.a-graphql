package repository

import (
	"fmt"
	"strings"
)

var (
	errGeneralClientReport = "error obteniendo el reporte general del cliente"
	errRecipesPetReport    = "error obteniendo el reporte de las recetas de la mascota"
)

func (r *MyRepository) GeneralClientReport(infoReplace map[string]string) (string, error) {
	sqlFile, err := content.ReadFile("queries/reports/clientReport.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoReplace {
		query = strings.ReplaceAll(query, marker, valor)
	}

	result, err := DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errGeneralClientReport, err)
	}

	return result, nil
}

func (r *MyRepository) RecipesPetReport(infoReplace map[string]interface{}) (string, error) {
	sqlFile, err := content.ReadFile("queries/reports/petReport.sql")
	if err != nil {
		return "", fmt.Errorf("%s : %s", errReadSqlFIle, err)
	}

	query := string(sqlFile)

	for marker, valor := range infoReplace {
		query = strings.ReplaceAll(query, marker, fmt.Sprintf("%v", valor))
	}

	result, err := DB.ExecMysqlQuery(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errRecipesPetReport, err)
	}

	return result, nil
}
