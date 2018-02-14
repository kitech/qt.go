package qtquick

// /usr/include/qt/QtQuick/qquickimageprovider.h
// #include <qquickimageprovider.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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

type QQuickTextureFactory struct {
	*qtcore.QObject
}
type QQuickTextureFactory_ITF interface {
	qtcore.QObject_ITF
	QQuickTextureFactory_PTR() *QQuickTextureFactory
}

func (ptr *QQuickTextureFactory) QQuickTextureFactory_PTR() *QQuickTextureFactory { return ptr }

func (this *QQuickTextureFactory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickTextureFactory) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQuickTextureFactoryFromPointer(cthis unsafe.Pointer) *QQuickTextureFactory {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQuickTextureFactory{bcthis0}
}
func (*QQuickTextureFactory) NewFromPointer(cthis unsafe.Pointer) *QQuickTextureFactory {
	return NewQQuickTextureFactoryFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickTextureFactory()
func NewQQuickTextureFactory() *QQuickTextureFactory {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickTextureFactoryC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickTextureFactoryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickTextureFactory")
	return gothis
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickTextureFactory()
func DeleteQQuickTextureFactory(this *QQuickTextureFactory) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickTextureFactoryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSGTexture * createTexture(QQuickWindow *)
func (this *QQuickTextureFactory) CreateTexture(window QQuickWindow_ITF /*777 QQuickWindow **/) *QSGTexture /*777 QSGTexture **/ {
	var convArg0 unsafe.Pointer
	if window != nil && window.QQuickWindow_PTR() != nil {
		convArg0 = window.QQuickWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickTextureFactory13createTextureEP12QQuickWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize textureSize()
func (this *QQuickTextureFactory) TextureSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickTextureFactory11textureSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int textureByteCount()
func (this *QQuickTextureFactory) TextureByteCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickTextureFactory16textureByteCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QImage image()
func (this *QQuickTextureFactory) Image() *qtgui.QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QQuickTextureFactory5imageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQImage)
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:68
// index:0
// Public static Visibility=Default Availability=Available
// [8] QQuickTextureFactory * textureFactoryForImage(const QImage &)
func (this *QQuickTextureFactory) TextureFactoryForImage(image qtgui.QImage_ITF) *QQuickTextureFactory /*777 QQuickTextureFactory **/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QQuickTextureFactory22textureFactoryForImageERK6QImage", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickTextureFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QQuickTextureFactory_TextureFactoryForImage(image qtgui.QImage_ITF) *QQuickTextureFactory /*777 QQuickTextureFactory **/ {
	var nilthis *QQuickTextureFactory
	rv := nilthis.TextureFactoryForImage(image)
	return rv
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
