//  header block begin
// /usr/include/qt/QtWidgets/qstackedwidget.h
// #include <qstackedwidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QStackedWidget struct {
	*QFrame
}

func (this *QStackedWidget) GetCthis() unsafe.Pointer {
	return this.QFrame.GetCthis()
}
func NewQStackedWidgetFromPointer(cthis unsafe.Pointer) *QStackedWidget {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QStackedWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QStackedWidget) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:59
// index:0
// Public
// void QStackedWidget(class QWidget *)
func NewQStackedWidget(parent unsafe.Pointer) *QStackedWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQStackedWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:60
// index:0
// Public virtual
// void ~QStackedWidget()
func DeleteQStackedWidget(*QStackedWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:62
// index:0
// Public
// int addWidget(class QWidget *)
func (this *QStackedWidget) AddWidget(w unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget9addWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:63
// index:0
// Public
// int insertWidget(int, class QWidget *)
func (this *QStackedWidget) InsertWidget(index int, w unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget12insertWidgetEiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, w)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:64
// index:0
// Public
// void removeWidget(class QWidget *)
func (this *QStackedWidget) RemoveWidget(w unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget12removeWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:66
// index:0
// Public
// QWidget * currentWidget()
func (this *QStackedWidget) CurrentWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget13currentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:67
// index:0
// Public
// int currentIndex()
func (this *QStackedWidget) CurrentIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:69
// index:0
// Public
// int indexOf(class QWidget *)
func (this *QStackedWidget) IndexOf(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget7indexOfEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:70
// index:0
// Public
// QWidget * widget(int)
func (this *QStackedWidget) Widget(arg0 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget6widgetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:71
// index:0
// Public
// int count()
func (this *QStackedWidget) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:74
// index:0
// Public
// void setCurrentIndex(int)
func (this *QStackedWidget) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:75
// index:0
// Public
// void setCurrentWidget(class QWidget *)
func (this *QStackedWidget) SetCurrentWidget(w unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:78
// index:0
// Public
// void currentChanged(int)
func (this *QStackedWidget) CurrentChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget14currentChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:79
// index:0
// Public
// void widgetRemoved(int)
func (this *QStackedWidget) WidgetRemoved(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget13widgetRemovedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:82
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QStackedWidget) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
