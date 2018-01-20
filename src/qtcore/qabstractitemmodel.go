//  header block begin
// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>
package qtcore

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
type QAbstractItemModel struct {
	*QObject
}

func (this *QAbstractItemModel) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:167
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractItemModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:174
// index:0
// void QAbstractItemModel(class QObject *)
func NewQAbstractItemModel(parent unsafe.Pointer) *QAbstractItemModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModelC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractItemModelFromPointer(cthis)
	return gothis
}
func NewQAbstractItemModelFromPointer(cthis unsafe.Pointer) *QAbstractItemModel {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractItemModel{bcthis0}
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:289
// index:1
// void QAbstractItemModel(class QAbstractItemModelPrivate &, class QObject *)
func NewQAbstractItemModel_1(dd unsafe.Pointer, parent unsafe.Pointer) *QAbstractItemModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModelC1ER25QAbstractItemModelPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractItemModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:175
// index:0
// virtual
// void ~QAbstractItemModel()
func DeleteQAbstractItemModel(*QAbstractItemModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:177
// index:0
// bool hasIndex(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) HasIndex(row int, column int, parent unsafe.Pointer) {
	// 0: (, row int, column int, parent const QModelIndex &), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:178
// index:0
// pure virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) Index(row int, column int, parent unsafe.Pointer) {
	// 0: (, row int, column int, parent const QModelIndex &), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:180
// index:0
// pure virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QAbstractItemModel) Parent(child unsafe.Pointer) {
	// 0: (, child const QModelIndex &), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:182
// index:0
// virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) Sibling(row int, column int, idx unsafe.Pointer) {
	// 0: (, row int, column int, idx const QModelIndex &), (&row, &column, idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:183
// index:0
// pure virtual
// int rowCount(const class QModelIndex &)
func (this *QAbstractItemModel) RowCount(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:184
// index:0
// pure virtual
// int columnCount(const class QModelIndex &)
func (this *QAbstractItemModel) ColumnCount(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:185
// index:0
// virtual
// bool hasChildren(const class QModelIndex &)
func (this *QAbstractItemModel) HasChildren(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:187
// index:0
// pure virtual
// QVariant data(const class QModelIndex &, int)
func (this *QAbstractItemModel) Data(index unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, role int), (index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:188
// index:0
// virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QAbstractItemModel) SetData(index unsafe.Pointer, value unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, value const QVariant &, role int), (index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:190
// index:0
// virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QAbstractItemModel) HeaderData(section int, orientation int, role int) {
	// 0: (, section int, orientation Qt::Orientation, role int), (&section, &orientation, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:192
// index:0
// virtual
// bool setHeaderData(int, Qt::Orientation, const class QVariant &, int)
func (this *QAbstractItemModel) SetHeaderData(section int, orientation int, value unsafe.Pointer, role int) {
	// 0: (, section int, orientation Qt::Orientation, value const QVariant &, role int), (&section, &orientation, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &section, &orientation, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:195
// index:0
// virtual
// QMap<int, QVariant> itemData(const class QModelIndex &)
func (this *QAbstractItemModel) ItemData(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel8itemDataERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:198
// index:0
// virtual
// QStringList mimeTypes()
func (this *QAbstractItemModel) MimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel9mimeTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:200
// index:0
// virtual
// bool canDropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QAbstractItemModel) CanDropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, data const QMimeData *, action Qt::DropAction, row int, column int, parent const QModelIndex &), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel15canDropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:202
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QAbstractItemModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, data const QMimeData *, action Qt::DropAction, row int, column int, parent const QModelIndex &), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:204
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QAbstractItemModel) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:206
// index:0
// virtual
// Qt::DropActions supportedDragActions()
func (this *QAbstractItemModel) SupportedDragActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel20supportedDragActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:212
// index:0
// virtual
// bool insertRows(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) InsertRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, row int, count int, parent const QModelIndex &), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:213
// index:0
// virtual
// bool insertColumns(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) InsertColumns(column int, count int, parent unsafe.Pointer) {
	// 0: (, column int, count int, parent const QModelIndex &), (&column, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:214
// index:0
// virtual
// bool removeRows(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) RemoveRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, row int, count int, parent const QModelIndex &), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:215
// index:0
// virtual
// bool removeColumns(int, int, const class QModelIndex &)
func (this *QAbstractItemModel) RemoveColumns(column int, count int, parent unsafe.Pointer) {
	// 0: (, column int, count int, parent const QModelIndex &), (&column, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:216
// index:0
// virtual
// bool moveRows(const class QModelIndex &, int, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) MoveRows(sourceParent unsafe.Pointer, sourceRow int, count int, destinationParent unsafe.Pointer, destinationChild int) {
	// 0: (, sourceParent const QModelIndex &, sourceRow int, count int, destinationParent const QModelIndex &, destinationChild int), (sourceParent, &sourceRow, &count, destinationParent, &destinationChild)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceParent, &sourceRow, &count, destinationParent, &destinationChild)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:218
// index:0
// virtual
// bool moveColumns(const class QModelIndex &, int, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) MoveColumns(sourceParent unsafe.Pointer, sourceColumn int, count int, destinationParent unsafe.Pointer, destinationChild int) {
	// 0: (, sourceParent const QModelIndex &, sourceColumn int, count int, destinationParent const QModelIndex &, destinationChild int), (sourceParent, &sourceColumn, &count, destinationParent, &destinationChild)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceParent, &sourceColumn, &count, destinationParent, &destinationChild)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:221
// index:0
// inline
// bool insertRow(int, const class QModelIndex &)
func (this *QAbstractItemModel) InsertRow(row int, parent unsafe.Pointer) {
	// 0: (, row int, parent const QModelIndex &), (&row, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:222
// index:0
// inline
// bool insertColumn(int, const class QModelIndex &)
func (this *QAbstractItemModel) InsertColumn(column int, parent unsafe.Pointer) {
	// 0: (, column int, parent const QModelIndex &), (&column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:223
// index:0
// inline
// bool removeRow(int, const class QModelIndex &)
func (this *QAbstractItemModel) RemoveRow(row int, parent unsafe.Pointer) {
	// 0: (, row int, parent const QModelIndex &), (&row, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:224
// index:0
// inline
// bool removeColumn(int, const class QModelIndex &)
func (this *QAbstractItemModel) RemoveColumn(column int, parent unsafe.Pointer) {
	// 0: (, column int, parent const QModelIndex &), (&column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:225
// index:0
// inline
// bool moveRow(const class QModelIndex &, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) MoveRow(sourceParent unsafe.Pointer, sourceRow int, destinationParent unsafe.Pointer, destinationChild int) {
	// 0: (, sourceParent const QModelIndex &, sourceRow int, destinationParent const QModelIndex &, destinationChild int), (sourceParent, &sourceRow, destinationParent, &destinationChild)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceParent, &sourceRow, destinationParent, &destinationChild)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:227
// index:0
// inline
// bool moveColumn(const class QModelIndex &, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) MoveColumn(sourceParent unsafe.Pointer, sourceColumn int, destinationParent unsafe.Pointer, destinationChild int) {
	// 0: (, sourceParent const QModelIndex &, sourceColumn int, destinationParent const QModelIndex &, destinationChild int), (sourceParent, &sourceColumn, destinationParent, &destinationChild)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceParent, &sourceColumn, destinationParent, &destinationChild)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:230
// index:0
// virtual
// void fetchMore(const class QModelIndex &)
func (this *QAbstractItemModel) FetchMore(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel9fetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:231
// index:0
// virtual
// bool canFetchMore(const class QModelIndex &)
func (this *QAbstractItemModel) CanFetchMore(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:232
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QAbstractItemModel) Flags(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:233
// index:0
// virtual
// void sort(int, Qt::SortOrder)
func (this *QAbstractItemModel) Sort(column int, order int) {
	// 0: (, column int, order Qt::SortOrder), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:234
// index:0
// virtual
// QModelIndex buddy(const class QModelIndex &)
func (this *QAbstractItemModel) Buddy(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel5buddyERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:235
// index:0
// virtual
// QModelIndexList match(const class QModelIndex &, int, const class QVariant &, int, Qt::MatchFlags)
func (this *QAbstractItemModel) Match(start unsafe.Pointer, role int, value unsafe.Pointer, hits int, flags int) {
	// 0: (, start const QModelIndex &, role int, value const QVariant &, hits int, QFlags<Qt::MatchFlag> flags), (start, &role, value, &hits, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), start, &role, value, &hits, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:239
// index:0
// virtual
// QSize span(const class QModelIndex &)
func (this *QAbstractItemModel) Span(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel4spanERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:241
// index:0
// virtual
// QHash<int, QByteArray> roleNames()
func (this *QAbstractItemModel) RoleNames() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel9roleNamesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:255
// index:0
// void headerDataChanged(Qt::Orientation, int, int)
func (this *QAbstractItemModel) HeaderDataChanged(orientation int, first int, last int) {
	// 0: (, orientation Qt::Orientation, first int, last int), (&orientation, &first, &last)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel17headerDataChangedEN2Qt11OrientationEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &orientation, &first, &last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:281
// index:0
// virtual
// bool submit()
func (this *QAbstractItemModel) Submit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel6submitEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:282
// index:0
// virtual
// void revert()
func (this *QAbstractItemModel) Revert() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel6revertEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:286
// index:0
// void resetInternalData()
func (this *QAbstractItemModel) ResetInternalData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel17resetInternalDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:291
// index:0
// inline
// QModelIndex createIndex(int, int, void *)
func (this *QAbstractItemModel) CreateIndex(row int, column int, data unsafe.Pointer) {
	// 0: (, row int, column int, data void *), (&row, &column, data)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel11createIndexEiiPv", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:292
// index:1
// inline
// QModelIndex createIndex(int, int, quintptr)
func (this *QAbstractItemModel) CreateIndex_1(row int, column int, id uint64) {
	// 1: (, row int, column int, id quintptr), (&row, &column, &id)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel11createIndexEiiy", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, &id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:295
// index:0
// bool decodeData(int, int, const class QModelIndex &, class QDataStream &)
func (this *QAbstractItemModel) DecodeData(row int, column int, parent unsafe.Pointer, stream unsafe.Pointer) {
	// 0: (, row int, column int, parent const QModelIndex &, stream QDataStream &), (&row, &column, parent, stream)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel10decodeDataEiiRK11QModelIndexR11QDataStream", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, parent, stream)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:297
// index:0
// void beginInsertRows(const class QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginInsertRows(parent unsafe.Pointer, first int, last int) {
	// 0: (, parent const QModelIndex &, first int, last int), (parent, &first, &last)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginInsertRowsERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &first, &last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:298
// index:0
// void endInsertRows()
func (this *QAbstractItemModel) EndInsertRows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13endInsertRowsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:300
// index:0
// void beginRemoveRows(const class QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginRemoveRows(parent unsafe.Pointer, first int, last int) {
	// 0: (, parent const QModelIndex &, first int, last int), (parent, &first, &last)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginRemoveRowsERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &first, &last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:301
// index:0
// void endRemoveRows()
func (this *QAbstractItemModel) EndRemoveRows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13endRemoveRowsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:303
// index:0
// bool beginMoveRows(const class QModelIndex &, int, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) BeginMoveRows(sourceParent unsafe.Pointer, sourceFirst int, sourceLast int, destinationParent unsafe.Pointer, destinationRow int) {
	// 0: (, sourceParent const QModelIndex &, sourceFirst int, sourceLast int, destinationParent const QModelIndex &, destinationRow int), (sourceParent, &sourceFirst, &sourceLast, destinationParent, &destinationRow)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13beginMoveRowsERK11QModelIndexiiS2_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceParent, &sourceFirst, &sourceLast, destinationParent, &destinationRow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:304
// index:0
// void endMoveRows()
func (this *QAbstractItemModel) EndMoveRows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel11endMoveRowsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:306
// index:0
// void beginInsertColumns(const class QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginInsertColumns(parent unsafe.Pointer, first int, last int) {
	// 0: (, parent const QModelIndex &, first int, last int), (parent, &first, &last)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel18beginInsertColumnsERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &first, &last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:307
// index:0
// void endInsertColumns()
func (this *QAbstractItemModel) EndInsertColumns() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel16endInsertColumnsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:309
// index:0
// void beginRemoveColumns(const class QModelIndex &, int, int)
func (this *QAbstractItemModel) BeginRemoveColumns(parent unsafe.Pointer, first int, last int) {
	// 0: (, parent const QModelIndex &, first int, last int), (parent, &first, &last)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel18beginRemoveColumnsERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &first, &last)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:310
// index:0
// void endRemoveColumns()
func (this *QAbstractItemModel) EndRemoveColumns() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel16endRemoveColumnsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:312
// index:0
// bool beginMoveColumns(const class QModelIndex &, int, int, const class QModelIndex &, int)
func (this *QAbstractItemModel) BeginMoveColumns(sourceParent unsafe.Pointer, sourceFirst int, sourceLast int, destinationParent unsafe.Pointer, destinationColumn int) {
	// 0: (, sourceParent const QModelIndex &, sourceFirst int, sourceLast int, destinationParent const QModelIndex &, destinationColumn int), (sourceParent, &sourceFirst, &sourceLast, destinationParent, &destinationColumn)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel16beginMoveColumnsERK11QModelIndexiiS2_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceParent, &sourceFirst, &sourceLast, destinationParent, &destinationColumn)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:313
// index:0
// void endMoveColumns()
func (this *QAbstractItemModel) EndMoveColumns() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel14endMoveColumnsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:324
// index:0
// void beginResetModel()
func (this *QAbstractItemModel) BeginResetModel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginResetModelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:325
// index:0
// void endResetModel()
func (this *QAbstractItemModel) EndResetModel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel13endResetModelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:327
// index:0
// void changePersistentIndex(const class QModelIndex &, const class QModelIndex &)
func (this *QAbstractItemModel) ChangePersistentIndex(from unsafe.Pointer, to unsafe.Pointer) {
	// 0: (, from const QModelIndex &, to const QModelIndex &), (from, to)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QAbstractItemModel21changePersistentIndexERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), from, to)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:329
// index:0
// QModelIndexList persistentIndexList()
func (this *QAbstractItemModel) PersistentIndexList() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QAbstractItemModel19persistentIndexListEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
