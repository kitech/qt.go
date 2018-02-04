package qtwidgets

// /usr/include/qt/QtWidgets/qundostack.h
// #include <qundostack.h>
// #include <QtWidgets>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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

// /usr/include/qt/QtWidgets/qundostack.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QUndoStack) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundostack.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUndoStack(QObject *)
func NewQUndoStack(parent *qtcore.QObject /*777 QObject **/) *QUndoStack {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStackC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQUndoStackFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qundostack.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QUndoStack()
func DeleteQUndoStack(this *QUndoStack) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStackD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qundostack.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QUndoStack) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void push(QUndoCommand *)
func (this *QUndoStack) Push(cmd *QUndoCommand /*777 QUndoCommand **/) {
	var convArg0 = cmd.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack4pushEP12QUndoCommand", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:101
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canUndo()
func (this *QUndoStack) CanUndo() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack7canUndoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canRedo()
func (this *QUndoStack) CanRedo() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack7canRedoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QString undoText()
func (this *QUndoStack) UndoText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack8undoTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qundostack.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QString redoText()
func (this *QUndoStack) RedoText() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack8redoTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qundostack.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QUndoStack) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qundostack.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int index()
func (this *QUndoStack) Index() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack5indexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qundostack.h:108
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text(int)
func (this *QUndoStack) Text(idx int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack4textEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), idx)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qundostack.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createUndoAction(QObject *, const QString &)
func (this *QUndoStack) CreateUndoAction(parent *qtcore.QObject /*777 QObject **/, prefix *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = parent.GetCthis()
	var convArg1 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack16createUndoActionEP7QObjectRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundostack.h:113
// index:0
// Public Visibility=Default Availability=Available
// [8] QAction * createRedoAction(QObject *, const QString &)
func (this *QUndoStack) CreateRedoAction(parent *qtcore.QObject /*777 QObject **/, prefix *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = parent.GetCthis()
	var convArg1 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack16createRedoActionEP7QObjectRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundostack.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive()
func (this *QUndoStack) IsActive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack8isActiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isClean()
func (this *QUndoStack) IsClean() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack7isCleanEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundostack.h:119
// index:0
// Public Visibility=Default Availability=Available
// [4] int cleanIndex()
func (this *QUndoStack) CleanIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack10cleanIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qundostack.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginMacro(const QString &)
func (this *QUndoStack) BeginMacro(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack10beginMacroERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endMacro()
func (this *QUndoStack) EndMacro() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack8endMacroEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUndoLimit(int)
func (this *QUndoStack) SetUndoLimit(limit int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack12setUndoLimitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), limit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:125
// index:0
// Public Visibility=Default Availability=Available
// [4] int undoLimit()
func (this *QUndoStack) UndoLimit() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack9undoLimitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qundostack.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] const QUndoCommand * command(int)
func (this *QUndoStack) Command(index int) *QUndoCommand /*777 const QUndoCommand **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoStack7commandEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQUndoCommandFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundostack.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClean()
func (this *QUndoStack) SetClean() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack8setCleanEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetClean()
func (this *QUndoStack) ResetClean() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack10resetCleanEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIndex(int)
func (this *QUndoStack) SetIndex(idx int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack8setIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undo()
func (this *QUndoStack) Undo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack4undoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redo()
func (this *QUndoStack) Redo() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack4redoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActive(_Bool)
func (this *QUndoStack) SetActive(active bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack9setActiveEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), active)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void indexChanged(int)
func (this *QUndoStack) IndexChanged(idx int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack12indexChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cleanChanged(_Bool)
func (this *QUndoStack) CleanChanged(clean bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack12cleanChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), clean)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canUndoChanged(_Bool)
func (this *QUndoStack) CanUndoChanged(canUndo bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack14canUndoChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), canUndo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void canRedoChanged(_Bool)
func (this *QUndoStack) CanRedoChanged(canRedo bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack14canRedoChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), canRedo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void undoTextChanged(const QString &)
func (this *QUndoStack) UndoTextChanged(undoText *qtcore.QString) {
	var convArg0 = undoText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack15undoTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundostack.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void redoTextChanged(const QString &)
func (this *QUndoStack) RedoTextChanged(redoText *qtcore.QString) {
	var convArg0 = redoText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoStack15redoTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
