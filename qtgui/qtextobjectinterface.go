package qtgui

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h
// #include <qabstracttextdocumentlayout.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QTextObjectInterface struct {
	*qtrt.CObject
}

func (this *QTextObjectInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextObjectInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextObjectInterfaceFromPointer(cthis unsafe.Pointer) *QTextObjectInterface {
	return &QTextObjectInterface{&qtrt.CObject{cthis}}
}
func (*QTextObjectInterface) NewFromPointer(cthis unsafe.Pointer) *QTextObjectInterface {
	return NewQTextObjectInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:141
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextObjectInterface()
func DeleteQTextObjectInterface(this *QTextObjectInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextObjectInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:142
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QSizeF intrinsicSize(QTextDocument *, int, const QTextFormat &)
func (this *QTextObjectInterface) IntrinsicSize(doc *QTextDocument /*777 QTextDocument **/, posInDocument int, format *QTextFormat) *qtcore.QSizeF /*123*/ {
	var convArg0 = doc.GetCthis()
	var convArg2 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextObjectInterface13intrinsicSizeEP13QTextDocumentiRK11QTextFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, posInDocument, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qabstracttextdocumentlayout.h:143
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void drawObject(QPainter *, const QRectF &, QTextDocument *, int, const QTextFormat &)
func (this *QTextObjectInterface) DrawObject(painter *QPainter /*777 QPainter **/, rect *qtcore.QRectF, doc *QTextDocument /*777 QTextDocument **/, posInDocument int, format *QTextFormat) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	var convArg2 = doc.GetCthis()
	var convArg4 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextObjectInterface10drawObjectEP8QPainterRK6QRectFP13QTextDocumentiRK11QTextFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, posInDocument, convArg4)
	qtrt.ErrPrint(err, rv)
}

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
