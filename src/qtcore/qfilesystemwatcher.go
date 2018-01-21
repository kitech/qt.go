package qtcore

// /usr/include/qt/QtCore/qfilesystemwatcher.h
// #include <qfilesystemwatcher.h>
// #include <QtCore>

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
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func NewQFileSystemWatcherFromPointer(cthis unsafe.Pointer) *QFileSystemWatcher {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFileSystemWatcher{bcthis0}
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFileSystemWatcher) MetaObject() *QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFileSystemWatcher10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:58
// index:0
// Public
// void QFileSystemWatcher(class QObject *)
func NewQFileSystemWatcher(parent *QObject /*444 QObject **/) *QFileSystemWatcher {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:59
// index:1
// Public
// void QFileSystemWatcher(const class QStringList &, class QObject *)
func NewQFileSystemWatcher_1(paths *QStringList, parent *QObject /*444 QObject **/) *QFileSystemWatcher {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = paths.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2ERK11QStringListP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
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
func (this *QFileSystemWatcher) AddPath(file *QString) bool {
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcher7addPathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:64
// index:0
// Public
// bool removePath(const class QString &)
func (this *QFileSystemWatcher) RemovePath(file *QString) bool {
	var convArg0 = file.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFileSystemWatcher10removePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
