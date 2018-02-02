package qtwidgets

// /usr/include/qt/QtWidgets/qtablewidget.h
// #include <qtablewidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 76
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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

type QTableWidgetSelectionRange struct {
	*qtrt.CObject
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QTableWidgetSelectionRangeC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTableWidgetSelectionRange)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:56
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTableWidgetSelectionRange(int, int, int, int)
func NewQTableWidgetSelectionRange_1(top int, left int, bottom int, right int) *QTableWidgetSelectionRange {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QTableWidgetSelectionRangeC2Eiiii", ffiqt.FFI_TYPE_POINTER, top, left, bottom, right)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetSelectionRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTableWidgetSelectionRange)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTableWidgetSelectionRange()
func DeleteQTableWidgetSelectionRange(this *QTableWidgetSelectionRange) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QTableWidgetSelectionRangeD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int topRow()
func (this *QTableWidgetSelectionRange) TopRow() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange6topRowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:61
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int bottomRow()
func (this *QTableWidgetSelectionRange) BottomRow() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange9bottomRowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int leftColumn()
func (this *QTableWidgetSelectionRange) LeftColumn() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange10leftColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int rightColumn()
func (this *QTableWidgetSelectionRange) RightColumn() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange11rightColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int rowCount()
func (this *QTableWidgetSelectionRange) RowCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange8rowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qtablewidget.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int columnCount()
func (this *QTableWidgetSelectionRange) ColumnCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
