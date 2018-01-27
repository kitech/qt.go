package qtwidgets

// /usr/include/qt/QtWidgets/qundogroup.h
// #include <qundogroup.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 65
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
type QUndoGroup struct {
	*qtcore.QObject
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QUndoGroup) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundogroup.h:61
// index:0
// Public
// void QUndoGroup(QObject *)
func NewQUndoGroup(parent *qtcore.QObject /*777 QObject **/) *QUndoGroup {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
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
// void addStack(QUndoStack *)
func (this *QUndoGroup) AddStack(stack *QUndoStack /*777 QUndoStack **/) {
	var convArg0 = stack.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup8addStackEP10QUndoStack", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:65
// index:0
// Public
// void removeStack(QUndoStack *)
func (this *QUndoGroup) RemoveStack(stack *QUndoStack /*777 QUndoStack **/) {
	var convArg0 = stack.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup11removeStackEP10QUndoStack", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:67
// index:0
// Public
// QUndoStack * activeStack()
func (this *QUndoGroup) ActiveStack() *QUndoStack /*777 QUndoStack **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup11activeStackEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQUndoStackFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundogroup.h:70
// index:0
// Public
// QAction * createUndoAction(QObject *, const QString &)
func (this *QUndoGroup) CreateUndoAction(parent *qtcore.QObject /*777 QObject **/, prefix *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = parent.GetCthis()
	var convArg1 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup16createUndoActionEP7QObjectRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundogroup.h:72
// index:0
// Public
// QAction * createRedoAction(QObject *, const QString &)
func (this *QUndoGroup) CreateRedoAction(parent *qtcore.QObject /*777 QObject **/, prefix *qtcore.QString) *QAction /*777 QAction **/ {
	var convArg0 = parent.GetCthis()
	var convArg1 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup16createRedoActionEP7QObjectRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qundogroup.h:75
// index:0
// Public
// bool canUndo()
func (this *QUndoGroup) CanUndo() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup7canUndoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundogroup.h:76
// index:0
// Public
// bool canRedo()
func (this *QUndoGroup) CanRedo() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup7canRedoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qundogroup.h:77
// index:0
// Public
// QString undoText()
func (this *QUndoGroup) UndoText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup8undoTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qundogroup.h:78
// index:0
// Public
// QString redoText()
func (this *QUndoGroup) RedoText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup8redoTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qundogroup.h:79
// index:0
// Public
// bool isClean()
func (this *QUndoGroup) IsClean() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QUndoGroup7isCleanEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
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
// void setActiveStack(QUndoStack *)
func (this *QUndoGroup) SetActiveStack(stack *QUndoStack /*777 QUndoStack **/) {
	var convArg0 = stack.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup14setActiveStackEP10QUndoStack", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:87
// index:0
// Public
// void activeStackChanged(QUndoStack *)
func (this *QUndoGroup) ActiveStackChanged(stack *QUndoStack /*777 QUndoStack **/) {
	var convArg0 = stack.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup18activeStackChangedEP10QUndoStack", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:88
// index:0
// Public
// void indexChanged(int)
func (this *QUndoGroup) IndexChanged(idx int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup12indexChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:89
// index:0
// Public
// void cleanChanged(bool)
func (this *QUndoGroup) CleanChanged(clean bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup12cleanChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), clean)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:90
// index:0
// Public
// void canUndoChanged(bool)
func (this *QUndoGroup) CanUndoChanged(canUndo bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup14canUndoChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), canUndo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:91
// index:0
// Public
// void canRedoChanged(bool)
func (this *QUndoGroup) CanRedoChanged(canRedo bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup14canRedoChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), canRedo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:92
// index:0
// Public
// void undoTextChanged(const QString &)
func (this *QUndoGroup) UndoTextChanged(undoText *qtcore.QString) {
	var convArg0 = undoText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup15undoTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qundogroup.h:93
// index:0
// Public
// void redoTextChanged(const QString &)
func (this *QUndoGroup) RedoTextChanged(redoText *qtcore.QString) {
	var convArg0 = redoText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QUndoGroup15redoTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
