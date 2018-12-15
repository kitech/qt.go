package qtcore

// /usr/include/qt/QtCore/qbuffer.h
// #include <qbuffer.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 58
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void connectNotify(const QMetaMethod &)
func (this *QBuffer) InheritConnectNotify(f func(arg0 *QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "connectNotify", f)
}

// void disconnectNotify(const QMetaMethod &)
func (this *QBuffer) InheritDisconnectNotify(f func(arg0 *QMetaMethod) /*void*/) {
	qtrt.SetAllInheritCallback(this, "disconnectNotify", f)
}

// long long readData(char *, qint64)
func (this *QBuffer) InheritReadData(f func(data string, maxlen int64) int64) {
	qtrt.SetAllInheritCallback(this, "readData", f)
}

// long long writeData(const char *, qint64)
func (this *QBuffer) InheritWriteData(f func(data string, len_ int64) int64) {
	qtrt.SetAllInheritCallback(this, "writeData", f)
}

/*

 */
type QBuffer struct {
	*QIODevice
}
type QBuffer_ITF interface {
	QIODevice_ITF
	QBuffer_PTR() *QBuffer
}

func (ptr *QBuffer) QBuffer_PTR() *QBuffer { return ptr }

func (this *QBuffer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QIODevice.GetCthis()
	}
}
func (this *QBuffer) SetCthis(cthis unsafe.Pointer) {
	this.QIODevice = NewQIODeviceFromPointer(cthis)
}
func NewQBufferFromPointer(cthis unsafe.Pointer) *QBuffer {
	bcthis0 := NewQIODeviceFromPointer(cthis)
	return &QBuffer{bcthis0}
}
func (*QBuffer) NewFromPointer(cthis unsafe.Pointer) *QBuffer {
	return NewQBufferFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbuffer.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QBuffer) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QBuffer10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qbuffer.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBuffer(QObject *)

/*
Constructs an empty buffer with the given parent. You can call setData() to fill the buffer with data, or you can open it in write mode and use write().

See also open().
*/
func (*QBuffer) NewForInherit(parent QObject_ITF /*777 QObject **/) *QBuffer {
	return NewQBuffer(parent)
}
func NewQBuffer(parent QObject_ITF /*777 QObject **/) *QBuffer {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBufferC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QBuffer")
	return gothis
}

// /usr/include/qt/QtCore/qbuffer.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBuffer(QObject *)

/*
Constructs an empty buffer with the given parent. You can call setData() to fill the buffer with data, or you can open it in write mode and use write().

See also open().
*/
func (*QBuffer) NewForInheritp() *QBuffer {
	return NewQBufferp()
}
func NewQBufferp() *QBuffer {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBufferC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QBuffer")
	return gothis
}

// /usr/include/qt/QtCore/qbuffer.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QBuffer(QByteArray *, QObject *)

