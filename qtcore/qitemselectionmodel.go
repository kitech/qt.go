package qtcore

// /usr/include/qt/QtCore/qitemselectionmodel.h
// #include <qitemselectionmodel.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
type QItemSelectionModel struct {
	*QObject
}

func (this *QItemSelectionModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QItemSelectionModel) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQItemSelectionModelFromPointer(cthis unsafe.Pointer) *QItemSelectionModel {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QItemSelectionModel{bcthis0}
}
func (*QItemSelectionModel) NewFromPointer(cthis unsafe.Pointer) *QItemSelectionModel {
	return NewQItemSelectionModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:139
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QItemSelectionModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:167
// index:0
// Public
// void QItemSelectionModel(QAbstractItemModel *)
func NewQItemSelectionModel(model *QAbstractItemModel /*777 QAbstractItemModel **/) *QItemSelectionModel {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = model.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModelC2EP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:168
// index:1
// Public
// void QItemSelectionModel(QAbstractItemModel *, QObject *)
func NewQItemSelectionModel_1(model *QAbstractItemModel /*777 QAbstractItemModel **/, parent *QObject /*777 QObject **/) *QItemSelectionModel {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = model.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModelC2EP18QAbstractItemModelP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQItemSelectionModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:169
// index:0
// Public virtual
// void ~QItemSelectionModel()
func DeleteQItemSelectionModel(*QItemSelectionModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:171
// index:0
// Public
// QModelIndex currentIndex()
func (this *QItemSelectionModel) CurrentIndex() *QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel12currentIndexEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:173
// index:0
// Public
// bool isSelected(const QModelIndex &)
func (this *QItemSelectionModel) IsSelected(index *QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel10isSelectedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:174
// index:0
// Public
// bool isRowSelected(int, const QModelIndex &)
func (this *QItemSelectionModel) IsRowSelected(row int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel13isRowSelectedEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:175
// index:0
// Public
// bool isColumnSelected(int, const QModelIndex &)
func (this *QItemSelectionModel) IsColumnSelected(column int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel16isColumnSelectedEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:177
// index:0
// Public
// bool rowIntersectsSelection(int, const QModelIndex &)
func (this *QItemSelectionModel) RowIntersectsSelection(row int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel22rowIntersectsSelectionEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:178
// index:0
// Public
// bool columnIntersectsSelection(int, const QModelIndex &)
func (this *QItemSelectionModel) ColumnIntersectsSelection(column int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel25columnIntersectsSelectionEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:180
// index:0
// Public
// bool hasSelection()
func (this *QItemSelectionModel) HasSelection() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel12hasSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:185
// index:0
// Public
// const QItemSelection selection()
func (this *QItemSelectionModel) Selection() *QItemSelection /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel9selectionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:188
// index:0
// Public
// const QAbstractItemModel * model()
func (this *QItemSelectionModel) Model() *QAbstractItemModel /*777 const QAbstractItemModel **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:189
// index:1
// Public
// QAbstractItemModel * model()
func (this *QItemSelectionModel) Model_1() *QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:191
// index:0
// Public
// void setModel(QAbstractItemModel *)
func (this *QItemSelectionModel) SetModel(model *QAbstractItemModel /*777 QAbstractItemModel **/) {
	var convArg0 = model.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:194
// index:0
// Public virtual
// void setCurrentIndex(const QModelIndex &, QItemSelectionModel::SelectionFlags)
func (this *QItemSelectionModel) SetCurrentIndex(index *QModelIndex, command int) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel15setCurrentIndexERK11QModelIndex6QFlagsINS_13SelectionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:195
// index:0
// Public virtual
// void select(const QModelIndex &, QItemSelectionModel::SelectionFlags)
func (this *QItemSelectionModel) Select(index *QModelIndex, command int) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel6selectERK11QModelIndex6QFlagsINS_13SelectionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:196
// index:1
// Public virtual
// void select(const QItemSelection &, QItemSelectionModel::SelectionFlags)
func (this *QItemSelectionModel) Select_1(selection *QItemSelection, command int) {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel6selectERK14QItemSelection6QFlagsINS_13SelectionFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:197
// index:0
// Public virtual
// void clear()
func (this *QItemSelectionModel) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:198
// index:0
// Public virtual
// void reset()
func (this *QItemSelectionModel) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:200
// index:0
// Public
// void clearSelection()
func (this *QItemSelectionModel) ClearSelection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel14clearSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:201
// index:0
// Public virtual
// void clearCurrentIndex()
func (this *QItemSelectionModel) ClearCurrentIndex() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel17clearCurrentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:204
// index:0
// Public
// void selectionChanged(const QItemSelection &, const QItemSelection &)
func (this *QItemSelectionModel) SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {
	var convArg0 = selected.GetCthis()
	var convArg1 = deselected.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:205
// index:0
// Public
// void currentChanged(const QModelIndex &, const QModelIndex &)
func (this *QItemSelectionModel) CurrentChanged(current *QModelIndex, previous *QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:206
// index:0
// Public
// void currentRowChanged(const QModelIndex &, const QModelIndex &)
func (this *QItemSelectionModel) CurrentRowChanged(current *QModelIndex, previous *QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel17currentRowChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:207
// index:0
// Public
// void currentColumnChanged(const QModelIndex &, const QModelIndex &)
func (this *QItemSelectionModel) CurrentColumnChanged(current *QModelIndex, previous *QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel20currentColumnChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:208
// index:0
// Public
// void modelChanged(QAbstractItemModel *)
func (this *QItemSelectionModel) ModelChanged(model *QAbstractItemModel /*777 QAbstractItemModel **/) {
	var convArg0 = model.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel12modelChangedEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:212
// index:0
// Protected
// void emitSelectionChanged(const QItemSelection &, const QItemSelection &)
func (this *QItemSelectionModel) EmitSelectionChanged(newSelection *QItemSelection, oldSelection *QItemSelection) {
	var convArg0 = newSelection.GetCthis()
	var convArg1 = oldSelection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel20emitSelectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

type QItemSelectionModel__SelectionFlag = int

const QItemSelectionModel__NoUpdate QItemSelectionModel__SelectionFlag = 0
const QItemSelectionModel__Clear QItemSelectionModel__SelectionFlag = 1
const QItemSelectionModel__Select QItemSelectionModel__SelectionFlag = 2
const QItemSelectionModel__Deselect QItemSelectionModel__SelectionFlag = 4
const QItemSelectionModel__Toggle QItemSelectionModel__SelectionFlag = 8
const QItemSelectionModel__Current QItemSelectionModel__SelectionFlag = 16
const QItemSelectionModel__Rows QItemSelectionModel__SelectionFlag = 32
const QItemSelectionModel__Columns QItemSelectionModel__SelectionFlag = 64
const QItemSelectionModel__SelectCurrent QItemSelectionModel__SelectionFlag = 18
const QItemSelectionModel__ToggleCurrent QItemSelectionModel__SelectionFlag = 24
const QItemSelectionModel__ClearAndSelect QItemSelectionModel__SelectionFlag = 3

//  body block end
