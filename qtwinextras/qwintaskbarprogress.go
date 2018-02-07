package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h
// #include <qwintaskbarprogress.h>
// #include <Qtwinextras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin

type QWinTaskbarProgress struct {
	*qtcore.QObject
}

func (this *QWinTaskbarProgress) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWinTaskbarProgress) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWinTaskbarProgressFromPointer(cthis unsafe.Pointer) *QWinTaskbarProgress {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWinTaskbarProgress{bcthis0}
}
func (*QWinTaskbarProgress) NewFromPointer(cthis unsafe.Pointer) *QWinTaskbarProgress {
	return NewQWinTaskbarProgressFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QWinTaskbarProgress) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWinTaskbarProgress10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinTaskbarProgress(QObject *)
func NewQWinTaskbarProgress(parent *qtcore.QObject /*777 QObject **/) *QWinTaskbarProgress {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgressC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWinTaskbarProgressFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinTaskbarProgress()
func DeleteQWinTaskbarProgress(this *QWinTaskbarProgress) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgressD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] int value()
func (this *QWinTaskbarProgress) Value() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWinTaskbarProgress5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimum()
func (this *QWinTaskbarProgress) Minimum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWinTaskbarProgress7minimumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximum()
func (this *QWinTaskbarProgress) Maximum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWinTaskbarProgress7maximumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QWinTaskbarProgress) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWinTaskbarProgress9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPaused()
func (this *QWinTaskbarProgress) IsPaused() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWinTaskbarProgress8isPausedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:70
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStopped()
func (this *QWinTaskbarProgress) IsStopped() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWinTaskbarProgress9isStoppedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(int)
func (this *QWinTaskbarProgress) SetValue(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress8setValueEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimum(int)
func (this *QWinTaskbarProgress) SetMinimum(minimum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress10setMinimumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximum(int)
func (this *QWinTaskbarProgress) SetMaximum(maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress10setMaximumEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRange(int, int)
func (this *QWinTaskbarProgress) SetRange(minimum int, maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress8setRangeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum, maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reset()
func (this *QWinTaskbarProgress) Reset() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void show()
func (this *QWinTaskbarProgress) Show() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress4showEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hide()
func (this *QWinTaskbarProgress) Hide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress4hideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QWinTaskbarProgress) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pause()
func (this *QWinTaskbarProgress) Pause() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress5pauseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resume()
func (this *QWinTaskbarProgress) Resume() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress6resumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPaused(_Bool)
func (this *QWinTaskbarProgress) SetPaused(paused bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress9setPausedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()
func (this *QWinTaskbarProgress) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void valueChanged(int)
func (this *QWinTaskbarProgress) ValueChanged(value int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress12valueChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void minimumChanged(int)
func (this *QWinTaskbarProgress) MinimumChanged(minimum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress14minimumChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minimum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void maximumChanged(int)
func (this *QWinTaskbarProgress) MaximumChanged(maximum int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress14maximumChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maximum)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibilityChanged(_Bool)
func (this *QWinTaskbarProgress) VisibilityChanged(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress17visibilityChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void pausedChanged(_Bool)
func (this *QWinTaskbarProgress) PausedChanged(paused bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress13pausedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), paused)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwintaskbarprogress.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stoppedChanged(_Bool)
func (this *QWinTaskbarProgress) StoppedChanged(stopped bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWinTaskbarProgress14stoppedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stopped)
	gopp.ErrPrint(err, rv)
}

//  body block end
