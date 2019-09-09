package events

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
)

func init() {
	eventController_QmlRegisterType2("Events", 1, 0, "EventController")
	fmt.Println("eventController initd")
}

type eventController struct {
	quick.QQuickItem
	_ func() `constructor:"init"`

	_ *core.QAbstractListModel `property:"eventModel"`
}

type event struct {
	_ string      `property:"name"`
	_ *core.QDate `property:"date"`
}

func (e *eventController) init() {
	fmt.Println("event model initd")
	eventModel := core.NewQAbstractListModel(nil)
	// event := &event{"a date", core.NewQDate3(2014, 2, 3)}

	e.SetEventModel(eventModel)
}
