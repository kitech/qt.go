//  header block begin
// /usr/include/qt/QtCore/qdatastream.h
// #include <qdatastream.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QDataStream struct {
	*qtrt.CObject
}

func (this *QDataStream) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQDataStreamFromPointer(cthis unsafe.Pointer) *QDataStream {
	return &QDataStream{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qdatastream.h:123
// index:0
// Public
// void QDataStream()
func NewQDataStream() *QDataStream {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStreamC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:124
// index:1
// Public
// void QDataStream(class QIODevice *)
func NewQDataStream_1(arg0 unsafe.Pointer) *QDataStream {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStreamC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:126
// index:2
// Public
// void QDataStream(const class QByteArray &)
func NewQDataStream_2(arg0 *QByteArray) *QDataStream {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStreamC2ERK10QByteArray", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:127
// index:0
// Public
// void ~QDataStream()
func DeleteQDataStream(*QDataStream) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStreamD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:129
// index:0
// Public
// QIODevice * device()
func (this *QDataStream) Device() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:130
// index:0
// Public
// void setDevice(class QIODevice *)
func (this *QDataStream) SetDevice(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:131
// index:0
// Public
// void unsetDevice()
func (this *QDataStream) UnsetDevice() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream11unsetDeviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:133
// index:0
// Public
// bool atEnd()
func (this *QDataStream) AtEnd() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:135
// index:0
// Public
// QDataStream::Status status()
func (this *QDataStream) Status() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream6statusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:136
// index:0
// Public
// void setStatus(enum QDataStream::Status)
func (this *QDataStream) SetStatus(status int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream9setStatusENS_6StatusE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &status)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:137
// index:0
// Public
// void resetStatus()
func (this *QDataStream) ResetStatus() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream11resetStatusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:139
// index:0
// Public
// QDataStream::FloatingPointPrecision floatingPointPrecision()
func (this *QDataStream) FloatingPointPrecision() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream22floatingPointPrecisionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:140
// index:0
// Public
// void setFloatingPointPrecision(enum QDataStream::FloatingPointPrecision)
func (this *QDataStream) SetFloatingPointPrecision(precision int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream25setFloatingPointPrecisionENS_22FloatingPointPrecisionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:142
// index:0
// Public
// QDataStream::ByteOrder byteOrder()
func (this *QDataStream) ByteOrder() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream9byteOrderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:143
// index:0
// Public
// void setByteOrder(enum QDataStream::ByteOrder)
func (this *QDataStream) SetByteOrder(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream12setByteOrderENS_9ByteOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:145
// index:0
// Public
// int version()
func (this *QDataStream) Version() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream7versionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:146
// index:0
// Public
// void setVersion(int)
func (this *QDataStream) SetVersion(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream10setVersionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:179
// index:0
// Public
// QDataStream & readBytes(char *&, uint &)
func (this *QDataStream) ReadBytes(arg0 unsafe.Pointer, len uint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream9readBytesERPcRj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:180
// index:0
// Public
// int readRawData(char *, int)
func (this *QDataStream) ReadRawData(arg0 string, len int) interface{} {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream11readRawDataEPci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:182
// index:0
// Public
// QDataStream & writeBytes(const char *, uint)
func (this *QDataStream) WriteBytes(arg0 string, len uint) interface{} {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream10writeBytesEPKcj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:183
// index:0
// Public
// int writeRawData(const char *, int)
func (this *QDataStream) WriteRawData(arg0 string, len int) interface{} {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream12writeRawDataEPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:185
// index:0
// Public
// int skipRawData(int)
func (this *QDataStream) SkipRawData(len int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream11skipRawDataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &len)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:187
// index:0
// Public
// void startTransaction()
func (this *QDataStream) StartTransaction() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream16startTransactionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:188
// index:0
// Public
// bool commitTransaction()
func (this *QDataStream) CommitTransaction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream17commitTransactionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qdatastream.h:189
// index:0
// Public
// void rollbackTransaction()
func (this *QDataStream) RollbackTransaction() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream19rollbackTransactionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:190
// index:0
// Public
// void abortTransaction()
func (this *QDataStream) AbortTransaction() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream16abortTransactionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
