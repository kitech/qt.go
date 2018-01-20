//  header block begin
// /usr/include/qt/QtCore/qiodevice.h
// #include <qiodevice.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
type QIODevice struct {
	*QObject
}

func (this *QIODevice) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtCore/qiodevice.h:68
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QIODevice) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:83
// index:0
// void QIODevice()
func NewQIODevice() *QIODevice {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODeviceC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQIODeviceFromPointer(cthis)
	return gothis
}
func NewQIODeviceFromPointer(cthis unsafe.Pointer) *QIODevice {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QIODevice{bcthis0}
}

// /usr/include/qt/QtCore/qiodevice.h:85
// index:1
// void QIODevice(class QObject *)
func NewQIODevice_1(parent unsafe.Pointer) *QIODevice {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODeviceC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQIODeviceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qiodevice.h:164
// index:2
// void QIODevice(class QIODevicePrivate &, class QObject *)
func NewQIODevice_2(dd unsafe.Pointer, parent unsafe.Pointer) *QIODevice {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODeviceC1ER16QIODevicePrivateP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQIODeviceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qiodevice.h:87
// index:0
// virtual
// void ~QIODevice()
func DeleteQIODevice(*QIODevice) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODeviceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:89
// index:0
// QIODevice::OpenMode openMode()
func (this *QIODevice) OpenMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice8openModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:91
// index:0
// void setTextModeEnabled(_Bool)
func (this *QIODevice) SetTextModeEnabled(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice18setTextModeEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:92
// index:0
// bool isTextModeEnabled()
func (this *QIODevice) IsTextModeEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice17isTextModeEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:94
// index:0
// bool isOpen()
func (this *QIODevice) IsOpen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice6isOpenEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:95
// index:0
// bool isReadable()
func (this *QIODevice) IsReadable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice10isReadableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:96
// index:0
// bool isWritable()
func (this *QIODevice) IsWritable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice10isWritableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:97
// index:0
// virtual
// bool isSequential()
func (this *QIODevice) IsSequential() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice12isSequentialEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:99
// index:0
// int readChannelCount()
func (this *QIODevice) ReadChannelCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice16readChannelCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:100
// index:0
// int writeChannelCount()
func (this *QIODevice) WriteChannelCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice17writeChannelCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:101
// index:0
// int currentReadChannel()
func (this *QIODevice) CurrentReadChannel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice18currentReadChannelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:102
// index:0
// void setCurrentReadChannel(int)
func (this *QIODevice) SetCurrentReadChannel(channel int) {
	// 0: (, channel int), (&channel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice21setCurrentReadChannelEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:103
// index:0
// int currentWriteChannel()
func (this *QIODevice) CurrentWriteChannel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice19currentWriteChannelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:104
// index:0
// void setCurrentWriteChannel(int)
func (this *QIODevice) SetCurrentWriteChannel(channel int) {
	// 0: (, channel int), (&channel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice22setCurrentWriteChannelEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:106
// index:0
// virtual
// bool open(QIODevice::OpenMode)
func (this *QIODevice) Open(mode int) {
	// 0: (, QFlags<QIODevice::OpenModeFlag> mode), (mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4openE6QFlagsINS_12OpenModeFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:107
// index:0
// virtual
// void close()
func (this *QIODevice) Close() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice5closeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:111
// index:0
// virtual
// qint64 pos()
func (this *QIODevice) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice3posEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:112
// index:0
// virtual
// qint64 size()
func (this *QIODevice) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice4sizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:113
// index:0
// virtual
// bool seek(qint64)
func (this *QIODevice) Seek(pos int64) {
	// 0: (, pos qint64), (&pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4seekEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:114
// index:0
// virtual
// bool atEnd()
func (this *QIODevice) AtEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice5atEndEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:115
// index:0
// virtual
// bool reset()
func (this *QIODevice) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice5resetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:117
// index:0
// virtual
// qint64 bytesAvailable()
func (this *QIODevice) BytesAvailable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice14bytesAvailableEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:118
// index:0
// virtual
// qint64 bytesToWrite()
func (this *QIODevice) BytesToWrite() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice12bytesToWriteEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:120
// index:0
// qint64 read(char *, qint64)
func (this *QIODevice) Read(data unsafe.Pointer, maxlen int64) {
	// 0: (, data char *, maxlen qint64), (data, &maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4readEPcx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:121
// index:1
// QByteArray read(qint64)
func (this *QIODevice) Read_1(maxlen int64) {
	// 1: (, maxlen qint64), (&maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4readEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:122
// index:0
// QByteArray readAll()
func (this *QIODevice) ReadAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice7readAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:123
// index:0
// qint64 readLine(char *, qint64)
func (this *QIODevice) ReadLine(data unsafe.Pointer, maxlen int64) {
	// 0: (, data char *, maxlen qint64), (data, &maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice8readLineEPcx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:124
// index:1
// QByteArray readLine(qint64)
func (this *QIODevice) ReadLine_1(maxlen int64) {
	// 1: (, maxlen qint64), (&maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice8readLineEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:125
// index:0
// virtual
// bool canReadLine()
func (this *QIODevice) CanReadLine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice11canReadLineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:127
// index:0
// void startTransaction()
func (this *QIODevice) StartTransaction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice16startTransactionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:128
// index:0
// void commitTransaction()
func (this *QIODevice) CommitTransaction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice17commitTransactionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:129
// index:0
// void rollbackTransaction()
func (this *QIODevice) RollbackTransaction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice19rollbackTransactionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:130
// index:0
// bool isTransactionStarted()
func (this *QIODevice) IsTransactionStarted() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice20isTransactionStartedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:132
// index:0
// qint64 write(const char *, qint64)
func (this *QIODevice) Write(data unsafe.Pointer, len int64) {
	// 0: (, data const char *, len qint64), (data, &len)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice5writeEPKcx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:133
// index:1
// qint64 write(const char *)
func (this *QIODevice) Write_1(data unsafe.Pointer) {
	// 1: (, data const char *), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice5writeEPKc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:134
// index:2
// inline
// qint64 write(const class QByteArray &)
func (this *QIODevice) Write_2(data unsafe.Pointer) {
	// 2: (, data const QByteArray &), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice5writeERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:137
// index:0
// qint64 peek(char *, qint64)
func (this *QIODevice) Peek(data unsafe.Pointer, maxlen int64) {
	// 0: (, data char *, maxlen qint64), (data, &maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4peekEPcx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:138
// index:1
// QByteArray peek(qint64)
func (this *QIODevice) Peek_1(maxlen int64) {
	// 1: (, maxlen qint64), (&maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4peekEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:139
// index:0
// qint64 skip(qint64)
func (this *QIODevice) Skip(maxSize int64) {
	// 0: (, maxSize qint64), (&maxSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4skipEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &maxSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:141
// index:0
// virtual
// bool waitForReadyRead(int)
func (this *QIODevice) WaitForReadyRead(msecs int) {
	// 0: (, msecs int), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice16waitForReadyReadEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:142
// index:0
// virtual
// bool waitForBytesWritten(int)
func (this *QIODevice) WaitForBytesWritten(msecs int) {
	// 0: (, msecs int), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice19waitForBytesWrittenEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:144
// index:0
// void ungetChar(char)
func (this *QIODevice) UngetChar(c byte) {
	// 0: (, c char), (&c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice9ungetCharEc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:145
// index:0
// bool putChar(char)
func (this *QIODevice) PutChar(c byte) {
	// 0: (, c char), (&c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice7putCharEc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:146
// index:0
// bool getChar(char *)
func (this *QIODevice) GetChar(c unsafe.Pointer) {
	// 0: (, c char *), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice7getCharEPc", ffiqt.FFI_TYPE_VOID, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:148
// index:0
// QString errorString()
func (this *QIODevice) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice11errorStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:152
// index:0
// void readyRead()
func (this *QIODevice) ReadyRead() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice9readyReadEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:153
// index:0
// void channelReadyRead(int)
func (this *QIODevice) ChannelReadyRead(channel int) {
	// 0: (, channel int), (&channel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice16channelReadyReadEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:154
// index:0
// void bytesWritten(qint64)
func (this *QIODevice) BytesWritten(bytes int64) {
	// 0: (, bytes qint64), (&bytes)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice12bytesWrittenEx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &bytes)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:155
// index:0
// void channelBytesWritten(int, qint64)
func (this *QIODevice) ChannelBytesWritten(channel int, bytes int64) {
	// 0: (, channel int, bytes qint64), (&channel, &bytes)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice19channelBytesWrittenEix", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &channel, &bytes)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:156
// index:0
// void aboutToClose()
func (this *QIODevice) AboutToClose() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice12aboutToCloseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:157
// index:0
// void readChannelFinished()
func (this *QIODevice) ReadChannelFinished() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice19readChannelFinishedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:166
// index:0
// pure virtual
// qint64 readData(char *, qint64)
func (this *QIODevice) ReadData(data unsafe.Pointer, maxlen int64) {
	// 0: (, data char *, maxlen qint64), (data, &maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice8readDataEPcx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:167
// index:0
// virtual
// qint64 readLineData(char *, qint64)
func (this *QIODevice) ReadLineData(data unsafe.Pointer, maxlen int64) {
	// 0: (, data char *, maxlen qint64), (data, &maxlen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice12readLineDataEPcx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &maxlen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:168
// index:0
// pure virtual
// qint64 writeData(const char *, qint64)
func (this *QIODevice) WriteData(data unsafe.Pointer, len int64) {
	// 0: (, data const char *, len qint64), (data, &len)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice9writeDataEPKcx", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:170
// index:0
// void setOpenMode(QIODevice::OpenMode)
func (this *QIODevice) SetOpenMode(openMode int) {
	// 0: (, QFlags<QIODevice::OpenModeFlag> openMode), (openMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice11setOpenModeE6QFlagsINS_12OpenModeFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), openMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:172
// index:0
// void setErrorString(const class QString &)
func (this *QIODevice) SetErrorString(errorString unsafe.Pointer) {
	// 0: (, errorString const QString &), (errorString)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice14setErrorStringERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), errorString)
	gopp.ErrPrint(err, rv)
}

//  body block end
