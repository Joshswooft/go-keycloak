package repositories

import (
	"simplegoapi/src/config"
	"simplegoapi/src/domains"
	"simplegoapi/src/errors"
)

func Save(event *domains.Event) (*domains.Event, *errors.HttpError) {

	e := config.Database.Create(&event)

	if e.Error != nil {
		return nil, errors.DataAccessLayerError(e.Error.Error())
	}

	return event, nil
}

func FindOneById(id int) *domains.Event {
	var event domains.Event

	config.Database.First(&event, id)

	return &event
}

func FindAll() []domains.Event {
	var events []domains.Event
	config.Database.Find(&events)

	return events
}
