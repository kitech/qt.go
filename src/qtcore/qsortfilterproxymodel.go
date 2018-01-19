//  header block begin
// /usr/include/qt/QtCore/qsortfilterproxymodel.h
// #include <qsortfilterproxymodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QSortFilterProxyModel struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSortFilterProxyModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:73
// index:0
// void QSortFilterProxyModel(class QObject *)
func NewQSortFilterProxyModel(parent unsafe.Pointer) *QSortFilterProxyModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QSortFilterProxyModel{cthis}
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:74
// index:0
// virtual
// void ~QSortFilterProxyModel()
func DeleteQSortFilterProxyModel(*QSortFilterProxyModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:76
// index:0
// virtual
// void setSourceModel(class QAbstractItemModel *)
func (this *QSortFilterProxyModel) SetSourceModel(sourceModel unsafe.Pointer) {
	// 0: (, QAbstractItemModel * sourceModel), (sourceModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.cthis, sourceModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:78
// index:0
// virtual
// QModelIndex mapToSource(const class QModelIndex &)
func (this *QSortFilterProxyModel) MapToSource(proxyIndex unsafe.Pointer) {
	// 0: (, const QModelIndex & proxyIndex), (proxyIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, proxyIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:79
// index:0
// virtual
// QModelIndex mapFromSource(const class QModelIndex &)
func (this *QSortFilterProxyModel) MapFromSource(sourceIndex unsafe.Pointer) {
	// 0: (, const QModelIndex & sourceIndex), (sourceIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, sourceIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:81
// index:0
// virtual
// QItemSelection mapSelectionToSource(const class QItemSelection &)
func (this *QSortFilterProxyModel) MapSelectionToSource(proxySelection unsafe.Pointer) {
	// 0: (, const QItemSelection & proxySelection), (proxySelection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.cthis, proxySelection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:82
// index:0
// virtual
// QItemSelection mapSelectionFromSource(const class QItemSelection &)
func (this *QSortFilterProxyModel) MapSelectionFromSource(sourceSelection unsafe.Pointer) {
	// 0: (, const QItemSelection & sourceSelection), (sourceSelection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.cthis, sourceSelection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:84
// index:0
// QRegExp filterRegExp()
func (this *QSortFilterProxyModel) FilterRegExp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel12filterRegExpEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:85
// index:0
// void setFilterRegExp(const class QRegExp &)
func (this *QSortFilterProxyModel) SetFilterRegExp(regExp unsafe.Pointer) {
	// 0: (, const QRegExp & regExp), (regExp)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp", ffiqt.FFI_TYPE_VOID, this.cthis, regExp)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:115
// index:1
// void setFilterRegExp(const class QString &)
func (this *QSortFilterProxyModel) SetFilterRegExp_1(pattern unsafe.Pointer) {
	// 1: (, const QString & pattern), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel15setFilterRegExpERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:87
// index:0
// int filterKeyColumn()
func (this *QSortFilterProxyModel) FilterKeyColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel15filterKeyColumnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:88
// index:0
// void setFilterKeyColumn(int)
func (this *QSortFilterProxyModel) SetFilterKeyColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel18setFilterKeyColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:90
// index:0
// Qt::CaseSensitivity filterCaseSensitivity()
func (this *QSortFilterProxyModel) FilterCaseSensitivity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel21filterCaseSensitivityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:91
// index:0
// void setFilterCaseSensitivity(Qt::CaseSensitivity)
func (this *QSortFilterProxyModel) SetFilterCaseSensitivity(cs int) {
	// 0: (, Qt::CaseSensitivity cs), (&cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel24setFilterCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:93
// index:0
// Qt::CaseSensitivity sortCaseSensitivity()
func (this *QSortFilterProxyModel) SortCaseSensitivity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel19sortCaseSensitivityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:94
// index:0
// void setSortCaseSensitivity(Qt::CaseSensitivity)
func (this *QSortFilterProxyModel) SetSortCaseSensitivity(cs int) {
	// 0: (, Qt::CaseSensitivity cs), (&cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel22setSortCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:96
// index:0
// bool isSortLocaleAware()
func (this *QSortFilterProxyModel) IsSortLocaleAware() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel17isSortLocaleAwareEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:97
// index:0
// void setSortLocaleAware(_Bool)
func (this *QSortFilterProxyModel) SetSortLocaleAware(on bool) {
	// 0: (, bool on), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel18setSortLocaleAwareEb", ffiqt.FFI_TYPE_VOID, this.cthis, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:99
// index:0
// int sortColumn()
func (this *QSortFilterProxyModel) SortColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10sortColumnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:100
// index:0
// Qt::SortOrder sortOrder()
func (this *QSortFilterProxyModel) SortOrder() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel9sortOrderEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:102
// index:0
// bool dynamicSortFilter()
func (this *QSortFilterProxyModel) DynamicSortFilter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel17dynamicSortFilterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:103
// index:0
// void setDynamicSortFilter(_Bool)
func (this *QSortFilterProxyModel) SetDynamicSortFilter(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel20setDynamicSortFilterEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:105
// index:0
// int sortRole()
func (this *QSortFilterProxyModel) SortRole() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8sortRoleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:106
// index:0
// void setSortRole(int)
func (this *QSortFilterProxyModel) SetSortRole(role int) {
	// 0: (, int role), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel11setSortRoleEi", ffiqt.FFI_TYPE_VOID, this.cthis, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:108
// index:0
// int filterRole()
func (this *QSortFilterProxyModel) FilterRole() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10filterRoleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:109
// index:0
// void setFilterRole(int)
func (this *QSortFilterProxyModel) SetFilterRole(role int) {
	// 0: (, int role), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13setFilterRoleEi", ffiqt.FFI_TYPE_VOID, this.cthis, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:111
// index:0
// bool isRecursiveFilteringEnabled()
func (this *QSortFilterProxyModel) IsRecursiveFilteringEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel27isRecursiveFilteringEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:112
// index:0
// void setRecursiveFilteringEnabled(_Bool)
func (this *QSortFilterProxyModel) SetRecursiveFilteringEnabled(recursive bool) {
	// 0: (, bool recursive), (&recursive)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel28setRecursiveFilteringEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &recursive)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:116
// index:0
// void setFilterWildcard(const class QString &)
func (this *QSortFilterProxyModel) SetFilterWildcard(pattern unsafe.Pointer) {
	// 0: (, const QString & pattern), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel17setFilterWildcardERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:117
// index:0
// void setFilterFixedString(const class QString &)
func (this *QSortFilterProxyModel) SetFilterFixedString(pattern unsafe.Pointer) {
	// 0: (, const QString & pattern), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:118
// index:0
// void clear()
func (this *QSortFilterProxyModel) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:119
// index:0
// void invalidate()
func (this *QSortFilterProxyModel) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10invalidateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:132
// index:0
// virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) Index(row int, column int, parent unsafe.Pointer) {
	// 0: (, int row, int column, const QModelIndex & parent), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:133
// index:0
// virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QSortFilterProxyModel) Parent(child unsafe.Pointer) {
	// 0: (, const QModelIndex & child), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:134
// index:0
// virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) Sibling(row int, column int, idx unsafe.Pointer) {
	// 0: (, int row, int column, const QModelIndex & idx), (&row, &column, idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:136
// index:0
// virtual
// int rowCount(const class QModelIndex &)
func (this *QSortFilterProxyModel) RowCount(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:137
// index:0
// virtual
// int columnCount(const class QModelIndex &)
func (this *QSortFilterProxyModel) ColumnCount(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:138
// index:0
// virtual
// bool hasChildren(const class QModelIndex &)
func (this *QSortFilterProxyModel) HasChildren(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:140
// index:0
// virtual
// QVariant data(const class QModelIndex &, int)
func (this *QSortFilterProxyModel) Data(index unsafe.Pointer, role int) {
	// 0: (, const QModelIndex & index, int role), (index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_VOID, this.cthis, index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:141
// index:0
// virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QSortFilterProxyModel) SetData(index unsafe.Pointer, value unsafe.Pointer, role int) {
	// 0: (, const QModelIndex & index, const QVariant & value, int role), (index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.cthis, index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:143
// index:0
// virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QSortFilterProxyModel) HeaderData(section int, orientation int, role int) {
	// 0: (, int section, Qt::Orientation orientation, int role), (&section, &orientation, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_VOID, this.cthis, &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:144
// index:0
// virtual
// bool setHeaderData(int, Qt::Orientation, const class QVariant &, int)
func (this *QSortFilterProxyModel) SetHeaderData(section int, orientation int, value unsafe.Pointer, role int) {
	// 0: (, int section, Qt::Orientation orientation, const QVariant & value, int role), (&section, &orientation, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", ffiqt.FFI_TYPE_VOID, this.cthis, &section, &orientation, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:148
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:151
// index:0
// virtual
// bool insertRows(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) InsertRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, int row, int count, const QModelIndex & parent), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:152
// index:0
// virtual
// bool insertColumns(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) InsertColumns(column int, count int, parent unsafe.Pointer) {
	// 0: (, int column, int count, const QModelIndex & parent), (&column, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:153
// index:0
// virtual
// bool removeRows(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) RemoveRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, int row, int count, const QModelIndex & parent), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:154
// index:0
// virtual
// bool removeColumns(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) RemoveColumns(column int, count int, parent unsafe.Pointer) {
	// 0: (, int column, int count, const QModelIndex & parent), (&column, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:156
// index:0
// virtual
// void fetchMore(const class QModelIndex &)
func (this *QSortFilterProxyModel) FetchMore(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:157
// index:0
// virtual
// bool canFetchMore(const class QModelIndex &)
func (this *QSortFilterProxyModel) CanFetchMore(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:158
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QSortFilterProxyModel) Flags(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:160
// index:0
// virtual
// QModelIndex buddy(const class QModelIndex &)
func (this *QSortFilterProxyModel) Buddy(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5buddyERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:165
// index:0
// virtual
// QSize span(const class QModelIndex &)
func (this *QSortFilterProxyModel) Span(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel4spanERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:166
// index:0
// virtual
// void sort(int, Qt::SortOrder)
func (this *QSortFilterProxyModel) Sort(column int, order int) {
	// 0: (, int column, Qt::SortOrder order), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:168
// index:0
// virtual
// QStringList mimeTypes()
func (this *QSortFilterProxyModel) MimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel9mimeTypesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:169
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QSortFilterProxyModel) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
