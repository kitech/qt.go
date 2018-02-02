package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
// void setFormat(const class QTextFormat &)
func (this *QTextObject) InheritSetFormat(f func(format *QTextFormat)) {
	ffiqt.SetAllInheritCallback(this, "setFormat", f)
}

type QTextObject struct {
	*qtcore.QObject
}

func (this *QTextObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QTextObject) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQTextObjectFromPointer(cthis unsafe.Pointer) *QTextObject {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QTextObject{bcthis0}
}
func (*QTextObject) NewFromPointer(cthis unsafe.Pointer) *QTextObject {
	return NewQTextObjectFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextobject.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTextObject) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:65
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QTextObject(QTextDocument *)
func NewQTextObject(doc *QTextDocument /*777 QTextDocument **/) *QTextObject {
	var convArg0 = doc.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextObjectC2EP13QTextDocument", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:66
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ~QTextObject()
func DeleteQTextObject(this *QTextObject) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextObjectD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextobject.h:68
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setFormat(const QTextFormat &)
func (this *QTextObject) SetFormat(format *QTextFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextObject9setFormatERK11QTextFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:71
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextFormat format()
func (this *QTextObject) Format() *QTextFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:72
// index:0
// Public Visibility=Default Availability=Available
// [4] int formatIndex()
func (this *QTextObject) FormatIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject11formatIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document()
func (this *QTextObject) Document() *QTextDocument /*777 QTextDocument **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject8documentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int objectIndex()
func (this *QTextObject) ObjectIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextObject11objectIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
