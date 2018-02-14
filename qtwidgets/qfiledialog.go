package qtwidgets

// /usr/include/qt/QtWidgets/qfiledialog.h
// #include <qfiledialog.h>
// #include <QtWidgets>

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
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void done(int)
func (this *QFileDialog) InheritDone(f func(result int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "done", f)
}

// void accept()
func (this *QFileDialog) InheritAccept(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "accept", f)
}

// void changeEvent(class QEvent *)
func (this *QFileDialog) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

type QFileDialog struct {
	*QDialog
}
type QFileDialog_ITF interface {
	QDialog_ITF
	QFileDialog_PTR() *QFileDialog
}

func (ptr *QFileDialog) QFileDialog_PTR() *QFileDialog { return ptr }

func (this *QFileDialog) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QFileDialog) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQFileDialogFromPointer(cthis unsafe.Pointer) *QFileDialog {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QFileDialog{bcthis0}
}
func (*QFileDialog) NewFromPointer(cthis unsafe.Pointer) *QFileDialog {
	return NewQFileDialogFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFileDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfiledialog.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileDialog(QWidget *, Qt::WindowFlags)
func NewQFileDialog(parent QWidget_ITF /*777 QWidget **/, f int) *QFileDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialogC2EP7QWidget6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qfiledialog.h:101
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileDialog(QWidget *, const QString &, const QString &, const QString &)
func NewQFileDialog_1(parent QWidget_ITF /*777 QWidget **/, caption string, directory string, filter string) *QFileDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(directory)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(filter)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialogC2EP7QWidgetRK7QStringS4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qfiledialog.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFileDialog()
func DeleteQFileDialog(this *QFileDialog) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialogD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDirectory(const QString &)
func (this *QFileDialog) SetDirectory(directory string) {
	var tmpArg0 = qtcore.NewQString_5(directory)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog12setDirectoryERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:108
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setDirectory(const QDir &)
func (this *QFileDialog) SetDirectory_1(directory qtcore.QDir_ITF) {
	var convArg0 unsafe.Pointer
	if directory != nil && directory.QDir_PTR() != nil {
		convArg0 = directory.QDir_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog12setDirectoryERK4QDir", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QDir directory()
func (this *QFileDialog) Directory() *qtcore.QDir /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog9directoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQDirFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQDir)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDirectoryUrl(const QUrl &)
func (this *QFileDialog) SetDirectoryUrl(directory qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if directory != nil && directory.QUrl_PTR() != nil {
		convArg0 = directory.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15setDirectoryUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl directoryUrl()
func (this *QFileDialog) DirectoryUrl() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog12directoryUrlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectFile(const QString &)
func (this *QFileDialog) SelectFile(filename string) {
	var tmpArg0 = qtcore.NewQString_5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog10selectFileERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList selectedFiles()
func (this *QFileDialog) SelectedFiles() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog13selectedFilesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectUrl(const QUrl &)
func (this *QFileDialog) SelectUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog9selectUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilterDetailsVisible(_Bool)
func (this *QFileDialog) SetNameFilterDetailsVisible(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog27setNameFilterDetailsVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNameFilterDetailsVisible()
func (this *QFileDialog) IsNameFilterDetailsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog26isNameFilterDetailsVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilter(const QString &)
func (this *QFileDialog) SetNameFilter(filter string) {
	var tmpArg0 = qtcore.NewQString_5(filter)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog13setNameFilterERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilters(const QStringList &)
func (this *QFileDialog) SetNameFilters(filters qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if filters != nil && filters.QStringList_PTR() != nil {
		convArg0 = filters.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14setNameFiltersERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList nameFilters()
func (this *QFileDialog) NameFilters() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog11nameFiltersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectNameFilter(const QString &)
func (this *QFileDialog) SelectNameFilter(filter string) {
	var tmpArg0 = qtcore.NewQString_5(filter)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16selectNameFilterERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedMimeTypeFilter()
func (this *QFileDialog) SelectedMimeTypeFilter() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog22selectedMimeTypeFilterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedNameFilter()
func (this *QFileDialog) SelectedNameFilter() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog18selectedNameFilterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMimeTypeFilters(const QStringList &)
func (this *QFileDialog) SetMimeTypeFilters(filters qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if filters != nil && filters.QStringList_PTR() != nil {
		convArg0 = filters.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog18setMimeTypeFiltersERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList mimeTypeFilters()
func (this *QFileDialog) MimeTypeFilters() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog15mimeTypeFiltersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void selectMimeTypeFilter(const QString &)
func (this *QFileDialog) SelectMimeTypeFilter(filter string) {
	var tmpArg0 = qtcore.NewQString_5(filter)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog20selectMimeTypeFilterERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::Filters filter()
func (this *QFileDialog) Filter() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog6filterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilter(QDir::Filters)
func (this *QFileDialog) SetFilter(filters int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog9setFilterE6QFlagsIN4QDir6FilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewMode(enum QFileDialog::ViewMode)
func (this *QFileDialog) SetViewMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog11setViewModeENS_8ViewModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] QFileDialog::ViewMode viewMode()
func (this *QFileDialog) ViewMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog8viewModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileMode(enum QFileDialog::FileMode)
func (this *QFileDialog) SetFileMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog11setFileModeENS_8FileModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] QFileDialog::FileMode fileMode()
func (this *QFileDialog) FileMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog8fileModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptMode(enum QFileDialog::AcceptMode)
func (this *QFileDialog) SetAcceptMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog13setAcceptModeENS_10AcceptModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] QFileDialog::AcceptMode acceptMode()
func (this *QFileDialog) AcceptMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog10acceptModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(_Bool)
func (this *QFileDialog) SetReadOnly(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog11setReadOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:149
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly()
func (this *QFileDialog) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolveSymlinks(_Bool)
func (this *QFileDialog) SetResolveSymlinks(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog18setResolveSymlinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:152
// index:0
// Public Visibility=Default Availability=Available
// [1] bool resolveSymlinks()
func (this *QFileDialog) ResolveSymlinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog15resolveSymlinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:157
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveState()
func (this *QFileDialog) SaveState() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog9saveStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:158
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreState(const QByteArray &)
func (this *QFileDialog) RestoreState(state qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if state != nil && state.QByteArray_PTR() != nil {
		convArg0 = state.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog12restoreStateERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setConfirmOverwrite(_Bool)
func (this *QFileDialog) SetConfirmOverwrite(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog19setConfirmOverwriteEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:161
// index:0
// Public Visibility=Default Availability=Available
// [1] bool confirmOverwrite()
func (this *QFileDialog) ConfirmOverwrite() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog16confirmOverwriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultSuffix(const QString &)
func (this *QFileDialog) SetDefaultSuffix(suffix string) {
	var tmpArg0 = qtcore.NewQString_5(suffix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16setDefaultSuffixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:164
// index:0
// Public Visibility=Default Availability=Available
// [8] QString defaultSuffix()
func (this *QFileDialog) DefaultSuffix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog13defaultSuffixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHistory(const QStringList &)
func (this *QFileDialog) SetHistory(paths qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if paths != nil && paths.QStringList_PTR() != nil {
		convArg0 = paths.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog10setHistoryERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList history()
func (this *QFileDialog) History() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog7historyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemDelegate(QAbstractItemDelegate *)
func (this *QFileDialog) SetItemDelegate(delegate QAbstractItemDelegate_ITF /*777 QAbstractItemDelegate **/) {
	var convArg0 unsafe.Pointer
	if delegate != nil && delegate.QAbstractItemDelegate_PTR() != nil {
		convArg0 = delegate.QAbstractItemDelegate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15setItemDelegateEP21QAbstractItemDelegate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:170
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegate()
func (this *QFileDialog) ItemDelegate() *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog12itemDelegateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfiledialog.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconProvider(QFileIconProvider *)
func (this *QFileDialog) SetIconProvider(provider QFileIconProvider_ITF /*777 QFileIconProvider **/) {
	var convArg0 unsafe.Pointer
	if provider != nil && provider.QFileIconProvider_PTR() != nil {
		convArg0 = provider.QFileIconProvider_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15setIconProviderEP17QFileIconProvider", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:173
// index:0
// Public Visibility=Default Availability=Available
// [8] QFileIconProvider * iconProvider()
func (this *QFileDialog) IconProvider() *QFileIconProvider /*777 QFileIconProvider **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog12iconProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQFileIconProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfiledialog.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLabelText(enum QFileDialog::DialogLabel, const QString &)
func (this *QFileDialog) SetLabelText(label int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog12setLabelTextENS_11DialogLabelERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), label, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:176
// index:0
// Public Visibility=Default Availability=Available
// [8] QString labelText(enum QFileDialog::DialogLabel)
func (this *QFileDialog) LabelText(label int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog9labelTextENS_11DialogLabelE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), label)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSupportedSchemes(const QStringList &)
func (this *QFileDialog) SetSupportedSchemes(schemes qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if schemes != nil && schemes.QStringList_PTR() != nil {
		convArg0 = schemes.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog19setSupportedSchemesERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList supportedSchemes()
func (this *QFileDialog) SupportedSchemes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog16supportedSchemesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProxyModel(QAbstractProxyModel *)
func (this *QFileDialog) SetProxyModel(model qtcore.QAbstractProxyModel_ITF /*777 QAbstractProxyModel **/) {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractProxyModel_PTR() != nil {
		convArg0 = model.QAbstractProxyModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog13setProxyModelEP19QAbstractProxyModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:183
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractProxyModel * proxyModel()
func (this *QFileDialog) ProxyModel() *qtcore.QAbstractProxyModel /*777 QAbstractProxyModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog10proxyModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQAbstractProxyModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfiledialog.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(enum QFileDialog::Option, _Bool)
func (this *QFileDialog) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog9setOptionENS_6OptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:187
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(enum QFileDialog::Option)
func (this *QFileDialog) TestOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog10testOptionENS_6OptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QFileDialog::Options)
func (this *QFileDialog) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog10setOptionsE6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:189
// index:0
// Public Visibility=Default Availability=Available
// [4] QFileDialog::Options options()
func (this *QFileDialog) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open(QObject *, const char *)
func (this *QFileDialog) Open(receiver qtcore.QObject_ITF /*777 QObject **/, member string) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog4openEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:193
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QFileDialog) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fileSelected(const QString &)
func (this *QFileDialog) FileSelected(file string) {
	var tmpArg0 = qtcore.NewQString_5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog12fileSelectedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void filesSelected(const QStringList &)
func (this *QFileDialog) FilesSelected(files qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if files != nil && files.QStringList_PTR() != nil {
		convArg0 = files.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog13filesSelectedERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(const QString &)
func (this *QFileDialog) CurrentChanged(path string) {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14currentChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void directoryEntered(const QString &)
func (this *QFileDialog) DirectoryEntered(directory string) {
	var tmpArg0 = qtcore.NewQString_5(directory)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16directoryEnteredERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void urlSelected(const QUrl &)
func (this *QFileDialog) UrlSelected(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog11urlSelectedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:203
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentUrlChanged(const QUrl &)
func (this *QFileDialog) CurrentUrlChanged(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog17currentUrlChangedERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void directoryUrlEntered(const QUrl &)
func (this *QFileDialog) DirectoryUrlEntered(directory qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if directory != nil && directory.QUrl_PTR() != nil {
		convArg0 = directory.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog19directoryUrlEnteredERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void filterSelected(const QString &)
func (this *QFileDialog) FilterSelected(filter string) {
	var tmpArg0 = qtcore.NewQString_5(filter)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14filterSelectedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getOpenFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)
func (this *QFileDialog) GetOpenFileName(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string, selectedFilter string, options int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(dir)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(filter)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(selectedFilter)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QFileDialog_GetOpenFileName(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string, selectedFilter string, options int) string {
	var nilthis *QFileDialog
	rv := nilthis.GetOpenFileName(parent, caption, dir, filter, selectedFilter, options)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:217
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getOpenFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)
func (this *QFileDialog) GetOpenFileUrl(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string, options int, supportedSchemes qtcore.QStringList_ITF) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if dir != nil && dir.QUrl_PTR() != nil {
		convArg2 = dir.QUrl_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString_5(filter)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(selectedFilter)
	var convArg4 = tmpArg4.GetCthis()
	var convArg6 unsafe.Pointer
	if supportedSchemes != nil && supportedSchemes.QStringList_PTR() != nil {
		convArg6 = supportedSchemes.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getOpenFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}
func QFileDialog_GetOpenFileUrl(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string, options int, supportedSchemes qtcore.QStringList_ITF) *qtcore.QUrl /*123*/ {
	var nilthis *QFileDialog
	rv := nilthis.GetOpenFileUrl(parent, caption, dir, filter, selectedFilter, options, supportedSchemes)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:225
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getSaveFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)
func (this *QFileDialog) GetSaveFileName(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string, selectedFilter string, options int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(dir)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(filter)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(selectedFilter)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getSaveFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QFileDialog_GetSaveFileName(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string, selectedFilter string, options int) string {
	var nilthis *QFileDialog
	rv := nilthis.GetSaveFileName(parent, caption, dir, filter, selectedFilter, options)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:232
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getSaveFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)
func (this *QFileDialog) GetSaveFileUrl(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string, options int, supportedSchemes qtcore.QStringList_ITF) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if dir != nil && dir.QUrl_PTR() != nil {
		convArg2 = dir.QUrl_PTR().GetCthis()
	}
	var tmpArg3 = qtcore.NewQString_5(filter)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(selectedFilter)
	var convArg4 = tmpArg4.GetCthis()
	var convArg6 unsafe.Pointer
	if supportedSchemes != nil && supportedSchemes.QStringList_PTR() != nil {
		convArg6 = supportedSchemes.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getSaveFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}
func QFileDialog_GetSaveFileUrl(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string, options int, supportedSchemes qtcore.QStringList_ITF) *qtcore.QUrl /*123*/ {
	var nilthis *QFileDialog
	rv := nilthis.GetSaveFileUrl(parent, caption, dir, filter, selectedFilter, options, supportedSchemes)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:240
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getExistingDirectory(QWidget *, const QString &, const QString &, QFileDialog::Options)
func (this *QFileDialog) GetExistingDirectory(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, options int) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(dir)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog20getExistingDirectoryEP7QWidgetRK7QStringS4_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QFileDialog_GetExistingDirectory(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, options int) string {
	var nilthis *QFileDialog
	rv := nilthis.GetExistingDirectory(parent, caption, dir, options)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:245
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getExistingDirectoryUrl(QWidget *, const QString &, const QUrl &, QFileDialog::Options, const QStringList &)
func (this *QFileDialog) GetExistingDirectoryUrl(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, options int, supportedSchemes qtcore.QStringList_ITF) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if dir != nil && dir.QUrl_PTR() != nil {
		convArg2 = dir.QUrl_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if supportedSchemes != nil && supportedSchemes.QStringList_PTR() != nil {
		convArg4 = supportedSchemes.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog23getExistingDirectoryUrlEP7QWidgetRK7QStringRK4QUrl6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options, convArg4)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}
func QFileDialog_GetExistingDirectoryUrl(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, options int, supportedSchemes qtcore.QStringList_ITF) *qtcore.QUrl /*123*/ {
	var nilthis *QFileDialog
	rv := nilthis.GetExistingDirectoryUrl(parent, caption, dir, options, supportedSchemes)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:251
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList getOpenFileNames(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)
func (this *QFileDialog) GetOpenFileNames(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string, selectedFilter string, options int) *qtcore.QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(dir)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = qtcore.NewQString_5(filter)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = qtcore.NewQString_5(selectedFilter)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16getOpenFileNamesEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QFileDialog_GetOpenFileNames(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string, selectedFilter string, options int) *qtcore.QStringList /*123*/ {
	var nilthis *QFileDialog
	rv := nilthis.GetOpenFileNames(parent, caption, dir, filter, selectedFilter, options)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:269
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void done(int)
func (this *QFileDialog) Done(result int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), result)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:270
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void accept()
func (this *QFileDialog) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:271
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QFileDialog) ChangeEvent(e qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QFileDialog__ViewMode = int

const QFileDialog__Detail QFileDialog__ViewMode = 0
const QFileDialog__List QFileDialog__ViewMode = 1

type QFileDialog__FileMode = int

const QFileDialog__AnyFile QFileDialog__FileMode = 0
const QFileDialog__ExistingFile QFileDialog__FileMode = 1
const QFileDialog__Directory QFileDialog__FileMode = 2
const QFileDialog__ExistingFiles QFileDialog__FileMode = 3
const QFileDialog__DirectoryOnly QFileDialog__FileMode = 4

type QFileDialog__AcceptMode = int

const QFileDialog__AcceptOpen QFileDialog__AcceptMode = 0
const QFileDialog__AcceptSave QFileDialog__AcceptMode = 1

type QFileDialog__DialogLabel = int

const QFileDialog__LookIn QFileDialog__DialogLabel = 0
const QFileDialog__FileName QFileDialog__DialogLabel = 1
const QFileDialog__FileType QFileDialog__DialogLabel = 2
const QFileDialog__Accept QFileDialog__DialogLabel = 3
const QFileDialog__Reject QFileDialog__DialogLabel = 4

type QFileDialog__Option = int

const QFileDialog__ShowDirsOnly QFileDialog__Option = 1
const QFileDialog__DontResolveSymlinks QFileDialog__Option = 2
const QFileDialog__DontConfirmOverwrite QFileDialog__Option = 4
const QFileDialog__DontUseSheet QFileDialog__Option = 8
const QFileDialog__DontUseNativeDialog QFileDialog__Option = 16
const QFileDialog__ReadOnly QFileDialog__Option = 32
const QFileDialog__HideNameFilterDetails QFileDialog__Option = 64
const QFileDialog__DontUseCustomDirectoryIcons QFileDialog__Option = 128

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
