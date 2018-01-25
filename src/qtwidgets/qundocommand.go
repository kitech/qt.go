package qtwidgets

// /usr/include/qt/QtWidgets/qundostack.h
// #include <qundostack.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
	*qtrt.CObject
}

func (this *QUndoCommand) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QUndoCommand) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQUndoCommandFromPointer(cthis unsafe.Pointer) *QUndoCommand {
	return &QUndoCommand{&qtrt.CObject{cthis}}
}
func (*QUndoCommand) NewFromPointer(cthis unsafe.Pointer) *QUndoCommand {
	return NewQUndoCommandFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qundostack.h:60
// index:0
// Public
// void QUndoCommand(class QUndoCommand *)
func NewQUndoCommand(parent *QUndoCommand /*444 QUndoCommand **/) *QUndoCommand {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommandC2EPS_", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoCommandFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qundostack.h:61
// index:1
// Public
// void QUndoCommand(const class QString &, class QUndoCommand *)
func NewQUndoCommand_1(text *qtcore.QString, parent *QUndoCommand /*444 QUndoCommand **/) *QUndoCommand {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommandC2ERK7QStringPS_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoCommandFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qundostack.h:62
// index:0
// Public virtual
// void ~QUndoCommand()
func DeleteQUndoCommand(*QUndoCommand) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommandD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:64
// index:0
// Public virtual
// void undo()
func (this *QUndoCommand) Undo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommand4undoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:65
// index:0
// Public virtual
// void redo()
func (this *QUndoCommand) Redo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommand4redoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:67
// index:0
// Public
// QString text()
func (this *QUndoCommand) Text() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand4textEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qundostack.h:68
// index:0
// Public
// QString actionText()
func (this *QUndoCommand) ActionText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand10actionTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qundostack.h:69
// index:0
// Public
// void setText(const class QString &)
func (this *QUndoCommand) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommand7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:71
// index:0
// Public
// bool isObsolete()
func (this *QUndoCommand) IsObsolete() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand10isObsoleteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:72
// index:0
// Public
// void setObsolete(_Bool)
func (this *QUndoCommand) SetObsolete(obsolete bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommand11setObsoleteEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), obsolete)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:74
// index:0
// Public virtual
// int id()
func (this *QUndoCommand) Id() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand2idEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qundostack.h:75
// index:0
// Public virtual
// bool mergeWith(const class QUndoCommand *)
func (this *QUndoCommand) MergeWith(other *QUndoCommand /*444 const QUndoCommand **/) bool {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QUndoCommand9mergeWithEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:77
// index:0
// Public
// int childCount()
func (this *QUndoCommand) ChildCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand10childCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qundostack.h:78
// index:0
// Public
// const QUndoCommand * child(int)
func (this *QUndoCommand) Child(index int) *QUndoCommand /*444 const QUndoCommand **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QUndoCommand5childEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQUndoCommandFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
