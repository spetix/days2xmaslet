package event

import (
	"fmt"
	"time"

	"github.com/spetix/days2xmasleft/internal/data/render"
	"github.com/spetix/days2xmasleft/internal/dateutil"
)

type timeReferenceFunc func() time.Time

type Event struct {
	eventName     string
	eventDate     time.Time
	renderOptions *render.RenderOptions
	timeHandler   timeReferenceFunc
}

func New(eventName string, eventDateString string, renderOptions *render.RenderOptions, timeHandler timeReferenceFunc) *Event {
	eventDate := dateutil.StringToDateNextOccurrence(eventDateString)
	return &Event{
		eventName:     eventName,
		eventDate:     eventDate,
		renderOptions: renderOptions,
		timeHandler:   timeHandler,
	}
}

func (r *Event) timeToEvent() time.Duration {
	currentTime := r.timeHandler()
	adjustedDate := currentTime.Truncate(r.renderOptions.Unit)
	if adjustedDate.After(r.eventDate) {
		adjustedDate = adjustedDate.AddDate(1, 0, 0)
	}
	return r.eventDate.Sub(adjustedDate)
}

func (r *Event) Short() string {
	return dateutil.Format(r.timeToEvent(), r.renderOptions.Unit)
}

func (r *Event) Long() string {
	return fmt.Sprint(r.Short(), " to ", r.eventName)
}

func (r *Event) Label() string {
	return r.renderOptions.Label
}

func (r *Event) BackgroundColor() string {
	return r.renderOptions.BackgroundColor
}

func (r *Event) ForegroundColor() string {
	return r.renderOptions.ForegroundColor
}
