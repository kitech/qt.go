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

/*

 */
type QUndoStack struct {
	*qtcore.QObject
}
type QUndoStack_ITF interface {
	qtcore.QObject_ITF
	QUndoStack_PTR() *QUndoStack
}

func (ptr *QUndoStack) QUndoStack_PTR() *QUndoStack { return ptr }

func (this *QUndoStack) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QUndoStack) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQUndoStackFromPointer(cthis unsafe.Pointer) *QUndoStack {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QUndoStack{bcthis0}
}
func (*QUndoStack) NewFromPointer(cthis unsafe.Pointer) *QUndoStack {
	return NewQUndoStackFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qundostack.h:92
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

// /usr/include/qt/QtWidgets/qundostack.h:98
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

// /usr/include/qt/QtWidgets/qundostack.h:98
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

// /usr/include/qt/QtWidgets/qundostack.h:99
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

// /usr/include/qt/QtWidgets/qundostack.h:100
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

// /usr/include/qt/QtWidgets/qundostack.h:102
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

// /usr/include/qt/QtWidgets/qundostack.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canUndo() const

/*
Returns true if there is a command available for undo; otherwise returns false.

This function returns false if the stack is empty, or if the bottom command on the stack has already been undone.

Synonymous with index() == 0.

See also index() and canRedo().
*/
func (this *QUndoStack) CanUndo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack7canUndoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canRedo() const

/*
Returns true if there is a command available for redo; otherwise returns false.

This function returns false if the stack is empty or if the top command on the stack has already been redone.

Synonymous with index() == count().

See also index() and canUndo().
*/
func (this *QUndoStack) CanRedo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack7canRedoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:106
// index:0
// Public Visibility=Default Availability=Available
// [8] QString undoText() const

/*
Returns the text of the command which will be undone in the next call to undo().

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

// /usr/include/qt/QtWidgets/qundostack.h:107
// index:0
// Public Visibility=Default Availability=Available
// [8] QString redoText() const

/*
Returns the text of the command which will be redone in the next call to redo().

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

// /usr/include/qt/QtWidgets/qundostack.h:109
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

// /usr/include/qt/QtWidgets/qundostack.h:110
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

// /usr/include/qt/QtWidgets/qundostack.h:111
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

// /usr/include/qt/QtWidgets/qundostack.h:114
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

// /usr/include/qt/QtWidgets/qundostack.h:114
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

// /usr/include/qt/QtWidgets/qundostack.h:116
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

// /usr/include/qt/QtWidgets/qundostack.h:116
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

// /usr/include/qt/QtWidgets/qundostack.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive() const

/*

 */
func (this *QUndoStack) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isClean() const

/*
If the stack is in the clean state, returns true; otherwise returns false.

See also setClean() and cleanIndex().
*/
func (this *QUndoStack) IsClean() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack7isCleanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:122
// index:0
// Public Visibility=Default Availability=Available
// [4] int cleanIndex() const

/*
Returns the clean index. This is the index at which setClean() was called.

A stack may not have a clean index. This happens if a document is saved, some commands are undone, then a new command is pushed. Since push() deletes all the undone commands before pushing the new command, the stack can't return to the clean state again. In this case, this function returns -1. The -1 may also be returned after an explicit call to resetClean().

See also isClean() and setClean().
*/
func (this *QUndoStack) CleanIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack10cleanIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qundostack.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginMacro(const QString &)

/*
Begins composition of a macro command with the given text description.

An empty command described by the specified text is pushed on the stack. Any subsequent commands pushed on the stack will be appended to the empty command's children until endMacro() is called.

Calls to beginMacro() and endMacro() may be nested, but every call to beginMacro() must have a matching call to endMacro().

While a macro is being composed, the stack is disabled. This means that:


indexChanged() and cleanChanged() are not emitted,
canUndo() and canRedo() return false,
calling undo() or redo() has no effect,
the undo/redo actions are disabled.


The stack becomes enabled and appropriate signals are emitted when endMacro() is called for the outermost macro.


  stack.beginMacro("insert red text");
  stack.push(new InsertText(document, idx, text));
  stack.push(new SetColor(document, idx, text.length(), Qt::red));
  stack.endMacro(); // indexChanged() is emitted



This code is equivalent to:


  QUndoCommand *insertRed = new QUndoCommand(); // an empty command
  insertRed->setText("insert red text");

  new InsertText(document, idx, text, insertRed); // becomes child of insertRed
  new SetColor(document, idx, text.length(), Qt::red, insertRed);

  stack.push(insertRed);



See also endMacro().
*/
func (this *QUndoStack) BeginMacro(text string) {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack10beginMacroERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:125
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endMacro()

/*
Ends composition of a macro command.

If this is the outermost macro in a set nested macros, this function emits indexChanged() once for the entire macro command.

See also beginMacro().
*/
func (this *QUndoStack) EndMacro() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack8endMacroEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUndoLimit(int)

/*

 */
func (this *QUndoStack) SetUndoLimit(limit int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack12setUndoLimitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), limit)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] int undoLimit() const

