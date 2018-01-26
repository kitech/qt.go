package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QTextBlockGroup struct {
	*QTextObject
}

func (this *QTextBlockGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextObject.GetCthis()
	}
}
func (this *QTextBlockGroup) SetCthis(cthis unsafe.Pointer) {
	this.QTextObject = NewQTextObjectFromPointer(cthis)
}
func NewQTextBlockGroupFromPointer(cthis unsafe.Pointer) *QTextBlockGroup {
	bcthis0 := NewQTextObjectFromPointer(cthis)
	return &QTextBlockGroup{bcthis0}
}
func (*QTextBlockGroup) NewFromPointer(cthis unsafe.Pointer) *QTextBlockGroup {
	return NewQTextBlockGroupFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextobject.h:92
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTextBlockGroup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextBlockGroup10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:95
// index:0
// Protected
// void QTextBlockGroup(class QTextDocument *)
func NewQTextBlockGroup(doc *QTextDocument /*777 QTextDocument **/) *QTextBlockGroup {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = doc.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroupC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBlockGroupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:96
// index:0
// Protected virtual
// void ~QTextBlockGroup()
func DeleteQTextBlockGroup(*QTextBlockGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:98
// index:0
// Protected virtual
// void blockInserted(const class QTextBlock &)
func (this *QTextBlockGroup) BlockInserted(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroup13blockInsertedERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:99
// index:0
// Protected virtual
// void blockRemoved(const class QTextBlock &)
func (this *QTextBlockGroup) BlockRemoved(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroup12blockRemovedERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:100
// index:0
// Protected virtual
// void blockFormatChanged(const class QTextBlock &)
func (this *QTextBlockGroup) BlockFormatChanged(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroup18blockFormatChangedERK10QTextBlock", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
