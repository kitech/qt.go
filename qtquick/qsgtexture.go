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

/*

 */
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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSGTexture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtexture.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGTexture()

/*
Constructs the QSGTexture base class.
*/
func (*QSGTexture) NewForInherit() *QSGTexture {
	return NewQSGTexture()
}
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

/*

 */
func DeleteQSGTexture(this *QSGTexture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTextureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgtexture.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int textureId() const

/*
Returns the OpenGL texture id for this texture.

The default value is 0, indicating that it is an invalid texture id.

The function should at all times return the correct texture id.

Warning: This function can only be called from the rendering thread.
*/
func (this *QSGTexture) TextureId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture9textureIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgtexture.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize textureSize() const

/*
Returns the size of the texture.
*/
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
// [1] bool hasAlphaChannel() const

/*
Returns true if the texture data contains an alpha channel.
*/
func (this *QSGTexture) HasAlphaChannel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture15hasAlphaChannelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgtexture.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool hasMipmaps() const

/*
Returns true if the texture data contains mipmap levels.
*/
func (this *QSGTexture) HasMipmaps() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture10hasMipmapsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgtexture.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF normalizedTextureSubRect() const

/*
Returns the rectangle inside textureSize() that this texture represents in normalized coordinates.

The default implementation returns a rect at position (0, 0) with width and height of 1.
*/
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
// [1] bool isAtlasTexture() const

/*
Returns weither this texture is part of an atlas or not.

The default implementation returns false.
*/
func (this *QSGTexture) IsAtlasTexture() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture14isAtlasTextureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgtexture.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGTexture * removedFromAtlas() const

/*
This function returns a copy of the current texture which is removed from its atlas.

The current texture remains unchanged, so texture coordinates do not need to be updated.

Removing a texture from an atlas is primarily useful when passing it to a shader that operates on the texture coordinates 0-1 instead of the texture subrect inside the atlas.

If the texture is not part of a texture atlas, this function returns 0.

Implementations of this function are recommended to return the same instance for multiple calls to limit memory usage.

Warning: This function can only be called from the rendering thread.
*/
func (this *QSGTexture) RemovedFromAtlas() *QSGTexture /*777 QSGTexture **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture16removedFromAtlasEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgtexture.h:90
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void bind()

/*
Call this function to bind this texture to the current texture target.

Binding a texture may also include uploading the texture data from a previously set QImage.

Warning: This function can only be called from the rendering thread.
*/
func (this *QSGTexture) Bind() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture4bindEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateBindOptions(bool)

/*
Update the texture state to match the filtering, mipmap and wrap options currently set.

If force is true, all properties will be updated regardless of weither they have changed or not.
*/
func (this *QSGTexture) UpdateBindOptions(force bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture17updateBindOptionsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), force)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateBindOptions(bool)

/*
Update the texture state to match the filtering, mipmap and wrap options currently set.

If force is true, all properties will be updated regardless of weither they have changed or not.
*/
func (this *QSGTexture) UpdateBindOptionsp() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	force := false
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture17updateBindOptionsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), force)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMipmapFiltering(QSGTexture::Filtering)

