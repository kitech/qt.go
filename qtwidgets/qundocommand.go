package qtwidgets

// /usr/include/qt/QtWidgets/qundostack.h
// #include <qundostack.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QUndoCommand struct {
	*qtrt.CObject
}
type QUndoCommand_ITF interface {
	QUndoCommand_PTR() *QUndoCommand
}

func (ptr *QUndoCommand) QUndoCommand_PTR() *QUndoCommand { return ptr }

func (this *QUndoCommand) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QUndoCommand) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQUndoCommandFromPointer(cthis unsafe.Pointer) *QUndoCommand {
	return &QUndoCommand{&qtrt.CObject{cthis}}
}
func (*QUndoCommand) NewFromPointer(cthis unsafe.Pointer) *QUndoCommand {
	return NewQUndoCommandFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qundostack.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoCommand(QUndoCommand *)
func NewQUndoCommand(parent QUndoCommand_ITF /*777 QUndoCommand **/) *QUndoCommand {
	var convArg0 = parent.QUndoCommand_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QUndoCommandC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoCommandFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUndoCommand)
	return gothis
}

// /usr/include/qt/QtWidgets/qundostack.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUndoCommand(const QString &, QUndoCommand *)
func NewQUndoCommand_1(text string, parent QUndoCommand_ITF /*777 QUndoCommand **/) *QUndoCommand {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.QUndoCommand_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QUndoCommandC2ERK7QStringPS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoCommandFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUndoCommand)
	return gothis
}

// /usr/include/qt/QtWidgets/qundostack.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUndoCommand()
func DeleteQUndoCommand(this *QUndoCommand) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QUndoCommandD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qundostack.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void undo()
func (this *QUndoCommand) Undo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QUndoCommand4undoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void redo()
func (this *QUndoCommand) Redo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QUndoCommand4redoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QUndoCommand) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QUndoCommand4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundostack.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QString actionText()
func (this *QUndoCommand) ActionText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QUndoCommand10actionTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundostack.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QUndoCommand) SetText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QUndoCommand7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:71
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isObsolete()
func (this *QUndoCommand) IsObsolete() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QUndoCommand10isObsoleteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObsolete(_Bool)
func (this *QUndoCommand) SetObsolete(obsolete bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QUndoCommand11setObsoleteEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), obsolete)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int id()
func (this *QUndoCommand) Id() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QUndoCommand2idEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qundostack.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool mergeWith(const QUndoCommand *)
func (this *QUndoCommand) MergeWith(other QUndoCommand_ITF /*777 const QUndoCommand **/) bool {
	var convArg0 = other.QUndoCommand_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QUndoCommand9mergeWithEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int childCount()
func (this *QUndoCommand) ChildCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QUndoCommand10childCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qundostack.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] const QUndoCommand * child(int)
func (this *QUndoCommand) Child(index int) *QUndoCommand /*777 const QUndoCommand **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QUndoCommand5childEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQUndoCommandFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
