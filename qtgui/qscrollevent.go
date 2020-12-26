package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
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
// size 64
type QScrollEvent struct {
	*qtcore.QEvent
}
type QScrollEvent_ITF interface {
	qtcore.QEvent_ITF
	QScrollEvent_PTR() *QScrollEvent
}

func (ptr *QScrollEvent) QScrollEvent_PTR() *QScrollEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QScrollEventFromptr(cthis Voidptr) *QScrollEvent {
	bcthis0 := qtcore.QEventFromptr(cthis)
	return &QScrollEvent{bcthis0}
}
func (*QScrollEvent) Fromptr(cthis Voidptr) *QScrollEvent {
	return QScrollEventFromptr(cthis)
}

func DeleteQScrollEvent(this *QScrollEvent) {
	rv, err := qtrt.Qtcc3(2663943836, "_ZN12QScrollEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QScrollEvent__ScrollState = int

//
const QScrollEvent__ScrollStarted QScrollEvent__ScrollState = 0

//
const QScrollEvent__ScrollUpdated QScrollEvent__ScrollState = 1

//
const QScrollEvent__ScrollFinished QScrollEvent__ScrollState = 2

func (this *QScrollEvent) ScrollStateItemName(val int) string {
	switch val {
	case QScrollEvent__ScrollStarted: // 0
		return "ScrollStarted"
	case QScrollEvent__ScrollUpdated: // 1
		return "ScrollUpdated"
	case QScrollEvent__ScrollFinished: // 2
		return "ScrollFinished"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QScrollEvent_ScrollStateItemName(val int) string {
	var nilthis *QScrollEvent
	return nilthis.ScrollStateItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10129() {
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