/*
Sets the mipmap sampling mode to be used for the upcoming bind() call to filter.

Setting the mipmap filtering has no effect it the texture does not have mipmaps.

See also mipmapFiltering() and hasMipmaps().
*/
func (this *QSGTexture) SetMipmapFiltering(filter int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture18setMipmapFilteringENS_9FilteringE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::Filtering mipmapFiltering() const

/*
Returns whether mipmapping should be used when sampling from this texture.

See also setMipmapFiltering().
*/
func (this *QSGTexture) MipmapFiltering() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture15mipmapFilteringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFiltering(QSGTexture::Filtering)

/*
Sets the sampling mode to be used for the upcoming bind() call to filter.

See also filtering().
*/
func (this *QSGTexture) SetFiltering(filter int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture12setFilteringENS_9FilteringE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::Filtering filtering() const

/*
Returns the sampling mode to be used for this texture.

See also setFiltering().
*/
func (this *QSGTexture) Filtering() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture9filteringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAnisotropyLevel(QSGTexture::AnisotropyLevel)

/*
Sets the level of anisotropic filtering to be used for the upcoming bind() call to level. The default value is QSGTexture::AnisotropyNone, which means no anisotropic filtering is enabled.

This function was introduced in  Qt 5.9.

See also anisotropyLevel().
*/
func (this *QSGTexture) SetAnisotropyLevel(level int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture18setAnisotropyLevelENS_15AnisotropyLevelE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), level)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::AnisotropyLevel anisotropyLevel() const

/*
Returns the anisotropy level in use for filtering this texture.

This function was introduced in  Qt 5.9.

See also setAnisotropyLevel().
*/
func (this *QSGTexture) AnisotropyLevel() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture15anisotropyLevelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalWrapMode(QSGTexture::WrapMode)

/*
Sets the horizontal wrap mode to be used for the upcoming bind() call to hwrap

See also horizontalWrapMode().
*/
func (this *QSGTexture) SetHorizontalWrapMode(hwrap int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture21setHorizontalWrapModeENS_8WrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hwrap)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:103
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::WrapMode horizontalWrapMode() const

/*
Returns the horizontal wrap mode to be used for this texture.

See also setHorizontalWrapMode().
*/
func (this *QSGTexture) HorizontalWrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture18horizontalWrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalWrapMode(QSGTexture::WrapMode)

/*
Sets the vertical wrap mode to be used for the upcoming bind() call to vwrap

See also verticalWrapMode().
*/
func (this *QSGTexture) SetVerticalWrapMode(vwrap int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSGTexture19setVerticalWrapModeENS_8WrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), vwrap)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGTexture::WrapMode verticalWrapMode() const

/*
Returns the vertical wrap mode to be used for this texture.

See also setVerticalWrapMode().
*/
func (this *QSGTexture) VerticalWrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSGTexture16verticalWrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgtexture.h:108
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF convertToNormalizedSourceRect(const QRectF &) const

/*
Returns rect converted to normalized coordinates.

See also normalizedTextureSubRect().
*/
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

/*
Specifies how the texture should treat texture coordinates.


*/
type QSGTexture__WrapMode = int

//
const QSGTexture__Repeat QSGTexture__WrapMode = 0

//
const QSGTexture__ClampToEdge QSGTexture__WrapMode = 1

//
const QSGTexture__MirroredRepeat QSGTexture__WrapMode = 2

func (this *QSGTexture) WrapModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSGTexture_WrapModeItemName(val int) string {
	var nilthis *QSGTexture
	return nilthis.WrapModeItemName(val)
}

/*
Specifies how sampling of texels should filter when texture coordinates are not pixel aligned.


*/
type QSGTexture__Filtering = int

// No filtering should occur. This value is only used together with setMipmapFiltering().
const QSGTexture__None QSGTexture__Filtering = 0

// Sampling returns the nearest texel.
const QSGTexture__Nearest QSGTexture__Filtering = 1

// Sampling returns a linear interpolation of the neighboring texels.
const QSGTexture__Linear QSGTexture__Filtering = 2

func (this *QSGTexture) FilteringItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSGTexture_FilteringItemName(val int) string {
	var nilthis *QSGTexture
	return nilthis.FilteringItemName(val)
}

/*
Specifies the anisotropic filtering level to be used when the texture is not screen aligned.



This enum was introduced or modified in  Qt 5.9.

*/
type QSGTexture__AnisotropyLevel = int

// No anisotropic filtering.
const QSGTexture__AnisotropyNone QSGTexture__AnisotropyLevel = 0

//
const QSGTexture__Anisotropy2x QSGTexture__AnisotropyLevel = 1

//
const QSGTexture__Anisotropy4x QSGTexture__AnisotropyLevel = 2

//
const QSGTexture__Anisotropy8x QSGTexture__AnisotropyLevel = 3

//
const QSGTexture__Anisotropy16x QSGTexture__AnisotropyLevel = 4

func (this *QSGTexture) AnisotropyLevelItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSGTexture_AnisotropyLevelItemName(val int) string {
	var nilthis *QSGTexture
	return nilthis.AnisotropyLevelItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11603() {
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
