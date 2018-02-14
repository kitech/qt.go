package qtcore

// /usr/include/qt/QtCore/qdatastream.h
// #include <qdatastream.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QDataStream struct {
	*qtrt.CObject
}
type QDataStream_ITF interface {
	QDataStream_PTR() *QDataStream
}

func (ptr *QDataStream) QDataStream_PTR() *QDataStream { return ptr }

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
	qtrt.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDataStream)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:124
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDataStream(QIODevice *)
func NewQDataStream_1(arg0 QIODevice_ITF /*777 QIODevice **/) *QDataStream {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QIODevice_PTR() != nil {
		convArg0 = arg0.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamC2EP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDataStream)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:125
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QDataStream(QByteArray *, QIODevice::OpenMode)
func NewQDataStream_2(arg0 QByteArray_ITF /*777 QByteArray **/, flags int) *QDataStream {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamC2EP10QByteArray6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, convArg0, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQDataStream)
	return gothis
}

// /usr/include/qt/QtCore/qdatastream.h:126
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QDataStream(const QByteArray &)
func NewQDataStream_3(arg0 QByteArray_ITF) *QDataStream {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qdatastream.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QDataStream) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qdatastream.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QDataStream) SetDevice(arg0 QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QIODevice_PTR() != nil {
		convArg0 = arg0.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetDevice()
func (this *QDataStream) UnsetDevice() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11unsetDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atEnd()
func (this *QDataStream) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatastream.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] QDataStream::Status status()
func (this *QDataStream) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatus(enum QDataStream::Status)
func (this *QDataStream) SetStatus(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream9setStatusENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetStatus()
func (this *QDataStream) ResetStatus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11resetStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] QDataStream::FloatingPointPrecision floatingPointPrecision()
func (this *QDataStream) FloatingPointPrecision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream22floatingPointPrecisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFloatingPointPrecision(enum QDataStream::FloatingPointPrecision)
func (this *QDataStream) SetFloatingPointPrecision(precision int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream25setFloatingPointPrecisionENS_22FloatingPointPrecisionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), precision)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:142
// index:0
// Public Visibility=Default Availability=Available
// [4] QDataStream::ByteOrder byteOrder()
func (this *QDataStream) ByteOrder() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream9byteOrderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setByteOrder(enum QDataStream::ByteOrder)
func (this *QDataStream) SetByteOrder(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream12setByteOrderENS_9ByteOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:145
// index:0
// Public Visibility=Default Availability=Available
// [4] int version()
func (this *QDataStream) Version() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream7versionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVersion(int)
func (this *QDataStream) SetVersion(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream10setVersionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:149
// index:0
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(quint8 &)
func (this *QDataStream) Operator_right_shift(i byte) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:150
// index:1
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(qint16 &)
func (this *QDataStream) Operator_right_shift_1(i int16) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERs", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:151
// index:2
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(quint16 &)
func (this *QDataStream) Operator_right_shift_2(i uint16) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:152
// index:3
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(qint32 &)
func (this *QDataStream) Operator_right_shift_3(i int) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:153
// index:4
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(quint32 &)
func (this *QDataStream) Operator_right_shift_4(i uint) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:154
// index:5
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(qint64 &)
func (this *QDataStream) Operator_right_shift_5(i int64) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:155
// index:6
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(quint64 &)
func (this *QDataStream) Operator_right_shift_6(i uint64) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:158
// index:7
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(_Bool &)
func (this *QDataStream) Operator_right_shift_7(i bool) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:160
// index:8
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(float &)
func (this *QDataStream) Operator_right_shift_8(f float32) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:161
// index:9
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(double &)
func (this *QDataStream) Operator_right_shift_9(f float64) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:162
// index:10
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(char *&)
func (this *QDataStream) Operator_right_shift_10(str unsafe.Pointer /*555*/) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERPc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), str)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:165
// index:0
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(quint8)
func (this *QDataStream) Operator_left_shift(i byte) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEh", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:166
// index:1
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(qint16)
func (this *QDataStream) Operator_left_shift_1(i int16) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEs", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:167
// index:2
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(quint16)
func (this *QDataStream) Operator_left_shift_2(i uint16) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:168
// index:3
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(qint32)
func (this *QDataStream) Operator_left_shift_3(i int) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:169
// index:4
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(quint32)
func (this *QDataStream) Operator_left_shift_4(i uint) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:170
// index:5
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(qint64)
func (this *QDataStream) Operator_left_shift_5(i int64) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:171
// index:6
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(quint64)
func (this *QDataStream) Operator_left_shift_6(i uint64) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:173
// index:7
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(_Bool)
func (this *QDataStream) Operator_left_shift_7(i bool) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:175
// index:8
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(float)
func (this *QDataStream) Operator_left_shift_8(f float32) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:176
// index:9
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(double)
func (this *QDataStream) Operator_left_shift_9(f float64) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), f)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:177
// index:10
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(const char *)
func (this *QDataStream) Operator_left_shift_10(str string) *QDataStream {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:179
// index:0
// Public Visibility=Default Availability=Available
// [32] QDataStream & readBytes(char *&, uint &)
func (this *QDataStream) ReadBytes(arg0 unsafe.Pointer /*555*/, len uint) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream9readBytesERPcRj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, &len)
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:185
// index:0
// Public Visibility=Default Availability=Available
// [4] int skipRawData(int)
func (this *QDataStream) SkipRawData(len int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11skipRawDataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), len)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startTransaction()
func (this *QDataStream) StartTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream16startTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:188
// index:0
// Public Visibility=Default Availability=Available
// [1] bool commitTransaction()
func (this *QDataStream) CommitTransaction() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream17commitTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatastream.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rollbackTransaction()
func (this *QDataStream) RollbackTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream19rollbackTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void abortTransaction()
func (this *QDataStream) AbortTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream16abortTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
