//  header block begin
// /usr/include/qt/QtGui/qtextdocumentfragment.h
// #include <qtextdocumentfragment.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QTextDocumentFragment struct {
	*qtrt.CObject
}

func (this *QTextDocumentFragment) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTextDocumentFragmentFromPointer(cthis unsafe.Pointer) *QTextDocumentFragment {
	return &QTextDocumentFragment{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:57
// index:0
// Public
// void QTextDocumentFragment()
func NewQTextDocumentFragment() *QTextDocumentFragment {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentFragmentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:58
// index:1
// Public
// void QTextDocumentFragment(const class QTextDocument *)
func NewQTextDocumentFragment_1(document unsafe.Pointer) *QTextDocumentFragment {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2EPK13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, document)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentFragmentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:59
// index:2
// Public
// void QTextDocumentFragment(const class QTextCursor &)
func NewQTextDocumentFragment_2(range_ *QTextCursor) *QTextDocumentFragment {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = range_.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2ERK11QTextCursor", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextDocumentFragmentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:62
// index:0
// Public
// void ~QTextDocumentFragment()
func DeleteQTextDocumentFragment(*QTextDocumentFragment) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:64
// index:0
// Public
// bool isEmpty()
func (this *QTextDocumentFragment) IsEmpty() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QTextDocumentFragment7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:66
// index:0
// Public
// QString toPlainText()
func (this *QTextDocumentFragment) ToPlainText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QTextDocumentFragment11toPlainTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:68
// index:0
// Public
// QString toHtml(const class QByteArray &)
func (this *QTextDocumentFragment) ToHtml(encoding *qtcore.QByteArray) interface{} {
	var convArg0 = encoding.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QTextDocumentFragment6toHtmlERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:71
// index:0
// Public static
// QTextDocumentFragment fromPlainText(const class QString &)
func (this *QTextDocumentFragment) FromPlainText(plainText *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragment13fromPlainTextERK7QString", ffiqt.FFI_TYPE_POINTER, plainText)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextDocumentFragment_FromPlainText(plainText *qtcore.QString) {
	var nilthis *QTextDocumentFragment
	nilthis.FromPlainText(plainText)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:73
// index:0
// Public static
// QTextDocumentFragment fromHtml(const class QString &)
func (this *QTextDocumentFragment) FromHtml(html *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragment8fromHtmlERK7QString", ffiqt.FFI_TYPE_POINTER, html)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextDocumentFragment_FromHtml(html *qtcore.QString) {
	var nilthis *QTextDocumentFragment
	nilthis.FromHtml(html)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:74
// index:1
// Public static
// QTextDocumentFragment fromHtml(const class QString &, const class QTextDocument *)
func (this *QTextDocumentFragment) FromHtml_1(html *qtcore.QString, resourceProvider unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragment8fromHtmlERK7QStringPK13QTextDocument", ffiqt.FFI_TYPE_POINTER, html, resourceProvider)
	gopp.ErrPrint(err, rv)
	return rv
}
func QTextDocumentFragment_FromHtml_1(html *qtcore.QString, resourceProvider unsafe.Pointer) {
	var nilthis *QTextDocumentFragment
	nilthis.FromHtml_1(html, resourceProvider)
}

//  body block end
