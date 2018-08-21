package qtgui

// /usr/include/qt/QtGui/qsurfaceformat.h
// #include <qsurfaceformat.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QSurfaceFormat struct {
	*qtrt.CObject
}
type QSurfaceFormat_ITF interface {
	QSurfaceFormat_PTR() *QSurfaceFormat
}

func (ptr *QSurfaceFormat) QSurfaceFormat_PTR() *QSurfaceFormat { return ptr }

func (this *QSurfaceFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSurfaceFormat) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSurfaceFormatFromPointer(cthis unsafe.Pointer) *QSurfaceFormat {
	return &QSurfaceFormat{&qtrt.CObject{cthis}}
}
func (*QSurfaceFormat) NewFromPointer(cthis unsafe.Pointer) *QSurfaceFormat {
	return NewQSurfaceFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSurfaceFormat()

/*
Constructs a default initialized QSurfaceFormat.

Note: By default OpenGL 2.0 is requested since this provides the highest grade of portability between platforms and OpenGL implementations.
*/
func NewQSurfaceFormat() *QSurfaceFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSurfaceFormat)
	return gothis
}

// /usr/include/qt/QtGui/qsurfaceformat.h:95
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSurfaceFormat(QSurfaceFormat::FormatOptions)

/*
Constructs a default initialized QSurfaceFormat.

Note: By default OpenGL 2.0 is requested since this provides the highest grade of portability between platforms and OpenGL implementations.
*/
func NewQSurfaceFormat_1(options int) *QSurfaceFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormatC2E6QFlagsINS_12FormatOptionEE", qtrt.FFI_TYPE_POINTER, options)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSurfaceFormat)
	return gothis
}

// /usr/include/qt/QtGui/qsurfaceformat.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QSurfaceFormat & operator=(const QSurfaceFormat &)

/*

 */
func (this *QSurfaceFormat) Operator_equal(other QSurfaceFormat_ITF) *QSurfaceFormat {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSurfaceFormat_PTR() != nil {
		convArg0 = other.QSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormataSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qsurfaceformat.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSurfaceFormat()

/*

 */
func DeleteQSurfaceFormat(this *QSurfaceFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDepthBufferSize(int)

/*
Set the minimum depth buffer size to size.

See also depthBufferSize().
*/
func (this *QSurfaceFormat) SetDepthBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat18setDepthBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int depthBufferSize() const

/*
Returns the depth buffer size.

See also setDepthBufferSize().
*/
func (this *QSurfaceFormat) DepthBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat15depthBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStencilBufferSize(int)

/*
Set the preferred stencil buffer size to size bits.

See also stencilBufferSize().
*/
func (this *QSurfaceFormat) SetStencilBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat20setStencilBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] int stencilBufferSize() const

/*
Returns the stencil buffer size in bits.

See also setStencilBufferSize().
*/
func (this *QSurfaceFormat) StencilBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat17stencilBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRedBufferSize(int)

/*
Set the desired size in bits of the red channel of the color buffer.

Note: On Mac OSX, be sure to set the buffer size of all color channels, otherwise this setting will have no effect. If one of the buffer sizes is not set, the current bit-depth of the screen is used.

See also redBufferSize().
*/
func (this *QSurfaceFormat) SetRedBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat16setRedBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int redBufferSize() const

/*
Get the size in bits of the red channel of the color buffer.

See also setRedBufferSize().
*/
func (this *QSurfaceFormat) RedBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat13redBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGreenBufferSize(int)

/*
Set the desired size in bits of the green channel of the color buffer.

Note: On Mac OSX, be sure to set the buffer size of all color channels, otherwise this setting will have no effect. If one of the buffer sizes is not set, the current bit-depth of the screen is used.

See also greenBufferSize().
*/
func (this *QSurfaceFormat) SetGreenBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat18setGreenBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] int greenBufferSize() const

/*
Get the size in bits of the green channel of the color buffer.

See also setGreenBufferSize().
*/
func (this *QSurfaceFormat) GreenBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat15greenBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBlueBufferSize(int)

/*
Set the desired size in bits of the blue channel of the color buffer.

Note: On Mac OSX, be sure to set the buffer size of all color channels, otherwise this setting will have no effect. If one of the buffer sizes is not set, the current bit-depth of the screen is used.

See also blueBufferSize().
*/
func (this *QSurfaceFormat) SetBlueBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat17setBlueBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] int blueBufferSize() const

