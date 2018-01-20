//  header block begin
// /usr/include/qt/QtCore/qabstractanimation.h
// #include <qabstractanimation.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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
type QAnimationDriver struct {
	*QObject
}

func (this *QAnimationDriver) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQAnimationDriverFromPointer(cthis unsafe.Pointer) *QAnimationDriver {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAnimationDriver{bcthis0}
}

// /usr/include/qt/QtCore/qabstractanimation.h:135
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAnimationDriver) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAnimationDriver10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractanimation.h:139
// index:0
// Public
// void QAnimationDriver(class QObject *)
func NewQAnimationDriver(parent unsafe.Pointer) *QAnimationDriver {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriverC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAnimationDriverFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qabstractanimation.h:140
// index:0
// Public virtual
// void ~QAnimationDriver()
func DeleteQAnimationDriver(*QAnimationDriver) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriverD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:142
// index:0
// Public virtual
// void advance()
func (this *QAnimationDriver) Advance() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver7advanceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:144
// index:0
// Public
// void install()
func (this *QAnimationDriver) Install() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver7installEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:145
// index:0
// Public
// void uninstall()
func (this *QAnimationDriver) Uninstall() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver9uninstallEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:147
// index:0
// Public
// bool isRunning()
func (this *QAnimationDriver) IsRunning() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAnimationDriver9isRunningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractanimation.h:149
// index:0
// Public virtual
// qint64 elapsed()
func (this *QAnimationDriver) Elapsed() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAnimationDriver7elapsedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractanimation.h:152
// index:0
// Public
// void setStartTime(qint64)
func (this *QAnimationDriver) SetStartTime(startTime int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver12setStartTimeEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &startTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:153
// index:0
// Public
// qint64 startTime()
func (this *QAnimationDriver) StartTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QAnimationDriver9startTimeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qabstractanimation.h:156
// index:0
// Public
// void started()
func (this *QAnimationDriver) Started() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver7startedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:157
// index:0
// Public
// void stopped()
func (this *QAnimationDriver) Stopped() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver7stoppedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:161
// index:0
// Protected
// void advanceAnimation(qint64)
func (this *QAnimationDriver) AdvanceAnimation(timeStep int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver16advanceAnimationEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &timeStep)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:162
// index:0
// Protected virtual
// void start()
func (this *QAnimationDriver) Start() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver5startEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractanimation.h:163
// index:0
// Protected virtual
// void stop()
func (this *QAnimationDriver) Stop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QAnimationDriver4stopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
