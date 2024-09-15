package application

import (
	"zoo/config"
	"zoo/libraries/logger"
	
	animalV1HtppDelivery "zoo/application/animal/delivery/v1http"
	animalHandler "zoo/application/animal/delivery/v1http/handler"
	animalRepository "zoo/application/animal/repository"
	animalUsecase "zoo/application/animal/usecase"
)

type (
	Dependency struct {
		Repositories Repositories
		Usecases     Usecases
		Handlers     Handlers
		Deliveries   Deliveries
		
		logger   logger.ILogger
		config   *config.Configurations
		postgres *config.PostgresGorm
	}
	Repositories struct {
		Animal animalRepository.IAnimalRepository
	}
	Usecases struct {
		Animal animalUsecase.IAnimalUsecase
	}
	Handlers struct {
		Animal animalHandler.IAnimalHandler
	}
	Deliveries struct {
		Animal *animalV1HtppDelivery.AnimalV1HttpDelivery
	}
)

func NewDependency(
	config *config.Configurations,
	logger logger.ILogger,
	postgres *config.PostgresGorm,
) *Dependency {
	dep := &Dependency{
		logger:   logger,
		config:   config,
		postgres: postgres,
	}
	
	dep.wireRepositories()
	dep.wireUsecases()
	dep.wireHandlers()
	dep.wireDeliveries()
	return dep
}

func (dep *Dependency) wireRepositories() {
	dep.Repositories.Animal = animalRepository.NewAnimalRepository(dep.postgres)
}

func (dep *Dependency) wireUsecases() {
	dep.Usecases.Animal = animalUsecase.NewAnimalUsecase(dep.logger, dep.Repositories.Animal)
}

func (dep *Dependency) wireHandlers() {
	dep.Handlers.Animal = animalHandler.NewAnimalHandler(dep.logger, dep.Usecases.Animal)
}

func (dep *Dependency) wireDeliveries() {
	dep.Deliveries.Animal = animalV1HtppDelivery.NewAnimalV1HttpDelivery(dep.Handlers.Animal)
}
