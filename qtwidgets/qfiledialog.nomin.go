//  header block begin

// +build !minimal

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
	var tmpArg0 = qtcore.NewQString5(filter)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QFileDialog20selectMimeTypeFilterERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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

//  body block end

//  keep block begin

func init_unused_11138() {
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
