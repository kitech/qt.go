package qtwidgets

// /usr/include/qt/QtWidgets/qfilesystemmodel.h
// #include <qfilesystemmodel.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 62
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
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
func NewQFileSystemModelFromPointer(cthis unsafe.Pointer) *QFileSystemModel {
	bcthis0 := qtcore.NewQAbstractItemModelFromPointer(cthis)
	return &QFileSystemModel{bcthis0}
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:60
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFileSystemModel) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:66
// index:0
// Public
// void rootPathChanged(const class QString &)
func (this *QFileSystemModel) RootPathChanged(newPath *qtcore.QString) {
	var convArg0 = newPath.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel15rootPathChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:67
// index:0
// Public
// void fileRenamed(const class QString &, const class QString &, const class QString &)
func (this *QFileSystemModel) FileRenamed(path *qtcore.QString, oldName *qtcore.QString, newName *qtcore.QString) {
	var convArg0 = path.GetCthis()
	var convArg1 = oldName.GetCthis()
	var convArg2 = newName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel11fileRenamedERK7QStringS2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:68
// index:0
// Public
// void directoryLoaded(const class QString &)
func (this *QFileSystemModel) DirectoryLoaded(path *qtcore.QString) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel15directoryLoadedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:78
// index:0
// Public
// void QFileSystemModel(class QObject *)
func NewQFileSystemModel(parent *qtcore.QObject /*444 QObject **/) *QFileSystemModel {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileSystemModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:79
// index:0
// Public virtual
// void ~QFileSystemModel()
func DeleteQFileSystemModel(*QFileSystemModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:81
// index:0
// Public virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QFileSystemModel) Index(row int, column int, parent *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:82
// index:1
// Public
// QModelIndex index(const class QString &, int)
func (this *QFileSystemModel) Index_1(path *qtcore.QString, column int) *qtcore.QModelIndex /*123*/ {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel5indexERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:83
// index:0
// Public virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QFileSystemModel) Parent(child *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:85
// index:0
// Public virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QFileSystemModel) Sibling(row int, column int, idx *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg2 = idx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, &column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:86
// index:0
// Public virtual
// bool hasChildren(const class QModelIndex &)
func (this *QFileSystemModel) HasChildren(parent *qtcore.QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:87
// index:0
// Public virtual
// bool canFetchMore(const class QModelIndex &)
func (this *QFileSystemModel) CanFetchMore(parent *qtcore.QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel12canFetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:88
// index:0
// Public virtual
// void fetchMore(const class QModelIndex &)
func (this *QFileSystemModel) FetchMore(parent *qtcore.QModelIndex) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel9fetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:90
// index:0
// Public virtual
// int rowCount(const class QModelIndex &)
func (this *QFileSystemModel) RowCount(parent *qtcore.QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:91
// index:0
// Public virtual
// int columnCount(const class QModelIndex &)
func (this *QFileSystemModel) ColumnCount(parent *qtcore.QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:93
// index:0
// Public
// QVariant myComputer(int)
func (this *QFileSystemModel) MyComputer(role int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel10myComputerEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:94
// index:0
// Public virtual
// QVariant data(const class QModelIndex &, int)
func (this *QFileSystemModel) Data(index *qtcore.QModelIndex, role int) *qtcore.QVariant /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:95
// index:0
// Public virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QFileSystemModel) SetData(index *qtcore.QModelIndex, value *qtcore.QVariant, role int) bool {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &role)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:97
// index:0
// Public virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QFileSystemModel) HeaderData(section int, orientation int, role int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:99
// index:0
// Public virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QFileSystemModel) Flags(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:101
// index:0
// Public virtual
// void sort(int, Qt::SortOrder)
func (this *QFileSystemModel) Sort(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:105
// index:0
// Public virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QFileSystemModel) DropMimeData(data *qtcore.QMimeData /*444 const QMimeData **/, action int, row int, column int, parent *qtcore.QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &action, &row, &column, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:107
// index:0
// Public virtual
// Qt::DropActions supportedDropActions()
func (this *QFileSystemModel) SupportedDropActions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:110
// index:0
// Public
// QModelIndex setRootPath(const class QString &)
func (this *QFileSystemModel) SetRootPath(path *qtcore.QString) *qtcore.QModelIndex /*123*/ {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel11setRootPathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:111
// index:0
// Public
// QString rootPath()
func (this *QFileSystemModel) RootPath() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8rootPathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:112
// index:0
// Public
// QDir rootDirectory()
func (this *QFileSystemModel) RootDirectory() *qtcore.QDir /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel13rootDirectoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:114
// index:0
// Public
// void setIconProvider(class QFileIconProvider *)
func (this *QFileSystemModel) SetIconProvider(provider *QFileIconProvider /*444 QFileIconProvider **/) {
	var convArg0 = provider.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel15setIconProviderEP17QFileIconProvider", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:115
// index:0
// Public
// QFileIconProvider * iconProvider()
func (this *QFileSystemModel) IconProvider() *QFileIconProvider /*444 QFileIconProvider **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel12iconProviderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQFileIconProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:118
// index:0
// Public
// QDir::Filters filter()
func (this *QFileSystemModel) Filter() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel6filterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:120
// index:0
// Public
// void setResolveSymlinks(_Bool)
func (this *QFileSystemModel) SetResolveSymlinks(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel18setResolveSymlinksEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:121
// index:0
// Public
// bool resolveSymlinks()
func (this *QFileSystemModel) ResolveSymlinks() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel15resolveSymlinksEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:123
// index:0
// Public
// void setReadOnly(_Bool)
func (this *QFileSystemModel) SetReadOnly(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel11setReadOnlyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:124
// index:0
// Public
// bool isReadOnly()
func (this *QFileSystemModel) IsReadOnly() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel10isReadOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:126
// index:0
// Public
// void setNameFilterDisables(_Bool)
func (this *QFileSystemModel) SetNameFilterDisables(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel21setNameFilterDisablesEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:127
// index:0
// Public
// bool nameFilterDisables()
func (this *QFileSystemModel) NameFilterDisables() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel18nameFilterDisablesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:129
// index:0
// Public
// void setNameFilters(const class QStringList &)
func (this *QFileSystemModel) SetNameFilters(filters *qtcore.QStringList) {
	var convArg0 = filters.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel14setNameFiltersERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:132
// index:0
// Public
// QString filePath(const class QModelIndex &)
func (this *QFileSystemModel) FilePath(index *qtcore.QModelIndex) *qtcore.QString /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8filePathERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:133
// index:0
// Public
// bool isDir(const class QModelIndex &)
func (this *QFileSystemModel) IsDir(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel5isDirERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:134
// index:0
// Public
// qint64 size(const class QModelIndex &)
func (this *QFileSystemModel) Size(index *qtcore.QModelIndex) int64 {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel4sizeERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:135
// index:0
// Public
// QString type(const class QModelIndex &)
func (this *QFileSystemModel) Type(index *qtcore.QModelIndex) *qtcore.QString /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel4typeERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:136
// index:0
// Public
// QDateTime lastModified(const class QModelIndex &)
func (this *QFileSystemModel) LastModified(index *qtcore.QModelIndex) *qtcore.QDateTime /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel12lastModifiedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQDateTimeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:138
// index:0
// Public
// QModelIndex mkdir(const class QModelIndex &, const class QString &)
func (this *QFileSystemModel) Mkdir(parent *qtcore.QModelIndex, name *qtcore.QString) *qtcore.QModelIndex /*123*/ {
	var convArg0 = parent.GetCthis()
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel5mkdirERK11QModelIndexRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:139
// index:0
// Public
// bool rmdir(const class QModelIndex &)
func (this *QFileSystemModel) Rmdir(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel5rmdirERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:140
// index:0
// Public inline
// QString fileName(const class QModelIndex &)
func (this *QFileSystemModel) FileName(index *qtcore.QModelIndex) *qtcore.QString /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8fileNameERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:141
// index:0
// Public inline
// QIcon fileIcon(const class QModelIndex &)
func (this *QFileSystemModel) FileIcon(index *qtcore.QModelIndex) *qtgui.QIcon /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8fileIconERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:142
// index:0
// Public
// QFile::Permissions permissions(const class QModelIndex &)
func (this *QFileSystemModel) Permissions(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel11permissionsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:143
// index:0
// Public
// QFileInfo fileInfo(const class QModelIndex &)
func (this *QFileSystemModel) FileInfo(index *qtcore.QModelIndex) *qtcore.QFileInfo /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8fileInfoERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:144
// index:0
// Public
// bool remove(const class QModelIndex &)
func (this *QFileSystemModel) Remove(index *qtcore.QModelIndex) bool {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel6removeERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:148
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QFileSystemModel) TimerEvent(event *qtcore.QTimerEvent /*444 QTimerEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:149
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QFileSystemModel) Event(event *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
