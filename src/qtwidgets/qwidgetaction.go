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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QWidgetAction) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QWidgetAction10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:59
// index:0
// void QWidgetAction(class QObject *)
func NewQWidgetAction(parent unsafe.Pointer) *QWidgetAction {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetActionC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QWidgetAction{cthis}
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
	// 0: (, QWidget * w), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetAction16setDefaultWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:63
// index:0
// QWidget * defaultWidget()
func (this *QWidgetAction) DefaultWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QWidgetAction13defaultWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:65
// index:0
// QWidget * requestWidget(class QWidget *)
func (this *QWidgetAction) RequestWidget(parent unsafe.Pointer) {
	// 0: (, QWidget * parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetAction13requestWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidgetaction.h:66
// index:0
// void releaseWidget(class QWidget *)
func (this *QWidgetAction) ReleaseWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QWidgetAction13releaseWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

//  body block end
