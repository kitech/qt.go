package qtandroidextras

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidserviceconnection.h
// #include <qandroidserviceconnection.h>
// #include <QtAndroidExtras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QAndroidServiceConnection struct {
	*qtrt.CObject
}

func (this *QAndroidServiceConnection) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidServiceConnection) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidServiceConnectionFromPointer(cthis unsafe.Pointer) *QAndroidServiceConnection {
	return &QAndroidServiceConnection{&qtrt.CObject{cthis}}
}
func (*QAndroidServiceConnection) NewFromPointer(cthis unsafe.Pointer) *QAndroidServiceConnection {
	return NewQAndroidServiceConnectionFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidserviceconnection.h:51
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidServiceConnection()
func NewQAndroidServiceConnection() *QAndroidServiceConnection {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAndroidServiceConnectionC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidServiceConnectionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidServiceConnection)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidserviceconnection.h:52
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAndroidServiceConnection(const QAndroidJniObject &)
func NewQAndroidServiceConnection_1(serviceConnection *QAndroidJniObject) *QAndroidServiceConnection {
	var convArg0 = serviceConnection.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAndroidServiceConnectionC2ERK17QAndroidJniObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidServiceConnectionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidServiceConnection)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidserviceconnection.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAndroidServiceConnection()
func DeleteQAndroidServiceConnection(this *QAndroidServiceConnection) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAndroidServiceConnectionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 1)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidserviceconnection.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void onServiceConnected(const QString &, const QAndroidBinder &)
func (this *QAndroidServiceConnection) OnServiceConnected(name string, serviceBinder *QAndroidBinder) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = serviceBinder.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAndroidServiceConnection18onServiceConnectedERK7QStringRK14QAndroidBinder", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidserviceconnection.h:56
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void onServiceDisconnected(const QString &)
func (this *QAndroidServiceConnection) OnServiceDisconnected(name string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN25QAndroidServiceConnection21onServiceDisconnectedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidserviceconnection.h:58
// index:0
// Public Visibility=Default Availability=Available
// [1] QAndroidJniObject handle()
func (this *QAndroidServiceConnection) Handle() *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QAndroidServiceConnection6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}

//  body block end
