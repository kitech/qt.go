//  header block begin
// /usr/include/qt/QtCore/qitemselectionmodel.h
// #include <qitemselectionmodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:139
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QItemSelectionModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:167
// index:0
// void QItemSelectionModel(class QAbstractItemModel *)
func NewQItemSelectionModel(model unsafe.Pointer) *QItemSelectionModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModelC2EP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, cthis, model)
	gopp.ErrPrint(err, rv)
	return &QItemSelectionModel{cthis}
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:168
// index:1
// void QItemSelectionModel(class QAbstractItemModel *, class QObject *)
func NewQItemSelectionModel_1(model unsafe.Pointer, parent unsafe.Pointer) *QItemSelectionModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModelC2EP18QAbstractItemModelP7QObject", ffiqt.FFI_TYPE_VOID, cthis, model, parent)
	gopp.ErrPrint(err, rv)
	return &QItemSelectionModel{cthis}
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:169
// index:0
// virtual
// void ~QItemSelectionModel()
func DeleteQItemSelectionModel(*QItemSelectionModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:171
// index:0
// QModelIndex currentIndex()
func (this *QItemSelectionModel) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:173
// index:0
// bool isSelected(const class QModelIndex &)
func (this *QItemSelectionModel) IsSelected(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel10isSelectedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:174
// index:0
// bool isRowSelected(int, const class QModelIndex &)
func (this *QItemSelectionModel) IsRowSelected(row int, parent unsafe.Pointer) {
	// 0: (, int row, const QModelIndex & parent), (&row, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel13isRowSelectedEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:175
// index:0
// bool isColumnSelected(int, const class QModelIndex &)
func (this *QItemSelectionModel) IsColumnSelected(column int, parent unsafe.Pointer) {
	// 0: (, int column, const QModelIndex & parent), (&column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel16isColumnSelectedEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:177
// index:0
// bool rowIntersectsSelection(int, const class QModelIndex &)
func (this *QItemSelectionModel) RowIntersectsSelection(row int, parent unsafe.Pointer) {
	// 0: (, int row, const QModelIndex & parent), (&row, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel22rowIntersectsSelectionEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:178
// index:0
// bool columnIntersectsSelection(int, const class QModelIndex &)
func (this *QItemSelectionModel) ColumnIntersectsSelection(column int, parent unsafe.Pointer) {
	// 0: (, int column, const QModelIndex & parent), (&column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel25columnIntersectsSelectionEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:180
// index:0
// bool hasSelection()
func (this *QItemSelectionModel) HasSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel12hasSelectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:182
// index:0
// QModelIndexList selectedIndexes()
func (this *QItemSelectionModel) SelectedIndexes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel15selectedIndexesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:183
// index:0
// QModelIndexList selectedRows(int)
func (this *QItemSelectionModel) SelectedRows(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel12selectedRowsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:184
// index:0
// QModelIndexList selectedColumns(int)
func (this *QItemSelectionModel) SelectedColumns(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel15selectedColumnsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:185
// index:0
// const QItemSelection selection()
func (this *QItemSelectionModel) Selection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel9selectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:188
// index:0
// const QAbstractItemModel * model()
func (this *QItemSelectionModel) Model() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QItemSelectionModel5modelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:189
// index:1
// QAbstractItemModel * model()
func (this *QItemSelectionModel) Model_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel5modelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:191
// index:0
// void setModel(class QAbstractItemModel *)
func (this *QItemSelectionModel) SetModel(model unsafe.Pointer) {
	// 0: (, QAbstractItemModel * model), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.cthis, model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:197
// index:0
// virtual
// void clear()
func (this *QItemSelectionModel) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:198
// index:0
// virtual
// void reset()
func (this *QItemSelectionModel) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel5resetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:200
// index:0
// void clearSelection()
func (this *QItemSelectionModel) ClearSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel14clearSelectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:201
// index:0
// virtual
// void clearCurrentIndex()
func (this *QItemSelectionModel) ClearCurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel17clearCurrentIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:204
// index:0
// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QItemSelectionModel) SelectionChanged(selected unsafe.Pointer, deselected unsafe.Pointer) {
	// 0: (, const QItemSelection & selected, const QItemSelection & deselected), (selected, deselected)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_VOID, this.cthis, selected, deselected)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:205
// index:0
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QItemSelectionModel) CurrentChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, const QModelIndex & current, const QModelIndex & previous), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.cthis, current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:206
// index:0
// void currentRowChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QItemSelectionModel) CurrentRowChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, const QModelIndex & current, const QModelIndex & previous), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel17currentRowChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.cthis, current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:207
// index:0
// void currentColumnChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QItemSelectionModel) CurrentColumnChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, const QModelIndex & current, const QModelIndex & previous), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel20currentColumnChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.cthis, current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:208
// index:0
// void modelChanged(class QAbstractItemModel *)
func (this *QItemSelectionModel) ModelChanged(model unsafe.Pointer) {
	// 0: (, QAbstractItemModel * model), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QItemSelectionModel12modelChangedEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.cthis, model)
	gopp.ErrPrint(err, rv)
}

//  body block end
