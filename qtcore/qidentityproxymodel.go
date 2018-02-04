package qtcore

// /usr/include/qt/QtCore/qidentityproxymodel.h
// #include <qidentityproxymodel.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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

type QIdentityProxyModel struct {
	*QAbstractProxyModel
}

func (this *QIdentityProxyModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractProxyModel.GetCthis()
	}
}
func (this *QIdentityProxyModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractProxyModel = NewQAbstractProxyModelFromPointer(cthis)
}
func NewQIdentityProxyModelFromPointer(cthis unsafe.Pointer) *QIdentityProxyModel {
	bcthis0 := NewQAbstractProxyModelFromPointer(cthis)
	return &QIdentityProxyModel{bcthis0}
}
func (*QIdentityProxyModel) NewFromPointer(cthis unsafe.Pointer) *QIdentityProxyModel {
	return NewQIdentityProxyModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QIdentityProxyModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIdentityProxyModel(QObject *)
func NewQIdentityProxyModel(parent *QObject /*777 QObject **/) *QIdentityProxyModel {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQIdentityProxyModelFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QIdentityProxyModel()
func DeleteQIdentityProxyModel(this *QIdentityProxyModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &)
func (this *QIdentityProxyModel) ColumnCount(parent *QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &)
func (this *QIdentityProxyModel) Index(row int, column int, parent *QModelIndex) *QModelIndex /*123*/ {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex mapFromSource(const QModelIndex &)
func (this *QIdentityProxyModel) MapFromSource(sourceIndex *QModelIndex) *QModelIndex /*123*/ {
	var convArg0 = sourceIndex.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex mapToSource(const QModelIndex &)
func (this *QIdentityProxyModel) MapToSource(proxyIndex *QModelIndex) *QModelIndex /*123*/ {
	var convArg0 = proxyIndex.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex parent(const QModelIndex &)
func (this *QIdentityProxyModel) Parent(child *QModelIndex) *QModelIndex /*123*/ {
	var convArg0 = child.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel6parentERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &)
func (this *QIdentityProxyModel) RowCount(parent *QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int)
func (this *QIdentityProxyModel) HeaderData(section int, orientation int, role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QIdentityProxyModel) DropMimeData(data *QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &)
func (this *QIdentityProxyModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex /*123*/ {
	var convArg2 = idx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QItemSelection mapSelectionFromSource(const QItemSelection &)
func (this *QIdentityProxyModel) MapSelectionFromSource(selection *QItemSelection) *QItemSelection /*123*/ {
	var convArg0 = selection.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QItemSelection mapSelectionToSource(const QItemSelection &)
func (this *QIdentityProxyModel) MapSelectionToSource(selection *QItemSelection) *QItemSelection /*123*/ {
	var convArg0 = selection.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSourceModel(QAbstractItemModel *)
func (this *QIdentityProxyModel) SetSourceModel(sourceModel *QAbstractItemModel /*777 QAbstractItemModel **/) {
	var convArg0 = sourceModel.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)
func (this *QIdentityProxyModel) InsertColumns(column int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)
func (this *QIdentityProxyModel) InsertRows(row int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)
func (this *QIdentityProxyModel) RemoveColumns(column int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)
func (this *QIdentityProxyModel) RemoveRows(row int, count int, parent *QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
