package qtqml

// /usr/include/qt/QtQml/qqmlnetworkaccessmanagerfactory.h
// #include <qqmlnetworkaccessmanagerfactory.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QQmlNetworkAccessManagerFactory struct {
	*qtrt.CObject
}
type QQmlNetworkAccessManagerFactory_ITF interface {
	QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory
}

func (ptr *QQmlNetworkAccessManagerFactory) QQmlNetworkAccessManagerFactory_PTR() *QQmlNetworkAccessManagerFactory {
	return ptr
}

func (this *QQmlNetworkAccessManagerFactory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlNetworkAccessManagerFactory) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlNetworkAccessManagerFactoryFromPointer(cthis unsafe.Pointer) *QQmlNetworkAccessManagerFactory {
	return &QQmlNetworkAccessManagerFactory{&qtrt.CObject{cthis}}
}
func (*QQmlNetworkAccessManagerFactory) NewFromPointer(cthis unsafe.Pointer) *QQmlNetworkAccessManagerFactory {
	return NewQQmlNetworkAccessManagerFactoryFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_11539() {
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
