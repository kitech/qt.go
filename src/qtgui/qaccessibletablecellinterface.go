//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

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
	return this.Cthis
}
func NewQAccessibleTableCellInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleTableCellInterface {
	return &QAccessibleTableCellInterface{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qaccessible.h:578
// index:0
// Public virtual
// void ~QAccessibleTableCellInterface()
func DeleteQAccessibleTableCellInterface(*QAccessibleTableCellInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN29QAccessibleTableCellInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:580
// index:0
// Public pure virtual
// bool isSelected()
func (this *QAccessibleTableCellInterface) IsSelected() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface10isSelectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:582
// index:0
// Public pure virtual
// QList<QAccessibleInterface *> columnHeaderCells()
func (this *QAccessibleTableCellInterface) ColumnHeaderCells() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface17columnHeaderCellsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:583
// index:0
// Public pure virtual
// QList<QAccessibleInterface *> rowHeaderCells()
func (this *QAccessibleTableCellInterface) RowHeaderCells() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface14rowHeaderCellsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:584
// index:0
// Public pure virtual
// int columnIndex()
func (this *QAccessibleTableCellInterface) ColumnIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface11columnIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:585
// index:0
// Public pure virtual
// int rowIndex()
func (this *QAccessibleTableCellInterface) RowIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface8rowIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:586
// index:0
// Public pure virtual
// int columnExtent()
func (this *QAccessibleTableCellInterface) ColumnExtent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface12columnExtentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:587
// index:0
// Public pure virtual
// int rowExtent()
func (this *QAccessibleTableCellInterface) RowExtent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface9rowExtentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:589
// index:0
// Public pure virtual
// QAccessibleInterface * table()
func (this *QAccessibleTableCellInterface) Table() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK29QAccessibleTableCellInterface5tableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
