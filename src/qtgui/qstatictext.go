//  header block begin
// /usr/include/qt/QtGui/qstatictext.h
// #include <qstatictext.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 55
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
type QStaticText struct {
	*qtrt.CObject
}

func (this *QStaticText) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qstatictext.h:64
// index:0
// void QStaticText()
func NewQStaticText() *QStaticText {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticTextC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStaticTextFromPointer(cthis)
	return gothis
}
func NewQStaticTextFromPointer(cthis unsafe.Pointer) *QStaticText {
	return &QStaticText{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qstatictext.h:65
// index:1
// void QStaticText(const class QString &)
func NewQStaticText_1(text unsafe.Pointer) *QStaticText {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticTextC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, text)
	gopp.ErrPrint(err, rv)
	gothis := NewQStaticTextFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qstatictext.h:71
// index:0
// void ~QStaticText()
func DeleteQStaticText(*QStaticText) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticTextD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:73
// index:0
// inline
// void swap(class QStaticText &)
func (this *QStaticText) Swap(other unsafe.Pointer) {
	// 0: (, other QStaticText &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:75
// index:0
// void setText(const class QString &)
func (this *QStaticText) SetText(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:76
// index:0
// QString text()
func (this *QStaticText) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText4textEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:78
// index:0
// void setTextFormat(Qt::TextFormat)
func (this *QStaticText) SetTextFormat(textFormat int) {
	// 0: (, textFormat Qt::TextFormat), (&textFormat)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText13setTextFormatEN2Qt10TextFormatE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &textFormat)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:79
// index:0
// Qt::TextFormat textFormat()
func (this *QStaticText) TextFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText10textFormatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:81
// index:0
// void setTextWidth(qreal)
func (this *QStaticText) SetTextWidth(textWidth float64) {
	// 0: (, textWidth qreal), (&textWidth)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText12setTextWidthEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &textWidth)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:82
// index:0
// qreal textWidth()
func (this *QStaticText) TextWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText9textWidthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:84
// index:0
// void setTextOption(const class QTextOption &)
func (this *QStaticText) SetTextOption(textOption unsafe.Pointer) {
	// 0: (, textOption const QTextOption &), (textOption)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText13setTextOptionERK11QTextOption", ffiqt.FFI_TYPE_VOID, this.GetCthis(), textOption)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:85
// index:0
// QTextOption textOption()
func (this *QStaticText) TextOption() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText10textOptionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:87
// index:0
// QSizeF size()
func (this *QStaticText) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:89
// index:0
// void prepare(const class QTransform &, const class QFont &)
func (this *QStaticText) Prepare(matrix unsafe.Pointer, font unsafe.Pointer) {
	// 0: (, matrix const QTransform &, font const QFont &), (matrix, font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText7prepareERK10QTransformRK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix, font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:91
// index:0
// void setPerformanceHint(enum QStaticText::PerformanceHint)
func (this *QStaticText) SetPerformanceHint(performanceHint int) {
	// 0: (, performanceHint QStaticText::PerformanceHint), (&performanceHint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStaticText18setPerformanceHintENS_15PerformanceHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &performanceHint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:92
// index:0
// QStaticText::PerformanceHint performanceHint()
func (this *QStaticText) PerformanceHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStaticText15performanceHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
