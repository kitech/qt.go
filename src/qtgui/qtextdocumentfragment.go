//  header block begin
// /usr/include/qt/QtGui/qtextdocumentfragment.h
// #include <qtextdocumentfragment.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:57
// index:0
// void QTextDocumentFragment()
func NewQTextDocumentFragment() *QTextDocumentFragment {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextDocumentFragment{cthis}
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:58
// index:1
// void QTextDocumentFragment(const class QTextDocument *)
func NewQTextDocumentFragment_1(document unsafe.Pointer) *QTextDocumentFragment {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2EPK13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, document)
	gopp.ErrPrint(err, rv)
	return &QTextDocumentFragment{cthis}
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:59
// index:2
// void QTextDocumentFragment(const class QTextCursor &)
func NewQTextDocumentFragment_2(range_ unsafe.Pointer) *QTextDocumentFragment {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentC2ERK11QTextCursor", ffiqt.FFI_TYPE_VOID, cthis, range_)
	gopp.ErrPrint(err, rv)
	return &QTextDocumentFragment{cthis}
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:62
// index:0
// void ~QTextDocumentFragment()
func DeleteQTextDocumentFragment(*QTextDocumentFragment) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragmentD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:64
// index:0
// bool isEmpty()
func (this *QTextDocumentFragment) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QTextDocumentFragment7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:66
// index:0
// QString toPlainText()
func (this *QTextDocumentFragment) ToPlainText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QTextDocumentFragment11toPlainTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:68
// index:0
// QString toHtml(const class QByteArray &)
func (this *QTextDocumentFragment) ToHtml(encoding unsafe.Pointer) {
	// 0: (, const QByteArray & encoding), (encoding)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QTextDocumentFragment6toHtmlERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, encoding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:71
// index:0
// static
// QTextDocumentFragment fromPlainText(const class QString &)
func (this *QTextDocumentFragment) FromPlainText(plainText unsafe.Pointer) {
	// 0: (const QString & plainText), (plainText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragment13fromPlainTextERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextDocumentFragment_FromPlainText(plainText unsafe.Pointer) {
	// 0: (const QString & plainText), (plainText)
	var nilthis *QTextDocumentFragment
	nilthis.FromPlainText(plainText)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:73
// index:0
// static
// QTextDocumentFragment fromHtml(const class QString &)
func (this *QTextDocumentFragment) FromHtml(html unsafe.Pointer) {
	// 0: (const QString & html), (html)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragment8fromHtmlERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextDocumentFragment_FromHtml(html unsafe.Pointer) {
	// 0: (const QString & html), (html)
	var nilthis *QTextDocumentFragment
	nilthis.FromHtml(html)
}

// /usr/include/qt/QtGui/qtextdocumentfragment.h:74
// index:1
// static
// QTextDocumentFragment fromHtml(const class QString &, const class QTextDocument *)
func (this *QTextDocumentFragment) FromHtml_1(html unsafe.Pointer, resourceProvider unsafe.Pointer) {
	// 1: (const QString & html, const QTextDocument * resourceProvider), (html, resourceProvider)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QTextDocumentFragment8fromHtmlERK7QStringPK13QTextDocument", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QTextDocumentFragment_FromHtml_1(html unsafe.Pointer, resourceProvider unsafe.Pointer) {
	// 1: (const QString & html, const QTextDocument * resourceProvider), (html, resourceProvider)
	var nilthis *QTextDocumentFragment
	nilthis.FromHtml_1(html, resourceProvider)
}

//  body block end
