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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QStackedWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:59
// index:0
// void QStackedWidget(class QWidget *)
func NewQStackedWidget(parent unsafe.Pointer) *QStackedWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidgetC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QStackedWidget{cthis}
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:60
// index:0
// virtual
// void ~QStackedWidget()
func DeleteQStackedWidget(*QStackedWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:62
// index:0
// int addWidget(class QWidget *)
func (this *QStackedWidget) AddWidget(w unsafe.Pointer) {
	// 0: (, QWidget * w), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget9addWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:63
// index:0
// int insertWidget(int, class QWidget *)
func (this *QStackedWidget) InsertWidget(index int, w unsafe.Pointer) {
	// 0: (, int index, QWidget * w), (&index, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget12insertWidgetEiP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &index, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:64
// index:0
// void removeWidget(class QWidget *)
func (this *QStackedWidget) RemoveWidget(w unsafe.Pointer) {
	// 0: (, QWidget * w), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget12removeWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:66
// index:0
// QWidget * currentWidget()
func (this *QStackedWidget) CurrentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget13currentWidgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:67
// index:0
// int currentIndex()
func (this *QStackedWidget) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:69
// index:0
// int indexOf(class QWidget *)
func (this *QStackedWidget) IndexOf(arg0 unsafe.Pointer) {
	// 0: (, QWidget * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget7indexOfEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:70
// index:0
// QWidget * widget(int)
func (this *QStackedWidget) Widget(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget6widgetEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:71
// index:0
// int count()
func (this *QStackedWidget) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedWidget5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:74
// index:0
// void setCurrentIndex(int)
func (this *QStackedWidget) SetCurrentIndex(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget15setCurrentIndexEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:75
// index:0
// void setCurrentWidget(class QWidget *)
func (this *QStackedWidget) SetCurrentWidget(w unsafe.Pointer) {
	// 0: (, QWidget * w), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:78
// index:0
// void currentChanged(int)
func (this *QStackedWidget) CurrentChanged(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget14currentChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:79
// index:0
// void widgetRemoved(int)
func (this *QStackedWidget) WidgetRemoved(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedWidget13widgetRemovedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

//  body block end
