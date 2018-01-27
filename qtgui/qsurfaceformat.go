package qtgui

// /usr/include/qt/QtGui/qsurfaceformat.h
// #include <qsurfaceformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin
type QSurfaceFormat struct {
	*qtrt.CObject
}

func (this *QSurfaceFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSurfaceFormat) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSurfaceFormatFromPointer(cthis unsafe.Pointer) *QSurfaceFormat {
	return &QSurfaceFormat{&qtrt.CObject{cthis}}
}
func (*QSurfaceFormat) NewFromPointer(cthis unsafe.Pointer) *QSurfaceFormat {
	return NewQSurfaceFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:94
// index:0
// Public
// void QSurfaceFormat()
func NewQSurfaceFormat() *QSurfaceFormat {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSurfaceFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qsurfaceformat.h:98
// index:0
// Public
// void ~QSurfaceFormat()
func DeleteQSurfaceFormat(*QSurfaceFormat) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormatD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:100
// index:0
// Public
// void setDepthBufferSize(int)
func (this *QSurfaceFormat) SetDepthBufferSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat18setDepthBufferSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:101
// index:0
// Public
// int depthBufferSize()
func (this *QSurfaceFormat) DepthBufferSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat15depthBufferSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:103
// index:0
// Public
// void setStencilBufferSize(int)
func (this *QSurfaceFormat) SetStencilBufferSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat20setStencilBufferSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:104
// index:0
// Public
// int stencilBufferSize()
func (this *QSurfaceFormat) StencilBufferSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat17stencilBufferSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:106
// index:0
// Public
// void setRedBufferSize(int)
func (this *QSurfaceFormat) SetRedBufferSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat16setRedBufferSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:107
// index:0
// Public
// int redBufferSize()
func (this *QSurfaceFormat) RedBufferSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat13redBufferSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:108
// index:0
// Public
// void setGreenBufferSize(int)
func (this *QSurfaceFormat) SetGreenBufferSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat18setGreenBufferSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:109
// index:0
// Public
// int greenBufferSize()
func (this *QSurfaceFormat) GreenBufferSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat15greenBufferSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:110
// index:0
// Public
// void setBlueBufferSize(int)
func (this *QSurfaceFormat) SetBlueBufferSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat17setBlueBufferSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:111
// index:0
// Public
// int blueBufferSize()
func (this *QSurfaceFormat) BlueBufferSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat14blueBufferSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:112
// index:0
// Public
// void setAlphaBufferSize(int)
func (this *QSurfaceFormat) SetAlphaBufferSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat18setAlphaBufferSizeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:113
// index:0
// Public
// int alphaBufferSize()
func (this *QSurfaceFormat) AlphaBufferSize() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat15alphaBufferSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:115
// index:0
// Public
// void setSamples(int)
func (this *QSurfaceFormat) SetSamples(numSamples int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat10setSamplesEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), numSamples)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:116
// index:0
// Public
// int samples()
func (this *QSurfaceFormat) Samples() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat7samplesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:118
// index:0
// Public
// void setSwapBehavior(QSurfaceFormat::SwapBehavior)
func (this *QSurfaceFormat) SetSwapBehavior(behavior int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat15setSwapBehaviorENS_12SwapBehaviorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), behavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:119
// index:0
// Public
// QSurfaceFormat::SwapBehavior swapBehavior()
func (this *QSurfaceFormat) SwapBehavior() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat12swapBehaviorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:121
// index:0
// Public
// bool hasAlpha()
func (this *QSurfaceFormat) HasAlpha() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat8hasAlphaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:123
// index:0
// Public
// void setProfile(QSurfaceFormat::OpenGLContextProfile)
func (this *QSurfaceFormat) SetProfile(profile int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat10setProfileENS_20OpenGLContextProfileE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), profile)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:124
// index:0
// Public
// QSurfaceFormat::OpenGLContextProfile profile()
func (this *QSurfaceFormat) Profile() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat7profileEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:126
// index:0
// Public
// void setRenderableType(QSurfaceFormat::RenderableType)
func (this *QSurfaceFormat) SetRenderableType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat17setRenderableTypeENS_14RenderableTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:127
// index:0
// Public
// QSurfaceFormat::RenderableType renderableType()
func (this *QSurfaceFormat) RenderableType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat14renderableTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:129
// index:0
// Public
// void setMajorVersion(int)
func (this *QSurfaceFormat) SetMajorVersion(majorVersion int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat15setMajorVersionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), majorVersion)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:130
// index:0
// Public
// int majorVersion()
func (this *QSurfaceFormat) MajorVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat12majorVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:132
// index:0
// Public
// void setMinorVersion(int)
func (this *QSurfaceFormat) SetMinorVersion(minorVersion int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat15setMinorVersionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), minorVersion)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:133
// index:0
// Public
// int minorVersion()
func (this *QSurfaceFormat) MinorVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat12minorVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:136
// index:0
// Public
// void setVersion(int, int)
func (this *QSurfaceFormat) SetVersion(major int, minor int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat10setVersionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), major, minor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:138
// index:0
// Public
// bool stereo()
func (this *QSurfaceFormat) Stereo() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat6stereoEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:139
// index:0
// Public
// void setStereo(bool)
func (this *QSurfaceFormat) SetStereo(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat9setStereoEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:147
// index:0
// Public
// void setOption(QSurfaceFormat::FormatOption, bool)
func (this *QSurfaceFormat) SetOption(option int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat9setOptionENS_12FormatOptionEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:148
// index:0
// Public
// bool testOption(QSurfaceFormat::FormatOption)
func (this *QSurfaceFormat) TestOption(option int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat10testOptionENS_12FormatOptionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:149
// index:0
// Public
// QSurfaceFormat::FormatOptions options()
func (this *QSurfaceFormat) Options() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat7optionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:151
// index:0
// Public
// int swapInterval()
func (this *QSurfaceFormat) SwapInterval() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat12swapIntervalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qsurfaceformat.h:152
// index:0
// Public
// void setSwapInterval(int)
func (this *QSurfaceFormat) SetSwapInterval(interval int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat15setSwapIntervalEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:154
// index:0
// Public
// QSurfaceFormat::ColorSpace colorSpace()
func (this *QSurfaceFormat) ColorSpace() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat10colorSpaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:155
// index:0
// Public
// void setColorSpace(QSurfaceFormat::ColorSpace)
func (this *QSurfaceFormat) SetColorSpace(colorSpace int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat13setColorSpaceENS_10ColorSpaceE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), colorSpace)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:157
// index:0
// Public static
// void setDefaultFormat(const QSurfaceFormat &)
func (this *QSurfaceFormat) SetDefaultFormat(format *QSurfaceFormat) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat16setDefaultFormatERKS_", ffiqt.FFI_TYPE_POINTER, format)
	gopp.ErrPrint(err, rv)
}
func QSurfaceFormat_SetDefaultFormat(format *QSurfaceFormat) {
	var nilthis *QSurfaceFormat
	nilthis.SetDefaultFormat(format)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:158
// index:0
// Public static
// QSurfaceFormat defaultFormat()
func (this *QSurfaceFormat) DefaultFormat() *QSurfaceFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat13defaultFormatEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QSurfaceFormat_DefaultFormat() *QSurfaceFormat /*123*/ {
	var nilthis *QSurfaceFormat
	rv := nilthis.DefaultFormat()
	return rv
}

type QSurfaceFormat__FormatOption = int

const QSurfaceFormat__StereoBuffers QSurfaceFormat__FormatOption = 1
const QSurfaceFormat__DebugContext QSurfaceFormat__FormatOption = 2
const QSurfaceFormat__DeprecatedFunctions QSurfaceFormat__FormatOption = 4
const QSurfaceFormat__ResetNotification QSurfaceFormat__FormatOption = 8

type QSurfaceFormat__SwapBehavior = int

const QSurfaceFormat__DefaultSwapBehavior QSurfaceFormat__SwapBehavior = 0
const QSurfaceFormat__SingleBuffer QSurfaceFormat__SwapBehavior = 1
const QSurfaceFormat__DoubleBuffer QSurfaceFormat__SwapBehavior = 2
const QSurfaceFormat__TripleBuffer QSurfaceFormat__SwapBehavior = 3

type QSurfaceFormat__RenderableType = int

const QSurfaceFormat__DefaultRenderableType QSurfaceFormat__RenderableType = 0
const QSurfaceFormat__OpenGL QSurfaceFormat__RenderableType = 1
const QSurfaceFormat__OpenGLES QSurfaceFormat__RenderableType = 2
const QSurfaceFormat__OpenVG QSurfaceFormat__RenderableType = 4

type QSurfaceFormat__OpenGLContextProfile = int

const QSurfaceFormat__NoProfile QSurfaceFormat__OpenGLContextProfile = 0
const QSurfaceFormat__CoreProfile QSurfaceFormat__OpenGLContextProfile = 1
const QSurfaceFormat__CompatibilityProfile QSurfaceFormat__OpenGLContextProfile = 2

type QSurfaceFormat__ColorSpace = int

const QSurfaceFormat__DefaultColorSpace QSurfaceFormat__ColorSpace = 0
const QSurfaceFormat__sRGBColorSpace QSurfaceFormat__ColorSpace = 1

//  body block end
