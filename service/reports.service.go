package service

func (s *MyService) GeneralClientReport(cedula string) (string, error) {
	values := map[string]string{
		"cedula_replace": cedula,
	}

	result, err := repo.GeneralClientReport(values)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (s *MyService) RecipesPetReport(pet int) (string, error) {
	values := map[string]interface{}{
		"id_mascota_replace": pet,
	}

	result, err := repo.RecipesPetReport(values)
	if err != nil {
		return "", err
	}

	return result, nil
}
