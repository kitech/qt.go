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
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQDataStreamFromPointer(cthis unsafe.Pointer) *QDataStream {
	return &QDataStream{&qtrt.CObject{cthis}}
}
func (*QDataStream) NewFromPointer(cthis unsafe.Pointer) *QDataStream {
	return NewQDataStreamFromPointer(cthis)
}

// /usr/include/qt/QtCore/qdatastream.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDataStream()
func NewQDataStream() *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDataStream)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:124
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDataStream(QIODevice *)
func NewQDataStream_1(arg0 *QIODevice /*777 QIODevice **/) *QDataStream {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamC2EP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDataStream)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:125
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDataStream(QByteArray *, QIODevice::OpenMode)
func NewQDataStream_2(arg0 *QByteArray /*777 QByteArray **/, flags int) *QDataStream {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamC2EP10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDataStream)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:126
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDataStream(const QByteArray &)
func NewQDataStream_3(arg0 *QByteArray) *QDataStream {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDataStream)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QDataStream()
func DeleteQDataStream(this *QDataStream) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qdatastream.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QDataStream) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QDataStream) SetDevice(arg0 *QIODevice /*777 QIODevice **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetDevice()
func (this *QDataStream) UnsetDevice() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11unsetDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atEnd()
func (this *QDataStream) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdatastream.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] QDataStream::Status status()
func (this *QDataStream) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatus(enum QDataStream::Status)
func (this *QDataStream) SetStatus(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream9setStatusENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetStatus()
func (this *QDataStream) ResetStatus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11resetStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] QDataStream::FloatingPointPrecision floatingPointPrecision()
func (this *QDataStream) FloatingPointPrecision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream22floatingPointPrecisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFloatingPointPrecision(enum QDataStream::FloatingPointPrecision)
func (this *QDataStream) SetFloatingPointPrecision(precision int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream25setFloatingPointPrecisionENS_22FloatingPointPrecisionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), precision)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:142
// index:0
// Public Visibility=Default Availability=Available
// [4] QDataStream::ByteOrder byteOrder()
func (this *QDataStream) ByteOrder() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream9byteOrderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setByteOrder(enum QDataStream::ByteOrder)
func (this *QDataStream) SetByteOrder(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream12setByteOrderENS_9ByteOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:145
// index:0
// Public Visibility=Default Availability=Available
// [4] int version()
func (this *QDataStream) Version() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream7versionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVersion(int)
func (this *QDataStream) SetVersion(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream10setVersionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:179
// index:0
// Public Visibility=Default Availability=Available
// [32] QDataStream & readBytes(char *&, uint &)
func (this *QDataStream) ReadBytes(arg0 unsafe.Pointer /*555*/, len uint) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream9readBytesERPcRj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, &len)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:180
// index:0
// Public Visibility=Default Availability=Available
// [4] int readRawData(char *, int)
func (this *QDataStream) ReadRawData(arg0 string, len int) int {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11readRawDataEPci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:182
// index:0
// Public Visibility=Default Availability=Available
// [32] QDataStream & writeBytes(const char *, uint)
func (this *QDataStream) WriteBytes(arg0 string, len uint) *QDataStream {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream10writeBytesEPKcj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:183
// index:0
// Public Visibility=Default Availability=Available
// [4] int writeRawData(const char *, int)
func (this *QDataStream) WriteRawData(arg0 string, len int) int {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream12writeRawDataEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:185
// index:0
// Public Visibility=Default Availability=Available
// [4] int skipRawData(int)
func (this *QDataStream) SkipRawData(len int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11skipRawDataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startTransaction()
func (this *QDataStream) StartTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream16startTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:188
// index:0
// Public Visibility=Default Availability=Available
// [1] bool commitTransaction()
func (this *QDataStream) CommitTransaction() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream17commitTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qdatastream.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rollbackTransaction()
func (this *QDataStream) RollbackTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream19rollbackTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void abortTransaction()
func (this *QDataStream) AbortTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream16abortTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
