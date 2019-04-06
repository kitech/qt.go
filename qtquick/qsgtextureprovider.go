package qtquick

// /usr/include/qt/QtQuick/qsgtextureprovider.h
// #include <qsgtextureprovider.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

/*

 */
type QSGTextureProvider struct {
	*qtcore.QObject
}
type QSGTextureProvider_ITF interface {
	qtcore.QObject_ITF
	QSGTextureProvider_PTR() *QSGTextureProvider
}

func (ptr *QSGTextureProvider) QSGTextureProvider_PTR() *QSGTextureProvider { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSGTextureProvider) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSGTextureProvider10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtextureprovider.h:52
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSGTexture * texture() const

/*
Returns a pointer to the texture object.
*/
func (this *QSGTextureProvider) Texture() *QSGTexture /*777 QSGTexture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QSGTextureProvider7textureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtextureprovider.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textureChanged()

/*
This signal is emitted when the texture changes.
*/
func (this *QSGTextureProvider) TextureChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSGTextureProvider14textureChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQSGTextureProvider(this *QSGTextureProvider) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QSGTextureProviderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11623() {
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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
