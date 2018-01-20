//  header block begin
// /usr/include/qt/QtCore/qprocess.h
// #include <qprocess.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QProcess struct {
	*QIODevice
}

func (this *QProcess) GetCthis() unsafe.Pointer {
	return this.QIODevice.GetCthis()
}
func NewQProcessFromPointer(cthis unsafe.Pointer) *QProcess {
	bcthis0 := NewQIODeviceFromPointer(cthis)
	return &QProcess{bcthis0}
}

// /usr/include/qt/QtCore/qprocess.h:112
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QProcess) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:158
// index:0
// Public
// void QProcess(class QObject *)
func NewQProcess(parent unsafe.Pointer) *QProcess {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcessC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQProcessFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qprocess.h:159
// index:0
// Public virtual
// void ~QProcess()
func DeleteQProcess(*QProcess) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcessD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:166
// index:0
// Public
// bool startDetached(qint64 *)
func (this *QProcess) StartDetached(pid unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedEPx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pid)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:251
// index:1
// Public static
// bool startDetached(const class QString &, const class QStringList &, const class QString &, qint64 *)
func (this *QProcess) StartDetached_1(program *QString, arguments *QStringList, workingDirectory *QString, pid unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px", ffiqt.FFI_TYPE_POINTER, program, arguments, workingDirectory, pid)
	gopp.ErrPrint(err, rv)
	return rv
}
func QProcess_StartDetached_1(program *QString, arguments *QStringList, workingDirectory *QString, pid unsafe.Pointer) {
	var nilthis *QProcess
	nilthis.StartDetached_1(program, arguments, workingDirectory, pid)
}

// /usr/include/qt/QtCore/qprocess.h:258
// index:2
// Public static
// bool startDetached(const class QString &, const class QStringList &)
func (this *QProcess) StartDetached_2(program *QString, arguments *QStringList) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QStringRK11QStringList", ffiqt.FFI_TYPE_POINTER, program, arguments)
	gopp.ErrPrint(err, rv)
	return rv
}
func QProcess_StartDetached_2(program *QString, arguments *QStringList) {
	var nilthis *QProcess
	nilthis.StartDetached_2(program, arguments)
}

// /usr/include/qt/QtCore/qprocess.h:260
// index:3
// Public static
// bool startDetached(const class QString &)
func (this *QProcess) StartDetached_3(command *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QString", ffiqt.FFI_TYPE_POINTER, command)
	gopp.ErrPrint(err, rv)
	return rv
}
func QProcess_StartDetached_3(command *QString) {
	var nilthis *QProcess
	nilthis.StartDetached_3(command)
}

