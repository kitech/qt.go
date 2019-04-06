//  header block begin

// +build !minimal

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
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// /usr/include/qt/QtWidgets/qundostack.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QUndoStack) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundostack.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoStack(QObject *)

/*
Constructs an empty undo stack with the parent parent. The stack will initially be in the clean state. If parent is a QUndoGroup object, the stack is automatically added to the group.

See also push().
*/
func (*QUndoStack) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QUndoStack {
	return NewQUndoStack(parent)
}
func NewQUndoStack(parent qtcore.QObject_ITF /*777 QObject **/) *QUndoStack {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStackC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoStackFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoStack")
	return gothis
}

// /usr/include/qt/QtWidgets/qundostack.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoStack(QObject *)

/*
Constructs an empty undo stack with the parent parent. The stack will initially be in the clean state. If parent is a QUndoGroup object, the stack is automatically added to the group.

See also push().
*/
func (*QUndoStack) NewForInheritp() *QUndoStack {
	return NewQUndoStackp()
}
func NewQUndoStackp() *QUndoStack {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStackC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoStackFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoStack")
	return gothis
}

// /usr/include/qt/QtWidgets/qundostack.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUndoStack()

/*

 */
func DeleteQUndoStack(this *QUndoStack) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStackD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qundostack.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the command stack by deleting all commands on it, and returns the stack to the clean state.

Commands are not undone or redone; the state of the edited object remains unchanged.

This function is usually used when the contents of the document are abandoned.

See also QUndoStack().
*/
func (this *QUndoStack) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void push(QUndoCommand *)

/*
Pushes cmd on the stack or merges it with the most recently executed command. In either case, executes cmd by calling its redo() function.

If cmd's id is not -1, and if the id is the same as that of the most recently executed command, QUndoStack will attempt to merge the two commands by calling QUndoCommand::mergeWith() on the most recently executed command. If QUndoCommand::mergeWith() returns true, cmd is deleted.

After calling QUndoCommand::redo() and, if applicable, QUndoCommand::mergeWith(), QUndoCommand::isObsolete() will be called for cmd or the merged command. If QUndoCommand::isObsolete() returns true, then cmd or the merged command will be deleted from the stack.

In all other cases cmd is simply pushed on the stack.

If commands were undone before cmd was pushed, the current command and all commands above it are deleted. Hence cmd always ends up being the top-most on the stack.

Once a command is pushed, the stack takes ownership of it. There are no getters to return the command, since modifying it after it has been executed will almost always lead to corruption of the document's state.

See also QUndoCommand::id() and QUndoCommand::mergeWith().
*/
func (this *QUndoStack) Push(cmd QUndoCommand_ITF /*777 QUndoCommand **/) {
	var convArg0 unsafe.Pointer
	if cmd != nil && cmd.QUndoCommand_PTR() != nil {
		convArg0 = cmd.QUndoCommand_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack4pushEP12QUndoCommand", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canUndo() const

/*
Returns true if there is a command available for undo; otherwise returns false.

This function returns false if the stack is empty, or if the bottom command on the stack has already been undone.

Synonymous with index() == 0.

Note: Getter function for property canUndo.

See also index() and canRedo().
*/
func (this *QUndoStack) CanUndo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack7canUndoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canRedo() const

/*
Returns true if there is a command available for redo; otherwise returns false.

This function returns false if the stack is empty or if the top command on the stack has already been redone.

Synonymous with index() == count().

Note: Getter function for property canRedo.

See also index() and canUndo().
*/
func (this *QUndoStack) CanRedo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack7canRedoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QString undoText() const

/*
Returns the text of the command which will be undone in the next call to undo().

Note: Getter function for property undoText.

See also QUndoCommand::actionText() and redoText().
*/
func (this *QUndoStack) UndoText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack8undoTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundostack.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QString redoText() const

/*
Returns the text of the command which will be redone in the next call to redo().

Note: Getter function for property redoText.

See also QUndoCommand::actionText() and undoText().
*/
func (this *QUndoStack) RedoText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack8redoTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundostack.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*
Returns the number of commands on the stack. Macro commands are counted as one command.

See also index(), setIndex(), and command().
*/
func (this *QUndoStack) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qundostack.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int index() const

/*
Returns the index of the current command. This is the command that will be executed on the next call to redo(). It is not always the top-most command on the stack, since a number of commands may have been undone.

See also setIndex(), undo(), redo(), and count().
*/
func (this *QUndoStack) Index() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack5indexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qundostack.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(int) const

/*
Returns the text of the command at index idx.

See also beginMacro().
*/
func (this *QUndoStack) Text(idx int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack4textEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), idx)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundostack.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createUndoAction(QObject *, const QString &) const

/*
Creates an undo QAction object with the given parent.

Triggering this action will cause a call to undo(). The text of this action is the text of the command which will be undone in the next call to undo(), prefixed by the specified prefix. If there is no command available for undo, this action will be disabled.

If prefix is empty, the default template "Undo %1" is used instead of prefix. Before Qt 4.8, the prefix "Undo" was used by default.

See also createRedoAction(), canUndo(), and QUndoCommand::text().
*/
func (this *QUndoStack) CreateUndoAction(parent qtcore.QObject_ITF /*777 QObject **/, prefix string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(prefix)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundostack.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createUndoAction(QObject *, const QString &) const

/*
Creates an undo QAction object with the given parent.

Triggering this action will cause a call to undo(). The text of this action is the text of the command which will be undone in the next call to undo(), prefixed by the specified prefix. If there is no command available for undo, this action will be disabled.

If prefix is empty, the default template "Undo %1" is used instead of prefix. Before Qt 4.8, the prefix "Undo" was used by default.

See also createRedoAction(), canUndo(), and QUndoCommand::text().
*/
func (this *QUndoStack) CreateUndoActionp(parent qtcore.QObject_ITF /*777 QObject **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundostack.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createRedoAction(QObject *, const QString &) const

/*
Creates an redo QAction object with the given parent.

Triggering this action will cause a call to redo(). The text of this action is the text of the command which will be redone in the next call to redo(), prefixed by the specified prefix. If there is no command available for redo, this action will be disabled.

If prefix is empty, the default template "Redo %1" is used instead of prefix. Before Qt 4.8, the prefix "Redo" was used by default.

See also createUndoAction(), canRedo(), and QUndoCommand::text().
*/
func (this *QUndoStack) CreateRedoAction(parent qtcore.QObject_ITF /*777 QObject **/, prefix string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(prefix)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundostack.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createRedoAction(QObject *, const QString &) const

/*
Creates an redo QAction object with the given parent.

Triggering this action will cause a call to redo(). The text of this action is the text of the command which will be redone in the next call to redo(), prefixed by the specified prefix. If there is no command available for redo, this action will be disabled.

If prefix is empty, the default template "Redo %1" is used instead of prefix. Before Qt 4.8, the prefix "Redo" was used by default.

See also createUndoAction(), canRedo(), and QUndoCommand::text().
*/
func (this *QUndoStack) CreateRedoActionp(parent qtcore.QObject_ITF /*777 QObject **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_11368() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