/*
Constructs an empty buffer with the given parent. You can call setData() to fill the buffer with data, or you can open it in write mode and use write().

See also open().
*/
func (*QBuffer) NewForInherit1(buf QByteArray_ITF /*777 QByteArray **/, parent QObject_ITF /*777 QObject **/) *QBuffer {
	return NewQBuffer1(buf, parent)
}
func NewQBuffer1(buf QByteArray_ITF /*777 QByteArray **/, parent QObject_ITF /*777 QObject **/) *QBuffer {
	var convArg0 unsafe.Pointer
	if buf != nil && buf.QByteArray_PTR() != nil {
		convArg0 = buf.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBufferC2EP10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QBuffer")
	return gothis
}

// /usr/include/qt/QtCore/qbuffer.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QBuffer(QByteArray *, QObject *)

/*
Constructs an empty buffer with the given parent. You can call setData() to fill the buffer with data, or you can open it in write mode and use write().

See also open().
*/
func (*QBuffer) NewForInherit1p(buf QByteArray_ITF /*777 QByteArray **/) *QBuffer {
	return NewQBuffer1p(buf)
}
func NewQBuffer1p(buf QByteArray_ITF /*777 QByteArray **/) *QBuffer {
	var convArg0 unsafe.Pointer
	if buf != nil && buf.QByteArray_PTR() != nil {
		convArg0 = buf.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBufferC2EP10QByteArrayP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBufferFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QBuffer")
	return gothis
}

// /usr/include/qt/QtCore/qbuffer.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QBuffer()

/*

 */
func DeleteQBuffer(this *QBuffer) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBufferD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qbuffer.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray & buffer()

/*
Returns a reference to the QBuffer's internal buffer. You can use it to modify the QByteArray behind the QBuffer's back.

See also setBuffer() and data().
*/
func (this *QBuffer) Buffer() *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer6bufferEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:69
// index:1
// Public Visibility=Default Availability=Available
// [8] const QByteArray & buffer() const

/*
Returns a reference to the QBuffer's internal buffer. You can use it to modify the QByteArray behind the QBuffer's back.

See also setBuffer() and data().
*/
func (this *QBuffer) Buffer1() *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QBuffer6bufferEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBuffer(QByteArray *)

/*
Makes QBuffer uses the QByteArray pointed to by byteArray as its internal buffer. The caller is responsible for ensuring that byteArray remains valid until the QBuffer is destroyed, or until setBuffer() is called to change the buffer. QBuffer doesn't take ownership of the QByteArray.

Does nothing if isOpen() is true.

If you open the buffer in write-only mode or read-write mode and write something into the QBuffer, byteArray will be modified.

Example:


      QByteArray byteArray("abc");
      QBuffer buffer;
      buffer.setBuffer(&byteArray);
      buffer.open(QIODevice::WriteOnly);
      buffer.seek(3);
      buffer.write("def", 3);
      buffer.close();
      // byteArray == "abcdef"



If byteArray is 0, the buffer creates its own internal QByteArray to work on. This byte array is initially empty.

See also buffer(), setData(), and open().
*/
func (this *QBuffer) SetBuffer(a QByteArray_ITF /*777 QByteArray **/) {
	var convArg0 unsafe.Pointer
	if a != nil && a.QByteArray_PTR() != nil {
		convArg0 = a.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer9setBufferEP10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setData(const QByteArray &)

/*
Sets the contents of the internal buffer to be data. This is the same as assigning data to buffer().

Does nothing if isOpen() is true.

See also data() and setBuffer().
*/
func (this *QBuffer) SetData(data QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer7setDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:73
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setData(const char *, int)

/*
Sets the contents of the internal buffer to be data. This is the same as assigning data to buffer().

Does nothing if isOpen() is true.

See also data() and setBuffer().
*/
func (this *QBuffer) SetData1(data string, len_ int) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer7setDataEPKci", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] const QByteArray & data() const

/*
Returns the data contained in the buffer.

This is the same as buffer().

See also setData() and setBuffer().
*/
func (this *QBuffer) Data() *QByteArray {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QBuffer4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool open(QIODevice::OpenMode)

/*
Reimplemented from QIODevice::open().
*/
func (this *QBuffer) Open(openMode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer4openE6QFlagsIN9QIODevice12OpenModeFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), openMode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void close()

/*
Reimplemented from QIODevice::close().
*/
func (this *QBuffer) Close() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 size() const

/*
Reimplemented from QIODevice::size().
*/
func (this *QBuffer) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QBuffer4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbuffer.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] qint64 pos() const

/*
Reimplemented from QIODevice::pos().
*/
func (this *QBuffer) Pos() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QBuffer3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbuffer.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool seek(qint64)

/*
Reimplemented from QIODevice::seek().
*/
func (this *QBuffer) Seek(off int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer4seekEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), off)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool atEnd() const

/*
Reimplemented from QIODevice::atEnd().
*/
func (this *QBuffer) AtEnd() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QBuffer5atEndEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canReadLine() const

/*
Reimplemented from QIODevice::canReadLine().
*/
func (this *QBuffer) CanReadLine() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QBuffer11canReadLineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:87
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void connectNotify(const QMetaMethod &)

/*

 */
func (this *QBuffer) ConnectNotify(arg0 QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMetaMethod_PTR() != nil {
		convArg0 = arg0.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer13connectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:88
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void disconnectNotify(const QMetaMethod &)

/*

 */
func (this *QBuffer) DisconnectNotify(arg0 QMetaMethod_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMetaMethod_PTR() != nil {
		convArg0 = arg0.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer16disconnectNotifyERK11QMetaMethod", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 readData(char *, qint64)

/*
Reimplemented from QIODevice::readData().
*/
func (this *QBuffer) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer8readDataEPcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbuffer.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] qint64 writeData(const char *, qint64)

/*
Reimplemented from QIODevice::writeData().
*/
func (this *QBuffer) WriteData(data string, len_ int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QBuffer9writeDataEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
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
