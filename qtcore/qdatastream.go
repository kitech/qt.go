package qtcore

// /usr/include/qt/QtCore/qdatastream.h
// #include <qdatastream.h>
// #include <QtCore>

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
type QDataStream struct {
	*qtrt.CObject
}

func (this *QDataStream) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QDataStream) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQDataStreamFromPointer(cthis unsafe.Pointer) *QDataStream {
	return &QDataStream{&qtrt.CObject{cthis}}
}
func (*QDataStream) NewFromPointer(cthis unsafe.Pointer) *QDataStream {
	return NewQDataStreamFromPointer(cthis)
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
// void QDataStream(QIODevice *)
func NewQDataStream_1(arg0 *QIODevice /*777 QIODevice **/) *QDataStream {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStreamC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:126
// index:2
// Public
// void QDataStream(const QByteArray &)
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
func (this *QDataStream) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream6deviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:130
// index:0
// Public
// void setDevice(QIODevice *)
func (this *QDataStream) SetDevice(arg0 *QIODevice /*777 QIODevice **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream9setDeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
func (this *QDataStream) AtEnd() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdatastream.h:135
// index:0
// Public
// QDataStream::Status status()
func (this *QDataStream) Status() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream6statusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:136
// index:0
// Public
// void setStatus(QDataStream::Status)
func (this *QDataStream) SetStatus(status int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream9setStatusENS_6StatusE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), status)
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
func (this *QDataStream) FloatingPointPrecision() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream22floatingPointPrecisionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:140
// index:0
// Public
// void setFloatingPointPrecision(QDataStream::FloatingPointPrecision)
func (this *QDataStream) SetFloatingPointPrecision(precision int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream25setFloatingPointPrecisionENS_22FloatingPointPrecisionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:142
// index:0
// Public
// QDataStream::ByteOrder byteOrder()
func (this *QDataStream) ByteOrder() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream9byteOrderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:143
// index:0
// Public
// void setByteOrder(QDataStream::ByteOrder)
func (this *QDataStream) SetByteOrder(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream12setByteOrderENS_9ByteOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:145
// index:0
// Public
// int version()
func (this *QDataStream) Version() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QDataStream7versionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatastream.h:146
// index:0
// Public
// void setVersion(int)
func (this *QDataStream) SetVersion(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream10setVersionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:179
// index:0
// Public
// QDataStream & readBytes(char *&, uint &)
func (this *QDataStream) ReadBytes(arg0 unsafe.Pointer /*555*/, len uint) *QDataStream {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream9readBytesERPcRj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, &len)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:180
// index:0
// Public
// int readRawData(char *, int)
func (this *QDataStream) ReadRawData(arg0 string, len int) int {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream11readRawDataEPci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatastream.h:182
// index:0
// Public
// QDataStream & writeBytes(const char *, uint)
func (this *QDataStream) WriteBytes(arg0 string, len uint) *QDataStream {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream10writeBytesEPKcj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:183
// index:0
// Public
// int writeRawData(const char *, int)
func (this *QDataStream) WriteRawData(arg0 string, len int) int {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream12writeRawDataEPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qdatastream.h:185
// index:0
// Public
// int skipRawData(int)
func (this *QDataStream) SkipRawData(len int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream11skipRawDataEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
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
func (this *QDataStream) CommitTransaction() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QDataStream17commitTransactionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
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

type QDataStream__Version = int

const QDataStream__Qt_1_0 QDataStream__Version = 1
const QDataStream__Qt_2_0 QDataStream__Version = 2
const QDataStream__Qt_2_1 QDataStream__Version = 3
const QDataStream__Qt_3_0 QDataStream__Version = 4
const QDataStream__Qt_3_1 QDataStream__Version = 5
const QDataStream__Qt_3_3 QDataStream__Version = 6
const QDataStream__Qt_4_0 QDataStream__Version = 7
const QDataStream__Qt_4_1 QDataStream__Version = 7
const QDataStream__Qt_4_2 QDataStream__Version = 8
const QDataStream__Qt_4_3 QDataStream__Version = 9
const QDataStream__Qt_4_4 QDataStream__Version = 10
const QDataStream__Qt_4_5 QDataStream__Version = 11
const QDataStream__Qt_4_6 QDataStream__Version = 12
const QDataStream__Qt_4_7 QDataStream__Version = 12
const QDataStream__Qt_4_8 QDataStream__Version = 12
const QDataStream__Qt_4_9 QDataStream__Version = 12
const QDataStream__Qt_5_0 QDataStream__Version = 13
const QDataStream__Qt_5_1 QDataStream__Version = 14
const QDataStream__Qt_5_2 QDataStream__Version = 15
const QDataStream__Qt_5_3 QDataStream__Version = 15
const QDataStream__Qt_5_4 QDataStream__Version = 16
const QDataStream__Qt_5_5 QDataStream__Version = 16
const QDataStream__Qt_5_6 QDataStream__Version = 17
const QDataStream__Qt_5_7 QDataStream__Version = 17
const QDataStream__Qt_5_8 QDataStream__Version = 17
const QDataStream__Qt_5_9 QDataStream__Version = 17
const QDataStream__Qt_5_10 QDataStream__Version = 17
const QDataStream__Qt_DefaultCompiledVersion QDataStream__Version = 17

type QDataStream__ByteOrder = int

const QDataStream__BigEndian QDataStream__ByteOrder = 0
const QDataStream__LittleEndian QDataStream__ByteOrder = 1

type QDataStream__Status = int

const QDataStream__Ok QDataStream__Status = 0
const QDataStream__ReadPastEnd QDataStream__Status = 1
const QDataStream__ReadCorruptData QDataStream__Status = 2
const QDataStream__WriteFailed QDataStream__Status = 3

type QDataStream__FloatingPointPrecision = int

const QDataStream__SinglePrecision QDataStream__FloatingPointPrecision = 0
const QDataStream__DoublePrecision QDataStream__FloatingPointPrecision = 1

//  body block end
