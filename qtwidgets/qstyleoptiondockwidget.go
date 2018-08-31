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
type QStyleOptionDockWidget struct {
	*QStyleOption
}
type QStyleOptionDockWidget_ITF interface {
	QStyleOption_ITF
	QStyleOptionDockWidget_PTR() *QStyleOptionDockWidget
}

func (ptr *QStyleOptionDockWidget) QStyleOptionDockWidget_PTR() *QStyleOptionDockWidget { return ptr }

func (this *QStyleOptionDockWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionDockWidget) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionDockWidgetFromPointer(cthis unsafe.Pointer) *QStyleOptionDockWidget {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionDockWidget{bcthis0}
}
func (*QStyleOptionDockWidget) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionDockWidget {
	return NewQStyleOptionDockWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:391
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionDockWidget()

/*

 */
func (*QStyleOptionDockWidget) NewForInherit() *QStyleOptionDockWidget {
	return NewQStyleOptionDockWidget()
}
func NewQStyleOptionDockWidget() *QStyleOptionDockWidget {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionDockWidgetC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionDockWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionDockWidget)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:395
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionDockWidget(int)

/*

 */
func (*QStyleOptionDockWidget) NewForInherit_1(version int) *QStyleOptionDockWidget {
	return NewQStyleOptionDockWidget_1(version)
}
func NewQStyleOptionDockWidget_1(version int) *QStyleOptionDockWidget {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionDockWidgetC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionDockWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionDockWidget)
	return gothis
}

func DeleteQStyleOptionDockWidget(this *QStyleOptionDockWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QStyleOptionDockWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionDockWidget__StyleOptionType = int

//
const QStyleOptionDockWidget__Type QStyleOptionDockWidget__StyleOptionType = 9

func (this *QStyleOptionDockWidget) StyleOptionTypeItemName(val int) string {
	switch val {
	case QStyleOptionDockWidget__Type: // 9
		return "Type"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionDockWidget_StyleOptionTypeItemName(val int) string {
	var nilthis *QStyleOptionDockWidget
	return nilthis.StyleOptionTypeItemName(val)
}

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionDockWidget__StyleOptionVersion = int

// 1
const QStyleOptionDockWidget__Version QStyleOptionDockWidget__StyleOptionVersion = 2

func (this *QStyleOptionDockWidget) StyleOptionVersionItemName(val int) string {
	switch val {
	case QStyleOptionDockWidget__Version: // 2
		return "Version"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStyleOptionDockWidget_StyleOptionVersionItemName(val int) string {
	var nilthis *QStyleOptionDockWidget
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
