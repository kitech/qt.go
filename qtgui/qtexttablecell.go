package qtgui

// /usr/include/qt/QtGui/qtexttable.h
// #include <qtexttable.h>
// #include <QtGui>

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

type QTextTableCell struct {
	*qtrt.CObject
}

func (this *QTextTableCell) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextTableCell) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextTableCellFromPointer(cthis unsafe.Pointer) *QTextTableCell {
	return &QTextTableCell{&qtrt.CObject{cthis}}
}
func (*QTextTableCell) NewFromPointer(cthis unsafe.Pointer) *QTextTableCell {
	return NewQTextTableCellFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtexttable.h:57
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextTableCell()
func NewQTextTableCell() *QTextTableCell {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTextTableCellC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextTableCellFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextTableCell)
	return gothis
}

// /usr/include/qt/QtGui/qtexttable.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QTextTableCell()
func DeleteQTextTableCell(this *QTextTableCell) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTextTableCellD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtexttable.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QTextCharFormat &)
func (this *QTextTableCell) SetFormat(format *QTextCharFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QTextTableCell9setFormatERK15QTextCharFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:64
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat format()
func (this *QTextTableCell) Format() *QTextCharFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] int row()
func (this *QTextTableCell) Row() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell3rowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int column()
func (this *QTextTableCell) Column() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell6columnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowSpan()
func (this *QTextTableCell) RowSpan() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell7rowSpanEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnSpan()
func (this *QTextTableCell) ColumnSpan() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell10columnSpanEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextTableCell) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtexttable.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor firstCursorPosition()
func (this *QTextTableCell) FirstCursorPosition() *QTextCursor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell19firstCursorPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor lastCursorPosition()
func (this *QTextTableCell) LastCursorPosition() *QTextCursor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell18lastCursorPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int firstPosition()
func (this *QTextTableCell) FirstPosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell13firstPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastPosition()
func (this *QTextTableCell) LastPosition() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell12lastPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:84
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextFrame::iterator begin()
func (this *QTextTableCell) Begin() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell5beginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtexttable.h:85
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextFrame::iterator end()
func (this *QTextTableCell) End() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell3endEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtexttable.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int tableCellFormatIndex()
func (this *QTextTableCell) TableCellFormatIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QTextTableCell20tableCellFormatIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
