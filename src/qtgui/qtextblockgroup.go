//  header block begin
// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
	return this.QTextObject.GetCthis()
}

// /usr/include/qt/QtGui/qtextobject.h:92
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTextBlockGroup) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextBlockGroup10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:95
// index:0
// void QTextBlockGroup(class QTextDocument *)
func NewQTextBlockGroup(doc unsafe.Pointer) *QTextBlockGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroupC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, doc)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBlockGroupFromPointer(cthis)
	return gothis
}
func NewQTextBlockGroupFromPointer(cthis unsafe.Pointer) *QTextBlockGroup {
	bcthis0 := NewQTextObjectFromPointer(cthis)
	return &QTextBlockGroup{bcthis0}
}

// /usr/include/qt/QtGui/qtextobject.h:105
// index:1
// void QTextBlockGroup(class QTextBlockGroupPrivate &, class QTextDocument *)
func NewQTextBlockGroup_1(p unsafe.Pointer, doc unsafe.Pointer) *QTextBlockGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroupC2ER22QTextBlockGroupPrivateP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, p, doc)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextBlockGroupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:96
// index:0
// virtual
// void ~QTextBlockGroup()
func DeleteQTextBlockGroup(*QTextBlockGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:98
// index:0
// virtual
// void blockInserted(const class QTextBlock &)
func (this *QTextBlockGroup) BlockInserted(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroup13blockInsertedERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:99
// index:0
// virtual
// void blockRemoved(const class QTextBlock &)
func (this *QTextBlockGroup) BlockRemoved(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroup12blockRemovedERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:100
// index:0
// virtual
// void blockFormatChanged(const class QTextBlock &)
func (this *QTextBlockGroup) BlockFormatChanged(block unsafe.Pointer) {
	// 0: (, block const QTextBlock &), (block)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextBlockGroup18blockFormatChangedERK10QTextBlock", ffiqt.FFI_TYPE_VOID, this.GetCthis(), block)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:102
// index:0
// QList<QTextBlock> blockList()
func (this *QTextBlockGroup) BlockList() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextBlockGroup9blockListEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
