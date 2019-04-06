// +build !minimal

package qtcore

// /usr/include/qt/QtCore/qprocess.h
// #include <qprocess.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void setProcessState(QProcess::ProcessState)
func (this *QProcess) InheritSetProcessState(f func(state int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setProcessState", f)
}

// void setupChildProcess()
func (this *QProcess) InheritSetupChildProcess(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "setupChildProcess", f)
}

// long long readData(char *, qint64)
func (this *QProcess) InheritReadData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readData", f)
}

// long long writeData(const char *, qint64)
func (this *QProcess) InheritWriteData(f func(data string, len_ int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

/*

 */
type QProcess struct {
	*QIODevice
}
type QProcess_ITF interface {
	QIODevice_ITF
	QProcess_PTR() *QProcess
}

func (ptr *QProcess) QProcess_PTR() *QProcess { return ptr }

func (this *QProcess) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QIODevice.GetCthis()
	}
}
func (this *QProcess) SetCthis(cthis unsafe.Pointer) {
	this.QIODevice = NewQIODeviceFromPointer(cthis)
}
func NewQProcessFromPointer(cthis unsafe.Pointer) *QProcess {
	bcthis0 := NewQIODeviceFromPointer(cthis)
	return &QProcess{bcthis0}
}
func (*QProcess) NewFromPointer(cthis unsafe.Pointer) *QProcess {
	return NewQProcessFromPointer(cthis)
}

/*
This enum describes the different types of errors that are reported by QProcess.



See also error().

*/
type QProcess__ProcessError = int

// The process failed to start. Either the invoked program is missing, or you may have insufficient permissions to invoke the program.
const QProcess__FailedToStart QProcess__ProcessError = 0

// The process crashed some time after starting successfully.
const QProcess__Crashed QProcess__ProcessError = 1

// The last waitFor...() function timed out. The state of QProcess is unchanged, and you can try calling waitFor...() again.
const QProcess__Timedout QProcess__ProcessError = 2

// An error occurred when attempting to read from the process. For example, the process may not be running.
const QProcess__ReadError QProcess__ProcessError = 3

// An error occurred when attempting to write to the process. For example, the process may not be running, or it may have closed its input channel.
const QProcess__WriteError QProcess__ProcessError = 4

// An unknown error occurred. This is the default return value of error().
const QProcess__UnknownError QProcess__ProcessError = 5

func (this *QProcess) ProcessErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QProcess_ProcessErrorItemName(val int) string {
	var nilthis *QProcess
	return nilthis.ProcessErrorItemName(val)
}

/*
This enum describes the different states of QProcess.



See also state().

*/
type QProcess__ProcessState = int

// The process is not running.
const QProcess__NotRunning QProcess__ProcessState = 0

// The process is starting, but the program has not yet been invoked.
const QProcess__Starting QProcess__ProcessState = 1

// The process is running and is ready for reading and writing.
const QProcess__Running QProcess__ProcessState = 2

func (this *QProcess) ProcessStateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QProcess_ProcessStateItemName(val int) string {
	var nilthis *QProcess
	return nilthis.ProcessStateItemName(val)
}

/*
This enum describes the process channels used by the running process. Pass one of these values to setReadChannel() to set the current read channel of QProcess.



See also setReadChannel().

*/
type QProcess__ProcessChannel = int

// The standard output (stdout) of the running process.
const QProcess__StandardOutput QProcess__ProcessChannel = 0

// The standard error (stderr) of the running process.
const QProcess__StandardError QProcess__ProcessChannel = 1

func (this *QProcess) ProcessChannelItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QProcess_ProcessChannelItemName(val int) string {
	var nilthis *QProcess
	return nilthis.ProcessChannelItemName(val)
}

/*
This enum describes the process output channel modes of QProcess. Pass one of these values to setProcessChannelMode() to set the current read channel mode.



Note: Windows intentionally suppresses output from GUI-only applications to inherited consoles. This does not apply to output redirected to files or pipes. To forward the output of GUI-only applications on the console nonetheless, you must use SeparateChannels and do the forwarding yourself by reading the output and writing it to the appropriate output channels.

See also setProcessChannelMode().

*/
type QProcess__ProcessChannelMode = int

// QProcess manages the output of the running process, keeping standard output and standard error data in separate internal buffers. You can select the QProcess's current read channel by calling setReadChannel(). This is the default channel mode of QProcess.
const QProcess__SeparateChannels QProcess__ProcessChannelMode = 0

// QProcess merges the output of the running process into the standard output channel (stdout). The standard error channel (stderr) will not receive any data. The standard output and standard error data of the running process are interleaved.
const QProcess__MergedChannels QProcess__ProcessChannelMode = 1

// QProcess forwards the output of the running process onto the main process. Anything the child process writes to its standard output and standard error will be written to the standard output and standard error of the main process.
const QProcess__ForwardedChannels QProcess__ProcessChannelMode = 2

//
const QProcess__ForwardedOutputChannel QProcess__ProcessChannelMode = 3

//
const QProcess__ForwardedErrorChannel QProcess__ProcessChannelMode = 4

func (this *QProcess) ProcessChannelModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QProcess_ProcessChannelModeItemName(val int) string {
	var nilthis *QProcess
	return nilthis.ProcessChannelModeItemName(val)
}

/*
This enum describes the process input channel modes of QProcess. Pass one of these values to setInputChannelMode() to set the current write channel mode.



This enum was introduced or modified in  Qt 5.2.

See also setInputChannelMode().

*/
type QProcess__InputChannelMode = int

// QProcess manages the input of the running process. This is the default input channel mode of QProcess.
const QProcess__ManagedInputChannel QProcess__InputChannelMode = 0

// QProcess forwards the input of the main process onto the running process. The child process reads its standard input from the same source as the main process. Note that the main process must not try to read its standard input while the child process is running.
const QProcess__ForwardedInputChannel QProcess__InputChannelMode = 1

func (this *QProcess) InputChannelModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QProcess_InputChannelModeItemName(val int) string {
	var nilthis *QProcess
	return nilthis.InputChannelModeItemName(val)
}

/*
This enum describes the different exit statuses of QProcess.



See also exitStatus().

*/
type QProcess__ExitStatus = int

// The process exited normally.
const QProcess__NormalExit QProcess__ExitStatus = 0

// The process crashed.
const QProcess__CrashExit QProcess__ExitStatus = 1

func (this *QProcess) ExitStatusItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QProcess_ExitStatusItemName(val int) string {
	var nilthis *QProcess
	return nilthis.ExitStatusItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10503() {
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
