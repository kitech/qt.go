//  header block begin
// /usr/include/qt/QtGui/qtextdocument.h
// #include <qtextdocument.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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
type QAbstractUndoItem struct {
	*qtrt.CObject
}

func (this *QAbstractUndoItem) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qtextdocument.h:86
// index:0
// pure virtual
// void ~QAbstractUndoItem()
func DeleteQAbstractUndoItem(*QAbstractUndoItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractUndoItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:87
// index:0
// pure virtual
// void undo()
func (this *QAbstractUndoItem) Undo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractUndoItem4undoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextdocument.h:88
// index:0
// pure virtual
// void redo()
func (this *QAbstractUndoItem) Redo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractUndoItem4redoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
