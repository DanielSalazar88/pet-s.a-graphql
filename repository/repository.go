package repository

import (
	"embed"
	"errors"

	"pet-s.a-graphql/database"
)

var (
	//go:embed queries/*
	content embed.FS

	errReadSqlFIle = errors.New("error leyendo el archivo sql")

	DB = database.MySQL{}
)

type MyRepository struct {
}
