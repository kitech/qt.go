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
type QStyleOptionTabBarBase struct {
	*QStyleOption
}
type QStyleOptionTabBarBase_ITF interface {
	QStyleOption_ITF
	QStyleOptionTabBarBase_PTR() *QStyleOptionTabBarBase
}

func (ptr *QStyleOptionTabBarBase) QStyleOptionTabBarBase_PTR() *QStyleOptionTabBarBase { return ptr }

func (this *QStyleOptionTabBarBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionTabBarBase) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionTabBarBaseFromPointer(cthis unsafe.Pointer) *QStyleOptionTabBarBase {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionTabBarBase{bcthis0}
}
func (*QStyleOptionTabBarBase) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionTabBarBase {
	return NewQStyleOptionTabBarBaseFromPointer(cthis)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionTabBarBase__StyleOptionType = int

//
const QStyleOptionTabBarBase__Type QStyleOptionTabBarBase__StyleOptionType = 12

func (this *QStyleOptionTabBarBase) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionTabBarBase__Type: // 12
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionTabBarBase_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionTabBarBase
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionTabBarBase__StyleOptionVersion = int

// 1
const QStyleOptionTabBarBase__Version QStyleOptionTabBarBase__StyleOptionVersion = 2

func (this *QStyleOptionTabBarBase) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionTabBarBase__Version: // 2
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionTabBarBase_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionTabBarBase
	return nilthis.StyleOptionVersionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11011() {
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
