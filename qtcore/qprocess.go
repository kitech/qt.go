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
func (this *QProcess) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qprocess.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QProcess(QObject *)
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
func (this *QProcess) Start_2(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5startE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:165
// index:2
// Public Visibility=Default Availability=Available
// [-2] void start(QIODevice::OpenMode)
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
func (this *QProcess) StartDetached(pid unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess13startDetachedEPx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pid)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool startDetached(qint64 *)
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
func (this *QProcess) Open(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)
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
func (this *QProcess) ReadChannelMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess15readChannelModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadChannelMode(enum QProcess::ProcessChannelMode)
func (this *QProcess) SetReadChannelMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess18setReadChannelModeENS_18ProcessChannelModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:177
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessChannelMode processChannelMode() const
func (this *QProcess) ProcessChannelMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess18processChannelModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProcessChannelMode(enum QProcess::ProcessChannelMode)
func (this *QProcess) SetProcessChannelMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess21setProcessChannelModeENS_18ProcessChannelModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:179
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::InputChannelMode inputChannelMode() const
func (this *QProcess) InputChannelMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess16inputChannelModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputChannelMode(enum QProcess::InputChannelMode)
func (this *QProcess) SetInputChannelMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess19setInputChannelModeENS_16InputChannelModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:182
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessChannel readChannel() const
func (this *QProcess) ReadChannel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess11readChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadChannel(enum QProcess::ProcessChannel)
func (this *QProcess) SetReadChannel(channel int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess14setReadChannelENS_14ProcessChannelE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeReadChannel(enum QProcess::ProcessChannel)
func (this *QProcess) CloseReadChannel(channel int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess16closeReadChannelENS_14ProcessChannelE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), channel)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void closeWriteChannel()
func (this *QProcess) CloseWriteChannel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess17closeWriteChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStandardInputFile(const QString &)
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
func (this *QProcess) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:275
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QProcess::ProcessError)
func (this *QProcess) Error_1(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5errorENS_12ProcessErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:223
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ProcessState state() const
func (this *QProcess) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:226
// index:0
// Public Visibility=Default Availability=Available
// [8] Q_PID pid() const
func (this *QProcess) Pid() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess3pidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qprocess.h:227
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 processId() const
func (this *QProcess) ProcessId() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess9processIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qprocess.h:229
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForStarted(int)
func (this *QProcess) WaitForStarted(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess14waitForStartedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:229
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForStarted(int)
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
func (this *QProcess) WaitForReadyRead(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess16waitForReadyReadEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:230
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForReadyRead(int)
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
func (this *QProcess) WaitForBytesWritten(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess19waitForBytesWrittenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:231
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool waitForBytesWritten(int)
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
func (this *QProcess) WaitForFinished(msecs int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess15waitForFinishedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msecs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:232
// index:0
// Public Visibility=Default Availability=Available
// [1] bool waitForFinished(int)
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
func (this *QProcess) ExitCode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess8exitCodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qprocess.h:238
// index:0
// Public Visibility=Default Availability=Available
// [4] QProcess::ExitStatus exitStatus() const
func (this *QProcess) ExitStatus() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess10exitStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qprocess.h:241
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesAvailable() const
func (this *QProcess) BytesAvailable() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess14bytesAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qprocess.h:242
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 bytesToWrite() const
func (this *QProcess) BytesToWrite() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess12bytesToWriteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qprocess.h:243
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSequential() const
func (this *QProcess) IsSequential() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess12isSequentialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:244
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine() const
func (this *QProcess) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:245
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()
func (this *QProcess) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:246
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd() const
func (this *QProcess) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QProcess5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qprocess.h:248
// index:0
// Public static Visibility=Default Availability=Available
// [4] int execute(const QString &, const QStringList &)
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
func (this *QProcess) Terminate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess9terminateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void kill()
func (this *QProcess) Kill() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess4killEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:272
// index:0
// Public Visibility=Default Availability=Available
// [-2] void finished(int)
func (this *QProcess) Finished(exitCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess8finishedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), exitCode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:273
// index:1
// Public Visibility=Default Availability=Available
// [-2] void finished(int, QProcess::ExitStatus)
func (this *QProcess) Finished_1(exitCode int, exitStatus int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess8finishedEiNS_10ExitStatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), exitCode, exitStatus)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:277
// index:0
// Public Visibility=Default Availability=Available
// [-2] void errorOccurred(QProcess::ProcessError)
func (this *QProcess) ErrorOccurred(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess13errorOccurredENS_12ProcessErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:284
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setProcessState(enum QProcess::ProcessState)
func (this *QProcess) SetProcessState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess15setProcessStateENS_12ProcessStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:286
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setupChildProcess()
func (this *QProcess) SetupChildProcess() {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess17setupChildProcessEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qprocess.h:289
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)
func (this *QProcess) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qprocess.h:290
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)
func (this *QProcess) WriteData(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QProcess9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
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
