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

var EventModel *eventModel

type eventController struct {
	quick.QQuickItem
	_ func() `constructor:"init"`

	events []*someEvent
	_      func(*core.QDate) []*someEvent `slot:"eventsForDate,auto"`

	_ *core.QAbstractListModel `property:"eventModel"`
	_ *core.QDate              `property:"selectedDate"`
	_ func()                   `signal:"calendarUpdate"`
}

type eventModel struct {
	core.QAbstractListModel
	_ *core.QDate `property:"selectedDate"`
	_ string      `property:"filterExpression"`
	_ func()      `signal:"update"`
}

func (e *eventModel) update() {
	fmt.Println("event controller update")
	e.BeginResetModel()
	e.EndResetModel()
}

//naming it "event" won't work at the moment for some reason
type someEvent struct {
	core.QObject

	_ string          `property:"name"`
	_ string          `property:"description"`
	_ *core.QDateTime `property:"startDate"`
	_ *core.QDateTime `property:"endDate"`
}

func (e *eventController) eventsForDate(d *core.QDate) (o []*someEvent) {
	fmt.Println("requested date ", d.Year(), d.Month(), d.Day())
	for _, diff := range e.events { //it all rides on where events comes from as to who owns this function, which is fine
		startDate := diff.StartDate().Date()
		//e is a QDateTime, need to get the Date() object to get the Year/Month/Day
		if startDate.Year() == d.Year() && int(startDate.Month()) == d.Month() && startDate.Day() == d.Day() {
			fmt.Println("event handler adding ", startDate.Year(), startDate.Month(), startDate.Day())
			o = append(o, diff)
		}
	}
	return
}

func (e *eventController) init() {
	fmt.Println("event model initd")
	eModel := NewEventModel(nil) //configure the event model
	e.SetEventModel(eModel)

	eModel.ConnectRowCount(func(*core.QModelIndex) int {
		if e.SelectedDate() == nil {
			return 0
		}
		return len(e.eventsForDate(eModel.SelectedDate()))
	})

	eModel.ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if eModel.SelectedDate() == nil || role != int(core.Qt__DisplayRole) {
			return core.NewQVariant()
		}
		return e.eventsForDate(e.SelectedDate())[index.Row()].ToVariant()
	})

	eModel.ConnectSelectedDateChanged(func(date *core.QDate) {
		fmt.Println("hahahahah selected date changed", date.Year(), date.Month(), date.Day())
		fmt.Println("eModel selected date set to", eModel.SelectedDate().Year(), eModel.SelectedDate().Month(), eModel.SelectedDate().Day())
		eModel.BeginResetModel()
		eModel.EndResetModel()
	})

	eModel.ConnectUpdate(eModel.update)
	e.ConnectCalendarUpdate(eModel.Update)
	e.ConnectSelectedDate(eModel.SelectedDate)
	e.ConnectSelectedDateChanged(func(date *core.QDate) {
		fmt.Println("controller date changed to ", date.Year(), date.Month(), date.Day())
		e.SetSelectedDate(date)
		eModel.SetSelectedDate(date)
		fmt.Println("eController selected date set to", e.SelectedDate().Year(), e.SelectedDate().Month(), e.SelectedDate().Day())
		fmt.Println("eModel selected date set to", eModel.SelectedDate().Year(), eModel.SelectedDate().Month(), eModel.SelectedDate().Day())
		eModel.BeginResetModel()
		eModel.EndResetModel()
		e.CalendarUpdate()
	})

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