/*
Get the size in bits of the blue channel of the color buffer.

See also setBlueBufferSize().
*/
func (this *QSurfaceFormat) BlueBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat14blueBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlphaBufferSize(int)

/*
Set the desired size in bits of the alpha channel of the color buffer.

See also alphaBufferSize().
*/
func (this *QSurfaceFormat) SetAlphaBufferSize(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat18setAlphaBufferSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:113
// index:0
// Public Visibility=Default Availability=Available
// [4] int alphaBufferSize() const

/*
Get the size in bits of the alpha channel of the color buffer.

See also setAlphaBufferSize().
*/
func (this *QSurfaceFormat) AlphaBufferSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat15alphaBufferSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSamples(int)

/*
Set the preferred number of samples per pixel when multisampling is enabled to numSamples. By default, multisampling is disabled.

See also samples().
*/
func (this *QSurfaceFormat) SetSamples(numSamples int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat10setSamplesEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), numSamples)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:116
// index:0
// Public Visibility=Default Availability=Available
// [4] int samples() const

/*
Returns the number of samples per pixel when multisampling is enabled. By default, multisampling is disabled.

See also setSamples().
*/
func (this *QSurfaceFormat) Samples() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat7samplesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSwapBehavior(QSurfaceFormat::SwapBehavior)

/*
Set the swap behavior of the surface.

The swap behavior specifies whether single, double, or triple buffering is desired. The default, DefaultSwapBehavior, gives the default swap behavior of the platform.

See also swapBehavior().
*/
func (this *QSurfaceFormat) SetSwapBehavior(behavior int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat15setSwapBehaviorENS_12SwapBehaviorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), behavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:119
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurfaceFormat::SwapBehavior swapBehavior() const

/*
Returns the configured swap behaviour.

See also setSwapBehavior().
*/
func (this *QSurfaceFormat) SwapBehavior() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat12swapBehaviorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:121
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasAlpha() const

/*
Returns true if the alpha buffer size is greater than zero.

This means that the surface might be used with per pixel translucency effects.
*/
func (this *QSurfaceFormat) HasAlpha() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat8hasAlphaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProfile(QSurfaceFormat::OpenGLContextProfile)

/*
Sets the desired OpenGL context profile.

This setting is ignored if the requested OpenGL version is less than 3.2.

See also profile().
*/
func (this *QSurfaceFormat) SetProfile(profile int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat10setProfileENS_20OpenGLContextProfileE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), profile)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurfaceFormat::OpenGLContextProfile profile() const

/*
Get the configured OpenGL context profile.

This setting is ignored if the requested OpenGL version is less than 3.2.

See also setProfile().
*/
func (this *QSurfaceFormat) Profile() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat7profileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderableType(QSurfaceFormat::RenderableType)

/*
Sets the desired renderable type.

Chooses between desktop OpenGL, OpenGL ES, and OpenVG.

See also renderableType().
*/
func (this *QSurfaceFormat) SetRenderableType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat17setRenderableTypeENS_14RenderableTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurfaceFormat::RenderableType renderableType() const

/*
Gets the renderable type.

Chooses between desktop OpenGL, OpenGL ES, and OpenVG.

See also setRenderableType().
*/
func (this *QSurfaceFormat) RenderableType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat14renderableTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMajorVersion(int)

/*
Sets the desired major OpenGL version.

See also majorVersion().
*/
func (this *QSurfaceFormat) SetMajorVersion(majorVersion int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat15setMajorVersionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), majorVersion)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:130
// index:0
// Public Visibility=Default Availability=Available
// [4] int majorVersion() const

/*
Returns the major OpenGL version.

The default version is 2.0.

See also setMajorVersion().
*/
func (this *QSurfaceFormat) MajorVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat12majorVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinorVersion(int)

/*
Sets the desired minor OpenGL version.

The default version is 2.0.

See also minorVersion().
*/
func (this *QSurfaceFormat) SetMinorVersion(minorVersion int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat15setMinorVersionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minorVersion)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:133
// index:0
// Public Visibility=Default Availability=Available
// [4] int minorVersion() const

/*
Returns the minor OpenGL version.

See also setMinorVersion().
*/
func (this *QSurfaceFormat) MinorVersion() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat12minorVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVersion(int, int)

/*
Sets the desired major and minor OpenGL versions.

The default version is 2.0.

See also version().
*/
func (this *QSurfaceFormat) SetVersion(major int, minor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat10setVersionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), major, minor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:138
// index:0
// Public Visibility=Default Availability=Available
// [1] bool stereo() const

