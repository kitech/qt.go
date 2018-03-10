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

// void setProcessState(enum QProcess::ProcessState)
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

// /usr/include/qt/QtCore/qprocess.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QProcess) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qprocess.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProcess(QObject *)

/*
Constructs a QProcess object with the given parent.
*/
func NewQProcess(parent QObject_ITF /*777 QObject **/) *QProcess {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcessC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProcessFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProcess")
	return gothis
}

// /usr/include/qt/QtCore/qprocess.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProcess(QObject *)

/*
Constructs a QProcess object with the given parent.
*/
func NewQProcess__() *QProcess {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcessC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQProcessFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QProcess")
	return gothis
}

// /usr/include/qt/QtCore/qprocess.h:159
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QProcess()

/*

 */
func DeleteQProcess(this *QProcess) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcessD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qprocess.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(const QString &, const QStringList &, QIODevice::OpenMode)

/*
Starts the given program in a new process, passing the command line arguments in arguments.

The QProcess object will immediately enter the Starting state. If the process starts successfully, QProcess will emit started(); otherwise, errorOccurred() will be emitted.

Note: Processes are started asynchronously, which means the started() and errorOccurred() signals may be delayed. Call waitForStarted() to make sure the process has started (or has failed to start) and those signals have been emitted.

Note: No further splitting of the arguments is performed.

Windows: The arguments are quoted and joined into a command line that is compatible with the CommandLineToArgvW() Windows function. For programs that have different command line quoting requirements, you need to use setNativeArguments(). One notable program that does not follow the CommandLineToArgvW() rules is cmd.exe and, by consequence, all batch scripts.

The OpenMode is set to mode.

If the QProcess object is already running a process, a warning may be printed at the console, and the existing process will continue running unaffected.

See also processId(), started(), waitForStarted(), and setNativeArguments().
*/
func (this *QProcess) Start(program string, arguments QStringList_ITF, mode int) {
	var tmpArg0 = NewQString_5(program)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if arguments != nil && arguments.QStringList_PTR() != nil {
		convArg1 = arguments.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5startERK7QStringRK11QStringList6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(const QString &, const QStringList &, QIODevice::OpenMode)

/*
Starts the given program in a new process, passing the command line arguments in arguments.

The QProcess object will immediately enter the Starting state. If the process starts successfully, QProcess will emit started(); otherwise, errorOccurred() will be emitted.

Note: Processes are started asynchronously, which means the started() and errorOccurred() signals may be delayed. Call waitForStarted() to make sure the process has started (or has failed to start) and those signals have been emitted.

Note: No further splitting of the arguments is performed.

Windows: The arguments are quoted and joined into a command line that is compatible with the CommandLineToArgvW() Windows function. For programs that have different command line quoting requirements, you need to use setNativeArguments(). One notable program that does not follow the CommandLineToArgvW() rules is cmd.exe and, by consequence, all batch scripts.

The OpenMode is set to mode.

If the QProcess object is already running a process, a warning may be printed at the console, and the existing process will continue running unaffected.

See also processId(), started(), waitForStarted(), and setNativeArguments().
*/
func (this *QProcess) Start__(program string, arguments QStringList_ITF) {
	var tmpArg0 = NewQString_5(program)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if arguments != nil && arguments.QStringList_PTR() != nil {
		convArg1 = arguments.QStringList_PTR().GetCthis()
	}
	// arg: 2, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5startERK7QStringRK11QStringList6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:163
// index:1
// Public Visibility=Default Availability=Available
// [-2] void start(const QString &, QIODevice::OpenMode)

/*
Starts the given program in a new process, passing the command line arguments in arguments.

The QProcess object will immediately enter the Starting state. If the process starts successfully, QProcess will emit started(); otherwise, errorOccurred() will be emitted.

Note: Processes are started asynchronously, which means the started() and errorOccurred() signals may be delayed. Call waitForStarted() to make sure the process has started (or has failed to start) and those signals have been emitted.

Note: No further splitting of the arguments is performed.

Windows: The arguments are quoted and joined into a command line that is compatible with the CommandLineToArgvW() Windows function. For programs that have different command line quoting requirements, you need to use setNativeArguments(). One notable program that does not follow the CommandLineToArgvW() rules is cmd.exe and, by consequence, all batch scripts.

The OpenMode is set to mode.

If the QProcess object is already running a process, a warning may be printed at the console, and the existing process will continue running unaffected.

See also processId(), started(), waitForStarted(), and setNativeArguments().
*/
func (this *QProcess) Start_1(command string, mode int) {
	var tmpArg0 = NewQString_5(command)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5startERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:163
// index:1
// Public Visibility=Default Availability=Available
// [-2] void start(const QString &, QIODevice::OpenMode)

/*
Starts the given program in a new process, passing the command line arguments in arguments.

The QProcess object will immediately enter the Starting state. If the process starts successfully, QProcess will emit started(); otherwise, errorOccurred() will be emitted.

Note: Processes are started asynchronously, which means the started() and errorOccurred() signals may be delayed. Call waitForStarted() to make sure the process has started (or has failed to start) and those signals have been emitted.

Note: No further splitting of the arguments is performed.

Windows: The arguments are quoted and joined into a command line that is compatible with the CommandLineToArgvW() Windows function. For programs that have different command line quoting requirements, you need to use setNativeArguments(). One notable program that does not follow the CommandLineToArgvW() rules is cmd.exe and, by consequence, all batch scripts.

The OpenMode is set to mode.

If the QProcess object is already running a process, a warning may be printed at the console, and the existing process will continue running unaffected.

See also processId(), started(), waitForStarted(), and setNativeArguments().
*/
func (this *QProcess) Start_1_(command string) {
	var tmpArg0 = NewQString_5(command)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5startERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:165
// index:2
// Public Visibility=Default Availability=Available
// [-2] void start(QIODevice::OpenMode)

/*
Starts the given program in a new process, passing the command line arguments in arguments.

The QProcess object will immediately enter the Starting state. If the process starts successfully, QProcess will emit started(); otherwise, errorOccurred() will be emitted.

Note: Processes are started asynchronously, which means the started() and errorOccurred() signals may be delayed. Call waitForStarted() to make sure the process has started (or has failed to start) and those signals have been emitted.

Note: No further splitting of the arguments is performed.

Windows: The arguments are quoted and joined into a command line that is compatible with the CommandLineToArgvW() Windows function. For programs that have different command line quoting requirements, you need to use setNativeArguments(). One notable program that does not follow the CommandLineToArgvW() rules is cmd.exe and, by consequence, all batch scripts.

The OpenMode is set to mode.

If the QProcess object is already running a process, a warning may be printed at the console, and the existing process will continue running unaffected.

See also processId(), started(), waitForStarted(), and setNativeArguments().
*/
func (this *QProcess) Start_2(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5startE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:165
// index:2
// Public Visibility=Default Availability=Available
// [-2] void start(QIODevice::OpenMode)

/*
Starts the given program in a new process, passing the command line arguments in arguments.

The QProcess object will immediately enter the Starting state. If the process starts successfully, QProcess will emit started(); otherwise, errorOccurred() will be emitted.

Note: Processes are started asynchronously, which means the started() and errorOccurred() signals may be delayed. Call waitForStarted() to make sure the process has started (or has failed to start) and those signals have been emitted.

Note: No further splitting of the arguments is performed.

Windows: The arguments are quoted and joined into a command line that is compatible with the CommandLineToArgvW() Windows function. For programs that have different command line quoting requirements, you need to use setNativeArguments(). One notable program that does not follow the CommandLineToArgvW() rules is cmd.exe and, by consequence, all batch scripts.

The OpenMode is set to mode.

If the QProcess object is already running a process, a warning may be printed at the console, and the existing process will continue running unaffected.

See also processId(), started(), waitForStarted(), and setNativeArguments().
*/
func (this *QProcess) Start_2_() {
	// arg: 0, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5startE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool startDetached(qint64 *)

/*
Starts the program set by setProgram() with arguments set by setArguments() in a new process, and detaches from it. Returns true on success; otherwise returns false. If the calling process exits, the detached process will continue to run unaffected.

Unix: The started process will run in its own session and act like a daemon.

The process will be started in the directory set by setWorkingDirectory(). If workingDirectory() is empty, the working directory is inherited from the calling process.

Note: On QNX, this may cause all application threads to temporarily freeze.

If the function is successful then *pid is set to the process identifier of the started process. Note that the child process may exit and the PID may become invalid without notice. Furthermore, after the child process exits, the same PID may be recycled and used by a completely different process. User code should be careful when using this variable, especially if one intends to forcibly terminate the process by operating system means.

Only the following property setters are supported by startDetached():


setArguments()
setCreateProcessArgumentsModifier()
setNativeArguments()
setProcessEnvironment()
setProgram()
setStandardErrorFile()
setStandardInputFile()
setStandardOutputFile()
setWorkingDirectory()


All other properties of the QProcess object are ignored.

This function was introduced in  Qt 5.10.

See also start(), startDetached(const QString &program, const QStringList &arguments, const QString &workingDirectory, qint64 *pid), and startDetached(const QString &command).
*/
func (this *QProcess) StartDetached(pid unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess13startDetachedEPx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pid)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool startDetached(qint64 *)

/*
Starts the program set by setProgram() with arguments set by setArguments() in a new process, and detaches from it. Returns true on success; otherwise returns false. If the calling process exits, the detached process will continue to run unaffected.

Unix: The started process will run in its own session and act like a daemon.

The process will be started in the directory set by setWorkingDirectory(). If workingDirectory() is empty, the working directory is inherited from the calling process.

Note: On QNX, this may cause all application threads to temporarily freeze.

If the function is successful then *pid is set to the process identifier of the started process. Note that the child process may exit and the PID may become invalid without notice. Furthermore, after the child process exits, the same PID may be recycled and used by a completely different process. User code should be careful when using this variable, especially if one intends to forcibly terminate the process by operating system means.

Only the following property setters are supported by startDetached():


setArguments()
setCreateProcessArgumentsModifier()
setNativeArguments()
setProcessEnvironment()
setProgram()
setStandardErrorFile()
setStandardInputFile()
setStandardOutputFile()
setWorkingDirectory()


All other properties of the QProcess object are ignored.

This function was introduced in  Qt 5.10.

See also start(), startDetached(const QString &program, const QStringList &arguments, const QString &workingDirectory, qint64 *pid), and startDetached(const QString &command).
*/
func (this *QProcess) StartDetached__() bool {
	// arg: 0, qint64 *=Pointer, qint64=Typedef, long long
	pid := unsafe.Pointer(nil)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess13startDetachedEPx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pid)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:251
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool startDetached(const QString &, const QStringList &, const QString &, qint64 *)

/*
Starts the program set by setProgram() with arguments set by setArguments() in a new process, and detaches from it. Returns true on success; otherwise returns false. If the calling process exits, the detached process will continue to run unaffected.

Unix: The started process will run in its own session and act like a daemon.

The process will be started in the directory set by setWorkingDirectory(). If workingDirectory() is empty, the working directory is inherited from the calling process.

Note: On QNX, this may cause all application threads to temporarily freeze.

If the function is successful then *pid is set to the process identifier of the started process. Note that the child process may exit and the PID may become invalid without notice. Furthermore, after the child process exits, the same PID may be recycled and used by a completely different process. User code should be careful when using this variable, especially if one intends to forcibly terminate the process by operating system means.

Only the following property setters are supported by startDetached():


setArguments()
setCreateProcessArgumentsModifier()
setNativeArguments()
setProcessEnvironment()
setProgram()
setStandardErrorFile()
setStandardInputFile()
setStandardOutputFile()
setWorkingDirectory()


All other properties of the QProcess object are ignored.

This function was introduced in  Qt 5.10.

See also start(), startDetached(const QString &program, const QStringList &arguments, const QString &workingDirectory, qint64 *pid), and startDetached(const QString &command).
*/
func (this *QProcess) StartDetached_1(program string, arguments QStringList_ITF, workingDirectory string, pid unsafe.Pointer /*666*/) bool {
	var tmpArg0 = NewQString_5(program)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if arguments != nil && arguments.QStringList_PTR() != nil {
		convArg1 = arguments.QStringList_PTR().GetCthis()
	}
	var tmpArg2 = NewQString_5(workingDirectory)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, pid)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QProcess_StartDetached_1(program string, arguments QStringList_ITF, workingDirectory string, pid unsafe.Pointer /*666*/) bool {
	var nilthis *QProcess
	rv := nilthis.StartDetached_1(program, arguments, workingDirectory, pid)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:251
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool startDetached(const QString &, const QStringList &, const QString &, qint64 *)

/*
Starts the program set by setProgram() with arguments set by setArguments() in a new process, and detaches from it. Returns true on success; otherwise returns false. If the calling process exits, the detached process will continue to run unaffected.

Unix: The started process will run in its own session and act like a daemon.

The process will be started in the directory set by setWorkingDirectory(). If workingDirectory() is empty, the working directory is inherited from the calling process.

Note: On QNX, this may cause all application threads to temporarily freeze.

If the function is successful then *pid is set to the process identifier of the started process. Note that the child process may exit and the PID may become invalid without notice. Furthermore, after the child process exits, the same PID may be recycled and used by a completely different process. User code should be careful when using this variable, especially if one intends to forcibly terminate the process by operating system means.

Only the following property setters are supported by startDetached():


setArguments()
setCreateProcessArgumentsModifier()
setNativeArguments()
setProcessEnvironment()
setProgram()
setStandardErrorFile()
setStandardInputFile()
setStandardOutputFile()
setWorkingDirectory()


All other properties of the QProcess object are ignored.

This function was introduced in  Qt 5.10.

See also start(), startDetached(const QString &program, const QStringList &arguments, const QString &workingDirectory, qint64 *pid), and startDetached(const QString &command).
*/
func (this *QProcess) StartDetached_1_(program string, arguments QStringList_ITF, workingDirectory string) bool {
	var tmpArg0 = NewQString_5(program)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if arguments != nil && arguments.QStringList_PTR() != nil {
		convArg1 = arguments.QStringList_PTR().GetCthis()
	}
	var tmpArg2 = NewQString_5(workingDirectory)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, qint64 *=Pointer, qint64=Typedef, long long
	pid := unsafe.Pointer(nil)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, pid)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:258
// index:2
// Public static Visibility=Default Availability=Available
// [1] bool startDetached(const QString &, const QStringList &)

/*
Starts the program set by setProgram() with arguments set by setArguments() in a new process, and detaches from it. Returns true on success; otherwise returns false. If the calling process exits, the detached process will continue to run unaffected.

Unix: The started process will run in its own session and act like a daemon.

The process will be started in the directory set by setWorkingDirectory(). If workingDirectory() is empty, the working directory is inherited from the calling process.

Note: On QNX, this may cause all application threads to temporarily freeze.

If the function is successful then *pid is set to the process identifier of the started process. Note that the child process may exit and the PID may become invalid without notice. Furthermore, after the child process exits, the same PID may be recycled and used by a completely different process. User code should be careful when using this variable, especially if one intends to forcibly terminate the process by operating system means.

Only the following property setters are supported by startDetached():


setArguments()
setCreateProcessArgumentsModifier()
setNativeArguments()
setProcessEnvironment()
setProgram()
setStandardErrorFile()
setStandardInputFile()
setStandardOutputFile()
setWorkingDirectory()


All other properties of the QProcess object are ignored.

This function was introduced in  Qt 5.10.

See also start(), startDetached(const QString &program, const QStringList &arguments, const QString &workingDirectory, qint64 *pid), and startDetached(const QString &command).
*/
func (this *QProcess) StartDetached_2(program string, arguments QStringList_ITF) bool {
	var tmpArg0 = NewQString_5(program)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if arguments != nil && arguments.QStringList_PTR() != nil {
		convArg1 = arguments.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QProcess_StartDetached_2(program string, arguments QStringList_ITF) bool {
	var nilthis *QProcess
	rv := nilthis.StartDetached_2(program, arguments)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:260
// index:3
// Public static Visibility=Default Availability=Available
// [1] bool startDetached(const QString &)

/*
Starts the program set by setProgram() with arguments set by setArguments() in a new process, and detaches from it. Returns true on success; otherwise returns false. If the calling process exits, the detached process will continue to run unaffected.

Unix: The started process will run in its own session and act like a daemon.

The process will be started in the directory set by setWorkingDirectory(). If workingDirectory() is empty, the working directory is inherited from the calling process.

Note: On QNX, this may cause all application threads to temporarily freeze.

If the function is successful then *pid is set to the process identifier of the started process. Note that the child process may exit and the PID may become invalid without notice. Furthermore, after the child process exits, the same PID may be recycled and used by a completely different process. User code should be careful when using this variable, especially if one intends to forcibly terminate the process by operating system means.

Only the following property setters are supported by startDetached():


setArguments()
setCreateProcessArgumentsModifier()
setNativeArguments()
setProcessEnvironment()
setProgram()
setStandardErrorFile()
setStandardInputFile()
setStandardOutputFile()
setWorkingDirectory()


All other properties of the QProcess object are ignored.

This function was introduced in  Qt 5.10.

See also start(), startDetached(const QString &program, const QStringList &arguments, const QString &workingDirectory, qint64 *pid), and startDetached(const QString &command).
*/
func (this *QProcess) StartDetached_3(command string) bool {
	var tmpArg0 = NewQString_5(command)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QProcess_StartDetached_3(command string) bool {
	var nilthis *QProcess
	rv := nilthis.StartDetached_3(command)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)

/*
Reimplemented from QIODevice::open().

Starts the program set by setProgram() with arguments set by setArguments(). The OpenMode is set to mode.

This method is an alias for start(), and exists only to fully implement the interface defined by QIODevice.

See also start(), setProgram(), and setArguments().
*/
func (this *QProcess) Open(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)

/*
Reimplemented from QIODevice::open().

Starts the program set by setProgram() with arguments set by setArguments(). The OpenMode is set to mode.

This method is an alias for start(), and exists only to fully implement the interface defined by QIODevice.

See also start(), setProgram(), and setArguments().
*/
func (this *QProcess) Open__() bool {
	// arg: 0, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QString program() const

/*
Returns the program the process was last started with.

This function was introduced in  Qt 5.0.

See also setProgram() and start().
*/
func (this *QProcess) Program() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess7programEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qprocess.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProgram(const QString &)

/*
Set the program to use when starting the process. This function must be called before start().

This function was introduced in  Qt 5.1.

See also start(), setArguments(), and program().
*/
func (this *QProcess) SetProgram(program string) {
	var tmpArg0 = NewQString_5(program)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess10setProgramERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:172
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList arguments() const

/*
Returns the command line arguments the process was last started with.

This function was introduced in  Qt 5.0.

See also setArguments() and start().
*/
func (this *QProcess) Arguments() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess9argumentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setArguments(const QStringList &)

/*
Set the arguments to pass to the called program when starting the process. This function must be called before start().

This function was introduced in  Qt 5.1.

See also start(), setProgram(), and arguments().
*/
func (this *QProcess) SetArguments(arguments QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if arguments != nil && arguments.QStringList_PTR() != nil {
		convArg0 = arguments.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess12setArgumentsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:175
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessChannelMode readChannelMode() const

/*

 */
func (this *QProcess) ReadChannelMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess15readChannelModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadChannelMode(enum QProcess::ProcessChannelMode)

/*

 */
func (this *QProcess) SetReadChannelMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess18setReadChannelModeENS_18ProcessChannelModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:177
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessChannelMode processChannelMode() const

/*
Returns the channel mode of the QProcess standard output and standard error channels.

This function was introduced in  Qt 4.2.

See also setProcessChannelMode(), ProcessChannelMode, and setReadChannel().
*/
func (this *QProcess) ProcessChannelMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess18processChannelModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProcessChannelMode(enum QProcess::ProcessChannelMode)

/*
Sets the channel mode of the QProcess standard output and standard error channels to the mode specified. This mode will be used the next time start() is called. For example:


  QProcess builder;
  builder.setProcessChannelMode(QProcess::MergedChannels);
  builder.start("make", QStringList() << "-j2");

  if (!builder.waitForFinished())
      qDebug() << "Make failed:" << builder.errorString();
  else
      qDebug() << "Make output:" << builder.readAll();



This function was introduced in  Qt 4.2.

See also processChannelMode(), ProcessChannelMode, and setReadChannel().
*/
func (this *QProcess) SetProcessChannelMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess21setProcessChannelModeENS_18ProcessChannelModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:179
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::InputChannelMode inputChannelMode() const

/*
Returns the channel mode of the QProcess standard input channel.

This function was introduced in  Qt 5.2.

See also setInputChannelMode() and InputChannelMode.
*/
func (this *QProcess) InputChannelMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess16inputChannelModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputChannelMode(enum QProcess::InputChannelMode)

/*
Sets the channel mode of the QProcess standard input channel to the mode specified. This mode will be used the next time start() is called.

This function was introduced in  Qt 5.2.

See also inputChannelMode() and InputChannelMode.
*/
func (this *QProcess) SetInputChannelMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess19setInputChannelModeENS_16InputChannelModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:182
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessChannel readChannel() const

/*
Returns the current read channel of the QProcess.

See also setReadChannel().
*/
func (this *QProcess) ReadChannel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess11readChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadChannel(enum QProcess::ProcessChannel)

/*
Sets the current read channel of the QProcess to the given channel. The current input channel is used by the functions read(), readAll(), readLine(), and getChar(). It also determines which channel triggers QProcess to emit readyRead().

See also readChannel().
*/
func (this *QProcess) SetReadChannel(channel int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess14setReadChannelENS_14ProcessChannelE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeReadChannel(enum QProcess::ProcessChannel)

/*
Closes the read channel channel. After calling this function, QProcess will no longer receive data on the channel. Any data that has already been received is still available for reading.

Call this function to save memory, if you are not interested in the output of the process.

See also closeWriteChannel() and setReadChannel().
*/
func (this *QProcess) CloseReadChannel(channel int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess16closeReadChannelENS_14ProcessChannelE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeWriteChannel()

/*
Schedules the write channel of QProcess to be closed. The channel will close once all data has been written to the process. After calling this function, any attempts to write to the process will fail.

Closing the write channel is necessary for programs that read input data until the channel has been closed. For example, the program "more" is used to display text data in a console on both Unix and Windows. But it will not display the text data until QProcess's write channel has been closed. Example:


  QProcess more;
  more.start("more");
  more.write("Text to display");
  more.closeWriteChannel();
  // QProcess will emit readyRead() once "more" starts printing



The write channel is implicitly opened when start() is called.

See also closeReadChannel().
*/
func (this *QProcess) CloseWriteChannel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess17closeWriteChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardInputFile(const QString &)

/*
Redirects the process' standard input to the file indicated by fileName. When an input redirection is in place, the QProcess object will be in read-only mode (calling write() will result in error).

To make the process read EOF right away, pass nullDevice() here. This is cleaner than using closeWriteChannel() before writing any data, because it can be set up prior to starting the process.

If the file fileName does not exist at the moment start() is called or is not readable, starting the process will fail.

Calling setStandardInputFile() after the process has started has no effect.

This function was introduced in  Qt 4.2.

See also setStandardOutputFile(), setStandardErrorFile(), and setStandardOutputProcess().
*/
func (this *QProcess) SetStandardInputFile(fileName string) {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess20setStandardInputFileERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardOutputFile(const QString &, QIODevice::OpenMode)

/*
Redirects the process' standard output to the file fileName. When the redirection is in place, the standard output read channel is closed: reading from it using read() will always fail, as will readAllStandardOutput().

To discard all standard output from the process, pass nullDevice() here. This is more efficient than simply never reading the standard output, as no QProcess buffers are filled.

If the file fileName doesn't exist at the moment start() is called, it will be created. If it cannot be created, the starting will fail.

If the file exists and mode is QIODevice::Truncate, the file will be truncated. Otherwise (if mode is QIODevice::Append), the file will be appended to.

Calling setStandardOutputFile() after the process has started has no effect.

This function was introduced in  Qt 4.2.

See also setStandardInputFile(), setStandardErrorFile(), and setStandardOutputProcess().
*/
func (this *QProcess) SetStandardOutputFile(fileName string, mode int) {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess21setStandardOutputFileERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardOutputFile(const QString &, QIODevice::OpenMode)

/*
Redirects the process' standard output to the file fileName. When the redirection is in place, the standard output read channel is closed: reading from it using read() will always fail, as will readAllStandardOutput().

To discard all standard output from the process, pass nullDevice() here. This is more efficient than simply never reading the standard output, as no QProcess buffers are filled.

If the file fileName doesn't exist at the moment start() is called, it will be created. If it cannot be created, the starting will fail.

If the file exists and mode is QIODevice::Truncate, the file will be truncated. Otherwise (if mode is QIODevice::Append), the file will be appended to.

Calling setStandardOutputFile() after the process has started has no effect.

This function was introduced in  Qt 4.2.

See also setStandardInputFile(), setStandardErrorFile(), and setStandardOutputProcess().
*/
func (this *QProcess) SetStandardOutputFile__(fileName string) {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess21setStandardOutputFileERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardErrorFile(const QString &, QIODevice::OpenMode)

/*
Redirects the process' standard error to the file fileName. When the redirection is in place, the standard error read channel is closed: reading from it using read() will always fail, as will readAllStandardError(). The file will be appended to if mode is Append, otherwise, it will be truncated.

See setStandardOutputFile() for more information on how the file is opened.

Note: if setProcessChannelMode() was called with an argument of QProcess::MergedChannels, this function has no effect.

This function was introduced in  Qt 4.2.

See also setStandardInputFile(), setStandardOutputFile(), and setStandardOutputProcess().
*/
func (this *QProcess) SetStandardErrorFile(fileName string, mode int) {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess20setStandardErrorFileERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardErrorFile(const QString &, QIODevice::OpenMode)

/*
Redirects the process' standard error to the file fileName. When the redirection is in place, the standard error read channel is closed: reading from it using read() will always fail, as will readAllStandardError(). The file will be appended to if mode is Append, otherwise, it will be truncated.

See setStandardOutputFile() for more information on how the file is opened.

Note: if setProcessChannelMode() was called with an argument of QProcess::MergedChannels, this function has no effect.

This function was introduced in  Qt 4.2.

See also setStandardInputFile(), setStandardOutputFile(), and setStandardOutputProcess().
*/
func (this *QProcess) SetStandardErrorFile__(fileName string) {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QIODevice::OpenMode=Typedef, QIODevice::OpenMode=Typedef, QFlags<QIODevice::OpenModeFlag>
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess20setStandardErrorFileERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardOutputProcess(QProcess *)

/*
Pipes the standard output stream of this process to the destination process' standard input.

The following shell command:


  command1 | command2



Can be accomplished with QProcess with the following code:


  QProcess process1;
  QProcess process2;

  process1.setStandardOutputProcess(&process2);

  process1.start("command1");
  process2.start("command2");



This function was introduced in  Qt 4.2.
*/
func (this *QProcess) SetStandardOutputProcess(destination QProcess_ITF /*777 QProcess **/) {
	var convArg0 unsafe.Pointer
	if destination != nil && destination.QProcess_PTR() != nil {
		convArg0 = destination.QProcess_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess24setStandardOutputProcessEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:214
// index:0
// Public Visibility=Default Availability=Available
// [8] QString workingDirectory() const

/*
If QProcess has been assigned a working directory, this function returns the working directory that the QProcess will enter before the program has started. Otherwise, (i.e., no directory has been assigned,) an empty string is returned, and QProcess will use the application's current working directory instead.

See also setWorkingDirectory().
*/
func (this *QProcess) WorkingDirectory() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess16workingDirectoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qprocess.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorkingDirectory(const QString &)

/*
Sets the working directory to dir. QProcess will start the process in this directory. The default behavior is to start the process in the working directory of the calling process.

Note: On QNX, this may cause all application threads to temporarily freeze.

See also workingDirectory() and start().
*/
func (this *QProcess) SetWorkingDirectory(dir string) {
	var tmpArg0 = NewQString_5(dir)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess19setWorkingDirectoryERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnvironment(const QStringList &)

/*

 */
func (this *QProcess) SetEnvironment(environment QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if environment != nil && environment.QStringList_PTR() != nil {
		convArg0 = environment.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess14setEnvironmentERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:218
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList environment() const

/*

 */
func (this *QProcess) Environment() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess11environmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProcessEnvironment(const QProcessEnvironment &)

/*
Sets the environment that QProcess will pass to the child process.

For example, the following code adds the environment variable TMPDIR:


  QProcess process;
  QProcessEnvironment env = QProcessEnvironment::systemEnvironment();
  env.insert("TMPDIR", "C:\\MyApp\\temp"); // Add an environment variable
  process.setProcessEnvironment(env);
  process.start("myapp");



Note how, on Windows, environment variable names are case-insensitive.

This function was introduced in  Qt 4.6.

See also processEnvironment(), QProcessEnvironment::systemEnvironment(), and setEnvironment().
*/
func (this *QProcess) SetProcessEnvironment(environment QProcessEnvironment_ITF) {
	var convArg0 unsafe.Pointer
	if environment != nil && environment.QProcessEnvironment_PTR() != nil {
		convArg0 = environment.QProcessEnvironment_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess21setProcessEnvironmentERK19QProcessEnvironment", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:220
// index:0
// Public Visibility=Default Availability=Available
// [8] QProcessEnvironment processEnvironment() const

/*
Returns the environment that QProcess will pass to its child process, or an empty object if no environment has been set using setEnvironment() or setProcessEnvironment(). If no environment has been set, the environment of the calling process will be used.

This function was introduced in  Qt 4.6.

See also setProcessEnvironment(), setEnvironment(), and QProcessEnvironment::isEmpty().
*/
func (this *QProcess) ProcessEnvironment() *QProcessEnvironment /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess18processEnvironmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQProcessEnvironmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQProcessEnvironment)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:222
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessError error() const

/*
Returns the type of error that occurred last.

See also state().
*/
func (this *QProcess) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:275
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QProcess::ProcessError)

/*
Returns the type of error that occurred last.

See also state().
*/
func (this *QProcess) Error_1(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5errorENS_12ProcessErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:223
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessState state() const

/*
Returns the current state of the process.

See also stateChanged() and error().
*/
func (this *QProcess) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:226
// index:0
// Public Visibility=Default Availability=Available
// [8] Q_PID pid() const

/*

 */
func (this *QProcess) Pid() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess3pidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
	// long long // 222
}

// /usr/include/qt/QtCore/qprocess.h:227
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 processId() const

/*
Returns the native process identifier for the running process, if available. If no process is currently running, 0 is returned.

This function was introduced in  Qt 5.3.
*/
func (this *QProcess) ProcessId() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess9processIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
	// long long // 222
}

// /usr/include/qt/QtCore/qprocess.h:229
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForStarted(int)

/*
Blocks until the process has started and the started() signal has been emitted, or until msecs milliseconds have passed.

Returns true if the process was started successfully; otherwise returns false (if the operation timed out or if an error occurred).

This function can operate without an event loop. It is useful when writing non-GUI applications and when performing I/O operations in a non-GUI thread.

Warning: Calling this function from the main (GUI) thread might cause your user interface to freeze.

If msecs is -1, this function will not time out.

Note: On some UNIX operating systems, this function may return true but the process may later report a QProcess::FailedToStart error.

See also started(), waitForReadyRead(), waitForBytesWritten(), and waitForFinished().
*/
func (this *QProcess) WaitForStarted(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess14waitForStartedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:229
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForStarted(int)

/*
Blocks until the process has started and the started() signal has been emitted, or until msecs milliseconds have passed.

Returns true if the process was started successfully; otherwise returns false (if the operation timed out or if an error occurred).

This function can operate without an event loop. It is useful when writing non-GUI applications and when performing I/O operations in a non-GUI thread.

Warning: Calling this function from the main (GUI) thread might cause your user interface to freeze.

If msecs is -1, this function will not time out.

Note: On some UNIX operating systems, this function may return true but the process may later report a QProcess::FailedToStart error.

See also started(), waitForReadyRead(), waitForBytesWritten(), and waitForFinished().
*/
func (this *QProcess) WaitForStarted__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess14waitForStartedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:230
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)

/*
Reimplemented from QIODevice::waitForReadyRead().
*/
func (this *QProcess) WaitForReadyRead(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:230
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)

/*
Reimplemented from QIODevice::waitForReadyRead().
*/
func (this *QProcess) WaitForReadyRead__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:231
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)

/*
Reimplemented from QIODevice::waitForBytesWritten().
*/
func (this *QProcess) WaitForBytesWritten(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:231
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)

/*
Reimplemented from QIODevice::waitForBytesWritten().
*/
func (this *QProcess) WaitForBytesWritten__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:232
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForFinished(int)

/*
Blocks until the process has finished and the finished() signal has been emitted, or until msecs milliseconds have passed.

Returns true if the process finished; otherwise returns false (if the operation timed out, if an error occurred, or if this QProcess is already finished).

This function can operate without an event loop. It is useful when writing non-GUI applications and when performing I/O operations in a non-GUI thread.

Warning: Calling this function from the main (GUI) thread might cause your user interface to freeze.

If msecs is -1, this function will not time out.

See also finished(), waitForStarted(), waitForReadyRead(), and waitForBytesWritten().
*/
func (this *QProcess) WaitForFinished(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess15waitForFinishedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:232
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForFinished(int)

/*
Blocks until the process has finished and the finished() signal has been emitted, or until msecs milliseconds have passed.

Returns true if the process finished; otherwise returns false (if the operation timed out, if an error occurred, or if this QProcess is already finished).

This function can operate without an event loop. It is useful when writing non-GUI applications and when performing I/O operations in a non-GUI thread.

Warning: Calling this function from the main (GUI) thread might cause your user interface to freeze.

If msecs is -1, this function will not time out.

See also finished(), waitForStarted(), waitForReadyRead(), and waitForBytesWritten().
*/
func (this *QProcess) WaitForFinished__() bool {
	// arg: 0, int=Int, =Invalid,
	msecs := int(30000)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess15waitForFinishedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:234
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray readAllStandardOutput()

/*
Regardless of the current read channel, this function returns all data available from the standard output of the process as a QByteArray.

See also readyReadStandardOutput(), readAllStandardError(), readChannel(), and setReadChannel().
*/
func (this *QProcess) ReadAllStandardOutput() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess21readAllStandardOutputEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:235
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray readAllStandardError()

/*
Regardless of the current read channel, this function returns all data available from the standard error of the process as a QByteArray.

See also readyReadStandardError(), readAllStandardOutput(), readChannel(), and setReadChannel().
*/
func (this *QProcess) ReadAllStandardError() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess20readAllStandardErrorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:237
// index:0
// Public Visibility=Default Availability=Available
// [4] int exitCode() const

/*
Returns the exit code of the last process that finished.

This value is not valid unless exitStatus() returns NormalExit.
*/
func (this *QProcess) ExitCode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess8exitCodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qprocess.h:238
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ExitStatus exitStatus() const

/*
Returns the exit status of the last process that finished.

On Windows, if the process was terminated with TerminateProcess() from another application, this function will still return NormalExit unless the exit code is less than 0.

This function was introduced in  Qt 4.1.
*/
func (this *QProcess) ExitStatus() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess10exitStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:241
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable() const

/*
Reimplemented from QIODevice::bytesAvailable().
*/
func (this *QProcess) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
	// long long // 222
}

// /usr/include/qt/QtCore/qprocess.h:242
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite() const

/*
Reimplemented from QIODevice::bytesToWrite().
*/
func (this *QProcess) BytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess12bytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
	// long long // 222
}

// /usr/include/qt/QtCore/qprocess.h:243
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential() const

/*
Reimplemented from QIODevice::isSequential().
*/
func (this *QProcess) IsSequential() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess12isSequentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:244
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine() const

/*
Reimplemented from QIODevice::canReadLine().

This function operates on the current read channel.

See also readChannel() and setReadChannel().
*/
func (this *QProcess) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:245
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()

/*
Reimplemented from QIODevice::close().

Closes all communication with the process and kills it. After calling this function, QProcess will no longer emit readyRead(), and data can no longer be read or written.
*/
func (this *QProcess) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:246
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd() const

/*
Reimplemented from QIODevice::atEnd().

Returns true if the process is not running, and no more data is available for reading; otherwise returns false.
*/
func (this *QProcess) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:248
// index:0
// Public static Visibility=Default Availability=Available
// [4] int execute(const QString &, const QStringList &)

/*
Starts the program program with the arguments arguments in a new process, waits for it to finish, and then returns the exit code of the process. Any data the new process writes to the console is forwarded to the calling process.

The environment and working directory are inherited from the calling process.

Argument handling is identical to the respective start() overload.

If the process cannot be started, -2 is returned. If the process crashes, -1 is returned. Otherwise, the process' exit code is returned.

See also start().
*/
func (this *QProcess) Execute(program string, arguments QStringList_ITF) int {
	var tmpArg0 = NewQString_5(program)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if arguments != nil && arguments.QStringList_PTR() != nil {
		convArg1 = arguments.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess7executeERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QProcess_Execute(program string, arguments QStringList_ITF) int {
	var nilthis *QProcess
	rv := nilthis.Execute(program, arguments)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:249
// index:1
// Public static Visibility=Default Availability=Available
// [4] int execute(const QString &)

/*
Starts the program program with the arguments arguments in a new process, waits for it to finish, and then returns the exit code of the process. Any data the new process writes to the console is forwarded to the calling process.

The environment and working directory are inherited from the calling process.

Argument handling is identical to the respective start() overload.

If the process cannot be started, -2 is returned. If the process crashes, -1 is returned. Otherwise, the process' exit code is returned.

See also start().
*/
func (this *QProcess) Execute_1(command string) int {
	var tmpArg0 = NewQString_5(command)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess7executeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QProcess_Execute_1(command string) int {
	var nilthis *QProcess
	rv := nilthis.Execute_1(command)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:262
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList systemEnvironment()

/*
Returns the environment of the calling process as a list of key=value pairs. Example:


  QStringList environment = QProcess::systemEnvironment();
  // environment = {"PATH=/usr/bin:/usr/local/bin",
  //                "USER=greg", "HOME=/home/greg"}



This function does not cache the system environment. Therefore, it's possible to obtain an updated version of the environment if low-level C library functions like setenv or putenv have been called.

However, note that repeated calls to this function will recreate the list of environment variables, which is a non-trivial operation.

Note: For new code, it is recommended to use QProcessEnvironment::systemEnvironment()

This function was introduced in  Qt 4.1.

See also QProcessEnvironment::systemEnvironment() and setProcessEnvironment().
*/
func (this *QProcess) SystemEnvironment() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess17systemEnvironmentEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QProcess_SystemEnvironment() *QStringList /*123*/ {
	var nilthis *QProcess
	rv := nilthis.SystemEnvironment()
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:264
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString nullDevice()

/*
The null device of the operating system.

The returned file path uses native directory separators.

This function was introduced in  Qt 5.2.

See also QProcess::setStandardInputFile(), QProcess::setStandardOutputFile(), and QProcess::setStandardErrorFile().
*/
func (this *QProcess) NullDevice() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess10nullDeviceEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QProcess_NullDevice() string {
	var nilthis *QProcess
	rv := nilthis.NullDevice()
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void terminate()

/*
Attempts to terminate the process.

The process may not exit as a result of calling this function (it is given the chance to prompt the user for any unsaved files, etc).

On Windows, terminate() posts a WM_CLOSE message to all top-level windows of the process and then to the main thread of the process itself. On Unix and macOS the SIGTERM signal is sent.

Console applications on Windows that do not run an event loop, or whose event loop does not handle the WM_CLOSE message, can only be terminated by calling kill().

See also kill().
*/
func (this *QProcess) Terminate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess9terminateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void kill()

/*
Kills the current process, causing it to exit immediately.

On Windows, kill() uses TerminateProcess, and on Unix and macOS, the SIGKILL signal is sent to the process.

See also terminate().
*/
func (this *QProcess) Kill() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess4killEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:272
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished(int)

/*
This signal is emitted when the process finishes. exitCode is the exit code of the process (only valid for normal exits), and exitStatus is the exit status. After the process has finished, the buffers in QProcess are still intact. You can still read any data that the process may have written before it finished.

Note: Signal finished is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(process, QOverload<int, QProcess::ExitStatus>::of(&QProcess::finished),
      [=](int exitCode, QProcess::ExitStatus exitStatus){ \/* ... *\/ });



See also exitStatus().
*/
func (this *QProcess) Finished(exitCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess8finishedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), exitCode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:273
// index:1
// Public Visibility=Default Availability=Available
// [-2] void finished(int, QProcess::ExitStatus)

/*
This signal is emitted when the process finishes. exitCode is the exit code of the process (only valid for normal exits), and exitStatus is the exit status. After the process has finished, the buffers in QProcess are still intact. You can still read any data that the process may have written before it finished.

Note: Signal finished is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(process, QOverload<int, QProcess::ExitStatus>::of(&QProcess::finished),
      [=](int exitCode, QProcess::ExitStatus exitStatus){ \/* ... *\/ });



See also exitStatus().
*/
func (this *QProcess) Finished_1(exitCode int, exitStatus int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess8finishedEiNS_10ExitStatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), exitCode, exitStatus)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:277
// index:0
// Public Visibility=Default Availability=Available
// [-2] void errorOccurred(QProcess::ProcessError)

/*
This signal is emitted when an error occurs with the process. The specified error describes the type of error that occurred.

This function was introduced in  Qt 5.6.
*/
func (this *QProcess) ErrorOccurred(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess13errorOccurredENS_12ProcessErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:284
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setProcessState(enum QProcess::ProcessState)

/*
Sets the current state of the QProcess to the state specified.

See also state().
*/
func (this *QProcess) SetProcessState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess15setProcessStateENS_12ProcessStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:286
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setupChildProcess()

/*
This function is called in the child process context just before the program is executed on Unix or macOS (i.e., after fork(), but before execve()). Reimplement this function to do last minute initialization of the child process. Example:


  class SandboxProcess : public QProcess
  {
      ...
   protected:
       void setupChildProcess();
      ...
  };

  void SandboxProcess::setupChildProcess()
  {
      // Drop all privileges in the child process, and enter
      // a chroot jail.
  #if defined Q_OS_UNIX
      ::setgroups(0, 0);
      ::chroot("/etc/safe");
      ::chdir("/");
      ::setgid(safeGid);
      ::setuid(safeUid);
      ::umask(0);
  #endif
  }



You cannot exit the process (by calling exit(), for instance) from this function. If you need to stop the program before it starts execution, your workaround is to emit finished() and then call exit().

Warning: This function is called by QProcess on Unix and macOS only. On Windows and QNX, it is not called.
*/
func (this *QProcess) SetupChildProcess() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess17setupChildProcessEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:289
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)

/*
Reimplemented from QIODevice::readData().
*/
func (this *QProcess) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
	// long long // 222
}

// /usr/include/qt/QtCore/qprocess.h:290
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)

/*
Reimplemented from QIODevice::writeData().
*/
func (this *QProcess) WriteData(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
	// long long // 222
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

/*
This enum describes the process channels used by the running process. Pass one of these values to setReadChannel() to set the current read channel of QProcess.



See also setReadChannel().

*/
type QProcess__ProcessChannel = int

// The standard output (stdout) of the running process.
const QProcess__StandardOutput QProcess__ProcessChannel = 0

// The standard error (stderr) of the running process.
const QProcess__StandardError QProcess__ProcessChannel = 1

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

/*
This enum describes the different exit statuses of QProcess.



See also exitStatus().

*/
type QProcess__ExitStatus = int

// The process exited normally.
const QProcess__NormalExit QProcess__ExitStatus = 0

// The process crashed.
const QProcess__CrashExit QProcess__ExitStatus = 1

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
