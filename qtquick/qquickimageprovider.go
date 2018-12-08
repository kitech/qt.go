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
// extern C begin: 7
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
type QQuickImageProvider struct {
	*qtrt.CObject
}
type QQuickImageProvider_ITF interface {
	QQuickImageProvider_PTR() *QQuickImageProvider
}

func (ptr *QQuickImageProvider) QQuickImageProvider_PTR() *QQuickImageProvider { return ptr }

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
// [-2] void QQuickImageProvider(QQmlImageProviderBase::ImageType, QQmlImageProviderBase::Flags)

/*
Creates an image provider that will provide images of the given type and behave according to the given flags.
*/
func (*QQuickImageProvider) NewForInherit(type_ int, flags int) *QQuickImageProvider {
	return NewQQuickImageProvider(type_, flags)
}
func NewQQuickImageProvider(type_ int, flags int) *QQuickImageProvider {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProviderC2EN21QQmlImageProviderBase9ImageTypeE6QFlagsINS0_4FlagEE", qtrt.FFI_TYPE_POINTER, type_, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickImageProviderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuickImageProvider)
	return gothis
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickImageProvider(QQmlImageProviderBase::ImageType, QQmlImageProviderBase::Flags)

/*
Creates an image provider that will provide images of the given type and behave according to the given flags.
*/
func (*QQuickImageProvider) NewForInheritp(type_ int) *QQuickImageProvider {
	return NewQQuickImageProviderp(type_)
}
func NewQQuickImageProviderp(type_ int) *QQuickImageProvider {
	// arg: 1, QQmlImageProviderBase::Flags=Typedef, QQmlImageProviderBase::Flags=Typedef, QFlags<QQmlImageProviderBase::Flag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProviderC2EN21QQmlImageProviderBase9ImageTypeE6QFlagsINS0_4FlagEE", qtrt.FFI_TYPE_POINTER, type_, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickImageProviderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQuickImageProvider)
	return gothis
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickImageProvider()

/*

 */
func DeleteQQuickImageProvider(this *QQuickImageProvider) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProviderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QQmlImageProviderBase::ImageType imageType() const

/*
Reimplemented from QQmlImageProviderBase::imageType().

Returns the image type supported by this provider.
*/
func (this *QQuickImageProvider) ImageType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QQuickImageProvider9imageTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QQmlImageProviderBase::Flags flags() const

/*
Reimplemented from QQmlImageProviderBase::flags().

Returns the flags set for this provider.
*/
func (this *QQuickImageProvider) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QQuickImageProvider5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QImage requestImage(const QString &, QSize *, const QSize &)

/*
Implement this method to return the image with id. The default implementation returns an empty image.

The id is the requested image source, with the "image:" scheme and provider identifier removed. For example, if the image source was "image://myprovider/icons/home", the given id would be "icons/home".

The requestedSize corresponds to the Image::sourceSize requested by an Image item. If requestedSize is a valid size, the image returned should be of that size.

In all cases, size must be set to the original size of the image. This is used to set the width and height of the relevant Image if these values have not been set explicitly.

Note: this method may be called by multiple threads, so ensure the implementation of this method is reentrant.
*/
func (this *QQuickImageProvider) RequestImage(id string, size qtcore.QSize_ITF /*777 QSize **/, requestedSize qtcore.QSize_ITF) *qtgui.QImage /*123*/ {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if requestedSize != nil && requestedSize.QSize_PTR() != nil {
		convArg2 = requestedSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProvider12requestImageERK7QStringP5QSizeRKS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQImage)
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:104
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QPixmap requestPixmap(const QString &, QSize *, const QSize &)

/*
Implement this method to return the pixmap with id. The default implementation returns an empty pixmap.

The id is the requested image source, with the "image:" scheme and provider identifier removed. For example, if the image source was "image://myprovider/icons/home", the given id would be "icons/home".

The requestedSize corresponds to the Image::sourceSize requested by an Image item. If requestedSize is a valid size, the image returned should be of that size.

In all cases, size must be set to the original size of the image. This is used to set the width and height of the relevant Image if these values have not been set explicitly.

Note: this method may be called by multiple threads, so ensure the implementation of this method is reentrant.
*/
func (this *QQuickImageProvider) RequestPixmap(id string, size qtcore.QSize_ITF /*777 QSize **/, requestedSize qtcore.QSize_ITF) *qtgui.QPixmap /*123*/ {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if requestedSize != nil && requestedSize.QSize_PTR() != nil {
		convArg2 = requestedSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProvider13requestPixmapERK7QStringP5QSizeRKS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtQuick/qquickimageprovider.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QQuickTextureFactory * requestTexture(const QString &, QSize *, const QSize &)

/*
Implement this method to return the texture with id. The default implementation returns 0.

The id is the requested image source, with the "image:" scheme and provider identifier removed. For example, if the image source was "image://myprovider/icons/home", the given id would be "icons/home".

The requestedSize corresponds to the Image::sourceSize requested by an Image item. If requestedSize is a valid size, the image returned should be of that size.

In all cases, size must be set to the original size of the image. This is used to set the width and height of the relevant Image if these values have not been set explicitly.

Note: this method may be called by multiple threads, so ensure the implementation of this method is reentrant.
*/
func (this *QQuickImageProvider) RequestTexture(id string, size qtcore.QSize_ITF /*777 QSize **/, requestedSize qtcore.QSize_ITF) *QQuickTextureFactory /*777 QQuickTextureFactory **/ {
	var tmpArg0 = qtcore.NewQString5(id)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if requestedSize != nil && requestedSize.QSize_PTR() != nil {
		convArg2 = requestedSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickImageProvider14requestTextureERK7QStringP5QSizeRKS3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickTextureFactoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
