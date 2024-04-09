package service

import (
	"encoding/json"

	"pet-s.a-graphql/graph/model"
)

func (s *MyService) InsertRecipe(recipe model.NewRecipe) (*model.Recipe, error) {
	values := map[string]interface{}{
		"id_mascota_replace":     recipe.IDMascota,
		"id_medicamento_replace": recipe.IDMedicamento,
	}

	_, err := repo.InsertRecipe(values)
	if err != nil {
		return &model.Recipe{}, err
	}

	return &model.Recipe{
		IDMedicamento: recipe.IDMedicamento,
		IDMascota:     recipe.IDMascota,
	}, nil
}

func (s *MyService) DeleteRecipe(recipe int) (int, error) {
	values := map[string]interface{}{
		"id_replace": recipe,
	}

	_, err := repo.DeleteRecipe(values)
	if err != nil {
		return 0, err
	}

	return recipe, nil
}

func (s *MyService) GetRecipes() ([]*model.Recipe, error) {
	result, err := repo.GetRecipes()
	if err != nil {
		return []*model.Recipe{}, err
	}

	var recipes []*model.Recipe

	json.Unmarshal([]byte(result), &recipes)

	return recipes, err
}
