package qtcore

// /usr/include/qt/QtCore/qprocess.h
// #include <qprocess.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// void setProcessState(enum QProcess::ProcessState)
func (this *QProcess) InheritSetProcessState(f func(state int)) {
	ffiqt.SetAllInheritCallback(this, "setProcessState", f)
}

// void setupChildProcess()
func (this *QProcess) InheritSetupChildProcess(f func()) {
	ffiqt.SetAllInheritCallback(this, "setupChildProcess", f)
}

// qint64 readData(char *, qint64)
func (this *QProcess) InheritReadData(f func(data string, maxlen int64) int64) {
	ffiqt.SetAllInheritCallback(this, "readData", f)
}

// qint64 writeData(const char *, qint64)
func (this *QProcess) InheritWriteData(f func(data string, len int64) int64) {
	ffiqt.SetAllInheritCallback(this, "writeData", f)
}

type QProcess struct {
	*QIODevice
}

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
// [8] const QMetaObject * metaObject()
func (this *QProcess) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProcess(QObject *)
func NewQProcess(parent *QObject /*777 QObject **/) *QProcess {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcessC2EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQProcessFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qprocess.h:159
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QProcess()
func DeleteQProcess(this *QProcess) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcessD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qprocess.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start(const QString &, const QStringList &, QIODevice::OpenMode)
func (this *QProcess) Start(program *QString, arguments *QStringList, mode int) {
	var convArg0 = program.GetCthis()
	var convArg1 = arguments.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess5startERK7QStringRK11QStringList6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:163
// index:1
// Public Visibility=Default Availability=Available
// [-2] void start(const QString &, QIODevice::OpenMode)
func (this *QProcess) Start_1(command *QString, mode int) {
	var convArg0 = command.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess5startERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:165
// index:2
// Public Visibility=Default Availability=Available
// [-2] void start(QIODevice::OpenMode)
func (this *QProcess) Start_2(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess5startE6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool startDetached(qint64 *)
func (this *QProcess) StartDetached(pid unsafe.Pointer /*666*/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedEPx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pid)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:251
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool startDetached(const QString &, const QStringList &, const QString &, qint64 *)
func (this *QProcess) StartDetached_1(program *QString, arguments *QStringList, workingDirectory *QString, pid unsafe.Pointer /*666*/) bool {
	var convArg0 = program.GetCthis()
	var convArg1 = arguments.GetCthis()
	var convArg2 = workingDirectory.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QStringRK11QStringListS2_Px", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, &pid)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QProcess_StartDetached_1(program *QString, arguments *QStringList, workingDirectory *QString, pid unsafe.Pointer /*666*/) bool {
	var nilthis *QProcess
	rv := nilthis.StartDetached_1(program, arguments, workingDirectory, pid)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:258
// index:2
// Public static Visibility=Default Availability=Available
// [1] bool startDetached(const QString &, const QStringList &)
func (this *QProcess) StartDetached_2(program *QString, arguments *QStringList) bool {
	var convArg0 = program.GetCthis()
	var convArg1 = arguments.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QStringRK11QStringList", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QProcess_StartDetached_2(program *QString, arguments *QStringList) bool {
	var nilthis *QProcess
	rv := nilthis.StartDetached_2(program, arguments)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:260
// index:3
// Public static Visibility=Default Availability=Available
// [1] bool startDetached(const QString &)
func (this *QProcess) StartDetached_3(command *QString) bool {
	var convArg0 = command.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13startDetachedERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QProcess_StartDetached_3(command *QString) bool {
	var nilthis *QProcess
	rv := nilthis.StartDetached_3(command)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)
func (this *QProcess) Open(mode int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess4openE6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QString program()
func (this *QProcess) Program() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess7programEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProgram(const QString &)
func (this *QProcess) SetProgram(program *QString) {
	var convArg0 = program.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess10setProgramERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setArguments(const QStringList &)
func (this *QProcess) SetArguments(arguments *QStringList) {
	var convArg0 = arguments.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess12setArgumentsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:175
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessChannelMode readChannelMode()
func (this *QProcess) ReadChannelMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess15readChannelModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadChannelMode(enum QProcess::ProcessChannelMode)
func (this *QProcess) SetReadChannelMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess18setReadChannelModeENS_18ProcessChannelModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:177
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessChannelMode processChannelMode()
func (this *QProcess) ProcessChannelMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess18processChannelModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProcessChannelMode(enum QProcess::ProcessChannelMode)
func (this *QProcess) SetProcessChannelMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess21setProcessChannelModeENS_18ProcessChannelModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:179
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::InputChannelMode inputChannelMode()
func (this *QProcess) InputChannelMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess16inputChannelModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputChannelMode(enum QProcess::InputChannelMode)
func (this *QProcess) SetInputChannelMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess19setInputChannelModeENS_16InputChannelModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:182
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessChannel readChannel()
func (this *QProcess) ReadChannel() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess11readChannelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadChannel(enum QProcess::ProcessChannel)
func (this *QProcess) SetReadChannel(channel int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess14setReadChannelENS_14ProcessChannelE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeReadChannel(enum QProcess::ProcessChannel)
func (this *QProcess) CloseReadChannel(channel int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess16closeReadChannelENS_14ProcessChannelE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeWriteChannel()
func (this *QProcess) CloseWriteChannel() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess17closeWriteChannelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardInputFile(const QString &)
func (this *QProcess) SetStandardInputFile(fileName *QString) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess20setStandardInputFileERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardOutputFile(const QString &, QIODevice::OpenMode)
func (this *QProcess) SetStandardOutputFile(fileName *QString, mode int) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess21setStandardOutputFileERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardErrorFile(const QString &, QIODevice::OpenMode)
func (this *QProcess) SetStandardErrorFile(fileName *QString, mode int) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess20setStandardErrorFileERK7QString6QFlagsIN9QIODevice12OpenModeFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardOutputProcess(QProcess *)
func (this *QProcess) SetStandardOutputProcess(destination *QProcess /*777 QProcess **/) {
	var convArg0 = destination.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess24setStandardOutputProcessEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:214
// index:0
// Public Visibility=Default Availability=Available
// [8] QString workingDirectory()
func (this *QProcess) WorkingDirectory() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess16workingDirectoryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWorkingDirectory(const QString &)
func (this *QProcess) SetWorkingDirectory(dir *QString) {
	var convArg0 = dir.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess19setWorkingDirectoryERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnvironment(const QStringList &)
func (this *QProcess) SetEnvironment(environment *QStringList) {
	var convArg0 = environment.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess14setEnvironmentERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProcessEnvironment(const QProcessEnvironment &)
func (this *QProcess) SetProcessEnvironment(environment *QProcessEnvironment) {
	var convArg0 = environment.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess21setProcessEnvironmentERK19QProcessEnvironment", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:220
// index:0
// Public Visibility=Default Availability=Available
// [8] QProcessEnvironment processEnvironment()
func (this *QProcess) ProcessEnvironment() *QProcessEnvironment /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess18processEnvironmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQProcessEnvironmentFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQProcessEnvironment)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:222
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessError error()
func (this *QProcess) Error() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess5errorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:275
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QProcess::ProcessError)
func (this *QProcess) Error_1(error int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess5errorENS_12ProcessErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), error)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:223
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessState state()
func (this *QProcess) State() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:226
// index:0
// Public Visibility=Default Availability=Available
// [8] Q_PID pid()
func (this *QProcess) Pid() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess3pidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qprocess.h:227
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 processId()
func (this *QProcess) ProcessId() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess9processIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qprocess.h:229
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForStarted(int)
func (this *QProcess) WaitForStarted(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess14waitForStartedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:230
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)
func (this *QProcess) WaitForReadyRead(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess16waitForReadyReadEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:231
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)
func (this *QProcess) WaitForBytesWritten(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess19waitForBytesWrittenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:232
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForFinished(int)
func (this *QProcess) WaitForFinished(msecs int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess15waitForFinishedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:234
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray readAllStandardOutput()
func (this *QProcess) ReadAllStandardOutput() *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess21readAllStandardOutputEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:235
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray readAllStandardError()
func (this *QProcess) ReadAllStandardError() *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess20readAllStandardErrorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qprocess.h:237
// index:0
// Public Visibility=Default Availability=Available
// [4] int exitCode()
func (this *QProcess) ExitCode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess8exitCodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qprocess.h:238
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ExitStatus exitStatus()
func (this *QProcess) ExitStatus() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess10exitStatusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:241
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable()
func (this *QProcess) BytesAvailable() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess14bytesAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qprocess.h:242
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite()
func (this *QProcess) BytesToWrite() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess12bytesToWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qprocess.h:243
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential()
func (this *QProcess) IsSequential() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess12isSequentialEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:244
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine()
func (this *QProcess) CanReadLine() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess11canReadLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:245
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()
func (this *QProcess) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:246
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd()
func (this *QProcess) AtEnd() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QProcess5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:248
// index:0
// Public static Visibility=Default Availability=Available
// [4] int execute(const QString &, const QStringList &)
func (this *QProcess) Execute(program *QString, arguments *QStringList) int {
	var convArg0 = program.GetCthis()
	var convArg1 = arguments.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess7executeERK7QStringRK11QStringList", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QProcess_Execute(program *QString, arguments *QStringList) int {
	var nilthis *QProcess
	rv := nilthis.Execute(program, arguments)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:249
// index:1
// Public static Visibility=Default Availability=Available
// [4] int execute(const QString &)
func (this *QProcess) Execute_1(command *QString) int {
	var convArg0 = command.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess7executeERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QProcess_Execute_1(command *QString) int {
	var nilthis *QProcess
	rv := nilthis.Execute_1(command)
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:264
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString nullDevice()
func (this *QProcess) NullDevice() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess10nullDeviceEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QProcess_NullDevice() *QString /*123*/ {
	var nilthis *QProcess
	rv := nilthis.NullDevice()
	return rv
}

// /usr/include/qt/QtCore/qprocess.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void terminate()
func (this *QProcess) Terminate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess9terminateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void kill()
func (this *QProcess) Kill() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess4killEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:272
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished(int)
func (this *QProcess) Finished(exitCode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess8finishedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), exitCode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:273
// index:1
// Public Visibility=Default Availability=Available
// [-2] void finished(int, QProcess::ExitStatus)
func (this *QProcess) Finished_1(exitCode int, exitStatus int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess8finishedEiNS_10ExitStatusE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), exitCode, exitStatus)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:277
// index:0
// Public Visibility=Default Availability=Available
// [-2] void errorOccurred(QProcess::ProcessError)
func (this *QProcess) ErrorOccurred(error int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess13errorOccurredENS_12ProcessErrorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), error)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:284
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setProcessState(enum QProcess::ProcessState)
func (this *QProcess) SetProcessState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess15setProcessStateENS_12ProcessStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:286
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setupChildProcess()
func (this *QProcess) SetupChildProcess() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess17setupChildProcessEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:289
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)
func (this *QProcess) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess8readDataEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qprocess.h:290
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)
func (this *QProcess) WriteData(data string, len int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QProcess9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

type QProcess__ProcessError = int

const QProcess__FailedToStart QProcess__ProcessError = 0
const QProcess__Crashed QProcess__ProcessError = 1
const QProcess__Timedout QProcess__ProcessError = 2
const QProcess__ReadError QProcess__ProcessError = 3
const QProcess__WriteError QProcess__ProcessError = 4
const QProcess__UnknownError QProcess__ProcessError = 5

type QProcess__ProcessState = int

const QProcess__NotRunning QProcess__ProcessState = 0
const QProcess__Starting QProcess__ProcessState = 1
const QProcess__Running QProcess__ProcessState = 2

type QProcess__ProcessChannel = int

const QProcess__StandardOutput QProcess__ProcessChannel = 0
const QProcess__StandardError QProcess__ProcessChannel = 1

type QProcess__ProcessChannelMode = int

const QProcess__SeparateChannels QProcess__ProcessChannelMode = 0
const QProcess__MergedChannels QProcess__ProcessChannelMode = 1
const QProcess__ForwardedChannels QProcess__ProcessChannelMode = 2
const QProcess__ForwardedOutputChannel QProcess__ProcessChannelMode = 3
const QProcess__ForwardedErrorChannel QProcess__ProcessChannelMode = 4

type QProcess__InputChannelMode = int

const QProcess__ManagedInputChannel QProcess__InputChannelMode = 0
const QProcess__ForwardedInputChannel QProcess__InputChannelMode = 1

type QProcess__ExitStatus = int

const QProcess__NormalExit QProcess__ExitStatus = 0
const QProcess__CrashExit QProcess__ExitStatus = 1

//  body block end
