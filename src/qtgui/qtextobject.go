//  header block begin
// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
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
type QTextObject struct {
	*qtcore.QObject
}

func (this *QTextObject) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qtextobject.h:62
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTextObject) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:65
// index:0
// void QTextObject(class QTextDocument *)
func NewQTextObject(doc unsafe.Pointer) *QTextObject {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextObjectC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, doc)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextObjectFromPointer(cthis)
	return gothis
}
func NewQTextObjectFromPointer(cthis unsafe.Pointer) *QTextObject {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QTextObject{bcthis0}
}

// /usr/include/qt/QtGui/qtextobject.h:81
// index:1
// void QTextObject(class QTextObjectPrivate &, class QTextDocument *)
func NewQTextObject_1(p unsafe.Pointer, doc unsafe.Pointer) *QTextObject {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextObjectC2ER18QTextObjectPrivateP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, p, doc)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextObjectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:66
// index:0
// virtual
// void ~QTextObject()
func DeleteQTextObject(*QTextObject) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextObjectD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:68
// index:0
// void setFormat(const class QTextFormat &)
func (this *QTextObject) SetFormat(format unsafe.Pointer) {
	// 0: (, format const QTextFormat &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextObject9setFormatERK11QTextFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:71
// index:0
// QTextFormat format()
func (this *QTextObject) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject6formatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:72
// index:0
// int formatIndex()
func (this *QTextObject) FormatIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject11formatIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:74
// index:0
// QTextDocument * document()
func (this *QTextObject) Document() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject8documentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:76
// index:0
// int objectIndex()
func (this *QTextObject) ObjectIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject11objectIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:78
// index:0
// QTextDocumentPrivate * docHandle()
func (this *QTextObject) DocHandle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject9docHandleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
