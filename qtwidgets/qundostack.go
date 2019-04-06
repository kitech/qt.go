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

// /usr/include/qt/QtWidgets/qundostack.h:122
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

// /usr/include/qt/QtWidgets/qundostack.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isClean() const

/*
If the stack is in the clean state, returns true; otherwise returns false.

Note: Getter function for property clean.

See also setClean() and cleanIndex().
*/
func (this *QUndoStack) IsClean() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoStack7isCleanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:124
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

// /usr/include/qt/QtWidgets/qundostack.h:126
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

// /usr/include/qt/QtWidgets/qundostack.h:127
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

// /usr/include/qt/QtWidgets/qundostack.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUndoLimit(int)

/*

 */
func (this *QUndoStack) SetUndoLimit(limit int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack12setUndoLimitEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), limit)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:130
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

// /usr/include/qt/QtWidgets/qundostack.h:132
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

// /usr/include/qt/QtWidgets/qundostack.h:135
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

// /usr/include/qt/QtWidgets/qundostack.h:136
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

// /usr/include/qt/QtWidgets/qundostack.h:137
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

// /usr/include/qt/QtWidgets/qundostack.h:138
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

// /usr/include/qt/QtWidgets/qundostack.h:139
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

// /usr/include/qt/QtWidgets/qundostack.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActive(bool)

/*

 */
func (this *QUndoStack) SetActive(active bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack9setActiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), active)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:140
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

// /usr/include/qt/QtWidgets/qundostack.h:143
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

// /usr/include/qt/QtWidgets/qundostack.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cleanChanged(bool)

/*
This signal is emitted whenever the stack enters or leaves the clean state. If clean is true, the stack is in a clean state; otherwise this signal indicates that it has left the clean state.

Note: Notifier signal for property clean.

See also isClean() and setClean().
*/
func (this *QUndoStack) CleanChanged(clean bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack12cleanChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clean)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canUndoChanged(bool)

/*
This signal is emitted whenever the value of canUndo() changes. It is used to enable or disable the undo action returned by createUndoAction(). canUndo specifies the new value.

Note: Notifier signal for property canUndo.
*/
func (this *QUndoStack) CanUndoChanged(canUndo bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack14canUndoChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), canUndo)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canRedoChanged(bool)

/*
This signal is emitted whenever the value of canRedo() changes. It is used to enable or disable the redo action returned by createRedoAction(). canRedo specifies the new value.

Note: Notifier signal for property canRedo.
*/
func (this *QUndoStack) CanRedoChanged(canRedo bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack14canRedoChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), canRedo)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoTextChanged(const QString &)

/*
This signal is emitted whenever the value of undoText() changes. It is used to update the text property of the undo action returned by createUndoAction(). undoText specifies the new text.

Note: Notifier signal for property undoText.
*/
func (this *QUndoStack) UndoTextChanged(undoText string) {
	var tmpArg0 = qtcore.NewQString5(undoText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack15undoTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redoTextChanged(const QString &)

/*
This signal is emitted whenever the value of redoText() changes. It is used to update the text property of the redo action returned by createRedoAction(). redoText specifies the new text.

Note: Notifier signal for property redoText.
*/
func (this *QUndoStack) RedoTextChanged(redoText string) {
	var tmpArg0 = qtcore.NewQString5(redoText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoStack15redoTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11367() {
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
