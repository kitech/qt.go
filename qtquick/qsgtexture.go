package qtquick

// /usr/include/qt/QtQuick/qsgtexture.h
// #include <qsgtexture.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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

type QSGTexture struct {
	*qtcore.QObject
}
type QSGTexture_ITF interface {
	qtcore.QObject_ITF
	QSGTexture_PTR() *QSGTexture
}

func (ptr *QSGTexture) QSGTexture_PTR() *QSGTexture { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSGTexture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtexture.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGTexture()
func NewQSGTexture() *QSGTexture {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTextureC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSGTexture")
	return gothis
}

// /usr/include/qt/QtQuick/qsgtexture.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGTexture()
func DeleteQSGTexture(this *QSGTexture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTextureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgtexture.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int textureId()
func (this *QSGTexture) TextureId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture9textureIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgtexture.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize textureSize()
func (this *QSGTexture) TextureSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture11textureSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtQuick/qsgtexture.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool hasAlphaChannel()
func (this *QSGTexture) HasAlphaChannel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture15hasAlphaChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgtexture.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool hasMipmaps()
func (this *QSGTexture) HasMipmaps() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture10hasMipmapsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgtexture.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF normalizedTextureSubRect()
func (this *QSGTexture) NormalizedTextureSubRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture24normalizedTextureSubRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qsgtexture.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isAtlasTexture()
func (this *QSGTexture) IsAtlasTexture() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture14isAtlasTextureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgtexture.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGTexture * removedFromAtlas()
func (this *QSGTexture) RemovedFromAtlas() *QSGTexture /*777 QSGTexture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture16removedFromAtlasEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtexture.h:90
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void bind()
func (this *QSGTexture) Bind() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture4bindEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateBindOptions(_Bool)
func (this *QSGTexture) UpdateBindOptions(force bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture17updateBindOptionsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), force)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMipmapFiltering(enum QSGTexture::Filtering)
func (this *QSGTexture) SetMipmapFiltering(filter int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture18setMipmapFilteringENS_9FilteringE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::Filtering mipmapFiltering()
func (this *QSGTexture) MipmapFiltering() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture15mipmapFilteringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFiltering(enum QSGTexture::Filtering)
func (this *QSGTexture) SetFiltering(filter int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture12setFilteringENS_9FilteringE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::Filtering filtering()
func (this *QSGTexture) Filtering() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture9filteringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAnisotropyLevel(enum QSGTexture::AnisotropyLevel)
func (this *QSGTexture) SetAnisotropyLevel(level int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture18setAnisotropyLevelENS_15AnisotropyLevelE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), level)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::AnisotropyLevel anisotropyLevel()
func (this *QSGTexture) AnisotropyLevel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture15anisotropyLevelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalWrapMode(enum QSGTexture::WrapMode)
func (this *QSGTexture) SetHorizontalWrapMode(hwrap int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture21setHorizontalWrapModeENS_8WrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hwrap)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::WrapMode horizontalWrapMode()
func (this *QSGTexture) HorizontalWrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture18horizontalWrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalWrapMode(enum QSGTexture::WrapMode)
func (this *QSGTexture) SetVerticalWrapMode(vwrap int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture19setVerticalWrapModeENS_8WrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), vwrap)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::WrapMode verticalWrapMode()
func (this *QSGTexture) VerticalWrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture16verticalWrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF convertToNormalizedSourceRect(const QRectF &)
func (this *QSGTexture) ConvertToNormalizedSourceRect(rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture29convertToNormalizedSourceRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
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
