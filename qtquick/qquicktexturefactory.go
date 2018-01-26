package qtquick

// /usr/include/qt/QtQuick/qquickimageprovider.h
// #include <qquickimageprovider.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QQuickTextureFactory struct {
	*qtcore.QObject
}

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
// Public
// void QQuickTextureFactory()
func NewQQuickTextureFactory() *QQuickTextureFactory {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QQuickTextureFactoryC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickTextureFactoryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:61
// index:0
// Public virtual
// void ~QQuickTextureFactory()
func DeleteQQuickTextureFactory(*QQuickTextureFactory) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QQuickTextureFactoryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:63
// index:0
// Public pure virtual
// QSGTexture * createTexture(class QQuickWindow *)
func (this *QQuickTextureFactory) CreateTexture(window *QQuickWindow /*777 QQuickWindow **/) *QSGTexture /*777 QSGTexture **/ {
	var convArg0 = window.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QQuickTextureFactory13createTextureEP12QQuickWindow", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:64
// index:0
// Public pure virtual
// QSize textureSize()
func (this *QQuickTextureFactory) TextureSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QQuickTextureFactory11textureSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:65
// index:0
// Public pure virtual
// int textureByteCount()
func (this *QQuickTextureFactory) TextureByteCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QQuickTextureFactory16textureByteCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:66
// index:0
// Public virtual
// QImage image()
func (this *QQuickTextureFactory) Image() *qtgui.QImage /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QQuickTextureFactory5imageEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:68
// index:0
// Public static
// QQuickTextureFactory * textureFactoryForImage(const class QImage &)
func (this *QQuickTextureFactory) TextureFactoryForImage(image *qtgui.QImage) *QQuickTextureFactory /*777 QQuickTextureFactory **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QQuickTextureFactory22textureFactoryForImageERK6QImage", ffiqt.FFI_TYPE_POINTER, image)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQQuickTextureFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QQuickTextureFactory_TextureFactoryForImage(image *qtgui.QImage) *QQuickTextureFactory /*777 QQuickTextureFactory **/ {
	var nilthis *QQuickTextureFactory
	rv := nilthis.TextureFactoryForImage(image)
	return rv
}

//  body block end
