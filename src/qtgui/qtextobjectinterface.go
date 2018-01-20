//  header block begin
// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h
// #include <qabstracttextdocumentlayout.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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
type QTextObjectInterface struct {
	*qtrt.CObject
}

func (this *QTextObjectInterface) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:141
// index:0
// virtual
// void ~QTextObjectInterface()
func DeleteQTextObjectInterface(*QTextObjectInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextObjectInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:142
// index:0
// pure virtual
// QSizeF intrinsicSize(class QTextDocument *, int, const class QTextFormat &)
func (this *QTextObjectInterface) IntrinsicSize(doc unsafe.Pointer, posInDocument int, format unsafe.Pointer) {
	// 0: (, doc QTextDocument *, posInDocument int, format const QTextFormat &), (doc, &posInDocument, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextObjectInterface13intrinsicSizeEP13QTextDocumentiRK11QTextFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), doc, &posInDocument, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:143
// index:0
// pure virtual
// void drawObject(class QPainter *, const class QRectF &, class QTextDocument *, int, const class QTextFormat &)
func (this *QTextObjectInterface) DrawObject(painter unsafe.Pointer, rect unsafe.Pointer, doc unsafe.Pointer, posInDocument int, format unsafe.Pointer) {
	// 0: (, painter QPainter *, rect const QRectF &, doc QTextDocument *, posInDocument int, format const QTextFormat &), (painter, rect, doc, &posInDocument, format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextObjectInterface10drawObjectEP8QPainterRK6QRectFP13QTextDocumentiRK11QTextFormat", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect, doc, &posInDocument, format)
	gopp.ErrPrint(err, rv)
}

//  body block end
