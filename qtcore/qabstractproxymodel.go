package qtcore

// /usr/include/qt/QtCore/qabstractproxymodel.h
// #include <qabstractproxymodel.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
func (this *QAbstractProxyModel) InheritResetInternalData(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "resetInternalData", f)
}

type QAbstractProxyModel struct {
	*QAbstractItemModel
}

func (this *QAbstractProxyModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemModel.GetCthis()
	}
}
func (this *QAbstractProxyModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemModel = NewQAbstractItemModelFromPointer(cthis)
}
func NewQAbstractProxyModelFromPointer(cthis unsafe.Pointer) *QAbstractProxyModel {
	bcthis0 := NewQAbstractItemModelFromPointer(cthis)
	return &QAbstractProxyModel{bcthis0}
}
func (*QAbstractProxyModel) NewFromPointer(cthis unsafe.Pointer) *QAbstractProxyModel {
	return NewQAbstractProxyModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAbstractProxyModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractProxyModel(QObject *)
func NewQAbstractProxyModel(parent *QObject /*777 QObject **/) *QAbstractProxyModel {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractProxyModelFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractProxyModel()
func DeleteQAbstractProxyModel(this *QAbstractProxyModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSourceModel(QAbstractItemModel *)
func (this *QAbstractProxyModel) SetSourceModel(sourceModel *QAbstractItemModel /*777 QAbstractItemModel **/) {
	var convArg0 = sourceModel.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel14setSourceModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * sourceModel()
func (this *QAbstractProxyModel) SourceModel() *QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11sourceModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex mapToSource(const QModelIndex &)
func (this *QAbstractProxyModel) MapToSource(proxyIndex *QModelIndex) *QModelIndex /*123*/ {
	var convArg0 = proxyIndex.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11mapToSourceERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex mapFromSource(const QModelIndex &)
func (this *QAbstractProxyModel) MapFromSource(sourceIndex *QModelIndex) *QModelIndex /*123*/ {
	var convArg0 = sourceIndex.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel13mapFromSourceERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QItemSelection mapSelectionToSource(const QItemSelection &)
func (this *QAbstractProxyModel) MapSelectionToSource(selection *QItemSelection) *QItemSelection /*123*/ {
	var convArg0 = selection.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20mapSelectionToSourceERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QItemSelection mapSelectionFromSource(const QItemSelection &)
func (this *QAbstractProxyModel) MapSelectionFromSource(selection *QItemSelection) *QItemSelection /*123*/ {
	var convArg0 = selection.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel22mapSelectionFromSourceERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool submit()
func (this *QAbstractProxyModel) Submit() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel6submitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void revert()
func (this *QAbstractProxyModel) Revert() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel6revertEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int)
func (this *QAbstractProxyModel) Data(proxyIndex *QModelIndex, role int) *QVariant /*123*/ {
	var convArg0 = proxyIndex.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int)
func (this *QAbstractProxyModel) HeaderData(section int, orientation int, role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &)
func (this *QAbstractProxyModel) Flags(index *QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)
func (this *QAbstractProxyModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)
func (this *QAbstractProxyModel) SetHeaderData(section int, orientation int, value *QVariant, role int) bool {
	var convArg2 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex buddy(const QModelIndex &)
func (this *QAbstractProxyModel) Buddy(index *QModelIndex) *QModelIndex /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel5buddyERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canFetchMore(const QModelIndex &)
func (this *QAbstractProxyModel) CanFetchMore(parent *QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel12canFetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fetchMore(const QModelIndex &)
func (this *QAbstractProxyModel) FetchMore(parent *QModelIndex) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel9fetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)
func (this *QAbstractProxyModel) Sort(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize span(const QModelIndex &)
func (this *QAbstractProxyModel) Span(index *QModelIndex) *QSize /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel4spanERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &)
func (this *QAbstractProxyModel) HasChildren(parent *QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &)
func (this *QAbstractProxyModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex /*123*/ {
	var convArg2 = idx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canDropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QAbstractProxyModel) CanDropMimeData(data *QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel15canDropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QAbstractProxyModel) DropMimeData(data *QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDragActions()
func (this *QAbstractProxyModel) SupportedDragActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20supportedDragActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions()
func (this *QAbstractProxyModel) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:104
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void resetInternalData()
func (this *QAbstractProxyModel) ResetInternalData() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel17resetInternalDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
