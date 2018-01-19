//  header block begin
// /usr/include/qt/QtCore/qidentityproxymodel.h
// #include <qidentityproxymodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QIdentityProxyModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:57
// index:0
// void QIdentityProxyModel(class QObject *)
func NewQIdentityProxyModel(parent unsafe.Pointer) *QIdentityProxyModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QIdentityProxyModel{cthis}
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:58
// index:0
// virtual
// void ~QIdentityProxyModel()
func DeleteQIdentityProxyModel(*QIdentityProxyModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:60
// index:0
// virtual
// int columnCount(const class QModelIndex &)
func (this *QIdentityProxyModel) ColumnCount(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:61
// index:0
// virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) Index(row int, column int, parent unsafe.Pointer) {
	// 0: (, int row, int column, const QModelIndex & parent), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:62
// index:0
// virtual
// QModelIndex mapFromSource(const class QModelIndex &)
func (this *QIdentityProxyModel) MapFromSource(sourceIndex unsafe.Pointer) {
	// 0: (, const QModelIndex & sourceIndex), (sourceIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, sourceIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:63
// index:0
// virtual
// QModelIndex mapToSource(const class QModelIndex &)
func (this *QIdentityProxyModel) MapToSource(proxyIndex unsafe.Pointer) {
	// 0: (, const QModelIndex & proxyIndex), (proxyIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, proxyIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:64
// index:0
// virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QIdentityProxyModel) Parent(child unsafe.Pointer) {
	// 0: (, const QModelIndex & child), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:66
// index:0
// virtual
// int rowCount(const class QModelIndex &)
func (this *QIdentityProxyModel) RowCount(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:67
// index:0
// virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QIdentityProxyModel) HeaderData(section int, orientation int, role int) {
	// 0: (, int section, Qt::Orientation orientation, int role), (&section, &orientation, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_VOID, this.cthis, &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:68
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:69
// index:0
// virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) Sibling(row int, column int, idx unsafe.Pointer) {
	// 0: (, int row, int column, const QModelIndex & idx), (&row, &column, idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:71
// index:0
// virtual
// QItemSelection mapSelectionFromSource(const class QItemSelection &)
func (this *QIdentityProxyModel) MapSelectionFromSource(selection unsafe.Pointer) {
	// 0: (, const QItemSelection & selection), (selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.cthis, selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:72
// index:0
// virtual
// QItemSelection mapSelectionToSource(const class QItemSelection &)
func (this *QIdentityProxyModel) MapSelectionToSource(selection unsafe.Pointer) {
	// 0: (, const QItemSelection & selection), (selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.cthis, selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:74
// index:0
// virtual
// void setSourceModel(class QAbstractItemModel *)
func (this *QIdentityProxyModel) SetSourceModel(sourceModel unsafe.Pointer) {
	// 0: (, QAbstractItemModel * sourceModel), (sourceModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.cthis, sourceModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:76
// index:0
// virtual
// bool insertColumns(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) InsertColumns(column int, count int, parent unsafe.Pointer) {
	// 0: (, int column, int count, const QModelIndex & parent), (&column, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:77
// index:0
// virtual
// bool insertRows(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) InsertRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, int row, int count, const QModelIndex & parent), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:78
// index:0
// virtual
// bool removeColumns(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) RemoveColumns(column int, count int, parent unsafe.Pointer) {
	// 0: (, int column, int count, const QModelIndex & parent), (&column, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:79
// index:0
// virtual
// bool removeRows(int, int, const class QModelIndex &)
func (this *QIdentityProxyModel) RemoveRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, int row, int count, const QModelIndex & parent), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

//  body block end
