package qtwidgets

// /usr/include/qt/QtWidgets/qsizepolicy.h
// #include <qsizepolicy.h>
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
// size 4
type QSizePolicy struct {
	*qtrt.CObject
}
type QSizePolicy_ITF interface {
	QSizePolicy_PTR() *QSizePolicy
}

func (ptr *QSizePolicy) QSizePolicy_PTR() *QSizePolicy { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QSizePolicyFromptr(cthis unsafe.Pointer) *QSizePolicy {
	return &QSizePolicy{&qtrt.CObject{cthis}}
}
func (*QSizePolicy) Fromptr(cthis unsafe.Pointer) *QSizePolicy {
	return QSizePolicyFromptr(cthis)
}

func DeleteQSizePolicy(this *QSizePolicy) {
	rv, err := qtrt.Qtcc1(0, "_ZN11QSizePolicyD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QSizePolicy__PolicyFlag = int

//
const QSizePolicy__GrowFlag QSizePolicy__PolicyFlag = 1

//
const QSizePolicy__ExpandFlag QSizePolicy__PolicyFlag = 2

//
const QSizePolicy__ShrinkFlag QSizePolicy__PolicyFlag = 4

//
const QSizePolicy__IgnoreFlag QSizePolicy__PolicyFlag = 8

func (this *QSizePolicy) PolicyFlagItemName(val int) string {
	switch val {
	case QSizePolicy__GrowFlag: // 1
		return "GrowFlag"
	case QSizePolicy__ExpandFlag: // 2
		return "ExpandFlag"
	case QSizePolicy__ShrinkFlag: // 4
		return "ShrinkFlag"
	case QSizePolicy__IgnoreFlag: // 8
		return "IgnoreFlag"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSizePolicy_PolicyFlagItemName(val int) string {
	var nilthis *QSizePolicy
	return nilthis.PolicyFlagItemName(val)
}

/*


 */
type QSizePolicy__Policy = int

//
const QSizePolicy__Fixed QSizePolicy__Policy = 0

//
const QSizePolicy__Minimum QSizePolicy__Policy = 1

//
const QSizePolicy__Maximum QSizePolicy__Policy = 4

//
const QSizePolicy__Preferred QSizePolicy__Policy = 5

//
const QSizePolicy__MinimumExpanding QSizePolicy__Policy = 3

//
const QSizePolicy__Expanding QSizePolicy__Policy = 7

//
const QSizePolicy__Ignored QSizePolicy__Policy = 13

func (this *QSizePolicy) PolicyItemName(val int) string {
	switch val {
	case QSizePolicy__Fixed: // 0
		return "Fixed"
	case QSizePolicy__Minimum: // 1
		return "Minimum"
	case QSizePolicy__Maximum: // 4
		return "Maximum"
	case QSizePolicy__Preferred: // 5
		return "Preferred"
	case QSizePolicy__MinimumExpanding: // 3
		return "MinimumExpanding"
	case QSizePolicy__Expanding: // 7
		return "Expanding"
	case QSizePolicy__Ignored: // 13
		return "Ignored"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSizePolicy_PolicyItemName(val int) string {
	var nilthis *QSizePolicy
	return nilthis.PolicyItemName(val)
}

/*


 */
type QSizePolicy__ControlType = int

//
const QSizePolicy__DefaultType QSizePolicy__ControlType = 1

//
const QSizePolicy__ButtonBox QSizePolicy__ControlType = 2

//
const QSizePolicy__CheckBox QSizePolicy__ControlType = 4

//
const QSizePolicy__ComboBox QSizePolicy__ControlType = 8

//
const QSizePolicy__Frame QSizePolicy__ControlType = 16

//
const QSizePolicy__GroupBox QSizePolicy__ControlType = 32

//
const QSizePolicy__Label QSizePolicy__ControlType = 64

//
const QSizePolicy__Line QSizePolicy__ControlType = 128

//
const QSizePolicy__LineEdit QSizePolicy__ControlType = 256

//
const QSizePolicy__PushButton QSizePolicy__ControlType = 512

//
const QSizePolicy__RadioButton QSizePolicy__ControlType = 1024

//
const QSizePolicy__Slider QSizePolicy__ControlType = 2048

//
const QSizePolicy__SpinBox QSizePolicy__ControlType = 4096

//
const QSizePolicy__TabWidget QSizePolicy__ControlType = 8192

//
const QSizePolicy__ToolButton QSizePolicy__ControlType = 16384

func (this *QSizePolicy) ControlTypeItemName(val int) string {
	switch val {
	case QSizePolicy__DefaultType: // 1
		return "DefaultType"
	case QSizePolicy__ButtonBox: // 2
		return "ButtonBox"
	case QSizePolicy__CheckBox: // 4
		return "CheckBox"
	case QSizePolicy__ComboBox: // 8
		return "ComboBox"
	case QSizePolicy__Frame: // 16
		return "Frame"
	case QSizePolicy__GroupBox: // 32
		return "GroupBox"
	case QSizePolicy__Label: // 64
		return "Label"
	case QSizePolicy__Line: // 128
		return "Line"
	case QSizePolicy__LineEdit: // 256
		return "LineEdit"
	case QSizePolicy__PushButton: // 512
		return "PushButton"
	case QSizePolicy__RadioButton: // 1024
		return "RadioButton"
	case QSizePolicy__Slider: // 2048
		return "Slider"
	case QSizePolicy__SpinBox: // 4096
		return "SpinBox"
	case QSizePolicy__TabWidget: // 8192
		return "TabWidget"
	case QSizePolicy__ToolButton: // 16384
		return "ToolButton"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSizePolicy_ControlTypeItemName(val int) string {
	var nilthis *QSizePolicy
	return nilthis.ControlTypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10081() {
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
