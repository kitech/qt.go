package qtcore

// /usr/include/qt/QtCore/qsortfilterproxymodel.h
// #include <qsortfilterproxymodel.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool filterAcceptsRow(int, const class QModelIndex &)
func (this *QSortFilterProxyModel) InheritFilterAcceptsRow(f func(source_row int, source_parent *QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "filterAcceptsRow", f)
}

// bool filterAcceptsColumn(int, const class QModelIndex &)
func (this *QSortFilterProxyModel) InheritFilterAcceptsColumn(f func(source_column int, source_parent *QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "filterAcceptsColumn", f)
}

// bool lessThan(const class QModelIndex &, const class QModelIndex &)
func (this *QSortFilterProxyModel) InheritLessThan(f func(source_left *QModelIndex, source_right *QModelIndex) bool) {
	qtrt.SetAllInheritCallback(this, "lessThan", f)
}

// void filterChanged()
func (this *QSortFilterProxyModel) InheritFilterChanged(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "filterChanged", f)
}

// void invalidateFilter()
func (this *QSortFilterProxyModel) InheritInvalidateFilter(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "invalidateFilter", f)
}

type QSortFilterProxyModel struct {
	*QAbstractProxyModel
}
type QSortFilterProxyModel_ITF interface {
	QAbstractProxyModel_ITF
	QSortFilterProxyModel_PTR() *QSortFilterProxyModel
}

func (ptr *QSortFilterProxyModel) QSortFilterProxyModel_PTR() *QSortFilterProxyModel { return ptr }

