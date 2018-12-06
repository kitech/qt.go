package main

import (
	"log"
	"os"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtmeta"
	"github.com/kitech/qt.go/qtrt"
	"github.com/kitech/qt.go/qtwidgets"
)

type MainWindow struct {
	*qtwidgets.QMainWindow `qt:"inherit"`
}

func NewMainWindow(parent qtwidgets.QWidget_ITF, fo qtcore.Qt__WindowType) *MainWindow {
	mainWindow := &MainWindow{}
	qtmeta.Derive(mainWindow)
	mainWindow.init()
	return mainWindow
}

func (m *MainWindow) init() {
	m.SetWindowTitle("housekeeper")
	m.SetFixedSize(qtcore.NewQSize_1(300, 200))

	frame := NewMainFrame(m, 0)
	m.SetCentralWidget(frame)
}

type MainFrame struct {
	*qtwidgets.QFrame `qt:"inherit"`

	_          qtmeta.Q_SIGNALS_BEGIN
	Clicked123 func() `qt:"signal"`
}

func NewMainFrame(parent qtwidgets.QWidget_ITF, fo qtcore.Qt__WindowType) *MainFrame {
	mainFrame := &MainFrame{}
	qtmeta.Derive(mainFrame)
	mainFrame.init()

	return mainFrame
}

func (m *MainFrame) init() {
	hboxLayout := qtwidgets.NewQHBoxLayout_1(m)

	button1 := qtwidgets.NewQPushButton_1("a", m)

	hboxLayout.AddStretch(1)
	hboxLayout.AddWidget(button1, 0, qtcore.Qt__AlignCenter)
	hboxLayout.AddWidget(qtwidgets.NewQPushButton_1("b", m), 0, qtcore.Qt__AlignCenter)
	hboxLayout.AddWidget(qtwidgets.NewQPushButton_1("c", m), 0, qtcore.Qt__AlignCenter)
	hboxLayout.AddStretch(1)

	m.SetLayout(hboxLayout)

	qtrt.Connect(button1, "clicked()", func() {
		log.Println("v")
		m.Clicked123() // call goside custom defined signal
	})

	// this is a custom signal defined in MainFrame struct
	qtrt.Connect(m, "Clicked123()", func() {
		log.Println("a")
	})
}

func main() {
	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)
	app.SetApplicationName("housekeeper")

	NewMainWindow(nil, 0).Show()

	os.Exit(app.Exec())
}
