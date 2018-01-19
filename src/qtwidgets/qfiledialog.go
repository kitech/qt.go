//  header block begin
// /usr/include/qt/QtWidgets/qfiledialog.h
// #include <qfiledialog.h>
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
type QFileDialog struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qfiledialog.h:63
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFileDialog) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:101
// index:0
// void QFileDialog(class QWidget *, const class QString &, const class QString &, const class QString &)
func NewQFileDialog(parent unsafe.Pointer, caption unsafe.Pointer, directory unsafe.Pointer, filter unsafe.Pointer) *QFileDialog {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialogC2EP7QWidgetRK7QStringS4_S4_", ffiqt.FFI_TYPE_VOID, cthis, parent, caption, directory, filter)
	gopp.ErrPrint(err, rv)
	return &QFileDialog{cthis}
}

// /usr/include/qt/QtWidgets/qfiledialog.h:105
// index:0
// virtual
// void ~QFileDialog()
func DeleteQFileDialog(*QFileDialog) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialogD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:107
// index:0
// void setDirectory(const class QString &)
func (this *QFileDialog) SetDirectory(directory unsafe.Pointer) {
	// 0: (, const QString & directory), (directory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog12setDirectoryERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, directory)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:108
// index:1
// inline
// void setDirectory(const class QDir &)
func (this *QFileDialog) SetDirectory_1(directory unsafe.Pointer) {
	// 1: (, const QDir & directory), (directory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog12setDirectoryERK4QDir", ffiqt.FFI_TYPE_VOID, this.cthis, directory)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:109
// index:0
// QDir directory()
func (this *QFileDialog) Directory() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog9directoryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:111
// index:0
// void setDirectoryUrl(const class QUrl &)
func (this *QFileDialog) SetDirectoryUrl(directory unsafe.Pointer) {
	// 0: (, const QUrl & directory), (directory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog15setDirectoryUrlERK4QUrl", ffiqt.FFI_TYPE_VOID, this.cthis, directory)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:112
// index:0
// QUrl directoryUrl()
func (this *QFileDialog) DirectoryUrl() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog12directoryUrlEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:114
// index:0
// void selectFile(const class QString &)
func (this *QFileDialog) SelectFile(filename unsafe.Pointer) {
	// 0: (, const QString & filename), (filename)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog10selectFileERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, filename)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:115
// index:0
// QStringList selectedFiles()
func (this *QFileDialog) SelectedFiles() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog13selectedFilesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:117
// index:0
// void selectUrl(const class QUrl &)
func (this *QFileDialog) SelectUrl(url unsafe.Pointer) {
	// 0: (, const QUrl & url), (url)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog9selectUrlERK4QUrl", ffiqt.FFI_TYPE_VOID, this.cthis, url)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:118
// index:0
// QList<QUrl> selectedUrls()
func (this *QFileDialog) SelectedUrls() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog12selectedUrlsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:120
// index:0
// void setNameFilterDetailsVisible(_Bool)
func (this *QFileDialog) SetNameFilterDetailsVisible(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog27setNameFilterDetailsVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:121
// index:0
// bool isNameFilterDetailsVisible()
func (this *QFileDialog) IsNameFilterDetailsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog26isNameFilterDetailsVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:123
// index:0
// void setNameFilter(const class QString &)
func (this *QFileDialog) SetNameFilter(filter unsafe.Pointer) {
	// 0: (, const QString & filter), (filter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog13setNameFilterERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, filter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:124
// index:0
// void setNameFilters(const class QStringList &)
func (this *QFileDialog) SetNameFilters(filters unsafe.Pointer) {
	// 0: (, const QStringList & filters), (filters)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog14setNameFiltersERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, filters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:125
// index:0
// QStringList nameFilters()
func (this *QFileDialog) NameFilters() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog11nameFiltersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:126
// index:0
// void selectNameFilter(const class QString &)
func (this *QFileDialog) SelectNameFilter(filter unsafe.Pointer) {
	// 0: (, const QString & filter), (filter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog16selectNameFilterERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, filter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:127
// index:0
// QString selectedMimeTypeFilter()
func (this *QFileDialog) SelectedMimeTypeFilter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog22selectedMimeTypeFilterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:128
// index:0
// QString selectedNameFilter()
func (this *QFileDialog) SelectedNameFilter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog18selectedNameFilterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:131
// index:0
// void setMimeTypeFilters(const class QStringList &)
func (this *QFileDialog) SetMimeTypeFilters(filters unsafe.Pointer) {
	// 0: (, const QStringList & filters), (filters)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog18setMimeTypeFiltersERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, filters)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:132
// index:0
// QStringList mimeTypeFilters()
func (this *QFileDialog) MimeTypeFilters() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog15mimeTypeFiltersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:133
// index:0
// void selectMimeTypeFilter(const class QString &)
func (this *QFileDialog) SelectMimeTypeFilter(filter unsafe.Pointer) {
	// 0: (, const QString & filter), (filter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog20selectMimeTypeFilterERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, filter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:136
// index:0
// QDir::Filters filter()
func (this *QFileDialog) Filter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog6filterEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:139
// index:0
// void setViewMode(enum QFileDialog::ViewMode)
func (this *QFileDialog) SetViewMode(mode int) {
	// 0: (, QFileDialog::ViewMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog11setViewModeENS_8ViewModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:140
// index:0
// QFileDialog::ViewMode viewMode()
func (this *QFileDialog) ViewMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog8viewModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:142
// index:0
// void setFileMode(enum QFileDialog::FileMode)
func (this *QFileDialog) SetFileMode(mode int) {
	// 0: (, QFileDialog::FileMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog11setFileModeENS_8FileModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:143
// index:0
// QFileDialog::FileMode fileMode()
func (this *QFileDialog) FileMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog8fileModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:145
// index:0
// void setAcceptMode(enum QFileDialog::AcceptMode)
func (this *QFileDialog) SetAcceptMode(mode int) {
	// 0: (, QFileDialog::AcceptMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog13setAcceptModeENS_10AcceptModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:146
// index:0
// QFileDialog::AcceptMode acceptMode()
func (this *QFileDialog) AcceptMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog10acceptModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:148
// index:0
// void setReadOnly(_Bool)
func (this *QFileDialog) SetReadOnly(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog11setReadOnlyEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:149
// index:0
// bool isReadOnly()
func (this *QFileDialog) IsReadOnly() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog10isReadOnlyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:151
// index:0
// void setResolveSymlinks(_Bool)
func (this *QFileDialog) SetResolveSymlinks(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog18setResolveSymlinksEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:152
// index:0
// bool resolveSymlinks()
func (this *QFileDialog) ResolveSymlinks() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog15resolveSymlinksEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:155
// index:0
// QList<QUrl> sidebarUrls()
func (this *QFileDialog) SidebarUrls() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog11sidebarUrlsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:157
// index:0
// QByteArray saveState()
func (this *QFileDialog) SaveState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog9saveStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:158
// index:0
// bool restoreState(const class QByteArray &)
func (this *QFileDialog) RestoreState(state unsafe.Pointer) {
	// 0: (, const QByteArray & state), (state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog12restoreStateERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:160
// index:0
// void setConfirmOverwrite(_Bool)
func (this *QFileDialog) SetConfirmOverwrite(enabled bool) {
	// 0: (, bool enabled), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog19setConfirmOverwriteEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:161
// index:0
// bool confirmOverwrite()
func (this *QFileDialog) ConfirmOverwrite() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog16confirmOverwriteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:163
// index:0
// void setDefaultSuffix(const class QString &)
func (this *QFileDialog) SetDefaultSuffix(suffix unsafe.Pointer) {
	// 0: (, const QString & suffix), (suffix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog16setDefaultSuffixERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, suffix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:164
// index:0
// QString defaultSuffix()
func (this *QFileDialog) DefaultSuffix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog13defaultSuffixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:166
// index:0
// void setHistory(const class QStringList &)
func (this *QFileDialog) SetHistory(paths unsafe.Pointer) {
	// 0: (, const QStringList & paths), (paths)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog10setHistoryERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, paths)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:167
// index:0
// QStringList history()
func (this *QFileDialog) History() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog7historyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:169
// index:0
// void setItemDelegate(class QAbstractItemDelegate *)
func (this *QFileDialog) SetItemDelegate(delegate unsafe.Pointer) {
	// 0: (, QAbstractItemDelegate * delegate), (delegate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog15setItemDelegateEP21QAbstractItemDelegate", ffiqt.FFI_TYPE_VOID, this.cthis, delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:170
// index:0
// QAbstractItemDelegate * itemDelegate()
func (this *QFileDialog) ItemDelegate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog12itemDelegateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:172
// index:0
// void setIconProvider(class QFileIconProvider *)
func (this *QFileDialog) SetIconProvider(provider unsafe.Pointer) {
	// 0: (, QFileIconProvider * provider), (provider)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog15setIconProviderEP17QFileIconProvider", ffiqt.FFI_TYPE_VOID, this.cthis, provider)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:173
// index:0
// QFileIconProvider * iconProvider()
func (this *QFileDialog) IconProvider() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog12iconProviderEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:175
// index:0
// void setLabelText(enum QFileDialog::DialogLabel, const class QString &)
func (this *QFileDialog) SetLabelText(label int, text unsafe.Pointer) {
	// 0: (, QFileDialog::DialogLabel label, const QString & text), (&label, text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog12setLabelTextENS_11DialogLabelERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &label, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:176
// index:0
// QString labelText(enum QFileDialog::DialogLabel)
func (this *QFileDialog) LabelText(label int) {
	// 0: (, QFileDialog::DialogLabel label), (&label)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog9labelTextENS_11DialogLabelE", ffiqt.FFI_TYPE_VOID, this.cthis, &label)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:178
// index:0
// void setSupportedSchemes(const class QStringList &)
func (this *QFileDialog) SetSupportedSchemes(schemes unsafe.Pointer) {
	// 0: (, const QStringList & schemes), (schemes)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog19setSupportedSchemesERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, schemes)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:179
// index:0
// QStringList supportedSchemes()
func (this *QFileDialog) SupportedSchemes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog16supportedSchemesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:182
// index:0
// void setProxyModel(class QAbstractProxyModel *)
func (this *QFileDialog) SetProxyModel(model unsafe.Pointer) {
	// 0: (, QAbstractProxyModel * model), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog13setProxyModelEP19QAbstractProxyModel", ffiqt.FFI_TYPE_VOID, this.cthis, model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:183
// index:0
// QAbstractProxyModel * proxyModel()
func (this *QFileDialog) ProxyModel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog10proxyModelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:186
// index:0
// void setOption(enum QFileDialog::Option, _Bool)
func (this *QFileDialog) SetOption(option int, on bool) {
	// 0: (, QFileDialog::Option option, bool on), (&option, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog9setOptionENS_6OptionEb", ffiqt.FFI_TYPE_VOID, this.cthis, &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:187
// index:0
// bool testOption(enum QFileDialog::Option)
func (this *QFileDialog) TestOption(option int) {
	// 0: (, QFileDialog::Option option), (&option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog10testOptionENS_6OptionE", ffiqt.FFI_TYPE_VOID, this.cthis, &option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:189
// index:0
// QFileDialog::Options options()
func (this *QFileDialog) Options() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QFileDialog7optionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:192
// index:0
// void open(class QObject *, const char *)
func (this *QFileDialog) Open(receiver unsafe.Pointer, member unsafe.Pointer) {
	// 0: (, QObject * receiver, const char * member), (receiver, member)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog4openEP7QObjectPKc", ffiqt.FFI_TYPE_VOID, this.cthis, receiver, member)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:193
// index:0
// virtual
// void setVisible(_Bool)
func (this *QFileDialog) SetVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:196
// index:0
// void fileSelected(const class QString &)
func (this *QFileDialog) FileSelected(file unsafe.Pointer) {
	// 0: (, const QString & file), (file)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog12fileSelectedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, file)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:197
// index:0
// void filesSelected(const class QStringList &)
func (this *QFileDialog) FilesSelected(files unsafe.Pointer) {
	// 0: (, const QStringList & files), (files)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog13filesSelectedERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, files)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:198
// index:0
// void currentChanged(const class QString &)
func (this *QFileDialog) CurrentChanged(path unsafe.Pointer) {
	// 0: (, const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog14currentChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:199
// index:0
// void directoryEntered(const class QString &)
func (this *QFileDialog) DirectoryEntered(directory unsafe.Pointer) {
	// 0: (, const QString & directory), (directory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog16directoryEnteredERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, directory)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:201
// index:0
// void urlSelected(const class QUrl &)
func (this *QFileDialog) UrlSelected(url unsafe.Pointer) {
	// 0: (, const QUrl & url), (url)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog11urlSelectedERK4QUrl", ffiqt.FFI_TYPE_VOID, this.cthis, url)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:203
// index:0
// void currentUrlChanged(const class QUrl &)
func (this *QFileDialog) CurrentUrlChanged(url unsafe.Pointer) {
	// 0: (, const QUrl & url), (url)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog17currentUrlChangedERK4QUrl", ffiqt.FFI_TYPE_VOID, this.cthis, url)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:204
// index:0
// void directoryUrlEntered(const class QUrl &)
func (this *QFileDialog) DirectoryUrlEntered(directory unsafe.Pointer) {
	// 0: (, const QUrl & directory), (directory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog19directoryUrlEnteredERK4QUrl", ffiqt.FFI_TYPE_VOID, this.cthis, directory)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:206
// index:0
// void filterSelected(const class QString &)
func (this *QFileDialog) FilterSelected(filter unsafe.Pointer) {
	// 0: (, const QString & filter), (filter)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QFileDialog14filterSelectedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, filter)
	gopp.ErrPrint(err, rv)
}

//  body block end
