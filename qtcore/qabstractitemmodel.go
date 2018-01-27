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
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractItemModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:174
// index:0
// Public
// void QAbstractItemModel(QObject *)
func NewQAbstractItemModel(parent *QObject /*777 QObject **/) *QAbstractItemModel {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModelC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractItemModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:175
// index:0
// Public virtual
// void ~QAbstractItemModel()
func DeleteQAbstractItemModel(*QAbstractItemModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:177
// index:0
// Public
// bool hasIndex(int, int, const QModelIndex &)
func (this *QAbstractItemModel) HasIndex(row int, column int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:178
// index:0
// Public pure virtual
// QModelIndex index(int, int, const QModelIndex &)
func (this *QAbstractItemModel) Index(row int, column int, parent *QModelIndex) *QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:180
// index:0
// Public pure virtual
// QModelIndex parent(const QModelIndex &)
func (this *QAbstractItemModel) Parent(child *QModelIndex) *QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:182
// index:0
// Public virtual
// QModelIndex sibling(int, int, const QModelIndex &)
func (this *QAbstractItemModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg2 = idx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:183
// index:0
// Public pure virtual
// int rowCount(const QModelIndex &)
func (this *QAbstractItemModel) RowCount(parent *QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:184
// index:0
// Public pure virtual
// int columnCount(const QModelIndex &)
func (this *QAbstractItemModel) ColumnCount(parent *QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:185
// index:0
// Public virtual
// bool hasChildren(const QModelIndex &)
func (this *QAbstractItemModel) HasChildren(parent *QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:187
// index:0
// Public pure virtual
// QVariant data(const QModelIndex &, int)
func (this *QAbstractItemModel) Data(index *QModelIndex, role int) *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:188
// index:0
// Public virtual
// bool setData(const QModelIndex &, const QVariant &, int)
func (this *QAbstractItemModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:190
// index:0
// Public virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QAbstractItemModel) HeaderData(section int, orientation int, role int) *QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), section, orientation, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:192
// index:0
// Public virtual
// bool setHeaderData(int, Qt::Orientation, const QVariant &, int)
func (this *QAbstractItemModel) SetHeaderData(section int, orientation int, value *QVariant, role int) bool {
	var convArg2 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:200
// index:0
// Public virtual
// bool canDropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QAbstractItemModel) CanDropMimeData(data *QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel15canDropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:202
// index:0
// Public virtual
// bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QAbstractItemModel) DropMimeData(data *QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:204
// index:0
// Public virtual
// Qt::DropActions supportedDropActions()
func (this *QAbstractItemModel) SupportedDropActions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:206
// index:0
// Public virtual
// Qt::DropActions supportedDragActions()
func (this *QAbstractItemModel) SupportedDragActions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel20supportedDragActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:212
// index:0
// Public virtual
// bool insertRows(int, int, const QModelIndex &)
func (this *QAbstractItemModel) InsertRows(row int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:213
// index:0
// Public virtual
// bool insertColumns(int, int, const QModelIndex &)
func (this *QAbstractItemModel) InsertColumns(column int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:214
// index:0
// Public virtual
// bool removeRows(int, int, const QModelIndex &)
func (this *QAbstractItemModel) RemoveRows(row int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:215
// index:0
// Public virtual
// bool removeColumns(int, int, const QModelIndex &)
func (this *QAbstractItemModel) RemoveColumns(column int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:216
// index:0
// Public virtual
// bool moveRows(const QModelIndex &, int, int, const QModelIndex &, int)
func (this *QAbstractItemModel) MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg3 = destinationParent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceRow, count, convArg3, destinationChild)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:218
// index:0
// Public virtual
// bool moveColumns(const QModelIndex &, int, int, const QModelIndex &, int)
func (this *QAbstractItemModel) MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg3 = destinationParent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceColumn, count, convArg3, destinationChild)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:221
// index:0
// Public inline
// bool insertRow(int, const QModelIndex &)
func (this *QAbstractItemModel) InsertRow(row int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:222
// index:0
// Public inline
// bool insertColumn(int, const QModelIndex &)
func (this *QAbstractItemModel) InsertColumn(column int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:223
// index:0
// Public inline
// bool removeRow(int, const QModelIndex &)
func (this *QAbstractItemModel) RemoveRow(row int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:224
// index:0
// Public inline
// bool removeColumn(int, const QModelIndex &)
func (this *QAbstractItemModel) RemoveColumn(column int, parent *QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:225
// index:0
// Public inline
// bool moveRow(const QModelIndex &, int, const QModelIndex &, int)
func (this *QAbstractItemModel) MoveRow(sourceParent *QModelIndex, sourceRow int, destinationParent *QModelIndex, destinationChild int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg2 = destinationParent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceRow, convArg2, destinationChild)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:227
// index:0
// Public inline
// bool moveColumn(const QModelIndex &, int, const QModelIndex &, int)
func (this *QAbstractItemModel) MoveColumn(sourceParent *QModelIndex, sourceColumn int, destinationParent *QModelIndex, destinationChild int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg2 = destinationParent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceColumn, convArg2, destinationChild)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:230
// index:0
// Public virtual
// void fetchMore(const QModelIndex &)
func (this *QAbstractItemModel) FetchMore(parent *QModelIndex) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel9fetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:231
// index:0
// Public virtual
// bool canFetchMore(const QModelIndex &)
func (this *QAbstractItemModel) CanFetchMore(parent *QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:232
// index:0
// Public virtual
// Qt::ItemFlags flags(const QModelIndex &)
func (this *QAbstractItemModel) Flags(index *QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:233
// index:0
// Public virtual
// void sort(int, Qt::SortOrder)
func (this *QAbstractItemModel) Sort(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:234
// index:0
// Public virtual
// QModelIndex buddy(const QModelIndex &)
func (this *QAbstractItemModel) Buddy(index *QModelIndex) *QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel5buddyERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:239
// index:0
// Public virtual
// QSize span(const QModelIndex &)
func (this *QAbstractItemModel) Span(index *QModelIndex) *QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel4spanERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:255
// index:0
// Public
// void headerDataChanged(Qt::Orientation, int, int)
func (this *QAbstractItemModel) HeaderDataChanged(orientation int, first int, last int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel17headerDataChangedEN2Qt11OrientationEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:281
// index:0
// Public virtual
// bool submit()
func (this *QAbstractItemModel) Submit() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel6submitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:282
// index:0
// Public virtual
// void revert()
func (this *QAbstractItemModel) Revert() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel6revertEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:286
// index:0
// Protected
// void resetInternalData()
func (this *QAbstractItemModel) ResetInternalData() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel17resetInternalDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:291
// index:0
// Protected inline
// QModelIndex createIndex(int, int, void *)
func (this *QAbstractItemModel) CreateIndex(row int, column int, data unsafe.Pointer /*666*/) *QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel11createIndexEiiPv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), row, column, data)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:292
// index:1
// Protected inline
// QModelIndex createIndex(int, int, quintptr)
func (this *QAbstractItemModel) CreateIndex_1(row int, column int, id uint64) *QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel11createIndexEiiy", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), row, column, id)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:295
// index:0
// Protected
// bool decodeData(int, int, const QModelIndex &, QDataStream &)
func (this *QAbstractItemModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	var convArg2 = parent.GetCthis()
	var convArg3 = stream.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel10decodeDataEiiRK11QModelIndexR11QDataStream", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:297
// index:0
// Protected
// void beginInsertRows(const QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginInsertRowsERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:298
// index:0
// Protected
// void endInsertRows()
func (this *QAbstractItemModel) EndInsertRows() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13endInsertRowsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:300
// index:0
// Protected
// void beginRemoveRows(const QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginRemoveRowsERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:301
// index:0
// Protected
// void endRemoveRows()
func (this *QAbstractItemModel) EndRemoveRows() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13endRemoveRowsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:303
// index:0
// Protected
// bool beginMoveRows(const QModelIndex &, int, int, const QModelIndex &, int)
func (this *QAbstractItemModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg3 = destinationParent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13beginMoveRowsERK11QModelIndexiiS2_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceFirst, sourceLast, convArg3, destinationRow)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:304
// index:0
// Protected
// void endMoveRows()
func (this *QAbstractItemModel) EndMoveRows() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel11endMoveRowsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:306
// index:0
// Protected
// void beginInsertColumns(const QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel18beginInsertColumnsERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:307
// index:0
// Protected
// void endInsertColumns()
func (this *QAbstractItemModel) EndInsertColumns() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel16endInsertColumnsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:309
// index:0
// Protected
// void beginRemoveColumns(const QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel18beginRemoveColumnsERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:310
// index:0
// Protected
// void endRemoveColumns()
func (this *QAbstractItemModel) EndRemoveColumns() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel16endRemoveColumnsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:312
// index:0
// Protected
// bool beginMoveColumns(const QModelIndex &, int, int, const QModelIndex &, int)
func (this *QAbstractItemModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	var convArg0 = sourceParent.GetCthis()
	var convArg3 = destinationParent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel16beginMoveColumnsERK11QModelIndexiiS2_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceFirst, sourceLast, convArg3, destinationColumn)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:313
// index:0
// Protected
// void endMoveColumns()
func (this *QAbstractItemModel) EndMoveColumns() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel14endMoveColumnsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:324
// index:0
// Protected
// void beginResetModel()
func (this *QAbstractItemModel) BeginResetModel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginResetModelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:325
// index:0
// Protected
// void endResetModel()
func (this *QAbstractItemModel) EndResetModel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13endResetModelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:327
// index:0
// Protected
// void changePersistentIndex(const QModelIndex &, const QModelIndex &)
func (this *QAbstractItemModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	var convArg0 = from.GetCthis()
	var convArg1 = to.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel21changePersistentIndexERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

type QAbstractItemModel__LayoutChangeHint = int

const QAbstractItemModel__NoLayoutChangeHint QAbstractItemModel__LayoutChangeHint = 0
const QAbstractItemModel__VerticalSortHint QAbstractItemModel__LayoutChangeHint = 1
const QAbstractItemModel__HorizontalSortHint QAbstractItemModel__LayoutChangeHint = 2

//  body block end
