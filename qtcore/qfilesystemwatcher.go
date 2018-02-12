package qtcore

// /usr/include/qt/QtCore/qfilesystemwatcher.h
// #include <qfilesystemwatcher.h>
// #include <QtCore>

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
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QFileSystemWatcher struct {
	*QObject
}

func (this *QFileSystemWatcher) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QFileSystemWatcher) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQFileSystemWatcherFromPointer(cthis unsafe.Pointer) *QFileSystemWatcher {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFileSystemWatcher{bcthis0}
}
func (*QFileSystemWatcher) NewFromPointer(cthis unsafe.Pointer) *QFileSystemWatcher {
	return NewQFileSystemWatcherFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QFileSystemWatcher) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFileSystemWatcher10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileSystemWatcher(QObject *)
func NewQFileSystemWatcher(parent *QObject /*777 QObject **/) *QFileSystemWatcher {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileSystemWatcher(const QStringList &, QObject *)
func NewQFileSystemWatcher_1(paths *QStringList, parent *QObject /*777 QObject **/) *QFileSystemWatcher {
	var convArg0 = paths.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2ERK11QStringListP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFileSystemWatcher()
func DeleteQFileSystemWatcher(this *QFileSystemWatcher) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcherD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:62
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addPath(const QString &)
func (this *QFileSystemWatcher) AddPath(file string) bool {
	var tmpArg0 = NewQString_5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcher7addPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList addPaths(const QStringList &)
func (this *QFileSystemWatcher) AddPaths(files *QStringList) *QStringList /*123*/ {
	var convArg0 = files.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcher8addPathsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool removePath(const QString &)
func (this *QFileSystemWatcher) RemovePath(file string) bool {
	var tmpArg0 = NewQString_5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcher10removePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList removePaths(const QStringList &)
func (this *QFileSystemWatcher) RemovePaths(files *QStringList) *QStringList /*123*/ {
	var convArg0 = files.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcher11removePathsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList files()
func (this *QFileSystemWatcher) Files() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFileSystemWatcher5filesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList directories()
func (this *QFileSystemWatcher) Directories() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFileSystemWatcher11directoriesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
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
		qtrt.KeepMe()
	}
}

//  keep block end
