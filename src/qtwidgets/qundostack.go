//  header block begin
// /usr/include/qt/QtWidgets/qundostack.h
// #include <qundostack.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QUndoStack struct {
	*qtcore.QObject
}

func (this *QUndoStack) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtWidgets/qundostack.h:89
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QUndoStack) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:95
// index:0
// void QUndoStack(class QObject *)
func NewQUndoStack(parent unsafe.Pointer) *QUndoStack {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStackC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoStackFromPointer(cthis)
	return gothis
}
func NewQUndoStackFromPointer(cthis unsafe.Pointer) *QUndoStack {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QUndoStack{bcthis0}
}

// /usr/include/qt/QtWidgets/qundostack.h:96
// index:0
// virtual
// void ~QUndoStack()
func DeleteQUndoStack(*QUndoStack) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStackD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:97
// index:0
// void clear()
func (this *QUndoStack) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:99
// index:0
// void push(class QUndoCommand *)
func (this *QUndoStack) Push(cmd unsafe.Pointer) {
	// 0: (, cmd QUndoCommand *), (cmd)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack4pushEP12QUndoCommand", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cmd)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:101
// index:0
// bool canUndo()
func (this *QUndoStack) CanUndo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack7canUndoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:102
// index:0
// bool canRedo()
func (this *QUndoStack) CanRedo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack7canRedoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:103
// index:0
// QString undoText()
func (this *QUndoStack) UndoText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack8undoTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:104
// index:0
// QString redoText()
func (this *QUndoStack) RedoText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack8redoTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:106
// index:0
// int count()
func (this *QUndoStack) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack5countEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:107
// index:0
// int index()
func (this *QUndoStack) Index() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack5indexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:108
// index:0
// QString text(int)
func (this *QUndoStack) Text(idx int) {
	// 0: (, idx int), (&idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack4textEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:111
// index:0
// QAction * createUndoAction(class QObject *, const class QString &)
func (this *QUndoStack) CreateUndoAction(parent unsafe.Pointer, prefix unsafe.Pointer) {
	// 0: (, parent QObject *, prefix const QString &), (parent, prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, prefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:113
// index:0
// QAction * createRedoAction(class QObject *, const class QString &)
func (this *QUndoStack) CreateRedoAction(parent unsafe.Pointer, prefix unsafe.Pointer) {
	// 0: (, parent QObject *, prefix const QString &), (parent, prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, prefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:117
// index:0
// bool isActive()
func (this *QUndoStack) IsActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack8isActiveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:118
// index:0
// bool isClean()
func (this *QUndoStack) IsClean() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack7isCleanEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:119
// index:0
// int cleanIndex()
func (this *QUndoStack) CleanIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack10cleanIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:121
// index:0
// void beginMacro(const class QString &)
func (this *QUndoStack) BeginMacro(text unsafe.Pointer) {
	// 0: (, text const QString &), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack10beginMacroERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:122
// index:0
// void endMacro()
func (this *QUndoStack) EndMacro() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack8endMacroEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:124
// index:0
// void setUndoLimit(int)
func (this *QUndoStack) SetUndoLimit(limit int) {
	// 0: (, limit int), (&limit)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack12setUndoLimitEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &limit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:125
// index:0
// int undoLimit()
func (this *QUndoStack) UndoLimit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack9undoLimitEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:127
// index:0
// const QUndoCommand * command(int)
func (this *QUndoStack) Command(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack7commandEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:130
// index:0
// void setClean()
func (this *QUndoStack) SetClean() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack8setCleanEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:131
// index:0
// void resetClean()
func (this *QUndoStack) ResetClean() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack10resetCleanEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:132
// index:0
// void setIndex(int)
func (this *QUndoStack) SetIndex(idx int) {
	// 0: (, idx int), (&idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack8setIndexEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:133
// index:0
// void undo()
func (this *QUndoStack) Undo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack4undoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:134
// index:0
// void redo()
func (this *QUndoStack) Redo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack4redoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:135
// index:0
// void setActive(_Bool)
func (this *QUndoStack) SetActive(active bool) {
	// 0: (, active bool), (&active)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack9setActiveEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &active)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:138
// index:0
// void indexChanged(int)
func (this *QUndoStack) IndexChanged(idx int) {
	// 0: (, idx int), (&idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack12indexChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:139
// index:0
// void cleanChanged(_Bool)
func (this *QUndoStack) CleanChanged(clean bool) {
	// 0: (, clean bool), (&clean)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack12cleanChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &clean)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:140
// index:0
// void canUndoChanged(_Bool)
func (this *QUndoStack) CanUndoChanged(canUndo bool) {
	// 0: (, canUndo bool), (&canUndo)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack14canUndoChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &canUndo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:141
// index:0
// void canRedoChanged(_Bool)
func (this *QUndoStack) CanRedoChanged(canRedo bool) {
	// 0: (, canRedo bool), (&canRedo)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack14canRedoChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &canRedo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:142
// index:0
// void undoTextChanged(const class QString &)
func (this *QUndoStack) UndoTextChanged(undoText unsafe.Pointer) {
	// 0: (, undoText const QString &), (undoText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack15undoTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), undoText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:143
// index:0
// void redoTextChanged(const class QString &)
func (this *QUndoStack) RedoTextChanged(redoText unsafe.Pointer) {
	// 0: (, redoText const QString &), (redoText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack15redoTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), redoText)
	gopp.ErrPrint(err, rv)
}

//  body block end
