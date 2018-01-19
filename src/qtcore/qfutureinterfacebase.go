//  header block begin
// /usr/include/qt/QtCore/qfutureinterface.h
// #include <qfutureinterface.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QFutureInterfaceBase struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qfutureinterface.h:73
// index:0
// void QFutureInterfaceBase(enum QFutureInterfaceBase::State)
func NewQFutureInterfaceBase(initialState int) *QFutureInterfaceBase {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBaseC2ENS_5StateE", ffiqt.FFI_TYPE_VOID, cthis, &initialState)
	gopp.ErrPrint(err, rv)
	return &QFutureInterfaceBase{cthis}
}

// /usr/include/qt/QtCore/qfutureinterface.h:75
// index:0
// virtual
// void ~QFutureInterfaceBase()
func DeleteQFutureInterfaceBase(*QFutureInterfaceBase) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBaseD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:78
// index:0
// void reportStarted()
func (this *QFutureInterfaceBase) ReportStarted() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13reportStartedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:79
// index:0
// void reportFinished()
func (this *QFutureInterfaceBase) ReportFinished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase14reportFinishedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:80
// index:0
// void reportCanceled()
func (this *QFutureInterfaceBase) ReportCanceled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase14reportCanceledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:84
// index:0
// void reportResultsReady(int, int)
func (this *QFutureInterfaceBase) ReportResultsReady(beginIndex int, endIndex int) {
	// 0: (, int beginIndex, int endIndex), (&beginIndex, &endIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase18reportResultsReadyEii", ffiqt.FFI_TYPE_VOID, this.cthis, &beginIndex, &endIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:86
// index:0
// void setRunnable(class QRunnable *)
func (this *QFutureInterfaceBase) SetRunnable(runnable unsafe.Pointer) {
	// 0: (, QRunnable * runnable), (runnable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase11setRunnableEP9QRunnable", ffiqt.FFI_TYPE_VOID, this.cthis, runnable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:87
// index:0
// void setThreadPool(class QThreadPool *)
func (this *QFutureInterfaceBase) SetThreadPool(pool unsafe.Pointer) {
	// 0: (, QThreadPool * pool), (pool)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13setThreadPoolEP11QThreadPool", ffiqt.FFI_TYPE_VOID, this.cthis, pool)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:88
// index:0
// void setFilterMode(_Bool)
func (this *QFutureInterfaceBase) SetFilterMode(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13setFilterModeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:89
// index:0
// void setProgressRange(int, int)
func (this *QFutureInterfaceBase) SetProgressRange(minimum int, maximum int) {
	// 0: (, int minimum, int maximum), (&minimum, &maximum)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase16setProgressRangeEii", ffiqt.FFI_TYPE_VOID, this.cthis, &minimum, &maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:90
// index:0
// int progressMinimum()
func (this *QFutureInterfaceBase) ProgressMinimum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15progressMinimumEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:91
// index:0
// int progressMaximum()
func (this *QFutureInterfaceBase) ProgressMaximum() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15progressMaximumEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:92
// index:0
// bool isProgressUpdateNeeded()
func (this *QFutureInterfaceBase) IsProgressUpdateNeeded() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase22isProgressUpdateNeededEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:93
// index:0
// void setProgressValue(int)
func (this *QFutureInterfaceBase) SetProgressValue(progressValue int) {
	// 0: (, int progressValue), (&progressValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase16setProgressValueEi", ffiqt.FFI_TYPE_VOID, this.cthis, &progressValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:94
// index:0
// int progressValue()
func (this *QFutureInterfaceBase) ProgressValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase13progressValueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:95
// index:0
// void setProgressValueAndText(int, const class QString &)
func (this *QFutureInterfaceBase) SetProgressValueAndText(progressValue int, progressText unsafe.Pointer) {
	// 0: (, int progressValue, const QString & progressText), (&progressValue, progressText)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase23setProgressValueAndTextEiRK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, &progressValue, progressText)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:96
// index:0
// QString progressText()
func (this *QFutureInterfaceBase) ProgressText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase12progressTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:98
// index:0
// void setExpectedResultCount(int)
func (this *QFutureInterfaceBase) SetExpectedResultCount(resultCount int) {
	// 0: (, int resultCount), (&resultCount)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase22setExpectedResultCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &resultCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:99
// index:0
// int expectedResultCount()
func (this *QFutureInterfaceBase) ExpectedResultCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase19expectedResultCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:100
// index:0
// int resultCount()
func (this *QFutureInterfaceBase) ResultCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase11resultCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:102
// index:0
// bool queryState(enum QFutureInterfaceBase::State)
func (this *QFutureInterfaceBase) QueryState(state int) {
	// 0: (, QFutureInterfaceBase::State state), (&state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase10queryStateENS_5StateE", ffiqt.FFI_TYPE_VOID, this.cthis, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:103
// index:0
// bool isRunning()
func (this *QFutureInterfaceBase) IsRunning() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase9isRunningEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:104
// index:0
// bool isStarted()
func (this *QFutureInterfaceBase) IsStarted() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase9isStartedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:105
// index:0
// bool isCanceled()
func (this *QFutureInterfaceBase) IsCanceled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase10isCanceledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:106
// index:0
// bool isFinished()
func (this *QFutureInterfaceBase) IsFinished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase10isFinishedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:107
// index:0
// bool isPaused()
func (this *QFutureInterfaceBase) IsPaused() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase8isPausedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:108
// index:0
// bool isThrottled()
func (this *QFutureInterfaceBase) IsThrottled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase11isThrottledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:109
// index:0
// bool isResultReadyAt(int)
func (this *QFutureInterfaceBase) IsResultReadyAt(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15isResultReadyAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:111
// index:0
// void cancel()
func (this *QFutureInterfaceBase) Cancel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase6cancelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:112
// index:0
// void setPaused(_Bool)
func (this *QFutureInterfaceBase) SetPaused(paused bool) {
	// 0: (, bool paused), (&paused)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase9setPausedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:113
// index:0
// void togglePaused()
func (this *QFutureInterfaceBase) TogglePaused() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase12togglePausedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:114
// index:0
// void setThrottled(_Bool)
func (this *QFutureInterfaceBase) SetThrottled(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase12setThrottledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:116
// index:0
// void waitForFinished()
func (this *QFutureInterfaceBase) WaitForFinished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase15waitForFinishedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:117
// index:0
// bool waitForNextResult()
func (this *QFutureInterfaceBase) WaitForNextResult() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase17waitForNextResultEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:118
// index:0
// void waitForResult(int)
func (this *QFutureInterfaceBase) WaitForResult(resultIndex int) {
	// 0: (, int resultIndex), (&resultIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13waitForResultEi", ffiqt.FFI_TYPE_VOID, this.cthis, &resultIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:119
// index:0
// void waitForResume()
func (this *QFutureInterfaceBase) WaitForResume() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13waitForResumeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:121
// index:0
// QMutex * mutex()
func (this *QFutureInterfaceBase) Mutex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase5mutexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:122
// index:0
// QtPrivate::ExceptionStore & exceptionStore()
func (this *QFutureInterfaceBase) ExceptionStore() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase14exceptionStoreEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:123
// index:0
// QtPrivate::ResultStoreBase & resultStoreBase()
func (this *QFutureInterfaceBase) ResultStoreBase() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase15resultStoreBaseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:124
// index:1
// const QtPrivate::ResultStoreBase & resultStoreBase()
func (this *QFutureInterfaceBase) ResultStoreBase_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15resultStoreBaseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
