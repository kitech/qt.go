package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QTextListFormat struct {
	*QTextFormat
}

func (this *QTextListFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextFormat.GetCthis()
	}
}
func (this *QTextListFormat) SetCthis(cthis unsafe.Pointer) {
	this.QTextFormat = NewQTextFormatFromPointer(cthis)
}
func NewQTextListFormatFromPointer(cthis unsafe.Pointer) *QTextListFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextListFormat{bcthis0}
}
func (*QTextListFormat) NewFromPointer(cthis unsafe.Pointer) *QTextListFormat {
	return NewQTextListFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:681
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextListFormat()
func NewQTextListFormat() *QTextListFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextListFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextListFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:714
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextListFormat(const QTextFormat &)
func NewQTextListFormat_1(fmt *QTextFormat) *QTextListFormat {
	var convArg0 = fmt.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormatC2ERK11QTextFormat", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextListFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextListFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:683
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextListFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextListFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:697
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setStyle(enum QTextListFormat::Style)
func (this *QTextListFormat) SetStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormat8setStyleENS_5StyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:698
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextListFormat::Style style()
func (this *QTextListFormat) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextListFormat5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:701
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setIndent(int)
func (this *QTextListFormat) SetIndent(indent int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormat9setIndentEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), indent)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:702
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int indent()
func (this *QTextListFormat) Indent() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextListFormat6indentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:705
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setNumberPrefix(const QString &)
func (this *QTextListFormat) SetNumberPrefix(numberPrefix string) {
	var tmpArg0 = qtcore.NewQString_5(numberPrefix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormat15setNumberPrefixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:706
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString numberPrefix()
func (this *QTextListFormat) NumberPrefix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextListFormat12numberPrefixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextformat.h:709
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setNumberSuffix(const QString &)
func (this *QTextListFormat) SetNumberSuffix(numberSuffix string) {
	var tmpArg0 = qtcore.NewQString_5(numberSuffix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormat15setNumberSuffixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:710
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString numberSuffix()
func (this *QTextListFormat) NumberSuffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextListFormat12numberSuffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

func DeleteQTextListFormat(this *QTextListFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextListFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QTextListFormat__Style = int

const QTextListFormat__ListDisc QTextListFormat__Style = -1
const QTextListFormat__ListCircle QTextListFormat__Style = -2
const QTextListFormat__ListSquare QTextListFormat__Style = -3
const QTextListFormat__ListDecimal QTextListFormat__Style = -4
const QTextListFormat__ListLowerAlpha QTextListFormat__Style = -5
const QTextListFormat__ListUpperAlpha QTextListFormat__Style = -6
const QTextListFormat__ListLowerRoman QTextListFormat__Style = -7
const QTextListFormat__ListUpperRoman QTextListFormat__Style = -8
const QTextListFormat__ListStyleUndefined QTextListFormat__Style = 0

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
