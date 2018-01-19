//  header block begin
// /usr/include/qt/QtCore/qbuffer.h
// #include <qbuffer.h>
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
type QBuffer struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qbuffer.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QBuffer) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:60
// index:0
// void QBuffer(class QObject *)
func NewQBuffer(parent unsafe.Pointer) *QBuffer {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QBuffer{cthis}
}

// /usr/include/qt/QtCore/qbuffer.h:61
// index:1
// void QBuffer(class QByteArray *, class QObject *)
func NewQBuffer_1(buf unsafe.Pointer, parent unsafe.Pointer) *QBuffer {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferC2EP10QByteArrayP7QObject", ffiqt.FFI_TYPE_VOID, cthis, buf, parent)
	gopp.ErrPrint(err, rv)
	return &QBuffer{cthis}
}

// /usr/include/qt/QtCore/qbuffer.h:66
// index:0
// virtual
// void ~QBuffer()
func DeleteQBuffer(*QBuffer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBufferD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:68
// index:0
// QByteArray & buffer()
func (this *QBuffer) Buffer() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer6bufferEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:69
// index:1
// const QByteArray & buffer()
func (this *QBuffer) Buffer_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer6bufferEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:70
// index:0
// void setBuffer(class QByteArray *)
func (this *QBuffer) SetBuffer(a unsafe.Pointer) {
	// 0: (, QByteArray * a), (a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer9setBufferEP10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:72
// index:0
// void setData(const class QByteArray &)
func (this *QBuffer) SetData(data unsafe.Pointer) {
	// 0: (, const QByteArray & data), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer7setDataERK10QByteArray", ffiqt.FFI_TYPE_VOID, this.cthis, data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:73
// index:1
// inline
// void setData(const char *, int)
func (this *QBuffer) SetData_1(data unsafe.Pointer, len int) {
	// 1: (, const char * data, int len), (data, &len)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer7setDataEPKci", ffiqt.FFI_TYPE_VOID, this.cthis, data, &len)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:74
// index:0
// const QByteArray & data()
func (this *QBuffer) Data() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer4dataEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:78
// index:0
// virtual
// void close()
func (this *QBuffer) Close() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer5closeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:79
// index:0
// virtual
// qint64 size()
func (this *QBuffer) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:80
// index:0
// virtual
// qint64 pos()
func (this *QBuffer) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:81
// index:0
// virtual
// bool seek(qint64)
func (this *QBuffer) Seek(off int64) {
	// 0: (, qint64 off), (&off)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QBuffer4seekEx", ffiqt.FFI_TYPE_VOID, this.cthis, &off)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:82
// index:0
// virtual
// bool atEnd()
func (this *QBuffer) AtEnd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer5atEndEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbuffer.h:83
// index:0
// virtual
// bool canReadLine()
func (this *QBuffer) CanReadLine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QBuffer11canReadLineEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
