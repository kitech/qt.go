package qtgui

// /usr/include/qt/QtGui/qwindow.h
// #include <qwindow.h>
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
// size 40
type QWindow struct {
	*qtcore.QObject
	*QSurface
}
type QWindow_ITF interface {
	qtcore.QObject_ITF
	QSurface_ITF
	QWindow_PTR() *QWindow
}

func (ptr *QWindow) QWindow_PTR() *QWindow { return ptr }

func (this *QWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWindow) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.QObjectFromptr(cthis)
	this.QSurface = QSurfaceFromptr(cthis)
}
func QWindowFromptr(cthis unsafe.Pointer) *QWindow {
	bcthis0 := qtcore.QObjectFromptr(cthis)
	bcthis1 := QSurfaceFromptr(cthis)
	return &QWindow{bcthis0, bcthis1}
}
func (*QWindow) Fromptr(cthis unsafe.Pointer) *QWindow {
	return QWindowFromptr(cthis)
}

func DeleteQWindow(this *QWindow) {
	rv, err := qtrt.Qtcc1(0, "_ZN7QWindowD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QWindow__Visibility = int

//
const QWindow__Hidden QWindow__Visibility = 0

//
const QWindow__AutomaticVisibility QWindow__Visibility = 1

//
const QWindow__Windowed QWindow__Visibility = 2

//
const QWindow__Minimized QWindow__Visibility = 3

//
const QWindow__Maximized QWindow__Visibility = 4

//
const QWindow__FullScreen QWindow__Visibility = 5

func (this *QWindow) VisibilityItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWindow_VisibilityItemName(val int) string {
	var nilthis *QWindow
	return nilthis.VisibilityItemName(val)
}

/*


 */
type QWindow__AncestorMode = int

//
const QWindow__ExcludeTransients QWindow__AncestorMode = 0

//
const QWindow__IncludeTransients QWindow__AncestorMode = 1

func (this *QWindow) AncestorModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QWindow_AncestorModeItemName(val int) string {
	var nilthis *QWindow
	return nilthis.AncestorModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10065() {
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
