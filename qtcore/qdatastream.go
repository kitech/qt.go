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

/*

 */
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

/*
Constructs a data stream that has no I/O device.

See also setDevice().
*/
func (*QDataStream) NewForInherit() *QDataStream {
	return NewQDataStream()
}
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

/*
Constructs a data stream that has no I/O device.

See also setDevice().
*/
func (*QDataStream) NewForInherit1(arg0 QIODevice_ITF /*777 QIODevice **/) *QDataStream {
	return NewQDataStream1(arg0)
}
func NewQDataStream1(arg0 QIODevice_ITF /*777 QIODevice **/) *QDataStream {
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

/*
Constructs a data stream that has no I/O device.

See also setDevice().
*/
func (*QDataStream) NewForInherit2(arg0 QByteArray_ITF /*777 QByteArray **/, flags int) *QDataStream {
	return NewQDataStream2(arg0, flags)
}
func NewQDataStream2(arg0 QByteArray_ITF /*777 QByteArray **/, flags int) *QDataStream {
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

/*
Constructs a data stream that has no I/O device.

See also setDevice().
*/
func (*QDataStream) NewForInherit3(arg0 QByteArray_ITF) *QDataStream {
	return NewQDataStream3(arg0)
}
func NewQDataStream3(arg0 QByteArray_ITF) *QDataStream {
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

/*

 */
func DeleteQDataStream(this *QDataStream) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qdatastream.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device() const

/*
Returns the I/O device currently set, or 0 if no device is currently set.

See also setDevice().
*/
func (this *QDataStream) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qdatastream.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)

/*
void QDataStream::setDevice(QIODevice *d)

Sets the I/O device to d, which can be 0 to unset to current I/O device.

See also device().
*/
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

/*

 */
func (this *QDataStream) UnsetDevice() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11unsetDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool atEnd() const

/*
Returns true if the I/O device has reached the end position (end of the stream or file) or if there is no I/O device set; otherwise returns false.

See also QIODevice::atEnd().
*/
func (this *QDataStream) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatastream.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] QDataStream::Status status() const

/*
Returns the status of the data stream.

See also Status, setStatus(), and resetStatus().
*/
func (this *QDataStream) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatus(QDataStream::Status)

/*
Sets the status of the data stream to the status given.

Subsequent calls to setStatus() are ignored until resetStatus() is called.

See also Status, status(), and resetStatus().
*/
func (this *QDataStream) SetStatus(status int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream9setStatusENS_6StatusE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), status)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetStatus()

/*
Resets the status of the data stream.

See also Status, status(), and setStatus().
*/
func (this *QDataStream) ResetStatus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11resetStatusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] QDataStream::FloatingPointPrecision floatingPointPrecision() const

/*
Returns the floating point precision of the data stream.

This function was introduced in  Qt 4.6.

See also FloatingPointPrecision and setFloatingPointPrecision().
*/
func (this *QDataStream) FloatingPointPrecision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream22floatingPointPrecisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFloatingPointPrecision(QDataStream::FloatingPointPrecision)

/*
Sets the floating point precision of the data stream to precision. If the floating point precision is DoublePrecision and the version of the data stream is Qt_4_6 or higher, all floating point numbers will be written and read with 64-bit precision. If the floating point precision is SinglePrecision and the version is Qt_4_6 or higher, all floating point numbers will be written and read with 32-bit precision.

For versions prior to Qt_4_6, the precision of floating point numbers in the data stream depends on the stream operator called.

The default is DoublePrecision.

Note that this property does not affect the serialization or deserialization of qfloat16 instances.

Warning: This property must be set to the same value on the object that writes and the object that reads the data stream.

This function was introduced in  Qt 4.6.

See also floatingPointPrecision().
*/
func (this *QDataStream) SetFloatingPointPrecision(precision int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream25setFloatingPointPrecisionENS_22FloatingPointPrecisionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), precision)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:142
// index:0
// Public Visibility=Default Availability=Available
// [4] QDataStream::ByteOrder byteOrder() const

