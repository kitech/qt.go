//  header block begin
// /usr/include/qt/QtWidgets/qtablewidget.h
// #include <qtablewidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 75
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	return this.Cthis
}
func NewQTableWidgetSelectionRangeFromPointer(cthis unsafe.Pointer) *QTableWidgetSelectionRange {
	return &QTableWidgetSelectionRange{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qtablewidget.h:55
// index:0
// Public
// void QTableWidgetSelectionRange()
func NewQTableWidgetSelectionRange() *QTableWidgetSelectionRange {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QTableWidgetSelectionRangeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetSelectionRangeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:56
// index:1
// Public
// void QTableWidgetSelectionRange(int, int, int, int)
func NewQTableWidgetSelectionRange_1(top int, left int, bottom int, right int) *QTableWidgetSelectionRange {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QTableWidgetSelectionRangeC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, &top, &left, &bottom, &right)
	gopp.ErrPrint(err, rv)
	gothis := NewQTableWidgetSelectionRangeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qtablewidget.h:58
// index:0
// Public
// void ~QTableWidgetSelectionRange()
func DeleteQTableWidgetSelectionRange(*QTableWidgetSelectionRange) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QTableWidgetSelectionRangeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qtablewidget.h:60
// index:0
// Public inline
// int topRow()
func (this *QTableWidgetSelectionRange) TopRow() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange6topRowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:61
// index:0
// Public inline
// int bottomRow()
func (this *QTableWidgetSelectionRange) BottomRow() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange9bottomRowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:62
// index:0
// Public inline
// int leftColumn()
func (this *QTableWidgetSelectionRange) LeftColumn() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange10leftColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:63
// index:0
// Public inline
// int rightColumn()
func (this *QTableWidgetSelectionRange) RightColumn() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange11rightColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:64
// index:0
// Public inline
// int rowCount()
func (this *QTableWidgetSelectionRange) RowCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange8rowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qtablewidget.h:65
// index:0
// Public inline
// int columnCount()
func (this *QTableWidgetSelectionRange) ColumnCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QTableWidgetSelectionRange11columnCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
