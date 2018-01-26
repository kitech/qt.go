package qtquick

// /usr/include/qt/QtQuick/qsgtextureprovider.h
// #include <qsgtextureprovider.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QSGTextureProvider struct {
	*qtcore.QObject
}

func (this *QSGTextureProvider) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSGTextureProvider) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSGTextureProviderFromPointer(cthis unsafe.Pointer) *QSGTextureProvider {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSGTextureProvider{bcthis0}
}
func (*QSGTextureProvider) NewFromPointer(cthis unsafe.Pointer) *QSGTextureProvider {
	return NewQSGTextureProviderFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgtextureprovider.h:50
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSGTextureProvider) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSGTextureProvider10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgtextureprovider.h:52
// index:0
// Public pure virtual
// QSGTexture * texture()
func (this *QSGTextureProvider) Texture() *QSGTexture /*777 QSGTexture **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QSGTextureProvider7textureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgtextureprovider.h:55
// index:0
// Public
// void textureChanged()
func (this *QSGTextureProvider) TextureChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QSGTextureProvider14textureChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
