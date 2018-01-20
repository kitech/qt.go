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
func NewQIODeviceFromPointer(cthis unsafe.Pointer) *QIODevice {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QIODevice{bcthis0}
}

// /usr/include/qt/QtCore/qiodevice.h:68
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QIODevice) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:83
// index:0
// Public
// void QIODevice()
func NewQIODevice() *QIODevice {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODeviceC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQIODeviceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qiodevice.h:85
// index:1
// Public
// void QIODevice(class QObject *)
func NewQIODevice_1(parent unsafe.Pointer) *QIODevice {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODeviceC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQIODeviceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qiodevice.h:87
// index:0
// Public virtual
// void ~QIODevice()
func DeleteQIODevice(*QIODevice) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODeviceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:89
// index:0
// Public
// QIODevice::OpenMode openMode()
func (this *QIODevice) OpenMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice8openModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:91
// index:0
// Public
// void setTextModeEnabled(_Bool)
func (this *QIODevice) SetTextModeEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice18setTextModeEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:92
// index:0
// Public
// bool isTextModeEnabled()
func (this *QIODevice) IsTextModeEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice17isTextModeEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:94
// index:0
// Public
// bool isOpen()
func (this *QIODevice) IsOpen() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice6isOpenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:95
// index:0
// Public
// bool isReadable()
func (this *QIODevice) IsReadable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice10isReadableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:96
// index:0
// Public
// bool isWritable()
func (this *QIODevice) IsWritable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice10isWritableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:97
// index:0
// Public virtual
// bool isSequential()
func (this *QIODevice) IsSequential() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice12isSequentialEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:99
// index:0
// Public
// int readChannelCount()
func (this *QIODevice) ReadChannelCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice16readChannelCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:100
// index:0
// Public
// int writeChannelCount()
func (this *QIODevice) WriteChannelCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice17writeChannelCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:101
// index:0
// Public
// int currentReadChannel()
func (this *QIODevice) CurrentReadChannel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice18currentReadChannelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:102
// index:0
// Public
// void setCurrentReadChannel(int)
func (this *QIODevice) SetCurrentReadChannel(channel int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice21setCurrentReadChannelEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:103
// index:0
// Public
// int currentWriteChannel()
func (this *QIODevice) CurrentWriteChannel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice19currentWriteChannelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:104
// index:0
// Public
// void setCurrentWriteChannel(int)
func (this *QIODevice) SetCurrentWriteChannel(channel int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice22setCurrentWriteChannelEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:107
// index:0
// Public virtual
// void close()
func (this *QIODevice) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:111
// index:0
// Public virtual
// qint64 pos()
func (this *QIODevice) Pos() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:112
// index:0
// Public virtual
// qint64 size()
func (this *QIODevice) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:113
// index:0
// Public virtual
// bool seek(qint64)
func (this *QIODevice) Seek(pos int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4seekEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:114
// index:0
// Public virtual
// bool atEnd()
func (this *QIODevice) AtEnd() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:115
// index:0
// Public virtual
// bool reset()
func (this *QIODevice) Reset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:117
// index:0
// Public virtual
// qint64 bytesAvailable()
func (this *QIODevice) BytesAvailable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice14bytesAvailableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:118
// index:0
// Public virtual
// qint64 bytesToWrite()
func (this *QIODevice) BytesToWrite() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice12bytesToWriteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:120
// index:0
// Public
// qint64 read(char *, qint64)
func (this *QIODevice) Read(data string, maxlen int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4readEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:121
// index:1
// Public
// QByteArray read(qint64)
func (this *QIODevice) Read_1(maxlen int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4readEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:122
// index:0
// Public
// QByteArray readAll()
func (this *QIODevice) ReadAll() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice7readAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:123
// index:0
// Public
// qint64 readLine(char *, qint64)
func (this *QIODevice) ReadLine(data string, maxlen int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice8readLineEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:124
// index:1
// Public
// QByteArray readLine(qint64)
func (this *QIODevice) ReadLine_1(maxlen int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice8readLineEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:125
// index:0
// Public virtual
// bool canReadLine()
func (this *QIODevice) CanReadLine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice11canReadLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:127
// index:0
// Public
// void startTransaction()
func (this *QIODevice) StartTransaction() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice16startTransactionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:128
// index:0
// Public
// void commitTransaction()
func (this *QIODevice) CommitTransaction() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice17commitTransactionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:129
// index:0
// Public
// void rollbackTransaction()
func (this *QIODevice) RollbackTransaction() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice19rollbackTransactionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:130
// index:0
// Public
// bool isTransactionStarted()
func (this *QIODevice) IsTransactionStarted() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice20isTransactionStartedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:132
// index:0
// Public
// qint64 write(const char *, qint64)
func (this *QIODevice) Write(data string, len int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice5writeEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:133
// index:1
// Public
// qint64 write(const char *)
func (this *QIODevice) Write_1(data string) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice5writeEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:134
// index:2
// Public inline
// qint64 write(const class QByteArray &)
func (this *QIODevice) Write_2(data *QByteArray) interface{} {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice5writeERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:137
// index:0
// Public
// qint64 peek(char *, qint64)
func (this *QIODevice) Peek(data string, maxlen int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4peekEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:138
// index:1
// Public
// QByteArray peek(qint64)
func (this *QIODevice) Peek_1(maxlen int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4peekEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:139
// index:0
// Public
// qint64 skip(qint64)
func (this *QIODevice) Skip(maxSize int64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice4skipEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maxSize)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:141
// index:0
// Public virtual
// bool waitForReadyRead(int)
func (this *QIODevice) WaitForReadyRead(msecs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice16waitForReadyReadEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:142
// index:0
// Public virtual
// bool waitForBytesWritten(int)
func (this *QIODevice) WaitForBytesWritten(msecs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice19waitForBytesWrittenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &msecs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:144
// index:0
// Public
// void ungetChar(char)
func (this *QIODevice) UngetChar(c byte) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice9ungetCharEc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:145
// index:0
// Public
// bool putChar(char)
func (this *QIODevice) PutChar(c byte) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice7putCharEc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:146
// index:0
// Public
// bool getChar(char *)
func (this *QIODevice) GetChar(c string) interface{} {
	var convArg0 = qtrt.CString(c)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice7getCharEPc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:148
// index:0
// Public
// QString errorString()
func (this *QIODevice) ErrorString() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QIODevice11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:152
// index:0
// Public
// void readyRead()
func (this *QIODevice) ReadyRead() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice9readyReadEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:153
// index:0
// Public
// void channelReadyRead(int)
func (this *QIODevice) ChannelReadyRead(channel int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice16channelReadyReadEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &channel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:154
// index:0
// Public
// void bytesWritten(qint64)
func (this *QIODevice) BytesWritten(bytes int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice12bytesWrittenEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &bytes)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:155
// index:0
// Public
// void channelBytesWritten(int, qint64)
func (this *QIODevice) ChannelBytesWritten(channel int, bytes int64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice19channelBytesWrittenEix", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &channel, &bytes)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:156
// index:0
// Public
// void aboutToClose()
func (this *QIODevice) AboutToClose() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice12aboutToCloseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:157
// index:0
// Public
// void readChannelFinished()
func (this *QIODevice) ReadChannelFinished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice19readChannelFinishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qiodevice.h:166
// index:0
// Protected pure virtual
// qint64 readData(char *, qint64)
func (this *QIODevice) ReadData(data string, maxlen int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice8readDataEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:167
// index:0
// Protected virtual
// qint64 readLineData(char *, qint64)
func (this *QIODevice) ReadLineData(data string, maxlen int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice12readLineDataEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &maxlen)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:168
// index:0
// Protected pure virtual
// qint64 writeData(const char *, qint64)
func (this *QIODevice) WriteData(data string, len int64) interface{} {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qiodevice.h:172
// index:0
// Protected
// void setErrorString(const class QString &)
func (this *QIODevice) SetErrorString(errorString *QString) {
	var convArg0 = errorString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QIODevice14setErrorStringERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
