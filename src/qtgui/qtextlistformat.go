//  header block begin
// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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
	return this.QTextFormat.GetCthis()
}

// /usr/include/qt/QtGui/qtextformat.h:681
// index:0
// void QTextListFormat()
func NewQTextListFormat() *QTextListFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextListFormatFromPointer(cthis)
	return gothis
}
func NewQTextListFormatFromPointer(cthis unsafe.Pointer) *QTextListFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextListFormat{bcthis0}
}

// /usr/include/qt/QtGui/qtextformat.h:714
// index:1
// void QTextListFormat(const class QTextFormat &)
func NewQTextListFormat_1(fmt unsafe.Pointer) *QTextListFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_VOID, cthis, fmt)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextListFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:683
// index:0
// inline
// bool isValid()
func (this *QTextListFormat) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextListFormat7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:697
// index:0
// inline
// void setStyle(enum QTextListFormat::Style)
func (this *QTextListFormat) SetStyle(style int) {
	// 0: (, style QTextListFormat::Style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormat8setStyleENS_5StyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:698
// index:0
// inline
// QTextListFormat::Style style()
func (this *QTextListFormat) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextListFormat5styleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:701
// index:0
// inline
// void setIndent(int)
func (this *QTextListFormat) SetIndent(indent int) {
	// 0: (, indent int), (&indent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormat9setIndentEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &indent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:702
// index:0
// inline
// int indent()
func (this *QTextListFormat) Indent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextListFormat6indentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:705
// index:0
// inline
// void setNumberPrefix(const class QString &)
func (this *QTextListFormat) SetNumberPrefix(numberPrefix unsafe.Pointer) {
	// 0: (, numberPrefix const QString &), (numberPrefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormat15setNumberPrefixERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), numberPrefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:706
// index:0
// inline
// QString numberPrefix()
func (this *QTextListFormat) NumberPrefix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextListFormat12numberPrefixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:709
// index:0
// inline
// void setNumberSuffix(const class QString &)
func (this *QTextListFormat) SetNumberSuffix(numberSuffix unsafe.Pointer) {
	// 0: (, numberSuffix const QString &), (numberSuffix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextListFormat15setNumberSuffixERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), numberSuffix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:710
// index:0
// inline
// QString numberSuffix()
func (this *QTextListFormat) NumberSuffix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextListFormat12numberSuffixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
