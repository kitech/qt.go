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
type QStyleOptionSpinBox struct {
	*QStyleOptionComplex
}
type QStyleOptionSpinBox_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionSpinBox_PTR() *QStyleOptionSpinBox
}

func (ptr *QStyleOptionSpinBox) QStyleOptionSpinBox_PTR() *QStyleOptionSpinBox { return ptr }

func (this *QStyleOptionSpinBox) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionSpinBox) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionSpinBoxFromPointer(cthis unsafe.Pointer) *QStyleOptionSpinBox {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSpinBox{bcthis0}
}
func (*QStyleOptionSpinBox) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionSpinBox {
	return NewQStyleOptionSpinBoxFromPointer(cthis)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionSpinBox__StyleOptionType = int

//
const QStyleOptionSpinBox__Type QStyleOptionSpinBox__StyleOptionType = 983042

func (this *QStyleOptionSpinBox) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionSpinBox__Type: // 983042
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionSpinBox_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionSpinBox
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionSpinBox__StyleOptionVersion = int

// 1
const QStyleOptionSpinBox__Version QStyleOptionSpinBox__StyleOptionVersion = 1

func (this *QStyleOptionSpinBox) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionSpinBox__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionSpinBox_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionSpinBox
	return nilthis.StyleOptionVersionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11037() {
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
