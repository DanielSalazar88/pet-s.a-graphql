package service

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("###############################")
	log.Println("CLIENTE INSERTADO: CEDULA: ", client.Cedula)
	fmt.Println("###############################")

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

	fmt.Println("###############################")
	log.Println("CLIENTE ELIMINADO: CEDULA: ", cedula)
	fmt.Println("###############################")

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

	fmt.Println("###############################")
	log.Println("CLIENTE ACTUALIZADO: CEDULA: ", client.Cedula)
	fmt.Println("###############################")

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

	fmt.Println("###############################")
	log.Println("CLIENTES OBTENIDOS ")
	fmt.Println("###############################")

	return clients
}
