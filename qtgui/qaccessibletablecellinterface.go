package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QAccessibleTableCellInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleTableCellInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleTableCellInterface) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
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
func DeleteQAccessibleTableCellInterface(*QAccessibleTableCellInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QAccessibleTableCellInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:580
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isSelected()
func (this *QAccessibleTableCellInterface) IsSelected() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface10isSelectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qaccessible.h:584
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int columnIndex()
func (this *QAccessibleTableCellInterface) ColumnIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface11columnIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:585
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int rowIndex()
func (this *QAccessibleTableCellInterface) RowIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface8rowIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:586
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int columnExtent()
func (this *QAccessibleTableCellInterface) ColumnExtent() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface12columnExtentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:587
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int rowExtent()
func (this *QAccessibleTableCellInterface) RowExtent() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface9rowExtentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:589
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * table()
func (this *QAccessibleTableCellInterface) Table() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface5tableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
