package qtmultimedia

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h
// #include <qvideowindowcontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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

/*

 */
type QVideoWindowControl struct {
	*QMediaControl
}
type QVideoWindowControl_ITF interface {
	QMediaControl_ITF
	QVideoWindowControl_PTR() *QVideoWindowControl
}

func (ptr *QVideoWindowControl) QVideoWindowControl_PTR() *QVideoWindowControl { return ptr }

func (this *QVideoWindowControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QVideoWindowControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQVideoWindowControlFromPointer(cthis unsafe.Pointer) *QVideoWindowControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QVideoWindowControl{bcthis0}
}
func (*QVideoWindowControl) NewFromPointer(cthis unsafe.Pointer) *QVideoWindowControl {
	return NewQVideoWindowControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QVideoWindowControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoWindowControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QVideoWindowControl()

/*

 */
func DeleteQVideoWindowControl(this *QVideoWindowControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] WId winId() const

/*
Returns the ID of the window a video overlay end point renders to.

See also setWinId().
*/
func (this *QVideoWindowControl) WinId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoWindowControl5winIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setWinId(WId)

/*
Sets the id of the window a video overlay end point renders to.

See also winId().
*/
func (this *QVideoWindowControl) SetWinId(id uint64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl8setWinIdEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QRect displayRect() const

/*
Returns the sub-rect of a window where video is displayed.

See also setDisplayRect().
*/
func (this *QVideoWindowControl) DisplayRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoWindowControl11displayRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setDisplayRect(const QRect &)

/*
Sets the sub-rect of a window where video is displayed.

See also displayRect().
*/
func (this *QVideoWindowControl) SetDisplayRect(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl14setDisplayRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isFullScreen() const

/*
Identifies if a video overlay is a fullScreen overlay.

Returns true if the video overlay is fullScreen, and false otherwise.
*/
func (this *QVideoWindowControl) IsFullScreen() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoWindowControl12isFullScreenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setFullScreen(bool)

/*
Sets whether a video overlay is a fullScreen overlay.

See also isFullScreen().
*/
func (this *QVideoWindowControl) SetFullScreen(fullScreen bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl13setFullScreenEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fullScreen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:69
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void repaint()

/*
Repaints the last frame.
*/
func (this *QVideoWindowControl) Repaint() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl7repaintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSize nativeSize() const

/*
Returns a suggested size for the video display based on the resolution and aspect ratio of the video.
*/
func (this *QVideoWindowControl) NativeSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoWindowControl10nativeSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:73
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] Qt::AspectRatioMode aspectRatioMode() const

/*
Returns how video is scaled to fit the display region with respect to its aspect ratio.

See also setAspectRatioMode().
*/
func (this *QVideoWindowControl) AspectRatioMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoWindowControl15aspectRatioModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setAspectRatioMode(Qt::AspectRatioMode)

/*
Sets the aspect ratio mode which determines how video is scaled to the fit the display region with respect to its aspect ratio.

See also aspectRatioMode().
*/
func (this *QVideoWindowControl) SetAspectRatioMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl18setAspectRatioModeEN2Qt15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int brightness() const

/*
Returns the brightness adjustment applied to a video overlay.

Valid brightness values range between -100 and 100, the default is 0.

See also setBrightness().
*/
func (this *QVideoWindowControl) Brightness() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoWindowControl10brightnessEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:77
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setBrightness(int)

/*
Sets a brightness adjustment for a video overlay.

Valid brightness values range between -100 and 100, the default is 0.

See also brightness().
*/
func (this *QVideoWindowControl) SetBrightness(brightness int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl13setBrightnessEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), brightness)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int contrast() const

/*
Returns the contrast adjustment applied to a video overlay.

Valid contrast values range between -100 and 100, the default is 0.

See also setContrast().
*/
func (this *QVideoWindowControl) Contrast() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoWindowControl8contrastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setContrast(int)

/*
Sets the contrast adjustment for a video overlay.

Valid contrast values range between -100 and 100, the default is 0.

See also contrast().
*/
func (this *QVideoWindowControl) SetContrast(contrast int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl11setContrastEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), contrast)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int hue() const

/*
Returns the hue adjustment applied to a video overlay.

Value hue values range between -100 and 100, the default is 0.

See also setHue().
*/
func (this *QVideoWindowControl) Hue() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoWindowControl3hueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:83
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setHue(int)

/*
Sets a hue adjustment for a video overlay.

Valid hue values range between -100 and 100, the default is 0.

See also hue().
*/
func (this *QVideoWindowControl) SetHue(hue int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl6setHueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hue)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:85
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int saturation() const

/*
Returns the saturation adjustment applied to a video overlay.

Value saturation values range between -100 and 100, the default is 0.

See also setSaturation().
*/
func (this *QVideoWindowControl) Saturation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QVideoWindowControl10saturationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:86
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setSaturation(int)

/*
Sets a saturation adjustment for a video overlay.

Valid saturation values range between -100 and 100, the default is 0.

See also saturation().
*/
func (this *QVideoWindowControl) SetSaturation(saturation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl13setSaturationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), saturation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void fullScreenChanged(bool)

/*
Signals that the fullScreen state of a video overlay has changed.
*/
func (this *QVideoWindowControl) FullScreenChanged(fullScreen bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl17fullScreenChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fullScreen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void brightnessChanged(int)

/*
Signals that a video overlay's brightness adjustment has changed.
*/
func (this *QVideoWindowControl) BrightnessChanged(brightness int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl17brightnessChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), brightness)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void contrastChanged(int)

/*
Signals that a video overlay's contrast adjustment has changed.
*/
func (this *QVideoWindowControl) ContrastChanged(contrast int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl15contrastChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), contrast)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hueChanged(int)

/*
Signals that a video overlay's hue adjustment has changed.
*/
func (this *QVideoWindowControl) HueChanged(hue int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl10hueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hue)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void saturationChanged(int)

/*
Signals that a video overlay's saturation adjustment has changed.
*/
func (this *QVideoWindowControl) SaturationChanged(saturation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl17saturationChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), saturation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void nativeSizeChanged()

/*
Signals that the native dimensions of the video have changed.
*/
func (this *QVideoWindowControl) NativeSizeChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControl17nativeSizeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:97
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QVideoWindowControl(QObject *)

/*
Constructs a new video window control with the given parent.
*/
func (*QVideoWindowControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QVideoWindowControl {
	return NewQVideoWindowControl(parent)
}
func NewQVideoWindowControl(parent qtcore.QObject_ITF /*777 QObject **/) *QVideoWindowControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoWindowControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVideoWindowControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qvideowindowcontrol.h:97
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QVideoWindowControl(QObject *)

/*
Constructs a new video window control with the given parent.
*/
func (*QVideoWindowControl) NewForInheritp() *QVideoWindowControl {
	return NewQVideoWindowControlp()
}
func NewQVideoWindowControlp() *QVideoWindowControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QVideoWindowControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQVideoWindowControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QVideoWindowControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11917() {
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
