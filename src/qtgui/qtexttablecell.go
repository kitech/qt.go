//  header block begin
// /usr/include/qt/QtGui/qtexttable.h
// #include <qtexttable.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QTextTableCell struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtexttable.h:57
// index:0
// inline
// void QTextTableCell()
func NewQTextTableCell() *QTextTableCell {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTextTableCellC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextTableCell{cthis}
}

// /usr/include/qt/QtGui/qtexttable.h:58
// index:0
// inline
// void ~QTextTableCell()
func DeleteQTextTableCell(*QTextTableCell) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTextTableCellD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:63
// index:0
// void setFormat(const class QTextCharFormat &)
func (this *QTextTableCell) SetFormat(format unsafe.Pointer) {
	// 0: (, const QTextCharFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTextTableCell9setFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:64
// index:0
// QTextCharFormat format()
func (this *QTextTableCell) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell6formatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:66
// index:0
// int row()
func (this *QTextTableCell) Row() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell3rowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:67
// index:0
// int column()
func (this *QTextTableCell) Column() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell6columnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:69
// index:0
// int rowSpan()
func (this *QTextTableCell) RowSpan() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell7rowSpanEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:70
// index:0
// int columnSpan()
func (this *QTextTableCell) ColumnSpan() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell10columnSpanEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:72
// index:0
// inline
// bool isValid()
func (this *QTextTableCell) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:74
// index:0
// QTextCursor firstCursorPosition()
func (this *QTextTableCell) FirstCursorPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell19firstCursorPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:75
// index:0
// QTextCursor lastCursorPosition()
func (this *QTextTableCell) LastCursorPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell18lastCursorPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:76
// index:0
// int firstPosition()
func (this *QTextTableCell) FirstPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell13firstPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:77
// index:0
// int lastPosition()
func (this *QTextTableCell) LastPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell12lastPositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:84
// index:0
// QTextFrame::iterator begin()
func (this *QTextTableCell) Begin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell5beginEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:85
// index:0
// QTextFrame::iterator end()
func (this *QTextTableCell) End() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell3endEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:87
// index:0
// int tableCellFormatIndex()
func (this *QTextTableCell) TableCellFormatIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell20tableCellFormatIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
