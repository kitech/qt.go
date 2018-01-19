//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

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
type QAccessibleTableModelChangeEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qaccessible.h:932
// index:0
// inline
// void QAccessibleTableModelChangeEvent(class QObject *, enum QAccessibleTableModelChangeEvent::ModelChangeType)
func NewQAccessibleTableModelChangeEvent(obj unsafe.Pointer, changeType int) *QAccessibleTableModelChangeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEventC2EP7QObjectNS_15ModelChangeTypeE", ffiqt.FFI_TYPE_VOID, cthis, obj, &changeType)
	gopp.ErrPrint(err, rv)
	return &QAccessibleTableModelChangeEvent{cthis}
}

// /usr/include/qt/QtGui/qaccessible.h:939
// index:1
// inline
// void QAccessibleTableModelChangeEvent(class QAccessibleInterface *, enum QAccessibleTableModelChangeEvent::ModelChangeType)
func NewQAccessibleTableModelChangeEvent_1(iface unsafe.Pointer, changeType int) *QAccessibleTableModelChangeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEventC2EP20QAccessibleInterfaceNS_15ModelChangeTypeE", ffiqt.FFI_TYPE_VOID, cthis, iface, &changeType)
	gopp.ErrPrint(err, rv)
	return &QAccessibleTableModelChangeEvent{cthis}
}

// /usr/include/qt/QtGui/qaccessible.h:947
// index:0
// virtual
// void ~QAccessibleTableModelChangeEvent()
func DeleteQAccessibleTableModelChangeEvent(*QAccessibleTableModelChangeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:949
// index:0
// inline
// void setModelChangeType(enum QAccessibleTableModelChangeEvent::ModelChangeType)
func (this *QAccessibleTableModelChangeEvent) SetModelChangeType(changeType int) {
	// 0: (, QAccessibleTableModelChangeEvent::ModelChangeType changeType), (&changeType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEvent18setModelChangeTypeENS_15ModelChangeTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &changeType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:950
// index:0
// inline
// QAccessibleTableModelChangeEvent::ModelChangeType modelChangeType()
func (this *QAccessibleTableModelChangeEvent) ModelChangeType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK32QAccessibleTableModelChangeEvent15modelChangeTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:952
// index:0
// inline
// void setFirstRow(int)
func (this *QAccessibleTableModelChangeEvent) SetFirstRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEvent11setFirstRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:953
// index:0
// inline
// void setFirstColumn(int)
func (this *QAccessibleTableModelChangeEvent) SetFirstColumn(col int) {
	// 0: (, int col), (&col)
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &col)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:954
// index:0
// inline
// void setLastRow(int)
func (this *QAccessibleTableModelChangeEvent) SetLastRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEvent10setLastRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:955
// index:0
// inline
// void setLastColumn(int)
func (this *QAccessibleTableModelChangeEvent) SetLastColumn(col int) {
	// 0: (, int col), (&col)
	rv, err := ffiqt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEvent13setLastColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &col)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:956
// index:0
// inline
// int firstRow()
func (this *QAccessibleTableModelChangeEvent) FirstRow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK32QAccessibleTableModelChangeEvent8firstRowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:957
// index:0
// inline
// int firstColumn()
func (this *QAccessibleTableModelChangeEvent) FirstColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK32QAccessibleTableModelChangeEvent11firstColumnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:958
// index:0
// inline
// int lastRow()
func (this *QAccessibleTableModelChangeEvent) LastRow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK32QAccessibleTableModelChangeEvent7lastRowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:959
// index:0
// inline
// int lastColumn()
func (this *QAccessibleTableModelChangeEvent) LastColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK32QAccessibleTableModelChangeEvent10lastColumnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
