package qtwidgets

// /usr/include/qt/QtWidgets/qfilesystemmodel.h
// #include <qfilesystemmodel.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 77
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// void timerEvent(class QTimerEvent *)
func (this *QFileSystemModel) InheritTimerEvent(f func(event *qtcore.QTimerEvent /*777 QTimerEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "timerEvent", f)
}

// bool event(class QEvent *)
func (this *QFileSystemModel) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QFileSystemModel struct {
	*qtcore.QAbstractItemModel
}

func (this *QFileSystemModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemModel.GetCthis()
	}
}
func (this *QFileSystemModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemModel = qtcore.NewQAbstractItemModelFromPointer(cthis)
}
func NewQFileSystemModelFromPointer(cthis unsafe.Pointer) *QFileSystemModel {
	bcthis0 := qtcore.NewQAbstractItemModelFromPointer(cthis)
	return &QFileSystemModel{bcthis0}
}
func (*QFileSystemModel) NewFromPointer(cthis unsafe.Pointer) *QFileSystemModel {
	return NewQFileSystemModelFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFileSystemModel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rootPathChanged(const QString &)
func (this *QFileSystemModel) RootPathChanged(newPath string) {
	var tmpArg0 = qtcore.NewQString_5(newPath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel15rootPathChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fileRenamed(const QString &, const QString &, const QString &)
func (this *QFileSystemModel) FileRenamed(path string, oldName string, newName string) {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(oldName)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(newName)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel11fileRenamedERK7QStringS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void directoryLoaded(const QString &)
func (this *QFileSystemModel) DirectoryLoaded(path string) {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel15directoryLoadedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileSystemModel(QObject *)
func NewQFileSystemModel(parent *qtcore.QObject /*777 QObject **/) *QFileSystemModel {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileSystemModelFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFileSystemModel()
func DeleteQFileSystemModel(this *QFileSystemModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &)
func (this *QFileSystemModel) Index(row int, column int, parent *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:82
// index:1
// Public Visibility=Default Availability=Available
// [24] QModelIndex index(const QString &, int)
func (this *QFileSystemModel) Index_1(path string, column int) *qtcore.QModelIndex /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel5indexERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex parent(const QModelIndex &)
func (this *QFileSystemModel) Parent(child *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg0 = child.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel6parentERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &)
func (this *QFileSystemModel) Sibling(row int, column int, idx *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg2 = idx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &)
func (this *QFileSystemModel) HasChildren(parent *qtcore.QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canFetchMore(const QModelIndex &)
func (this *QFileSystemModel) CanFetchMore(parent *qtcore.QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel12canFetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fetchMore(const QModelIndex &)
func (this *QFileSystemModel) FetchMore(parent *qtcore.QModelIndex) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel9fetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &)
func (this *QFileSystemModel) RowCount(parent *qtcore.QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &)
func (this *QFileSystemModel) ColumnCount(parent *qtcore.QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:93
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant myComputer(int)
func (this *QFileSystemModel) MyComputer(role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel10myComputerEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int)
func (this *QFileSystemModel) Data(index *qtcore.QModelIndex, role int) *qtcore.QVariant /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)
func (this *QFileSystemModel) SetData(index *qtcore.QModelIndex, value *qtcore.QVariant, role int) bool {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int)
func (this *QFileSystemModel) HeaderData(section int, orientation int, role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &)
func (this *QFileSystemModel) Flags(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:101
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)
func (this *QFileSystemModel) Sort(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes()
func (this *QFileSystemModel) MimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QFileSystemModel) DropMimeData(data *qtcore.QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *qtcore.QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:107
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions()
func (this *QFileSystemModel) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:110
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex setRootPath(const QString &)
func (this *QFileSystemModel) SetRootPath(path string) *qtcore.QModelIndex /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel11setRootPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QString rootPath()
func (this *QFileSystemModel) RootPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel8rootPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QDir rootDirectory()
func (this *QFileSystemModel) RootDirectory() *qtcore.QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel13rootDirectoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDir)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconProvider(QFileIconProvider *)
func (this *QFileSystemModel) SetIconProvider(provider *QFileIconProvider /*777 QFileIconProvider **/) {
	var convArg0 = provider.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel15setIconProviderEP17QFileIconProvider", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QFileIconProvider * iconProvider()
func (this *QFileSystemModel) IconProvider() *QFileIconProvider /*777 QFileIconProvider **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel12iconProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQFileIconProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilter(QDir::Filters)
func (this *QFileSystemModel) SetFilter(filters int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel9setFilterE6QFlagsIN4QDir6FilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:118
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::Filters filter()
func (this *QFileSystemModel) Filter() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel6filterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolveSymlinks(_Bool)
func (this *QFileSystemModel) SetResolveSymlinks(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel18setResolveSymlinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool resolveSymlinks()
func (this *QFileSystemModel) ResolveSymlinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel15resolveSymlinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(_Bool)
func (this *QFileSystemModel) SetReadOnly(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel11setReadOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:124
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly()
func (this *QFileSystemModel) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilterDisables(_Bool)
func (this *QFileSystemModel) SetNameFilterDisables(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel21setNameFilterDisablesEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:127
// index:0
// Public Visibility=Default Availability=Available
// [1] bool nameFilterDisables()
func (this *QFileSystemModel) NameFilterDisables() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel18nameFilterDisablesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilters(const QStringList &)
func (this *QFileSystemModel) SetNameFilters(filters *qtcore.QStringList) {
	var convArg0 = filters.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel14setNameFiltersERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList nameFilters()
func (this *QFileSystemModel) NameFilters() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel11nameFiltersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath(const QModelIndex &)
func (this *QFileSystemModel) FilePath(index *qtcore.QModelIndex) string {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel8filePathERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDir(const QModelIndex &)
func (this *QFileSystemModel) IsDir(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel5isDirERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 size(const QModelIndex &)
func (this *QFileSystemModel) Size(index *qtcore.QModelIndex) int64 {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel4sizeERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QString type(const QModelIndex &)
func (this *QFileSystemModel) Type(index *qtcore.QModelIndex) string {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel4typeERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:136
// index:0
// Public Visibility=Default Availability=Available
// [8] QDateTime lastModified(const QModelIndex &)
func (this *QFileSystemModel) LastModified(index *qtcore.QModelIndex) *qtcore.QDateTime /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel12lastModifiedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDateTime)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:138
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex mkdir(const QModelIndex &, const QString &)
func (this *QFileSystemModel) Mkdir(parent *qtcore.QModelIndex, name string) *qtcore.QModelIndex /*123*/ {
	var convArg0 = parent.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel5mkdirERK11QModelIndexRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:139
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rmdir(const QModelIndex &)
func (this *QFileSystemModel) Rmdir(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel5rmdirERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString fileName(const QModelIndex &)
func (this *QFileSystemModel) FileName(index *qtcore.QModelIndex) string {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel8fileNameERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:141
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QIcon fileIcon(const QModelIndex &)
func (this *QFileSystemModel) FileIcon(index *qtcore.QModelIndex) *qtgui.QIcon /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel8fileIconERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:142
// index:0
// Public Visibility=Default Availability=Available
// [4] QFile::Permissions permissions(const QModelIndex &)
func (this *QFileSystemModel) Permissions(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel11permissionsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QFileInfo fileInfo(const QModelIndex &)
func (this *QFileSystemModel) FileInfo(index *qtcore.QModelIndex) *qtcore.QFileInfo /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QFileSystemModel8fileInfoERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQFileInfo)
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:144
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove(const QModelIndex &)
func (this *QFileSystemModel) Remove(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel6removeERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:148
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void timerEvent(QTimerEvent *)
func (this *QFileSystemModel) TimerEvent(event *qtcore.QTimerEvent /*777 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel10timerEventEP11QTimerEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:149
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QFileSystemModel) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QFileSystemModel5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

type QFileSystemModel__Roles = int

const QFileSystemModel__FileIconRole QFileSystemModel__Roles = 1
const QFileSystemModel__FilePathRole QFileSystemModel__Roles = 257
const QFileSystemModel__FileNameRole QFileSystemModel__Roles = 258
const QFileSystemModel__FilePermissions QFileSystemModel__Roles = 259

//  body block end
