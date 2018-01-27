package qtcore

// /usr/include/qt/QtCore/qbuffer.h
// #include <qbuffer.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 57
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
type QBuffer struct {
	*QIODevice
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QBuffer) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:60
// index:0
// Public
// void QBuffer(QObject *)
func NewQBuffer(parent *QObject /*777 QObject **/) *QBuffer {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQBufferFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbuffer.h:61
// index:1
// Public
// void QBuffer(QByteArray *, QObject *)
func NewQBuffer_1(buf *QByteArray /*777 QByteArray **/, parent *QObject /*777 QObject **/) *QBuffer {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = buf.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferC2EP10QByteArrayP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQBufferFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qbuffer.h:66
// index:0
// Public virtual
// void ~QBuffer()
func DeleteQBuffer(*QBuffer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:68
// index:0
// Public
// QByteArray & buffer()
func (this *QBuffer) Buffer() *QByteArray {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer6bufferEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:69
// index:1
// Public
// const QByteArray & buffer()
func (this *QBuffer) Buffer_1() *QByteArray {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer6bufferEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:70
// index:0
// Public
// void setBuffer(QByteArray *)
func (this *QBuffer) SetBuffer(a *QByteArray /*777 QByteArray **/) {
	var convArg0 = a.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer9setBufferEP10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:72
// index:0
// Public
// void setData(const QByteArray &)
func (this *QBuffer) SetData(data *QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer7setDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:73
// index:1
// Public inline
// void setData(const char *, int)
func (this *QBuffer) SetData_1(data string, len int) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer7setDataEPKci", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:74
// index:0
// Public
// const QByteArray & data()
func (this *QBuffer) Data() *QByteArray {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer4dataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtCore/qbuffer.h:78
// index:0
// Public virtual
// void close()
func (this *QBuffer) Close() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:79
// index:0
// Public virtual
// qint64 size()
func (this *QBuffer) Size() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbuffer.h:80
// index:0
// Public virtual
// qint64 pos()
func (this *QBuffer) Pos() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer3posEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbuffer.h:81
// index:0
// Public virtual
// bool seek(qint64)
func (this *QBuffer) Seek(off int64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer4seekEx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), off)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:82
// index:0
// Public virtual
// bool atEnd()
func (this *QBuffer) AtEnd() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer5atEndEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:83
// index:0
// Public virtual
// bool canReadLine()
func (this *QBuffer) CanReadLine() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer11canReadLineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qbuffer.h:87
// index:0
// Protected virtual
// void connectNotify(const QMetaMethod &)
func (this *QBuffer) ConnectNotify(arg0 *QMetaMethod) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer13connectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:88
// index:0
// Protected virtual
// void disconnectNotify(const QMetaMethod &)
func (this *QBuffer) DisconnectNotify(arg0 *QMetaMethod) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer16disconnectNotifyERK11QMetaMethod", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:90
// index:0
// Protected virtual
// qint64 readData(char *, qint64)
func (this *QBuffer) ReadData(data string, maxlen int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer8readDataEPcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, maxlen)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qbuffer.h:91
// index:0
// Protected virtual
// qint64 writeData(const char *, qint64)
func (this *QBuffer) WriteData(data string, len int64) int64 {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer9writeDataEPKcx", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

//  body block end
