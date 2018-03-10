package qtquick

// /usr/include/qt/QtQuick/qquicktextdocument.h
// #include <qquicktextdocument.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

/*

 */
type QQuickTextDocument struct {
	*qtcore.QObject
}
type QQuickTextDocument_ITF interface {
	qtcore.QObject_ITF
	QQuickTextDocument_PTR() *QQuickTextDocument
}

func (ptr *QQuickTextDocument) QQuickTextDocument_PTR() *QQuickTextDocument { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQuickTextDocument) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QQuickTextDocument10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquicktextdocument.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickTextDocument(QQuickItem *)

/*
Constructs a QQuickTextDocument object with parent as the parent object.
*/
func NewQQuickTextDocument(parent QQuickItem_ITF /*777 QQuickItem **/) *QQuickTextDocument {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QQuickItem_PTR() != nil {
		convArg0 = parent.QQuickItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QQuickTextDocumentC2EP10QQuickItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickTextDocumentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickTextDocument")
	return gothis
}

// /usr/include/qt/QtQuick/qquicktextdocument.h:55
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * textDocument() const

/*
Returns a pointer to the QTextDocument object.
*/
func (this *QQuickTextDocument) TextDocument() *qtgui.QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QQuickTextDocument12textDocumentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

func DeleteQQuickTextDocument(this *QQuickTextDocument) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QQuickTextDocumentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
