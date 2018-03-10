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
type QStyleOptionTitleBar struct {
	*QStyleOptionComplex
}
type QStyleOptionTitleBar_ITF interface {
	QStyleOptionComplex_ITF
	QStyleOptionTitleBar_PTR() *QStyleOptionTitleBar
}

func (ptr *QStyleOptionTitleBar) QStyleOptionTitleBar_PTR() *QStyleOptionTitleBar { return ptr }

func (this *QStyleOptionTitleBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOptionComplex.GetCthis()
	}
}
func (this *QStyleOptionTitleBar) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOptionComplex = NewQStyleOptionComplexFromPointer(cthis)
}
func NewQStyleOptionTitleBarFromPointer(cthis unsafe.Pointer) *QStyleOptionTitleBar {
	bcthis0 := NewQStyleOptionComplexFromPointer(cthis)
	return &QStyleOptionTitleBar{bcthis0}
}
func (*QStyleOptionTitleBar) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionTitleBar {
	return NewQStyleOptionTitleBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:619
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionTitleBar()

/*

 */
func NewQStyleOptionTitleBar() *QStyleOptionTitleBar {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionTitleBarC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTitleBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTitleBar)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:623
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionTitleBar(int)

/*

 */
func NewQStyleOptionTitleBar_1(version int) *QStyleOptionTitleBar {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionTitleBarC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionTitleBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionTitleBar)
	return gothis
}

func DeleteQStyleOptionTitleBar(this *QStyleOptionTitleBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionTitleBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionTitleBar__StyleOptionType = int

//
const QStyleOptionTitleBar__Type QStyleOptionTitleBar__StyleOptionType = 983045

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionTitleBar__StyleOptionVersion = int

// 1
const QStyleOptionTitleBar__Version QStyleOptionTitleBar__StyleOptionVersion = 1

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
