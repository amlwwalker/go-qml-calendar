package main

import (
	"os"
	"sync"

	_ "github.com/amlwwalker/go-qml-calendar/events"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

type Controller struct {
	core.QObject
	qApp *widgets.QApplication
}

var once sync.Once
var instance *Controller

func Instance() *Controller {
	once.Do(func() {
		instance = NewController(nil)
	})
	return instance
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	qApp := widgets.NewQApplication(len(os.Args), os.Args)
	Instance()
	Instance().qApp = qApp

	app := qml.NewQQmlApplicationEngine(nil)
	app.AddImportPath("./events/qml")
	app.Load(core.NewQUrl3("./view/Main.qml", 0))
	widgets.QApplication_Exec()
}
