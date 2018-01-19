//  header block begin
// /usr/include/qt/QtWidgets/qundostack.h
// #include <qundostack.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QUndoCommand struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qundostack.h:60
// index:0
// void QUndoCommand(class QUndoCommand *)
func NewQUndoCommand(parent unsafe.Pointer) *QUndoCommand {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommandC2EPS_", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QUndoCommand{cthis}
}

// /usr/include/qt/QtWidgets/qundostack.h:61
// index:1
// void QUndoCommand(const class QString &, class QUndoCommand *)
func NewQUndoCommand_1(text unsafe.Pointer, parent unsafe.Pointer) *QUndoCommand {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommandC2ERK7QStringPS_", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	return &QUndoCommand{cthis}
}

// /usr/include/qt/QtWidgets/qundostack.h:62
// index:0
// virtual
// void ~QUndoCommand()
func DeleteQUndoCommand(*QUndoCommand) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommandD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:64
// index:0
// virtual
// void undo()
func (this *QUndoCommand) Undo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommand4undoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:65
// index:0
// virtual
// void redo()
func (this *QUndoCommand) Redo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommand4redoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:67
// index:0
// QString text()
func (this *QUndoCommand) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:68
// index:0
// QString actionText()
func (this *QUndoCommand) ActionText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand10actionTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:69
// index:0
// void setText(const class QString &)
func (this *QUndoCommand) SetText(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommand7setTextERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:71
// index:0
// bool isObsolete()
func (this *QUndoCommand) IsObsolete() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand10isObsoleteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:72
// index:0
// void setObsolete(_Bool)
func (this *QUndoCommand) SetObsolete(obsolete bool) {
	// 0: (, bool obsolete), (&obsolete)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommand11setObsoleteEb", ffiqt.FFI_TYPE_VOID, this.cthis, &obsolete)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:74
// index:0
// virtual
// int id()
func (this *QUndoCommand) Id() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand2idEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:75
// index:0
// virtual
// bool mergeWith(const class QUndoCommand *)
func (this *QUndoCommand) MergeWith(other unsafe.Pointer) {
	// 0: (, const QUndoCommand * other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommand9mergeWithEPKS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:77
// index:0
// int childCount()
func (this *QUndoCommand) ChildCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand10childCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:78
// index:0
// const QUndoCommand * child(int)
func (this *QUndoCommand) Child(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand5childEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

//  body block end