/*
Returns the current byte order setting -- either BigEndian or LittleEndian.

See also setByteOrder().
*/
func (this *QDataStream) ByteOrder() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream9byteOrderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qdatastream.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setByteOrder(QDataStream::ByteOrder)

/*
Sets the serialization byte order to bo.

The bo parameter can be QDataStream::BigEndian or QDataStream::LittleEndian.

The default setting is big endian. We recommend leaving this setting unless you have special requirements.

See also byteOrder().
*/
func (this *QDataStream) SetByteOrder(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream12setByteOrderENS_9ByteOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:145
// index:0
// Public Visibility=Default Availability=Available
// [4] int version() const

/*
Returns the version number of the data serialization format.

See also setVersion() and Version.
*/
func (this *QDataStream) Version() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QDataStream7versionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVersion(int)

/*
Sets the version number of the data serialization format to v, a value of the Version enum.

You don't have to set a version if you are using the current version of Qt, but for your own custom binary formats we recommend that you do; see Versioning in the Detailed Description.

To accommodate new functionality, the datastream serialization format of some Qt classes has changed in some versions of Qt. If you want to read data that was created by an earlier version of Qt, or write data that can be read by a program that was compiled with an earlier version of Qt, use this function to modify the serialization format used by QDataStream.

The Version enum provides symbolic constants for the different versions of Qt. For example:


  QDataStream out(file);
  out.setVersion(QDataStream::Qt_4_0);



See also version() and Version.
*/
func (this *QDataStream) SetVersion(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream10setVersionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:149
// index:0
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(quint8 &)

/*

 */
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

/*

 */
func (this *QDataStream) Operator_right_shift1(i int16) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_right_shift2(i uint16) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_right_shift3(i int) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_right_shift4(i uint) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_right_shift5(i int64) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_right_shift6(i uint64) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamrsERy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:158
// index:7
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator>>(bool &)

/*

 */
func (this *QDataStream) Operator_right_shift7(i bool) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_right_shift8(f float32) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_right_shift9(f float64) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_right_shift10(str unsafe.Pointer /*555*/) *QDataStream {
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

/*

 */
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

/*

 */
func (this *QDataStream) Operator_left_shift1(i int16) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_left_shift2(i uint16) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_left_shift3(i int) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_left_shift4(i uint) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_left_shift5(i int64) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_left_shift6(i uint64) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStreamlsEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:173
// index:7
// Public Visibility=Default Availability=Available
// [32] QDataStream & operator<<(bool)

/*

 */
func (this *QDataStream) Operator_left_shift7(i bool) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_left_shift8(f float32) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_left_shift9(f float64) *QDataStream {
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

/*

 */
func (this *QDataStream) Operator_left_shift10(str string) *QDataStream {
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

/*
Reads the buffer s from the stream and returns a reference to the stream.

The buffer s is allocated using new []. Destroy it with the delete [] operator.

The l parameter is set to the length of the buffer. If the string read is empty, l is set to 0 and s is set to a null pointer.

The serialization format is a quint32 length specifier first, then l bytes of data.

See also readRawData() and writeBytes().
*/
func (this *QDataStream) ReadBytes(arg0 unsafe.Pointer /*555*/, len_ uint) *QDataStream {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream9readBytesERPcRj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, &len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:180
// index:0
// Public Visibility=Default Availability=Available
// [4] int readRawData(char *, int)

/*
Reads at most len bytes from the stream into s and returns the number of bytes read. If an error occurs, this function returns -1.

The buffer s must be preallocated. The data is not encoded.

See also readBytes(), QIODevice::read(), and writeRawData().
*/
func (this *QDataStream) ReadRawData(arg0 string, len_ int) int {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11readRawDataEPci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:182
// index:0
// Public Visibility=Default Availability=Available
// [32] QDataStream & writeBytes(const char *, uint)

/*
Writes the length specifier len and the buffer s to the stream and returns a reference to the stream.

The len is serialized as a quint32, followed by len bytes from s. Note that the data is not encoded.

See also writeRawData() and readBytes().
*/
func (this *QDataStream) WriteBytes(arg0 string, len_ uint) *QDataStream {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream10writeBytesEPKcj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQDataStreamFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQDataStream)
	return rv2
}

// /usr/include/qt/QtCore/qdatastream.h:183
// index:0
// Public Visibility=Default Availability=Available
// [4] int writeRawData(const char *, int)

/*
Writes len bytes from s to the stream. Returns the number of bytes actually written, or -1 on error. The data is not encoded.

See also writeBytes(), QIODevice::write(), and readRawData().
*/
func (this *QDataStream) WriteRawData(arg0 string, len_ int) int {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream12writeRawDataEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:185
// index:0
// Public Visibility=Default Availability=Available
// [4] int skipRawData(int)

/*
Skips len bytes from the device. Returns the number of bytes actually skipped, or -1 on error.

This is equivalent to calling readRawData() on a buffer of length len and ignoring the buffer.

This function was introduced in  Qt 4.1.

See also QIODevice::seek().
*/
func (this *QDataStream) SkipRawData(len_ int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream11skipRawDataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), len_)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qdatastream.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startTransaction()

/*
Starts a new read transaction on the stream.

Defines a restorable point within the sequence of read operations. For sequential devices, read data will be duplicated internally to allow recovery in case of incomplete reads. For random-access devices, this function saves the current position of the stream. Call commitTransaction(), rollbackTransaction(), or abortTransaction() to finish the current transaction.

Once a transaction is started, subsequent calls to this function will make the transaction recursive. Inner transactions act as agents of the outermost transaction (i.e., report the status of read operations to the outermost transaction, which can restore the position of the stream).

Note: Restoring to the point of the nested startTransaction() call is not supported.

When an error occurs during a transaction (including an inner transaction failing), reading from the data stream is suspended (all subsequent read operations return empty/zero values) and subsequent inner transactions are forced to fail. Starting a new outermost transaction recovers from this state. This behavior makes it unnecessary to error-check every read operation separately.

This function was introduced in  Qt 5.7.

See also commitTransaction(), rollbackTransaction(), and abortTransaction().
*/
func (this *QDataStream) StartTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream16startTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:188
// index:0
// Public Visibility=Default Availability=Available
// [1] bool commitTransaction()

/*
Completes a read transaction. Returns true if no read errors have occurred during the transaction; otherwise returns false.

If called on an inner transaction, committing will be postponed until the outermost commitTransaction(), rollbackTransaction(), or abortTransaction() call occurs.

Otherwise, if the stream status indicates reading past the end of the data, this function restores the stream data to the point of the startTransaction() call. When this situation occurs, you need to wait for more data to arrive, after which you start a new transaction. If the data stream has read corrupt data or any of the inner transactions was aborted, this function aborts the transaction.

This function was introduced in  Qt 5.7.

See also startTransaction(), rollbackTransaction(), and abortTransaction().
*/
func (this *QDataStream) CommitTransaction() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream17commitTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qdatastream.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rollbackTransaction()

/*
Reverts a read transaction.

This function is commonly used to rollback the transaction when an incomplete read was detected prior to committing the transaction.

If called on an inner transaction, reverting is delegated to the outermost transaction, and subsequently started inner transactions are forced to fail.

For the outermost transaction, restores the stream data to the point of the startTransaction() call. If the data stream has read corrupt data or any of the inner transactions was aborted, this function aborts the transaction.

If the preceding stream operations were successful, sets the status of the data stream to

ConstantDescription
ReadPastEnd.


This function was introduced in  Qt 5.7.

See also startTransaction(), commitTransaction(), and abortTransaction().
*/
func (this *QDataStream) RollbackTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream19rollbackTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qdatastream.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void abortTransaction()

/*
Aborts a read transaction.

This function is commonly used to discard the transaction after higher-level protocol errors or loss of stream synchronization.

If called on an inner transaction, aborting is delegated to the outermost transaction, and subsequently started inner transactions are forced to fail.

For the outermost transaction, discards the restoration point and any internally duplicated data of the stream. Will not affect the current read position of the stream.

Sets the status of the data stream to

ConstantDescription
ReadCorruptData.


This function was introduced in  Qt 5.7.

See also startTransaction(), commitTransaction(), and rollbackTransaction().
*/
func (this *QDataStream) AbortTransaction() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QDataStream16abortTransactionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*
This enum provides symbolic synonyms for the data serialization format version numbers.



See also setVersion() and version().

*/
type QDataStream__Version = int

//
const QDataStream__Qt_1_0 QDataStream__Version = 1

//
const QDataStream__Qt_2_0 QDataStream__Version = 2

//
const QDataStream__Qt_2_1 QDataStream__Version = 3

//
const QDataStream__Qt_3_0 QDataStream__Version = 4

//
const QDataStream__Qt_3_1 QDataStream__Version = 5

//
const QDataStream__Qt_3_3 QDataStream__Version = 6

//
const QDataStream__Qt_4_0 QDataStream__Version = 7

//
const QDataStream__Qt_4_1 QDataStream__Version = 7

//
const QDataStream__Qt_4_2 QDataStream__Version = 8

//
const QDataStream__Qt_4_3 QDataStream__Version = 9

//
const QDataStream__Qt_4_4 QDataStream__Version = 10

//
const QDataStream__Qt_4_5 QDataStream__Version = 11

//
const QDataStream__Qt_4_6 QDataStream__Version = 12

//
const QDataStream__Qt_4_7 QDataStream__Version = 12

//
const QDataStream__Qt_4_8 QDataStream__Version = 12

//
const QDataStream__Qt_4_9 QDataStream__Version = 12

//
const QDataStream__Qt_5_0 QDataStream__Version = 13

//
const QDataStream__Qt_5_1 QDataStream__Version = 14

//
const QDataStream__Qt_5_2 QDataStream__Version = 15

//
const QDataStream__Qt_5_3 QDataStream__Version = 15

//
const QDataStream__Qt_5_4 QDataStream__Version = 16

//
const QDataStream__Qt_5_5 QDataStream__Version = 16

//
const QDataStream__Qt_5_6 QDataStream__Version = 17

//
const QDataStream__Qt_5_7 QDataStream__Version = 17

//
const QDataStream__Qt_5_8 QDataStream__Version = 17

//
const QDataStream__Qt_5_9 QDataStream__Version = 17

//
const QDataStream__Qt_5_10 QDataStream__Version = 17

//
const QDataStream__Qt_DefaultCompiledVersion QDataStream__Version = 17

func (this *QDataStream) VersionItemName(val int) string {
	switch val {
	case QDataStream__Qt_1_0: // 1
		return "Qt_1_0"
	case QDataStream__Qt_2_0: // 2
		return "Qt_2_0"
	case QDataStream__Qt_2_1: // 3
		return "Qt_2_1"
	case QDataStream__Qt_3_0: // 4
		return "Qt_3_0"
	case QDataStream__Qt_3_1: // 5
		return "Qt_3_1"
	case QDataStream__Qt_3_3: // 6
		return "Qt_3_3"
	case QDataStream__Qt_4_0: // 7
		return "Qt_4_0,Qt_4_1"
		// case QDataStream__Qt_4_1: // 7
		// return ""
	case QDataStream__Qt_4_2: // 8
		return "Qt_4_2"
	case QDataStream__Qt_4_3: // 9
		return "Qt_4_3"
	case QDataStream__Qt_4_4: // 10
		return "Qt_4_4"
	case QDataStream__Qt_4_5: // 11
		return "Qt_4_5"
	case QDataStream__Qt_4_6: // 12
		return "Qt_4_6,Qt_4_7,Qt_4_8,Qt_4_9"
		// case QDataStream__Qt_4_7: // 12
		// return ""
		// case QDataStream__Qt_4_8: // 12
		// return ""
		// case QDataStream__Qt_4_9: // 12
		// return ""
	case QDataStream__Qt_5_0: // 13
		return "Qt_5_0"
	case QDataStream__Qt_5_1: // 14
		return "Qt_5_1"
	case QDataStream__Qt_5_2: // 15
		return "Qt_5_2,Qt_5_3"
		// case QDataStream__Qt_5_3: // 15
		// return ""
	case QDataStream__Qt_5_4: // 16
		return "Qt_5_4,Qt_5_5"
		// case QDataStream__Qt_5_5: // 16
		// return ""
	case QDataStream__Qt_5_6: // 17
		return "Qt_5_6,Qt_5_7,Qt_5_8,Qt_5_9,Qt_5_10,Qt_DefaultCompiledVersion"
		// case QDataStream__Qt_5_7: // 17
		// return ""
		// case QDataStream__Qt_5_8: // 17
		// return ""
		// case QDataStream__Qt_5_9: // 17
		// return ""
		// case QDataStream__Qt_5_10: // 17
		// return ""
		// case QDataStream__Qt_DefaultCompiledVersion: // 17
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QDataStream_VersionItemName(val int) string {
	var nilthis *QDataStream
	return nilthis.VersionItemName(val)
}

/*
The byte order used for reading/writing the data.

QDataStream::BigEndianQSysInfo::BigEndianMost significant byte first (the default)
QDataStream::LittleEndianQSysInfo::LittleEndianLeast significant byte first

*/
type QDataStream__ByteOrder = int

//
const QDataStream__BigEndian QDataStream__ByteOrder = 0

//
const QDataStream__LittleEndian QDataStream__ByteOrder = 1

func (this *QDataStream) ByteOrderItemName(val int) string {
	switch val {
	case QDataStream__BigEndian: // 0
		return "BigEndian"
	case QDataStream__LittleEndian: // 1
		return "LittleEndian"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QDataStream_ByteOrderItemName(val int) string {
	var nilthis *QDataStream
	return nilthis.ByteOrderItemName(val)
}

/*
This enum describes the current status of the data stream.


*/
type QDataStream__Status = int

// The data stream is operating normally.
const QDataStream__Ok QDataStream__Status = 0

// The data stream has read past the end of the data in the underlying device.
const QDataStream__ReadPastEnd QDataStream__Status = 1

// The data stream has read corrupt data.
const QDataStream__ReadCorruptData QDataStream__Status = 2

// The data stream cannot write to the underlying device.
const QDataStream__WriteFailed QDataStream__Status = 3

func (this *QDataStream) StatusItemName(val int) string {
	switch val {
	case QDataStream__Ok: // 0
		return "Ok"
	case QDataStream__ReadPastEnd: // 1
		return "ReadPastEnd"
	case QDataStream__ReadCorruptData: // 2
		return "ReadCorruptData"
	case QDataStream__WriteFailed: // 3
		return "WriteFailed"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QDataStream_StatusItemName(val int) string {
	var nilthis *QDataStream
	return nilthis.StatusItemName(val)
}

/*
The precision of floating point numbers used for reading/writing the data. This will only have an effect if the version of the data stream is Qt_4_6 or higher.

Warning: The floating point precision must be set to the same value on the object that writes and the object that reads the data stream.



See also setFloatingPointPrecision() and floatingPointPrecision().

*/
type QDataStream__FloatingPointPrecision = int

//
const QDataStream__SinglePrecision QDataStream__FloatingPointPrecision = 0

//
const QDataStream__DoublePrecision QDataStream__FloatingPointPrecision = 1

func (this *QDataStream) FloatingPointPrecisionItemName(val int) string {
	switch val {
	case QDataStream__SinglePrecision: // 0
		return "SinglePrecision"
	case QDataStream__DoublePrecision: // 1
		return "DoublePrecision"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QDataStream_FloatingPointPrecisionItemName(val int) string {
	var nilthis *QDataStream
	return nilthis.FloatingPointPrecisionItemName(val)
}

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
