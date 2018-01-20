//  header block begin
// /usr/include/qt/QtWidgets/qundogroup.h
// #include <qundogroup.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 67
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
	*qtcore.QObject
}

func (this *QUndoGroup) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQUndoGroupFromPointer(cthis unsafe.Pointer) *QUndoGroup {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QUndoGroup{bcthis0}
}

// /usr/include/qt/QtWidgets/qundogroup.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QUndoGroup) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundogroup.h:61
// index:0
// Public
// void QUndoGroup(class QObject *)
func NewQUndoGroup(parent unsafe.Pointer) *QUndoGroup {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoGroupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qundogroup.h:62
// index:0
// Public virtual
// void ~QUndoGroup()
func DeleteQUndoGroup(*QUndoGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:64
// index:0
// Public
// void addStack(class QUndoStack *)
func (this *QUndoGroup) AddStack(stack unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup8addStackEP10QUndoStack", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stack)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:65
// index:0
// Public
// void removeStack(class QUndoStack *)
func (this *QUndoGroup) RemoveStack(stack unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup11removeStackEP10QUndoStack", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stack)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:66
// index:0
// Public
// QList<QUndoStack *> stacks()
func (this *QUndoGroup) Stacks() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup6stacksEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundogroup.h:67
// index:0
// Public
// QUndoStack * activeStack()
func (this *QUndoGroup) ActiveStack() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup11activeStackEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundogroup.h:70
// index:0
// Public
// QAction * createUndoAction(class QObject *, const class QString &)
func (this *QUndoGroup) CreateUndoAction(parent unsafe.Pointer, prefix *qtcore.QString) interface{} {
	var convArg1 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), parent, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundogroup.h:72
// index:0
// Public
// QAction * createRedoAction(class QObject *, const class QString &)
func (this *QUndoGroup) CreateRedoAction(parent unsafe.Pointer, prefix *qtcore.QString) interface{} {
	var convArg1 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), parent, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundogroup.h:75
// index:0
// Public
// bool canUndo()
func (this *QUndoGroup) CanUndo() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup7canUndoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundogroup.h:76
// index:0
// Public
// bool canRedo()
func (this *QUndoGroup) CanRedo() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup7canRedoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundogroup.h:77
// index:0
// Public
// QString undoText()
func (this *QUndoGroup) UndoText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup8undoTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundogroup.h:78
// index:0
// Public
// QString redoText()
func (this *QUndoGroup) RedoText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup8redoTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundogroup.h:79
// index:0
// Public
// bool isClean()
func (this *QUndoGroup) IsClean() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup7isCleanEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qundogroup.h:82
// index:0
// Public
// void undo()
func (this *QUndoGroup) Undo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup4undoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:83
// index:0
// Public
// void redo()
func (this *QUndoGroup) Redo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup4redoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:84
// index:0
// Public
// void setActiveStack(class QUndoStack *)
func (this *QUndoGroup) SetActiveStack(stack unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup14setActiveStackEP10QUndoStack", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stack)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:87
// index:0
// Public
// void activeStackChanged(class QUndoStack *)
func (this *QUndoGroup) ActiveStackChanged(stack unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup18activeStackChangedEP10QUndoStack", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stack)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:88
// index:0
// Public
// void indexChanged(int)
func (this *QUndoGroup) IndexChanged(idx int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup12indexChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:89
// index:0
// Public
// void cleanChanged(_Bool)
func (this *QUndoGroup) CleanChanged(clean bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup12cleanChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &clean)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:90
// index:0
// Public
// void canUndoChanged(_Bool)
func (this *QUndoGroup) CanUndoChanged(canUndo bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup14canUndoChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &canUndo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:91
// index:0
// Public
// void canRedoChanged(_Bool)
func (this *QUndoGroup) CanRedoChanged(canRedo bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup14canRedoChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &canRedo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:92
// index:0
// Public
// void undoTextChanged(const class QString &)
func (this *QUndoGroup) UndoTextChanged(undoText *qtcore.QString) {
	var convArg0 = undoText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup15undoTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:93
// index:0
// Public
// void redoTextChanged(const class QString &)
func (this *QUndoGroup) RedoTextChanged(redoText *qtcore.QString) {
	var convArg0 = redoText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup15redoTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
