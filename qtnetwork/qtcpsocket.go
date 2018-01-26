package qtnetwork

// /usr/include/qt/QtNetwork/qtcpsocket.h
// #include <qtcpsocket.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 30
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTcpSocket struct {
	*QAbstractSocket
}

func (this *QTcpSocket) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractSocket.GetCthis()
	}
}
func (this *QTcpSocket) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractSocket = NewQAbstractSocketFromPointer(cthis)
}
func NewQTcpSocketFromPointer(cthis unsafe.Pointer) *QTcpSocket {
	bcthis0 := NewQAbstractSocketFromPointer(cthis)
	return &QTcpSocket{bcthis0}
}
func (*QTcpSocket) NewFromPointer(cthis unsafe.Pointer) *QTcpSocket {
	return NewQTcpSocketFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qtcpsocket.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QTcpSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTcpSocket10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtNetwork/qtcpsocket.h:56
// index:0
// Public
// void QTcpSocket(class QObject *)
func NewQTcpSocket(parent *qtcore.QObject /*777 QObject **/) *QTcpSocket {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpSocketC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTcpSocketFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qtcpsocket.h:57
// index:0
// Public virtual
// void ~QTcpSocket()
func DeleteQTcpSocket(*QTcpSocket) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTcpSocketD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
