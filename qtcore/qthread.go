package qtcore

// /usr/include/qt/QtCore/qthread.h
// #include <qthread.h>
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

/*
 */
// size 16
type QThread struct {
	*QObject
}
type QThread_ITF interface {
	QObject_ITF
	QThread_PTR() *QThread
}

func (ptr *QThread) QThread_PTR() *QThread { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QThreadFromptr(cthis Voidptr) *QThread {
	bcthis0 := QObjectFromptr(cthis)
	return &QThread{bcthis0}
}
func (*QThread) Fromptr(cthis Voidptr) *QThread {
	return QThreadFromptr(cthis)
}

func DeleteQThread(this *QThread) {
	rv, err := qtrt.Qtcc1(0, "_ZN7QThreadD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QThread__Priority = int

//
const QThread__IdlePriority QThread__Priority = 0

//
const QThread__LowestPriority QThread__Priority = 1

//
const QThread__LowPriority QThread__Priority = 2

//
const QThread__NormalPriority QThread__Priority = 3

//
const QThread__HighPriority QThread__Priority = 4

//
const QThread__HighestPriority QThread__Priority = 5

//
const QThread__TimeCriticalPriority QThread__Priority = 6

//
const QThread__InheritPriority QThread__Priority = 7

func (this *QThread) PriorityItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QThread_PriorityItemName(val int) string {
	var nilthis *QThread
	return nilthis.PriorityItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10055() {
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
