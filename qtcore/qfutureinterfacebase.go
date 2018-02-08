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
import "gopp"
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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool refT()
func (this *QFutureInterfaceBase) InheritRefT(f func() bool) {
	qtrt.SetAllInheritCallback(this, "refT", f)
}

// bool derefT()
func (this *QFutureInterfaceBase) InheritDerefT(f func() bool) {
	qtrt.SetAllInheritCallback(this, "derefT", f)
}

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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFutureInterfaceBaseFromPointer(cthis unsafe.Pointer) *QFutureInterfaceBase {
	return &QFutureInterfaceBase{&qtrt.CObject{cthis}}
}
func (*QFutureInterfaceBase) NewFromPointer(cthis unsafe.Pointer) *QFutureInterfaceBase {
	return NewQFutureInterfaceBaseFromPointer(cthis)
}

// /usr/include/qt/QtCore/qfutureinterface.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFutureInterfaceBase(enum QFutureInterfaceBase::State)
func NewQFutureInterfaceBase(initialState int) *QFutureInterfaceBase {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBaseC2ENS_5StateE", qtrt.FFI_TYPE_POINTER, initialState)
	gopp.ErrPrint(err, rv)
	gothis := NewQFutureInterfaceBaseFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFutureInterfaceBase)
	return gothis
}

// /usr/include/qt/QtCore/qfutureinterface.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QFutureInterfaceBase()
func DeleteQFutureInterfaceBase(this *QFutureInterfaceBase) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBaseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qfutureinterface.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reportStarted()
func (this *QFutureInterfaceBase) ReportStarted() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13reportStartedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reportFinished()
func (this *QFutureInterfaceBase) ReportFinished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase14reportFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reportCanceled()
func (this *QFutureInterfaceBase) ReportCanceled() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase14reportCanceledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reportResultsReady(int, int)
func (this *QFutureInterfaceBase) ReportResultsReady(beginIndex int, endIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase18reportResultsReadyEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), beginIndex, endIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRunnable(QRunnable *)
func (this *QFutureInterfaceBase) SetRunnable(runnable *QRunnable /*777 QRunnable **/) {
	var convArg0 = runnable.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase11setRunnableEP9QRunnable", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setThreadPool(QThreadPool *)
func (this *QFutureInterfaceBase) SetThreadPool(pool *QThreadPool /*777 QThreadPool **/) {
	var convArg0 = pool.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13setThreadPoolEP11QThreadPool", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilterMode(_Bool)
func (this *QFutureInterfaceBase) SetFilterMode(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13setFilterModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProgressRange(int, int)
func (this *QFutureInterfaceBase) SetProgressRange(minimum int, maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase16setProgressRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum, maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:90
// index:0
// Public Visibility=Default Availability=Available
// [4] int progressMinimum()
func (this *QFutureInterfaceBase) ProgressMinimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15progressMinimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfutureinterface.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] int progressMaximum()
func (this *QFutureInterfaceBase) ProgressMaximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15progressMaximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfutureinterface.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isProgressUpdateNeeded()
func (this *QFutureInterfaceBase) IsProgressUpdateNeeded() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase22isProgressUpdateNeededEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProgressValue(int)
func (this *QFutureInterfaceBase) SetProgressValue(progressValue int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase16setProgressValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), progressValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] int progressValue()
func (this *QFutureInterfaceBase) ProgressValue() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase13progressValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfutureinterface.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProgressValueAndText(int, const QString &)
func (this *QFutureInterfaceBase) SetProgressValueAndText(progressValue int, progressText string) {
	var tmpArg1 = NewQString_5(progressText)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase23setProgressValueAndTextEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), progressValue, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QString progressText()
func (this *QFutureInterfaceBase) ProgressText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase12progressTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qfutureinterface.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExpectedResultCount(int)
func (this *QFutureInterfaceBase) SetExpectedResultCount(resultCount int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase22setExpectedResultCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resultCount)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int expectedResultCount()
func (this *QFutureInterfaceBase) ExpectedResultCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase19expectedResultCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfutureinterface.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int resultCount()
func (this *QFutureInterfaceBase) ResultCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase11resultCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qfutureinterface.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool queryState(enum QFutureInterfaceBase::State)
func (this *QFutureInterfaceBase) QueryState(state int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase10queryStateENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRunning()
func (this *QFutureInterfaceBase) IsRunning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase9isRunningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:104
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStarted()
func (this *QFutureInterfaceBase) IsStarted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase9isStartedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCanceled()
func (this *QFutureInterfaceBase) IsCanceled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase10isCanceledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinished()
func (this *QFutureInterfaceBase) IsFinished() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase10isFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPaused()
func (this *QFutureInterfaceBase) IsPaused() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase8isPausedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isThrottled()
func (this *QFutureInterfaceBase) IsThrottled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase11isThrottledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isResultReadyAt(int)
func (this *QFutureInterfaceBase) IsResultReadyAt(index int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15isResultReadyAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancel()
func (this *QFutureInterfaceBase) Cancel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase6cancelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaused(_Bool)
func (this *QFutureInterfaceBase) SetPaused(paused bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase9setPausedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void togglePaused()
func (this *QFutureInterfaceBase) TogglePaused() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase12togglePausedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setThrottled(_Bool)
func (this *QFutureInterfaceBase) SetThrottled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase12setThrottledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void waitForFinished()
func (this *QFutureInterfaceBase) WaitForFinished() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase15waitForFinishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForNextResult()
func (this *QFutureInterfaceBase) WaitForNextResult() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase17waitForNextResultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void waitForResult(int)
func (this *QFutureInterfaceBase) WaitForResult(resultIndex int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13waitForResultEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resultIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void waitForResume()
func (this *QFutureInterfaceBase) WaitForResume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase13waitForResumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qfutureinterface.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] QMutex * mutex()
func (this *QFutureInterfaceBase) Mutex() *QMutex /*777 QMutex **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase5mutexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQMutexFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qfutureinterface.h:122
// index:0
// Public Visibility=Default Availability=Available
// [1] QtPrivate::ExceptionStore & exceptionStore()
func (this *QFutureInterfaceBase) ExceptionStore() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase14exceptionStoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

// /usr/include/qt/QtCore/qfutureinterface.h:123
// index:0
// Public Visibility=Default Availability=Available
// [48] QtPrivate::ResultStoreBase & resultStoreBase()
func (this *QFutureInterfaceBase) ResultStoreBase() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QFutureInterfaceBase15resultStoreBaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

// /usr/include/qt/QtCore/qfutureinterface.h:124
// index:1
// Public Visibility=Default Availability=Available
// [48] const QtPrivate::ResultStoreBase & resultStoreBase()
func (this *QFutureInterfaceBase) ResultStoreBase_1() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase15resultStoreBaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cpretval2go("int", rv).(int) // 3331
}

// /usr/include/qt/QtCore/qfutureinterface.h:131
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool refT()
func (this *QFutureInterfaceBase) RefT() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase4refTEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qfutureinterface.h:132
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool derefT()
func (this *QFutureInterfaceBase) DerefT() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QFutureInterfaceBase6derefTEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
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
