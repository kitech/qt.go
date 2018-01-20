//  header block begin
// /usr/include/qt/QtCore/qfuturewatcher.h
// #include <qfuturewatcher.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
type QFutureWatcherBase struct {
	*QObject
}

func (this *QFutureWatcherBase) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQFutureWatcherBaseFromPointer(cthis unsafe.Pointer) *QFutureWatcherBase {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFutureWatcherBase{bcthis0}
}

// /usr/include/qt/QtCore/qfuturewatcher.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QFutureWatcherBase) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:61
// index:0
// Public
// void QFutureWatcherBase(class QObject *)
func NewQFutureWatcherBase(parent unsafe.Pointer) *QFutureWatcherBase {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBaseC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFutureWatcherBaseFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfuturewatcher.h:64
// index:0
// Public
// int progressValue()
func (this *QFutureWatcherBase) ProgressValue() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase13progressValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:65
// index:0
// Public
// int progressMinimum()
func (this *QFutureWatcherBase) ProgressMinimum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase15progressMinimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:66
// index:0
// Public
// int progressMaximum()
func (this *QFutureWatcherBase) ProgressMaximum() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase15progressMaximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:67
// index:0
// Public
// QString progressText()
func (this *QFutureWatcherBase) ProgressText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase12progressTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:69
// index:0
// Public
// bool isStarted()
func (this *QFutureWatcherBase) IsStarted() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase9isStartedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:70
// index:0
// Public
// bool isFinished()
func (this *QFutureWatcherBase) IsFinished() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10isFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:71
// index:0
// Public
// bool isRunning()
func (this *QFutureWatcherBase) IsRunning() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase9isRunningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:72
// index:0
// Public
// bool isCanceled()
func (this *QFutureWatcherBase) IsCanceled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10isCanceledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:73
// index:0
// Public
// bool isPaused()
func (this *QFutureWatcherBase) IsPaused() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase8isPausedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:75
// index:0
// Public
// void waitForFinished()
func (this *QFutureWatcherBase) WaitForFinished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase15waitForFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:77
// index:0
// Public
// void setPendingResultsLimit(int)
func (this *QFutureWatcherBase) SetPendingResultsLimit(limit int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase22setPendingResultsLimitEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &limit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:79
// index:0
// Public virtual
// bool event(class QEvent *)
func (this *QFutureWatcherBase) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qfuturewatcher.h:82
// index:0
// Public
// void started()
func (this *QFutureWatcherBase) Started() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase7startedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:83
// index:0
// Public
// void finished()
func (this *QFutureWatcherBase) Finished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase8finishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:84
// index:0
// Public
// void canceled()
func (this *QFutureWatcherBase) Canceled() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase8canceledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:85
// index:0
// Public
// void paused()
func (this *QFutureWatcherBase) Paused() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase6pausedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:86
// index:0
// Public
// void resumed()
func (this *QFutureWatcherBase) Resumed() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase7resumedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:87
// index:0
// Public
// void resultReadyAt(int)
func (this *QFutureWatcherBase) ResultReadyAt(resultIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase13resultReadyAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &resultIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:88
// index:0
// Public
// void resultsReadyAt(int, int)
func (this *QFutureWatcherBase) ResultsReadyAt(beginIndex int, endIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase14resultsReadyAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &beginIndex, &endIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:89
// index:0
// Public
// void progressRangeChanged(int, int)
func (this *QFutureWatcherBase) ProgressRangeChanged(minimum int, maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase20progressRangeChangedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &minimum, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:90
// index:0
// Public
// void progressValueChanged(int)
func (this *QFutureWatcherBase) ProgressValueChanged(progressValue int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase20progressValueChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &progressValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:91
// index:0
// Public
// void progressTextChanged(const class QString &)
func (this *QFutureWatcherBase) ProgressTextChanged(progressText *QString) {
	var convArg0 = progressText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase19progressTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:94
// index:0
// Public
// void cancel()
func (this *QFutureWatcherBase) Cancel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase6cancelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:95
// index:0
// Public
// void setPaused(_Bool)
func (this *QFutureWatcherBase) SetPaused(paused bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase9setPausedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:96
// index:0
// Public
// void pause()
func (this *QFutureWatcherBase) Pause() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase5pauseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:97
// index:0
// Public
// void resume()
func (this *QFutureWatcherBase) Resume() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase6resumeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:98
// index:0
// Public
// void togglePaused()
func (this *QFutureWatcherBase) TogglePaused() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase12togglePausedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:101
// index:0
// Protected virtual
// void connectNotify(const class QMetaMethod &)
func (this *QFutureWatcherBase) ConnectNotify(signal *QMetaMethod) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase13connectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:102
// index:0
// Protected virtual
// void disconnectNotify(const class QMetaMethod &)
func (this *QFutureWatcherBase) DisconnectNotify(signal *QMetaMethod) {
	var convArg0 = signal.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase16disconnectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:105
// index:0
// Protected
// void connectOutputInterface()
func (this *QFutureWatcherBase) ConnectOutputInterface() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase22connectOutputInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:106
// index:0
// Protected
// void disconnectOutputInterface(_Bool)
func (this *QFutureWatcherBase) DisconnectOutputInterface(pendingAssignment bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase25disconnectOutputInterfaceEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pendingAssignment)
	gopp.ErrPrint(err, rv)
}

//  body block end
