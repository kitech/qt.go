package qtgui

// /usr/include/qt/QtGui/qicon.h
// #include <qicon.h>
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
type QIcon struct {
	*qtrt.CObject
}
type QIcon_ITF interface {
	QIcon_PTR() *QIcon
}

func (ptr *QIcon) QIcon_PTR() *QIcon { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QIconFromptr(cthis Voidptr) *QIcon {
	return &QIcon{&qtrt.CObject{cthis}}
}
func (*QIcon) Fromptr(cthis Voidptr) *QIcon {
	return QIconFromptr(cthis)
}

// /usr/include/qt/QtGui/qicon.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIcon()

/*
 */
func (*QIcon) NewForInherit() *QIcon {
	return NewQIcon()
}
func NewQIcon() *QIcon {
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc3(1359309426, "_ZN5QIconC2Ev", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, Voidptr(&cthis))
	qtrt.ErrPrint2(err, rv)
	gothis := QIconFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

func DeleteQIcon(this *QIcon) {
	rv, err := qtrt.Qtcc1(0, "_ZN5QIconD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QIcon__Mode = int

//
const QIcon__Normal QIcon__Mode = 0

//
const QIcon__Disabled QIcon__Mode = 1

//
const QIcon__Active QIcon__Mode = 2

//
const QIcon__Selected QIcon__Mode = 3

func (this *QIcon) ModeItemName(val int) string {
	switch val {
	case QIcon__Normal: // 0
		return "Normal"
	case QIcon__Disabled: // 1
		return "Disabled"
	case QIcon__Active: // 2
		return "Active"
	case QIcon__Selected: // 3
		return "Selected"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QIcon_ModeItemName(val int) string {
	var nilthis *QIcon
	return nilthis.ModeItemName(val)
}

/*


 */
type QIcon__State = int

//
const QIcon__On QIcon__State = 0

//
const QIcon__Off QIcon__State = 1

func (this *QIcon) StateItemName(val int) string {
	switch val {
	case QIcon__On: // 0
		return "On"
	case QIcon__Off: // 1
		return "Off"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QIcon_StateItemName(val int) string {
	var nilthis *QIcon
	return nilthis.StateItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10149() {
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
