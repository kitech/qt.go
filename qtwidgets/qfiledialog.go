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

// void changeEvent(QEvent *)
func (this *QFileDialog) InheritChangeEvent(f func(e *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

/*

 */
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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QFileDialog) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfiledialog.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileDialog(QWidget *, Qt::WindowFlags)

/*
Constructs a file dialog with the given parent and widget flags.
*/
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

/*
Constructs a file dialog with the given parent and widget flags.
*/
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

// /usr/include/qt/QtWidgets/qfiledialog.h:101
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileDialog(QWidget *, const QString &, const QString &, const QString &)

/*
Constructs a file dialog with the given parent and widget flags.
*/
func NewQFileDialog_1_() *QFileDialog {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialogC2EP7QWidgetRK7QStringS4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qfiledialog.h:101
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileDialog(QWidget *, const QString &, const QString &, const QString &)

/*
Constructs a file dialog with the given parent and widget flags.
*/
func NewQFileDialog_1_1(parent QWidget_ITF /*777 QWidget **/) *QFileDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialogC2EP7QWidgetRK7QStringS4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qfiledialog.h:101
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileDialog(QWidget *, const QString &, const QString &, const QString &)

/*
Constructs a file dialog with the given parent and widget flags.
*/
func NewQFileDialog_1_2(parent QWidget_ITF /*777 QWidget **/, caption string) *QFileDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialogC2EP7QWidgetRK7QStringS4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileDialogFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileDialog")
	return gothis
}

// /usr/include/qt/QtWidgets/qfiledialog.h:101
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileDialog(QWidget *, const QString &, const QString &, const QString &)

/*
Constructs a file dialog with the given parent and widget flags.
*/
func NewQFileDialog_1_3(parent QWidget_ITF /*777 QWidget **/, caption string, directory string) *QFileDialog {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(directory)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
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

/*

 */
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

/*
Sets the file dialog's current directory.

Note: On iOS, if you set directory to QStandardPaths::standardLocations(QStandardPaths::PicturesLocation).last(), a native image picker dialog will be used for accessing the user's photo album. The filename returned can be loaded using QFile and related APIs. For this to be enabled, the Info.plist assigned to QMAKE_INFO_PLIST in the project file must contain the key NSPhotoLibraryUsageDescription. See Info.plist documentation from Apple for more information regarding this key. This feature was added in Qt 5.5.

See also directory().
*/
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

/*
Sets the file dialog's current directory.

Note: On iOS, if you set directory to QStandardPaths::standardLocations(QStandardPaths::PicturesLocation).last(), a native image picker dialog will be used for accessing the user's photo album. The filename returned can be loaded using QFile and related APIs. For this to be enabled, the Info.plist assigned to QMAKE_INFO_PLIST in the project file must contain the key NSPhotoLibraryUsageDescription. See Info.plist documentation from Apple for more information regarding this key. This feature was added in Qt 5.5.

See also directory().
*/
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
// [8] QDir directory() const

/*
Returns the directory currently being displayed in the dialog.

See also setDirectory().
*/
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

/*
Sets the file dialog's current directory url.

Note: The non-native QFileDialog supports only local files.

Note: On Windows, it is possible to pass URLs representing one of the virtual folders, such as "Computer" or "Network". This is done by passing a QUrl using the scheme clsid followed by the CLSID value with the curly braces removed. For example the URL clsid:374DE290-123F-4565-9164-39C4925E467B denotes the download location. For a complete list of possible values, see the MSDN documentation on KNOWNFOLDERID. This feature was added in Qt 5.5.

This function was introduced in  Qt 5.2.

See also directoryUrl() and QUuid.
*/
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
// [8] QUrl directoryUrl() const

/*
Returns the url of the directory currently being displayed in the dialog.

This function was introduced in  Qt 5.2.

See also setDirectoryUrl().
*/
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

/*
Selects the given filename in the file dialog.

See also selectedFiles().
*/
func (this *QFileDialog) SelectFile(filename string) {
	var tmpArg0 = qtcore.NewQString_5(filename)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog10selectFileERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList selectedFiles() const

/*
Returns a list of strings containing the absolute paths of the selected files in the dialog. If no files are selected, or the mode is not ExistingFiles or ExistingFile, selectedFiles() contains the current path in the viewport.

See also selectedNameFilter() and selectFile().
*/
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

/*
Selects the given url in the file dialog.

Note: The non-native QFileDialog supports only local files.

This function was introduced in  Qt 5.2.

See also selectedUrls().
*/
func (this *QFileDialog) SelectUrl(url qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog9selectUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] QList<QUrl> selectedUrls() const

/*
Returns a list of urls containing the selected files in the dialog. If no files are selected, or the mode is not ExistingFiles or ExistingFile, selectedUrls() contains the current path in the viewport.

This function was introduced in  Qt 5.2.

See also selectedNameFilter() and selectUrl().
*/
func (this *QFileDialog) SelectedUrls() *qtcore.QUrlList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog12selectedUrlsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilterDetailsVisible(bool)

/*

 */
func (this *QFileDialog) SetNameFilterDetailsVisible(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog27setNameFilterDetailsVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNameFilterDetailsVisible() const

/*

 */
func (this *QFileDialog) IsNameFilterDetailsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog26isNameFilterDetailsVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilter(const QString &)

/*
Sets the filter used in the file dialog to the given filter.

If filter contains a pair of parentheses containing one or more filename-wildcard patterns, separated by spaces, then only the text contained in the parentheses is used as the filter. This means that these calls are all equivalent:


  dialog.setNameFilter("All C++ files (*.cpp *.cc *.C *.cxx *.c++)");
  dialog.setNameFilter("*.cpp *.cc *.C *.cxx *.c++");



This function was introduced in  Qt 4.4.

See also setMimeTypeFilters() and setNameFilters().
*/
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

/*
Sets the filters used in the file dialog.

Note that the filter *.* is not portable, because the historical assumption that the file extension determines the file type is not consistent on every operating system. It is possible to have a file with no dot in its name (for example, Makefile). In a native Windows file dialog, *.* will match such files, while in other types of file dialogs it may not. So it is better to use * if you mean to select any file.


  QStringList filters;
  filters << "Image files (*.png *.xpm *.jpg)"
          << "Text files (*.txt)"
          << "Any files (*)";

  QFileDialog dialog(this);
  dialog.setNameFilters(filters);
  dialog.exec();



setMimeTypeFilters() has the advantage of providing all possible name filters for each file type. For example, JPEG images have three possible extensions; if your application can open such files, selecting the image/jpeg mime type as a filter will allow you to open all of them.

This function was introduced in  Qt 4.4.

See also nameFilters().
*/
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
// [8] QStringList nameFilters() const

/*
Returns the file type filters that are in operation on this file dialog.

This function was introduced in  Qt 4.4.

See also setNameFilters().
*/
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

/*
Sets the current file type filter. Multiple filters can be passed in filter by separating them with semicolons or spaces.

This function was introduced in  Qt 4.4.

See also setNameFilter(), setNameFilters(), and selectedNameFilter().
*/
func (this *QFileDialog) SelectNameFilter(filter string) {
	var tmpArg0 = qtcore.NewQString_5(filter)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16selectNameFilterERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QString selectedMimeTypeFilter() const

/*
Returns The mimetype of the file that the user selected in the file dialog.

This function was introduced in  Qt 5.9.
*/
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
// [8] QString selectedNameFilter() const

/*
Returns the filter that the user selected in the file dialog.

This function was introduced in  Qt 4.4.

See also selectedFiles().
*/
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

/*
Sets the filters used in the file dialog, from a list of MIME types.

Convenience method for setNameFilters(). Uses QMimeType to create a name filter from the glob patterns and description defined in each MIME type.

Use application/octet-stream for the "All files (*)" filter, since that is the base MIME type for all files.

Calling setMimeTypeFilters overrides any previously set name filters, and changes the return value of nameFilters().


  QStringList mimeTypeFilters;
  mimeTypeFilters << "image/jpeg" // will show "JPEG image (*.jpeg *.jpg *.jpe)
              << "image/png"  // will show "PNG image (*.png)"
              << "application/octet-stream"; // will show "All files (*)"

  QFileDialog dialog(this);
  dialog.setMimeTypeFilters(mimeTypeFilters);
  dialog.exec();



This function was introduced in  Qt 5.2.

See also mimeTypeFilters().
*/
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
// [8] QStringList mimeTypeFilters() const

/*
Returns the MIME type filters that are in operation on this file dialog.

This function was introduced in  Qt 5.2.

See also setMimeTypeFilters().
*/
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

/*
Sets the current MIME type filter.

This function was introduced in  Qt 5.2.
*/
func (this *QFileDialog) SelectMimeTypeFilter(filter string) {
	var tmpArg0 = qtcore.NewQString_5(filter)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog20selectMimeTypeFilterERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::Filters filter() const

/*
Returns the filter that is used when displaying files.

This function was introduced in  Qt 4.4.

See also setFilter().
*/
func (this *QFileDialog) Filter() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog6filterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilter(QDir::Filters)

/*
Sets the filter used by the model to filters. The filter is used to specify the kind of files that should be shown.

This function was introduced in  Qt 4.4.

See also filter().
*/
func (this *QFileDialog) SetFilter(filters int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog9setFilterE6QFlagsIN4QDir6FilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewMode(QFileDialog::ViewMode)

/*

 */
func (this *QFileDialog) SetViewMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog11setViewModeENS_8ViewModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] QFileDialog::ViewMode viewMode() const

/*

 */
func (this *QFileDialog) ViewMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog8viewModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileMode(QFileDialog::FileMode)

/*

 */
func (this *QFileDialog) SetFileMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog11setFileModeENS_8FileModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] QFileDialog::FileMode fileMode() const

/*

 */
func (this *QFileDialog) FileMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog8fileModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptMode(QFileDialog::AcceptMode)

/*

 */
func (this *QFileDialog) SetAcceptMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog13setAcceptModeENS_10AcceptModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:146
// index:0
// Public Visibility=Default Availability=Available
// [4] QFileDialog::AcceptMode acceptMode() const

/*

 */
func (this *QFileDialog) AcceptMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog10acceptModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(bool)

/*

 */
func (this *QFileDialog) SetReadOnly(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog11setReadOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:149
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly() const

/*

 */
func (this *QFileDialog) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolveSymlinks(bool)

/*

 */
func (this *QFileDialog) SetResolveSymlinks(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog18setResolveSymlinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:152
// index:0
// Public Visibility=Default Availability=Available
// [1] bool resolveSymlinks() const

/*

 */
func (this *QFileDialog) ResolveSymlinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog15resolveSymlinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] QList<QUrl> sidebarUrls() const

/*
Returns a list of urls that are currently in the sidebar

This function was introduced in  Qt 4.3.

See also setSidebarUrls().
*/
func (this *QFileDialog) SidebarUrls() *qtcore.QUrlList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog11sidebarUrlsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:157
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveState() const

/*
Saves the state of the dialog's layout, history and current directory.

Typically this is used in conjunction with QSettings to remember the size for a future session. A version number is stored as part of the data.

This function was introduced in  Qt 4.3.
*/
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

/*
Restores the dialogs's layout, history and current directory to the state specified.

Typically this is used in conjunction with QSettings to restore the size from a past session.

Returns false if there are errors

This function was introduced in  Qt 4.3.
*/
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
// [-2] void setConfirmOverwrite(bool)

/*

 */
func (this *QFileDialog) SetConfirmOverwrite(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog19setConfirmOverwriteEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:161
// index:0
// Public Visibility=Default Availability=Available
// [1] bool confirmOverwrite() const

/*

 */
func (this *QFileDialog) ConfirmOverwrite() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog16confirmOverwriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultSuffix(const QString &)

/*

 */
func (this *QFileDialog) SetDefaultSuffix(suffix string) {
	var tmpArg0 = qtcore.NewQString_5(suffix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16setDefaultSuffixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:164
// index:0
// Public Visibility=Default Availability=Available
// [8] QString defaultSuffix() const

/*

 */
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

/*
Sets the browsing history of the filedialog to contain the given paths.

See also history().
*/
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
// [8] QStringList history() const

/*
Returns the browsing history of the filedialog as a list of paths.

See also setHistory().
*/
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

/*
Sets the item delegate used to render items in the views in the file dialog to the given delegate.

Warning: You should not share the same instance of a delegate between views. Doing so can cause incorrect or unintuitive editing behavior since each view connected to a given delegate may receive the closeEditor() signal, and attempt to access, modify or close an editor that has already been closed.

Note that the model used is QFileSystemModel. It has custom item data roles, which is described by the Roles enum. You can use a QFileIconProvider if you only want custom icons.

See also itemDelegate(), setIconProvider(), and QFileSystemModel.
*/
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
// [8] QAbstractItemDelegate * itemDelegate() const

/*
Returns the item delegate used to render the items in the views in the filedialog.

See also setItemDelegate().
*/
func (this *QFileDialog) ItemDelegate() *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog12itemDelegateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfiledialog.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconProvider(QFileIconProvider *)

/*
Sets the icon provider used by the filedialog to the specified provider.

See also iconProvider().
*/
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
// [8] QFileIconProvider * iconProvider() const

/*
Returns the icon provider used by the filedialog.

See also setIconProvider().
*/
func (this *QFileDialog) IconProvider() *QFileIconProvider /*777 QFileIconProvider **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog12iconProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQFileIconProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfiledialog.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLabelText(QFileDialog::DialogLabel, const QString &)

/*
Sets the text shown in the filedialog in the specified label.

See also labelText().
*/
func (this *QFileDialog) SetLabelText(label int, text string) {
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog12setLabelTextENS_11DialogLabelERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), label, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:176
// index:0
// Public Visibility=Default Availability=Available
// [8] QString labelText(QFileDialog::DialogLabel) const

/*
Returns the text shown in the filedialog in the specified label.

See also setLabelText().
*/
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

/*

 */
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
// [8] QStringList supportedSchemes() const

/*

 */
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

/*
Sets the model for the views to the given proxyModel. This is useful if you want to modify the underlying model; for example, to add columns, filter data or add drives.

Any existing proxy model will be removed, but not deleted. The file dialog will take ownership of the proxyModel.

This function was introduced in  Qt 4.3.

See also proxyModel().
*/
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
// [8] QAbstractProxyModel * proxyModel() const

/*
Returns the proxy model used by the file dialog. By default no proxy is set.

See also setProxyModel().
*/
func (this *QFileDialog) ProxyModel() *qtcore.QAbstractProxyModel /*777 QAbstractProxyModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog10proxyModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQAbstractProxyModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qfiledialog.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QFileDialog::Option, bool)

/*
Sets the given option to be enabled if on is true; otherwise, clears the given option.

This function was introduced in  Qt 4.5.

See also options and testOption().
*/
func (this *QFileDialog) SetOption(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog9setOptionENS_6OptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QFileDialog::Option, bool)

/*
Sets the given option to be enabled if on is true; otherwise, clears the given option.

This function was introduced in  Qt 4.5.

See also options and testOption().
*/
func (this *QFileDialog) SetOption__(option int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog9setOptionENS_6OptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:187
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(QFileDialog::Option) const

/*
Returns true if the given option is enabled; otherwise, returns false.

This function was introduced in  Qt 4.5.

See also options and setOption().
*/
func (this *QFileDialog) TestOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog10testOptionENS_6OptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qfiledialog.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QFileDialog::Options)

/*

 */
func (this *QFileDialog) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog10setOptionsE6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:189
// index:0
// Public Visibility=Default Availability=Available
// [4] QFileDialog::Options options() const

/*

 */
func (this *QFileDialog) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QFileDialog7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void open(QObject *, const char *)

/*
This is an overloaded function.

This function connects one of its signals to the slot specified by receiver and member. The specific signal depends is filesSelected() if fileMode is ExistingFiles and fileSelected() if fileMode is anything else.

The signal will be disconnected from the slot when the dialog is closed.

This function was introduced in  Qt 4.5.
*/
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
// [-2] void setVisible(bool)

/*
Reimplemented from QWidget::setVisible().
*/
func (this *QFileDialog) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fileSelected(const QString &)

/*
When the selection changes for local operations and the dialog is accepted, this signal is emitted with the (possibly empty) selected file.

See also currentChanged() and QDialog::Accepted.
*/
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

/*
When the selection changes for local operations and the dialog is accepted, this signal is emitted with the (possibly empty) list of selected files.

See also currentChanged() and QDialog::Accepted.
*/
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

/*
When the current file changes for local operations, this signal is emitted with the new file name as the path parameter.

See also filesSelected().
*/
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

/*
This signal is emitted for local operations when the user enters a directory.

This function was introduced in  Qt 4.3.
*/
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

/*
When the selection changes and the dialog is accepted, this signal is emitted with the (possibly empty) selected url.

This function was introduced in  Qt 5.2.

See also currentUrlChanged() and QDialog::Accepted.
*/
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

/*
When the current file changes, this signal is emitted with the new file URL as the url parameter.

This function was introduced in  Qt 5.2.

See also urlsSelected().
*/
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

/*
This signal is emitted when the user enters a directory.

This function was introduced in  Qt 5.2.
*/
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

/*
This signal is emitted when the user selects a filter.

This function was introduced in  Qt 4.3.
*/
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

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns a null string.


  QString fileName = QFileDialog::getOpenFileName(this, tr("Open File"),
                                                  "/home",
                                                  tr("Images (*.png *.xpm *.jpg)"));



The function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the given filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. If you want multiple filters, separate them with ';;', for example:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileNames(), getSaveFileName(), and getExistingDirectory().
*/
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

// /usr/include/qt/QtWidgets/qfiledialog.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getOpenFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns a null string.


  QString fileName = QFileDialog::getOpenFileName(this, tr("Open File"),
                                                  "/home",
                                                  tr("Images (*.png *.xpm *.jpg)"));



The function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the given filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. If you want multiple filters, separate them with ';;', for example:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileNames(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileName__() string {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getOpenFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns a null string.


  QString fileName = QFileDialog::getOpenFileName(this, tr("Open File"),
                                                  "/home",
                                                  tr("Images (*.png *.xpm *.jpg)"));



The function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the given filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. If you want multiple filters, separate them with ';;', for example:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileNames(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileName__1(parent QWidget_ITF /*777 QWidget **/) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getOpenFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns a null string.


  QString fileName = QFileDialog::getOpenFileName(this, tr("Open File"),
                                                  "/home",
                                                  tr("Images (*.png *.xpm *.jpg)"));



The function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the given filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. If you want multiple filters, separate them with ';;', for example:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileNames(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileName__2(parent QWidget_ITF /*777 QWidget **/, caption string) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getOpenFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns a null string.


  QString fileName = QFileDialog::getOpenFileName(this, tr("Open File"),
                                                  "/home",
                                                  tr("Images (*.png *.xpm *.jpg)"));



The function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the given filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. If you want multiple filters, separate them with ';;', for example:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileNames(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileName__3(parent QWidget_ITF /*777 QWidget **/, caption string, dir string) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(dir)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getOpenFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns a null string.


  QString fileName = QFileDialog::getOpenFileName(this, tr("Open File"),
                                                  "/home",
                                                  tr("Images (*.png *.xpm *.jpg)"));



The function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the given filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. If you want multiple filters, separate them with ';;', for example:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileNames(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileName__4(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string) string {
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
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getOpenFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns a null string.


  QString fileName = QFileDialog::getOpenFileName(this, tr("Open File"),
                                                  "/home",
                                                  tr("Images (*.png *.xpm *.jpg)"));



The function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the given filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. If you want multiple filters, separate them with ';;', for example:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileNames(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileName__5(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string, selectedFilter string) string {
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
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:217
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getOpenFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getOpenFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileName(), getOpenFileUrls(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
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

// /usr/include/qt/QtWidgets/qfiledialog.h:217
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getOpenFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getOpenFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileName(), getOpenFileUrls(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrl__() *qtcore.QUrl /*123*/ {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getOpenFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:217
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getOpenFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getOpenFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileName(), getOpenFileUrls(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrl__1(parent QWidget_ITF /*777 QWidget **/) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getOpenFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:217
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getOpenFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getOpenFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileName(), getOpenFileUrls(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrl__2(parent QWidget_ITF /*777 QWidget **/, caption string) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getOpenFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:217
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getOpenFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getOpenFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileName(), getOpenFileUrls(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrl__3(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF) *qtcore.QUrl /*123*/ {
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
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getOpenFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:217
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getOpenFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getOpenFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileName(), getOpenFileUrls(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrl__4(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string) *qtcore.QUrl /*123*/ {
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
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getOpenFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:217
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getOpenFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getOpenFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileName(), getOpenFileUrls(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrl__5(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string) *qtcore.QUrl /*123*/ {
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
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getOpenFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:217
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getOpenFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns an existing file selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getOpenFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileName(), getOpenFileUrls(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrl__6(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string, options int) *qtcore.QUrl /*123*/ {
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
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getOpenFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:225
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getSaveFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return a file name selected by the user. The file does not have to exist.

It creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.


  QString fileName = QFileDialog::getSaveFileName(this, tr("Save File"),
                             "/home/jana/untitled.png",
                             tr("Images (*.png *.xpm *.jpg)"));



The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. Multiple filters are separated with ';;'. For instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The default filter can be chosen by setting selectedFilter to the desired value.

The dialog's caption is set to caption. If caption is not specified, a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar. On macOS, with its native file dialog, the filter argument is ignored.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getExistingDirectory().
*/
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

// /usr/include/qt/QtWidgets/qfiledialog.h:225
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getSaveFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return a file name selected by the user. The file does not have to exist.

It creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.


  QString fileName = QFileDialog::getSaveFileName(this, tr("Save File"),
                             "/home/jana/untitled.png",
                             tr("Images (*.png *.xpm *.jpg)"));



The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. Multiple filters are separated with ';;'. For instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The default filter can be chosen by setting selectedFilter to the desired value.

The dialog's caption is set to caption. If caption is not specified, a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar. On macOS, with its native file dialog, the filter argument is ignored.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getExistingDirectory().
*/
func (this *QFileDialog) GetSaveFileName__() string {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getSaveFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:225
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getSaveFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return a file name selected by the user. The file does not have to exist.

It creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.


  QString fileName = QFileDialog::getSaveFileName(this, tr("Save File"),
                             "/home/jana/untitled.png",
                             tr("Images (*.png *.xpm *.jpg)"));



The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. Multiple filters are separated with ';;'. For instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The default filter can be chosen by setting selectedFilter to the desired value.

The dialog's caption is set to caption. If caption is not specified, a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar. On macOS, with its native file dialog, the filter argument is ignored.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getExistingDirectory().
*/
func (this *QFileDialog) GetSaveFileName__1(parent QWidget_ITF /*777 QWidget **/) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getSaveFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:225
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getSaveFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return a file name selected by the user. The file does not have to exist.

It creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.


  QString fileName = QFileDialog::getSaveFileName(this, tr("Save File"),
                             "/home/jana/untitled.png",
                             tr("Images (*.png *.xpm *.jpg)"));



The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. Multiple filters are separated with ';;'. For instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The default filter can be chosen by setting selectedFilter to the desired value.

The dialog's caption is set to caption. If caption is not specified, a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar. On macOS, with its native file dialog, the filter argument is ignored.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getExistingDirectory().
*/
func (this *QFileDialog) GetSaveFileName__2(parent QWidget_ITF /*777 QWidget **/, caption string) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getSaveFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:225
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getSaveFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return a file name selected by the user. The file does not have to exist.

It creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.


  QString fileName = QFileDialog::getSaveFileName(this, tr("Save File"),
                             "/home/jana/untitled.png",
                             tr("Images (*.png *.xpm *.jpg)"));



The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. Multiple filters are separated with ';;'. For instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The default filter can be chosen by setting selectedFilter to the desired value.

The dialog's caption is set to caption. If caption is not specified, a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar. On macOS, with its native file dialog, the filter argument is ignored.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getExistingDirectory().
*/
func (this *QFileDialog) GetSaveFileName__3(parent QWidget_ITF /*777 QWidget **/, caption string, dir string) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(dir)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getSaveFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:225
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getSaveFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return a file name selected by the user. The file does not have to exist.

It creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.


  QString fileName = QFileDialog::getSaveFileName(this, tr("Save File"),
                             "/home/jana/untitled.png",
                             tr("Images (*.png *.xpm *.jpg)"));



The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. Multiple filters are separated with ';;'. For instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The default filter can be chosen by setting selectedFilter to the desired value.

The dialog's caption is set to caption. If caption is not specified, a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar. On macOS, with its native file dialog, the filter argument is ignored.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getExistingDirectory().
*/
func (this *QFileDialog) GetSaveFileName__4(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string) string {
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
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getSaveFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:225
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getSaveFileName(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return a file name selected by the user. The file does not have to exist.

It creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.


  QString fileName = QFileDialog::getSaveFileName(this, tr("Save File"),
                             "/home/jana/untitled.png",
                             tr("Images (*.png *.xpm *.jpg)"));



The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. Only files that match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter, and filter may be empty strings. Multiple filters are separated with ';;'. For instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

The default filter can be chosen by setting selectedFilter to the desired value.

The dialog's caption is set to caption. If caption is not specified, a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar. On macOS, with its native file dialog, the filter argument is ignored.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks the file dialog will treat symlinks as regular directories.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getExistingDirectory().
*/
func (this *QFileDialog) GetSaveFileName__5(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string, selectedFilter string) string {
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
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getSaveFileNameEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:232
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getSaveFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns a file selected by the user. The file does not have to exist. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getSaveFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getSaveFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to save the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getSaveFileName(), getOpenFileUrl(), getOpenFileUrls(), and getExistingDirectoryUrl().
*/
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

// /usr/include/qt/QtWidgets/qfiledialog.h:232
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getSaveFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns a file selected by the user. The file does not have to exist. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getSaveFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getSaveFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to save the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getSaveFileName(), getOpenFileUrl(), getOpenFileUrls(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetSaveFileUrl__() *qtcore.QUrl /*123*/ {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getSaveFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:232
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getSaveFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns a file selected by the user. The file does not have to exist. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getSaveFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getSaveFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to save the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getSaveFileName(), getOpenFileUrl(), getOpenFileUrls(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetSaveFileUrl__1(parent QWidget_ITF /*777 QWidget **/) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getSaveFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:232
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getSaveFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns a file selected by the user. The file does not have to exist. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getSaveFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getSaveFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to save the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getSaveFileName(), getOpenFileUrl(), getOpenFileUrls(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetSaveFileUrl__2(parent QWidget_ITF /*777 QWidget **/, caption string) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getSaveFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:232
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getSaveFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns a file selected by the user. The file does not have to exist. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getSaveFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getSaveFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to save the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getSaveFileName(), getOpenFileUrl(), getOpenFileUrls(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetSaveFileUrl__3(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF) *qtcore.QUrl /*123*/ {
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
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getSaveFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:232
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getSaveFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns a file selected by the user. The file does not have to exist. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getSaveFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getSaveFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to save the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getSaveFileName(), getOpenFileUrl(), getOpenFileUrls(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetSaveFileUrl__4(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string) *qtcore.QUrl /*123*/ {
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
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getSaveFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:232
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getSaveFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns a file selected by the user. The file does not have to exist. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getSaveFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getSaveFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to save the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getSaveFileName(), getOpenFileUrl(), getOpenFileUrls(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetSaveFileUrl__5(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string) *qtcore.QUrl /*123*/ {
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
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getSaveFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:232
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getSaveFileUrl(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that returns a file selected by the user. The file does not have to exist. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getSaveFileName(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getSaveFileName() comes from the ability offered to the user to select a remote file. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to save the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getSaveFileName(), getOpenFileUrl(), getOpenFileUrls(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetSaveFileUrl__6(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string, options int) *qtcore.QUrl /*123*/ {
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
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog14getSaveFileUrlEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:240
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getExistingDirectory(QWidget *, const QString &, const QString &, QFileDialog::Options)

/*
This is a convenience static function that will return an existing directory selected by the user.


  QString dir = QFileDialog::getExistingDirectory(this, tr("Open Directory"),
                                                  "/home",
                                                  QFileDialog::ShowDirsOnly
                                                  | QFileDialog::DontResolveSymlinks);



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The dialog's working directory is set to dir, and the caption is set to caption. Either of these may be an empty string in which case the current directory and a default caption will be used respectively.

The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass. To ensure a native file dialog, ShowDirsOnly must be set.

On Windows and macOS, this static function will use the native file dialog and not a QFileDialog. However, the native Windows file dialog does not support displaying files in the directory chooser. You need to pass DontUseNativeDialog to display files using a QFileDialog.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

On Windows, the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getSaveFileName().
*/
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

// /usr/include/qt/QtWidgets/qfiledialog.h:240
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getExistingDirectory(QWidget *, const QString &, const QString &, QFileDialog::Options)

/*
This is a convenience static function that will return an existing directory selected by the user.


  QString dir = QFileDialog::getExistingDirectory(this, tr("Open Directory"),
                                                  "/home",
                                                  QFileDialog::ShowDirsOnly
                                                  | QFileDialog::DontResolveSymlinks);



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The dialog's working directory is set to dir, and the caption is set to caption. Either of these may be an empty string in which case the current directory and a default caption will be used respectively.

The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass. To ensure a native file dialog, ShowDirsOnly must be set.

On Windows and macOS, this static function will use the native file dialog and not a QFileDialog. However, the native Windows file dialog does not support displaying files in the directory chooser. You need to pass DontUseNativeDialog to display files using a QFileDialog.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

On Windows, the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getSaveFileName().
*/
func (this *QFileDialog) GetExistingDirectory__() string {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog20getExistingDirectoryEP7QWidgetRK7QStringS4_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:240
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getExistingDirectory(QWidget *, const QString &, const QString &, QFileDialog::Options)

/*
This is a convenience static function that will return an existing directory selected by the user.


  QString dir = QFileDialog::getExistingDirectory(this, tr("Open Directory"),
                                                  "/home",
                                                  QFileDialog::ShowDirsOnly
                                                  | QFileDialog::DontResolveSymlinks);



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The dialog's working directory is set to dir, and the caption is set to caption. Either of these may be an empty string in which case the current directory and a default caption will be used respectively.

The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass. To ensure a native file dialog, ShowDirsOnly must be set.

On Windows and macOS, this static function will use the native file dialog and not a QFileDialog. However, the native Windows file dialog does not support displaying files in the directory chooser. You need to pass DontUseNativeDialog to display files using a QFileDialog.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

On Windows, the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getSaveFileName().
*/
func (this *QFileDialog) GetExistingDirectory__1(parent QWidget_ITF /*777 QWidget **/) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog20getExistingDirectoryEP7QWidgetRK7QStringS4_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:240
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getExistingDirectory(QWidget *, const QString &, const QString &, QFileDialog::Options)

/*
This is a convenience static function that will return an existing directory selected by the user.


  QString dir = QFileDialog::getExistingDirectory(this, tr("Open Directory"),
                                                  "/home",
                                                  QFileDialog::ShowDirsOnly
                                                  | QFileDialog::DontResolveSymlinks);



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The dialog's working directory is set to dir, and the caption is set to caption. Either of these may be an empty string in which case the current directory and a default caption will be used respectively.

The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass. To ensure a native file dialog, ShowDirsOnly must be set.

On Windows and macOS, this static function will use the native file dialog and not a QFileDialog. However, the native Windows file dialog does not support displaying files in the directory chooser. You need to pass DontUseNativeDialog to display files using a QFileDialog.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

On Windows, the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getSaveFileName().
*/
func (this *QFileDialog) GetExistingDirectory__2(parent QWidget_ITF /*777 QWidget **/, caption string) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog20getExistingDirectoryEP7QWidgetRK7QStringS4_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:240
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString getExistingDirectory(QWidget *, const QString &, const QString &, QFileDialog::Options)

/*
This is a convenience static function that will return an existing directory selected by the user.


  QString dir = QFileDialog::getExistingDirectory(this, tr("Open Directory"),
                                                  "/home",
                                                  QFileDialog::ShowDirsOnly
                                                  | QFileDialog::DontResolveSymlinks);



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The dialog's working directory is set to dir, and the caption is set to caption. Either of these may be an empty string in which case the current directory and a default caption will be used respectively.

The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass. To ensure a native file dialog, ShowDirsOnly must be set.

On Windows and macOS, this static function will use the native file dialog and not a QFileDialog. However, the native Windows file dialog does not support displaying files in the directory chooser. You need to pass DontUseNativeDialog to display files using a QFileDialog.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. If options includes DontResolveSymlinks, the file dialog will treat symlinks as regular directories.

On Windows, the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getOpenFileNames(), and getSaveFileName().
*/
func (this *QFileDialog) GetExistingDirectory__3(parent QWidget_ITF /*777 QWidget **/, caption string, dir string) string {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(dir)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog20getExistingDirectoryEP7QWidgetRK7QStringS4_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qfiledialog.h:245
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getExistingDirectoryUrl(QWidget *, const QString &, const QUrl &, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return an existing directory selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getExistingDirectory(). In particular parent, caption, dir and options are used in the exact same way.

The main difference with QFileDialog::getExistingDirectory() comes from the ability offered to the user to select a remote directory. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getExistingDirectory(), getOpenFileUrl(), getOpenFileUrls(), and getSaveFileUrl().
*/
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

// /usr/include/qt/QtWidgets/qfiledialog.h:245
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getExistingDirectoryUrl(QWidget *, const QString &, const QUrl &, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return an existing directory selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getExistingDirectory(). In particular parent, caption, dir and options are used in the exact same way.

The main difference with QFileDialog::getExistingDirectory() comes from the ability offered to the user to select a remote directory. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getExistingDirectory(), getOpenFileUrl(), getOpenFileUrls(), and getSaveFileUrl().
*/
func (this *QFileDialog) GetExistingDirectoryUrl__() *qtcore.QUrl /*123*/ {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 4, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog23getExistingDirectoryUrlEP7QWidgetRK7QStringRK4QUrl6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options, convArg4)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:245
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getExistingDirectoryUrl(QWidget *, const QString &, const QUrl &, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return an existing directory selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getExistingDirectory(). In particular parent, caption, dir and options are used in the exact same way.

The main difference with QFileDialog::getExistingDirectory() comes from the ability offered to the user to select a remote directory. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getExistingDirectory(), getOpenFileUrl(), getOpenFileUrls(), and getSaveFileUrl().
*/
func (this *QFileDialog) GetExistingDirectoryUrl__1(parent QWidget_ITF /*777 QWidget **/) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 4, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog23getExistingDirectoryUrlEP7QWidgetRK7QStringRK4QUrl6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options, convArg4)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:245
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getExistingDirectoryUrl(QWidget *, const QString &, const QUrl &, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return an existing directory selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getExistingDirectory(). In particular parent, caption, dir and options are used in the exact same way.

The main difference with QFileDialog::getExistingDirectory() comes from the ability offered to the user to select a remote directory. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getExistingDirectory(), getOpenFileUrl(), getOpenFileUrls(), and getSaveFileUrl().
*/
func (this *QFileDialog) GetExistingDirectoryUrl__2(parent QWidget_ITF /*777 QWidget **/, caption string) *qtcore.QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 4, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog23getExistingDirectoryUrlEP7QWidgetRK7QStringRK4QUrl6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options, convArg4)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:245
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getExistingDirectoryUrl(QWidget *, const QString &, const QUrl &, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return an existing directory selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getExistingDirectory(). In particular parent, caption, dir and options are used in the exact same way.

The main difference with QFileDialog::getExistingDirectory() comes from the ability offered to the user to select a remote directory. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getExistingDirectory(), getOpenFileUrl(), getOpenFileUrls(), and getSaveFileUrl().
*/
func (this *QFileDialog) GetExistingDirectoryUrl__3(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF) *qtcore.QUrl /*123*/ {
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
	// arg: 3, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 4, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog23getExistingDirectoryUrlEP7QWidgetRK7QStringRK4QUrl6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options, convArg4)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:245
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl getExistingDirectoryUrl(QWidget *, const QString &, const QUrl &, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return an existing directory selected by the user. If the user presses Cancel, it returns an empty url.

The function is used similarly to QFileDialog::getExistingDirectory(). In particular parent, caption, dir and options are used in the exact same way.

The main difference with QFileDialog::getExistingDirectory() comes from the ability offered to the user to select a remote directory. That's why the return type and the type of dir is QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getExistingDirectory(), getOpenFileUrl(), getOpenFileUrls(), and getSaveFileUrl().
*/
func (this *QFileDialog) GetExistingDirectoryUrl__4(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, options int) *qtcore.QUrl /*123*/ {
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
	// arg: 4, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog23getExistingDirectoryUrlEP7QWidgetRK7QStringRK4QUrl6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options, convArg4)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:251
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList getOpenFileNames(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return one or more existing files selected by the user.


  QStringList files = QFileDialog::getOpenFileNames(
                          this,
                          "Select one or more files to open",
                          "/home",
                          "Images (*.png *.xpm *.jpg)");



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. The filter is set to filter so that only those files which match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter and filter may be empty strings. If you need multiple filters, separate them with ';;', for instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getSaveFileName(), and getExistingDirectory().
*/
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

// /usr/include/qt/QtWidgets/qfiledialog.h:251
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList getOpenFileNames(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return one or more existing files selected by the user.


  QStringList files = QFileDialog::getOpenFileNames(
                          this,
                          "Select one or more files to open",
                          "/home",
                          "Images (*.png *.xpm *.jpg)");



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. The filter is set to filter so that only those files which match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter and filter may be empty strings. If you need multiple filters, separate them with ';;', for instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileNames__() *qtcore.QStringList /*123*/ {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16getOpenFileNamesEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:251
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList getOpenFileNames(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return one or more existing files selected by the user.


  QStringList files = QFileDialog::getOpenFileNames(
                          this,
                          "Select one or more files to open",
                          "/home",
                          "Images (*.png *.xpm *.jpg)");



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. The filter is set to filter so that only those files which match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter and filter may be empty strings. If you need multiple filters, separate them with ';;', for instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileNames__1(parent QWidget_ITF /*777 QWidget **/) *qtcore.QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16getOpenFileNamesEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:251
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList getOpenFileNames(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return one or more existing files selected by the user.


  QStringList files = QFileDialog::getOpenFileNames(
                          this,
                          "Select one or more files to open",
                          "/home",
                          "Images (*.png *.xpm *.jpg)");



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. The filter is set to filter so that only those files which match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter and filter may be empty strings. If you need multiple filters, separate them with ';;', for instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileNames__2(parent QWidget_ITF /*777 QWidget **/, caption string) *qtcore.QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16getOpenFileNamesEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:251
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList getOpenFileNames(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return one or more existing files selected by the user.


  QStringList files = QFileDialog::getOpenFileNames(
                          this,
                          "Select one or more files to open",
                          "/home",
                          "Images (*.png *.xpm *.jpg)");



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. The filter is set to filter so that only those files which match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter and filter may be empty strings. If you need multiple filters, separate them with ';;', for instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileNames__3(parent QWidget_ITF /*777 QWidget **/, caption string, dir string) *qtcore.QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(dir)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16getOpenFileNamesEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:251
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList getOpenFileNames(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return one or more existing files selected by the user.


  QStringList files = QFileDialog::getOpenFileNames(
                          this,
                          "Select one or more files to open",
                          "/home",
                          "Images (*.png *.xpm *.jpg)");



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. The filter is set to filter so that only those files which match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter and filter may be empty strings. If you need multiple filters, separate them with ';;', for instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileNames__4(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string) *qtcore.QStringList /*123*/ {
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
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16getOpenFileNamesEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:251
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList getOpenFileNames(QWidget *, const QString &, const QString &, const QString &, QString *, QFileDialog::Options)

/*
This is a convenience static function that will return one or more existing files selected by the user.


  QStringList files = QFileDialog::getOpenFileNames(
                          this,
                          "Select one or more files to open",
                          "/home",
                          "Images (*.png *.xpm *.jpg)");



This function creates a modal file dialog with the given parent widget. If parent is not 0, the dialog will be shown centered over the parent widget.

The file dialog's working directory will be set to dir. If dir includes a file name, the file will be selected. The filter is set to filter so that only those files which match the filter are shown. The filter selected is set to selectedFilter. The parameters dir, selectedFilter and filter may be empty strings. If you need multiple filters, separate them with ';;', for instance:


  "Images (*.png *.xpm *.jpg);;Text files (*.txt);;XML files (*.xml)"



The dialog's caption is set to caption. If caption is not specified then a default caption will be used.

On Windows, and macOS, this static function will use the native file dialog and not a QFileDialog.

On Windows the dialog will spin a blocking modal event loop that will not dispatch any QTimers, and if parent is not 0 then it will position the dialog just below the parent's title bar.

On Unix/X11, the normal behavior of the file dialog is to resolve and follow symlinks. For example, if /usr/tmp is a symlink to /var/tmp, the file dialog will change to /var/tmp after entering /usr/tmp. The options argument holds various options about how to run the dialog, see the QFileDialog::Option enum for more information on the flags you can pass.

Warning: Do not delete parent during the execution of the dialog. If you want to do this, you should create the dialog yourself using one of the QFileDialog constructors.

See also getOpenFileName(), getSaveFileName(), and getExistingDirectory().
*/
func (this *QFileDialog) GetOpenFileNames__5(parent QWidget_ITF /*777 QWidget **/, caption string, dir string, filter string, selectedFilter string) *qtcore.QStringList /*123*/ {
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
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog16getOpenFileNamesEP7QWidgetRK7QStringS4_S4_PS2_6QFlagsINS_6OptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:258
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QUrl> getOpenFileUrls(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return one or more existing files selected by the user. If the user presses Cancel, it returns an empty list.

The function is used similarly to QFileDialog::getOpenFileNames(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileNames() comes from the ability offered to the user to select remote files. That's why the return type and the type of dir are respectively QList<QUrl> and QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileNames(), getOpenFileUrl(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrls(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string, options int, supportedSchemes qtcore.QStringList_ITF) *qtcore.QUrlList /*lll*/ {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileUrlsEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}
func QFileDialog_GetOpenFileUrls(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string, options int, supportedSchemes qtcore.QStringList_ITF) *qtcore.QUrlList /*lll*/ {
	var nilthis *QFileDialog
	rv := nilthis.GetOpenFileUrls(parent, caption, dir, filter, selectedFilter, options, supportedSchemes)
	return rv
}

// /usr/include/qt/QtWidgets/qfiledialog.h:258
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QUrl> getOpenFileUrls(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return one or more existing files selected by the user. If the user presses Cancel, it returns an empty list.

The function is used similarly to QFileDialog::getOpenFileNames(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileNames() comes from the ability offered to the user to select remote files. That's why the return type and the type of dir are respectively QList<QUrl> and QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileNames(), getOpenFileUrl(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrls__() *qtcore.QUrlList /*lll*/ {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileUrlsEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:258
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QUrl> getOpenFileUrls(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return one or more existing files selected by the user. If the user presses Cancel, it returns an empty list.

The function is used similarly to QFileDialog::getOpenFileNames(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileNames() comes from the ability offered to the user to select remote files. That's why the return type and the type of dir are respectively QList<QUrl> and QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileNames(), getOpenFileUrl(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrls__1(parent QWidget_ITF /*777 QWidget **/) *qtcore.QUrlList /*lll*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileUrlsEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:258
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QUrl> getOpenFileUrls(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return one or more existing files selected by the user. If the user presses Cancel, it returns an empty list.

The function is used similarly to QFileDialog::getOpenFileNames(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileNames() comes from the ability offered to the user to select remote files. That's why the return type and the type of dir are respectively QList<QUrl> and QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileNames(), getOpenFileUrl(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrls__2(parent QWidget_ITF /*777 QWidget **/, caption string) *qtcore.QUrlList /*lll*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(caption)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QUrl &=LValueReference, QUrl=Record, , Invalid
	var convArg2 = qtcore.NewQUrl()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileUrlsEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:258
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QUrl> getOpenFileUrls(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return one or more existing files selected by the user. If the user presses Cancel, it returns an empty list.

The function is used similarly to QFileDialog::getOpenFileNames(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileNames() comes from the ability offered to the user to select remote files. That's why the return type and the type of dir are respectively QList<QUrl> and QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileNames(), getOpenFileUrl(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrls__3(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF) *qtcore.QUrlList /*lll*/ {
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
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = qtcore.NewQString()
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileUrlsEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:258
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QUrl> getOpenFileUrls(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return one or more existing files selected by the user. If the user presses Cancel, it returns an empty list.

The function is used similarly to QFileDialog::getOpenFileNames(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileNames() comes from the ability offered to the user to select remote files. That's why the return type and the type of dir are respectively QList<QUrl> and QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileNames(), getOpenFileUrl(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrls__4(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string) *qtcore.QUrlList /*lll*/ {
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
	// arg: 4, QString *=Pointer, QString=Record, , Invalid
	var convArg4 = qtcore.NewQString()
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileUrlsEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:258
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QUrl> getOpenFileUrls(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return one or more existing files selected by the user. If the user presses Cancel, it returns an empty list.

The function is used similarly to QFileDialog::getOpenFileNames(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileNames() comes from the ability offered to the user to select remote files. That's why the return type and the type of dir are respectively QList<QUrl> and QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileNames(), getOpenFileUrl(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrls__5(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string) *qtcore.QUrlList /*lll*/ {
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
	// arg: 5, QFileDialog::Options=Typedef, QFileDialog::Options=Typedef, QFlags<QFileDialog::Option>, Unexposed
	options := 0
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileUrlsEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:258
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QUrl> getOpenFileUrls(QWidget *, const QString &, const QUrl &, const QString &, QString *, QFileDialog::Options, const QStringList &)

/*
This is a convenience static function that will return one or more existing files selected by the user. If the user presses Cancel, it returns an empty list.

The function is used similarly to QFileDialog::getOpenFileNames(). In particular parent, caption, dir, filter, selectedFilter and options are used in the exact same way.

The main difference with QFileDialog::getOpenFileNames() comes from the ability offered to the user to select remote files. That's why the return type and the type of dir are respectively QList<QUrl> and QUrl.

The supportedSchemes argument allows to restrict the type of URLs the user will be able to select. It is a way for the application to declare the protocols it will support to fetch the file content. An empty list means that no restriction is applied (the default). Supported for local files ("file" scheme) is implicit and always enabled; it is not necessary to include it in the restriction.

When possible, this static function will use the native file dialog and not a QFileDialog. On platforms which don't support selecting remote files, Qt will allow to select only local files.

This function was introduced in  Qt 5.2.

See also getOpenFileNames(), getOpenFileUrl(), getSaveFileUrl(), and getExistingDirectoryUrl().
*/
func (this *QFileDialog) GetOpenFileUrls__6(parent QWidget_ITF /*777 QWidget **/, caption string, dir qtcore.QUrl_ITF, filter string, selectedFilter string, options int) *qtcore.QUrlList /*lll*/ {
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
	// arg: 6, const QStringList &=LValueReference, QStringList=Record, , Invalid
	var convArg6 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog15getOpenFileUrlsEP7QWidgetRK7QStringRK4QUrlS4_PS2_6QFlagsINS_6OptionEERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, options, convArg6)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qfiledialog.h:269
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void done(int)

/*
Reimplemented from QDialog::done().
*/
func (this *QFileDialog) Done(result int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog4doneEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), result)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:270
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void accept()

/*
Reimplemented from QDialog::accept().
*/
func (this *QFileDialog) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qfiledialog.h:271
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QWidget::changeEvent().
*/
func (this *QFileDialog) ChangeEvent(e qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the view mode of the file dialog; i.e. what information about each file will be displayed.



See also setViewMode().

*/
type QFileDialog__ViewMode = int

// Displays an icon, a name, and details for each item in the directory.
const QFileDialog__Detail QFileDialog__ViewMode = 0

// Displays only an icon and a name for each item in the directory.
const QFileDialog__List QFileDialog__ViewMode = 1

func (this *QFileDialog) ViewModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFileDialog_ViewModeItemName(val int) string {
	var nilthis *QFileDialog
	return nilthis.ViewModeItemName(val)
}

/*
This enum is used to indicate what the user may select in the file dialog; i.e. what the dialog will return if the user clicks OK.



This value is obsolete since Qt 4.5:



See also setFileMode().

*/
type QFileDialog__FileMode = int

// The name of a file, whether it exists or not.
const QFileDialog__AnyFile QFileDialog__FileMode = 0

// The name of a single existing file.
const QFileDialog__ExistingFile QFileDialog__FileMode = 1

// The name of a directory. Both files and directories are displayed. However, the native Windows file dialog does not support displaying files in the directory chooser.
const QFileDialog__Directory QFileDialog__FileMode = 2

// The names of zero or more existing files.
const QFileDialog__ExistingFiles QFileDialog__FileMode = 3

// Use Directory and setOption(ShowDirsOnly, true) instead.
const QFileDialog__DirectoryOnly QFileDialog__FileMode = 4

func (this *QFileDialog) FileModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFileDialog_FileModeItemName(val int) string {
	var nilthis *QFileDialog
	return nilthis.FileModeItemName(val)
}

/*
ConstantValue
QFileDialog::AcceptOpen0
QFileDialog::AcceptSave1

*/
type QFileDialog__AcceptMode = int

//
const QFileDialog__AcceptOpen QFileDialog__AcceptMode = 0

//
const QFileDialog__AcceptSave QFileDialog__AcceptMode = 1

func (this *QFileDialog) AcceptModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFileDialog_AcceptModeItemName(val int) string {
	var nilthis *QFileDialog
	return nilthis.AcceptModeItemName(val)
}

/*
ConstantValue
QFileDialog::LookIn0
QFileDialog::FileName1
QFileDialog::FileType2
QFileDialog::Accept3
QFileDialog::Reject4

*/
type QFileDialog__DialogLabel = int

//
const QFileDialog__LookIn QFileDialog__DialogLabel = 0

//
const QFileDialog__FileName QFileDialog__DialogLabel = 1

//
const QFileDialog__FileType QFileDialog__DialogLabel = 2

//
const QFileDialog__Accept QFileDialog__DialogLabel = 3

//
const QFileDialog__Reject QFileDialog__DialogLabel = 4

func (this *QFileDialog) DialogLabelItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFileDialog_DialogLabelItemName(val int) string {
	var nilthis *QFileDialog
	return nilthis.DialogLabelItemName(val)
}

/*


 */
type QFileDialog__Option = int

//
const QFileDialog__ShowDirsOnly QFileDialog__Option = 1

//
const QFileDialog__DontResolveSymlinks QFileDialog__Option = 2

//
const QFileDialog__DontConfirmOverwrite QFileDialog__Option = 4

//
const QFileDialog__DontUseSheet QFileDialog__Option = 8

//
const QFileDialog__DontUseNativeDialog QFileDialog__Option = 16

//
const QFileDialog__ReadOnly QFileDialog__Option = 32

//
const QFileDialog__HideNameFilterDetails QFileDialog__Option = 64

//
const QFileDialog__DontUseCustomDirectoryIcons QFileDialog__Option = 128

func (this *QFileDialog) OptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QFileDialog_OptionItemName(val int) string {
	var nilthis *QFileDialog
	return nilthis.OptionItemName(val)
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
