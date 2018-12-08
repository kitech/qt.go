package qtcore

// /usr/include/qt/QtCore/qitemselectionmodel.h
// #include <qitemselectionmodel.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void emitSelectionChanged(const QItemSelection &, const QItemSelection &)
func (this *QItemSelectionModel) InheritEmitSelectionChanged(f func(newSelection *QItemSelection, oldSelection *QItemSelection) /*void*/) {
	qtrt.SetAllInheritCallback(this, "emitSelectionChanged", f)
}

/*

 */
type QItemSelectionModel struct {
	*QObject
}
type QItemSelectionModel_ITF interface {
	QObject_ITF
	QItemSelectionModel_PTR() *QItemSelectionModel
}

func (ptr *QItemSelectionModel) QItemSelectionModel_PTR() *QItemSelectionModel { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QItemSelectionModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QItemSelectionModel(QAbstractItemModel *)

/*
Constructs a selection model that operates on the specified item model.
*/
func (*QItemSelectionModel) NewForInherit(model QAbstractItemModel_ITF /*777 QAbstractItemModel **/) *QItemSelectionModel {
	return NewQItemSelectionModel(model)
}
func NewQItemSelectionModel(model QAbstractItemModel_ITF /*777 QAbstractItemModel **/) *QItemSelectionModel {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModelC2EP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QItemSelectionModel")
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:167
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QItemSelectionModel(QAbstractItemModel *)

/*
Constructs a selection model that operates on the specified item model.
*/
func (*QItemSelectionModel) NewForInheritp() *QItemSelectionModel {
	return NewQItemSelectionModelp()
}
func NewQItemSelectionModelp() *QItemSelectionModel {
	// arg: 0, QAbstractItemModel *=Pointer, QAbstractItemModel=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModelC2EP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QItemSelectionModel")
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:168
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QItemSelectionModel(QAbstractItemModel *, QObject *)

/*
Constructs a selection model that operates on the specified item model.
*/
func (*QItemSelectionModel) NewForInherit1(model QAbstractItemModel_ITF /*777 QAbstractItemModel **/, parent QObject_ITF /*777 QObject **/) *QItemSelectionModel {
	return NewQItemSelectionModel1(model, parent)
}
func NewQItemSelectionModel1(model QAbstractItemModel_ITF /*777 QAbstractItemModel **/, parent QObject_ITF /*777 QObject **/) *QItemSelectionModel {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModelC2EP18QAbstractItemModelP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQItemSelectionModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QItemSelectionModel")
	return gothis
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:169
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QItemSelectionModel()

/*

 */
func DeleteQItemSelectionModel(this *QItemSelectionModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:171
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex currentIndex() const

/*
Returns the model item index for the current item, or an invalid index if there is no current item.

See also setCurrentIndex().
*/
func (this *QItemSelectionModel) CurrentIndex() *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSelected(const QModelIndex &) const

/*
Returns true if the given model item index is selected.
*/
func (this *QItemSelectionModel) IsSelected(index QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel10isSelectedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:174
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRowSelected(int, const QModelIndex &) const

/*
Returns true if all items are selected in the row with the given parent.

Note that this function is usually faster than calling isSelected() on all items in the same row and that unselectable items are ignored.
*/
func (this *QItemSelectionModel) IsRowSelected(row int, parent QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel13isRowSelectedEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:175
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isColumnSelected(int, const QModelIndex &) const

/*
Returns true if all items are selected in the column with the given parent.

Note that this function is usually faster than calling isSelected() on all items in the same column and that unselectable items are ignored.
*/
func (this *QItemSelectionModel) IsColumnSelected(column int, parent QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel16isColumnSelectedEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:177
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rowIntersectsSelection(int, const QModelIndex &) const

/*
Returns true if there are any items selected in the row with the given parent.
*/
func (this *QItemSelectionModel) RowIntersectsSelection(row int, parent QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel22rowIntersectsSelectionEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:178
// index:0
// Public Visibility=Default Availability=Available
// [1] bool columnIntersectsSelection(int, const QModelIndex &) const

/*
Returns true if there are any items selected in the column with the given parent.
*/
func (this *QItemSelectionModel) ColumnIntersectsSelection(column int, parent QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel25columnIntersectsSelectionEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:180
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasSelection() const

/*
Returns true if the selection model contains any selection ranges; otherwise returns false.

This function was introduced in  Qt 4.2.
*/
func (this *QItemSelectionModel) HasSelection() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel12hasSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] QModelIndexList selectedIndexes() const

/*
Returns a list of all selected model item indexes. The list contains no duplicates, and is not sorted.

Note: Getter function for property selectedIndexes.
*/
func (this *QItemSelectionModel) SelectedIndexes() *QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel15selectedIndexesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:183
// index:0
// Public Visibility=Default Availability=Available
// [8] QModelIndexList selectedRows(int) const

/*
Returns the indexes in the given column for the rows where all columns are selected.

This function was introduced in  Qt 4.2.

See also selectedIndexes() and selectedColumns().
*/
func (this *QItemSelectionModel) SelectedRows(column int) *QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel12selectedRowsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:183
// index:0
// Public Visibility=Default Availability=Available
// [8] QModelIndexList selectedRows(int) const

/*
Returns the indexes in the given column for the rows where all columns are selected.

This function was introduced in  Qt 4.2.

See also selectedIndexes() and selectedColumns().
*/
func (this *QItemSelectionModel) SelectedRowsp() *QModelIndexList /*667*/ {
	// arg: 0, int=Int, =Invalid, , Invalid
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel12selectedRowsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:184
// index:0
// Public Visibility=Default Availability=Available
// [8] QModelIndexList selectedColumns(int) const

/*
Returns the indexes in the given row for columns where all rows are selected.

This function was introduced in  Qt 4.2.

See also selectedIndexes() and selectedRows().
*/
func (this *QItemSelectionModel) SelectedColumns(row int) *QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel15selectedColumnsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:184
// index:0
// Public Visibility=Default Availability=Available
// [8] QModelIndexList selectedColumns(int) const

/*
Returns the indexes in the given row for columns where all rows are selected.

This function was introduced in  Qt 4.2.

See also selectedIndexes() and selectedRows().
*/
func (this *QItemSelectionModel) SelectedColumnsp() *QModelIndexList /*667*/ {
	// arg: 0, int=Int, =Invalid, , Invalid
	row := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel15selectedColumnsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] const QItemSelection selection() const

/*
Returns the selection ranges stored in the selection model.
*/
func (this *QItemSelectionModel) Selection() *QItemSelection /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel9selectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:188
// index:0
// Public Visibility=Default Availability=Available
// [8] const QAbstractItemModel * model() const

/*
Returns the item model operated on by the selection model.

See also setModel().
*/
func (this *QItemSelectionModel) Model() *QAbstractItemModel /*777 const QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QItemSelectionModel5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:189
// index:1
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * model()

/*
Returns the item model operated on by the selection model.

See also setModel().
*/
func (this *QItemSelectionModel) Model1() *QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)

/*
Sets the model to model. The modelChanged() signal will be emitted.

This function was introduced in  Qt 5.5.

See also model() and modelChanged().
*/
func (this *QItemSelectionModel) SetModel(model QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:194
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setCurrentIndex(const QModelIndex &, QItemSelectionModel::SelectionFlags)

/*
Sets the model item index to be the current item, and emits currentChanged(). The current item is used for keyboard navigation and focus indication; it is independent of any selected items, although a selected item can also be the current item.

Depending on the specified command, the index can also become part of the current selection.

See also currentIndex() and select().
*/
func (this *QItemSelectionModel) SetCurrentIndex(index QModelIndex_ITF, command int) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel15setCurrentIndexERK11QModelIndex6QFlagsINS_13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:195
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void select(const QModelIndex &, QItemSelectionModel::SelectionFlags)

/*
Selects the model item index using the specified command, and emits selectionChanged().

See also QItemSelectionModel::SelectionFlags.
*/
func (this *QItemSelectionModel) Select(index QModelIndex_ITF, command int) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel6selectERK11QModelIndex6QFlagsINS_13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:196
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void select(const QItemSelection &, QItemSelectionModel::SelectionFlags)

/*
Selects the model item index using the specified command, and emits selectionChanged().

See also QItemSelectionModel::SelectionFlags.
*/
func (this *QItemSelectionModel) Select1(selection QItemSelection_ITF, command int) {
	var convArg0 unsafe.Pointer
	if selection != nil && selection.QItemSelection_PTR() != nil {
		convArg0 = selection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel6selectERK14QItemSelection6QFlagsINS_13SelectionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, command)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:197
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clear()

/*
Clears the selection model. Emits selectionChanged() and currentChanged().
*/
func (this *QItemSelectionModel) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:198
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void reset()

/*
Clears the selection model. Does not emit any signals.
*/
func (this *QItemSelectionModel) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearSelection()

/*
Clears the selection in the selection model. Emits selectionChanged().

This function was introduced in  Qt 4.2.
*/
func (this *QItemSelectionModel) ClearSelection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel14clearSelectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:201
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void clearCurrentIndex()

/*
Clears the current index. Emits currentChanged().
*/
func (this *QItemSelectionModel) ClearCurrentIndex() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel17clearCurrentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectionChanged(const QItemSelection &, const QItemSelection &)

/*
This signal is emitted whenever the selection changes. The change in the selection is represented as an item selection of deselected items and an item selection of selected items.

Note the that the current index changes independently from the selection. Also note that this signal will not be emitted when the item model is reset.

Note: Notifier signal for property selectedIndexes.

See also select() and currentChanged().
*/
func (this *QItemSelectionModel) SelectionChanged(selected QItemSelection_ITF, deselected QItemSelection_ITF) {
	var convArg0 unsafe.Pointer
	if selected != nil && selected.QItemSelection_PTR() != nil {
		convArg0 = selected.QItemSelection_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if deselected != nil && deselected.QItemSelection_PTR() != nil {
		convArg1 = deselected.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel16selectionChangedERK14QItemSelectionS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(const QModelIndex &, const QModelIndex &)

/*
This signal is emitted whenever the current item changes. The previous model item index is replaced by the current index as the selection's current item.

Note that this signal will not be emitted when the item model is reset.

See also currentIndex(), setCurrentIndex(), and selectionChanged().
*/
func (this *QItemSelectionModel) CurrentChanged(current QModelIndex_ITF, previous QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QModelIndex_PTR() != nil {
		convArg0 = current.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QModelIndex_PTR() != nil {
		convArg1 = previous.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel14currentChangedERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentRowChanged(const QModelIndex &, const QModelIndex &)

/*
This signal is emitted if the current item changes and its row is different to the row of the previous current item.

Note that this signal will not be emitted when the item model is reset.

See also currentChanged(), currentColumnChanged(), currentIndex(), and setCurrentIndex().
*/
func (this *QItemSelectionModel) CurrentRowChanged(current QModelIndex_ITF, previous QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QModelIndex_PTR() != nil {
		convArg0 = current.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QModelIndex_PTR() != nil {
		convArg1 = previous.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel17currentRowChangedERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:207
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentColumnChanged(const QModelIndex &, const QModelIndex &)

/*
This signal is emitted if the current item changes and its column is different to the column of the previous current item.

Note that this signal will not be emitted when the item model is reset.

See also currentChanged(), currentRowChanged(), currentIndex(), and setCurrentIndex().
*/
func (this *QItemSelectionModel) CurrentColumnChanged(current QModelIndex_ITF, previous QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if current != nil && current.QModelIndex_PTR() != nil {
		convArg0 = current.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if previous != nil && previous.QModelIndex_PTR() != nil {
		convArg1 = previous.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel20currentColumnChangedERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void modelChanged(QAbstractItemModel *)

/*
This signal is emitted when the model is successfully set with setModel().

This function was introduced in  Qt 5.5.

See also model() and setModel().
*/
func (this *QItemSelectionModel) ModelChanged(model QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel12modelChangedEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:212
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void emitSelectionChanged(const QItemSelection &, const QItemSelection &)

/*
Compares the two selections newSelection and oldSelection and emits selectionChanged() with the deselected and selected items.
*/
func (this *QItemSelectionModel) EmitSelectionChanged(newSelection QItemSelection_ITF, oldSelection QItemSelection_ITF) {
	var convArg0 unsafe.Pointer
	if newSelection != nil && newSelection.QItemSelection_PTR() != nil {
		convArg0 = newSelection.QItemSelection_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if oldSelection != nil && oldSelection.QItemSelection_PTR() != nil {
		convArg1 = oldSelection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QItemSelectionModel20emitSelectionChangedERK14QItemSelectionS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QItemSelectionModel__SelectionFlag = int

//
const QItemSelectionModel__NoUpdate QItemSelectionModel__SelectionFlag = 0

//
const QItemSelectionModel__Clear QItemSelectionModel__SelectionFlag = 1

//
const QItemSelectionModel__Select QItemSelectionModel__SelectionFlag = 2

//
const QItemSelectionModel__Deselect QItemSelectionModel__SelectionFlag = 4

//
const QItemSelectionModel__Toggle QItemSelectionModel__SelectionFlag = 8

//
const QItemSelectionModel__Current QItemSelectionModel__SelectionFlag = 16

//
const QItemSelectionModel__Rows QItemSelectionModel__SelectionFlag = 32

//
const QItemSelectionModel__Columns QItemSelectionModel__SelectionFlag = 64

//
const QItemSelectionModel__SelectCurrent QItemSelectionModel__SelectionFlag = 18

//
const QItemSelectionModel__ToggleCurrent QItemSelectionModel__SelectionFlag = 24

//
const QItemSelectionModel__ClearAndSelect QItemSelectionModel__SelectionFlag = 3

func (this *QItemSelectionModel) SelectionFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QItemSelectionModel_SelectionFlagItemName(val int) string {
	var nilthis *QItemSelectionModel
	return nilthis.SelectionFlagItemName(val)
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
}

//  keep block end
