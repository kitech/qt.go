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

// /usr/include/qt/QtCore/qfuturewatcher.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QFutureWatcherBase) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:61
// index:0
// void QFutureWatcherBase(class QObject *)
func NewQFutureWatcherBase(parent unsafe.Pointer) *QFutureWatcherBase {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBaseC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQFutureWatcherBaseFromPointer(cthis)
	return gothis
}
func NewQFutureWatcherBaseFromPointer(cthis unsafe.Pointer) *QFutureWatcherBase {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QFutureWatcherBase{bcthis0}
}

// /usr/include/qt/QtCore/qfuturewatcher.h:64
// index:0
// int progressValue()
func (this *QFutureWatcherBase) ProgressValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase13progressValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:65
// index:0
// int progressMinimum()
func (this *QFutureWatcherBase) ProgressMinimum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase15progressMinimumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:66
// index:0
// int progressMaximum()
func (this *QFutureWatcherBase) ProgressMaximum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase15progressMaximumEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:67
// index:0
// QString progressText()
func (this *QFutureWatcherBase) ProgressText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase12progressTextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:69
// index:0
// bool isStarted()
func (this *QFutureWatcherBase) IsStarted() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase9isStartedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:70
// index:0
// bool isFinished()
func (this *QFutureWatcherBase) IsFinished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10isFinishedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:71
// index:0
// bool isRunning()
func (this *QFutureWatcherBase) IsRunning() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase9isRunningEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:72
// index:0
// bool isCanceled()
func (this *QFutureWatcherBase) IsCanceled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase10isCanceledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:73
// index:0
// bool isPaused()
func (this *QFutureWatcherBase) IsPaused() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QFutureWatcherBase8isPausedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:75
// index:0
// void waitForFinished()
func (this *QFutureWatcherBase) WaitForFinished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase15waitForFinishedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:77
// index:0
// void setPendingResultsLimit(int)
func (this *QFutureWatcherBase) SetPendingResultsLimit(limit int) {
	// 0: (, limit int), (&limit)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase22setPendingResultsLimitEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &limit)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:79
// index:0
// virtual
// bool event(class QEvent *)
func (this *QFutureWatcherBase) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:82
// index:0
// void started()
func (this *QFutureWatcherBase) Started() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase7startedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:83
// index:0
// void finished()
func (this *QFutureWatcherBase) Finished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase8finishedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:84
// index:0
// void canceled()
func (this *QFutureWatcherBase) Canceled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase8canceledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:85
// index:0
// void paused()
func (this *QFutureWatcherBase) Paused() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase6pausedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:86
// index:0
// void resumed()
func (this *QFutureWatcherBase) Resumed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase7resumedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:87
// index:0
// void resultReadyAt(int)
func (this *QFutureWatcherBase) ResultReadyAt(resultIndex int) {
	// 0: (, resultIndex int), (&resultIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase13resultReadyAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &resultIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:88
// index:0
// void resultsReadyAt(int, int)
func (this *QFutureWatcherBase) ResultsReadyAt(beginIndex int, endIndex int) {
	// 0: (, beginIndex int, endIndex int), (&beginIndex, &endIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase14resultsReadyAtEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &beginIndex, &endIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:89
// index:0
// void progressRangeChanged(int, int)
func (this *QFutureWatcherBase) ProgressRangeChanged(minimum int, maximum int) {
	// 0: (, minimum int, maximum int), (&minimum, &maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase20progressRangeChangedEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &minimum, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:90
// index:0
// void progressValueChanged(int)
func (this *QFutureWatcherBase) ProgressValueChanged(progressValue int) {
	// 0: (, progressValue int), (&progressValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase20progressValueChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &progressValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:91
// index:0
// void progressTextChanged(const class QString &)
func (this *QFutureWatcherBase) ProgressTextChanged(progressText unsafe.Pointer) {
	// 0: (, progressText const QString &), (progressText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase19progressTextChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), progressText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:94
// index:0
// void cancel()
func (this *QFutureWatcherBase) Cancel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase6cancelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:95
// index:0
// void setPaused(_Bool)
func (this *QFutureWatcherBase) SetPaused(paused bool) {
	// 0: (, paused bool), (&paused)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase9setPausedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:96
// index:0
// void pause()
func (this *QFutureWatcherBase) Pause() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase5pauseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:97
// index:0
// void resume()
func (this *QFutureWatcherBase) Resume() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase6resumeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:98
// index:0
// void togglePaused()
func (this *QFutureWatcherBase) TogglePaused() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase12togglePausedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:101
// index:0
// virtual
// void connectNotify(const class QMetaMethod &)
func (this *QFutureWatcherBase) ConnectNotify(signal unsafe.Pointer) {
	// 0: (, signal const QMetaMethod &), (signal)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase13connectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_VOID, this.GetCthis(), signal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:102
// index:0
// virtual
// void disconnectNotify(const class QMetaMethod &)
func (this *QFutureWatcherBase) DisconnectNotify(signal unsafe.Pointer) {
	// 0: (, signal const QMetaMethod &), (signal)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase16disconnectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_VOID, this.GetCthis(), signal)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:105
// index:0
// void connectOutputInterface()
func (this *QFutureWatcherBase) ConnectOutputInterface() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase22connectOutputInterfaceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfuturewatcher.h:106
// index:0
// void disconnectOutputInterface(_Bool)
func (this *QFutureWatcherBase) DisconnectOutputInterface(pendingAssignment bool) {
	// 0: (, pendingAssignment bool), (&pendingAssignment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QFutureWatcherBase25disconnectOutputInterfaceEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pendingAssignment)
	gopp.ErrPrint(err, rv)
}

//  body block end
