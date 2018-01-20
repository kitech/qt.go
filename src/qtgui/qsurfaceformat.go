//  header block begin
// /usr/include/qt/QtGui/qsurfaceformat.h
// #include <qsurfaceformat.h>
// #include <QtGui>
package qtgui

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	return this.Cthis
}

// /usr/include/qt/QtGui/qsurfaceformat.h:94
// index:0
// void QSurfaceFormat()
func NewQSurfaceFormat() *QSurfaceFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSurfaceFormatFromPointer(cthis)
	return gothis
}
func NewQSurfaceFormatFromPointer(cthis unsafe.Pointer) *QSurfaceFormat {
	return &QSurfaceFormat{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qsurfaceformat.h:95
// index:1
// void QSurfaceFormat(QSurfaceFormat::FormatOptions)
func NewQSurfaceFormat_1(options int) *QSurfaceFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormatC2E6QFlagsINS_12FormatOptionEE", ffiqt.FFI_TYPE_VOID, cthis, options)
	gopp.ErrPrint(err, rv)
	gothis := NewQSurfaceFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qsurfaceformat.h:98
// index:0
// void ~QSurfaceFormat()
func DeleteQSurfaceFormat(*QSurfaceFormat) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormatD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:100
// index:0
// void setDepthBufferSize(int)
func (this *QSurfaceFormat) SetDepthBufferSize(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat18setDepthBufferSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:101
// index:0
// int depthBufferSize()
func (this *QSurfaceFormat) DepthBufferSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat15depthBufferSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:103
// index:0
// void setStencilBufferSize(int)
func (this *QSurfaceFormat) SetStencilBufferSize(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat20setStencilBufferSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:104
// index:0
// int stencilBufferSize()
func (this *QSurfaceFormat) StencilBufferSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat17stencilBufferSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:106
// index:0
// void setRedBufferSize(int)
func (this *QSurfaceFormat) SetRedBufferSize(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat16setRedBufferSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:107
// index:0
// int redBufferSize()
func (this *QSurfaceFormat) RedBufferSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat13redBufferSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:108
// index:0
// void setGreenBufferSize(int)
func (this *QSurfaceFormat) SetGreenBufferSize(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat18setGreenBufferSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:109
// index:0
// int greenBufferSize()
func (this *QSurfaceFormat) GreenBufferSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat15greenBufferSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:110
// index:0
// void setBlueBufferSize(int)
func (this *QSurfaceFormat) SetBlueBufferSize(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat17setBlueBufferSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:111
// index:0
// int blueBufferSize()
func (this *QSurfaceFormat) BlueBufferSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat14blueBufferSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:112
// index:0
// void setAlphaBufferSize(int)
func (this *QSurfaceFormat) SetAlphaBufferSize(size int) {
	// 0: (, size int), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat18setAlphaBufferSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:113
// index:0
// int alphaBufferSize()
func (this *QSurfaceFormat) AlphaBufferSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat15alphaBufferSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:115
// index:0
// void setSamples(int)
func (this *QSurfaceFormat) SetSamples(numSamples int) {
	// 0: (, numSamples int), (&numSamples)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat10setSamplesEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &numSamples)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:116
// index:0
// int samples()
func (this *QSurfaceFormat) Samples() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat7samplesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:118
// index:0
// void setSwapBehavior(enum QSurfaceFormat::SwapBehavior)
func (this *QSurfaceFormat) SetSwapBehavior(behavior int) {
	// 0: (, behavior QSurfaceFormat::SwapBehavior), (&behavior)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat15setSwapBehaviorENS_12SwapBehaviorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &behavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:119
// index:0
// QSurfaceFormat::SwapBehavior swapBehavior()
func (this *QSurfaceFormat) SwapBehavior() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat12swapBehaviorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:121
// index:0
// bool hasAlpha()
func (this *QSurfaceFormat) HasAlpha() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat8hasAlphaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:123
// index:0
// void setProfile(enum QSurfaceFormat::OpenGLContextProfile)
func (this *QSurfaceFormat) SetProfile(profile int) {
	// 0: (, profile QSurfaceFormat::OpenGLContextProfile), (&profile)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat10setProfileENS_20OpenGLContextProfileE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &profile)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:124
// index:0
// QSurfaceFormat::OpenGLContextProfile profile()
func (this *QSurfaceFormat) Profile() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat7profileEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:126
// index:0
// void setRenderableType(enum QSurfaceFormat::RenderableType)
func (this *QSurfaceFormat) SetRenderableType(type_ int) {
	// 0: (, type QSurfaceFormat::RenderableType), (&type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat17setRenderableTypeENS_14RenderableTypeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:127
// index:0
// QSurfaceFormat::RenderableType renderableType()
func (this *QSurfaceFormat) RenderableType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat14renderableTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:129
// index:0
// void setMajorVersion(int)
func (this *QSurfaceFormat) SetMajorVersion(majorVersion int) {
	// 0: (, majorVersion int), (&majorVersion)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat15setMajorVersionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &majorVersion)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:130
// index:0
// int majorVersion()
func (this *QSurfaceFormat) MajorVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat12majorVersionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:132
// index:0
// void setMinorVersion(int)
func (this *QSurfaceFormat) SetMinorVersion(minorVersion int) {
	// 0: (, minorVersion int), (&minorVersion)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat15setMinorVersionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &minorVersion)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:133
// index:0
// int minorVersion()
func (this *QSurfaceFormat) MinorVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat12minorVersionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:135
// index:0
// QPair<int, int> version()
func (this *QSurfaceFormat) Version() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat7versionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:136
// index:0
// void setVersion(int, int)
func (this *QSurfaceFormat) SetVersion(major int, minor int) {
	// 0: (, major int, minor int), (&major, &minor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat10setVersionEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &major, &minor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:138
// index:0
// bool stereo()
func (this *QSurfaceFormat) Stereo() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat6stereoEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:139
// index:0
// void setStereo(_Bool)
func (this *QSurfaceFormat) SetStereo(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat9setStereoEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:142
// index:0
// void setOption(class QSurfaceFormat::FormatOptions)
func (this *QSurfaceFormat) SetOption(opt int) {
	// 0: (, QFlags<QSurfaceFormat::FormatOption> opt), (&opt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat9setOptionE6QFlagsINS_12FormatOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &opt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:147
// index:1
// void setOption(enum QSurfaceFormat::FormatOption, _Bool)
func (this *QSurfaceFormat) SetOption_1(option int, on bool) {
	// 1: (, option QSurfaceFormat::FormatOption, on bool), (&option, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat9setOptionENS_12FormatOptionEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:143
// index:0
// bool testOption(class QSurfaceFormat::FormatOptions)
func (this *QSurfaceFormat) TestOption(opt int) {
	// 0: (, QFlags<QSurfaceFormat::FormatOption> opt), (&opt)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat10testOptionE6QFlagsINS_12FormatOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &opt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:148
// index:1
// bool testOption(enum QSurfaceFormat::FormatOption)
func (this *QSurfaceFormat) TestOption_1(option int) {
	// 1: (, option QSurfaceFormat::FormatOption), (&option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat10testOptionENS_12FormatOptionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:146
// index:0
// void setOptions(class QSurfaceFormat::FormatOptions)
func (this *QSurfaceFormat) SetOptions(options int) {
	// 0: (, QFlags<QSurfaceFormat::FormatOption> options), (&options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat10setOptionsE6QFlagsINS_12FormatOptionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &options)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:149
// index:0
// QSurfaceFormat::FormatOptions options()
func (this *QSurfaceFormat) Options() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat7optionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:151
// index:0
// int swapInterval()
func (this *QSurfaceFormat) SwapInterval() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat12swapIntervalEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:152
// index:0
// void setSwapInterval(int)
func (this *QSurfaceFormat) SetSwapInterval(interval int) {
	// 0: (, interval int), (&interval)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat15setSwapIntervalEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &interval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:154
// index:0
// QSurfaceFormat::ColorSpace colorSpace()
func (this *QSurfaceFormat) ColorSpace() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QSurfaceFormat10colorSpaceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:155
// index:0
// void setColorSpace(enum QSurfaceFormat::ColorSpace)
func (this *QSurfaceFormat) SetColorSpace(colorSpace int) {
	// 0: (, colorSpace QSurfaceFormat::ColorSpace), (&colorSpace)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat13setColorSpaceENS_10ColorSpaceE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &colorSpace)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:157
// index:0
// static
// void setDefaultFormat(const class QSurfaceFormat &)
func (this *QSurfaceFormat) SetDefaultFormat(format unsafe.Pointer) {
	// 0: (format const QSurfaceFormat &), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat16setDefaultFormatERKS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QSurfaceFormat_SetDefaultFormat(format unsafe.Pointer) {
	// 0: (format const QSurfaceFormat &), (format)
	var nilthis *QSurfaceFormat
	nilthis.SetDefaultFormat(format)
}

// /usr/include/qt/QtGui/qsurfaceformat.h:158
// index:0
// static
// QSurfaceFormat defaultFormat()
func (this *QSurfaceFormat) DefaultFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSurfaceFormat13defaultFormatEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QSurfaceFormat_DefaultFormat() {
	// 0: (), ()
	var nilthis *QSurfaceFormat
	nilthis.DefaultFormat()
}

//  body block end
