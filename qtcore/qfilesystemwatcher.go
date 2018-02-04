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
import "gopp"
import "qt.go/qtrt"

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
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileSystemWatcher(QObject *)
func NewQFileSystemWatcher(parent *QObject /*777 QObject **/) *QFileSystemWatcher {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
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
	gopp.ErrPrint(err, rv)
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
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:62
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addPath(const QString &)
func (this *QFileSystemWatcher) AddPath(file *QString) bool {
	var convArg0 = file.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcher7addPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool removePath(const QString &)
func (this *QFileSystemWatcher) RemovePath(file *QString) bool {
	var convArg0 = file.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcher10removePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
