package qtwidgets

// /usr/include/qt/QtWidgets/qdatetimeedit.h
// #include <qdatetimeedit.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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

type QDateEdit struct {
	*QDateTimeEdit
}

func (this *QDateEdit) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDateTimeEdit.GetCthis()
	}
}
func (this *QDateEdit) SetCthis(cthis unsafe.Pointer) {
	this.QDateTimeEdit = NewQDateTimeEditFromPointer(cthis)
}
func NewQDateEditFromPointer(cthis unsafe.Pointer) *QDateEdit {
	bcthis0 := NewQDateTimeEditFromPointer(cthis)
	return &QDateEdit{bcthis0}
}
func (*QDateEdit) NewFromPointer(cthis unsafe.Pointer) *QDateEdit {
	return NewQDateEditFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:217
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QDateEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateEdit10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDateEdit(QWidget *)
func NewQDateEdit(parent *QWidget /*777 QWidget **/) *QDateEdit {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:221
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDateEdit(const QDate &, QWidget *)
func NewQDateEdit_1(date *qtcore.QDate, parent *QWidget /*777 QWidget **/) *QDateEdit {
	var convArg0 = date.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateEditC2ERK5QDateP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQDateEditFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:222
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDateEdit()
func DeleteQDateEdit(this *QDateEdit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateEditD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void userDateChanged(const QDate &)
func (this *QDateEdit) UserDateChanged(date *qtcore.QDate) {
	var convArg0 = date.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateEdit15userDateChangedERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
