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
type QStyleOptionComplex struct {
	*QStyleOption
}
type QStyleOptionComplex_ITF interface {
	QStyleOption_ITF
	QStyleOptionComplex_PTR() *QStyleOptionComplex
}

func (ptr *QStyleOptionComplex) QStyleOptionComplex_PTR() *QStyleOptionComplex { return ptr }

func (this *QStyleOptionComplex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionComplex) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionComplexFromPointer(cthis unsafe.Pointer) *QStyleOptionComplex {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionComplex{bcthis0}
}
func (*QStyleOptionComplex) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionComplex {
	return NewQStyleOptionComplexFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:509
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionComplex(int, int)

/*

 */
func NewQStyleOptionComplex(version int, type_ int) *QStyleOptionComplex {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionComplexC2Eii", qtrt.FFI_TYPE_POINTER, version, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionComplexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionComplex)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:509
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionComplex(int, int)

/*

 */
func NewQStyleOptionComplex__() *QStyleOptionComplex {
	// arg: 0, int=Int, =Invalid,
	version := 0 /*QStyleOptionComplex::Version*/
	// arg: 1, int=Int, =Invalid,
	type_ := QStyleOption__SO_Complex
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionComplexC2Eii", qtrt.FFI_TYPE_POINTER, version, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionComplexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionComplex)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:509
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionComplex(int, int)

/*

 */
func NewQStyleOptionComplex__1(version int) *QStyleOptionComplex {
	// arg: 1, int=Int, =Invalid,
	type_ := QStyleOption__SO_Complex
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionComplexC2Eii", qtrt.FFI_TYPE_POINTER, version, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionComplexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionComplex)
	return gothis
}

func DeleteQStyleOptionComplex(this *QStyleOptionComplex) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QStyleOptionComplexD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionComplex__StyleOptionType = int

//
const QStyleOptionComplex__Type QStyleOptionComplex__StyleOptionType = 983040

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionComplex__StyleOptionVersion = int

// 1
const QStyleOptionComplex__Version QStyleOptionComplex__StyleOptionVersion = 1

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