/*
Returns true if stereo buffering is enabled; otherwise returns false. Stereo buffering is disabled by default.

See also setStereo().
*/
func (this *QSurfaceFormat) Stereo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat6stereoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStereo(bool)

/*
If enable is true enables stereo buffering; otherwise disables stereo buffering.

Stereo buffering is disabled by default.

Stereo buffering provides extra color buffers to generate left-eye and right-eye images.

See also stereo().
*/
func (this *QSurfaceFormat) SetStereo(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat9setStereoEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOption(QSurfaceFormat::FormatOptions)

/*
Sets the format option option if on is true; otherwise, clears the option.

This function was introduced in  Qt 5.3.

See also setOptions(), options(), and testOption().
*/
func (this *QSurfaceFormat) SetOption(opt int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat9setOptionE6QFlagsINS_12FormatOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opt)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:147
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setOption(QSurfaceFormat::FormatOption, bool)

/*
Sets the format option option if on is true; otherwise, clears the option.

This function was introduced in  Qt 5.3.

See also setOptions(), options(), and testOption().
*/
func (this *QSurfaceFormat) SetOption_1(option int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat9setOptionENS_12FormatOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:147
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setOption(QSurfaceFormat::FormatOption, bool)

/*
Sets the format option option if on is true; otherwise, clears the option.

This function was introduced in  Qt 5.3.

See also setOptions(), options(), and testOption().
*/
func (this *QSurfaceFormat) SetOption_1_(option int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat9setOptionENS_12FormatOptionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:143
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testOption(QSurfaceFormat::FormatOptions) const

/*
Returns true if the format option option is set; otherwise returns false.

This function was introduced in  Qt 5.3.

See also options().
*/
func (this *QSurfaceFormat) TestOption(opt int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat10testOptionE6QFlagsINS_12FormatOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opt)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:148
// index:1
// Public Visibility=Default Availability=Available
// [1] bool testOption(QSurfaceFormat::FormatOption) const

/*
Returns true if the format option option is set; otherwise returns false.

This function was introduced in  Qt 5.3.

See also options().
*/
func (this *QSurfaceFormat) TestOption_1(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat10testOptionENS_12FormatOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptions(QSurfaceFormat::FormatOptions)

/*
Sets the format options to options.

This function was introduced in  Qt 5.3.

See also options() and testOption().
*/
func (this *QSurfaceFormat) SetOptions(options int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat10setOptionsE6QFlagsINS_12FormatOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), options)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:149
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurfaceFormat::FormatOptions options() const

/*
Returns the currently set format options.

This function was introduced in  Qt 5.3.

See also setOption(), setOptions(), and testOption().
*/
func (this *QSurfaceFormat) Options() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat7optionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:151
// index:0
// Public Visibility=Default Availability=Available
// [4] int swapInterval() const

/*
Returns the swap interval.

This function was introduced in  Qt 5.3.

See also setSwapInterval().
*/
func (this *QSurfaceFormat) SwapInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat12swapIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSwapInterval(int)

/*
Sets the preferred swap interval. The swap interval specifies the minimum number of video frames that are displayed before a buffer swap occurs. This can be used to sync the GL drawing into a window to the vertical refresh of the screen.

Setting an interval value of 0 will turn the vertical refresh syncing off, any value higher than 0 will turn the vertical syncing on. Setting interval to a higher value, for example 10, results in having 10 vertical retraces between every buffer swap.

The default interval is 1.

Changing the swap interval may not be supported by the underlying platform. In this case, the request will be silently ignored.

This function was introduced in  Qt 5.3.

See also swapInterval().
*/
func (this *QSurfaceFormat) SetSwapInterval(interval int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat15setSwapIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), interval)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:154
// index:0
// Public Visibility=Default Availability=Available
// [4] QSurfaceFormat::ColorSpace colorSpace() const

/*
Returns the color space.

This function was introduced in  Qt 5.10.

See also setColorSpace().
*/
func (this *QSurfaceFormat) ColorSpace() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSurfaceFormat10colorSpaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColorSpace(QSurfaceFormat::ColorSpace)

/*
Sets the preferred colorSpace.

For example, this allows requesting windows with default framebuffers that are sRGB-capable on platforms that support it.

Note: When the requested color space is not supported by the platform, the request is ignored. Query the QSurfaceFormat after window creation to verify if the color space request could be honored or not.

Note: This setting controls if the default framebuffer of the window is capable of updates and blending in a given color space. It does not change applications' output by itself. The applications' rendering code will still have to opt in via the appropriate OpenGL calls to enable updates and blending to be performed in the given color space instead of using the standard linear operations.

This function was introduced in  Qt 5.10.

See also colorSpace().
*/
func (this *QSurfaceFormat) SetColorSpace(colorSpace int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat13setColorSpaceENS_10ColorSpaceE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), colorSpace)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:157
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultFormat(const QSurfaceFormat &)

/*
Sets the global default surface format.

This format is used by default in QOpenGLContext, QWindow, QOpenGLWidget and similar classes.

It can always be overridden on a per-instance basis by using the class in question's own setFormat() function. However, it is often more convenient to set the format for all windows once at the start of the application. It also guarantees proper behavior in cases where shared contexts are required, because settings the format via this function guarantees that all contexts and surfaces, even the ones created internally by Qt, will use the same format.

Note: When setting Qt::AA_ShareOpenGLContexts, it is strongly recommended to place the call to this function before the construction of the QGuiApplication or QApplication. Otherwise format will not be applied to the global share context and therefore issues may arise with context sharing afterwards.

This function was introduced in  Qt 5.4.

See also defaultFormat().
*/
func (this *QSurfaceFormat) SetDefaultFormat(format QSurfaceFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QSurfaceFormat_PTR() != nil {
		convArg0 = format.QSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat16setDefaultFormatERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QSurfaceFormat_SetDefaultFormat(format QSurfaceFormat_ITF) {
	var nilthis *QSurfaceFormat
	nilthis.SetDefaultFormat(format)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:158
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSurfaceFormat defaultFormat()

/*
Returns the global default surface format.

When setDefaultFormat() is not called, this is a default-constructed QSurfaceFormat.

This function was introduced in  Qt 5.4.

See also setDefaultFormat().
*/
func (this *QSurfaceFormat) DefaultFormat() *QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSurfaceFormat13defaultFormatEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}
func QSurfaceFormat_DefaultFormat() *QSurfaceFormat /*123*/ {
	var nilthis *QSurfaceFormat
	rv := nilthis.DefaultFormat()
	return rv
}

/*


 */
type QSurfaceFormat__FormatOption = int

//
const QSurfaceFormat__StereoBuffers QSurfaceFormat__FormatOption = 1

//
const QSurfaceFormat__DebugContext QSurfaceFormat__FormatOption = 2

//
const QSurfaceFormat__DeprecatedFunctions QSurfaceFormat__FormatOption = 4

//
const QSurfaceFormat__ResetNotification QSurfaceFormat__FormatOption = 8

func (this *QSurfaceFormat) FormatOptionItemName(val int) string {
	switch val {
	case QSurfaceFormat__StereoBuffers: // 1
		return "StereoBuffers"
	case QSurfaceFormat__DebugContext: // 2
		return "DebugContext"
	case QSurfaceFormat__DeprecatedFunctions: // 4
		return "DeprecatedFunctions"
	case QSurfaceFormat__ResetNotification: // 8
		return "ResetNotification"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSurfaceFormat_FormatOptionItemName(val int) string {
	var nilthis *QSurfaceFormat
	return nilthis.FormatOptionItemName(val)
}

/*
This enum is used by QSurfaceFormat to specify the swap behaviour of a surface. The swap behaviour is mostly transparent to the application, but it affects factors such as rendering latency and throughput.


*/
type QSurfaceFormat__SwapBehavior = int

// The default, unspecified swap behaviour of the platform.
const QSurfaceFormat__DefaultSwapBehavior QSurfaceFormat__SwapBehavior = 0

// Used to request single buffering, which might result in flickering when OpenGL rendering is done directly to screen without an intermediate offscreen buffer.
const QSurfaceFormat__SingleBuffer QSurfaceFormat__SwapBehavior = 1

// This is typically the default swap behaviour on desktop platforms, consisting of one back buffer and one front buffer. Rendering is done to the back buffer, and then the back buffer and front buffer are swapped, or the contents of the back buffer are copied to the front buffer, depending on the implementation.
const QSurfaceFormat__DoubleBuffer QSurfaceFormat__SwapBehavior = 2

// This swap behaviour is sometimes used in order to decrease the risk of skipping a frame when the rendering rate is just barely keeping up with the screen refresh rate. Depending on the platform it might also lead to slightly more efficient use of the GPU due to improved pipelining behaviour. Triple buffering comes at the cost of an extra frame of memory usage and latency, and might not be supported depending on the underlying platform.
const QSurfaceFormat__TripleBuffer QSurfaceFormat__SwapBehavior = 3

func (this *QSurfaceFormat) SwapBehaviorItemName(val int) string {
	switch val {
	case QSurfaceFormat__DefaultSwapBehavior: // 0
		return "DefaultSwapBehavior"
	case QSurfaceFormat__SingleBuffer: // 1
		return "SingleBuffer"
	case QSurfaceFormat__DoubleBuffer: // 2
		return "DoubleBuffer"
	case QSurfaceFormat__TripleBuffer: // 3
		return "TripleBuffer"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSurfaceFormat_SwapBehaviorItemName(val int) string {
	var nilthis *QSurfaceFormat
	return nilthis.SwapBehaviorItemName(val)
}

/*
This enum specifies the rendering backend for the surface.


*/
type QSurfaceFormat__RenderableType = int

//
const QSurfaceFormat__DefaultRenderableType QSurfaceFormat__RenderableType = 0

//
const QSurfaceFormat__OpenGL QSurfaceFormat__RenderableType = 1

//
const QSurfaceFormat__OpenGLES QSurfaceFormat__RenderableType = 2

//
const QSurfaceFormat__OpenVG QSurfaceFormat__RenderableType = 4

func (this *QSurfaceFormat) RenderableTypeItemName(val int) string {
	switch val {
	case QSurfaceFormat__DefaultRenderableType: // 0
		return "DefaultRenderableType"
	case QSurfaceFormat__OpenGL: // 1
		return "OpenGL"
	case QSurfaceFormat__OpenGLES: // 2
		return "OpenGLES"
	case QSurfaceFormat__OpenVG: // 4
		return "OpenVG"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSurfaceFormat_RenderableTypeItemName(val int) string {
	var nilthis *QSurfaceFormat
	return nilthis.RenderableTypeItemName(val)
}

/*
This enum is used to specify the OpenGL context profile, in conjunction with QSurfaceFormat::setMajorVersion() and QSurfaceFormat::setMinorVersion().

Profiles are exposed in OpenGL 3.2 and above, and are used to choose between a restricted core profile, and a compatibility profile which might contain deprecated support functionality.

Note that the core profile might still contain functionality that is deprecated and scheduled for removal in a higher version. To get access to the deprecated functionality for the core profile in the set OpenGL version you can use the QSurfaceFormat format option QSurfaceFormat::DeprecatedFunctions.


*/
type QSurfaceFormat__OpenGLContextProfile = int

//
const QSurfaceFormat__NoProfile QSurfaceFormat__OpenGLContextProfile = 0

//
const QSurfaceFormat__CoreProfile QSurfaceFormat__OpenGLContextProfile = 1

// Functionality from earlier OpenGL versions is available.
const QSurfaceFormat__CompatibilityProfile QSurfaceFormat__OpenGLContextProfile = 2

func (this *QSurfaceFormat) OpenGLContextProfileItemName(val int) string {
	switch val {
	case QSurfaceFormat__NoProfile: // 0
		return "NoProfile"
	case QSurfaceFormat__CoreProfile: // 1
		return "CoreProfile"
	case QSurfaceFormat__CompatibilityProfile: // 2
		return "CompatibilityProfile"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSurfaceFormat_OpenGLContextProfileItemName(val int) string {
	var nilthis *QSurfaceFormat
	return nilthis.OpenGLContextProfileItemName(val)
}

/*
This enum is used to specify the preferred color space, controlling if the window's associated default framebuffer is able to do updates and blending in a given encoding instead of the standard linear operations.


*/
type QSurfaceFormat__ColorSpace = int

// The default, unspecified color space.
const QSurfaceFormat__DefaultColorSpace QSurfaceFormat__ColorSpace = 0

// When GL_ARB_framebuffer_sRGB or GL_EXT_framebuffer_sRGB is supported by the platform and this value is set, the window will be created with an sRGB-capable default framebuffer. Note that some platforms may return windows with a sRGB-capable default framebuffer even when not requested explicitly.
const QSurfaceFormat__sRGBColorSpace QSurfaceFormat__ColorSpace = 1

func (this *QSurfaceFormat) ColorSpaceItemName(val int) string {
	switch val {
	case QSurfaceFormat__DefaultColorSpace: // 0
		return "DefaultColorSpace"
	case QSurfaceFormat__sRGBColorSpace: // 1
		return "sRGBColorSpace"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSurfaceFormat_ColorSpaceItemName(val int) string {
	var nilthis *QSurfaceFormat
	return nilthis.ColorSpaceItemName(val)
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
}

//  keep block end
