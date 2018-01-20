//  header block begin
// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QTextImageFormat struct {
	*QTextCharFormat
}

func (this *QTextImageFormat) GetCthis() unsafe.Pointer {
	return this.QTextCharFormat.GetCthis()
}

// /usr/include/qt/QtGui/qtextformat.h:735
// index:0
// void QTextImageFormat()
func NewQTextImageFormat() *QTextImageFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextImageFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextImageFormatFromPointer(cthis)
	return gothis
}
func NewQTextImageFormatFromPointer(cthis unsafe.Pointer) *QTextImageFormat {
	bcthis0 := NewQTextCharFormatFromPointer(cthis)
	return &QTextImageFormat{bcthis0}
}

// /usr/include/qt/QtGui/qtextformat.h:752
// index:1
// void QTextImageFormat(const class QTextFormat &)
func NewQTextImageFormat_1(format unsafe.Pointer) *QTextImageFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextImageFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_VOID, cthis, format)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextImageFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:737
// index:0
// inline
// bool isValid()
func (this *QTextImageFormat) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextImageFormat7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:739
// index:0
// inline
// void setName(const class QString &)
func (this *QTextImageFormat) SetName(name unsafe.Pointer) {
	// 0: (, name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextImageFormat7setNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:740
// index:0
// inline
// QString name()
func (this *QTextImageFormat) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextImageFormat4nameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:743
// index:0
// inline
// void setWidth(qreal)
func (this *QTextImageFormat) SetWidth(width float64) {
	// 0: (, width qreal), (&width)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextImageFormat8setWidthEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:744
// index:0
// inline
// qreal width()
func (this *QTextImageFormat) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextImageFormat5widthEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:747
// index:0
// inline
// void setHeight(qreal)
func (this *QTextImageFormat) SetHeight(height float64) {
	// 0: (, height qreal), (&height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QTextImageFormat9setHeightEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:748
// index:0
// inline
// qreal height()
func (this *QTextImageFormat) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QTextImageFormat6heightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
