package qtmultimedia

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h
// #include <qabstractvideofilter.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
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
type QVideoFilterRunnable struct {
	*qtrt.CObject
}
type QVideoFilterRunnable_ITF interface {
	QVideoFilterRunnable_PTR() *QVideoFilterRunnable
}

func (ptr *QVideoFilterRunnable) QVideoFilterRunnable_PTR() *QVideoFilterRunnable { return ptr }

func (this *QVideoFilterRunnable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QVideoFilterRunnable) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQVideoFilterRunnableFromPointer(cthis unsafe.Pointer) *QVideoFilterRunnable {
	return &QVideoFilterRunnable{&qtrt.CObject{cthis}}
}
func (*QVideoFilterRunnable) NewFromPointer(cthis unsafe.Pointer) *QVideoFilterRunnable {
	return NewQVideoFilterRunnableFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QVideoFilterRunnable()

/*

 */
func DeleteQVideoFilterRunnable(this *QVideoFilterRunnable) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QVideoFilterRunnableD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qabstractvideofilter.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QVideoFrame run(QVideoFrame *, const QVideoSurfaceFormat &, QVideoFilterRunnable::RunFlags)

/*

 */
func (this *QVideoFilterRunnable) Run(input QVideoFrame_ITF /*777 QVideoFrame **/, surfaceFormat QVideoSurfaceFormat_ITF, flags int) *QVideoFrame /*123*/ {
	var convArg0 unsafe.Pointer
	if input != nil && input.QVideoFrame_PTR() != nil {
		convArg0 = input.QVideoFrame_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if surfaceFormat != nil && surfaceFormat.QVideoSurfaceFormat_PTR() != nil {
		convArg1 = surfaceFormat.QVideoSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QVideoFilterRunnable3runEP11QVideoFrameRK19QVideoSurfaceFormat6QFlagsINS_7RunFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVideoFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVideoFrame)
	return rv2
}

/*


 */
type QVideoFilterRunnable__RunFlag = int

//
const QVideoFilterRunnable__LastInChain QVideoFilterRunnable__RunFlag = 1

func (this *QVideoFilterRunnable) RunFlagItemName(val int) string {
	switch val {
	case QVideoFilterRunnable__LastInChain: // 1
		return "LastInChain"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QVideoFilterRunnable_RunFlagItemName(val int) string {
	var nilthis *QVideoFilterRunnable
	return nilthis.RunFlagItemName(val)
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
