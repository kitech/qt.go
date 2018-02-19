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
// extern C begin: 5
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

type QAccessibleTableModelChangeEvent struct {
	*QAccessibleEvent
}
type QAccessibleTableModelChangeEvent_ITF interface {
	QAccessibleEvent_ITF
	QAccessibleTableModelChangeEvent_PTR() *QAccessibleTableModelChangeEvent
}

func (ptr *QAccessibleTableModelChangeEvent) QAccessibleTableModelChangeEvent_PTR() *QAccessibleTableModelChangeEvent {
	return ptr
}

func (this *QAccessibleTableModelChangeEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAccessibleEvent.GetCthis()
	}
}
func (this *QAccessibleTableModelChangeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QAccessibleEvent = NewQAccessibleEventFromPointer(cthis)
}
func NewQAccessibleTableModelChangeEventFromPointer(cthis unsafe.Pointer) *QAccessibleTableModelChangeEvent {
	bcthis0 := NewQAccessibleEventFromPointer(cthis)
	return &QAccessibleTableModelChangeEvent{bcthis0}
}
func (*QAccessibleTableModelChangeEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleTableModelChangeEvent {
	return NewQAccessibleTableModelChangeEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:932
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTableModelChangeEvent(QObject *, enum QAccessibleTableModelChangeEvent::ModelChangeType)
func NewQAccessibleTableModelChangeEvent(obj qtcore.QObject_ITF /*777 QObject **/, changeType int) *QAccessibleTableModelChangeEvent {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEventC2EP7QObjectNS_15ModelChangeTypeE", qtrt.FFI_TYPE_POINTER, convArg0, changeType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTableModelChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTableModelChangeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:939
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleTableModelChangeEvent(QAccessibleInterface *, enum QAccessibleTableModelChangeEvent::ModelChangeType)
func NewQAccessibleTableModelChangeEvent_1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, changeType int) *QAccessibleTableModelChangeEvent {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QAccessibleInterface_PTR() != nil {
		convArg0 = iface.QAccessibleInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEventC2EP20QAccessibleInterfaceNS_15ModelChangeTypeE", qtrt.FFI_TYPE_POINTER, convArg0, changeType)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleTableModelChangeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleTableModelChangeEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:947
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleTableModelChangeEvent()
func DeleteQAccessibleTableModelChangeEvent(this *QAccessibleTableModelChangeEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:949
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setModelChangeType(enum QAccessibleTableModelChangeEvent::ModelChangeType)
func (this *QAccessibleTableModelChangeEvent) SetModelChangeType(changeType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEvent18setModelChangeTypeENS_15ModelChangeTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), changeType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:950
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QAccessibleTableModelChangeEvent::ModelChangeType modelChangeType() const
func (this *QAccessibleTableModelChangeEvent) ModelChangeType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QAccessibleTableModelChangeEvent15modelChangeTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:952
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFirstRow(int)
func (this *QAccessibleTableModelChangeEvent) SetFirstRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEvent11setFirstRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:953
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFirstColumn(int)
func (this *QAccessibleTableModelChangeEvent) SetFirstColumn(col int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), col)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:954
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLastRow(int)
func (this *QAccessibleTableModelChangeEvent) SetLastRow(row int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEvent10setLastRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:955
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLastColumn(int)
func (this *QAccessibleTableModelChangeEvent) SetLastColumn(col int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN32QAccessibleTableModelChangeEvent13setLastColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), col)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:956
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int firstRow() const
func (this *QAccessibleTableModelChangeEvent) FirstRow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QAccessibleTableModelChangeEvent8firstRowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:957
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int firstColumn() const
func (this *QAccessibleTableModelChangeEvent) FirstColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QAccessibleTableModelChangeEvent11firstColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:958
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lastRow() const
func (this *QAccessibleTableModelChangeEvent) LastRow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QAccessibleTableModelChangeEvent7lastRowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:959
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int lastColumn() const
func (this *QAccessibleTableModelChangeEvent) LastColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK32QAccessibleTableModelChangeEvent10lastColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

type QAccessibleTableModelChangeEvent__ModelChangeType = int

const QAccessibleTableModelChangeEvent__ModelReset QAccessibleTableModelChangeEvent__ModelChangeType = 0
const QAccessibleTableModelChangeEvent__DataChanged QAccessibleTableModelChangeEvent__ModelChangeType = 1
const QAccessibleTableModelChangeEvent__RowsInserted QAccessibleTableModelChangeEvent__ModelChangeType = 2
const QAccessibleTableModelChangeEvent__ColumnsInserted QAccessibleTableModelChangeEvent__ModelChangeType = 3
const QAccessibleTableModelChangeEvent__RowsRemoved QAccessibleTableModelChangeEvent__ModelChangeType = 4
const QAccessibleTableModelChangeEvent__ColumnsRemoved QAccessibleTableModelChangeEvent__ModelChangeType = 5

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
