package qtwidgets

// /usr/include/qt/QtWidgets/qstackedwidget.h
// #include <qstackedwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QStackedWidget) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QStackedWidget struct {
	*QFrame
}

func (this *QStackedWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFrame.GetCthis()
	}
}
func (this *QStackedWidget) SetCthis(cthis unsafe.Pointer) {
	this.QFrame = NewQFrameFromPointer(cthis)
}
func NewQStackedWidgetFromPointer(cthis unsafe.Pointer) *QStackedWidget {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QStackedWidget{bcthis0}
}
func (*QStackedWidget) NewFromPointer(cthis unsafe.Pointer) *QStackedWidget {
	return NewQStackedWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QStackedWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStackedWidget(QWidget *)
func NewQStackedWidget(parent *QWidget /*777 QWidget **/) *QStackedWidget {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidgetC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStackedWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStackedWidget()
func DeleteQStackedWidget(this *QStackedWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:62
// index:0
// Public Visibility=Default Availability=Available
// [4] int addWidget(QWidget *)
func (this *QStackedWidget) AddWidget(w *QWidget /*777 QWidget **/) int {
	var convArg0 = w.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget9addWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertWidget(int, QWidget *)
func (this *QStackedWidget) InsertWidget(index int, w *QWidget /*777 QWidget **/) int {
	var convArg1 = w.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget12insertWidgetEiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeWidget(QWidget *)
func (this *QStackedWidget) RemoveWidget(w *QWidget /*777 QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget12removeWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * currentWidget()
func (this *QStackedWidget) CurrentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget13currentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex()
func (this *QStackedWidget) CurrentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(QWidget *)
func (this *QStackedWidget) IndexOf(arg0 *QWidget /*777 QWidget **/) int {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget7indexOfEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget(int)
func (this *QStackedWidget) Widget(arg0 int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget6widgetEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QStackedWidget) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QStackedWidget5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)
func (this *QStackedWidget) SetCurrentIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget15setCurrentIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentWidget(QWidget *)
func (this *QStackedWidget) SetCurrentWidget(w *QWidget /*777 QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget16setCurrentWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(int)
func (this *QStackedWidget) CurrentChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget14currentChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void widgetRemoved(int)
func (this *QStackedWidget) WidgetRemoved(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget13widgetRemovedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedwidget.h:82
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QStackedWidget) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStackedWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

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
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
