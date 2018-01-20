//  header block begin
// /usr/include/qt/QtWidgets/qdirmodel.h
// #include <qdirmodel.h>
// #include <QtWidgets>
package qtwidgets

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
	*qtcore.QAbstractItemModel
}

func (this *QDirModel) GetCthis() unsafe.Pointer {
	return this.QAbstractItemModel.GetCthis()
}

// /usr/include/qt/QtWidgets/qdirmodel.h:56
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDirModel) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:68
// index:0
// void QDirModel(const class QStringList &, class QDir::Filters, class QDir::SortFlags, class QObject *)
func NewQDirModel(nameFilters unsafe.Pointer, filters int, sort int, parent unsafe.Pointer) *QDirModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModelC2ERK11QStringList6QFlagsIN4QDir6FilterEES3_INS4_8SortFlagEEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, nameFilters, &filters, &sort, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirModelFromPointer(cthis)
	return gothis
}
func NewQDirModelFromPointer(cthis unsafe.Pointer) *QDirModel {
	bcthis0 := qtcore.NewQAbstractItemModelFromPointer(cthis)
	return &QDirModel{bcthis0}
}

// /usr/include/qt/QtWidgets/qdirmodel.h:70
// index:1
// void QDirModel(class QObject *)
func NewQDirModel_1(parent unsafe.Pointer) *QDirModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdirmodel.h:136
// index:2
// void QDirModel(class QDirModelPrivate &, class QObject *)
func NewQDirModel_2(arg0 unsafe.Pointer, parent unsafe.Pointer) *QDirModel {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModelC2ER16QDirModelPrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDirModelFromPointer(cthis)
	return gothis
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
	// 0: (, row int, column int, parent const QModelIndex &), (&row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:118
// index:1
// QModelIndex index(const class QString &, int)
func (this *QDirModel) Index_1(path unsafe.Pointer, column int) {
	// 1: (, path const QString &, column int), (path, &column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5indexERK7QStringi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:74
// index:0
// virtual
// QModelIndex parent(const class QModelIndex &)
func (this *QDirModel) Parent(child unsafe.Pointer) {
	// 0: (, child const QModelIndex &), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:76
// index:0
// virtual
// int rowCount(const class QModelIndex &)
func (this *QDirModel) RowCount(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:77
// index:0
// virtual
// int columnCount(const class QModelIndex &)
func (this *QDirModel) ColumnCount(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:79
// index:0
// virtual
// QVariant data(const class QModelIndex &, int)
func (this *QDirModel) Data(index unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, role int), (index, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:80
// index:0
// virtual
// bool setData(const class QModelIndex &, const class QVariant &, int)
func (this *QDirModel) SetData(index unsafe.Pointer, value unsafe.Pointer, role int) {
	// 0: (, index const QModelIndex &, value const QVariant &, role int), (index, value, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, value, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:82
// index:0
// virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QDirModel) HeaderData(section int, orientation int, role int) {
	// 0: (, section int, orientation Qt::Orientation, role int), (&section, &orientation, &role)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &section, &orientation, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:84
// index:0
// virtual
// bool hasChildren(const class QModelIndex &)
func (this *QDirModel) HasChildren(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:85
// index:0
// virtual
// Qt::ItemFlags flags(const class QModelIndex &)
func (this *QDirModel) Flags(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:87
// index:0
// virtual
// void sort(int, Qt::SortOrder)
func (this *QDirModel) Sort(column int, order int) {
	// 0: (, column int, order Qt::SortOrder), (&column, &order)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, &order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:89
// index:0
// virtual
// QStringList mimeTypes()
func (this *QDirModel) MimeTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel9mimeTypesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:91
// index:0
// virtual
// bool dropMimeData(const class QMimeData *, Qt::DropAction, int, int, const class QModelIndex &)
func (this *QDirModel) DropMimeData(data unsafe.Pointer, action int, row int, column int, parent unsafe.Pointer) {
	// 0: (, data const QMimeData *, action Qt::DropAction, row int, column int, parent const QModelIndex &), (data, &action, &row, &column, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &action, &row, &column, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:93
// index:0
// virtual
// Qt::DropActions supportedDropActions()
func (this *QDirModel) SupportedDropActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel20supportedDropActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:97
// index:0
// void setIconProvider(class QFileIconProvider *)
func (this *QDirModel) SetIconProvider(provider unsafe.Pointer) {
	// 0: (, provider QFileIconProvider *), (provider)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel15setIconProviderEP17QFileIconProvider", ffiqt.FFI_TYPE_VOID, this.GetCthis(), provider)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:98
// index:0
// QFileIconProvider * iconProvider()
func (this *QDirModel) IconProvider() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel12iconProviderEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:100
// index:0
// void setNameFilters(const class QStringList &)
func (this *QDirModel) SetNameFilters(filters unsafe.Pointer) {
	// 0: (, filters const QStringList &), (filters)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel14setNameFiltersERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:101
// index:0
// QStringList nameFilters()
func (this *QDirModel) NameFilters() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel11nameFiltersEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:103
// index:0
// void setFilter(class QDir::Filters)
func (this *QDirModel) SetFilter(filters int) {
	// 0: (, QFlags<QDir::Filter> filters), (&filters)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel9setFilterE6QFlagsIN4QDir6FilterEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &filters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:104
// index:0
// QDir::Filters filter()
func (this *QDirModel) Filter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel6filterEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:106
// index:0
// void setSorting(class QDir::SortFlags)
func (this *QDirModel) SetSorting(sort int) {
	// 0: (, QFlags<QDir::SortFlag> sort), (&sort)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel10setSortingE6QFlagsIN4QDir8SortFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &sort)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:107
// index:0
// QDir::SortFlags sorting()
func (this *QDirModel) Sorting() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel7sortingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:109
// index:0
// void setResolveSymlinks(_Bool)
func (this *QDirModel) SetResolveSymlinks(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel18setResolveSymlinksEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:110
// index:0
// bool resolveSymlinks()
func (this *QDirModel) ResolveSymlinks() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel15resolveSymlinksEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:112
// index:0
// void setReadOnly(_Bool)
func (this *QDirModel) SetReadOnly(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel11setReadOnlyEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:113
// index:0
// bool isReadOnly()
func (this *QDirModel) IsReadOnly() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel10isReadOnlyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:115
// index:0
// void setLazyChildCount(_Bool)
func (this *QDirModel) SetLazyChildCount(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel17setLazyChildCountEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:116
// index:0
// bool lazyChildCount()
func (this *QDirModel) LazyChildCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel14lazyChildCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:120
// index:0
// bool isDir(const class QModelIndex &)
func (this *QDirModel) IsDir(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel5isDirERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:121
// index:0
// QModelIndex mkdir(const class QModelIndex &, const class QString &)
func (this *QDirModel) Mkdir(parent unsafe.Pointer, name unsafe.Pointer) {
	// 0: (, parent const QModelIndex &, name const QString &), (parent, name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel5mkdirERK11QModelIndexRK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:122
// index:0
// bool rmdir(const class QModelIndex &)
func (this *QDirModel) Rmdir(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel5rmdirERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:123
// index:0
// bool remove(const class QModelIndex &)
func (this *QDirModel) Remove(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel6removeERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:125
// index:0
// QString filePath(const class QModelIndex &)
func (this *QDirModel) FilePath(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8filePathERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:126
// index:0
// QString fileName(const class QModelIndex &)
func (this *QDirModel) FileName(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8fileNameERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:127
// index:0
// QIcon fileIcon(const class QModelIndex &)
func (this *QDirModel) FileIcon(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8fileIconERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:128
// index:0
// QFileInfo fileInfo(const class QModelIndex &)
func (this *QDirModel) FileInfo(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QDirModel8fileInfoERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:133
// index:0
// void refresh(const class QModelIndex &)
func (this *QDirModel) Refresh(parent unsafe.Pointer) {
	// 0: (, parent const QModelIndex &), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QDirModel7refreshERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

//  body block end
