package service

import (
	"encoding/json"
	"log"

	"pet-s.a-graphql/graph/model"
)

func (s *MyService) InsertClient(client model.NewClient) (*model.Client, error) {
	values := map[string]string{
		"cedula_replace":    client.Cedula,
		"nombres_replace":   client.Nombres,
		"apellidos_replace": client.Apellidos,
		"direccion_replace": client.Direccion,
		"telefono_replace":  client.Telefono,
		"correo_replace":    client.Correo,
	}

	_, err := repo.InsertClient(values)
	if err != nil {
		log.Println(err)
		return &model.Client{}, err
	}

	return &model.Client{
		Cedula:    client.Cedula,
		Nombres:   client.Nombres,
		Apellidos: client.Apellidos,
		Direccion: client.Direccion,
		Telefono:  client.Telefono,
		Correo:    client.Correo,
	}, err
}

func (s *MyService) DeleteClient(cedula string) (string, error) {
	values := map[string]interface{}{
		"cedula_replace": cedula,
	}

	_, err := repo.DeleteClient(values)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return cedula, err
}

func (s *MyService) UpdateClient(client model.NewClient) (*model.Client, error) {
	values := map[string]string{
		"cedula_replace":    client.Cedula,
		"nombres_replace":   client.Nombres,
		"apellidos_replace": client.Apellidos,
		"direccion_replace": client.Direccion,
		"telefono_replace":  client.Telefono,
		"correo_replace":    client.Correo,
	}

	_, err := repo.UpdateClient(values)
	if err != nil {
		log.Println(err)
		return &model.Client{}, err
	}

	return &model.Client{
		Cedula:    client.Cedula,
		Nombres:   client.Nombres,
		Apellidos: client.Apellidos,
		Direccion: client.Direccion,
		Telefono:  client.Telefono,
		Correo:    client.Correo,
	}, err
}

func (s *MyService) GetClients() []*model.Client {
	result, err := repo.GetClients()
	if err != nil {
		log.Println(err)
	}

	var clients []*model.Client

	json.Unmarshal([]byte(result), &clients)

	return clients
}
