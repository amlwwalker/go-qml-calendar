
import Events 1.0
import "." as This

EventController {
  This.Calendar {
      id: calendar
      width: (parent.width > parent.height ? parent.width * 0.6 - parent.spacing : parent.width)
      height: (parent.height > parent.width ? parent.height * 0.6 - parent.spacing : parent.height)
      frameVisible: true
      weekNumbersVisible: true
      selectedDate: new Date(2014, 0, 1)
      focus: true
  }
  visible: {
    console.log(calendar.selectedDate)
    console.log(eventsForDate(calendar.selectedDate).length)
    return true
  }
  selectedDate: calendar.selectedDate
}