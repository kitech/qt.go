//  header block begin
// /usr/include/qt/QtWidgets/qcolormap.h
// #include <qcolormap.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QColormap struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qcolormap.h:60
// index:0
// static
// void initialize()
func (this *QColormap) Initialize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QColormap10initializeEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColormap_Initialize() {
	// 0: (), ()
	var nilthis *QColormap
	nilthis.Initialize()
}

// /usr/include/qt/QtWidgets/qcolormap.h:61
// index:0
// static
// void cleanup()
func (this *QColormap) Cleanup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QColormap7cleanupEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColormap_Cleanup() {
	// 0: (), ()
	var nilthis *QColormap
	nilthis.Cleanup()
}

// /usr/include/qt/QtWidgets/qcolormap.h:63
// index:0
// static
// QColormap instance(int)
func (this *QColormap) Instance(screen int) {
	// 0: (int screen), (screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QColormap8instanceEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColormap_Instance(screen int) {
	// 0: (int screen), (screen)
	var nilthis *QColormap
	nilthis.Instance(screen)
}

// /usr/include/qt/QtWidgets/qcolormap.h:66
// index:0
// void ~QColormap()
func DeleteQColormap(*QColormap) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QColormapD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolormap.h:70
// index:0
// QColormap::Mode mode()
func (this *QColormap) Mode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap4modeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolormap.h:72
// index:0
// int depth()
func (this *QColormap) Depth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap5depthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolormap.h:73
// index:0
// int size()
func (this *QColormap) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolormap.h:75
// index:0
// uint pixel(const class QColor &)
func (this *QColormap) Pixel(color unsafe.Pointer) {
	// 0: (, const QColor & color), (color)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap5pixelERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolormap.h:76
// index:0
// const QColor colorAt(uint)
func (this *QColormap) ColorAt(pixel uint) {
	// 0: (, uint pixel), (&pixel)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap7colorAtEj", ffiqt.FFI_TYPE_VOID, this.cthis, &pixel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcolormap.h:78
// index:0
// const QVector<QColor> colormap()
func (this *QColormap) Colormap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QColormap8colormapEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
