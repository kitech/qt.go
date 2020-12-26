package qtgui

// /usr/include/qt/QtGui/qpaintengine.h
// #include <qpaintengine.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
// size 32
type QPaintEngine struct {
	*qtrt.CObject
}
type QPaintEngine_ITF interface {
	QPaintEngine_PTR() *QPaintEngine
}

func (ptr *QPaintEngine) QPaintEngine_PTR() *QPaintEngine { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QPaintEngineFromptr(cthis Voidptr) *QPaintEngine {
	return &QPaintEngine{&qtrt.CObject{cthis}}
}
func (*QPaintEngine) Fromptr(cthis Voidptr) *QPaintEngine {
	return QPaintEngineFromptr(cthis)
}

func DeleteQPaintEngine(this *QPaintEngine) {
	rv, err := qtrt.Qtcc1(0, "_ZN12QPaintEngineD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QPaintEngine__PaintEngineFeature = int

//
const QPaintEngine__PrimitiveTransform QPaintEngine__PaintEngineFeature = 1

//
const QPaintEngine__PatternTransform QPaintEngine__PaintEngineFeature = 2

//
const QPaintEngine__PixmapTransform QPaintEngine__PaintEngineFeature = 4

//
const QPaintEngine__PatternBrush QPaintEngine__PaintEngineFeature = 8

//
const QPaintEngine__LinearGradientFill QPaintEngine__PaintEngineFeature = 16

//
const QPaintEngine__RadialGradientFill QPaintEngine__PaintEngineFeature = 32

//
const QPaintEngine__ConicalGradientFill QPaintEngine__PaintEngineFeature = 64

//
const QPaintEngine__AlphaBlend QPaintEngine__PaintEngineFeature = 128

//
const QPaintEngine__PorterDuff QPaintEngine__PaintEngineFeature = 256

//
const QPaintEngine__PainterPaths QPaintEngine__PaintEngineFeature = 512

//
const QPaintEngine__Antialiasing QPaintEngine__PaintEngineFeature = 1024

//
const QPaintEngine__BrushStroke QPaintEngine__PaintEngineFeature = 2048

//
const QPaintEngine__ConstantOpacity QPaintEngine__PaintEngineFeature = 4096

//
const QPaintEngine__MaskedBrush QPaintEngine__PaintEngineFeature = 8192

//
const QPaintEngine__PerspectiveTransform QPaintEngine__PaintEngineFeature = 16384

//
const QPaintEngine__BlendModes QPaintEngine__PaintEngineFeature = 32768

//
const QPaintEngine__ObjectBoundingModeGradients QPaintEngine__PaintEngineFeature = 65536

//
const QPaintEngine__RasterOpModes QPaintEngine__PaintEngineFeature = 131072

//
const QPaintEngine__PaintOutsidePaintEvent QPaintEngine__PaintEngineFeature = 536870912

//
const QPaintEngine__AllFeatures QPaintEngine__PaintEngineFeature = -1

func (this *QPaintEngine) PaintEngineFeatureItemName(val int) string {
	switch val {
	case QPaintEngine__PrimitiveTransform: // 1
		return "PrimitiveTransform"
	case QPaintEngine__PatternTransform: // 2
		return "PatternTransform"
	case QPaintEngine__PixmapTransform: // 4
		return "PixmapTransform"
	case QPaintEngine__PatternBrush: // 8
		return "PatternBrush"
	case QPaintEngine__LinearGradientFill: // 16
		return "LinearGradientFill"
	case QPaintEngine__RadialGradientFill: // 32
		return "RadialGradientFill"
	case QPaintEngine__ConicalGradientFill: // 64
		return "ConicalGradientFill"
	case QPaintEngine__AlphaBlend: // 128
		return "AlphaBlend"
	case QPaintEngine__PorterDuff: // 256
		return "PorterDuff"
	case QPaintEngine__PainterPaths: // 512
		return "PainterPaths"
	case QPaintEngine__Antialiasing: // 1024
		return "Antialiasing"
	case QPaintEngine__BrushStroke: // 2048
		return "BrushStroke"
	case QPaintEngine__ConstantOpacity: // 4096
		return "ConstantOpacity"
	case QPaintEngine__MaskedBrush: // 8192
		return "MaskedBrush"
	case QPaintEngine__PerspectiveTransform: // 16384
		return "PerspectiveTransform"
	case QPaintEngine__BlendModes: // 32768
		return "BlendModes"
	case QPaintEngine__ObjectBoundingModeGradients: // 65536
		return "ObjectBoundingModeGradients"
	case QPaintEngine__RasterOpModes: // 131072
		return "RasterOpModes"
	case QPaintEngine__PaintOutsidePaintEvent: // 536870912
		return "PaintOutsidePaintEvent"
	case QPaintEngine__AllFeatures: // -1
		return "AllFeatures"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPaintEngine_PaintEngineFeatureItemName(val int) string {
	var nilthis *QPaintEngine
	return nilthis.PaintEngineFeatureItemName(val)
}

/*


 */
type QPaintEngine__DirtyFlag = int

//
const QPaintEngine__DirtyPen QPaintEngine__DirtyFlag = 1

//
const QPaintEngine__DirtyBrush QPaintEngine__DirtyFlag = 2

//
const QPaintEngine__DirtyBrushOrigin QPaintEngine__DirtyFlag = 4

//
const QPaintEngine__DirtyFont QPaintEngine__DirtyFlag = 8

//
const QPaintEngine__DirtyBackground QPaintEngine__DirtyFlag = 16

//
const QPaintEngine__DirtyBackgroundMode QPaintEngine__DirtyFlag = 32

//
const QPaintEngine__DirtyTransform QPaintEngine__DirtyFlag = 64

//
const QPaintEngine__DirtyClipRegion QPaintEngine__DirtyFlag = 128

//
const QPaintEngine__DirtyClipPath QPaintEngine__DirtyFlag = 256

//
const QPaintEngine__DirtyHints QPaintEngine__DirtyFlag = 512

//
const QPaintEngine__DirtyCompositionMode QPaintEngine__DirtyFlag = 1024

//
const QPaintEngine__DirtyClipEnabled QPaintEngine__DirtyFlag = 2048

//
const QPaintEngine__DirtyOpacity QPaintEngine__DirtyFlag = 4096

//
const QPaintEngine__AllDirty QPaintEngine__DirtyFlag = 65535

func (this *QPaintEngine) DirtyFlagItemName(val int) string {
	switch val {
	case QPaintEngine__DirtyPen: // 1
		return "DirtyPen"
	case QPaintEngine__DirtyBrush: // 2
		return "DirtyBrush"
	case QPaintEngine__DirtyBrushOrigin: // 4
		return "DirtyBrushOrigin"
	case QPaintEngine__DirtyFont: // 8
		return "DirtyFont"
	case QPaintEngine__DirtyBackground: // 16
		return "DirtyBackground"
	case QPaintEngine__DirtyBackgroundMode: // 32
		return "DirtyBackgroundMode"
	case QPaintEngine__DirtyTransform: // 64
		return "DirtyTransform"
	case QPaintEngine__DirtyClipRegion: // 128
		return "DirtyClipRegion"
	case QPaintEngine__DirtyClipPath: // 256
		return "DirtyClipPath"
	case QPaintEngine__DirtyHints: // 512
		return "DirtyHints"
	case QPaintEngine__DirtyCompositionMode: // 1024
		return "DirtyCompositionMode"
	case QPaintEngine__DirtyClipEnabled: // 2048
		return "DirtyClipEnabled"
	case QPaintEngine__DirtyOpacity: // 4096
		return "DirtyOpacity"
	case QPaintEngine__AllDirty: // 65535
		return "AllDirty"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPaintEngine_DirtyFlagItemName(val int) string {
	var nilthis *QPaintEngine
	return nilthis.DirtyFlagItemName(val)
}

/*


 */
type QPaintEngine__PolygonDrawMode = int

//
const QPaintEngine__OddEvenMode QPaintEngine__PolygonDrawMode = 0

//
const QPaintEngine__WindingMode QPaintEngine__PolygonDrawMode = 1

//
const QPaintEngine__ConvexMode QPaintEngine__PolygonDrawMode = 2

//
const QPaintEngine__PolylineMode QPaintEngine__PolygonDrawMode = 3

func (this *QPaintEngine) PolygonDrawModeItemName(val int) string {
	switch val {
	case QPaintEngine__OddEvenMode: // 0
		return "OddEvenMode"
	case QPaintEngine__WindingMode: // 1
		return "WindingMode"
	case QPaintEngine__ConvexMode: // 2
		return "ConvexMode"
	case QPaintEngine__PolylineMode: // 3
		return "PolylineMode"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPaintEngine_PolygonDrawModeItemName(val int) string {
	var nilthis *QPaintEngine
	return nilthis.PolygonDrawModeItemName(val)
}

/*


 */
type QPaintEngine__Type = int

//
const QPaintEngine__X11 QPaintEngine__Type = 0

//
const QPaintEngine__Windows QPaintEngine__Type = 1

//
const QPaintEngine__QuickDraw QPaintEngine__Type = 2

//
const QPaintEngine__CoreGraphics QPaintEngine__Type = 3

//
const QPaintEngine__MacPrinter QPaintEngine__Type = 4

//
const QPaintEngine__QWindowSystem QPaintEngine__Type = 5

//
const QPaintEngine__PostScript QPaintEngine__Type = 6

//
const QPaintEngine__OpenGL QPaintEngine__Type = 7

//
const QPaintEngine__Picture QPaintEngine__Type = 8

//
const QPaintEngine__SVG QPaintEngine__Type = 9

//
const QPaintEngine__Raster QPaintEngine__Type = 10

//
const QPaintEngine__Direct3D QPaintEngine__Type = 11

//
const QPaintEngine__Pdf QPaintEngine__Type = 12

//
const QPaintEngine__OpenVG QPaintEngine__Type = 13

//
const QPaintEngine__OpenGL2 QPaintEngine__Type = 14

//
const QPaintEngine__PaintBuffer QPaintEngine__Type = 15

//
const QPaintEngine__Blitter QPaintEngine__Type = 16

//
const QPaintEngine__Direct2D QPaintEngine__Type = 17

//
const QPaintEngine__User QPaintEngine__Type = 50

//
const QPaintEngine__MaxUser QPaintEngine__Type = 100

func (this *QPaintEngine) TypeItemName(val int) string {
	switch val {
	case QPaintEngine__X11: // 0
		return "X11"
	case QPaintEngine__Windows: // 1
		return "Windows"
	case QPaintEngine__QuickDraw: // 2
		return "QuickDraw"
	case QPaintEngine__CoreGraphics: // 3
		return "CoreGraphics"
	case QPaintEngine__MacPrinter: // 4
		return "MacPrinter"
	case QPaintEngine__QWindowSystem: // 5
		return "QWindowSystem"
	case QPaintEngine__PostScript: // 6
		return "PostScript"
	case QPaintEngine__OpenGL: // 7
		return "OpenGL"
	case QPaintEngine__Picture: // 8
		return "Picture"
	case QPaintEngine__SVG: // 9
		return "SVG"
	case QPaintEngine__Raster: // 10
		return "Raster"
	case QPaintEngine__Direct3D: // 11
		return "Direct3D"
	case QPaintEngine__Pdf: // 12
		return "Pdf"
	case QPaintEngine__OpenVG: // 13
		return "OpenVG"
	case QPaintEngine__OpenGL2: // 14
		return "OpenGL2"
	case QPaintEngine__PaintBuffer: // 15
		return "PaintBuffer"
	case QPaintEngine__Blitter: // 16
		return "Blitter"
	case QPaintEngine__Direct2D: // 17
		return "Direct2D"
	case QPaintEngine__User: // 50
		return "User"
	case QPaintEngine__MaxUser: // 100
		return "MaxUser"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPaintEngine_TypeItemName(val int) string {
	var nilthis *QPaintEngine
	return nilthis.TypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10165() {
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
