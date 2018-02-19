package qtwidgets

// /usr/include/qt/QtWidgets/qtablewidget.h
// #include <qtablewidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 77
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

type QTableWidgetSelectionRange struct {
	*qtrt.CObject
}
type QTableWidgetSelectionRange_ITF interface {
	QTableWidgetSelectionRange_PTR() *QTableWidgetSelectionRange
}

func (ptr *QTableWidgetSelectionRange) QTableWidgetSelectionRange_PTR() *QTableWidgetSelectionRange {
	return ptr
}

func (this *QTableWidgetSelectionRange) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTableWidgetSelectionRange) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTableWidgetSelectionRangeFromPointer(cthis unsafe.Pointer) *QTableWidgetSelectionRange {
	return &QTableWidgetSelectionRange{&qtrt.CObject{cthis}}
}
func (*QTableWidgetSelectionRange) NewFromPointer(cthis unsafe.Pointer) *QTableWidgetSelectionRange {
	return NewQTableWidgetSelectionRangeFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTableWidgetSelectionRange()
func NewQTableWidgetSelectionRange() *QTableWidgetSelectionRange {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QTableWidgetSelectionRangeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTableWidgetSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTableWidgetSelectionRange)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:56
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTableWidgetSelectionRange(int, int, int, int)
func NewQTableWidgetSelectionRange_1(top int, left int, bottom int, right int) *QTableWidgetSelectionRange {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QTableWidgetSelectionRangeC2Eiiii", qtrt.FFI_TYPE_POINTER, top, left, bottom, right)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTableWidgetSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTableWidgetSelectionRange)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTableWidgetSelectionRange()
func DeleteQTableWidgetSelectionRange(this *QTableWidgetSelectionRange) {
	rv, err := qtrt.InvokeQtFunc6("_ZN26QTableWidgetSelectionRangeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int topRow() const
func (this *QTableWidgetSelectionRange) TopRow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange6topRowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int bottomRow() const
func (this *QTableWidgetSelectionRange) BottomRow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange9bottomRowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int leftColumn() const
func (this *QTableWidgetSelectionRange) LeftColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange10leftColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int rightColumn() const
func (this *QTableWidgetSelectionRange) RightColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange11rightColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int rowCount() const
func (this *QTableWidgetSelectionRange) RowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange8rowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int columnCount() const
func (this *QTableWidgetSelectionRange) ColumnCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange11columnCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
