//  header block begin
// /usr/include/qt/QtCore/qidentityproxymodel.h
// #include <qidentityproxymodel.h>
// #include <QtCore>
package qtcore

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
type QIdentityProxyModel struct {
	*QAbstractProxyModel
}

func (this *QIdentityProxyModel) GetCthis() unsafe.Pointer {
	return this.QAbstractProxyModel.GetCthis()
}
func NewQIdentityProxyModelFromPointer(cthis unsafe.Pointer) *QIdentityProxyModel {
	bcthis0 := NewQAbstractProxyModelFromPointer(cthis)
	return &QIdentityProxyModel{bcthis0}
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QIdentityProxyModel) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:57
// index:0
// Public
// void QIdentityProxyModel(class QObject *)
func NewQIdentityProxyModel(parent unsafe.Pointer) *QIdentityProxyModel {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQIdentityProxyModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:58
// index:0
// Public virtual
// void ~QIdentityProxyModel()
func DeleteQIdentityProxyModel(*QIdentityProxyModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:60
// index:0
// Public virtual
// int columnCount(const class QModelIndex &)
func (this *QIdentityProxyModel) ColumnCount(parent *QModelIndex) interface{} {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:61
// index:0
// Public virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) Index(row int, column int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:62
// index:0
// Public virtual
// QModelIndex mapFromSource(const class QModelIndex &)
func (this *QIdentityProxyModel) MapFromSource(sourceIndex *QModelIndex) interface{} {
	var convArg0 = sourceIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:63
// index:0
// Public virtual
// QModelIndex mapToSource(const class QModelIndex &)
func (this *QIdentityProxyModel) MapToSource(proxyIndex *QModelIndex) interface{} {
	var convArg0 = proxyIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:64
// index:0
// Public virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QIdentityProxyModel) Parent(child *QModelIndex) interface{} {
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:66
// index:0
// Public virtual
// int rowCount(const class QModelIndex &)
func (this *QIdentityProxyModel) RowCount(parent *QModelIndex) interface{} {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:67
// index:0
// Public virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QIdentityProxyModel) HeaderData(section int, orientation int, role int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:68
// index:0
// Public virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent *QModelIndex) interface{} {
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data, &action, &row, &column, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:69
// index:0
// Public virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) Sibling(row int, column int, idx *QModelIndex) interface{} {
	var convArg2 = idx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:71
// index:0
// Public virtual
// QItemSelection mapSelectionFromSource(const class QItemSelection &)
func (this *QIdentityProxyModel) MapSelectionFromSource(selection *QItemSelection) interface{} {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:72
// index:0
// Public virtual
// QItemSelection mapSelectionToSource(const class QItemSelection &)
func (this *QIdentityProxyModel) MapSelectionToSource(selection *QItemSelection) interface{} {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:74
// index:0
// Public virtual
// void setSourceModel(class QAbstractItemModel *)
func (this *QIdentityProxyModel) SetSourceModel(sourceModel unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sourceModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:76
// index:0
// Public virtual
// bool insertColumns(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) InsertColumns(column int, count int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:77
// index:0
// Public virtual
// bool insertRows(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) InsertRows(row int, count int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:78
// index:0
// Public virtual
// bool removeColumns(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) RemoveColumns(column int, count int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:79
// index:0
// Public virtual
// bool removeRows(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) RemoveRows(row int, count int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
