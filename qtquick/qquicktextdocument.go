package qtquick

// /usr/include/qt/QtQuick/qquicktextdocument.h
// #include <qquicktextdocument.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQuickTextDocument struct {
	*qtcore.QObject
}

func (this *QQuickTextDocument) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickTextDocument) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQuickTextDocumentFromPointer(cthis unsafe.Pointer) *QQuickTextDocument {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQuickTextDocument{bcthis0}
}
func (*QQuickTextDocument) NewFromPointer(cthis unsafe.Pointer) *QQuickTextDocument {
	return NewQQuickTextDocumentFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquicktextdocument.h:51
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QQuickTextDocument) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QQuickTextDocument10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquicktextdocument.h:54
// index:0
// Public
// void QQuickTextDocument(class QQuickItem *)
func NewQQuickTextDocument(parent *QQuickItem /*777 QQuickItem **/) *QQuickTextDocument {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QQuickTextDocumentC2EP10QQuickItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickTextDocumentFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquicktextdocument.h:55
// index:0
// Public
// QTextDocument * textDocument()
func (this *QQuickTextDocument) TextDocument() *qtgui.QTextDocument /*777 QTextDocument **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QQuickTextDocument12textDocumentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
