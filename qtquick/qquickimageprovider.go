package qtquick

// /usr/include/qt/QtQuick/qquickimageprovider.h
// #include <qquickimageprovider.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"
import "qt.go/qtgui"
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
		gopp.KeepMe()
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

//  ext block end

//  body block begin

type QQuickImageProvider struct {
	*qtrt.CObject
}

func (this *QQuickImageProvider) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQuickImageProvider) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQuickImageProviderFromPointer(cthis unsafe.Pointer) *QQuickImageProvider {
	return &QQuickImageProvider{&qtrt.CObject{cthis}}
}
func (*QQuickImageProvider) NewFromPointer(cthis unsafe.Pointer) *QQuickImageProvider {
	return NewQQuickImageProviderFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickImageProvider(enum QQmlImageProviderBase::ImageType, QQmlImageProviderBase::Flags)
func NewQQuickImageProvider(type_ int, flags int) *QQuickImageProvider {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProviderC2EN21QQmlImageProviderBase9ImageTypeE6QFlagsINS0_4FlagEE", qtrt.FFI_TYPE_POINTER, type_, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickImageProviderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuickImageProvider)
	return gothis
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickImageProvider()
func DeleteQQuickImageProvider(this *QQuickImageProvider) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProviderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QQmlImageProviderBase::ImageType imageType()
func (this *QQuickImageProvider) ImageType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QQuickImageProvider9imageTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QQmlImageProviderBase::Flags flags()
func (this *QQuickImageProvider) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QQuickImageProvider5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QImage requestImage(const QString &, QSize *, const QSize &)
func (this *QQuickImageProvider) RequestImage(id *qtcore.QString, size *qtcore.QSize /*777 QSize **/, requestedSize *qtcore.QSize) *qtgui.QImage /*123*/ {
	var convArg0 = id.GetCthis()
	var convArg1 = size.GetCthis()
	var convArg2 = requestedSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProvider12requestImageERK7QStringP5QSizeRKS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQImage)
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:104
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap requestPixmap(const QString &, QSize *, const QSize &)
func (this *QQuickImageProvider) RequestPixmap(id *qtcore.QString, size *qtcore.QSize /*777 QSize **/, requestedSize *qtcore.QSize) *qtgui.QPixmap /*123*/ {
	var convArg0 = id.GetCthis()
	var convArg1 = size.GetCthis()
	var convArg2 = requestedSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProvider13requestPixmapERK7QStringP5QSizeRKS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QQuickTextureFactory * requestTexture(const QString &, QSize *, const QSize &)
func (this *QQuickImageProvider) RequestTexture(id *qtcore.QString, size *qtcore.QSize /*777 QSize **/, requestedSize *qtcore.QSize) *QQuickTextureFactory /*777 QQuickTextureFactory **/ {
	var convArg0 = id.GetCthis()
	var convArg1 = size.GetCthis()
	var convArg2 = requestedSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProvider14requestTextureERK7QStringP5QSizeRKS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickTextureFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
