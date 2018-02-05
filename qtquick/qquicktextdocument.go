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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQuickTextDocument) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QQuickTextDocument10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquicktextdocument.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickTextDocument(QQuickItem *)
func NewQQuickTextDocument(parent *QQuickItem /*777 QQuickItem **/) *QQuickTextDocument {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QQuickTextDocumentC2EP10QQuickItem", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickTextDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQuick/qquicktextdocument.h:55
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * textDocument()
func (this *QQuickTextDocument) TextDocument() *qtgui.QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QQuickTextDocument12textDocumentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

func DeleteQQuickTextDocument(this *QQuickTextDocument) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QQuickTextDocumentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
