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
// extern C begin: 8
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
type QTextBlockGroup_ITF interface {
	QTextObject_ITF
	QTextBlockGroup_PTR() *QTextBlockGroup
}

func (ptr *QTextBlockGroup) QTextBlockGroup_PTR() *QTextBlockGroup { return ptr }

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
// [8] const QMetaObject * metaObject() const
func (this *QTextBlockGroup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextBlockGroup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:95
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QTextBlockGroup(QTextDocument *)
func NewQTextBlockGroup(doc QTextDocument_ITF /*777 QTextDocument **/) *QTextBlockGroup {
	var convArg0 unsafe.Pointer
	if doc != nil && doc.QTextDocument_PTR() != nil {
		convArg0 = doc.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextBlockGroupC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextBlockGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextBlockGroup")
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ~QTextBlockGroup()
func DeleteQTextBlockGroup(this *QTextBlockGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextBlockGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextobject.h:98
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void blockInserted(const QTextBlock &)
func (this *QTextBlockGroup) BlockInserted(block QTextBlock_ITF) {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextBlockGroup13blockInsertedERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:99
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void blockRemoved(const QTextBlock &)
func (this *QTextBlockGroup) BlockRemoved(block QTextBlock_ITF) {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextBlockGroup12blockRemovedERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:100
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void blockFormatChanged(const QTextBlock &)
func (this *QTextBlockGroup) BlockFormatChanged(block QTextBlock_ITF) {
	var convArg0 unsafe.Pointer
	if block != nil && block.QTextBlock_PTR() != nil {
		convArg0 = block.QTextBlock_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextBlockGroup18blockFormatChangedERK10QTextBlock", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
