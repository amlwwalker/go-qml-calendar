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
  }
}
