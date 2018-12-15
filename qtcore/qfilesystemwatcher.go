// +build !minimal

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
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QFileSystemWatcher struct {
	*QObject
}
type QFileSystemWatcher_ITF interface {
	QObject_ITF
	QFileSystemWatcher_PTR() *QFileSystemWatcher
}

func (ptr *QFileSystemWatcher) QFileSystemWatcher_PTR() *QFileSystemWatcher { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QFileSystemWatcher) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QFileSystemWatcher10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileSystemWatcher(QObject *)

/*
Constructs a new file system watcher object with the given parent.
*/
func (*QFileSystemWatcher) NewForInherit(parent QObject_ITF /*777 QObject **/) *QFileSystemWatcher {
	return NewQFileSystemWatcher(parent)
}
func NewQFileSystemWatcher(parent QObject_ITF /*777 QObject **/) *QFileSystemWatcher {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileSystemWatcher")
	return gothis
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFileSystemWatcher(QObject *)

/*
Constructs a new file system watcher object with the given parent.
*/
func (*QFileSystemWatcher) NewForInheritp() *QFileSystemWatcher {
	return NewQFileSystemWatcherp()
}
func NewQFileSystemWatcherp() *QFileSystemWatcher {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileSystemWatcher")
	return gothis
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileSystemWatcher(const QStringList &, QObject *)

/*
Constructs a new file system watcher object with the given parent.
*/
func (*QFileSystemWatcher) NewForInherit1(paths QStringList_ITF, parent QObject_ITF /*777 QObject **/) *QFileSystemWatcher {
	return NewQFileSystemWatcher1(paths, parent)
}
func NewQFileSystemWatcher1(paths QStringList_ITF, parent QObject_ITF /*777 QObject **/) *QFileSystemWatcher {
	var convArg0 unsafe.Pointer
	if paths != nil && paths.QStringList_PTR() != nil {
		convArg0 = paths.QStringList_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2ERK11QStringListP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileSystemWatcher")
	return gothis
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFileSystemWatcher(const QStringList &, QObject *)

/*
Constructs a new file system watcher object with the given parent.
*/
func (*QFileSystemWatcher) NewForInherit1p(paths QStringList_ITF) *QFileSystemWatcher {
	return NewQFileSystemWatcher1p(paths)
}
func NewQFileSystemWatcher1p(paths QStringList_ITF) *QFileSystemWatcher {
	var convArg0 unsafe.Pointer
	if paths != nil && paths.QStringList_PTR() != nil {
		convArg0 = paths.QStringList_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcherC2ERK11QStringListP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFileSystemWatcherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QFileSystemWatcher")
	return gothis
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFileSystemWatcher()

/*

 */
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

/*
Adds path to the file system watcher if path exists. The path is not added if it does not exist, or if it is already being monitored by the file system watcher.

If path specifies a directory, the directoryChanged() signal will be emitted when path is modified or removed from disk; otherwise the fileChanged() signal is emitted when path is modified, renamed or removed.

If the watch was successful, true is returned.

Reasons for a watch failure are generally system-dependent, but may include the resource not existing, access failures, or the total watch count limit, if the platform has one.

Note: There may be a system dependent limit to the number of files and directories that can be monitored simultaneously. If this limit is been reached, path will not be monitored, and false is returned.

See also addPaths() and removePath().
*/
func (this *QFileSystemWatcher) AddPath(file string) bool {
	var tmpArg0 = NewQString5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcher7addPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList addPaths(const QStringList &)

/*
Adds each path in paths to the file system watcher. Paths are not added if they not exist, or if they are already being monitored by the file system watcher.

If a path specifies a directory, the directoryChanged() signal will be emitted when the path is modified or removed from disk; otherwise the fileChanged() signal is emitted when the path is modified, renamed, or removed.

The return value is a list of paths that could not be watched.

Reasons for a watch failure are generally system-dependent, but may include the resource not existing, access failures, or the total watch count limit, if the platform has one.

Note: There may be a system dependent limit to the number of files and directories that can be monitored simultaneously. If this limit has been reached, the excess paths will not be monitored, and they will be added to the returned QStringList.

See also addPath() and removePaths().
*/
func (this *QFileSystemWatcher) AddPaths(files QStringList_ITF) *QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if files != nil && files.QStringList_PTR() != nil {
		convArg0 = files.QStringList_PTR().GetCthis()
	}
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

/*
Removes the specified path from the file system watcher.

If the watch is successfully removed, true is returned.

Reasons for watch removal failing are generally system-dependent, but may be due to the path having already been deleted, for example.

See also removePaths() and addPath().
*/
func (this *QFileSystemWatcher) RemovePath(file string) bool {
	var tmpArg0 = NewQString5(file)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcher10removePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList removePaths(const QStringList &)

/*
Removes the specified paths from the file system watcher.

The return value is a list of paths which were not able to be unwatched successfully.

Reasons for watch removal failing are generally system-dependent, but may be due to the path having already been deleted, for example.

See also removePath() and addPaths().
*/
func (this *QFileSystemWatcher) RemovePaths(files QStringList_ITF) *QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if files != nil && files.QStringList_PTR() != nil {
		convArg0 = files.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QFileSystemWatcher11removePathsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qfilesystemwatcher.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList files() const

/*
Returns a list of paths to files that are being watched.

See also directories().
*/
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
// [8] QStringList directories() const

/*
Returns a list of paths to directories that are being watched.

See also files().
*/
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
