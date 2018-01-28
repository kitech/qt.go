package qtcore

// /usr/include/qt/QtCore/qlockfile.h
// #include <qlockfile.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QLockFile struct {
	*qtrt.CObject
}

func (this *QLockFile) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLockFile) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQLockFileFromPointer(cthis unsafe.Pointer) *QLockFile {
	return &QLockFile{&qtrt.CObject{cthis}}
}
func (*QLockFile) NewFromPointer(cthis unsafe.Pointer) *QLockFile {
	return NewQLockFileFromPointer(cthis)
}

// /usr/include/qt/QtCore/qlockfile.h:53
// index:0
// Public
// void QLockFile(const QString &)
func NewQLockFile(fileName *QString) *QLockFile {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFileC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLockFileFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qlockfile.h:54
// index:0
// Public
// void ~QLockFile()
func DeleteQLockFile(*QLockFile) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFileD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:56
// index:0
// Public
// bool lock()
func (this *QLockFile) Lock() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFile4lockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:57
// index:0
// Public
// bool tryLock(int)
func (this *QLockFile) TryLock(timeout int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFile7tryLockEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:58
// index:0
// Public
// void unlock()
func (this *QLockFile) Unlock() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFile6unlockEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:60
// index:0
// Public
// void setStaleLockTime(int)
func (this *QLockFile) SetStaleLockTime(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFile16setStaleLockTimeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:61
// index:0
// Public
// int staleLockTime()
func (this *QLockFile) StaleLockTime() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLockFile13staleLockTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qlockfile.h:63
// index:0
// Public
// bool isLocked()
func (this *QLockFile) IsLocked() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLockFile8isLockedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:64
// index:0
// Public
// bool getLockInfo(qint64 *, QString *, QString *)
func (this *QLockFile) GetLockInfo(pid unsafe.Pointer /*666*/, hostname *QString /*777 QString **/, appname *QString /*777 QString **/) bool {
	var convArg1 = hostname.GetCthis()
	var convArg2 = appname.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLockFile11getLockInfoEPxP7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pid, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:65
// index:0
// Public
// bool removeStaleLockFile()
func (this *QLockFile) RemoveStaleLockFile() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFile19removeStaleLockFileEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qlockfile.h:73
// index:0
// Public
// QLockFile::LockError error()
func (this *QLockFile) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLockFile5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

type QLockFile__LockError = int

const QLockFile__NoError QLockFile__LockError = 0
const QLockFile__LockFailedError QLockFile__LockError = 1
const QLockFile__PermissionError QLockFile__LockError = 2
const QLockFile__UnknownError QLockFile__LockError = 3

//  body block end
