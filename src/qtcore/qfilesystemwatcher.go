//  header block begin
// /usr/include/qt/QtCore/qfilesystemwatcher.h
// #include <qfilesystemwatcher.h>
// #include <QtCore>
package qtcore

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
}

//  ext block end

//  body block begin
type QFileSystemWatcher struct {
	*QObject
}

func (this *QFileSystemWatcher) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQFileSystemWatcherFromPointer(cthis unsafe.Pointer) *QFileSystemWatcher {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFileSystemWatcher{bcthis0}
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFileSystemWatcher) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFileSystemWatcher10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:58
// index:0
// Public
// void QFileSystemWatcher(class QObject *)
func NewQFileSystemWatcher(parent unsafe.Pointer) *QFileSystemWatcher {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:59
// index:1
// Public
// void QFileSystemWatcher(const class QStringList &, class QObject *)
func NewQFileSystemWatcher_1(paths *QStringList, parent unsafe.Pointer) *QFileSystemWatcher {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = paths.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2ERK11QStringListP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:60
// index:0
// Public virtual
// void ~QFileSystemWatcher()
func DeleteQFileSystemWatcher(*QFileSystemWatcher) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcherD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:62
// index:0
// Public
// bool addPath(const class QString &)
func (this *QFileSystemWatcher) AddPath(file *QString) interface{} {
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcher7addPathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:63
// index:0
// Public
// QStringList addPaths(const class QStringList &)
func (this *QFileSystemWatcher) AddPaths(files *QStringList) interface{} {
	var convArg0 = files.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcher8addPathsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:64
// index:0
// Public
// bool removePath(const class QString &)
func (this *QFileSystemWatcher) RemovePath(file *QString) interface{} {
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcher10removePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:65
// index:0
// Public
// QStringList removePaths(const class QStringList &)
func (this *QFileSystemWatcher) RemovePaths(files *QStringList) interface{} {
	var convArg0 = files.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcher11removePathsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:67
// index:0
// Public
// QStringList files()
func (this *QFileSystemWatcher) Files() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFileSystemWatcher5filesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:68
// index:0
// Public
// QStringList directories()
func (this *QFileSystemWatcher) Directories() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFileSystemWatcher11directoriesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
