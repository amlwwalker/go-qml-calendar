package events

import (
	"fmt"
	"time"

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

	_ func(*core.QDate) []*someEvent `slot:"eventsForDate,auto"`
}

//naming it "event" won't work at the moment for some reason
type someEvent struct {
	core.QObject

	_ string          `property:"name"`
	_ *core.QDateTime `property:"startDate"`
	_ *core.QDateTime `property:"endDate"`
}

func (e *eventController) init() {
	fmt.Println("event model initd")
}

func (m *eventController) eventsForDate(d *core.QDate) (o []*someEvent) {

	for i := 0; i < d.Day(); i++ {
		e := NewSomeEvent(nil)
		e.SetName(fmt.Sprintf("event (%v) on the %v.%v.%v", i+1, d.Year(), d.Month(), d.Day()))

		st := core.NewQDateTime()
		st.SetDate(core.NewQDate3(2014, 2, 3))
		st.SetTime(core.NewQTime3(time.Now().Hour(), time.Now().Minute(), 0, 0))
		e.SetStartDate(st)

		et := core.NewQDateTime()
		et.SetDate(core.NewQDate3(2014, 2, 3))
		et.SetTime(core.NewQTime3(time.Now().Hour(), time.Now().Minute(), 0, 0))
		e.SetEndDate(et)

		o = append(o, e)
	}

	return
}
