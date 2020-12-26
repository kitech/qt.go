package qtgui

// /usr/include/qt/QtGui/qsurface.h
// #include <qsurface.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
// size 24
type QSurface struct {
	*qtrt.CObject
}
type QSurface_ITF interface {
	QSurface_PTR() *QSurface
}

func (ptr *QSurface) QSurface_PTR() *QSurface { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QSurfaceFromptr(cthis Voidptr) *QSurface {
	return &QSurface{&qtrt.CObject{cthis}}
}
func (*QSurface) Fromptr(cthis Voidptr) *QSurface {
	return QSurfaceFromptr(cthis)
}

func DeleteQSurface(this *QSurface) {
	rv, err := qtrt.Qtcc3(4285993900, "_ZN8QSurfaceD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QSurface__SurfaceClass = int

//
const QSurface__Window QSurface__SurfaceClass = 0

//
const QSurface__Offscreen QSurface__SurfaceClass = 1

func (this *QSurface) SurfaceClassItemName(val int) string {
	switch val {
	case QSurface__Window: // 0
		return "Window"
	case QSurface__Offscreen: // 1
		return "Offscreen"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSurface_SurfaceClassItemName(val int) string {
	var nilthis *QSurface
	return nilthis.SurfaceClassItemName(val)
}

/*


 */
type QSurface__SurfaceType = int

//
const QSurface__RasterSurface QSurface__SurfaceType = 0

//
const QSurface__OpenGLSurface QSurface__SurfaceType = 1

//
const QSurface__RasterGLSurface QSurface__SurfaceType = 2

//
const QSurface__OpenVGSurface QSurface__SurfaceType = 3

//
const QSurface__VulkanSurface QSurface__SurfaceType = 4

//
const QSurface__MetalSurface QSurface__SurfaceType = 5

func (this *QSurface) SurfaceTypeItemName(val int) string {
	switch val {
	case QSurface__RasterSurface: // 0
		return "RasterSurface"
	case QSurface__OpenGLSurface: // 1
		return "OpenGLSurface"
	case QSurface__RasterGLSurface: // 2
		return "RasterGLSurface"
	case QSurface__OpenVGSurface: // 3
		return "OpenVGSurface"
	case QSurface__VulkanSurface: // 4
		return "VulkanSurface"
	case QSurface__MetalSurface: // 5
		return "MetalSurface"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSurface_SurfaceTypeItemName(val int) string {
	var nilthis *QSurface
	return nilthis.SurfaceTypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10147() {
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
