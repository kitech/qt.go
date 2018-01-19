//  header block begin
// /usr/include/qt/QtWidgets/qdirmodel.h
// #include <qdirmodel.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QDirModel struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qdirmodel.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDirModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:70
// index:0
// void QDirModel(class QObject *)
func NewQDirModel(parent unsafe.Pointer) *QDirModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QDirModel{cthis}
}

// /usr/include/qt/QtWidgets/qdirmodel.h:71
// index:0
// virtual
// void ~QDirModel()
func DeleteQDirModel(*QDirModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:73
// index:0
// virtual
// QModelIndex index(int, int, const class QModelIndex &)
func (this *QDirModel) Index(row int, column int, parent unsafe.Pointer) {
	// 0: (, int row, int column, const QModelIndex & parent), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:118
// index:1
// QModelIndex index(const class QString &, int)
func (this *QDirModel) Index_1(path unsafe.Pointer, column int) {
	// 1: (, const QString & path, int column), (path, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5indexERK7QStringi", ffiqt.FFI_TYPE_VOID, this.cthis, path, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:74
// index:0
// virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QDirModel) Parent(child unsafe.Pointer) {
	// 0: (, const QModelIndex & child), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:76
// index:0
// virtual
// int rowCount(const class QModelIndex &)
func (this *QDirModel) RowCount(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:77
// index:0
// virtual
// int columnCount(const class QModelIndex &)
func (this *QDirModel) ColumnCount(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:79
// index:0
// virtual
// QVariant data(const class QModelIndex &, int)
func (this *QDirModel) Data(index unsafe.Pointer, role int) {
	// 0: (, const QModelIndex & index, int role), (index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_VOID, this.cthis, index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:80
// index:0
// virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QDirModel) SetData(index unsafe.Pointer, value unsafe.Pointer, role int) {
	// 0: (, const QModelIndex & index, const QVariant & value, int role), (index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.cthis, index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:82
// index:0
// virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QDirModel) HeaderData(section int, orientation int, role int) {
	// 0: (, int section, Qt::Orientation orientation, int role), (&section, &orientation, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_VOID, this.cthis, &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:84
// index:0
// virtual
// bool hasChildren(const class QModelIndex &)
func (this *QDirModel) HasChildren(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:85
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QDirModel) Flags(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:87
// index:0
// virtual
// void sort(int, Qt::SortOrder)
func (this *QDirModel) Sort(column int, order int) {
	// 0: (, int column, Qt::SortOrder order), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.cthis, &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:89
// index:0
// virtual
// QStringList mimeTypes()
func (this *QDirModel) MimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel9mimeTypesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:91
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QDirModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:93
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QDirModel) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:97
// index:0
// void setIconProvider(class QFileIconProvider *)
func (this *QDirModel) SetIconProvider(provider unsafe.Pointer) {
	// 0: (, QFileIconProvider * provider), (provider)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel15setIconProviderEP17QFileIconProvider", ffiqt.FFI_TYPE_VOID, this.cthis, provider)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:98
// index:0
// QFileIconProvider * iconProvider()
func (this *QDirModel) IconProvider() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel12iconProviderEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:100
// index:0
// void setNameFilters(const class QStringList &)
func (this *QDirModel) SetNameFilters(filters unsafe.Pointer) {
	// 0: (, const QStringList & filters), (filters)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel14setNameFiltersERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, filters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:101
// index:0
// QStringList nameFilters()
func (this *QDirModel) NameFilters() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel11nameFiltersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:104
// index:0
// QDir::Filters filter()
func (this *QDirModel) Filter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel6filterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:107
// index:0
// QDir::SortFlags sorting()
func (this *QDirModel) Sorting() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel7sortingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:109
// index:0
// void setResolveSymlinks(_Bool)
func (this *QDirModel) SetResolveSymlinks(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel18setResolveSymlinksEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:110
// index:0
// bool resolveSymlinks()
func (this *QDirModel) ResolveSymlinks() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel15resolveSymlinksEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:112
// index:0
// void setReadOnly(_Bool)
func (this *QDirModel) SetReadOnly(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel11setReadOnlyEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:113
// index:0
// bool isReadOnly()
func (this *QDirModel) IsReadOnly() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel10isReadOnlyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:115
// index:0
// void setLazyChildCount(_Bool)
func (this *QDirModel) SetLazyChildCount(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel17setLazyChildCountEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:116
// index:0
// bool lazyChildCount()
func (this *QDirModel) LazyChildCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel14lazyChildCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:120
// index:0
// bool isDir(const class QModelIndex &)
func (this *QDirModel) IsDir(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5isDirERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:121
// index:0
// QModelIndex mkdir(const class QModelIndex &, const class QString &)
func (this *QDirModel) Mkdir(parent unsafe.Pointer, name unsafe.Pointer) {
	// 0: (, const QModelIndex & parent, const QString & name), (parent, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel5mkdirERK11QModelIndexRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, parent, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:122
// index:0
// bool rmdir(const class QModelIndex &)
func (this *QDirModel) Rmdir(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel5rmdirERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:123
// index:0
// bool remove(const class QModelIndex &)
func (this *QDirModel) Remove(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel6removeERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:125
// index:0
// QString filePath(const class QModelIndex &)
func (this *QDirModel) FilePath(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8filePathERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:126
// index:0
// QString fileName(const class QModelIndex &)
func (this *QDirModel) FileName(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8fileNameERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:127
// index:0
// QIcon fileIcon(const class QModelIndex &)
func (this *QDirModel) FileIcon(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8fileIconERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:128
// index:0
// QFileInfo fileInfo(const class QModelIndex &)
func (this *QDirModel) FileInfo(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8fileInfoERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:133
// index:0
// void refresh(const class QModelIndex &)
func (this *QDirModel) Refresh(parent unsafe.Pointer) {
	// 0: (, const QModelIndex & parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel7refreshERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

//  body block end
