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
type QStyleOptionTabWidgetFrame struct {
	*QStyleOption
}
type QStyleOptionTabWidgetFrame_ITF interface {
	QStyleOption_ITF
	QStyleOptionTabWidgetFrame_PTR() *QStyleOptionTabWidgetFrame
}

func (ptr *QStyleOptionTabWidgetFrame) QStyleOptionTabWidgetFrame_PTR() *QStyleOptionTabWidgetFrame {
	return ptr
}

func (this *QStyleOptionTabWidgetFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionTabWidgetFrame) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionTabWidgetFrameFromPointer(cthis unsafe.Pointer) *QStyleOptionTabWidgetFrame {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionTabWidgetFrame{bcthis0}
}
func (*QStyleOptionTabWidgetFrame) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionTabWidgetFrame {
	return NewQStyleOptionTabWidgetFrameFromPointer(cthis)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionTabWidgetFrame__StyleOptionType = int

//
const QStyleOptionTabWidgetFrame__Type QStyleOptionTabWidgetFrame__StyleOptionType = 11

func (this *QStyleOptionTabWidgetFrame) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionTabWidgetFrame__Type: // 11
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionTabWidgetFrame_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionTabWidgetFrame
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionTabWidgetFrame__StyleOptionVersion = int

// 1
const QStyleOptionTabWidgetFrame__Version QStyleOptionTabWidgetFrame__StyleOptionVersion = 2

func (this *QStyleOptionTabWidgetFrame) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionTabWidgetFrame__Version: // 2
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionTabWidgetFrame_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionTabWidgetFrame
	return nilthis.StyleOptionVersionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11009() {
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