func (this *QSortFilterProxyModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractProxyModel.GetCthis()
	}
}
func (this *QSortFilterProxyModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractProxyModel = NewQAbstractProxyModelFromPointer(cthis)
}
func NewQSortFilterProxyModelFromPointer(cthis unsafe.Pointer) *QSortFilterProxyModel {
	bcthis0 := NewQAbstractProxyModelFromPointer(cthis)
	return &QSortFilterProxyModel{bcthis0}
}
func (*QSortFilterProxyModel) NewFromPointer(cthis unsafe.Pointer) *QSortFilterProxyModel {
	return NewQSortFilterProxyModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QSortFilterProxyModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSortFilterProxyModel(QObject *)
func NewQSortFilterProxyModel(parent QObject_ITF /*777 QObject **/) *QSortFilterProxyModel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSortFilterProxyModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSortFilterProxyModel")
	return gothis
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSortFilterProxyModel(QObject *)
func NewQSortFilterProxyModel__() *QSortFilterProxyModel {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSortFilterProxyModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSortFilterProxyModel")
	return gothis
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSortFilterProxyModel()
func DeleteQSortFilterProxyModel(this *QSortFilterProxyModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSourceModel(QAbstractItemModel *)
func (this *QSortFilterProxyModel) SetSourceModel(sourceModel QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if sourceModel != nil && sourceModel.QAbstractItemModel_PTR() != nil {
		convArg0 = sourceModel.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel14setSourceModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex mapToSource(const QModelIndex &) const
func (this *QSortFilterProxyModel) MapToSource(proxyIndex QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if proxyIndex != nil && proxyIndex.QModelIndex_PTR() != nil {
		convArg0 = proxyIndex.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11mapToSourceERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex mapFromSource(const QModelIndex &) const
func (this *QSortFilterProxyModel) MapFromSource(sourceIndex QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if sourceIndex != nil && sourceIndex.QModelIndex_PTR() != nil {
		convArg0 = sourceIndex.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel13mapFromSourceERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QItemSelection mapSelectionToSource(const QItemSelection &) const
func (this *QSortFilterProxyModel) MapSelectionToSource(proxySelection QItemSelection_ITF) *QItemSelection /*123*/ {
	var convArg0 unsafe.Pointer
	if proxySelection != nil && proxySelection.QItemSelection_PTR() != nil {
		convArg0 = proxySelection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel20mapSelectionToSourceERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QItemSelection mapSelectionFromSource(const QItemSelection &) const
func (this *QSortFilterProxyModel) MapSelectionFromSource(sourceSelection QItemSelection_ITF) *QItemSelection /*123*/ {
	var convArg0 unsafe.Pointer
	if sourceSelection != nil && sourceSelection.QItemSelection_PTR() != nil {
		convArg0 = sourceSelection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel22mapSelectionFromSourceERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegExp filterRegExp() const
func (this *QSortFilterProxyModel) FilterRegExp() *QRegExp /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel12filterRegExpEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegExpFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegExp)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilterRegExp(const QRegExp &)
func (this *QSortFilterProxyModel) SetFilterRegExp(regExp QRegExp_ITF) {
	var convArg0 unsafe.Pointer
	if regExp != nil && regExp.QRegExp_PTR() != nil {
		convArg0 = regExp.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel15setFilterRegExpERK7QRegExp", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:115
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFilterRegExp(const QString &)
func (this *QSortFilterProxyModel) SetFilterRegExp_1(pattern string) {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel15setFilterRegExpERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int filterKeyColumn() const
func (this *QSortFilterProxyModel) FilterKeyColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel15filterKeyColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilterKeyColumn(int)
func (this *QSortFilterProxyModel) SetFilterKeyColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel18setFilterKeyColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity filterCaseSensitivity() const
func (this *QSortFilterProxyModel) FilterCaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel21filterCaseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilterCaseSensitivity(Qt::CaseSensitivity)
func (this *QSortFilterProxyModel) SetFilterCaseSensitivity(cs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel24setFilterCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity sortCaseSensitivity() const
func (this *QSortFilterProxyModel) SortCaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel19sortCaseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortCaseSensitivity(Qt::CaseSensitivity)
func (this *QSortFilterProxyModel) SetSortCaseSensitivity(cs int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel22setSortCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cs)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSortLocaleAware() const
func (this *QSortFilterProxyModel) IsSortLocaleAware() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel17isSortLocaleAwareEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortLocaleAware(_Bool)
func (this *QSortFilterProxyModel) SetSortLocaleAware(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel18setSortLocaleAwareEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int sortColumn() const
func (this *QSortFilterProxyModel) SortColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10sortColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::SortOrder sortOrder() const
func (this *QSortFilterProxyModel) SortOrder() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel9sortOrderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool dynamicSortFilter() const
func (this *QSortFilterProxyModel) DynamicSortFilter() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel17dynamicSortFilterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDynamicSortFilter(_Bool)
func (this *QSortFilterProxyModel) SetDynamicSortFilter(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel20setDynamicSortFilterEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] int sortRole() const
func (this *QSortFilterProxyModel) SortRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8sortRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortRole(int)
func (this *QSortFilterProxyModel) SetSortRole(role int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel11setSortRoleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] int filterRole() const
func (this *QSortFilterProxyModel) FilterRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10filterRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilterRole(int)
func (this *QSortFilterProxyModel) SetFilterRole(role int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13setFilterRoleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRecursiveFilteringEnabled() const
func (this *QSortFilterProxyModel) IsRecursiveFilteringEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel27isRecursiveFilteringEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRecursiveFilteringEnabled(_Bool)
func (this *QSortFilterProxyModel) SetRecursiveFilteringEnabled(recursive bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel28setRecursiveFilteringEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), recursive)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilterWildcard(const QString &)
func (this *QSortFilterProxyModel) SetFilterWildcard(pattern string) {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel17setFilterWildcardERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilterFixedString(const QString &)
func (this *QSortFilterProxyModel) SetFilterFixedString(pattern string) {
	var tmpArg0 = NewQString_5(pattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel20setFilterFixedStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QSortFilterProxyModel) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QSortFilterProxyModel) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:122
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool filterAcceptsRow(int, const QModelIndex &) const
func (this *QSortFilterProxyModel) FilterAcceptsRow(source_row int, source_parent QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if source_parent != nil && source_parent.QModelIndex_PTR() != nil {
		convArg1 = source_parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel16filterAcceptsRowEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), source_row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:123
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool filterAcceptsColumn(int, const QModelIndex &) const
func (this *QSortFilterProxyModel) FilterAcceptsColumn(source_column int, source_parent QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if source_parent != nil && source_parent.QModelIndex_PTR() != nil {
		convArg1 = source_parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel19filterAcceptsColumnEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), source_column, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:124
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool lessThan(const QModelIndex &, const QModelIndex &) const
func (this *QSortFilterProxyModel) LessThan(source_left QModelIndex_ITF, source_right QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if source_left != nil && source_left.QModelIndex_PTR() != nil {
		convArg0 = source_left.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if source_right != nil && source_right.QModelIndex_PTR() != nil {
		convArg1 = source_right.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8lessThanERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:126
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void filterChanged()
func (this *QSortFilterProxyModel) FilterChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13filterChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:127
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void invalidateFilter()
func (this *QSortFilterProxyModel) InvalidateFilter() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel16invalidateFilterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:132
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const
func (this *QSortFilterProxyModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:132
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const
func (this *QSortFilterProxyModel) Index__(row int, column int) *QModelIndex /*123*/ {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:133
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex parent(const QModelIndex &) const
func (this *QSortFilterProxyModel) Parent(child QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if child != nil && child.QModelIndex_PTR() != nil {
		convArg0 = child.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel6parentERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:134
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &) const
func (this *QSortFilterProxyModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if idx != nil && idx.QModelIndex_PTR() != nil {
		convArg2 = idx.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:136
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const
func (this *QSortFilterProxyModel) RowCount(parent QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:136
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const
func (this *QSortFilterProxyModel) RowCount__() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:137
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const
func (this *QSortFilterProxyModel) ColumnCount(parent QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:137
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const
func (this *QSortFilterProxyModel) ColumnCount__() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const
func (this *QSortFilterProxyModel) HasChildren(parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:138
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const
func (this *QSortFilterProxyModel) HasChildren__() bool {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:140
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const
func (this *QSortFilterProxyModel) Data(index QModelIndex_ITF, role int) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:140
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const
func (this *QSortFilterProxyModel) Data__(index QModelIndex_ITF) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:141
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)
func (this *QSortFilterProxyModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:141
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)
func (this *QSortFilterProxyModel) SetData__(index QModelIndex_ITF, value QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	role := 0 /*Qt::EditRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:143
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const
func (this *QSortFilterProxyModel) HeaderData(section int, orientation int, role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:143
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const
func (this *QSortFilterProxyModel) HeaderData__(section int, orientation int) *QVariant /*123*/ {
	// arg: 2, int=Int, =Invalid,
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)
func (this *QSortFilterProxyModel) SetHeaderData(section int, orientation int, value QVariant_ITF, role int) bool {
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)
func (this *QSortFilterProxyModel) SetHeaderData__(section int, orientation int, value QVariant_ITF) bool {
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid,
	role := 0 /*Qt::EditRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:148
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QSortFilterProxyModel) DropMimeData(data QMimeData_ITF /*777 const QMimeData **/, action int, row int, column int, parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg4 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)
func (this *QSortFilterProxyModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)
func (this *QSortFilterProxyModel) InsertRows__(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:152
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)
func (this *QSortFilterProxyModel) InsertColumns(column int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:152
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)
func (this *QSortFilterProxyModel) InsertColumns__(column int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:153
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)
func (this *QSortFilterProxyModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:153
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)
func (this *QSortFilterProxyModel) RemoveRows__(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)
func (this *QSortFilterProxyModel) RemoveColumns(column int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:154
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)
func (this *QSortFilterProxyModel) RemoveColumns__(column int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:156
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fetchMore(const QModelIndex &)
func (this *QSortFilterProxyModel) FetchMore(parent QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel9fetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:157
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canFetchMore(const QModelIndex &) const
func (this *QSortFilterProxyModel) CanFetchMore(parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel12canFetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:158
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &) const
func (this *QSortFilterProxyModel) Flags(index QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:160
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex buddy(const QModelIndex &) const
func (this *QSortFilterProxyModel) Buddy(index QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5buddyERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:161
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QModelIndexList match(const QModelIndex &, int, const QVariant &, int, Qt::MatchFlags) const
func (this *QSortFilterProxyModel) Match(start QModelIndex_ITF, role int, value QVariant_ITF, hits int, flags int) *QModelIndexList /*667*/ {
	var convArg0 unsafe.Pointer
	if start != nil && start.QModelIndex_PTR() != nil {
		convArg0 = start.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, convArg2, hits, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:161
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QModelIndexList match(const QModelIndex &, int, const QVariant &, int, Qt::MatchFlags) const
func (this *QSortFilterProxyModel) Match__(start QModelIndex_ITF, role int, value QVariant_ITF) *QModelIndexList /*667*/ {
	var convArg0 unsafe.Pointer
	if start != nil && start.QModelIndex_PTR() != nil {
		convArg0 = start.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid,
	hits := 1
	// arg: 4, Qt::MatchFlags=Elaborated, Qt::MatchFlags=Typedef, QFlags<Qt::MatchFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, convArg2, hits, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:161
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QModelIndexList match(const QModelIndex &, int, const QVariant &, int, Qt::MatchFlags) const
func (this *QSortFilterProxyModel) Match__1(start QModelIndex_ITF, role int, value QVariant_ITF, hits int) *QModelIndexList /*667*/ {
	var convArg0 unsafe.Pointer
	if start != nil && start.QModelIndex_PTR() != nil {
		convArg0 = start.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	// arg: 4, Qt::MatchFlags=Elaborated, Qt::MatchFlags=Typedef, QFlags<Qt::MatchFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, convArg2, hits, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:165
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize span(const QModelIndex &) const
func (this *QSortFilterProxyModel) Span(index QModelIndex_ITF) *QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel4spanERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:166
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)
func (this *QSortFilterProxyModel) Sort(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:166
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)
func (this *QSortFilterProxyModel) Sort__(column int) {
	// arg: 1, Qt::SortOrder=Elaborated, Qt::SortOrder=Enum,
	order := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN21QSortFilterProxyModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:168
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes() const
func (this *QSortFilterProxyModel) MimeTypes() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qsortfilterproxymodel.h:169
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const
func (this *QSortFilterProxyModel) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QSortFilterProxyModel20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
