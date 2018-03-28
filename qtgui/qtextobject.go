package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// void setFormat(const QTextFormat &)
func (this *QTextObject) InheritSetFormat(f func(format *QTextFormat) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setFormat", f)
}

/*

 */
type QTextObject struct {
	*qtcore.QObject
}
type QTextObject_ITF interface {
	qtcore.QObject_ITF
	QTextObject_PTR() *QTextObject
}

func (ptr *QTextObject) QTextObject_PTR() *QTextObject { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTextObject) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextObject10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:65
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QTextObject(QTextDocument *)

/*
Creates a new QTextObject for the given document.

Warning: This function should never be called directly, but only from QTextDocument::createObject().
*/
func NewQTextObject(doc QTextDocument_ITF /*777 QTextDocument **/) *QTextObject {
	var convArg0 unsafe.Pointer
	if doc != nil && doc.QTextDocument_PTR() != nil {
		convArg0 = doc.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextObjectC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextObject")
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:66
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ~QTextObject()

/*

 */
func DeleteQTextObject(this *QTextObject) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextObjectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextobject.h:68
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setFormat(const QTextFormat &)

/*
Sets the text object's format.

See also format().
*/
func (this *QTextObject) SetFormat(format QTextFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextFormat_PTR() != nil {
		convArg0 = format.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextObject9setFormatERK11QTextFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:71
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextFormat format() const

/*
Returns the text object's format.

See also setFormat() and document().
*/
func (this *QTextObject) Format() *QTextFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextObject6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:72
// index:0
// Public Visibility=Default Availability=Available
// [4] int formatIndex() const

/*
Returns the index of the object's format in the document's internal list of formats.

See also QTextDocument::allFormats().
*/
func (this *QTextObject) FormatIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextObject11formatIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDocument * document() const

/*
Returns the document this object belongs to.

See also format().
*/
func (this *QTextObject) Document() *QTextDocument /*777 QTextDocument **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextObject8documentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDocumentFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int objectIndex() const

/*
Returns the object index of this object. This can be used together with QTextFormat::setObjectIndex().
*/
func (this *QTextObject) ObjectIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextObject11objectIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
}

//  keep block end
