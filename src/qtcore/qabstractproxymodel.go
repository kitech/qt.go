//  header block begin
// /usr/include/qt/QtCore/qabstractproxymodel.h
// #include <qabstractproxymodel.h>
// #include <QtCore>
package qtcore

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
type QAbstractProxyModel struct {
	*QAbstractItemModel
}

func (this *QAbstractProxyModel) GetCthis() unsafe.Pointer {
	return this.QAbstractItemModel.GetCthis()
}
func NewQAbstractProxyModelFromPointer(cthis unsafe.Pointer) *QAbstractProxyModel {
	bcthis0 := NewQAbstractItemModelFromPointer(cthis)
	return &QAbstractProxyModel{bcthis0}
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractProxyModel) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:59
// index:0
// Public
// void QAbstractProxyModel(class QObject *)
func NewQAbstractProxyModel(parent unsafe.Pointer) *QAbstractProxyModel {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModelC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractProxyModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:60
// index:0
// Public virtual
// void ~QAbstractProxyModel()
func DeleteQAbstractProxyModel(*QAbstractProxyModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:62
// index:0
// Public virtual
// void setSourceModel(class QAbstractItemModel *)
func (this *QAbstractProxyModel) SetSourceModel(sourceModel unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel14setSourceModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sourceModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:63
// index:0
// Public
// QAbstractItemModel * sourceModel()
func (this *QAbstractProxyModel) SourceModel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11sourceModelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:65
// index:0
// Public pure virtual
// QModelIndex mapToSource(const class QModelIndex &)
func (this *QAbstractProxyModel) MapToSource(proxyIndex *QModelIndex) interface{} {
	var convArg0 = proxyIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11mapToSourceERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:66
// index:0
// Public pure virtual
// QModelIndex mapFromSource(const class QModelIndex &)
func (this *QAbstractProxyModel) MapFromSource(sourceIndex *QModelIndex) interface{} {
	var convArg0 = sourceIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel13mapFromSourceERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:68
// index:0
// Public virtual
// QItemSelection mapSelectionToSource(const class QItemSelection &)
func (this *QAbstractProxyModel) MapSelectionToSource(selection *QItemSelection) interface{} {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20mapSelectionToSourceERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:69
// index:0
// Public virtual
// QItemSelection mapSelectionFromSource(const class QItemSelection &)
func (this *QAbstractProxyModel) MapSelectionFromSource(selection *QItemSelection) interface{} {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel22mapSelectionFromSourceERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:71
// index:0
// Public virtual
// bool submit()
func (this *QAbstractProxyModel) Submit() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel6submitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:72
// index:0
// Public virtual
// void revert()
func (this *QAbstractProxyModel) Revert() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel6revertEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:74
// index:0
// Public virtual
// QVariant data(const class QModelIndex &, int)
func (this *QAbstractProxyModel) Data(proxyIndex *QModelIndex, role int) interface{} {
	var convArg0 = proxyIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:75
// index:0
// Public virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QAbstractProxyModel) HeaderData(section int, orientation int, role int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:76
// index:0
// Public virtual
// QMap<int, QVariant> itemData(const class QModelIndex &)
func (this *QAbstractProxyModel) ItemData(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel8itemDataERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:77
// index:0
// Public virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QAbstractProxyModel) Flags(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:79
// index:0
// Public virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QAbstractProxyModel) SetData(index *QModelIndex, value *QVariant, role int) interface{} {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:81
// index:0
// Public virtual
// bool setHeaderData(int, Qt::Orientation, const class QVariant &, int)
func (this *QAbstractProxyModel) SetHeaderData(section int, orientation int, value *QVariant, role int) interface{} {
	var convArg2 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &section, &orientation, convArg2, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:83
// index:0
// Public virtual
// QModelIndex buddy(const class QModelIndex &)
func (this *QAbstractProxyModel) Buddy(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel5buddyERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:84
// index:0
// Public virtual
// bool canFetchMore(const class QModelIndex &)
func (this *QAbstractProxyModel) CanFetchMore(parent *QModelIndex) interface{} {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel12canFetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:85
// index:0
// Public virtual
// void fetchMore(const class QModelIndex &)
func (this *QAbstractProxyModel) FetchMore(parent *QModelIndex) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel9fetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:86
// index:0
// Public virtual
// void sort(int, Qt::SortOrder)
func (this *QAbstractProxyModel) Sort(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:87
// index:0
// Public virtual
// QSize span(const class QModelIndex &)
func (this *QAbstractProxyModel) Span(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel4spanERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:88
// index:0
// Public virtual
// bool hasChildren(const class QModelIndex &)
func (this *QAbstractProxyModel) HasChildren(parent *QModelIndex) interface{} {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:89
// index:0
// Public virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QAbstractProxyModel) Sibling(row int, column int, idx *QModelIndex) interface{} {
	var convArg2 = idx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:92
// index:0
// Public virtual
// bool canDropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QAbstractProxyModel) CanDropMimeData(data unsafe.Pointer, action int, row int, column int, parent *QModelIndex) interface{} {
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel15canDropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data, &action, &row, &column, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:94
// index:0
// Public virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QAbstractProxyModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent *QModelIndex) interface{} {
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data, &action, &row, &column, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:96
// index:0
// Public virtual
// QStringList mimeTypes()
func (this *QAbstractProxyModel) MimeTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel9mimeTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:97
// index:0
// Public virtual
// Qt::DropActions supportedDragActions()
func (this *QAbstractProxyModel) SupportedDragActions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20supportedDragActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:98
// index:0
// Public virtual
// Qt::DropActions supportedDropActions()
func (this *QAbstractProxyModel) SupportedDropActions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:104
// index:0
// Protected
// void resetInternalData()
func (this *QAbstractProxyModel) ResetInternalData() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel17resetInternalDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
