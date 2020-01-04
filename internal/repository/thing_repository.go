//go:generate mockgen -source thing_repository.go -destination mocks/thing_repository.go
package repository

import (
	"github.com/gocraft/dbr/v2"
	"github.com/knappjf/quickquestion/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type ThingRepository interface {
	GetThing(int) (models.Thing, error)
	CreateThing(models.Thing) error
	DeleteThing(int) error
	UpdateThing(int, models.Thing) error
}

type thingRepository struct {
	sessionRunner dbr.SessionRunner
}

func NewThingRepository(sessionRunner dbr.SessionRunner) *thingRepository {
	return &thingRepository{
		sessionRunner: sessionRunner,
	}
}

func (tr *thingRepository) CreateThing(thing models.Thing) error {
	_, err := tr.sessionRunner.
		InsertInto("things").
		Pair("name", thing.Name).
		Pair("description", thing.Description).
		Pair("enabled", thing.Enabled).
		Exec()

	return err
}

func (tr *thingRepository) GetThing(id int) (models.Thing, error) {
	var thing models.Thing

	err := tr.sessionRunner.
		Select("name", "description", "enabled").
		From("things").
		Where("thing_id = ?", id).
		LoadOne(&thing)

	if err != nil {
		return models.Thing{}, err
	}

	return thing, nil
}

func (tr *thingRepository) DeleteThing(id int) error {
	_, err := tr.sessionRunner.DeleteFrom("things").Where("thing_id = ?", id).Exec()
	return err
}

func (tr *thingRepository) UpdateThing(id int, thing models.Thing) error {
	_, err := tr.sessionRunner.
		Update("things").
		Set("name", thing.Name).
		Set("description", thing.Description).
		Set("enabled", thing.Enabled).
		Where("thing_id = ?", id).
		Exec()

	return err
}
