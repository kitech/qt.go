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
func NewQSortFilterProxyModelFromPointer(cthis unsafe.Pointer) *QSortFilterProxyModel {
	bcthis0 := NewQAbstractProxyModelFromPointer(cthis)
	return &QSortFilterProxyModel{bcthis0}
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:61
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSortFilterProxyModel) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:73
// index:0
// Public
// void QSortFilterProxyModel(class QObject *)
func NewQSortFilterProxyModel(parent unsafe.Pointer) *QSortFilterProxyModel {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSortFilterProxyModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:74
// index:0
// Public virtual
// void ~QSortFilterProxyModel()
func DeleteQSortFilterProxyModel(*QSortFilterProxyModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:76
// index:0
// Public virtual
// void setSourceModel(class QAbstractItemModel *)
func (this *QSortFilterProxyModel) SetSourceModel(sourceModel unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), sourceModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:78
// index:0
// Public virtual
// QModelIndex mapToSource(const class QModelIndex &)
func (this *QSortFilterProxyModel) MapToSource(proxyIndex *QModelIndex) interface{} {
	var convArg0 = proxyIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:79
// index:0
// Public virtual
// QModelIndex mapFromSource(const class QModelIndex &)
func (this *QSortFilterProxyModel) MapFromSource(sourceIndex *QModelIndex) interface{} {
	var convArg0 = sourceIndex.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:81
// index:0
// Public virtual
// QItemSelection mapSelectionToSource(const class QItemSelection &)
func (this *QSortFilterProxyModel) MapSelectionToSource(proxySelection *QItemSelection) interface{} {
	var convArg0 = proxySelection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:82
// index:0
// Public virtual
// QItemSelection mapSelectionFromSource(const class QItemSelection &)
func (this *QSortFilterProxyModel) MapSelectionFromSource(sourceSelection *QItemSelection) interface{} {
	var convArg0 = sourceSelection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:84
// index:0
// Public
// QRegExp filterRegExp()
func (this *QSortFilterProxyModel) FilterRegExp() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel12filterRegExpEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:85
// index:0
// Public
// void setFilterRegExp(const class QRegExp &)
func (this *QSortFilterProxyModel) SetFilterRegExp(regExp *QRegExp) {
	var convArg0 = regExp.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:115
// index:1
// Public
// void setFilterRegExp(const class QString &)
func (this *QSortFilterProxyModel) SetFilterRegExp_1(pattern *QString) {
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel15setFilterRegExpERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:87
// index:0
// Public
// int filterKeyColumn()
func (this *QSortFilterProxyModel) FilterKeyColumn() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel15filterKeyColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:88
// index:0
// Public
// void setFilterKeyColumn(int)
func (this *QSortFilterProxyModel) SetFilterKeyColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel18setFilterKeyColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:90
// index:0
// Public
// Qt::CaseSensitivity filterCaseSensitivity()
func (this *QSortFilterProxyModel) FilterCaseSensitivity() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel21filterCaseSensitivityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:91
// index:0
// Public
// void setFilterCaseSensitivity(Qt::CaseSensitivity)
func (this *QSortFilterProxyModel) SetFilterCaseSensitivity(cs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel24setFilterCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:93
// index:0
// Public
// Qt::CaseSensitivity sortCaseSensitivity()
func (this *QSortFilterProxyModel) SortCaseSensitivity() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel19sortCaseSensitivityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:94
// index:0
// Public
// void setSortCaseSensitivity(Qt::CaseSensitivity)
func (this *QSortFilterProxyModel) SetSortCaseSensitivity(cs int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel22setSortCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &cs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:96
// index:0
// Public
// bool isSortLocaleAware()
func (this *QSortFilterProxyModel) IsSortLocaleAware() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel17isSortLocaleAwareEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:97
// index:0
// Public
// void setSortLocaleAware(_Bool)
func (this *QSortFilterProxyModel) SetSortLocaleAware(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel18setSortLocaleAwareEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:99
// index:0
// Public
// int sortColumn()
func (this *QSortFilterProxyModel) SortColumn() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10sortColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:100
// index:0
// Public
// Qt::SortOrder sortOrder()
func (this *QSortFilterProxyModel) SortOrder() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel9sortOrderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:102
// index:0
// Public
// bool dynamicSortFilter()
func (this *QSortFilterProxyModel) DynamicSortFilter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel17dynamicSortFilterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:103
// index:0
// Public
// void setDynamicSortFilter(_Bool)
func (this *QSortFilterProxyModel) SetDynamicSortFilter(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel20setDynamicSortFilterEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:105
// index:0
// Public
// int sortRole()
func (this *QSortFilterProxyModel) SortRole() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8sortRoleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:106
// index:0
// Public
// void setSortRole(int)
func (this *QSortFilterProxyModel) SetSortRole(role int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel11setSortRoleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:108
// index:0
// Public
// int filterRole()
func (this *QSortFilterProxyModel) FilterRole() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10filterRoleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:109
// index:0
// Public
// void setFilterRole(int)
func (this *QSortFilterProxyModel) SetFilterRole(role int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13setFilterRoleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:111
// index:0
// Public
// bool isRecursiveFilteringEnabled()
func (this *QSortFilterProxyModel) IsRecursiveFilteringEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel27isRecursiveFilteringEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:112
// index:0
// Public
// void setRecursiveFilteringEnabled(_Bool)
func (this *QSortFilterProxyModel) SetRecursiveFilteringEnabled(recursive bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel28setRecursiveFilteringEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &recursive)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:116
// index:0
// Public
// void setFilterWildcard(const class QString &)
func (this *QSortFilterProxyModel) SetFilterWildcard(pattern *QString) {
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel17setFilterWildcardERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:117
// index:0
// Public
// void setFilterFixedString(const class QString &)
func (this *QSortFilterProxyModel) SetFilterFixedString(pattern *QString) {
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:118
// index:0
// Public
// void clear()
func (this *QSortFilterProxyModel) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:119
// index:0
// Public
// void invalidate()
func (this *QSortFilterProxyModel) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:122
// index:0
// Protected virtual
// bool filterAcceptsRow(int, const class QModelIndex &)
func (this *QSortFilterProxyModel) FilterAcceptsRow(source_row int, source_parent *QModelIndex) interface{} {
	var convArg1 = source_parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel16filterAcceptsRowEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &source_row, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:123
// index:0
// Protected virtual
// bool filterAcceptsColumn(int, const class QModelIndex &)
func (this *QSortFilterProxyModel) FilterAcceptsColumn(source_column int, source_parent *QModelIndex) interface{} {
	var convArg1 = source_parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel19filterAcceptsColumnEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &source_column, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:124
// index:0
// Protected virtual
// bool lessThan(const class QModelIndex &, const class QModelIndex &)
func (this *QSortFilterProxyModel) LessThan(source_left *QModelIndex, source_right *QModelIndex) interface{} {
	var convArg0 = source_left.GetCthis()
	var convArg1 = source_right.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8lessThanERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:126
// index:0
// Protected
// void filterChanged()
func (this *QSortFilterProxyModel) FilterChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13filterChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:127
// index:0
// Protected
// void invalidateFilter()
func (this *QSortFilterProxyModel) InvalidateFilter() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel16invalidateFilterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:132
// index:0
// Public virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) Index(row int, column int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:133
// index:0
// Public virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QSortFilterProxyModel) Parent(child *QModelIndex) interface{} {
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:134
// index:0
// Public virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) Sibling(row int, column int, idx *QModelIndex) interface{} {
	var convArg2 = idx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:136
// index:0
// Public virtual
// int rowCount(const class QModelIndex &)
func (this *QSortFilterProxyModel) RowCount(parent *QModelIndex) interface{} {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:137
// index:0
// Public virtual
// int columnCount(const class QModelIndex &)
func (this *QSortFilterProxyModel) ColumnCount(parent *QModelIndex) interface{} {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:138
// index:0
// Public virtual
// bool hasChildren(const class QModelIndex &)
func (this *QSortFilterProxyModel) HasChildren(parent *QModelIndex) interface{} {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:140
// index:0
// Public virtual
// QVariant data(const class QModelIndex &, int)
func (this *QSortFilterProxyModel) Data(index *QModelIndex, role int) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:141
// index:0
// Public virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QSortFilterProxyModel) SetData(index *QModelIndex, value *QVariant, role int) interface{} {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:143
// index:0
// Public virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QSortFilterProxyModel) HeaderData(section int, orientation int, role int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:144
// index:0
// Public virtual
// bool setHeaderData(int, Qt::Orientation, const class QVariant &, int)
func (this *QSortFilterProxyModel) SetHeaderData(section int, orientation int, value *QVariant, role int) interface{} {
	var convArg2 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &section, &orientation, convArg2, &role)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:148
// index:0
// Public virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent *QModelIndex) interface{} {
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), data, &action, &row, &column, convArg4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:151
// index:0
// Public virtual
// bool insertRows(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) InsertRows(row int, count int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:152
// index:0
// Public virtual
// bool insertColumns(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) InsertColumns(column int, count int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:153
// index:0
// Public virtual
// bool removeRows(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) RemoveRows(row int, count int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:154
// index:0
// Public virtual
// bool removeColumns(int, int, const class QModelIndex &)
func (this *QSortFilterProxyModel) RemoveColumns(column int, count int, parent *QModelIndex) interface{} {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:156
// index:0
// Public virtual
// void fetchMore(const class QModelIndex &)
func (this *QSortFilterProxyModel) FetchMore(parent *QModelIndex) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:157
// index:0
// Public virtual
// bool canFetchMore(const class QModelIndex &)
func (this *QSortFilterProxyModel) CanFetchMore(parent *QModelIndex) interface{} {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:158
// index:0
// Public virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QSortFilterProxyModel) Flags(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:160
// index:0
// Public virtual
// QModelIndex buddy(const class QModelIndex &)
func (this *QSortFilterProxyModel) Buddy(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5buddyERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:165
// index:0
// Public virtual
// QSize span(const class QModelIndex &)
func (this *QSortFilterProxyModel) Span(index *QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel4spanERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:166
// index:0
// Public virtual
// void sort(int, Qt::SortOrder)
func (this *QSortFilterProxyModel) Sort(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN21QSortFilterProxyModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:168
// index:0
// Public virtual
// QStringList mimeTypes()
func (this *QSortFilterProxyModel) MimeTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel9mimeTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:169
// index:0
// Public virtual
// Qt::DropActions supportedDropActions()
func (this *QSortFilterProxyModel) SupportedDropActions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
