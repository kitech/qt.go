//  header block begin
// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h
// #include <qabstracttextdocumentlayout.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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
func NewQTextObjectInterfaceFromPointer(cthis unsafe.Pointer) *QTextObjectInterface {
	return &QTextObjectInterface{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:141
// index:0
// Public virtual
// void ~QTextObjectInterface()
func DeleteQTextObjectInterface(*QTextObjectInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextObjectInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:142
// index:0
// Public pure virtual
// QSizeF intrinsicSize(class QTextDocument *, int, const class QTextFormat &)
func (this *QTextObjectInterface) IntrinsicSize(doc unsafe.Pointer, posInDocument int, format *QTextFormat) interface{} {
	var convArg2 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextObjectInterface13intrinsicSizeEP13QTextDocumentiRK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), doc, &posInDocument, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:143
// index:0
// Public pure virtual
// void drawObject(class QPainter *, const class QRectF &, class QTextDocument *, int, const class QTextFormat &)
func (this *QTextObjectInterface) DrawObject(painter unsafe.Pointer, rect *qtcore.QRectF, doc unsafe.Pointer, posInDocument int, format *QTextFormat) {
	var convArg1 = rect.GetCthis()
	var convArg4 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextObjectInterface10drawObjectEP8QPainterRK6QRectFP13QTextDocumentiRK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, convArg1, doc, &posInDocument, convArg4)
	gopp.ErrPrint(err, rv)
}

//  body block end
