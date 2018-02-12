package qtnetwork

// /usr/include/qt/QtNetwork/qtcpsocket.h
// #include <qtcpsocket.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTcpSocket) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTcpSocket10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtNetwork/qtcpsocket.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTcpSocket(QObject *)
func NewQTcpSocket(parent *qtcore.QObject /*777 QObject **/) *QTcpSocket {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpSocketC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTcpSocketFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtNetwork/qtcpsocket.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTcpSocket()
func DeleteQTcpSocket(this *QTcpSocket) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTcpSocketD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
