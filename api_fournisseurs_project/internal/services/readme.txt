Ce dossier contient la logique métier de votre application.

package services

import (
	"mon-projet/internal/models"
	"mon-projet/internal/repositories"
)

type ExampleService struct {
	exampleRepository *repositories.ExampleRepository
}

func NewExampleService() *ExampleService {
	return &ExampleService{
		exampleRepository: repositories.NewExampleRepository(),
	}
}

func (s *ExampleService) GetExample() (*models.ExampleModel, error) {
	// Logique pour obtenir un exemple de modèle
	return s.exampleRepository.GetExample()
}