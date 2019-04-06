package qtgui

// /usr/include/qt/QtGui/qtextdocument.h
// #include <qtextdocument.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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

/*

 */
type QAbstractUndoItem struct {
	*qtrt.CObject
}
type QAbstractUndoItem_ITF interface {
	QAbstractUndoItem_PTR() *QAbstractUndoItem
}

func (ptr *QAbstractUndoItem) QAbstractUndoItem_PTR() *QAbstractUndoItem { return ptr }

func (this *QAbstractUndoItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAbstractUndoItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAbstractUndoItemFromPointer(cthis unsafe.Pointer) *QAbstractUndoItem {
	return &QAbstractUndoItem{&qtrt.CObject{cthis}}
}
func (*QAbstractUndoItem) NewFromPointer(cthis unsafe.Pointer) *QAbstractUndoItem {
	return NewQAbstractUndoItemFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextdocument.h:86
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractUndoItem()

/*

 */
func DeleteQAbstractUndoItem(this *QAbstractUndoItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractUndoItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextdocument.h:87
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void undo()

/*
Undoes the last editing operation on the document if undo is available. The provided cursor is positioned at the end of the location where the edition operation was undone.

See the Qt Undo Framework documentation for details.

This function was introduced in  Qt 4.2.

See also undoAvailable() and isUndoRedoEnabled().
*/
func (this *QAbstractUndoItem) Undo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractUndoItem4undoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:88
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void redo()

/*
Redoes the last editing operation on the document if redo is available.

The provided cursor is positioned at the end of the location where the edition operation was redone.

This function was introduced in  Qt 4.2.
*/
func (this *QAbstractUndoItem) Redo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAbstractUndoItem4redoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10791() {
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
