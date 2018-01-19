//  header block begin
// /usr/include/qt/QtWidgets/qfilesystemmodel.h
// #include <qfilesystemmodel.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 66
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:60
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFileSystemModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:66
// index:0
// void rootPathChanged(const class QString &)
func (this *QFileSystemModel) RootPathChanged(newPath unsafe.Pointer) {
	// 0: (, const QString & newPath), (newPath)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel15rootPathChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, newPath)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:67
// index:0
// void fileRenamed(const class QString &, const class QString &, const class QString &)
func (this *QFileSystemModel) FileRenamed(path unsafe.Pointer, oldName unsafe.Pointer, newName unsafe.Pointer) {
	// 0: (, const QString & path, const QString & oldName, const QString & newName), (path, oldName, newName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel11fileRenamedERK7QStringS2_S2_", ffiqt.FFI_TYPE_VOID, this.cthis, path, oldName, newName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:68
// index:0
// void directoryLoaded(const class QString &)
func (this *QFileSystemModel) DirectoryLoaded(path unsafe.Pointer) {
	// 0: (, const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel15directoryLoadedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:78
// index:0
// void QFileSystemModel(class QObject *)
func NewQFileSystemModel(parent unsafe.Pointer) *QFileSystemModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QFileSystemModel{cthis}
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:79
// index:0
// virtual
// void ~QFileSystemModel()
func DeleteQFileSystemModel(*QFileSystemModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:81
// index:0
// virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QFileSystemModel) Index(row int, column int, parent unsafe.Pointer) {
	// 0: (, int row, int column, const QModelIndex & parent), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:82
// index:1
// QModelIndex index(const class QString &, int)
func (this *QFileSystemModel) Index_1(path unsafe.Pointer, column int) {
	// 1: (, const QString & path, int column), (path, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel5indexERK7QStringi", ffiqt.FFI_TYPE_VOID, this.cthis, path, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:83
// index:0
// virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QFileSystemModel) Parent(child unsafe.Pointer) {
	// 0: (, const QModelIndex & child), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:85
// index:0
// virtual
// QModelIndex sibling(int, int, const class QModelIndex &)
func (this *QFileSystemModel) Sibling(row int, column int, idx unsafe.Pointer) {
	// 0: (, int row, int column, const QModelIndex & idx), (&row, &column, idx)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, idx)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:86
// index:0
// virtual
// bool hasChildren(const class QModelIndex &)
func (this *QFileSystemModel) HasChildren(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:87
// index:0
// virtual
// bool canFetchMore(const class QModelIndex &)
func (this *QFileSystemModel) CanFetchMore(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel12canFetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:88
// index:0
// virtual
// void fetchMore(const class QModelIndex &)
func (this *QFileSystemModel) FetchMore(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel9fetchMoreERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:90
// index:0
// virtual
// int rowCount(const class QModelIndex &)
func (this *QFileSystemModel) RowCount(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:91
// index:0
// virtual
// int columnCount(const class QModelIndex &)
func (this *QFileSystemModel) ColumnCount(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:93
// index:0
// QVariant myComputer(int)
func (this *QFileSystemModel) MyComputer(role int) {
	// 0: (, int role), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel10myComputerEi", ffiqt.FFI_TYPE_VOID, this.cthis, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:94
// index:0
// virtual
// QVariant data(const class QModelIndex &, int)
func (this *QFileSystemModel) Data(index unsafe.Pointer, role int) {
	// 0: (, const QModelIndex & index, int role), (index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_VOID, this.cthis, index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:95
// index:0
// virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QFileSystemModel) SetData(index unsafe.Pointer, value unsafe.Pointer, role int) {
	// 0: (, const QModelIndex & index, const QVariant & value, int role), (index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.cthis, index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:97
// index:0
// virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QFileSystemModel) HeaderData(section int, orientation int, role int) {
	// 0: (, int section, Qt::Orientation orientation, int role), (&section, &orientation, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_VOID, this.cthis, &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:99
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QFileSystemModel) Flags(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:101
// index:0
// virtual
// void sort(int, Qt::SortOrder)
func (this *QFileSystemModel) Sort(column int, order int) {
	// 0: (, int column, Qt::SortOrder order), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:103
// index:0
// virtual
// QStringList mimeTypes()
func (this *QFileSystemModel) MimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel9mimeTypesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:105
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QFileSystemModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:107
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QFileSystemModel) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:110
// index:0
// QModelIndex setRootPath(const class QString &)
func (this *QFileSystemModel) SetRootPath(path unsafe.Pointer) {
	// 0: (, const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel11setRootPathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:111
// index:0
// QString rootPath()
func (this *QFileSystemModel) RootPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8rootPathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:112
// index:0
// QDir rootDirectory()
func (this *QFileSystemModel) RootDirectory() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel13rootDirectoryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:114
// index:0
// void setIconProvider(class QFileIconProvider *)
func (this *QFileSystemModel) SetIconProvider(provider unsafe.Pointer) {
	// 0: (, QFileIconProvider * provider), (provider)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel15setIconProviderEP17QFileIconProvider", ffiqt.FFI_TYPE_VOID, this.cthis, provider)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:115
// index:0
// QFileIconProvider * iconProvider()
func (this *QFileSystemModel) IconProvider() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel12iconProviderEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:118
// index:0
// QDir::Filters filter()
func (this *QFileSystemModel) Filter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel6filterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:120
// index:0
// void setResolveSymlinks(_Bool)
func (this *QFileSystemModel) SetResolveSymlinks(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel18setResolveSymlinksEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:121
// index:0
// bool resolveSymlinks()
func (this *QFileSystemModel) ResolveSymlinks() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel15resolveSymlinksEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:123
// index:0
// void setReadOnly(_Bool)
func (this *QFileSystemModel) SetReadOnly(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel11setReadOnlyEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:124
// index:0
// bool isReadOnly()
func (this *QFileSystemModel) IsReadOnly() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel10isReadOnlyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:126
// index:0
// void setNameFilterDisables(_Bool)
func (this *QFileSystemModel) SetNameFilterDisables(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel21setNameFilterDisablesEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:127
// index:0
// bool nameFilterDisables()
func (this *QFileSystemModel) NameFilterDisables() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel18nameFilterDisablesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:129
// index:0
// void setNameFilters(const class QStringList &)
func (this *QFileSystemModel) SetNameFilters(filters unsafe.Pointer) {
	// 0: (, const QStringList & filters), (filters)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel14setNameFiltersERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, filters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:130
// index:0
// QStringList nameFilters()
func (this *QFileSystemModel) NameFilters() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel11nameFiltersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:132
// index:0
// QString filePath(const class QModelIndex &)
func (this *QFileSystemModel) FilePath(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8filePathERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:133
// index:0
// bool isDir(const class QModelIndex &)
func (this *QFileSystemModel) IsDir(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel5isDirERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:134
// index:0
// qint64 size(const class QModelIndex &)
func (this *QFileSystemModel) Size(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel4sizeERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:135
// index:0
// QString type(const class QModelIndex &)
func (this *QFileSystemModel) Type(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel4typeERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:136
// index:0
// QDateTime lastModified(const class QModelIndex &)
func (this *QFileSystemModel) LastModified(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel12lastModifiedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:138
// index:0
// QModelIndex mkdir(const class QModelIndex &, const class QString &)
func (this *QFileSystemModel) Mkdir(parent unsafe.Pointer, name unsafe.Pointer) {
	// 0: (, const QModelIndex & parent, const QString & name), (parent, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel5mkdirERK11QModelIndexRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, parent, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:139
// index:0
// bool rmdir(const class QModelIndex &)
func (this *QFileSystemModel) Rmdir(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel5rmdirERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:140
// index:0
// inline
// QString fileName(const class QModelIndex &)
func (this *QFileSystemModel) FileName(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8fileNameERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:141
// index:0
// inline
// QIcon fileIcon(const class QModelIndex &)
func (this *QFileSystemModel) FileIcon(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8fileIconERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:142
// index:0
// QFile::Permissions permissions(const class QModelIndex &)
func (this *QFileSystemModel) Permissions(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel11permissionsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:143
// index:0
// QFileInfo fileInfo(const class QModelIndex &)
func (this *QFileSystemModel) FileInfo(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QFileSystemModel8fileInfoERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfilesystemmodel.h:144
// index:0
// bool remove(const class QModelIndex &)
func (this *QFileSystemModel) Remove(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QFileSystemModel6removeERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

//  body block end
