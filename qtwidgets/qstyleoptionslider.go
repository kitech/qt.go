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
// extern C begin: 1
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
type QStyleOptionSlider struct {
	*QStyleOptionComplex
}
type QStyleOptionSlider_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionSlider_PTR() *QStyleOptionSlider
}

func (ptr *QStyleOptionSlider) QStyleOptionSlider_PTR() *QStyleOptionSlider { return ptr }

func (this *QStyleOptionSlider) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionSlider) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionSliderFromPointer(cthis unsafe.Pointer) *QStyleOptionSlider {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionSlider{bcthis0}
}
func (*QStyleOptionSlider) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionSlider {
	return NewQStyleOptionSliderFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:533
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionSlider()

/*

 */
func NewQStyleOptionSlider() *QStyleOptionSlider {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionSliderC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionSlider)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:537
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionSlider(int)

/*

 */
func NewQStyleOptionSlider_1(version int) *QStyleOptionSlider {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionSliderC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionSliderFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionSlider)
	return gothis
}

func DeleteQStyleOptionSlider(this *QStyleOptionSlider) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStyleOptionSliderD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionSlider__StyleOptionType = int

//
const QStyleOptionSlider__Type QStyleOptionSlider__StyleOptionType = 983041

func (this *QStyleOptionSlider) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionSlider__Type: // 983041
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionSlider_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionSlider
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionSlider__StyleOptionVersion = int

// 1
const QStyleOptionSlider__Version QStyleOptionSlider__StyleOptionVersion = 1

func (this *QStyleOptionSlider) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionSlider__Version: // 1
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionSlider_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionSlider
	return nilthis.StyleOptionVersionItemName(val)
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
