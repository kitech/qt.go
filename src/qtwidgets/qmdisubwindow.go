//  header block begin
// /usr/include/qt/QtWidgets/qmdisubwindow.h
// #include <qmdisubwindow.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 35
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QMdiSubWindow struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QMdiSubWindow) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:70
// index:0
// virtual
// void ~QMdiSubWindow()
func DeleteQMdiSubWindow(*QMdiSubWindow) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindowD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:72
// index:0
// virtual
// QSize sizeHint()
func (this *QMdiSubWindow) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:73
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QMdiSubWindow) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:75
// index:0
// void setWidget(class QWidget *)
func (this *QMdiSubWindow) SetWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:76
// index:0
// QWidget * widget()
func (this *QMdiSubWindow) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow6widgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:78
// index:0
// QWidget * maximizedButtonsWidget()
func (this *QMdiSubWindow) MaximizedButtonsWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow22maximizedButtonsWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:79
// index:0
// QWidget * maximizedSystemMenuIconWidget()
func (this *QMdiSubWindow) MaximizedSystemMenuIconWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:81
// index:0
// bool isShaded()
func (this *QMdiSubWindow) IsShaded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow8isShadedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:83
// index:0
// void setOption(enum QMdiSubWindow::SubWindowOption, _Bool)
func (this *QMdiSubWindow) SetOption(option int, on bool) {
	// 0: (, QMdiSubWindow::SubWindowOption option, bool on), (&option, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow9setOptionENS_15SubWindowOptionEb", ffiqt.FFI_TYPE_VOID, this.cthis, &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:84
// index:0
// bool testOption(enum QMdiSubWindow::SubWindowOption)
func (this *QMdiSubWindow) TestOption(arg0 int) {
	// 0: (, QMdiSubWindow::SubWindowOption arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow10testOptionENS_15SubWindowOptionE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:86
// index:0
// void setKeyboardSingleStep(int)
func (this *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	// 0: (, int step), (&step)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow21setKeyboardSingleStepEi", ffiqt.FFI_TYPE_VOID, this.cthis, &step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:87
// index:0
// int keyboardSingleStep()
func (this *QMdiSubWindow) KeyboardSingleStep() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow18keyboardSingleStepEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:89
// index:0
// void setKeyboardPageStep(int)
func (this *QMdiSubWindow) SetKeyboardPageStep(step int) {
	// 0: (, int step), (&step)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow19setKeyboardPageStepEi", ffiqt.FFI_TYPE_VOID, this.cthis, &step)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:90
// index:0
// int keyboardPageStep()
func (this *QMdiSubWindow) KeyboardPageStep() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow16keyboardPageStepEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:93
// index:0
// void setSystemMenu(class QMenu *)
func (this *QMdiSubWindow) SetSystemMenu(systemMenu unsafe.Pointer) {
	// 0: (, QMenu * systemMenu), (systemMenu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow13setSystemMenuEP5QMenu", ffiqt.FFI_TYPE_VOID, this.cthis, systemMenu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:94
// index:0
// QMenu * systemMenu()
func (this *QMdiSubWindow) SystemMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow10systemMenuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:97
// index:0
// QMdiArea * mdiArea()
func (this *QMdiSubWindow) MdiArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMdiSubWindow7mdiAreaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:101
// index:0
// void aboutToActivate()
func (this *QMdiSubWindow) AboutToActivate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow15aboutToActivateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:105
// index:0
// void showSystemMenu()
func (this *QMdiSubWindow) ShowSystemMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow14showSystemMenuEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qmdisubwindow.h:107
// index:0
// void showShaded()
func (this *QMdiSubWindow) ShowShaded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMdiSubWindow10showShadedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
