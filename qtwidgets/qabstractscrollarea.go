package qtwidgets

// /usr/include/qt/QtWidgets/qabstractscrollarea.h
// #include <qabstractscrollarea.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
// size 48
type QAbstractScrollArea struct {
	*QFrame
}
type QAbstractScrollArea_ITF interface {
	QFrame_ITF
	QAbstractScrollArea_PTR() *QAbstractScrollArea
}

func (ptr *QAbstractScrollArea) QAbstractScrollArea_PTR() *QAbstractScrollArea { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
func QAbstractScrollAreaFromptr(cthis unsafe.Pointer) *QAbstractScrollArea {
	bcthis0 := QFrameFromptr(cthis)
	return &QAbstractScrollArea{bcthis0}
}
func (*QAbstractScrollArea) Fromptr(cthis unsafe.Pointer) *QAbstractScrollArea {
	return QAbstractScrollAreaFromptr(cthis)
}

/*


 */
type QAbstractScrollArea__SizeAdjustPolicy = int

//
const QAbstractScrollArea__AdjustIgnored QAbstractScrollArea__SizeAdjustPolicy = 0

//
const QAbstractScrollArea__AdjustToContentsOnFirstShow QAbstractScrollArea__SizeAdjustPolicy = 1

//
const QAbstractScrollArea__AdjustToContents QAbstractScrollArea__SizeAdjustPolicy = 2

func (this *QAbstractScrollArea) SizeAdjustPolicyItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractScrollArea_SizeAdjustPolicyItemName(val int) string {
	var nilthis *QAbstractScrollArea
	return nilthis.SizeAdjustPolicyItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10091() {
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