// /usr/include/qt/QtCore/qprocess.h:169
// index:0
// Public
// QString program()
func (this *QProcess) Program() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess7programEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:170
// index:0
// Public
// void setProgram(const class QString &)
func (this *QProcess) SetProgram(program *QString) {
	var convArg0 = program.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess10setProgramERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:172
// index:0
// Public
// QStringList arguments()
func (this *QProcess) Arguments() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess9argumentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:173
// index:0
// Public
// void setArguments(const class QStringList &)
func (this *QProcess) SetArguments(arguments *QStringList) {
	var convArg0 = arguments.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess12setArgumentsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:175
// index:0
// Public
// QProcess::ProcessChannelMode readChannelMode()
func (this *QProcess) ReadChannelMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess15readChannelModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:176
// index:0
// Public
// void setReadChannelMode(enum QProcess::ProcessChannelMode)
func (this *QProcess) SetReadChannelMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess18setReadChannelModeENS_18ProcessChannelModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:177
// index:0
// Public
// QProcess::ProcessChannelMode processChannelMode()
func (this *QProcess) ProcessChannelMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess18processChannelModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:178
// index:0
// Public
// void setProcessChannelMode(enum QProcess::ProcessChannelMode)
func (this *QProcess) SetProcessChannelMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess21setProcessChannelModeENS_18ProcessChannelModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:179
// index:0
// Public
// QProcess::InputChannelMode inputChannelMode()
func (this *QProcess) InputChannelMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess16inputChannelModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:180
// index:0
// Public
// void setInputChannelMode(enum QProcess::InputChannelMode)
func (this *QProcess) SetInputChannelMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess19setInputChannelModeENS_16InputChannelModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:182
// index:0
// Public
// QProcess::ProcessChannel readChannel()
func (this *QProcess) ReadChannel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess11readChannelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:183
// index:0
// Public
// void setReadChannel(enum QProcess::ProcessChannel)
func (this *QProcess) SetReadChannel(channel int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess14setReadChannelENS_14ProcessChannelE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:185
// index:0
// Public
// void closeReadChannel(enum QProcess::ProcessChannel)
func (this *QProcess) CloseReadChannel(channel int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess16closeReadChannelENS_14ProcessChannelE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:186
// index:0
// Public
// void closeWriteChannel()
func (this *QProcess) CloseWriteChannel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess17closeWriteChannelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:188
// index:0
// Public
// void setStandardInputFile(const class QString &)
func (this *QProcess) SetStandardInputFile(fileName *QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess20setStandardInputFileERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:191
// index:0
// Public
// void setStandardOutputProcess(class QProcess *)
func (this *QProcess) SetStandardOutputProcess(destination unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess24setStandardOutputProcessEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), destination)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:214
// index:0
// Public
// QString workingDirectory()
func (this *QProcess) WorkingDirectory() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess16workingDirectoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:215
// index:0
// Public
// void setWorkingDirectory(const class QString &)
func (this *QProcess) SetWorkingDirectory(dir *QString) {
	var convArg0 = dir.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess19setWorkingDirectoryERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:217
// index:0
// Public
// void setEnvironment(const class QStringList &)
func (this *QProcess) SetEnvironment(environment *QStringList) {
	var convArg0 = environment.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess14setEnvironmentERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:218
// index:0
// Public
// QStringList environment()
func (this *QProcess) Environment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess11environmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:219
// index:0
// Public
// void setProcessEnvironment(const class QProcessEnvironment &)
func (this *QProcess) SetProcessEnvironment(environment *QProcessEnvironment) {
	var convArg0 = environment.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess21setProcessEnvironmentERK19QProcessEnvironment", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:220
// index:0
// Public
// QProcessEnvironment processEnvironment()
func (this *QProcess) ProcessEnvironment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess18processEnvironmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:222
// index:0
// Public
// QProcess::ProcessError error()
func (this *QProcess) Error() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:275
// index:1
// Public
// void error(class QProcess::ProcessError)
func (this *QProcess) Error_1(error int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess5errorENS_12ProcessErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &error)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:223
// index:0
// Public
// QProcess::ProcessState state()
func (this *QProcess) State() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:226
// index:0
// Public
// Q_PID pid()
func (this *QProcess) Pid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess3pidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:227
// index:0
// Public
// qint64 processId()
func (this *QProcess) ProcessId() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess9processIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:229
// index:0
// Public
// bool waitForStarted(int)
func (this *QProcess) WaitForStarted(msecs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess14waitForStartedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:230
// index:0
// Public virtual
// bool waitForReadyRead(int)
func (this *QProcess) WaitForReadyRead(msecs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess16waitForReadyReadEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:231
// index:0
// Public virtual
// bool waitForBytesWritten(int)
func (this *QProcess) WaitForBytesWritten(msecs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess19waitForBytesWrittenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:232
// index:0
// Public
// bool waitForFinished(int)
func (this *QProcess) WaitForFinished(msecs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess15waitForFinishedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:234
// index:0
// Public
// QByteArray readAllStandardOutput()
func (this *QProcess) ReadAllStandardOutput() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess21readAllStandardOutputEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:235
// index:0
// Public
// QByteArray readAllStandardError()
func (this *QProcess) ReadAllStandardError() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess20readAllStandardErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:237
// index:0
// Public
// int exitCode()
func (this *QProcess) ExitCode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess8exitCodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:238
// index:0
// Public
// QProcess::ExitStatus exitStatus()
func (this *QProcess) ExitStatus() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess10exitStatusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:241
// index:0
// Public virtual
// qint64 bytesAvailable()
func (this *QProcess) BytesAvailable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess14bytesAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:242
// index:0
// Public virtual
// qint64 bytesToWrite()
func (this *QProcess) BytesToWrite() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess12bytesToWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:243
// index:0
// Public virtual
// bool isSequential()
func (this *QProcess) IsSequential() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess12isSequentialEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:244
// index:0
// Public virtual
// bool canReadLine()
func (this *QProcess) CanReadLine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess11canReadLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:245
// index:0
// Public virtual
// void close()
func (this *QProcess) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:246
// index:0
// Public virtual
// bool atEnd()
func (this *QProcess) AtEnd() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:248
// index:0
// Public static
// int execute(const class QString &, const class QStringList &)
func (this *QProcess) Execute(program *QString, arguments *QStringList) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess7executeERK7QStringRK11QStringList", ffiqt.FFI_TYPE_POINTER, program, arguments)
	gopp.ErrPrint(err, rv)
	return rv
}
func QProcess_Execute(program *QString, arguments *QStringList) {
	var nilthis *QProcess
	nilthis.Execute(program, arguments)
}

// /usr/include/qt/QtCore/qprocess.h:249
// index:1
// Public static
// int execute(const class QString &)
func (this *QProcess) Execute_1(command *QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess7executeERK7QString", ffiqt.FFI_TYPE_POINTER, command)
	gopp.ErrPrint(err, rv)
	return rv
}
func QProcess_Execute_1(command *QString) {
	var nilthis *QProcess
	nilthis.Execute_1(command)
}

// /usr/include/qt/QtCore/qprocess.h:262
// index:0
// Public static
// QStringList systemEnvironment()
func (this *QProcess) SystemEnvironment() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess17systemEnvironmentEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QProcess_SystemEnvironment() {
	var nilthis *QProcess
	nilthis.SystemEnvironment()
}

// /usr/include/qt/QtCore/qprocess.h:264
// index:0
// Public static
// QString nullDevice()
func (this *QProcess) NullDevice() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess10nullDeviceEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QProcess_NullDevice() {
	var nilthis *QProcess
	nilthis.NullDevice()
}

// /usr/include/qt/QtCore/qprocess.h:267
// index:0
// Public
// void terminate()
func (this *QProcess) Terminate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess9terminateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:268
// index:0
// Public
// void kill()
func (this *QProcess) Kill() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess4killEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:272
// index:0
// Public
// void finished(int)
func (this *QProcess) Finished(exitCode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess8finishedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &exitCode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:273
// index:1
// Public
// void finished(int, class QProcess::ExitStatus)
func (this *QProcess) Finished_1(exitCode int, exitStatus int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess8finishedEiNS_10ExitStatusE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &exitCode, &exitStatus)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:277
// index:0
// Public
// void errorOccurred(class QProcess::ProcessError)
func (this *QProcess) ErrorOccurred(error int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13errorOccurredENS_12ProcessErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &error)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:284
// index:0
// Protected
// void setProcessState(enum QProcess::ProcessState)
func (this *QProcess) SetProcessState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess15setProcessStateENS_12ProcessStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:286
// index:0
// Protected virtual
// void setupChildProcess()
func (this *QProcess) SetupChildProcess() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess17setupChildProcessEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:289
// index:0
// Protected virtual
// qint64 readData(char *, qint64)
func (this *QProcess) ReadData(data string, maxlen int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess8readDataEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:290
// index:0
// Protected virtual
// qint64 writeData(const char *, qint64)
func (this *QProcess) WriteData(data string, len int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
