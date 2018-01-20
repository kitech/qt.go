//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QAccessibleTableInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleTableInterface) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qaccessible.h:595
// index:0
// virtual
// void ~QAccessibleTableInterface()
func DeleteQAccessibleTableInterface(*QAccessibleTableInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:597
// index:0
// pure virtual
// QAccessibleInterface * caption()
func (this *QAccessibleTableInterface) Caption() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface7captionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:598
// index:0
// pure virtual
// QAccessibleInterface * summary()
func (this *QAccessibleTableInterface) Summary() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface7summaryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:600
// index:0
// pure virtual
// QAccessibleInterface * cellAt(int, int)
func (this *QAccessibleTableInterface) CellAt(row int, column int) {
	// 0: (, row int, column int), (&row, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface6cellAtEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:601
// index:0
// pure virtual
// int selectedCellCount()
func (this *QAccessibleTableInterface) SelectedCellCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface17selectedCellCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:602
// index:0
// pure virtual
// QList<QAccessibleInterface *> selectedCells()
func (this *QAccessibleTableInterface) SelectedCells() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface13selectedCellsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:604
// index:0
// pure virtual
// QString columnDescription(int)
func (this *QAccessibleTableInterface) ColumnDescription(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface17columnDescriptionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:605
// index:0
// pure virtual
// QString rowDescription(int)
func (this *QAccessibleTableInterface) RowDescription(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface14rowDescriptionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:606
// index:0
// pure virtual
// int selectedColumnCount()
func (this *QAccessibleTableInterface) SelectedColumnCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface19selectedColumnCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:607
// index:0
// pure virtual
// int selectedRowCount()
func (this *QAccessibleTableInterface) SelectedRowCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface16selectedRowCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:608
// index:0
// pure virtual
// int columnCount()
func (this *QAccessibleTableInterface) ColumnCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface11columnCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:609
// index:0
// pure virtual
// int rowCount()
func (this *QAccessibleTableInterface) RowCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface8rowCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:610
// index:0
// pure virtual
// QList<int> selectedColumns()
func (this *QAccessibleTableInterface) SelectedColumns() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface15selectedColumnsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:611
// index:0
// pure virtual
// QList<int> selectedRows()
func (this *QAccessibleTableInterface) SelectedRows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface12selectedRowsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:612
// index:0
// pure virtual
// bool isColumnSelected(int)
func (this *QAccessibleTableInterface) IsColumnSelected(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface16isColumnSelectedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:613
// index:0
// pure virtual
// bool isRowSelected(int)
func (this *QAccessibleTableInterface) IsRowSelected(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleTableInterface13isRowSelectedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:614
// index:0
// pure virtual
// bool selectRow(int)
func (this *QAccessibleTableInterface) SelectRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterface9selectRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:615
// index:0
// pure virtual
// bool selectColumn(int)
func (this *QAccessibleTableInterface) SelectColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterface12selectColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:616
// index:0
// pure virtual
// bool unselectRow(int)
func (this *QAccessibleTableInterface) UnselectRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterface11unselectRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:617
// index:0
// pure virtual
// bool unselectColumn(int)
func (this *QAccessibleTableInterface) UnselectColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterface14unselectColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:619
// index:0
// pure virtual
// void modelChange(class QAccessibleTableModelChangeEvent *)
func (this *QAccessibleTableInterface) ModelChange(event unsafe.Pointer) {
	// 0: (, event QAccessibleTableModelChangeEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleTableInterface11modelChangeEP32QAccessibleTableModelChangeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

//  body block end
