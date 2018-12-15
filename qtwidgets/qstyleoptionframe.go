package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QStyleOptionFrame struct {
	*QStyleOption
}
type QStyleOptionFrame_ITF interface {
	QStyleOption_ITF
	QStyleOptionFrame_PTR() *QStyleOptionFrame
}

func (ptr *QStyleOptionFrame) QStyleOptionFrame_PTR() *QStyleOptionFrame { return ptr }

func (this *QStyleOptionFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionFrame) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionFrameFromPointer(cthis unsafe.Pointer) *QStyleOptionFrame {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionFrame{bcthis0}
}
func (*QStyleOptionFrame) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionFrame {
	return NewQStyleOptionFrameFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionFrame()

/*

 */
func (*QStyleOptionFrame) NewForInherit() *QStyleOptionFrame {
	return NewQStyleOptionFrame()
}
func NewQStyleOptionFrame() *QStyleOptionFrame {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QStyleOptionFrameC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionFrame)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:136
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionFrame(int)

/*

 */
func (*QStyleOptionFrame) NewForInherit1(version int) *QStyleOptionFrame {
	return NewQStyleOptionFrame1(version)
}
func NewQStyleOptionFrame1(version int) *QStyleOptionFrame {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QStyleOptionFrameC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionFrame)
	return gothis
}

func DeleteQStyleOptionFrame(this *QStyleOptionFrame) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QStyleOptionFrameD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionFrame__StyleOptionType = int

//
const QStyleOptionFrame__Type QStyleOptionFrame__StyleOptionType = 5

func (this *QStyleOptionFrame) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionFrame__Type: // 5
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionFrame_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionFrame
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionFrame__StyleOptionVersion = int

// 1
const QStyleOptionFrame__Version QStyleOptionFrame__StyleOptionVersion = 3

func (this *QStyleOptionFrame) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionFrame__Version: // 3
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionFrame_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionFrame
	return nilthis.StyleOptionVersionItemName(val)
}

/*


 */
type QStyleOptionFrame__FrameFeature = int

//
const QStyleOptionFrame__None QStyleOptionFrame__FrameFeature = 0

//
const QStyleOptionFrame__Flat QStyleOptionFrame__FrameFeature = 1

//
const QStyleOptionFrame__Rounded QStyleOptionFrame__FrameFeature = 2

func (this *QStyleOptionFrame) FrameFeatureItemName(val int) string {
	switch val {
	case QStyleOptionFrame__None: // 0
		return "None"
	case QStyleOptionFrame__Flat: // 1
		return "Flat"
	case QStyleOptionFrame__Rounded: // 2
		return "Rounded"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionFrame_FrameFeatureItemName(val int) string {
	var nilthis *QStyleOptionFrame
	return nilthis.FrameFeatureItemName(val)
}

//  body block end

//  keep block begin

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
