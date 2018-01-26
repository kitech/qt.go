package qtcore

// /usr/include/qt/QtCore/qfutureinterface.h
// #include <qfutureinterface.h>
// #include <QtCore>

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
	*qtrt.CObject
}

func (this *QFutureInterfaceBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFutureInterfaceBase) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQFutureInterfaceBaseFromPointer(cthis unsafe.Pointer) *QFutureInterfaceBase {
	return &QFutureInterfaceBase{&qtrt.CObject{cthis}}
}
func (*QFutureInterfaceBase) NewFromPointer(cthis unsafe.Pointer) *QFutureInterfaceBase {
	return NewQFutureInterfaceBaseFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfutureinterface.h:73
// index:0
// Public
// void QFutureInterfaceBase(enum QFutureInterfaceBase::State)
func NewQFutureInterfaceBase(initialState int) *QFutureInterfaceBase {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBaseC2ENS_5StateE", ffiqt.FFI_TYPE_VOID, cthis, initialState)
	gopp.ErrPrint(err, rv)
	gothis := NewQFutureInterfaceBaseFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qfutureinterface.h:75
// index:0
// Public virtual
// void ~QFutureInterfaceBase()
func DeleteQFutureInterfaceBase(*QFutureInterfaceBase) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBaseD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:78
// index:0
// Public
// void reportStarted()
func (this *QFutureInterfaceBase) ReportStarted() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13reportStartedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:79
// index:0
// Public
// void reportFinished()
func (this *QFutureInterfaceBase) ReportFinished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase14reportFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:80
// index:0
// Public
// void reportCanceled()
func (this *QFutureInterfaceBase) ReportCanceled() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase14reportCanceledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:84
// index:0
// Public
// void reportResultsReady(int, int)
func (this *QFutureInterfaceBase) ReportResultsReady(beginIndex int, endIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase18reportResultsReadyEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), beginIndex, endIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:86
// index:0
// Public
// void setRunnable(class QRunnable *)
func (this *QFutureInterfaceBase) SetRunnable(runnable *QRunnable /*777 QRunnable **/) {
	var convArg0 = runnable.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase11setRunnableEP9QRunnable", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:87
// index:0
// Public
// void setThreadPool(class QThreadPool *)
func (this *QFutureInterfaceBase) SetThreadPool(pool *QThreadPool /*777 QThreadPool **/) {
	var convArg0 = pool.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13setThreadPoolEP11QThreadPool", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:88
// index:0
// Public
// void setFilterMode(_Bool)
func (this *QFutureInterfaceBase) SetFilterMode(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13setFilterModeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:89
// index:0
// Public
// void setProgressRange(int, int)
func (this *QFutureInterfaceBase) SetProgressRange(minimum int, maximum int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase16setProgressRangeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), minimum, maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:90
// index:0
// Public
// int progressMinimum()
func (this *QFutureInterfaceBase) ProgressMinimum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15progressMinimumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qfutureinterface.h:91
// index:0
// Public
// int progressMaximum()
func (this *QFutureInterfaceBase) ProgressMaximum() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15progressMaximumEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qfutureinterface.h:92
// index:0
// Public
// bool isProgressUpdateNeeded()
func (this *QFutureInterfaceBase) IsProgressUpdateNeeded() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase22isProgressUpdateNeededEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:93
// index:0
// Public
// void setProgressValue(int)
func (this *QFutureInterfaceBase) SetProgressValue(progressValue int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase16setProgressValueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), progressValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:94
// index:0
// Public
// int progressValue()
func (this *QFutureInterfaceBase) ProgressValue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase13progressValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qfutureinterface.h:95
// index:0
// Public
// void setProgressValueAndText(int, const class QString &)
func (this *QFutureInterfaceBase) SetProgressValueAndText(progressValue int, progressText *QString) {
	var convArg1 = progressText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase23setProgressValueAndTextEiRK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), progressValue, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:96
// index:0
// Public
// QString progressText()
func (this *QFutureInterfaceBase) ProgressText() *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase12progressTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qfutureinterface.h:98
// index:0
// Public
// void setExpectedResultCount(int)
func (this *QFutureInterfaceBase) SetExpectedResultCount(resultCount int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase22setExpectedResultCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), resultCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:99
// index:0
// Public
// int expectedResultCount()
func (this *QFutureInterfaceBase) ExpectedResultCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase19expectedResultCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qfutureinterface.h:100
// index:0
// Public
// int resultCount()
func (this *QFutureInterfaceBase) ResultCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase11resultCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qfutureinterface.h:102
// index:0
// Public
// bool queryState(enum QFutureInterfaceBase::State)
func (this *QFutureInterfaceBase) QueryState(state int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase10queryStateENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:103
// index:0
// Public
// bool isRunning()
func (this *QFutureInterfaceBase) IsRunning() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase9isRunningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:104
// index:0
// Public
// bool isStarted()
func (this *QFutureInterfaceBase) IsStarted() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase9isStartedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:105
// index:0
// Public
// bool isCanceled()
func (this *QFutureInterfaceBase) IsCanceled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase10isCanceledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:106
// index:0
// Public
// bool isFinished()
func (this *QFutureInterfaceBase) IsFinished() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase10isFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:107
// index:0
// Public
// bool isPaused()
func (this *QFutureInterfaceBase) IsPaused() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase8isPausedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:108
// index:0
// Public
// bool isThrottled()
func (this *QFutureInterfaceBase) IsThrottled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase11isThrottledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:109
// index:0
// Public
// bool isResultReadyAt(int)
func (this *QFutureInterfaceBase) IsResultReadyAt(index int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15isResultReadyAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:111
// index:0
// Public
// void cancel()
func (this *QFutureInterfaceBase) Cancel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase6cancelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:112
// index:0
// Public
// void setPaused(_Bool)
func (this *QFutureInterfaceBase) SetPaused(paused bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase9setPausedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:113
// index:0
// Public
// void togglePaused()
func (this *QFutureInterfaceBase) TogglePaused() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase12togglePausedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:114
// index:0
// Public
// void setThrottled(_Bool)
func (this *QFutureInterfaceBase) SetThrottled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase12setThrottledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:116
// index:0
// Public
// void waitForFinished()
func (this *QFutureInterfaceBase) WaitForFinished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase15waitForFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:117
// index:0
// Public
// bool waitForNextResult()
func (this *QFutureInterfaceBase) WaitForNextResult() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase17waitForNextResultEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:118
// index:0
// Public
// void waitForResult(int)
func (this *QFutureInterfaceBase) WaitForResult(resultIndex int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13waitForResultEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), resultIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:119
// index:0
// Public
// void waitForResume()
func (this *QFutureInterfaceBase) WaitForResume() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13waitForResumeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:121
// index:0
// Public
// QMutex * mutex()
func (this *QFutureInterfaceBase) Mutex() *QMutex /*777 QMutex **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase5mutexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMutexFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qfutureinterface.h:122
// index:0
// Public
// QtPrivate::ExceptionStore & exceptionStore()
func (this *QFutureInterfaceBase) ExceptionStore() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase14exceptionStoreEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 3331
}

// /usr/include/qt/QtCore/qfutureinterface.h:123
// index:0
// Public
// QtPrivate::ResultStoreBase & resultStoreBase()
func (this *QFutureInterfaceBase) ResultStoreBase() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QFutureInterfaceBase15resultStoreBaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 3331
}

// /usr/include/qt/QtCore/qfutureinterface.h:124
// index:1
// Public
// const QtPrivate::ResultStoreBase & resultStoreBase()
func (this *QFutureInterfaceBase) ResultStoreBase_1() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15resultStoreBaseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 3331
}

// /usr/include/qt/QtCore/qfutureinterface.h:131
// index:0
// Protected
// bool refT()
func (this *QFutureInterfaceBase) RefT() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase4refTEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:132
// index:0
// Protected
// bool derefT()
func (this *QFutureInterfaceBase) DerefT() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase6derefTEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QFutureInterfaceBase__State = int

const QFutureInterfaceBase__NoState QFutureInterfaceBase__State = 0
const QFutureInterfaceBase__Running QFutureInterfaceBase__State = 1
const QFutureInterfaceBase__Started QFutureInterfaceBase__State = 2
const QFutureInterfaceBase__Finished QFutureInterfaceBase__State = 4
const QFutureInterfaceBase__Canceled QFutureInterfaceBase__State = 8
const QFutureInterfaceBase__Paused QFutureInterfaceBase__State = 16
const QFutureInterfaceBase__Throttled QFutureInterfaceBase__State = 32

//  body block end
