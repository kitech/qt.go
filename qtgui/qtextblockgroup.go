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
import "gopp"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
// void blockInserted(const class QTextBlock &)
func (this *QTextBlockGroup) InheritBlockInserted(f func(block *QTextBlock) /*void*/) {
	qtrt.SetAllInheritCallback(this, "blockInserted", f)
}

// void blockRemoved(const class QTextBlock &)
func (this *QTextBlockGroup) InheritBlockRemoved(f func(block *QTextBlock) /*void*/) {
	qtrt.SetAllInheritCallback(this, "blockRemoved", f)
}

// void blockFormatChanged(const class QTextBlock &)
func (this *QTextBlockGroup) InheritBlockFormatChanged(f func(block *QTextBlock) /*void*/) {
	qtrt.SetAllInheritCallback(this, "blockFormatChanged", f)
}

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTextBlockGroup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextBlockGroup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:95
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QTextBlockGroup(QTextDocument *)
func NewQTextBlockGroup(doc *QTextDocument /*777 QTextDocument **/) *QTextBlockGroup {
	var convArg0 = doc.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextBlockGroupC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBlockGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ~QTextBlockGroup()
func DeleteQTextBlockGroup(this *QTextBlockGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextBlockGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextobject.h:98
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void blockInserted(const QTextBlock &)
func (this *QTextBlockGroup) BlockInserted(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextBlockGroup13blockInsertedERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:99
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void blockRemoved(const QTextBlock &)
func (this *QTextBlockGroup) BlockRemoved(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextBlockGroup12blockRemovedERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:100
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void blockFormatChanged(const QTextBlock &)
func (this *QTextBlockGroup) BlockFormatChanged(block *QTextBlock) {
	var convArg0 = block.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextBlockGroup18blockFormatChangedERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
