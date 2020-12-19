// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicseffect.h
// #include <qgraphicseffect.h>
// #include <QtWidgets>

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
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*
 */
// size 16
type QGraphicsEffect struct {
	*qtcore.QObject
}
type QGraphicsEffect_ITF interface {
	qtcore.QObject_ITF
	QGraphicsEffect_PTR() *QGraphicsEffect
}

func (ptr *QGraphicsEffect) QGraphicsEffect_PTR() *QGraphicsEffect { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QGraphicsEffectFromptr(cthis unsafe.Pointer) *QGraphicsEffect {
	bcthis0 := qtcore.QObjectFromptr(cthis)
	return &QGraphicsEffect{bcthis0}
}
func (*QGraphicsEffect) Fromptr(cthis unsafe.Pointer) *QGraphicsEffect {
	return QGraphicsEffectFromptr(cthis)
}

func DeleteQGraphicsEffect(this *QGraphicsEffect) {
	rv, err := qtrt.Qtcc1(0, "_ZN15QGraphicsEffectD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QGraphicsEffect__ChangeFlag = int

//
const QGraphicsEffect__SourceAttached QGraphicsEffect__ChangeFlag = 1

//
const QGraphicsEffect__SourceDetached QGraphicsEffect__ChangeFlag = 2

//
const QGraphicsEffect__SourceBoundingRectChanged QGraphicsEffect__ChangeFlag = 4

//
const QGraphicsEffect__SourceInvalidated QGraphicsEffect__ChangeFlag = 8

func (this *QGraphicsEffect) ChangeFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsEffect_ChangeFlagItemName(val int) string {
	var nilthis *QGraphicsEffect
	return nilthis.ChangeFlagItemName(val)
}

/*


 */
type QGraphicsEffect__PixmapPadMode = int

//
const QGraphicsEffect__NoPad QGraphicsEffect__PixmapPadMode = 0

//
const QGraphicsEffect__PadToTransparentBorder QGraphicsEffect__PixmapPadMode = 1

//
const QGraphicsEffect__PadToEffectiveBoundingRect QGraphicsEffect__PixmapPadMode = 2

func (this *QGraphicsEffect) PixmapPadModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGraphicsEffect_PixmapPadModeItemName(val int) string {
	var nilthis *QGraphicsEffect
	return nilthis.PixmapPadModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10125() {
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
}

//  keep block end
