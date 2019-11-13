import QtQuick.Controls 2.4
import QtQuick.Dialogs 1.2
import QtQuick.Layouts 1.3
import QtQuick 2.2
import QtQuick.Controls 1.4 as Old
import QtQuick.Controls.Styles 1.4
import QtQuick.Controls.Private 1.0

Old.Calendar {
  id: calendar
  style: CalendarStyle {
    gridVisible: true
    dayOfWeekDelegate: Rectangle {
        color: gridVisible ? "#fcfcfc" : "transparent"
        implicitHeight: Math.round(TextSingleton.implicitHeight * 2.25)
        Label {
            text: control.locale.dayName(styleData.dayOfWeek, control.dayOfWeekFormat)
            anchors.centerIn: parent
        }
    }
    navigationBar: Rectangle {
      height: Math.round(TextSingleton.implicitHeight * 2.73)
      color: "#f9f9f9"

      Rectangle {
          color: Qt.rgba(1,1,1,0.6)
          height: 1
          width: parent.width
      }

      Rectangle {
          anchors.bottom: parent.bottom
          height: 1
          width: parent.width
          color: "#ddd"
      }

      HoverButton {
          id: previousMonth
          width: parent.height
          height: width
          anchors.verticalCenter: parent.verticalCenter
          anchors.left: parent.left
          source: "./assets/leftanglearrow.png"
          onClicked: control.showPreviousMonth()
      }
      Label {
          id: dateText
          text: styleData.title
          elide: Text.ElideRight
          horizontalAlignment: Text.AlignHCenter
          font.pixelSize: TextSingleton.implicitHeight * 1.25
          anchors.verticalCenter: parent.verticalCenter
          anchors.left: previousMonth.right
          anchors.leftMargin: 2
          anchors.right: nextMonth.left
          anchors.rightMargin: 2
      }
      HoverButton {
          id: nextMonth
          width: parent.height
          height: width
          anchors.verticalCenter: parent.verticalCenter
          anchors.right: parent.right
          source: "./assets/rightanglearrow.png"
          onClicked: control.showNextMonth()
      }
    }
    dayDelegate: Item {
      readonly property color sameMonthDateTextColor: "#444"
      readonly property color selectedDateColor: Qt.platform.os === "osx" ? "#3778d0" : systemPalette.highlight
      readonly property color selectedDateTextColor: "white"
      readonly property color differentMonthDateTextColor: "#bbb"
      readonly property color invalidDatecolor: "#dddddd"

      Rectangle {
          anchors.fill: parent
          border.color: "transparent"
          color: {
              console.log("styleData.selected ", styleData.selected)
              return styleData.date !== undefined && styleData.selected ? selectedDateColor : "transparent"
            }
          anchors.margins: styleData.selected ? -1 : 0
      }

      Image {
          visible: calendar.parent.eventsForDate(styleData.date).length > 0
          anchors.top: parent.top
          anchors.left: parent.left
          anchors.margins: -1
          width: 12
          height: width
          source: "./assets/eventindicator.png"
      }

      Label {
          id: dayDelegateText
          text: styleData.date.getDate()
          anchors.centerIn: parent
          color: {
              var color = invalidDatecolor;
              if (styleData.valid) {
                  // Date is within the valid range.
                  color = styleData.visibleMonth ? sameMonthDateTextColor : differentMonthDateTextColor;
                  if (styleData.selected) {
                      color = selectedDateTextColor;
                  }
              }
              color;
          }
      }
    }
  }
}
