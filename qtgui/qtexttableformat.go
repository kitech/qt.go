package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
type QTextTableFormat struct {
	*QTextFrameFormat
}

func (this *QTextTableFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextFrameFormat.GetCthis()
	}
}
func (this *QTextTableFormat) SetCthis(cthis unsafe.Pointer) {
	this.QTextFrameFormat = NewQTextFrameFormatFromPointer(cthis)
}
func NewQTextTableFormatFromPointer(cthis unsafe.Pointer) *QTextTableFormat {
	bcthis0 := NewQTextFrameFormatFromPointer(cthis)
	return &QTextTableFormat{bcthis0}
}
func (*QTextTableFormat) NewFromPointer(cthis unsafe.Pointer) *QTextTableFormat {
	return NewQTextTableFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:887
// index:0
// Public
// void QTextTableFormat()
func NewQTextTableFormat() *QTextTableFormat {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextTableFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:923
// index:1
// Protected
// void QTextTableFormat(const QTextFormat &)
func NewQTextTableFormat_1(fmt *QTextFormat) *QTextTableFormat {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = fmt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextTableFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:889
// index:0
// Public inline
// bool isValid()
func (this *QTextTableFormat) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:891
// index:0
// Public inline
// int columns()
func (this *QTextTableFormat) Columns() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat7columnsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextformat.h:893
// index:0
// Public inline
// void setColumns(int)
func (this *QTextTableFormat) SetColumns(columns int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormat10setColumnsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:901
// index:0
// Public inline
// void clearColumnWidthConstraints()
func (this *QTextTableFormat) ClearColumnWidthConstraints() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormat27clearColumnWidthConstraintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:904
// index:0
// Public inline
// qreal cellSpacing()
func (this *QTextTableFormat) CellSpacing() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat11cellSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextformat.h:906
// index:0
// Public inline
// void setCellSpacing(qreal)
func (this *QTextTableFormat) SetCellSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormat14setCellSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:909
// index:0
// Public inline
// qreal cellPadding()
func (this *QTextTableFormat) CellPadding() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat11cellPaddingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextformat.h:911
// index:0
// Public inline
// void setCellPadding(qreal)
func (this *QTextTableFormat) SetCellPadding(padding float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormat14setCellPaddingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:914
// index:0
// Public inline
// Qt::Alignment alignment()
func (this *QTextTableFormat) Alignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat9alignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:917
// index:0
// Public inline
// void setHeaderRowCount(int)
func (this *QTextTableFormat) SetHeaderRowCount(count int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextTableFormat17setHeaderRowCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:919
// index:0
// Public inline
// int headerRowCount()
func (this *QTextTableFormat) HeaderRowCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextTableFormat14headerRowCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
