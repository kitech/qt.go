package qtquick

// /usr/include/qt/QtQuick/qsgtexture.h
// #include <qsgtexture.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QSGTexture struct {
	*qtcore.QObject
}

func (this *QSGTexture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSGTexture) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSGTextureFromPointer(cthis unsafe.Pointer) *QSGTexture {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSGTexture{bcthis0}
}
func (*QSGTexture) NewFromPointer(cthis unsafe.Pointer) *QSGTexture {
	return NewQSGTextureFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgtexture.h:52
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSGTexture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgtexture.h:56
// index:0
// Public
// void QSGTexture()
func NewQSGTexture() *QSGTexture {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSGTextureC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGTextureFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgtexture.h:57
// index:0
// Public virtual
// void ~QSGTexture()
func DeleteQSGTexture(*QSGTexture) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSGTextureD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:79
// index:0
// Public pure virtual
// int textureId()
func (this *QSGTexture) TextureId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture9textureIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtQuick/qsgtexture.h:80
// index:0
// Public pure virtual
// QSize textureSize()
func (this *QSGTexture) TextureSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture11textureSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgtexture.h:81
// index:0
// Public pure virtual
// bool hasAlphaChannel()
func (this *QSGTexture) HasAlphaChannel() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture15hasAlphaChannelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgtexture.h:82
// index:0
// Public pure virtual
// bool hasMipmaps()
func (this *QSGTexture) HasMipmaps() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture10hasMipmapsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgtexture.h:84
// index:0
// Public virtual
// QRectF normalizedTextureSubRect()
func (this *QSGTexture) NormalizedTextureSubRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture24normalizedTextureSubRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgtexture.h:86
// index:0
// Public virtual
// bool isAtlasTexture()
func (this *QSGTexture) IsAtlasTexture() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture14isAtlasTextureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgtexture.h:88
// index:0
// Public virtual
// QSGTexture * removedFromAtlas()
func (this *QSGTexture) RemovedFromAtlas() *QSGTexture /*777 QSGTexture **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture16removedFromAtlasEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgtexture.h:90
// index:0
// Public pure virtual
// void bind()
func (this *QSGTexture) Bind() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSGTexture4bindEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:91
// index:0
// Public
// void updateBindOptions(_Bool)
func (this *QSGTexture) UpdateBindOptions(force bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSGTexture17updateBindOptionsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), force)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:93
// index:0
// Public
// void setMipmapFiltering(enum QSGTexture::Filtering)
func (this *QSGTexture) SetMipmapFiltering(filter int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSGTexture18setMipmapFilteringENS_9FilteringE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:94
// index:0
// Public
// QSGTexture::Filtering mipmapFiltering()
func (this *QSGTexture) MipmapFiltering() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture15mipmapFilteringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:96
// index:0
// Public
// void setFiltering(enum QSGTexture::Filtering)
func (this *QSGTexture) SetFiltering(filter int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSGTexture12setFilteringENS_9FilteringE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:97
// index:0
// Public
// QSGTexture::Filtering filtering()
func (this *QSGTexture) Filtering() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture9filteringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:99
// index:0
// Public
// void setAnisotropyLevel(enum QSGTexture::AnisotropyLevel)
func (this *QSGTexture) SetAnisotropyLevel(level int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSGTexture18setAnisotropyLevelENS_15AnisotropyLevelE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), level)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:100
// index:0
// Public
// QSGTexture::AnisotropyLevel anisotropyLevel()
func (this *QSGTexture) AnisotropyLevel() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture15anisotropyLevelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:102
// index:0
// Public
// void setHorizontalWrapMode(enum QSGTexture::WrapMode)
func (this *QSGTexture) SetHorizontalWrapMode(hwrap int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSGTexture21setHorizontalWrapModeENS_8WrapModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hwrap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:103
// index:0
// Public
// QSGTexture::WrapMode horizontalWrapMode()
func (this *QSGTexture) HorizontalWrapMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture18horizontalWrapModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:105
// index:0
// Public
// void setVerticalWrapMode(enum QSGTexture::WrapMode)
func (this *QSGTexture) SetVerticalWrapMode(vwrap int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QSGTexture19setVerticalWrapModeENS_8WrapModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), vwrap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:106
// index:0
// Public
// QSGTexture::WrapMode verticalWrapMode()
func (this *QSGTexture) VerticalWrapMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture16verticalWrapModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:108
// index:0
// Public inline
// QRectF convertToNormalizedSourceRect(const class QRectF &)
func (this *QSGTexture) ConvertToNormalizedSourceRect(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QSGTexture29convertToNormalizedSourceRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QSGTexture__WrapMode = int

const QSGTexture__Repeat QSGTexture__WrapMode = 0
const QSGTexture__ClampToEdge QSGTexture__WrapMode = 1
const QSGTexture__MirroredRepeat QSGTexture__WrapMode = 2

type QSGTexture__Filtering = int

const QSGTexture__None QSGTexture__Filtering = 0
const QSGTexture__Nearest QSGTexture__Filtering = 1
const QSGTexture__Linear QSGTexture__Filtering = 2

type QSGTexture__AnisotropyLevel = int

const QSGTexture__AnisotropyNone QSGTexture__AnisotropyLevel = 0
const QSGTexture__Anisotropy2x QSGTexture__AnisotropyLevel = 1
const QSGTexture__Anisotropy4x QSGTexture__AnisotropyLevel = 2
const QSGTexture__Anisotropy8x QSGTexture__AnisotropyLevel = 3
const QSGTexture__Anisotropy16x QSGTexture__AnisotropyLevel = 4

//  body block end
