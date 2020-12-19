package qtgui

// /usr/include/qt/QtGui/qpixmap.h
// #include <qpixmap.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QPixmap struct {
	*QPaintDevice
}
type QPixmap_ITF interface {
	QPaintDevice_ITF
	QPixmap_PTR() *QPixmap
}

func (ptr *QPixmap) QPixmap_PTR() *QPixmap { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QPixmapFromptr(cthis unsafe.Pointer) *QPixmap {
	bcthis0 := QPaintDeviceFromptr(cthis)
	return &QPixmap{bcthis0}
}
func (*QPixmap) Fromptr(cthis unsafe.Pointer) *QPixmap {
	return QPixmapFromptr(cthis)
}

// /usr/include/qt/QtGui/qpixmap.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPixmap()

/*
 */
func (*QPixmap) NewForInherit() *QPixmap {
	return NewQPixmap()
}
func NewQPixmap() *QPixmap {
	cthis := qtrt.Malloc(32)
	rv, err := qtrt.Qtcc1(440989401, "_ZN7QPixmapC2Ev", qtrt.FFITY_POINTER, cthis)
	qtrt.ErrPrint(err, rv)
	gothis := QPixmapFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQPixmap)
	return gothis
}

func DeleteQPixmap(this *QPixmap) {
	rv, err := qtrt.Qtcc1(0, "_ZN7QPixmapD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10055() {
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
