package qtcore

// /usr/include/qt/QtCore/qlockfile.h
// #include <qlockfile.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QLockFile struct {
	*qtrt.CObject
}
type QLockFile_ITF interface {
	QLockFile_PTR() *QLockFile
}

func (ptr *QLockFile) QLockFile_PTR() *QLockFile { return ptr }

func (this *QLockFile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLockFile) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQLockFileFromPointer(cthis unsafe.Pointer) *QLockFile {
	return &QLockFile{&qtrt.CObject{cthis}}
}
func (*QLockFile) NewFromPointer(cthis unsafe.Pointer) *QLockFile {
	return NewQLockFileFromPointer(cthis)
}

// /usr/include/qt/QtCore/qlockfile.h:53
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLockFile(const QString &)
func NewQLockFile(fileName string) *QLockFile {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFileC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLockFileFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQLockFile)
	return gothis
}

// /usr/include/qt/QtCore/qlockfile.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QLockFile()
func DeleteQLockFile(this *QLockFile) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFileD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qlockfile.h:56
// index:0
// Public Visibility=Default Availability=Available
// [1] bool lock()
func (this *QLockFile) Lock() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile4lockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:57
// index:0
// Public Visibility=Default Availability=Available
// [1] bool tryLock(int)
func (this *QLockFile) TryLock(timeout int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile7tryLockEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unlock()
func (this *QLockFile) Unlock() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile6unlockEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStaleLockTime(int)
func (this *QLockFile) SetStaleLockTime(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile16setStaleLockTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:61
// index:0
// Public Visibility=Default Availability=Available
// [4] int staleLockTime()
func (this *QLockFile) StaleLockTime() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLockFile13staleLockTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlockfile.h:63
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLocked()
func (this *QLockFile) IsLocked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLockFile8isLockedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool getLockInfo(qint64 *, QString *, QString *)
func (this *QLockFile) GetLockInfo(pid unsafe.Pointer /*666*/, hostname string, appname string) bool {
	var tmpArg1 = NewQString_5(hostname)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(appname)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLockFile11getLockInfoEPxP7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pid, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool removeStaleLockFile()
func (this *QLockFile) RemoveStaleLockFile() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QLockFile19removeStaleLockFileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] QLockFile::LockError error()
func (this *QLockFile) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QLockFile5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

type QLockFile__LockError = int

const QLockFile__NoError QLockFile__LockError = 0
const QLockFile__LockFailedError QLockFile__LockError = 1
const QLockFile__PermissionError QLockFile__LockError = 2
const QLockFile__UnknownError QLockFile__LockError = 3

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
