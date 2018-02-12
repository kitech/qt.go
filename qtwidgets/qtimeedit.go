package qtwidgets

// /usr/include/qt/QtWidgets/qdatetimeedit.h
// #include <qdatetimeedit.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 70
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
func (this *QTimeEdit) SetCthis(cthis unsafe.Pointer) {
	this.QDateTimeEdit = NewQDateTimeEditFromPointer(cthis)
}
func NewQTimeEditFromPointer(cthis unsafe.Pointer) *QTimeEdit {
	bcthis0 := NewQDateTimeEditFromPointer(cthis)
	return &QTimeEdit{bcthis0}
}
func (*QTimeEdit) NewFromPointer(cthis unsafe.Pointer) *QTimeEdit {
	return NewQTimeEditFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:204
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTimeEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QTimeEdit10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTimeEdit(QWidget *)
func NewQTimeEdit(parent *QWidget /*777 QWidget **/) *QTimeEdit {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:208
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTimeEdit(const QTime &, QWidget *)
func NewQTimeEdit_1(time *qtcore.QTime, parent *QWidget /*777 QWidget **/) *QTimeEdit {
	var convArg0 = time.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeEditC2ERK5QTimeP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTimeEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:209
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTimeEdit()
func DeleteQTimeEdit(this *QTimeEdit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeEditD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:212
// index:0
// Public Visibility=Default Availability=Available
// [-2] void userTimeChanged(const QTime &)
func (this *QTimeEdit) UserTimeChanged(time *qtcore.QTime) {
	var convArg0 = time.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QTimeEdit15userTimeChangedERK5QTime", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
