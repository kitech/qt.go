package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QAccessibleTableCellInterface struct {
	*qtrt.CObject
}
type QAccessibleTableCellInterface_ITF interface {
	QAccessibleTableCellInterface_PTR() *QAccessibleTableCellInterface
}

func (ptr *QAccessibleTableCellInterface) QAccessibleTableCellInterface_PTR() *QAccessibleTableCellInterface {
	return ptr
}

func (this *QAccessibleTableCellInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleTableCellInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleTableCellInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleTableCellInterface {
	return &QAccessibleTableCellInterface{&qtrt.CObject{cthis}}
}
func (*QAccessibleTableCellInterface) NewFromPointer(cthis unsafe.Pointer) *QAccessibleTableCellInterface {
	return NewQAccessibleTableCellInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:578
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTableCellInterface()
func DeleteQAccessibleTableCellInterface(this *QAccessibleTableCellInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN29QAccessibleTableCellInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:580
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isSelected()
func (this *QAccessibleTableCellInterface) IsSelected() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface10isSelectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qaccessible.h:584
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int columnIndex()
func (this *QAccessibleTableCellInterface) ColumnIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface11columnIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:585
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int rowIndex()
func (this *QAccessibleTableCellInterface) RowIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface8rowIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:586
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int columnExtent()
func (this *QAccessibleTableCellInterface) ColumnExtent() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface12columnExtentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:587
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int rowExtent()
func (this *QAccessibleTableCellInterface) RowExtent() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface9rowExtentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:589
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * table()
func (this *QAccessibleTableCellInterface) Table() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface5tableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
}

//  keep block end
