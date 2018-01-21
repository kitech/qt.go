package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
func NewQTextListFormatFromPointer(cthis unsafe.Pointer) *QTextListFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextListFormat{bcthis0}
}

// /usr/include/qt/QtGui/qtextformat.h:681
// index:0
// Public
// void QTextListFormat()
func NewQTextListFormat() *QTextListFormat {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextListFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:714
// index:1
// Protected
// void QTextListFormat(const class QTextFormat &)
func NewQTextListFormat_1(fmt *QTextFormat) *QTextListFormat {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = fmt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextListFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:683
// index:0
// Public inline
// bool isValid()
func (this *QTextListFormat) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextListFormat7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:697
// index:0
// Public inline
// void setStyle(enum QTextListFormat::Style)
func (this *QTextListFormat) SetStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormat8setStyleENS_5StyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:698
// index:0
// Public inline
// QTextListFormat::Style style()
func (this *QTextListFormat) Style() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextListFormat5styleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:701
// index:0
// Public inline
// void setIndent(int)
func (this *QTextListFormat) SetIndent(indent int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormat9setIndentEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &indent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:702
// index:0
// Public inline
// int indent()
func (this *QTextListFormat) Indent() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextListFormat6indentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextformat.h:705
// index:0
// Public inline
// void setNumberPrefix(const class QString &)
func (this *QTextListFormat) SetNumberPrefix(numberPrefix *qtcore.QString) {
	var convArg0 = numberPrefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormat15setNumberPrefixERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:706
// index:0
// Public inline
// QString numberPrefix()
func (this *QTextListFormat) NumberPrefix() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextListFormat12numberPrefixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:709
// index:0
// Public inline
// void setNumberSuffix(const class QString &)
func (this *QTextListFormat) SetNumberSuffix(numberSuffix *qtcore.QString) {
	var convArg0 = numberSuffix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormat15setNumberSuffixERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:710
// index:0
// Public inline
// QString numberSuffix()
func (this *QTextListFormat) NumberSuffix() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextListFormat12numberSuffixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
