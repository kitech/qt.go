package qtwidgets

// /usr/include/qt/QtWidgets/qundogroup.h
// #include <qundogroup.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 66
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
type QUndoGroup struct {
	*qtcore.QObject
}
type QUndoGroup_ITF interface {
	qtcore.QObject_ITF
	QUndoGroup_PTR() *QUndoGroup
}

func (ptr *QUndoGroup) QUndoGroup_PTR() *QUndoGroup { return ptr }

func (this *QUndoGroup) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QUndoGroup) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQUndoGroupFromPointer(cthis unsafe.Pointer) *QUndoGroup {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QUndoGroup{bcthis0}
}
func (*QUndoGroup) NewFromPointer(cthis unsafe.Pointer) *QUndoGroup {
	return NewQUndoGroupFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qundogroup.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QUndoGroup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoGroup(QObject *)

/*
Creates an empty QUndoGroup object with parent parent.

See also addStack().
*/
func (*QUndoGroup) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QUndoGroup {
	return NewQUndoGroup(parent)
}
func NewQUndoGroup(parent qtcore.QObject_ITF /*777 QObject **/) *QUndoGroup {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoGroup")
	return gothis
}

// /usr/include/qt/QtWidgets/qundogroup.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoGroup(QObject *)

/*
Creates an empty QUndoGroup object with parent parent.

See also addStack().
*/
func (*QUndoGroup) NewForInheritp() *QUndoGroup {
	return NewQUndoGroupp()
}
func NewQUndoGroupp() *QUndoGroup {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroupC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUndoGroupFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QUndoGroup")
	return gothis
}

// /usr/include/qt/QtWidgets/qundogroup.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUndoGroup()

/*

 */
func DeleteQUndoGroup(this *QUndoGroup) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroupD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qundogroup.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addStack(QUndoStack *)

/*
Adds stack to this group. The group does not take ownership of the stack. Another way of adding a stack to a group is by specifying the group as the stack's parent QObject in QUndoStack::QUndoStack(). In this case, the stack is deleted when the group is deleted, in the usual manner of QObjects.

See also removeStack(), stacks(), and QUndoStack::QUndoStack().
*/
func (this *QUndoGroup) AddStack(stack QUndoStack_ITF /*777 QUndoStack **/) {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup8addStackEP10QUndoStack", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeStack(QUndoStack *)

/*
Removes stack from this group. If the stack was the active stack in the group, the active stack becomes 0.

See also addStack(), stacks(), and QUndoStack::~QUndoStack().
*/
func (this *QUndoGroup) RemoveStack(stack QUndoStack_ITF /*777 QUndoStack **/) {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup11removeStackEP10QUndoStack", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QUndoStack * activeStack() const

/*
Returns the active stack of this group.

If none of the stacks are active, or if the group is empty, this function returns 0.

See also setActiveStack() and QUndoStack::setActive().
*/
func (this *QUndoGroup) ActiveStack() *QUndoStack /*777 QUndoStack **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup11activeStackEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQUndoStackFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createUndoAction(QObject *, const QString &) const

/*
Creates an undo QAction object with parent parent.

Triggering this action will cause a call to QUndoStack::undo() on the active stack. The text of this action will always be the text of the command which will be undone in the next call to undo(), prefixed by prefix. If there is no command available for undo, if the group is empty or if none of the stacks are active, this action will be disabled.

If prefix is empty, the default template "Undo %1" is used instead of prefix. Before Qt 4.8, the prefix "Undo" was used by default.

See also createRedoAction(), canUndo(), and QUndoCommand::text().
*/
func (this *QUndoGroup) CreateUndoAction(parent qtcore.QObject_ITF /*777 QObject **/, prefix string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(prefix)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createUndoAction(QObject *, const QString &) const

/*
Creates an undo QAction object with parent parent.

Triggering this action will cause a call to QUndoStack::undo() on the active stack. The text of this action will always be the text of the command which will be undone in the next call to undo(), prefixed by prefix. If there is no command available for undo, if the group is empty or if none of the stacks are active, this action will be disabled.

If prefix is empty, the default template "Undo %1" is used instead of prefix. Before Qt 4.8, the prefix "Undo" was used by default.

See also createRedoAction(), canUndo(), and QUndoCommand::text().
*/
func (this *QUndoGroup) CreateUndoActionp(parent qtcore.QObject_ITF /*777 QObject **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createRedoAction(QObject *, const QString &) const

/*
Creates an redo QAction object with parent parent.

Triggering this action will cause a call to QUndoStack::redo() on the active stack. The text of this action will always be the text of the command which will be redone in the next call to redo(), prefixed by prefix. If there is no command available for redo, if the group is empty or if none of the stacks are active, this action will be disabled.

If prefix is empty, the default template "Redo %1" is used instead of prefix. Before Qt 4.8, the prefix "Redo" was used by default.

See also createUndoAction(), canRedo(), and QUndoCommand::text().
*/
func (this *QUndoGroup) CreateRedoAction(parent qtcore.QObject_ITF /*777 QObject **/, prefix string) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(prefix)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createRedoAction(QObject *, const QString &) const

/*
Creates an redo QAction object with parent parent.

Triggering this action will cause a call to QUndoStack::redo() on the active stack. The text of this action will always be the text of the command which will be redone in the next call to redo(), prefixed by prefix. If there is no command available for redo, if the group is empty or if none of the stacks are active, this action will be disabled.

If prefix is empty, the default template "Redo %1" is used instead of prefix. Before Qt 4.8, the prefix "Redo" was used by default.

See also createUndoAction(), canRedo(), and QUndoCommand::text().
*/
func (this *QUndoGroup) CreateRedoActionp(parent qtcore.QObject_ITF /*777 QObject **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qundogroup.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canUndo() const

/*
Returns the value of the active stack's QUndoStack::canUndo().

If none of the stacks are active, or if the group is empty, this function returns false.

See also canRedo() and setActiveStack().
*/
func (this *QUndoGroup) CanUndo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup7canUndoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundogroup.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canRedo() const

/*
Returns the value of the active stack's QUndoStack::canRedo().

If none of the stacks are active, or if the group is empty, this function returns false.

See also canUndo() and setActiveStack().
*/
func (this *QUndoGroup) CanRedo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup7canRedoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundogroup.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QString undoText() const

/*
Returns the value of the active stack's QUndoStack::undoText().

If none of the stacks are active, or if the group is empty, this function returns an empty string.

See also redoText() and setActiveStack().
*/
func (this *QUndoGroup) UndoText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup8undoTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundogroup.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QString redoText() const

/*
Returns the value of the active stack's QUndoStack::redoText().

If none of the stacks are active, or if the group is empty, this function returns an empty string.

See also undoText() and setActiveStack().
*/
func (this *QUndoGroup) RedoText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup8redoTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qundogroup.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isClean() const

/*
Returns the value of the active stack's QUndoStack::isClean().

If none of the stacks are active, or if the group is empty, this function returns true.

See also setActiveStack().
*/
func (this *QUndoGroup) IsClean() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QUndoGroup7isCleanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundogroup.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undo()

/*
Calls QUndoStack::undo() on the active stack.

If none of the stacks are active, or if the group is empty, this function does nothing.

See also redo(), canUndo(), and setActiveStack().
*/
func (this *QUndoGroup) Undo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup4undoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redo()

/*
Calls QUndoStack::redo() on the active stack.

If none of the stacks are active, or if the group is empty, this function does nothing.

See also undo(), canRedo(), and setActiveStack().
*/
func (this *QUndoGroup) Redo() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup4redoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActiveStack(QUndoStack *)

/*
Sets the active stack of this group to stack.

If the stack is not a member of this group, this function does nothing.

Synonymous with calling QUndoStack::setActive() on stack.

The actions returned by createUndoAction() and createRedoAction() will now behave in the same way as those returned by stack's QUndoStack::createUndoAction() and QUndoStack::createRedoAction().

See also QUndoStack::setActive() and activeStack().
*/
func (this *QUndoGroup) SetActiveStack(stack QUndoStack_ITF /*777 QUndoStack **/) {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup14setActiveStackEP10QUndoStack", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeStackChanged(QUndoStack *)

/*
This signal is emitted whenever the active stack of the group changes. This can happen when setActiveStack() or QUndoStack::setActive() is called, or when the active stack is removed form the group. stack is the new active stack. If no stack is active, stack is 0.

See also setActiveStack() and QUndoStack::setActive().
*/
func (this *QUndoGroup) ActiveStackChanged(stack QUndoStack_ITF /*777 QUndoStack **/) {
	var convArg0 unsafe.Pointer
	if stack != nil && stack.QUndoStack_PTR() != nil {
		convArg0 = stack.QUndoStack_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup18activeStackChangedEP10QUndoStack", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void indexChanged(int)

/*
This signal is emitted whenever the active stack emits QUndoStack::indexChanged() or the active stack changes.

idx is the new current index, or 0 if the active stack is 0.

See also QUndoStack::indexChanged() and setActiveStack().
*/
func (this *QUndoGroup) IndexChanged(idx int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup12indexChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), idx)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cleanChanged(bool)

/*
This signal is emitted whenever the active stack emits QUndoStack::cleanChanged() or the active stack changes.

clean is the new state, or true if the active stack is 0.

See also QUndoStack::cleanChanged() and setActiveStack().
*/
func (this *QUndoGroup) CleanChanged(clean bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup12cleanChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), clean)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canUndoChanged(bool)

/*
This signal is emitted whenever the active stack emits QUndoStack::canUndoChanged() or the active stack changes.

canUndo is the new state, or false if the active stack is 0.

See also QUndoStack::canUndoChanged() and setActiveStack().
*/
func (this *QUndoGroup) CanUndoChanged(canUndo bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup14canUndoChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), canUndo)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canRedoChanged(bool)

/*
This signal is emitted whenever the active stack emits QUndoStack::canRedoChanged() or the active stack changes.

canRedo is the new state, or false if the active stack is 0.

See also QUndoStack::canRedoChanged() and setActiveStack().
*/
func (this *QUndoGroup) CanRedoChanged(canRedo bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup14canRedoChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), canRedo)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoTextChanged(const QString &)

/*
This signal is emitted whenever the active stack emits QUndoStack::undoTextChanged() or the active stack changes.

undoText is the new state, or an empty string if the active stack is 0.

See also QUndoStack::undoTextChanged() and setActiveStack().
*/
func (this *QUndoGroup) UndoTextChanged(undoText string) {
	var tmpArg0 = qtcore.NewQString5(undoText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup15undoTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redoTextChanged(const QString &)

/*
This signal is emitted whenever the active stack emits QUndoStack::redoTextChanged() or the active stack changes.

redoText is the new state, or an empty string if the active stack is 0.

See also QUndoStack::redoTextChanged() and setActiveStack().
*/
func (this *QUndoGroup) RedoTextChanged(redoText string) {
	var tmpArg0 = qtcore.NewQString5(redoText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QUndoGroup15redoTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
