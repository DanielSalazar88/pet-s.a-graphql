package database

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"pet-s.a-graphql/settings"
)

type MySQL struct {
}

var DB = settings.ReadYaml()

var (
	errMysqlConn = errors.New("error conectando a mysql")
	errMysqlExec = errors.New("error ejecutando consulta a mysql")
	errMysqlScan = errors.New("error al escanear el resultado")
)

func (db *MySQL) mysqlConn() (*sql.DB, error) {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		DB.Mysql.Username,
		DB.Mysql.Password,
		DB.Mysql.Address,
		DB.Mysql.Port,
		DB.Mysql.Database,
	)

	mysql, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, fmt.Errorf("%s : %s", errMysqlConn, err)
	}

	return mysql, nil
}

func (db *MySQL) ExecMysqlQuery(query string) (string, error) {
	mysqlClient, err := db.mysqlConn()
	if err != nil {
		return "", err
	}

	err = mysqlClient.Ping()
	if err != nil {
		return "", fmt.Errorf("%s : %s", errMysqlConn, err)
	}

	rows, err := mysqlClient.Query(query)
	if err != nil {
		return "", fmt.Errorf("%s : %s", errMysqlExec, err)
	}

	var result string
	for rows.Next() {
		if err := rows.Scan(&result); err != nil {
			return "", fmt.Errorf("%s : %s", errMysqlScan, err)
		}
	}

	return result, nil
}
