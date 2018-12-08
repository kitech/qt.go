package qtmultimedia

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h
// #include <qabstractvideosurface.h>
// #include <QtMultimedia>

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
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

// void setError(QAbstractVideoSurface::Error)
func (this *QAbstractVideoSurface) InheritSetError(f func(error int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setError", f)
}

// void setNativeResolution(const QSize &)
func (this *QAbstractVideoSurface) InheritSetNativeResolution(f func(resolution *qtcore.QSize) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setNativeResolution", f)
}

/*

 */
type QAbstractVideoSurface struct {
	*qtcore.QObject
}
type QAbstractVideoSurface_ITF interface {
	qtcore.QObject_ITF
	QAbstractVideoSurface_PTR() *QAbstractVideoSurface
}

func (ptr *QAbstractVideoSurface) QAbstractVideoSurface_PTR() *QAbstractVideoSurface { return ptr }

func (this *QAbstractVideoSurface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractVideoSurface) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAbstractVideoSurfaceFromPointer(cthis unsafe.Pointer) *QAbstractVideoSurface {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractVideoSurface{bcthis0}
}
func (*QAbstractVideoSurface) NewFromPointer(cthis unsafe.Pointer) *QAbstractVideoSurface {
	return NewQAbstractVideoSurfaceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractVideoSurface) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractVideoSurface10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractVideoSurface(QObject *)

/*
Constructs a video surface with the given parent.
*/
func (*QAbstractVideoSurface) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QAbstractVideoSurface {
	return NewQAbstractVideoSurface(parent)
}
func NewQAbstractVideoSurface(parent qtcore.QObject_ITF /*777 QObject **/) *QAbstractVideoSurface {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurfaceC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractVideoSurfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractVideoSurface")
	return gothis
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractVideoSurface(QObject *)

/*
Constructs a video surface with the given parent.
*/
func (*QAbstractVideoSurface) NewForInheritp() *QAbstractVideoSurface {
	return NewQAbstractVideoSurfacep()
}
func NewQAbstractVideoSurfacep() *QAbstractVideoSurface {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurfaceC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractVideoSurfaceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractVideoSurface")
	return gothis
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractVideoSurface()

/*

 */
func DeleteQAbstractVideoSurface(this *QAbstractVideoSurface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isFormatSupported(const QVideoSurfaceFormat &) const

/*
Tests a video surface format to determine if a surface can accept it.

Returns true if the format is supported by the surface, and false otherwise.
*/
func (this *QAbstractVideoSurface) IsFormatSupported(format QVideoSurfaceFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if format != nil && format.QVideoSurfaceFormat_PTR() != nil {
		convArg0 = format.QVideoSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractVideoSurface17isFormatSupportedERK19QVideoSurfaceFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QVideoSurfaceFormat nearestFormat(const QVideoSurfaceFormat &) const

/*
Returns a supported video surface format that is similar to format.

A similar surface format is one that has the same pixel format and handle type but may differ in some of the other properties. For example, if there are restrictions on the frame sizes a video surface can accept it may suggest a format with a larger frame size and a viewport the size of the original frame size.

If the format is already supported it will be returned unchanged, or if there is no similar supported format an invalid format will be returned.
*/
func (this *QAbstractVideoSurface) NearestFormat(format QVideoSurfaceFormat_ITF) *QVideoSurfaceFormat /*123*/ {
	var convArg0 unsafe.Pointer
	if format != nil && format.QVideoSurfaceFormat_PTR() != nil {
		convArg0 = format.QVideoSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractVideoSurface13nearestFormatERK19QVideoSurfaceFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVideoSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVideoSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QVideoSurfaceFormat surfaceFormat() const

/*
Returns the format of a video surface.
*/
func (this *QAbstractVideoSurface) SurfaceFormat() *QVideoSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractVideoSurface13surfaceFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVideoSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVideoSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize nativeResolution() const

/*

 */
func (this *QAbstractVideoSurface) NativeResolution() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractVideoSurface16nativeResolutionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool start(const QVideoSurfaceFormat &)

/*
Starts a video surface presenting format frames.

Returns true if the surface was started, and false if an error occurred.

Note: You must call the base class implementation of start() at the end of your implementation.

See also isActive() and stop().
*/
func (this *QAbstractVideoSurface) Start(format QVideoSurfaceFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if format != nil && format.QVideoSurfaceFormat_PTR() != nil {
		convArg0 = format.QVideoSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurface5startERK19QVideoSurfaceFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void stop()

/*
Stops a video surface presenting frames and releases any resources acquired in start().

Note: You must call the base class implementation of stop() at the start of your implementation.

See also isActive() and start().
*/
func (this *QAbstractVideoSurface) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurface4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive() const

/*
Indicates whether a video surface has been started.

Returns true if the surface has been started, and false otherwise.
*/
func (this *QAbstractVideoSurface) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractVideoSurface8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:84
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool present(const QVideoFrame &)

/*
Presents a video frame.

Returns true if the frame was presented, and false if an error occurred.

Not all surfaces will block until the presentation of a frame has completed. Calling present() on a non-blocking surface may fail if called before the presentation of a previous frame has completed. In such cases the surface may not return to a ready state until it has had an opportunity to process events.

If present() fails for any other reason the surface should immediately enter the stopped state and an error() value will be set.

A video surface must be in the started state for present() to succeed, and the format of the video frame must be compatible with the current video surface format.

See also error().
*/
func (this *QAbstractVideoSurface) Present(frame QVideoFrame_ITF) bool {
	var convArg0 unsafe.Pointer
	if frame != nil && frame.QVideoFrame_PTR() != nil {
		convArg0 = frame.QVideoFrame_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurface7presentERK11QVideoFrame", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] QAbstractVideoSurface::Error error() const

/*
Returns the last error that occurred.

If a surface fails to start(), or stops unexpectedly this function can be called to discover what error occurred.

See also setError().
*/
func (this *QAbstractVideoSurface) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QAbstractVideoSurface5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeChanged(bool)

/*
Signals that the active state of a video surface has changed.

See also isActive(), start(), and stop().
*/
func (this *QAbstractVideoSurface) ActiveChanged(active bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurface13activeChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), active)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void surfaceFormatChanged(const QVideoSurfaceFormat &)

/*
Signals that the configured format of a video surface has changed.

See also surfaceFormat() and start().
*/
func (this *QAbstractVideoSurface) SurfaceFormatChanged(format QVideoSurfaceFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QVideoSurfaceFormat_PTR() != nil {
		convArg0 = format.QVideoSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurface20surfaceFormatChangedERK19QVideoSurfaceFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void supportedFormatsChanged()

/*
Signals that the set of formats supported by a video surface has changed.

See also supportedPixelFormats() and isFormatSupported().
*/
func (this *QAbstractVideoSurface) SupportedFormatsChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurface23supportedFormatsChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void nativeResolutionChanged(const QSize &)

/*
Signals the native resolution of video surface has changed.

Note: Notifier signal for property nativeResolution.
*/
func (this *QAbstractVideoSurface) NativeResolutionChanged(arg0 qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSize_PTR() != nil {
		convArg0 = arg0.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurface23nativeResolutionChangedERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:95
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setError(QAbstractVideoSurface::Error)

/*
Sets the value of error() to error.

This can be called by implementors of this interface to communicate what the most recent error was.

See also error().
*/
func (this *QAbstractVideoSurface) SetError(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurface8setErrorENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qabstractvideosurface.h:96
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setNativeResolution(const QSize &)

/*
Set the video surface native resolution.

This function can be called by implementors of this interface to specify to frame producers what the native resolution of this surface is.

See also nativeResolution().
*/
func (this *QAbstractVideoSurface) SetNativeResolution(resolution qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if resolution != nil && resolution.QSize_PTR() != nil {
		convArg0 = resolution.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QAbstractVideoSurface19setNativeResolutionERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the errors that may be returned by the error() function.


*/
type QAbstractVideoSurface__Error = int

// No error occurred.
const QAbstractVideoSurface__NoError QAbstractVideoSurface__Error = 0

// A video format was not supported.
const QAbstractVideoSurface__UnsupportedFormatError QAbstractVideoSurface__Error = 1

// A video frame was not compatible with the format of the surface.
const QAbstractVideoSurface__IncorrectFormatError QAbstractVideoSurface__Error = 2

// The surface has not been started.
const QAbstractVideoSurface__StoppedError QAbstractVideoSurface__Error = 3

// The surface could not allocate some resource.
const QAbstractVideoSurface__ResourceError QAbstractVideoSurface__Error = 4

func (this *QAbstractVideoSurface) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractVideoSurface_ErrorItemName(val int) string {
	var nilthis *QAbstractVideoSurface
	return nilthis.ErrorItemName(val)
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
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