/*

 */
func (this *QUndoStack) UndoLimit() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack9undoLimitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qundostack.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] const QUndoCommand * command(int) const

/*
Returns a const pointer to the command at index.

This function returns a const pointer, because modifying a command, once it has been pushed onto the stack and executed, almost always causes corruption of the state of the document, if the command is later undone or redone.

This function was introduced in  Qt 4.4.

See also QUndoCommand::child().
*/
func (this *QUndoStack) Command(index int) *QUndoCommand /*777 const QUndoCommand **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack7commandEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQUndoCommandFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundostack.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClean()

/*
Marks the stack as clean and emits cleanChanged() if the stack was not already clean.

This is typically called when a document is saved, for example.

Whenever the stack returns to this state through the use of undo/redo commands, it emits the signal cleanChanged(). This signal is also emitted when the stack leaves the clean state.

See also isClean(), resetClean(), and cleanIndex().
*/
func (this *QUndoStack) SetClean() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack8setCleanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetClean()

/*
Leaves the clean state and emits cleanChanged() if the stack was clean. This method resets the clean index to -1.

This is typically called in the following cases, when a document has been:


created basing on some template and has not been saved, so no filename has been associated with the document yet.
restored from a backup file.
changed outside of the editor and the user did not reload it.


This function was introduced in  Qt 5.8.

See also isClean(), setClean(), and cleanIndex().
*/
func (this *QUndoStack) ResetClean() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack10resetCleanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndex(int)

/*
Repeatedly calls undo() or redo() until the current command index reaches idx. This function can be used to roll the state of the document forwards of backwards. indexChanged() is emitted only once.

See also index(), count(), undo(), and redo().
*/
func (this *QUndoStack) SetIndex(idx int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack8setIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), idx)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undo()

/*
Undoes the command below the current command by calling QUndoCommand::undo(). Decrements the current command index.

If the stack is empty, or if the bottom command on the stack has already been undone, this function does nothing.

After the command is undone, if QUndoCommand::isObsolete() returns true, then the command will be deleted from the stack. Additionally, if the clean index is greater than or equal to the current command index, then the clean index is reset.

See also redo() and index().
*/
func (this *QUndoStack) Undo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack4undoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redo()

/*
Redoes the current command by calling QUndoCommand::redo(). Increments the current command index.

If the stack is empty, or if the top command on the stack has already been redone, this function does nothing.

If QUndoCommand::isObsolete() returns true for the current command, then the command will be deleted from the stack. Additionally, if the clean index is greater than or equal to the current command index, then the clean index is reset.

See also undo() and index().
*/
func (this *QUndoStack) Redo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack4redoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActive(bool)

/*

 */
func (this *QUndoStack) SetActive(active bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack9setActiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), active)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActive(bool)

/*

 */
func (this *QUndoStack) SetActivep() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	active := true
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack9setActiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), active)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void indexChanged(int)

/*
This signal is emitted whenever a command modifies the state of the document. This happens when a command is undone or redone. When a macro command is undone or redone, or setIndex() is called, this signal is emitted only once.

idx specifies the index of the current command, ie. the command which will be executed on the next call to redo().

See also index() and setIndex().
*/
func (this *QUndoStack) IndexChanged(idx int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack12indexChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), idx)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cleanChanged(bool)

/*
This signal is emitted whenever the stack enters or leaves the clean state. If clean is true, the stack is in a clean state; otherwise this signal indicates that it has left the clean state.

See also isClean() and setClean().
*/
func (this *QUndoStack) CleanChanged(clean bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack12cleanChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clean)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canUndoChanged(bool)

/*
This signal is emitted whenever the value of canUndo() changes. It is used to enable or disable the undo action returned by createUndoAction(). canUndo specifies the new value.
*/
func (this *QUndoStack) CanUndoChanged(canUndo bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack14canUndoChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), canUndo)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canRedoChanged(bool)

/*
This signal is emitted whenever the value of canRedo() changes. It is used to enable or disable the redo action returned by createRedoAction(). canRedo specifies the new value.
*/
func (this *QUndoStack) CanRedoChanged(canRedo bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack14canRedoChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), canRedo)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoTextChanged(const QString &)

/*
This signal is emitted whenever the value of undoText() changes. It is used to update the text property of the undo action returned by createUndoAction(). undoText specifies the new text.
*/
func (this *QUndoStack) UndoTextChanged(undoText string) {
	var tmpArg0 = qtcore.NewQString5(undoText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack15undoTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redoTextChanged(const QString &)

/*
This signal is emitted whenever the value of redoText() changes. It is used to update the text property of the redo action returned by createRedoAction(). redoText specifies the new text.
*/
func (this *QUndoStack) RedoTextChanged(redoText string) {
	var tmpArg0 = qtcore.NewQString5(redoText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack15redoTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
