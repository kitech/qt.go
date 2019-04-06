package qtcore

// /usr/include/qt/QtCore/qreadwritelock.h
// #include <qreadwritelock.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QReadWriteLock struct {
	*qtrt.CObject
}
type QReadWriteLock_ITF interface {
	QReadWriteLock_PTR() *QReadWriteLock
}

func (ptr *QReadWriteLock) QReadWriteLock_PTR() *QReadWriteLock { return ptr }

func (this *QReadWriteLock) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QReadWriteLock) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQReadWriteLockFromPointer(cthis unsafe.Pointer) *QReadWriteLock {
	return &QReadWriteLock{&qtrt.CObject{cthis}}
}
func (*QReadWriteLock) NewFromPointer(cthis unsafe.Pointer) *QReadWriteLock {
	return NewQReadWriteLockFromPointer(cthis)
}

/*


This enum was introduced or modified in  Qt 4.4.

See also QReadWriteLock().

*/
type QReadWriteLock__RecursionMode = int

// In this mode, a thread may only lock a QReadWriteLock once.
const QReadWriteLock__NonRecursive QReadWriteLock__RecursionMode = 0

// In this mode, a thread can lock the same QReadWriteLock multiple times. The QReadWriteLock won't be unlocked until a corresponding number of unlock() calls have been made.
const QReadWriteLock__Recursive QReadWriteLock__RecursionMode = 1

func (this *QReadWriteLock) RecursionModeItemName(val int) string {
	switch val {
	case QReadWriteLock__NonRecursive: // 0
		return "NonRecursive"
	case QReadWriteLock__Recursive: // 1
		return "Recursive"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QReadWriteLock_RecursionModeItemName(val int) string {
	var nilthis *QReadWriteLock
	return nilthis.RecursionModeItemName(val)
}

/*


 */
type QReadWriteLock__StateForWaitCondition = int

//
const QReadWriteLock__LockedForRead QReadWriteLock__StateForWaitCondition = 0

//
const QReadWriteLock__LockedForWrite QReadWriteLock__StateForWaitCondition = 1

//
const QReadWriteLock__Unlocked QReadWriteLock__StateForWaitCondition = 2

//
const QReadWriteLock__RecursivelyLocked QReadWriteLock__StateForWaitCondition = 3

func (this *QReadWriteLock) StateForWaitConditionItemName(val int) string {
	switch val {
	case QReadWriteLock__LockedForRead: // 0
		return "LockedForRead"
	case QReadWriteLock__LockedForWrite: // 1
		return "LockedForWrite"
	case QReadWriteLock__Unlocked: // 2
		return "Unlocked"
	case QReadWriteLock__RecursivelyLocked: // 3
		return "RecursivelyLocked"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QReadWriteLock_StateForWaitConditionItemName(val int) string {
	var nilthis *QReadWriteLock
	return nilthis.StateForWaitConditionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10513() {
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
