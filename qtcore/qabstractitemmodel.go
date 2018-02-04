package qtcore

// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// void resetInternalData()
func (this *QAbstractItemModel) InheritResetInternalData(f func()) {
	qtrt.SetAllInheritCallback(this, "resetInternalData", f)
}

// QModelIndex createIndex(int, int, void *)
func (this *QAbstractItemModel) InheritCreateIndex(f func(row int, column int, data unsafe.Pointer /*666*/) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "createIndex", f)
}

// bool decodeData(int, int, const class QModelIndex &, class QDataStream &)
func (this *QAbstractItemModel) InheritDecodeData(f func(row int, column int, parent *QModelIndex, stream *QDataStream) bool) {
	qtrt.SetAllInheritCallback(this, "decodeData", f)
}

// void beginInsertRows(const class QModelIndex &, int, int)
func (this *QAbstractItemModel) InheritBeginInsertRows(f func(parent *QModelIndex, first int, last int)) {
	qtrt.SetAllInheritCallback(this, "beginInsertRows", f)
}

// void endInsertRows()
func (this *QAbstractItemModel) InheritEndInsertRows(f func()) {
	qtrt.SetAllInheritCallback(this, "endInsertRows", f)
}

// void beginRemoveRows(const class QModelIndex &, int, int)
func (this *QAbstractItemModel) InheritBeginRemoveRows(f func(parent *QModelIndex, first int, last int)) {
	qtrt.SetAllInheritCallback(this, "beginRemoveRows", f)
}

// void endRemoveRows()
func (this *QAbstractItemModel) InheritEndRemoveRows(f func()) {
	qtrt.SetAllInheritCallback(this, "endRemoveRows", f)
}

// bool beginMoveRows(const class QModelIndex &, int, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) InheritBeginMoveRows(f func(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool) {
	qtrt.SetAllInheritCallback(this, "beginMoveRows", f)
}

// void endMoveRows()
func (this *QAbstractItemModel) InheritEndMoveRows(f func()) {
	qtrt.SetAllInheritCallback(this, "endMoveRows", f)
}

// void beginInsertColumns(const class QModelIndex &, int, int)
func (this *QAbstractItemModel) InheritBeginInsertColumns(f func(parent *QModelIndex, first int, last int)) {
	qtrt.SetAllInheritCallback(this, "beginInsertColumns", f)
}

// void endInsertColumns()
func (this *QAbstractItemModel) InheritEndInsertColumns(f func()) {
	qtrt.SetAllInheritCallback(this, "endInsertColumns", f)
}

// void beginRemoveColumns(const class QModelIndex &, int, int)
func (this *QAbstractItemModel) InheritBeginRemoveColumns(f func(parent *QModelIndex, first int, last int)) {
	qtrt.SetAllInheritCallback(this, "beginRemoveColumns", f)
}

// void endRemoveColumns()
func (this *QAbstractItemModel) InheritEndRemoveColumns(f func()) {
	qtrt.SetAllInheritCallback(this, "endRemoveColumns", f)
}

// bool beginMoveColumns(const class QModelIndex &, int, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) InheritBeginMoveColumns(f func(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool) {
	qtrt.SetAllInheritCallback(this, "beginMoveColumns", f)
}

// void endMoveColumns()
func (this *QAbstractItemModel) InheritEndMoveColumns(f func()) {
	qtrt.SetAllInheritCallback(this, "endMoveColumns", f)
}

// void beginResetModel()
func (this *QAbstractItemModel) InheritBeginResetModel(f func()) {
	qtrt.SetAllInheritCallback(this, "beginResetModel", f)
}

// void endResetModel()
func (this *QAbstractItemModel) InheritEndResetModel(f func()) {
	qtrt.SetAllInheritCallback(this, "endResetModel", f)
}

// void changePersistentIndex(const class QModelIndex &, const class QModelIndex &)
func (this *QAbstractItemModel) InheritChangePersistentIndex(f func(from *QModelIndex, to *QModelIndex)) {
	qtrt.SetAllInheritCallback(this, "changePersistentIndex", f)
}

type QAbstractItemModel struct {
	*QObject
}

