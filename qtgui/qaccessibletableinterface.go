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
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QAccessibleTableInterface struct {
	*qtrt.CObject
}
type QAccessibleTableInterface_ITF interface {
	QAccessibleTableInterface_PTR() *QAccessibleTableInterface
}

func (ptr *QAccessibleTableInterface) QAccessibleTableInterface_PTR() *QAccessibleTableInterface {
	return ptr
}

func (this *QAccessibleTableInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleTableInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleTableInterfaceFromPointer(cthis unsafe.Pointer) *QAccessibleTableInterface {
	return &QAccessibleTableInterface{&qtrt.CObject{cthis}}
}
func (*QAccessibleTableInterface) NewFromPointer(cthis unsafe.Pointer) *QAccessibleTableInterface {
	return NewQAccessibleTableInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:595
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTableInterface()
func DeleteQAccessibleTableInterface(this *QAccessibleTableInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAccessibleTableInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:597
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * caption()
func (this *QAccessibleTableInterface) Caption() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface7captionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:598
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * summary()
func (this *QAccessibleTableInterface) Summary() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface7summaryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:600
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * cellAt(int, int)
func (this *QAccessibleTableInterface) CellAt(row int, column int) *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface6cellAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:601
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int selectedCellCount()
func (this *QAccessibleTableInterface) SelectedCellCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface17selectedCellCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:604
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString columnDescription(int)
func (this *QAccessibleTableInterface) ColumnDescription(column int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface17columnDescriptionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:605
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString rowDescription(int)
func (this *QAccessibleTableInterface) RowDescription(row int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface14rowDescriptionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:606
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int selectedColumnCount()
func (this *QAccessibleTableInterface) SelectedColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface19selectedColumnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:607
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int selectedRowCount()
func (this *QAccessibleTableInterface) SelectedRowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface16selectedRowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:608
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int columnCount()
func (this *QAccessibleTableInterface) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:609
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int rowCount()
func (this *QAccessibleTableInterface) RowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface8rowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:612
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isColumnSelected(int)
func (this *QAccessibleTableInterface) IsColumnSelected(column int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface16isColumnSelectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qaccessible.h:613
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isRowSelected(int)
func (this *QAccessibleTableInterface) IsRowSelected(row int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface13isRowSelectedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qaccessible.h:614
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool selectRow(int)
func (this *QAccessibleTableInterface) SelectRow(row int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAccessibleTableInterface9selectRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qaccessible.h:615
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool selectColumn(int)
func (this *QAccessibleTableInterface) SelectColumn(column int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAccessibleTableInterface12selectColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qaccessible.h:616
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool unselectRow(int)
func (this *QAccessibleTableInterface) UnselectRow(row int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAccessibleTableInterface11unselectRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qaccessible.h:617
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool unselectColumn(int)
func (this *QAccessibleTableInterface) UnselectColumn(column int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAccessibleTableInterface14unselectColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qaccessible.h:619
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void modelChange(QAccessibleTableModelChangeEvent *)
func (this *QAccessibleTableInterface) ModelChange(event QAccessibleTableModelChangeEvent_ITF /*777 QAccessibleTableModelChangeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QAccessibleTableModelChangeEvent_PTR() != nil {
		convArg0 = event.QAccessibleTableModelChangeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAccessibleTableInterface11modelChangeEP32QAccessibleTableModelChangeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
}

//  keep block end
