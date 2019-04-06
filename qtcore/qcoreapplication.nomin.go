//  header block begin

// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qcoreapplication.h
// #include <qcoreapplication.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// /usr/include/qt/QtCore/qcoreapplication.h:145
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setLibraryPaths(const QStringList &)

/*
Sets the list of directories to search when loading libraries to paths. All existing paths will be deleted and the path list will consist of the paths given in paths.

The library paths are reset to the default when an instance of QCoreApplication is destructed.

See also libraryPaths(), addLibraryPath(), removeLibraryPath(), and QLibrary.
*/
func (this *QCoreApplication) SetLibraryPaths(arg0 QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStringList_PTR() != nil {
		convArg0 = arg0.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication15setLibraryPathsERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SetLibraryPaths(arg0 QStringList_ITF) {
	var nilthis *QCoreApplication
	nilthis.SetLibraryPaths(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:146
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList libraryPaths()

/*
Returns a list of paths that the application will search when dynamically loading libraries.

The return value of this function may change when a QCoreApplication is created. It is not recommended to call it before creating a QCoreApplication. The directory of the application executable (not the working directory) is part of the list if it is known. In order to make it known a QCoreApplication has to be constructed as it will use argv[0] to find it.

Qt provides default library paths, but they can also be set using a qt.conf file. Paths specified in this file will override default values. Note that if the qt.conf file is in the directory of the application executable, it may not be found until a QCoreApplication is created. If it is not found when calling this function, the default library paths will be used.

The list will include the installation directory for plugins if it exists (the default installation directory for plugins is INSTALL/plugins, where INSTALL is the directory where Qt was installed). The colon separated entries of the QT_PLUGIN_PATH environment variable are always added. The plugin installation directory (and its existence) may change when the directory of the application executable becomes known.

If you want to iterate over the list, you can use the foreach pseudo-keyword:


  foreach (const QString &path, app.libraryPaths())
      do_something(path);



See also setLibraryPaths(), addLibraryPath(), removeLibraryPath(), QLibrary, and How to Create Qt Plugins.
*/
func (this *QCoreApplication) LibraryPaths() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication12libraryPathsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QCoreApplication_LibraryPaths() *QStringList /*123*/ {
	var nilthis *QCoreApplication
	rv := nilthis.LibraryPaths()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:147
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addLibraryPath(const QString &)

/*
Prepends path to the beginning of the library path list, ensuring that it is searched for libraries first. If path is empty or already in the path list, the path list is not changed.

The default path list consists of a single entry, the installation directory for plugins. The default installation directory for plugins is INSTALL/plugins, where INSTALL is the directory where Qt was installed.

The library paths are reset to the default when an instance of QCoreApplication is destructed.

See also removeLibraryPath(), libraryPaths(), and setLibraryPaths().
*/
func (this *QCoreApplication) AddLibraryPath(arg0 string) {
	var tmpArg0 = NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication14addLibraryPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_AddLibraryPath(arg0 string) {
	var nilthis *QCoreApplication
	nilthis.AddLibraryPath(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:148
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void removeLibraryPath(const QString &)

/*
Removes path from the library path list. If path is empty or not in the path list, the list is not changed.

The library paths are reset to the default when an instance of QCoreApplication is destructed.

See also addLibraryPath(), libraryPaths(), and setLibraryPaths().
*/
func (this *QCoreApplication) RemoveLibraryPath(arg0 string) {
	var tmpArg0 = NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication17removeLibraryPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_RemoveLibraryPath(arg0 string) {
	var nilthis *QCoreApplication
	nilthis.RemoveLibraryPath(arg0)
}

//  body block end

//  keep block begin

func init_unused_10380() {
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
}

//  keep block end