func (this *QAbstractItemModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractItemModel) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQAbstractItemModelFromPointer(cthis unsafe.Pointer) *QAbstractItemModel {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractItemModel{bcthis0}
}
func (*QAbstractItemModel) NewFromPointer(cthis unsafe.Pointer) *QAbstractItemModel {
	return NewQAbstractItemModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAbstractItemModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemModel(QObject *)
func NewQAbstractItemModel(parent *QObject /*777 QObject **/) *QAbstractItemModel {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModelC1EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:175
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractItemModel()
func DeleteQAbstractItemModel(this *QAbstractItemModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:177
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasIndex(int, int, const QModelIndex &)
func (this *QAbstractItemModel) HasIndex(row int, column int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:178
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &)
func (this *QAbstractItemModel) Index(row int, column int, parent *QModelIndex) *QModelIndex /*123*/ {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:180
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex parent(const QModelIndex &)
func (this *QAbstractItemModel) Parent(child *QModelIndex) *QModelIndex /*123*/ {
	var convArg0 = child.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel6parentERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:182
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &)
func (this *QAbstractItemModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex /*123*/ {
	var convArg2 = idx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:183
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &)
func (this *QAbstractItemModel) RowCount(parent *QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:184
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &)
func (this *QAbstractItemModel) ColumnCount(parent *QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:185
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &)
func (this *QAbstractItemModel) HasChildren(parent *QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:187
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int)
func (this *QAbstractItemModel) Data(index *QModelIndex, role int) *QVariant /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:188
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)
func (this *QAbstractItemModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:190
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int)
func (this *QAbstractItemModel) HeaderData(section int, orientation int, role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:192
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)
func (this *QAbstractItemModel) SetHeaderData(section int, orientation int, value *QVariant, role int) bool {
	var convArg2 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:200
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canDropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QAbstractItemModel) CanDropMimeData(data *QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel15canDropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:202
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QAbstractItemModel) DropMimeData(data *QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:204
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions()
func (this *QAbstractItemModel) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:206
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDragActions()
func (this *QAbstractItemModel) SupportedDragActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel20supportedDragActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:212
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)
func (this *QAbstractItemModel) InsertRows(row int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:213
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)
func (this *QAbstractItemModel) InsertColumns(column int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:214
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)
func (this *QAbstractItemModel) RemoveRows(row int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:215
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)
func (this *QAbstractItemModel) RemoveColumns(column int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:216
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool moveRows(const QModelIndex &, int, int, const QModelIndex &, int)
func (this *QAbstractItemModel) MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg3 = destinationParent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceRow, count, convArg3, destinationChild)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:218
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool moveColumns(const QModelIndex &, int, int, const QModelIndex &, int)
func (this *QAbstractItemModel) MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg3 = destinationParent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceColumn, count, convArg3, destinationChild)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool insertRow(int, const QModelIndex &)
func (this *QAbstractItemModel) InsertRow(row int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool insertColumn(int, const QModelIndex &)
func (this *QAbstractItemModel) InsertColumn(column int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:223
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool removeRow(int, const QModelIndex &)
func (this *QAbstractItemModel) RemoveRow(row int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:224
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool removeColumn(int, const QModelIndex &)
func (this *QAbstractItemModel) RemoveColumn(column int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:225
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool moveRow(const QModelIndex &, int, const QModelIndex &, int)
func (this *QAbstractItemModel) MoveRow(sourceParent *QModelIndex, sourceRow int, destinationParent *QModelIndex, destinationChild int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg2 = destinationParent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceRow, convArg2, destinationChild)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:227
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool moveColumn(const QModelIndex &, int, const QModelIndex &, int)
func (this *QAbstractItemModel) MoveColumn(sourceParent *QModelIndex, sourceColumn int, destinationParent *QModelIndex, destinationChild int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg2 = destinationParent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceColumn, convArg2, destinationChild)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:230
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fetchMore(const QModelIndex &)
func (this *QAbstractItemModel) FetchMore(parent *QModelIndex) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel9fetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:231
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canFetchMore(const QModelIndex &)
func (this *QAbstractItemModel) CanFetchMore(parent *QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:232
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &)
func (this *QAbstractItemModel) Flags(index *QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:233
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)
func (this *QAbstractItemModel) Sort(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:234
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex buddy(const QModelIndex &)
func (this *QAbstractItemModel) Buddy(index *QModelIndex) *QModelIndex /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel5buddyERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:239
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize span(const QModelIndex &)
func (this *QAbstractItemModel) Span(index *QModelIndex) *QSize /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel4spanERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void headerDataChanged(Qt::Orientation, int, int)
func (this *QAbstractItemModel) HeaderDataChanged(orientation int, first int, last int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel17headerDataChangedEN2Qt11OrientationEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:281
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool submit()
func (this *QAbstractItemModel) Submit() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel6submitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:282
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void revert()
func (this *QAbstractItemModel) Revert() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel6revertEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:286
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void resetInternalData()
func (this *QAbstractItemModel) ResetInternalData() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel17resetInternalDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:291
// index:0
// Protected inline Visibility=Default Availability=Available
// [24] QModelIndex createIndex(int, int, void *)
func (this *QAbstractItemModel) CreateIndex(row int, column int, data unsafe.Pointer /*666*/) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11createIndexEiiPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, data)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:292
// index:1
// Protected inline Visibility=Default Availability=Available
// [24] QModelIndex createIndex(int, int, quintptr)
func (this *QAbstractItemModel) CreateIndex_1(row int, column int, id uint64) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11createIndexEiiy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, id)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:295
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool decodeData(int, int, const QModelIndex &, QDataStream &)
func (this *QAbstractItemModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	var convArg2 = parent.GetCthis()
	var convArg3 = stream.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel10decodeDataEiiRK11QModelIndexR11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:297
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void beginInsertRows(const QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginInsertRowsERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:298
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endInsertRows()
func (this *QAbstractItemModel) EndInsertRows() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13endInsertRowsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:300
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void beginRemoveRows(const QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginRemoveRowsERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:301
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endRemoveRows()
func (this *QAbstractItemModel) EndRemoveRows() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13endRemoveRowsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:303
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool beginMoveRows(const QModelIndex &, int, int, const QModelIndex &, int)
func (this *QAbstractItemModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg3 = destinationParent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13beginMoveRowsERK11QModelIndexiiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceFirst, sourceLast, convArg3, destinationRow)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:304
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endMoveRows()
func (this *QAbstractItemModel) EndMoveRows() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel11endMoveRowsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:306
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void beginInsertColumns(const QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel18beginInsertColumnsERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:307
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endInsertColumns()
func (this *QAbstractItemModel) EndInsertColumns() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel16endInsertColumnsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:309
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void beginRemoveColumns(const QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel18beginRemoveColumnsERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:310
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endRemoveColumns()
func (this *QAbstractItemModel) EndRemoveColumns() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel16endRemoveColumnsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:312
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool beginMoveColumns(const QModelIndex &, int, int, const QModelIndex &, int)
func (this *QAbstractItemModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg3 = destinationParent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel16beginMoveColumnsERK11QModelIndexiiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceFirst, sourceLast, convArg3, destinationColumn)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:313
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endMoveColumns()
func (this *QAbstractItemModel) EndMoveColumns() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel14endMoveColumnsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:324
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void beginResetModel()
func (this *QAbstractItemModel) BeginResetModel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginResetModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:325
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endResetModel()
func (this *QAbstractItemModel) EndResetModel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13endResetModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:327
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void changePersistentIndex(const QModelIndex &, const QModelIndex &)
func (this *QAbstractItemModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	var convArg0 = from.GetCthis()
	var convArg1 = to.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel21changePersistentIndexERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

type QAbstractItemModel__LayoutChangeHint = int

const QAbstractItemModel__NoLayoutChangeHint QAbstractItemModel__LayoutChangeHint = 0
const QAbstractItemModel__VerticalSortHint QAbstractItemModel__LayoutChangeHint = 1
const QAbstractItemModel__HorizontalSortHint QAbstractItemModel__LayoutChangeHint = 2

//  body block end
