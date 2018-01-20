//  header block begin
// /usr/include/qt/QtCore/qsortfilterproxymodel.h
// #include <qsortfilterproxymodel.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
	*QAbstractProxyModel
}

func (this *QSortFilterProxyModel) GetCthis() unsafe.Pointer {
	return this.QAbstractProxyModel.GetCthis()
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSortFilterProxyModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:73
// index:0
// void QSortFilterProxyModel(class QObject *)
func NewQSortFilterProxyModel(parent unsafe.Pointer) *QSortFilterProxyModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSortFilterProxyModelFromPointer(cthis)
	return gothis
}
func NewQSortFilterProxyModelFromPointer(cthis unsafe.Pointer) *QSortFilterProxyModel {
	bcthis0 := NewQAbstractProxyModelFromPointer(cthis)
	return &QSortFilterProxyModel{bcthis0}
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
	// 0: (, sourceModel QAbstractItemModel *), (sourceModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:78
// index:0
// virtual
// QModelIndex mapToSource(const class QModelIndex &)
func (this *QSortFilterProxyModel) MapToSource(proxyIndex unsafe.Pointer) {
	// 0: (, proxyIndex const QModelIndex &), (proxyIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), proxyIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:79
// index:0
// virtual
// QModelIndex mapFromSource(const class QModelIndex &)
func (this *QSortFilterProxyModel) MapFromSource(sourceIndex unsafe.Pointer) {
	// 0: (, sourceIndex const QModelIndex &), (sourceIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:81
// index:0
// virtual
// QItemSelection mapSelectionToSource(const class QItemSelection &)
func (this *QSortFilterProxyModel) MapSelectionToSource(proxySelection unsafe.Pointer) {
	// 0: (, proxySelection const QItemSelection &), (proxySelection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.GetCthis(), proxySelection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:82
// index:0
// virtual
// QItemSelection mapSelectionFromSource(const class QItemSelection &)
func (this *QSortFilterProxyModel) MapSelectionFromSource(sourceSelection unsafe.Pointer) {
	// 0: (, sourceSelection const QItemSelection &), (sourceSelection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sourceSelection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:84
// index:0
// QRegExp filterRegExp()
func (this *QSortFilterProxyModel) FilterRegExp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel12filterRegExpEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:85
// index:0
// void setFilterRegExp(const class QRegExp &)
func (this *QSortFilterProxyModel) SetFilterRegExp(regExp unsafe.Pointer) {
	// 0: (, regExp const QRegExp &), (regExp)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp", ffiqt.FFI_TYPE_VOID, this.GetCthis(), regExp)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:115
// index:1
// void setFilterRegExp(const class QString &)
func (this *QSortFilterProxyModel) SetFilterRegExp_1(pattern unsafe.Pointer) {
	// 1: (, pattern const QString &), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel15setFilterRegExpERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:87
// index:0
// int filterKeyColumn()
func (this *QSortFilterProxyModel) FilterKeyColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel15filterKeyColumnEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:88
// index:0
// void setFilterKeyColumn(int)
func (this *QSortFilterProxyModel) SetFilterKeyColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel18setFilterKeyColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:90
// index:0
// Qt::CaseSensitivity filterCaseSensitivity()
func (this *QSortFilterProxyModel) FilterCaseSensitivity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel21filterCaseSensitivityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:91
// index:0
// void setFilterCaseSensitivity(Qt::CaseSensitivity)
func (this *QSortFilterProxyModel) SetFilterCaseSensitivity(cs int) {
	// 0: (, cs Qt::CaseSensitivity), (&cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel24setFilterCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:93
// index:0
// Qt::CaseSensitivity sortCaseSensitivity()
func (this *QSortFilterProxyModel) SortCaseSensitivity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel19sortCaseSensitivityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:94
// index:0
// void setSortCaseSensitivity(Qt::CaseSensitivity)
func (this *QSortFilterProxyModel) SetSortCaseSensitivity(cs int) {
	// 0: (, cs Qt::CaseSensitivity), (&cs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel22setSortCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:96
// index:0
// bool isSortLocaleAware()
func (this *QSortFilterProxyModel) IsSortLocaleAware() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel17isSortLocaleAwareEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:97
// index:0
// void setSortLocaleAware(_Bool)
func (this *QSortFilterProxyModel) SetSortLocaleAware(on bool) {
	// 0: (, on bool), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel18setSortLocaleAwareEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:99
// index:0
// int sortColumn()
func (this *QSortFilterProxyModel) SortColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10sortColumnEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:100
// index:0
// Qt::SortOrder sortOrder()
func (this *QSortFilterProxyModel) SortOrder() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel9sortOrderEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:102
// index:0
// bool dynamicSortFilter()
func (this *QSortFilterProxyModel) DynamicSortFilter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel17dynamicSortFilterEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:103
// index:0
// void setDynamicSortFilter(_Bool)
func (this *QSortFilterProxyModel) SetDynamicSortFilter(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel20setDynamicSortFilterEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:105
// index:0
// int sortRole()
func (this *QSortFilterProxyModel) SortRole() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8sortRoleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:106
// index:0
// void setSortRole(int)
func (this *QSortFilterProxyModel) SetSortRole(role int) {
	// 0: (, role int), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel11setSortRoleEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:108
// index:0
// int filterRole()
func (this *QSortFilterProxyModel) FilterRole() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10filterRoleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:109
// index:0
// void setFilterRole(int)
func (this *QSortFilterProxyModel) SetFilterRole(role int) {
	// 0: (, role int), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13setFilterRoleEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:111
// index:0
// bool isRecursiveFilteringEnabled()
func (this *QSortFilterProxyModel) IsRecursiveFilteringEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel27isRecursiveFilteringEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:112
// index:0
// void setRecursiveFilteringEnabled(_Bool)
func (this *QSortFilterProxyModel) SetRecursiveFilteringEnabled(recursive bool) {
	// 0: (, recursive bool), (&recursive)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel28setRecursiveFilteringEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &recursive)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:116
// index:0
// void setFilterWildcard(const class QString &)
func (this *QSortFilterProxyModel) SetFilterWildcard(pattern unsafe.Pointer) {
	// 0: (, pattern const QString &), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel17setFilterWildcardERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:117
// index:0
// void setFilterFixedString(const class QString &)
func (this *QSortFilterProxyModel) SetFilterFixedString(pattern unsafe.Pointer) {
	// 0: (, pattern const QString &), (pattern)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pattern)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:118
// index:0
// void clear()
func (this *QSortFilterProxyModel) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:119
// index:0
// void invalidate()
func (this *QSortFilterProxyModel) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10invalidateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:122
// index:0
// virtual
// bool filterAcceptsRow(int, const class QModelIndex &)
func (this *QSortFilterProxyModel) FilterAcceptsRow(source_row int, source_parent unsafe.Pointer) {
	// 0: (, source_row int, source_parent const QModelIndex &), (&source_row, source_parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel16filterAcceptsRowEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &source_row, source_parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:123
// index:0
// virtual
// bool filterAcceptsColumn(int, const class QModelIndex &)
func (this *QSortFilterProxyModel) FilterAcceptsColumn(source_column int, source_parent unsafe.Pointer) {
	// 0: (, source_column int, source_parent const QModelIndex &), (&source_column, source_parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel19filterAcceptsColumnEiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &source_column, source_parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:124
// index:0
// virtual
// bool lessThan(const class QModelIndex &, const class QModelIndex &)
func (this *QSortFilterProxyModel) LessThan(source_left unsafe.Pointer, source_right unsafe.Pointer) {
	// 0: (, source_left const QModelIndex &, source_right const QModelIndex &), (source_left, source_right)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8lessThanERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), source_left, source_right)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:126
// index:0
// void filterChanged()
func (this *QSortFilterProxyModel) FilterChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13filterChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:127
// index:0
// void invalidateFilter()
func (this *QSortFilterProxyModel) InvalidateFilter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel16invalidateFilterEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:132
// index:0
// virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) Index(row int, column int, parent unsafe.Pointer) {
	// 0: (, row int, column int, parent const QModelIndex &), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:133
// index:0
// virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QSortFilterProxyModel) Parent(child unsafe.Pointer) {
	// 0: (, child const QModelIndex &), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:134
// index:0
// virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) Sibling(row int, column int, idx unsafe.Pointer) {
	// 0: (, row int, column int, idx const QModelIndex &), (&row, &column, idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:136
// index:0
// virtual
// int rowCount(const class QModelIndex &)
func (this *QSortFilterProxyModel) RowCount(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:137
// index:0
// virtual
// int columnCount(const class QModelIndex &)
func (this *QSortFilterProxyModel) ColumnCount(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:138
// index:0
// virtual
// bool hasChildren(const class QModelIndex &)
func (this *QSortFilterProxyModel) HasChildren(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:140
// index:0
// virtual
// QVariant data(const class QModelIndex &, int)
func (this *QSortFilterProxyModel) Data(index unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, role int), (index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:141
// index:0
// virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QSortFilterProxyModel) SetData(index unsafe.Pointer, value unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, value const QVariant &, role int), (index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:143
// index:0
// virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QSortFilterProxyModel) HeaderData(section int, orientation int, role int) {
	// 0: (, section int, orientation Qt::Orientation, role int), (&section, &orientation, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:144
// index:0
// virtual
// bool setHeaderData(int, Qt::Orientation, const class QVariant &, int)
func (this *QSortFilterProxyModel) SetHeaderData(section int, orientation int, value unsafe.Pointer, role int) {
	// 0: (, section int, orientation Qt::Orientation, value const QVariant &, role int), (&section, &orientation, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &section, &orientation, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:148
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, data const QMimeData *, action Qt::DropAction, row int, column int, parent const QModelIndex &), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:151
// index:0
// virtual
// bool insertRows(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) InsertRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, row int, count int, parent const QModelIndex &), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:152
// index:0
// virtual
// bool insertColumns(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) InsertColumns(column int, count int, parent unsafe.Pointer) {
	// 0: (, column int, count int, parent const QModelIndex &), (&column, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:153
// index:0
// virtual
// bool removeRows(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) RemoveRows(row int, count int, parent unsafe.Pointer) {
	// 0: (, row int, count int, parent const QModelIndex &), (&row, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:154
// index:0
// virtual
// bool removeColumns(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) RemoveColumns(column int, count int, parent unsafe.Pointer) {
	// 0: (, column int, count int, parent const QModelIndex &), (&column, &count, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &count, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:156
// index:0
// virtual
// void fetchMore(const class QModelIndex &)
func (this *QSortFilterProxyModel) FetchMore(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:157
// index:0
// virtual
// bool canFetchMore(const class QModelIndex &)
func (this *QSortFilterProxyModel) CanFetchMore(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:158
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QSortFilterProxyModel) Flags(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:160
// index:0
// virtual
// QModelIndex buddy(const class QModelIndex &)
func (this *QSortFilterProxyModel) Buddy(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5buddyERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:161
// index:0
// virtual
// QModelIndexList match(const class QModelIndex &, int, const class QVariant &, int, Qt::MatchFlags)
func (this *QSortFilterProxyModel) Match(start unsafe.Pointer, role int, value unsafe.Pointer, hits int, flags int) {
	// 0: (, start const QModelIndex &, role int, value const QVariant &, hits int, QFlags<Qt::MatchFlag> flags), (start, &role, value, &hits, &flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), start, &role, value, &hits, &flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:165
// index:0
// virtual
// QSize span(const class QModelIndex &)
func (this *QSortFilterProxyModel) Span(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel4spanERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:166
// index:0
// virtual
// void sort(int, Qt::SortOrder)
func (this *QSortFilterProxyModel) Sort(column int, order int) {
	// 0: (, column int, order Qt::SortOrder), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:168
// index:0
// virtual
// QStringList mimeTypes()
func (this *QSortFilterProxyModel) MimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel9mimeTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:169
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QSortFilterProxyModel) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
