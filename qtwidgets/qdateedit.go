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
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QDateEdit struct {
	*QDateTimeEdit
}
type QDateEdit_ITF interface {
	QDateTimeEdit_ITF
	QDateEdit_PTR() *QDateEdit
}

func (ptr *QDateEdit) QDateEdit_PTR() *QDateEdit { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDateEdit) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDateEdit10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDateEdit(QWidget *)

/*

 */
func NewQDateEdit(parent QWidget_ITF /*777 QWidget **/) *QDateEdit {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDateEdit(QWidget *)

/*

 */
func NewQDateEdit__() *QDateEdit {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateEditC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:221
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDateEdit(const QDate &, QWidget *)

/*

 */
func NewQDateEdit_1(date qtcore.QDate_ITF, parent QWidget_ITF /*777 QWidget **/) *QDateEdit {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateEditC2ERK5QDateP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:221
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDateEdit(const QDate &, QWidget *)

/*

 */
func NewQDateEdit_1_(date qtcore.QDate_ITF) *QDateEdit {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	// arg: 1, QWidget *=Pointer, QWidget=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateEditC2ERK5QDateP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDateEditFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDateEdit")
	return gothis
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:222
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDateEdit()

/*

 */
func DeleteQDateEdit(this *QDateEdit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateEditD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdatetimeedit.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void userDateChanged(const QDate &)

/*

 */
func (this *QDateEdit) UserDateChanged(date qtcore.QDate_ITF) {
	var convArg0 unsafe.Pointer
	if date != nil && date.QDate_PTR() != nil {
		convArg0 = date.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDateEdit15userDateChangedERK5QDate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
		log.Println(123)
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
