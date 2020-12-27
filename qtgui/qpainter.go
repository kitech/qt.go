package qtgui

// /usr/include/qt/QtGui/qpainter.h
// #include <qpainter.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
// size 8
type QPainter struct {
	*qtrt.CObject
}
type QPainter_ITF interface {
	QPainter_PTR() *QPainter
}

func (ptr *QPainter) QPainter_PTR() *QPainter { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QPainterFromptr(cthis Voidptr) *QPainter {
	return &QPainter{&qtrt.CObject{cthis}}
}
func (*QPainter) Fromptr(cthis Voidptr) *QPainter {
	return QPainterFromptr(cthis)
}

// /usr/include/qt/QtGui/qpainter.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPainter(QPaintDevice *)

/*
 */
func (*QPainter) NewForInherit(arg0 QPaintDevice_ITF /*777 QPaintDevice **/) *QPainter {
	return NewQPainter(arg0)
}
func NewQPainter(arg0 QPaintDevice_ITF /*777 QPaintDevice **/) *QPainter {
	var convArg0 Voidptr
	if arg0 != nil && arg0.QPaintDevice_PTR() != nil {
		convArg0 = arg0.QPaintDevice_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc3(1768399324, "_ZN8QPainterC2EP12QPaintDevice", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QPainterFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQPainter)
	return gothis
}

// /usr/include/qt/QtGui/qpainter.h:130
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QPaintDevice * device() const

/*
 */
func (this *QPainter) Device() *QPaintDevice /*777 QPaintDevice **/ {
	rv, err := qtrt.Qtcc3(3490243267, "_ZNK8QPainter6deviceEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return /*==*/ QPaintDeviceFromptr(rv.Ptr()) // 444
}

func DeleteQPainter(this *QPainter) {
	rv, err := qtrt.Qtcc3(2229242048, "_ZN8QPainterD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QPainter__RenderHint = int

//
const QPainter__Antialiasing QPainter__RenderHint = 1

//
const QPainter__TextAntialiasing QPainter__RenderHint = 2

//
const QPainter__SmoothPixmapTransform QPainter__RenderHint = 4

//
const QPainter__HighQualityAntialiasing QPainter__RenderHint = 8

//
const QPainter__NonCosmeticDefaultPen QPainter__RenderHint = 16

//
const QPainter__Qt4CompatiblePainting QPainter__RenderHint = 32

//
const QPainter__LosslessImageRendering QPainter__RenderHint = 64

func (this *QPainter) RenderHintItemName(val int) string {
	switch val {
	case QPainter__Antialiasing: // 1
		return "Antialiasing"
	case QPainter__TextAntialiasing: // 2
		return "TextAntialiasing"
	case QPainter__SmoothPixmapTransform: // 4
		return "SmoothPixmapTransform"
	case QPainter__HighQualityAntialiasing: // 8
		return "HighQualityAntialiasing"
	case QPainter__NonCosmeticDefaultPen: // 16
		return "NonCosmeticDefaultPen"
	case QPainter__Qt4CompatiblePainting: // 32
		return "Qt4CompatiblePainting"
	case QPainter__LosslessImageRendering: // 64
		return "LosslessImageRendering"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPainter_RenderHintItemName(val int) string {
	var nilthis *QPainter
	return nilthis.RenderHintItemName(val)
}

/*


 */
type QPainter__PixmapFragmentHint = int

//
const QPainter__OpaqueHint QPainter__PixmapFragmentHint = 1

func (this *QPainter) PixmapFragmentHintItemName(val int) string {
	switch val {
	case QPainter__OpaqueHint: // 1
		return "OpaqueHint"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPainter_PixmapFragmentHintItemName(val int) string {
	var nilthis *QPainter
	return nilthis.PixmapFragmentHintItemName(val)
}

/*


 */
type QPainter__CompositionMode = int

//
const QPainter__CompositionMode_SourceOver QPainter__CompositionMode = 0

//
const QPainter__CompositionMode_DestinationOver QPainter__CompositionMode = 1

//
const QPainter__CompositionMode_Clear QPainter__CompositionMode = 2

//
const QPainter__CompositionMode_Source QPainter__CompositionMode = 3

//
const QPainter__CompositionMode_Destination QPainter__CompositionMode = 4

//
const QPainter__CompositionMode_SourceIn QPainter__CompositionMode = 5

//
const QPainter__CompositionMode_DestinationIn QPainter__CompositionMode = 6

//
const QPainter__CompositionMode_SourceOut QPainter__CompositionMode = 7

//
const QPainter__CompositionMode_DestinationOut QPainter__CompositionMode = 8

//
const QPainter__CompositionMode_SourceAtop QPainter__CompositionMode = 9

//
const QPainter__CompositionMode_DestinationAtop QPainter__CompositionMode = 10

//
const QPainter__CompositionMode_Xor QPainter__CompositionMode = 11

//
const QPainter__CompositionMode_Plus QPainter__CompositionMode = 12

//
const QPainter__CompositionMode_Multiply QPainter__CompositionMode = 13

//
const QPainter__CompositionMode_Screen QPainter__CompositionMode = 14

//
const QPainter__CompositionMode_Overlay QPainter__CompositionMode = 15

//
const QPainter__CompositionMode_Darken QPainter__CompositionMode = 16

//
const QPainter__CompositionMode_Lighten QPainter__CompositionMode = 17

//
const QPainter__CompositionMode_ColorDodge QPainter__CompositionMode = 18

//
const QPainter__CompositionMode_ColorBurn QPainter__CompositionMode = 19

//
const QPainter__CompositionMode_HardLight QPainter__CompositionMode = 20

//
const QPainter__CompositionMode_SoftLight QPainter__CompositionMode = 21

//
const QPainter__CompositionMode_Difference QPainter__CompositionMode = 22

//
const QPainter__CompositionMode_Exclusion QPainter__CompositionMode = 23

//
const QPainter__RasterOp_SourceOrDestination QPainter__CompositionMode = 24

//
const QPainter__RasterOp_SourceAndDestination QPainter__CompositionMode = 25

//
const QPainter__RasterOp_SourceXorDestination QPainter__CompositionMode = 26

//
const QPainter__RasterOp_NotSourceAndNotDestination QPainter__CompositionMode = 27

//
const QPainter__RasterOp_NotSourceOrNotDestination QPainter__CompositionMode = 28

//
const QPainter__RasterOp_NotSourceXorDestination QPainter__CompositionMode = 29

//
const QPainter__RasterOp_NotSource QPainter__CompositionMode = 30

//
const QPainter__RasterOp_NotSourceAndDestination QPainter__CompositionMode = 31

//
const QPainter__RasterOp_SourceAndNotDestination QPainter__CompositionMode = 32

//
const QPainter__RasterOp_NotSourceOrDestination QPainter__CompositionMode = 33

//
const QPainter__RasterOp_SourceOrNotDestination QPainter__CompositionMode = 34

//
const QPainter__RasterOp_ClearDestination QPainter__CompositionMode = 35

//
const QPainter__RasterOp_SetDestination QPainter__CompositionMode = 36

//
const QPainter__RasterOp_NotDestination QPainter__CompositionMode = 37

func (this *QPainter) CompositionModeItemName(val int) string {
	switch val {
	case QPainter__CompositionMode_SourceOver: // 0
		return "CompositionMode_SourceOver"
	case QPainter__CompositionMode_DestinationOver: // 1
		return "CompositionMode_DestinationOver"
	case QPainter__CompositionMode_Clear: // 2
		return "CompositionMode_Clear"
	case QPainter__CompositionMode_Source: // 3
		return "CompositionMode_Source"
	case QPainter__CompositionMode_Destination: // 4
		return "CompositionMode_Destination"
	case QPainter__CompositionMode_SourceIn: // 5
		return "CompositionMode_SourceIn"
	case QPainter__CompositionMode_DestinationIn: // 6
		return "CompositionMode_DestinationIn"
	case QPainter__CompositionMode_SourceOut: // 7
		return "CompositionMode_SourceOut"
	case QPainter__CompositionMode_DestinationOut: // 8
		return "CompositionMode_DestinationOut"
	case QPainter__CompositionMode_SourceAtop: // 9
		return "CompositionMode_SourceAtop"
	case QPainter__CompositionMode_DestinationAtop: // 10
		return "CompositionMode_DestinationAtop"
	case QPainter__CompositionMode_Xor: // 11
		return "CompositionMode_Xor"
	case QPainter__CompositionMode_Plus: // 12
		return "CompositionMode_Plus"
	case QPainter__CompositionMode_Multiply: // 13
		return "CompositionMode_Multiply"
	case QPainter__CompositionMode_Screen: // 14
		return "CompositionMode_Screen"
	case QPainter__CompositionMode_Overlay: // 15
		return "CompositionMode_Overlay"
	case QPainter__CompositionMode_Darken: // 16
		return "CompositionMode_Darken"
	case QPainter__CompositionMode_Lighten: // 17
		return "CompositionMode_Lighten"
	case QPainter__CompositionMode_ColorDodge: // 18
		return "CompositionMode_ColorDodge"
	case QPainter__CompositionMode_ColorBurn: // 19
		return "CompositionMode_ColorBurn"
	case QPainter__CompositionMode_HardLight: // 20
		return "CompositionMode_HardLight"
	case QPainter__CompositionMode_SoftLight: // 21
		return "CompositionMode_SoftLight"
	case QPainter__CompositionMode_Difference: // 22
		return "CompositionMode_Difference"
	case QPainter__CompositionMode_Exclusion: // 23
		return "CompositionMode_Exclusion"
	case QPainter__RasterOp_SourceOrDestination: // 24
		return "RasterOp_SourceOrDestination"
	case QPainter__RasterOp_SourceAndDestination: // 25
		return "RasterOp_SourceAndDestination"
	case QPainter__RasterOp_SourceXorDestination: // 26
		return "RasterOp_SourceXorDestination"
	case QPainter__RasterOp_NotSourceAndNotDestination: // 27
		return "RasterOp_NotSourceAndNotDestination"
	case QPainter__RasterOp_NotSourceOrNotDestination: // 28
		return "RasterOp_NotSourceOrNotDestination"
	case QPainter__RasterOp_NotSourceXorDestination: // 29
		return "RasterOp_NotSourceXorDestination"
	case QPainter__RasterOp_NotSource: // 30
		return "RasterOp_NotSource"
	case QPainter__RasterOp_NotSourceAndDestination: // 31
		return "RasterOp_NotSourceAndDestination"
	case QPainter__RasterOp_SourceAndNotDestination: // 32
		return "RasterOp_SourceAndNotDestination"
	case QPainter__RasterOp_NotSourceOrDestination: // 33
		return "RasterOp_NotSourceOrDestination"
	case QPainter__RasterOp_SourceOrNotDestination: // 34
		return "RasterOp_SourceOrNotDestination"
	case QPainter__RasterOp_ClearDestination: // 35
		return "RasterOp_ClearDestination"
	case QPainter__RasterOp_SetDestination: // 36
		return "RasterOp_SetDestination"
	case QPainter__RasterOp_NotDestination: // 37
		return "RasterOp_NotDestination"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPainter_CompositionModeItemName(val int) string {
	var nilthis *QPainter
	return nilthis.CompositionModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10171() {
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
