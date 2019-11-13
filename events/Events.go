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

	events []*someEvent
	_      func(*core.QDate) []*someEvent `slot:"eventsForDate,auto"`

	_ *core.QAbstractListModel `property:"listModel"`
	_ *core.QDate              `property:"selectedDate"`
	_ func()                   `signal:"update"`
}

//naming it "event" won't work at the moment for some reason
type someEvent struct {
	core.QObject

	_ string          `property:"name"`
	_ string          `property:"description"`
	_ *core.QDateTime `property:"startDate"`
	_ *core.QDateTime `property:"endDate"`
}

func (e *eventController) init() {
	fmt.Println("event model initd")
	e.SetListModel(core.NewQAbstractListModel(nil))

	e.ListModel().ConnectRowCount(func(*core.QModelIndex) int {
		if e.SelectedDate() == nil {
			return 0
		}
		return len(e.eventsForDate(e.SelectedDate()))
	})

	e.ListModel().ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if e.SelectedDate() == nil || role != int(core.Qt__DisplayRole) {
			return core.NewQVariant()
		}
		return e.eventsForDate(e.SelectedDate())[index.Row()].ToVariant()
	})

	e.ConnectSelectedDateChanged(func(*core.QDate) {
		e.ListModel().BeginResetModel()
		e.ListModel().EndResetModel()
	})

	e.ConnectUpdate(e.update)

	for i := 0; i < 3; i++ {
		ev := NewSomeEvent(nil)
		ev.SetName(fmt.Sprintf("event (%v) on the %v.%v.%v", i+1, 2014, 1, i+1))
		ev.SetDescription(fmt.Sprintf("started adding properties for day (%+v)", i+1))
		st := core.NewQDateTime()
		st.SetDate(core.NewQDate3(2014, 2, 3))
		st.SetTime(core.NewQTime3(time.Now().Hour(), time.Now().Minute(), 0, 0))
		ev.SetStartDate(st)

		et := core.NewQDateTime()
		et.SetDate(core.NewQDate3(2014, 2, 3))
		et.SetTime(core.NewQTime3(time.Now().Hour(), time.Now().Minute(), 0, 0))
		ev.SetEndDate(et)

		e.events = append(e.events, ev)
	}
	// for i := 0; i < 1; i++ {
	// 	ev := NewSomeEvent(nil)
	// 	ev.SetName(fmt.Sprintf("event (%v) on the %v.%v.%v", i+1, 2014, 1, i+1))
	// 	ev.SetDescription(fmt.Sprintf("started adding properties for day (%+v)", i+1))
	// 	st := core.NewQDateTime()
	// 	st.SetDate(core.NewQDate3(2014, 2, 5))
	// 	st.SetTime(core.NewQTime3(time.Now().Hour(), time.Now().Minute(), 0, 0))
	// 	ev.SetStartDate(st)

	// 	et := core.NewQDateTime()
	// 	et.SetDate(core.NewQDate3(2014, 2, 5))
	// 	et.SetTime(core.NewQTime3(time.Now().Hour(), time.Now().Minute(), 0, 0))
	// 	ev.SetEndDate(et)

	// 	e.events = append(e.events, ev)
	// }

	go func() {
		time.Sleep(3 * time.Second)
		goEve := NewSomeEvent(nil)
		goEve.SetName(fmt.Sprintf("thread event (%v) on the %v.%v.%v", 10+1, 2014, 1, 10+1))
		goEve.SetDescription(fmt.Sprintf("started adding properties for day (%+v)", 10+1))
		st := core.NewQDateTime()
		st.SetDate(core.NewQDate3(2014, 2, 5))
		st.SetTime(core.NewQTime3(time.Now().Hour(), time.Now().Minute(), 0, 0))
		goEve.SetStartDate(st)

		et := core.NewQDateTime()
		et.SetDate(core.NewQDate3(2014, 2, 5))
		et.SetTime(core.NewQTime3(time.Now().Hour(), time.Now().Minute(), 0, 0))
		goEve.SetEndDate(et)
		goEve.MoveToThread(e.Thread())
		e.events = append(e.events, goEve)
		e.Update()
		fmt.Println("updating")
	}()
}

func (e *eventController) update() {
	e.ListModel().BeginResetModel()
	e.ListModel().EndResetModel()
}

func (e *eventController) eventsForDate(d *core.QDate) (o []*someEvent) {

	for _, diff := range e.events {
		startDate := diff.StartDate().Date()
		//e is a QDateTime, need to get the Date() object to get the Year/Month/Day
		if startDate.Year() == d.Year() && int(startDate.Month()) == d.Month() && startDate.Day() == d.Day() {
			fmt.Println("event handler adding ", startDate.Year())
			o = append(o, diff)
		}
	}
	return
}
