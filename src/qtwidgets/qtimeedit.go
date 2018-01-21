package qtwidgets

// /usr/include/qt/QtWidgets/qdatetimeedit.h
// #include <qdatetimeedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 70
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
type QTimeEdit struct {
	*QDateTimeEdit
}

func (this *QTimeEdit) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDateTimeEdit.GetCthis()
	}
}
func NewQTimeEditFromPointer(cthis unsafe.Pointer) *QTimeEdit {
	bcthis0 := NewQDateTimeEditFromPointer(cthis)
	return &QTimeEdit{bcthis0}
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:204
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTimeEdit) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QTimeEdit10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:207
// index:0
// Public
// void QTimeEdit(class QWidget *)
func NewQTimeEdit(parent *QWidget /*444 QWidget **/) *QTimeEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeEditC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:208
// index:1
// Public
// void QTimeEdit(const class QTime &, class QWidget *)
func NewQTimeEdit_1(time *qtcore.QTime, parent *QWidget /*444 QWidget **/) *QTimeEdit {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = time.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeEditC2ERK5QTimeP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQTimeEditFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:209
// index:0
// Public virtual
// void ~QTimeEdit()
func DeleteQTimeEdit(*QTimeEdit) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeEditD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:212
// index:0
// Public
// void userTimeChanged(const class QTime &)
func (this *QTimeEdit) UserTimeChanged(time *qtcore.QTime) {
	var convArg0 = time.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QTimeEdit15userTimeChangedERK5QTime", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
