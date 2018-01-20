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

// /usr/include/qt/QtCore/qfilesystemwatcher.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFileSystemWatcher) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFileSystemWatcher10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:58
// index:0
// void QFileSystemWatcher(class QObject *)
func NewQFileSystemWatcher(parent unsafe.Pointer) *QFileSystemWatcher {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(cthis)
	return gothis
}
func NewQFileSystemWatcherFromPointer(cthis unsafe.Pointer) *QFileSystemWatcher {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFileSystemWatcher{bcthis0}
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:59
// index:1
// void QFileSystemWatcher(const class QStringList &, class QObject *)
func NewQFileSystemWatcher_1(paths unsafe.Pointer, parent unsafe.Pointer) *QFileSystemWatcher {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2ERK11QStringListP7QObject", ffiqt.FFI_TYPE_VOID, cthis, paths, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:60
// index:0
// virtual
// void ~QFileSystemWatcher()
func DeleteQFileSystemWatcher(*QFileSystemWatcher) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcherD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:62
// index:0
// bool addPath(const class QString &)
func (this *QFileSystemWatcher) AddPath(file unsafe.Pointer) {
	// 0: (, file const QString &), (file)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcher7addPathERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), file)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:63
// index:0
// QStringList addPaths(const class QStringList &)
func (this *QFileSystemWatcher) AddPaths(files unsafe.Pointer) {
	// 0: (, files const QStringList &), (files)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcher8addPathsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), files)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:64
// index:0
// bool removePath(const class QString &)
func (this *QFileSystemWatcher) RemovePath(file unsafe.Pointer) {
	// 0: (, file const QString &), (file)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcher10removePathERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), file)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:65
// index:0
// QStringList removePaths(const class QStringList &)
func (this *QFileSystemWatcher) RemovePaths(files unsafe.Pointer) {
	// 0: (, files const QStringList &), (files)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcher11removePathsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), files)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:67
// index:0
// QStringList files()
func (this *QFileSystemWatcher) Files() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFileSystemWatcher5filesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:68
// index:0
// QStringList directories()
func (this *QFileSystemWatcher) Directories() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFileSystemWatcher11directoriesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
