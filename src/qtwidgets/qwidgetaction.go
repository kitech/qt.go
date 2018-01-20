//  header block begin
// /usr/include/qt/QtWidgets/qwidgetaction.h
// #include <qwidgetaction.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QWidgetAction struct {
	*QAction
}

func (this *QWidgetAction) GetCthis() unsafe.Pointer {
	return this.QAction.GetCthis()
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QWidgetAction) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QWidgetAction10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:59
// index:0
// void QWidgetAction(class QObject *)
func NewQWidgetAction(parent unsafe.Pointer) *QWidgetAction {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetActionC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQWidgetActionFromPointer(cthis)
	return gothis
}
func NewQWidgetActionFromPointer(cthis unsafe.Pointer) *QWidgetAction {
	bcthis0 := NewQActionFromPointer(cthis)
	return &QWidgetAction{bcthis0}
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:60
// index:0
// virtual
// void ~QWidgetAction()
func DeleteQWidgetAction(*QWidgetAction) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetActionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:62
// index:0
// void setDefaultWidget(class QWidget *)
func (this *QWidgetAction) SetDefaultWidget(w unsafe.Pointer) {
	// 0: (, w QWidget *), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetAction16setDefaultWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:63
// index:0
// QWidget * defaultWidget()
func (this *QWidgetAction) DefaultWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QWidgetAction13defaultWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:65
// index:0
// QWidget * requestWidget(class QWidget *)
func (this *QWidgetAction) RequestWidget(parent unsafe.Pointer) {
	// 0: (, parent QWidget *), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetAction13requestWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:66
// index:0
// void releaseWidget(class QWidget *)
func (this *QWidgetAction) ReleaseWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetAction13releaseWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:69
// index:0
// virtual
// bool event(class QEvent *)
func (this *QWidgetAction) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetAction5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:70
// index:0
// virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QWidgetAction) EventFilter(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, QObject * arg0, QEvent * arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetAction11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:71
// index:0
// virtual
// QWidget * createWidget(class QWidget *)
func (this *QWidgetAction) CreateWidget(parent unsafe.Pointer) {
	// 0: (, parent QWidget *), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetAction12createWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:72
// index:0
// virtual
// void deleteWidget(class QWidget *)
func (this *QWidgetAction) DeleteWidget(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetAction12deleteWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:73
// index:0
// QList<QWidget *> createdWidgets()
func (this *QWidgetAction) CreatedWidgets() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QWidgetAction14createdWidgetsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
