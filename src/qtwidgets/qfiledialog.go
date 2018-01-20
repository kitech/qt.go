//  header block begin
// /usr/include/qt/QtWidgets/qfiledialog.h
// #include <qfiledialog.h>
// #include <QtWidgets>
package qtwidgets

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
type QFileDialog struct {
	*QDialog
}

func (this *QFileDialog) GetCthis() unsafe.Pointer {
	return this.QDialog.GetCthis()
}
func NewQFileDialogFromPointer(cthis unsafe.Pointer) *QFileDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QFileDialog{bcthis0}
}

// /usr/include/qt/QtWidgets/qfiledialog.h:63
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFileDialog) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:101
// index:0
// Public
// void QFileDialog(class QWidget *, const class QString &, const class QString &, const class QString &)
func NewQFileDialog(parent unsafe.Pointer, caption *qtcore.QString, directory *qtcore.QString, filter *qtcore.QString) *QFileDialog {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg1 = caption.GetCthis()
	var convArg2 = directory.GetCthis()
	var convArg3 = filter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialogC2EP7QWidgetRK7QStringS4_S4_", ffiqt.FFI_TYPE_VOID, cthis, parent, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileDialogFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qfiledialog.h:105
// index:0
// Public virtual
// void ~QFileDialog()
func DeleteQFileDialog(*QFileDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:107
// index:0
// Public
// void setDirectory(const class QString &)
func (this *QFileDialog) SetDirectory(directory *qtcore.QString) {
	var convArg0 = directory.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog12setDirectoryERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:108
// index:1
// Public inline
// void setDirectory(const class QDir &)
func (this *QFileDialog) SetDirectory_1(directory *qtcore.QDir) {
	var convArg0 = directory.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog12setDirectoryERK4QDir", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:109
// index:0
// Public
// QDir directory()
func (this *QFileDialog) Directory() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog9directoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:111
// index:0
// Public
// void setDirectoryUrl(const class QUrl &)
func (this *QFileDialog) SetDirectoryUrl(directory *qtcore.QUrl) {
	var convArg0 = directory.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog15setDirectoryUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:112
// index:0
// Public
// QUrl directoryUrl()
func (this *QFileDialog) DirectoryUrl() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog12directoryUrlEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:114
// index:0
// Public
// void selectFile(const class QString &)
func (this *QFileDialog) SelectFile(filename *qtcore.QString) {
	var convArg0 = filename.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog10selectFileERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:115
// index:0
// Public
// QStringList selectedFiles()
func (this *QFileDialog) SelectedFiles() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog13selectedFilesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:117
// index:0
// Public
// void selectUrl(const class QUrl &)
func (this *QFileDialog) SelectUrl(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog9selectUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:118
// index:0
// Public
// QList<QUrl> selectedUrls()
func (this *QFileDialog) SelectedUrls() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog12selectedUrlsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:120
// index:0
// Public
// void setNameFilterDetailsVisible(_Bool)
func (this *QFileDialog) SetNameFilterDetailsVisible(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog27setNameFilterDetailsVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:121
// index:0
// Public
// bool isNameFilterDetailsVisible()
func (this *QFileDialog) IsNameFilterDetailsVisible() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog26isNameFilterDetailsVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:123
// index:0
// Public
// void setNameFilter(const class QString &)
func (this *QFileDialog) SetNameFilter(filter *qtcore.QString) {
	var convArg0 = filter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog13setNameFilterERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:124
// index:0
// Public
// void setNameFilters(const class QStringList &)
func (this *QFileDialog) SetNameFilters(filters *qtcore.QStringList) {
	var convArg0 = filters.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog14setNameFiltersERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:125
// index:0
// Public
// QStringList nameFilters()
func (this *QFileDialog) NameFilters() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog11nameFiltersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:126
// index:0
// Public
// void selectNameFilter(const class QString &)
func (this *QFileDialog) SelectNameFilter(filter *qtcore.QString) {
	var convArg0 = filter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog16selectNameFilterERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:127
// index:0
// Public
// QString selectedMimeTypeFilter()
func (this *QFileDialog) SelectedMimeTypeFilter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog22selectedMimeTypeFilterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:128
// index:0
// Public
// QString selectedNameFilter()
func (this *QFileDialog) SelectedNameFilter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog18selectedNameFilterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:131
// index:0
// Public
// void setMimeTypeFilters(const class QStringList &)
func (this *QFileDialog) SetMimeTypeFilters(filters *qtcore.QStringList) {
	var convArg0 = filters.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog18setMimeTypeFiltersERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:132
// index:0
// Public
// QStringList mimeTypeFilters()
func (this *QFileDialog) MimeTypeFilters() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog15mimeTypeFiltersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:133
// index:0
// Public
// void selectMimeTypeFilter(const class QString &)
func (this *QFileDialog) SelectMimeTypeFilter(filter *qtcore.QString) {
	var convArg0 = filter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog20selectMimeTypeFilterERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:136
// index:0
// Public
// QDir::Filters filter()
func (this *QFileDialog) Filter() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog6filterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:139
// index:0
// Public
// void setViewMode(enum QFileDialog::ViewMode)
func (this *QFileDialog) SetViewMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog11setViewModeENS_8ViewModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:140
// index:0
// Public
// QFileDialog::ViewMode viewMode()
func (this *QFileDialog) ViewMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog8viewModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:142
// index:0
// Public
// void setFileMode(enum QFileDialog::FileMode)
func (this *QFileDialog) SetFileMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog11setFileModeENS_8FileModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:143
// index:0
// Public
// QFileDialog::FileMode fileMode()
func (this *QFileDialog) FileMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog8fileModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:145
// index:0
// Public
// void setAcceptMode(enum QFileDialog::AcceptMode)
func (this *QFileDialog) SetAcceptMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog13setAcceptModeENS_10AcceptModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:146
// index:0
// Public
// QFileDialog::AcceptMode acceptMode()
func (this *QFileDialog) AcceptMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog10acceptModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:148
// index:0
// Public
// void setReadOnly(_Bool)
func (this *QFileDialog) SetReadOnly(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog11setReadOnlyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:149
// index:0
// Public
// bool isReadOnly()
func (this *QFileDialog) IsReadOnly() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog10isReadOnlyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:151
// index:0
// Public
// void setResolveSymlinks(_Bool)
func (this *QFileDialog) SetResolveSymlinks(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog18setResolveSymlinksEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:152
// index:0
// Public
// bool resolveSymlinks()
func (this *QFileDialog) ResolveSymlinks() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog15resolveSymlinksEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:155
// index:0
// Public
// QList<QUrl> sidebarUrls()
func (this *QFileDialog) SidebarUrls() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog11sidebarUrlsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:157
// index:0
// Public
// QByteArray saveState()
func (this *QFileDialog) SaveState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog9saveStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:158
// index:0
// Public
// bool restoreState(const class QByteArray &)
func (this *QFileDialog) RestoreState(state *qtcore.QByteArray) interface{} {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog12restoreStateERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:160
// index:0
// Public
// void setConfirmOverwrite(_Bool)
func (this *QFileDialog) SetConfirmOverwrite(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog19setConfirmOverwriteEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:161
// index:0
// Public
// bool confirmOverwrite()
func (this *QFileDialog) ConfirmOverwrite() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog16confirmOverwriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:163
// index:0
// Public
// void setDefaultSuffix(const class QString &)
func (this *QFileDialog) SetDefaultSuffix(suffix *qtcore.QString) {
	var convArg0 = suffix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog16setDefaultSuffixERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:164
// index:0
// Public
// QString defaultSuffix()
func (this *QFileDialog) DefaultSuffix() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog13defaultSuffixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:166
// index:0
// Public
// void setHistory(const class QStringList &)
func (this *QFileDialog) SetHistory(paths *qtcore.QStringList) {
	var convArg0 = paths.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog10setHistoryERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:167
// index:0
// Public
// QStringList history()
func (this *QFileDialog) History() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog7historyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:169
// index:0
// Public
// void setItemDelegate(class QAbstractItemDelegate *)
func (this *QFileDialog) SetItemDelegate(delegate unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog15setItemDelegateEP21QAbstractItemDelegate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:170
// index:0
// Public
// QAbstractItemDelegate * itemDelegate()
func (this *QFileDialog) ItemDelegate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog12itemDelegateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:172
// index:0
// Public
// void setIconProvider(class QFileIconProvider *)
func (this *QFileDialog) SetIconProvider(provider unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog15setIconProviderEP17QFileIconProvider", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), provider)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:173
// index:0
// Public
// QFileIconProvider * iconProvider()
func (this *QFileDialog) IconProvider() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog12iconProviderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:175
// index:0
// Public
// void setLabelText(enum QFileDialog::DialogLabel, const class QString &)
func (this *QFileDialog) SetLabelText(label int, text *qtcore.QString) {
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog12setLabelTextENS_11DialogLabelERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &label, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:176
// index:0
// Public
// QString labelText(enum QFileDialog::DialogLabel)
func (this *QFileDialog) LabelText(label int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog9labelTextENS_11DialogLabelE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &label)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:178
// index:0
// Public
// void setSupportedSchemes(const class QStringList &)
func (this *QFileDialog) SetSupportedSchemes(schemes *qtcore.QStringList) {
	var convArg0 = schemes.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog19setSupportedSchemesERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:179
// index:0
// Public
// QStringList supportedSchemes()
func (this *QFileDialog) SupportedSchemes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog16supportedSchemesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:182
// index:0
// Public
// void setProxyModel(class QAbstractProxyModel *)
func (this *QFileDialog) SetProxyModel(model unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog13setProxyModelEP19QAbstractProxyModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:183
// index:0
// Public
// QAbstractProxyModel * proxyModel()
func (this *QFileDialog) ProxyModel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog10proxyModelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:186
// index:0
// Public
// void setOption(enum QFileDialog::Option, _Bool)
func (this *QFileDialog) SetOption(option int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog9setOptionENS_6OptionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:187
// index:0
// Public
// bool testOption(enum QFileDialog::Option)
func (this *QFileDialog) TestOption(option int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog10testOptionENS_6OptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:189
// index:0
// Public
// QFileDialog::Options options()
func (this *QFileDialog) Options() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog7optionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:192
// index:0
// Public
// void open(class QObject *, const char *)
func (this *QFileDialog) Open(receiver unsafe.Pointer, member string) {
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), receiver, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:193
// index:0
// Public virtual
// void setVisible(_Bool)
func (this *QFileDialog) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:196
// index:0
// Public
// void fileSelected(const class QString &)
func (this *QFileDialog) FileSelected(file *qtcore.QString) {
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog12fileSelectedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:197
// index:0
// Public
// void filesSelected(const class QStringList &)
func (this *QFileDialog) FilesSelected(files *qtcore.QStringList) {
	var convArg0 = files.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog13filesSelectedERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:198
// index:0
// Public
// void currentChanged(const class QString &)
func (this *QFileDialog) CurrentChanged(path *qtcore.QString) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog14currentChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:199
// index:0
// Public
// void directoryEntered(const class QString &)
func (this *QFileDialog) DirectoryEntered(directory *qtcore.QString) {
	var convArg0 = directory.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog16directoryEnteredERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:201
// index:0
// Public
// void urlSelected(const class QUrl &)
func (this *QFileDialog) UrlSelected(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog11urlSelectedERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:203
// index:0
// Public
// void currentUrlChanged(const class QUrl &)
func (this *QFileDialog) CurrentUrlChanged(url *qtcore.QUrl) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog17currentUrlChangedERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:204
// index:0
// Public
// void directoryUrlEntered(const class QUrl &)
func (this *QFileDialog) DirectoryUrlEntered(directory *qtcore.QUrl) {
	var convArg0 = directory.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog19directoryUrlEnteredERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:206
// index:0
// Public
// void filterSelected(const class QString &)
func (this *QFileDialog) FilterSelected(filter *qtcore.QString) {
	var convArg0 = filter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog14filterSelectedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:269
// index:0
// Protected virtual
// void done(int)
func (this *QFileDialog) Done(result int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog4doneEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &result)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:270
// index:0
// Protected virtual
// void accept()
func (this *QFileDialog) Accept() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog6acceptEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:271
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QFileDialog) ChangeEvent(e unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

//  body block end
