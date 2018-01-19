//  header block begin
// /usr/include/qt/QtCore/qlockfile.h
// #include <qlockfile.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
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
type QLockFile struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qlockfile.h:53
// index:0
// void QLockFile(const class QString &)
func NewQLockFile(fileName unsafe.Pointer) *QLockFile {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFileC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, fileName)
	gopp.ErrPrint(err, rv)
	return &QLockFile{cthis}
}

// /usr/include/qt/QtCore/qlockfile.h:54
// index:0
// void ~QLockFile()
func DeleteQLockFile(*QLockFile) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFileD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:56
// index:0
// bool lock()
func (this *QLockFile) Lock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFile4lockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:57
// index:0
// bool tryLock(int)
func (this *QLockFile) TryLock(timeout int) {
	// 0: (, int timeout), (&timeout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFile7tryLockEi", ffiqt.FFI_TYPE_VOID, this.cthis, &timeout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:58
// index:0
// void unlock()
func (this *QLockFile) Unlock() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFile6unlockEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:60
// index:0
// void setStaleLockTime(int)
func (this *QLockFile) SetStaleLockTime(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFile16setStaleLockTimeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:61
// index:0
// int staleLockTime()
func (this *QLockFile) StaleLockTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLockFile13staleLockTimeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:63
// index:0
// bool isLocked()
func (this *QLockFile) IsLocked() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLockFile8isLockedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:64
// index:0
// bool getLockInfo(qint64 *, class QString *, class QString *)
func (this *QLockFile) GetLockInfo(pid unsafe.Pointer, hostname unsafe.Pointer, appname unsafe.Pointer) {
	// 0: (, qint64 * pid, QString * hostname, QString * appname), (pid, hostname, appname)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLockFile11getLockInfoEPxP7QStringS2_", ffiqt.FFI_TYPE_VOID, this.cthis, pid, hostname, appname)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:65
// index:0
// bool removeStaleLockFile()
func (this *QLockFile) RemoveStaleLockFile() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QLockFile19removeStaleLockFileEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlockfile.h:73
// index:0
// QLockFile::LockError error()
func (this *QLockFile) Error() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QLockFile5errorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
