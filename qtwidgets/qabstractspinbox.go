// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractspinbox.h
// #include <qabstractspinbox.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QAbstractSpinBox struct {
	*QWidget
}
type QAbstractSpinBox_ITF interface {
	QWidget_ITF
	QAbstractSpinBox_PTR() *QAbstractSpinBox
}

func (ptr *QAbstractSpinBox) QAbstractSpinBox_PTR() *QAbstractSpinBox { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QAbstractSpinBoxFromptr(cthis Voidptr) *QAbstractSpinBox {
	bcthis0 := QWidgetFromptr(cthis)
	return &QAbstractSpinBox{bcthis0}
}
func (*QAbstractSpinBox) Fromptr(cthis Voidptr) *QAbstractSpinBox {
	return QAbstractSpinBoxFromptr(cthis)
}

func DeleteQAbstractSpinBox(this *QAbstractSpinBox) {
	rv, err := qtrt.Qtcc1(0, "_ZN16QAbstractSpinBoxD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QAbstractSpinBox__StepEnabledFlag = int

//
const QAbstractSpinBox__StepNone QAbstractSpinBox__StepEnabledFlag = 0

//
const QAbstractSpinBox__StepUpEnabled QAbstractSpinBox__StepEnabledFlag = 1

//
const QAbstractSpinBox__StepDownEnabled QAbstractSpinBox__StepEnabledFlag = 2

func (this *QAbstractSpinBox) StepEnabledFlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSpinBox_StepEnabledFlagItemName(val int) string {
	var nilthis *QAbstractSpinBox
	return nilthis.StepEnabledFlagItemName(val)
}

/*


 */
type QAbstractSpinBox__ButtonSymbols = int

//
const QAbstractSpinBox__UpDownArrows QAbstractSpinBox__ButtonSymbols = 0

//
const QAbstractSpinBox__PlusMinus QAbstractSpinBox__ButtonSymbols = 1

//
const QAbstractSpinBox__NoButtons QAbstractSpinBox__ButtonSymbols = 2

func (this *QAbstractSpinBox) ButtonSymbolsItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSpinBox_ButtonSymbolsItemName(val int) string {
	var nilthis *QAbstractSpinBox
	return nilthis.ButtonSymbolsItemName(val)
}

/*


 */
type QAbstractSpinBox__CorrectionMode = int

//
const QAbstractSpinBox__CorrectToPreviousValue QAbstractSpinBox__CorrectionMode = 0

//
const QAbstractSpinBox__CorrectToNearestValue QAbstractSpinBox__CorrectionMode = 1

func (this *QAbstractSpinBox) CorrectionModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSpinBox_CorrectionModeItemName(val int) string {
	var nilthis *QAbstractSpinBox
	return nilthis.CorrectionModeItemName(val)
}

/*


 */
type QAbstractSpinBox__StepType = int

//
const QAbstractSpinBox__DefaultStepType QAbstractSpinBox__StepType = 0

//
const QAbstractSpinBox__AdaptiveDecimalStepType QAbstractSpinBox__StepType = 1

func (this *QAbstractSpinBox) StepTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractSpinBox_StepTypeItemName(val int) string {
	var nilthis *QAbstractSpinBox
	return nilthis.StepTypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10087() {
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
