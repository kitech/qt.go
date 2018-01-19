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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qprocess.h:112
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QProcess) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:158
// index:0
// void QProcess(class QObject *)
func NewQProcess(parent unsafe.Pointer) *QProcess {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcessC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QProcess{cthis}
}

// /usr/include/qt/QtCore/qprocess.h:159
// index:0
// virtual
// void ~QProcess()
func DeleteQProcess(*QProcess) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcessD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:166
// index:0
// bool startDetached(qint64 *)
func (this *QProcess) StartDetached(pid unsafe.Pointer) {
	// 0: (, qint64 * pid), (pid)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedEPx", ffiqt.FFI_TYPE_VOID, this.cthis, pid)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:251
// index:1
// static
// bool startDetached(const class QString &, const class QStringList &, const class QString &, qint64 *)
func (this *QProcess) StartDetached_1(program unsafe.Pointer, arguments unsafe.Pointer, workingDirectory unsafe.Pointer, pid unsafe.Pointer) {
	// 1: (const QString & program, const QStringList & arguments, const QString & workingDirectory, qint64 * pid), (program, arguments, workingDirectory, pid)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QProcess_StartDetached_1(program unsafe.Pointer, arguments unsafe.Pointer, workingDirectory unsafe.Pointer, pid unsafe.Pointer) {
	// 1: (const QString & program, const QStringList & arguments, const QString & workingDirectory, qint64 * pid), (program, arguments, workingDirectory, pid)
	var nilthis *QProcess
	nilthis.StartDetached_1(program, arguments, workingDirectory, pid)
}

// /usr/include/qt/QtCore/qprocess.h:258
// index:2
// static
// bool startDetached(const class QString &, const class QStringList &)
func (this *QProcess) StartDetached_2(program unsafe.Pointer, arguments unsafe.Pointer) {
	// 2: (const QString & program, const QStringList & arguments), (program, arguments)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QStringRK11QStringList", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QProcess_StartDetached_2(program unsafe.Pointer, arguments unsafe.Pointer) {
	// 2: (const QString & program, const QStringList & arguments), (program, arguments)
	var nilthis *QProcess
	nilthis.StartDetached_2(program, arguments)
}

// /usr/include/qt/QtCore/qprocess.h:260
// index:3
// static
// bool startDetached(const class QString &)
func (this *QProcess) StartDetached_3(command unsafe.Pointer) {
	// 3: (const QString & command), (command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QProcess_StartDetached_3(command unsafe.Pointer) {
	// 3: (const QString & command), (command)
	var nilthis *QProcess
	nilthis.StartDetached_3(command)
}

// /usr/include/qt/QtCore/qprocess.h:169
// index:0
// QString program()
func (this *QProcess) Program() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess7programEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:170
// index:0
// void setProgram(const class QString &)
func (this *QProcess) SetProgram(program unsafe.Pointer) {
	// 0: (, const QString & program), (program)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess10setProgramERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, program)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:172
// index:0
// QStringList arguments()
func (this *QProcess) Arguments() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess9argumentsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:173
// index:0
// void setArguments(const class QStringList &)
func (this *QProcess) SetArguments(arguments unsafe.Pointer) {
	// 0: (, const QStringList & arguments), (arguments)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess12setArgumentsERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, arguments)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:175
// index:0
// QProcess::ProcessChannelMode readChannelMode()
func (this *QProcess) ReadChannelMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess15readChannelModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:176
// index:0
// void setReadChannelMode(enum QProcess::ProcessChannelMode)
func (this *QProcess) SetReadChannelMode(mode int) {
	// 0: (, QProcess::ProcessChannelMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess18setReadChannelModeENS_18ProcessChannelModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:177
// index:0
// QProcess::ProcessChannelMode processChannelMode()
func (this *QProcess) ProcessChannelMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess18processChannelModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:178
// index:0
// void setProcessChannelMode(enum QProcess::ProcessChannelMode)
func (this *QProcess) SetProcessChannelMode(mode int) {
	// 0: (, QProcess::ProcessChannelMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess21setProcessChannelModeENS_18ProcessChannelModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:179
// index:0
// QProcess::InputChannelMode inputChannelMode()
func (this *QProcess) InputChannelMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess16inputChannelModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:180
// index:0
// void setInputChannelMode(enum QProcess::InputChannelMode)
func (this *QProcess) SetInputChannelMode(mode int) {
	// 0: (, QProcess::InputChannelMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess19setInputChannelModeENS_16InputChannelModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:182
// index:0
// QProcess::ProcessChannel readChannel()
func (this *QProcess) ReadChannel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess11readChannelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:183
// index:0
// void setReadChannel(enum QProcess::ProcessChannel)
func (this *QProcess) SetReadChannel(channel int) {
	// 0: (, QProcess::ProcessChannel channel), (&channel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess14setReadChannelENS_14ProcessChannelE", ffiqt.FFI_TYPE_VOID, this.cthis, &channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:185
// index:0
// void closeReadChannel(enum QProcess::ProcessChannel)
func (this *QProcess) CloseReadChannel(channel int) {
	// 0: (, QProcess::ProcessChannel channel), (&channel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess16closeReadChannelENS_14ProcessChannelE", ffiqt.FFI_TYPE_VOID, this.cthis, &channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:186
// index:0
// void closeWriteChannel()
func (this *QProcess) CloseWriteChannel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess17closeWriteChannelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:188
// index:0
// void setStandardInputFile(const class QString &)
func (this *QProcess) SetStandardInputFile(fileName unsafe.Pointer) {
	// 0: (, const QString & fileName), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess20setStandardInputFileERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, fileName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:191
// index:0
// void setStandardOutputProcess(class QProcess *)
func (this *QProcess) SetStandardOutputProcess(destination unsafe.Pointer) {
	// 0: (, QProcess * destination), (destination)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess24setStandardOutputProcessEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, destination)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:214
// index:0
// QString workingDirectory()
func (this *QProcess) WorkingDirectory() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess16workingDirectoryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:215
// index:0
// void setWorkingDirectory(const class QString &)
func (this *QProcess) SetWorkingDirectory(dir unsafe.Pointer) {
	// 0: (, const QString & dir), (dir)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess19setWorkingDirectoryERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, dir)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:217
// index:0
// void setEnvironment(const class QStringList &)
func (this *QProcess) SetEnvironment(environment unsafe.Pointer) {
	// 0: (, const QStringList & environment), (environment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess14setEnvironmentERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, environment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:218
// index:0
// QStringList environment()
func (this *QProcess) Environment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess11environmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:219
// index:0
// void setProcessEnvironment(const class QProcessEnvironment &)
func (this *QProcess) SetProcessEnvironment(environment unsafe.Pointer) {
	// 0: (, const QProcessEnvironment & environment), (environment)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess21setProcessEnvironmentERK19QProcessEnvironment", ffiqt.FFI_TYPE_VOID, this.cthis, environment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:220
// index:0
// QProcessEnvironment processEnvironment()
func (this *QProcess) ProcessEnvironment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess18processEnvironmentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:222
// index:0
// QProcess::ProcessError error()
func (this *QProcess) Error() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess5errorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:275
// index:1
// void error(class QProcess::ProcessError)
func (this *QProcess) Error_1(error int) {
	// 1: (, QProcess::ProcessError error), (&error)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess5errorENS_12ProcessErrorE", ffiqt.FFI_TYPE_VOID, this.cthis, &error)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:223
// index:0
// QProcess::ProcessState state()
func (this *QProcess) State() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess5stateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:226
// index:0
// Q_PID pid()
func (this *QProcess) Pid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess3pidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:227
// index:0
// qint64 processId()
func (this *QProcess) ProcessId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess9processIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:229
// index:0
// bool waitForStarted(int)
func (this *QProcess) WaitForStarted(msecs int) {
	// 0: (, int msecs), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess14waitForStartedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:230
// index:0
// virtual
// bool waitForReadyRead(int)
func (this *QProcess) WaitForReadyRead(msecs int) {
	// 0: (, int msecs), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess16waitForReadyReadEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:231
// index:0
// virtual
// bool waitForBytesWritten(int)
func (this *QProcess) WaitForBytesWritten(msecs int) {
	// 0: (, int msecs), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess19waitForBytesWrittenEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:232
// index:0
// bool waitForFinished(int)
func (this *QProcess) WaitForFinished(msecs int) {
	// 0: (, int msecs), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess15waitForFinishedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:234
// index:0
// QByteArray readAllStandardOutput()
func (this *QProcess) ReadAllStandardOutput() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess21readAllStandardOutputEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:235
// index:0
// QByteArray readAllStandardError()
func (this *QProcess) ReadAllStandardError() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess20readAllStandardErrorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:237
// index:0
// int exitCode()
func (this *QProcess) ExitCode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess8exitCodeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:238
// index:0
// QProcess::ExitStatus exitStatus()
func (this *QProcess) ExitStatus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess10exitStatusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:241
// index:0
// virtual
// qint64 bytesAvailable()
func (this *QProcess) BytesAvailable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess14bytesAvailableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:242
// index:0
// virtual
// qint64 bytesToWrite()
func (this *QProcess) BytesToWrite() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess12bytesToWriteEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:243
// index:0
// virtual
// bool isSequential()
func (this *QProcess) IsSequential() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess12isSequentialEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:244
// index:0
// virtual
// bool canReadLine()
func (this *QProcess) CanReadLine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess11canReadLineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:245
// index:0
// virtual
// void close()
func (this *QProcess) Close() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess5closeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:246
// index:0
// virtual
// bool atEnd()
func (this *QProcess) AtEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess5atEndEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:248
// index:0
// static
// int execute(const class QString &, const class QStringList &)
func (this *QProcess) Execute(program unsafe.Pointer, arguments unsafe.Pointer) {
	// 0: (const QString & program, const QStringList & arguments), (program, arguments)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess7executeERK7QStringRK11QStringList", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QProcess_Execute(program unsafe.Pointer, arguments unsafe.Pointer) {
	// 0: (const QString & program, const QStringList & arguments), (program, arguments)
	var nilthis *QProcess
	nilthis.Execute(program, arguments)
}

// /usr/include/qt/QtCore/qprocess.h:249
// index:1
// static
// int execute(const class QString &)
func (this *QProcess) Execute_1(command unsafe.Pointer) {
	// 1: (const QString & command), (command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess7executeERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QProcess_Execute_1(command unsafe.Pointer) {
	// 1: (const QString & command), (command)
	var nilthis *QProcess
	nilthis.Execute_1(command)
}

// /usr/include/qt/QtCore/qprocess.h:262
// index:0
// static
// QStringList systemEnvironment()
func (this *QProcess) SystemEnvironment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess17systemEnvironmentEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QProcess_SystemEnvironment() {
	// 0: (), ()
	var nilthis *QProcess
	nilthis.SystemEnvironment()
}

// /usr/include/qt/QtCore/qprocess.h:264
// index:0
// static
// QString nullDevice()
func (this *QProcess) NullDevice() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess10nullDeviceEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QProcess_NullDevice() {
	// 0: (), ()
	var nilthis *QProcess
	nilthis.NullDevice()
}

// /usr/include/qt/QtCore/qprocess.h:267
// index:0
// void terminate()
func (this *QProcess) Terminate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess9terminateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:268
// index:0
// void kill()
func (this *QProcess) Kill() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess4killEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:272
// index:0
// void finished(int)
func (this *QProcess) Finished(exitCode int) {
	// 0: (, int exitCode), (&exitCode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess8finishedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &exitCode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:273
// index:1
// void finished(int, class QProcess::ExitStatus)
func (this *QProcess) Finished_1(exitCode int, exitStatus int) {
	// 1: (, int exitCode, QProcess::ExitStatus exitStatus), (&exitCode, &exitStatus)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess8finishedEiNS_10ExitStatusE", ffiqt.FFI_TYPE_VOID, this.cthis, &exitCode, &exitStatus)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:277
// index:0
// void errorOccurred(class QProcess::ProcessError)
func (this *QProcess) ErrorOccurred(error int) {
	// 0: (, QProcess::ProcessError error), (&error)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13errorOccurredENS_12ProcessErrorE", ffiqt.FFI_TYPE_VOID, this.cthis, &error)
	gopp.ErrPrint(err, rv)
}

//  body block end
