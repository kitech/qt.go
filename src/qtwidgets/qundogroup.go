//  header block begin
// /usr/include/qt/QtWidgets/qundogroup.h
// #include <qundogroup.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 58
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
type QUndoGroup struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qundogroup.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QUndoGroup) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:61
// index:0
// void QUndoGroup(class QObject *)
func NewQUndoGroup(parent unsafe.Pointer) *QUndoGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QUndoGroup{cthis}
}

// /usr/include/qt/QtWidgets/qundogroup.h:62
// index:0
// virtual
// void ~QUndoGroup()
func DeleteQUndoGroup(*QUndoGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:64
// index:0
// void addStack(class QUndoStack *)
func (this *QUndoGroup) AddStack(stack unsafe.Pointer) {
	// 0: (, QUndoStack * stack), (stack)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup8addStackEP10QUndoStack", ffiqt.FFI_TYPE_VOID, this.cthis, stack)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:65
// index:0
// void removeStack(class QUndoStack *)
func (this *QUndoGroup) RemoveStack(stack unsafe.Pointer) {
	// 0: (, QUndoStack * stack), (stack)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup11removeStackEP10QUndoStack", ffiqt.FFI_TYPE_VOID, this.cthis, stack)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:66
// index:0
// QList<QUndoStack *> stacks()
func (this *QUndoGroup) Stacks() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup6stacksEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:67
// index:0
// QUndoStack * activeStack()
func (this *QUndoGroup) ActiveStack() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup11activeStackEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:70
// index:0
// QAction * createUndoAction(class QObject *, const class QString &)
func (this *QUndoGroup) CreateUndoAction(parent unsafe.Pointer, prefix unsafe.Pointer) {
	// 0: (, QObject * parent, const QString & prefix), (parent, prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, parent, prefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:72
// index:0
// QAction * createRedoAction(class QObject *, const class QString &)
func (this *QUndoGroup) CreateRedoAction(parent unsafe.Pointer, prefix unsafe.Pointer) {
	// 0: (, QObject * parent, const QString & prefix), (parent, prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, parent, prefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:75
// index:0
// bool canUndo()
func (this *QUndoGroup) CanUndo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup7canUndoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:76
// index:0
// bool canRedo()
func (this *QUndoGroup) CanRedo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup7canRedoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:77
// index:0
// QString undoText()
func (this *QUndoGroup) UndoText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup8undoTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:78
// index:0
// QString redoText()
func (this *QUndoGroup) RedoText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup8redoTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:79
// index:0
// bool isClean()
func (this *QUndoGroup) IsClean() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup7isCleanEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:82
// index:0
// void undo()
func (this *QUndoGroup) Undo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup4undoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:83
// index:0
// void redo()
func (this *QUndoGroup) Redo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup4redoEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:84
// index:0
// void setActiveStack(class QUndoStack *)
func (this *QUndoGroup) SetActiveStack(stack unsafe.Pointer) {
	// 0: (, QUndoStack * stack), (stack)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup14setActiveStackEP10QUndoStack", ffiqt.FFI_TYPE_VOID, this.cthis, stack)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:87
// index:0
// void activeStackChanged(class QUndoStack *)
func (this *QUndoGroup) ActiveStackChanged(stack unsafe.Pointer) {
	// 0: (, QUndoStack * stack), (stack)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup18activeStackChangedEP10QUndoStack", ffiqt.FFI_TYPE_VOID, this.cthis, stack)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:88
// index:0
// void indexChanged(int)
func (this *QUndoGroup) IndexChanged(idx int) {
	// 0: (, int idx), (&idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup12indexChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:89
// index:0
// void cleanChanged(_Bool)
func (this *QUndoGroup) CleanChanged(clean bool) {
	// 0: (, bool clean), (&clean)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup12cleanChangedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &clean)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:90
// index:0
// void canUndoChanged(_Bool)
func (this *QUndoGroup) CanUndoChanged(canUndo bool) {
	// 0: (, bool canUndo), (&canUndo)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup14canUndoChangedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &canUndo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:91
// index:0
// void canRedoChanged(_Bool)
func (this *QUndoGroup) CanRedoChanged(canRedo bool) {
	// 0: (, bool canRedo), (&canRedo)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup14canRedoChangedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &canRedo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:92
// index:0
// void undoTextChanged(const class QString &)
func (this *QUndoGroup) UndoTextChanged(undoText unsafe.Pointer) {
	// 0: (, const QString & undoText), (undoText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup15undoTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, undoText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:93
// index:0
// void redoTextChanged(const class QString &)
func (this *QUndoGroup) RedoTextChanged(redoText unsafe.Pointer) {
	// 0: (, const QString & redoText), (redoText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup15redoTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, redoText)
	gopp.ErrPrint(err, rv)
}

//  body block end
