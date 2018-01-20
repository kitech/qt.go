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

// /usr/include/qt/QtCore/qabstractproxymodel.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractProxyModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:59
// index:0
// void QAbstractProxyModel(class QObject *)
func NewQAbstractProxyModel(parent unsafe.Pointer) *QAbstractProxyModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModelC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractProxyModelFromPointer(cthis)
	return gothis
}
func NewQAbstractProxyModelFromPointer(cthis unsafe.Pointer) *QAbstractProxyModel {
	bcthis0 := NewQAbstractItemModelFromPointer(cthis)
	return &QAbstractProxyModel{bcthis0}
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:107
// index:1
// void QAbstractProxyModel(class QAbstractProxyModelPrivate &, class QObject *)
func NewQAbstractProxyModel_1(arg0 unsafe.Pointer, parent unsafe.Pointer) *QAbstractProxyModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModelC1ER26QAbstractProxyModelPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractProxyModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:60
// index:0
// virtual
// void ~QAbstractProxyModel()
func DeleteQAbstractProxyModel(*QAbstractProxyModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:62
// index:0
// virtual
// void setSourceModel(class QAbstractItemModel *)
func (this *QAbstractProxyModel) SetSourceModel(sourceModel unsafe.Pointer) {
	// 0: (, sourceModel QAbstractItemModel *), (sourceModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel14setSourceModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:63
// index:0
// QAbstractItemModel * sourceModel()
func (this *QAbstractProxyModel) SourceModel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11sourceModelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:65
// index:0
// pure virtual
// QModelIndex mapToSource(const class QModelIndex &)
func (this *QAbstractProxyModel) MapToSource(proxyIndex unsafe.Pointer) {
	// 0: (, proxyIndex const QModelIndex &), (proxyIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11mapToSourceERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), proxyIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:66
// index:0
// pure virtual
// QModelIndex mapFromSource(const class QModelIndex &)
func (this *QAbstractProxyModel) MapFromSource(sourceIndex unsafe.Pointer) {
	// 0: (, sourceIndex const QModelIndex &), (sourceIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel13mapFromSourceERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:68
// index:0
// virtual
// QItemSelection mapSelectionToSource(const class QItemSelection &)
func (this *QAbstractProxyModel) MapSelectionToSource(selection unsafe.Pointer) {
	// 0: (, selection const QItemSelection &), (selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20mapSelectionToSourceERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:69
// index:0
// virtual
// QItemSelection mapSelectionFromSource(const class QItemSelection &)
func (this *QAbstractProxyModel) MapSelectionFromSource(selection unsafe.Pointer) {
	// 0: (, selection const QItemSelection &), (selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel22mapSelectionFromSourceERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:71
// index:0
// virtual
// bool submit()
func (this *QAbstractProxyModel) Submit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel6submitEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:72
// index:0
// virtual
// void revert()
func (this *QAbstractProxyModel) Revert() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel6revertEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:74
// index:0
// virtual
// QVariant data(const class QModelIndex &, int)
func (this *QAbstractProxyModel) Data(proxyIndex unsafe.Pointer, role int) {
	// 0: (, proxyIndex const QModelIndex &, role int), (proxyIndex, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), proxyIndex, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:75
// index:0
// virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QAbstractProxyModel) HeaderData(section int, orientation int, role int) {
	// 0: (, section int, orientation Qt::Orientation, role int), (&section, &orientation, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:76
// index:0
// virtual
// QMap<int, QVariant> itemData(const class QModelIndex &)
func (this *QAbstractProxyModel) ItemData(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel8itemDataERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:77
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QAbstractProxyModel) Flags(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:79
// index:0
// virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QAbstractProxyModel) SetData(index unsafe.Pointer, value unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, value const QVariant &, role int), (index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:81
// index:0
// virtual
// bool setHeaderData(int, Qt::Orientation, const class QVariant &, int)
func (this *QAbstractProxyModel) SetHeaderData(section int, orientation int, value unsafe.Pointer, role int) {
	// 0: (, section int, orientation Qt::Orientation, value const QVariant &, role int), (&section, &orientation, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &section, &orientation, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:83
// index:0
// virtual
// QModelIndex buddy(const class QModelIndex &)
func (this *QAbstractProxyModel) Buddy(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel5buddyERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:84
// index:0
// virtual
// bool canFetchMore(const class QModelIndex &)
func (this *QAbstractProxyModel) CanFetchMore(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel12canFetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:85
// index:0
// virtual
// void fetchMore(const class QModelIndex &)
func (this *QAbstractProxyModel) FetchMore(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel9fetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:86
// index:0
// virtual
// void sort(int, Qt::SortOrder)
func (this *QAbstractProxyModel) Sort(column int, order int) {
	// 0: (, column int, order Qt::SortOrder), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:87
// index:0
// virtual
// QSize span(const class QModelIndex &)
func (this *QAbstractProxyModel) Span(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel4spanERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:88
// index:0
// virtual
// bool hasChildren(const class QModelIndex &)
func (this *QAbstractProxyModel) HasChildren(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:89
// index:0
// virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QAbstractProxyModel) Sibling(row int, column int, idx unsafe.Pointer) {
	// 0: (, row int, column int, idx const QModelIndex &), (&row, &column, idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:92
// index:0
// virtual
// bool canDropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QAbstractProxyModel) CanDropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, data const QMimeData *, action Qt::DropAction, row int, column int, parent const QModelIndex &), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel15canDropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:94
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QAbstractProxyModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, data const QMimeData *, action Qt::DropAction, row int, column int, parent const QModelIndex &), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:96
// index:0
// virtual
// QStringList mimeTypes()
func (this *QAbstractProxyModel) MimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel9mimeTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:97
// index:0
// virtual
// Qt::DropActions supportedDragActions()
func (this *QAbstractProxyModel) SupportedDragActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20supportedDragActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:98
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QAbstractProxyModel) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:104
// index:0
// void resetInternalData()
func (this *QAbstractProxyModel) ResetInternalData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QAbstractProxyModel17resetInternalDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
