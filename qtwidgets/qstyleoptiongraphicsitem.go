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
type QStyleOptionGraphicsItem struct {
	*QStyleOption
}
type QStyleOptionGraphicsItem_ITF interface {
	QStyleOption_ITF
	QStyleOptionGraphicsItem_PTR() *QStyleOptionGraphicsItem
}

func (ptr *QStyleOptionGraphicsItem) QStyleOptionGraphicsItem_PTR() *QStyleOptionGraphicsItem {
	return ptr
}

func (this *QStyleOptionGraphicsItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionGraphicsItem) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionGraphicsItemFromPointer(cthis unsafe.Pointer) *QStyleOptionGraphicsItem {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionGraphicsItem{bcthis0}
}
func (*QStyleOptionGraphicsItem) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionGraphicsItem {
	return NewQStyleOptionGraphicsItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:669
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionGraphicsItem()

/*

 */
func NewQStyleOptionGraphicsItem() *QStyleOptionGraphicsItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItemC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionGraphicsItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:673
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionGraphicsItem(int)

/*

 */
func NewQStyleOptionGraphicsItem_1(version int) *QStyleOptionGraphicsItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItemC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionGraphicsItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:671
// index:0
// Public static Visibility=Default Availability=Available
// [8] qreal levelOfDetailFromTransform(const QTransform &)

/*

 */
func (this *QStyleOptionGraphicsItem) LevelOfDetailFromTransform(worldTransform qtgui.QTransform_ITF) float64 {
	var convArg0 unsafe.Pointer
	if worldTransform != nil && worldTransform.QTransform_PTR() != nil {
		convArg0 = worldTransform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItem26levelOfDetailFromTransformERK10QTransform", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QStyleOptionGraphicsItem_LevelOfDetailFromTransform(worldTransform qtgui.QTransform_ITF) float64 {
	var nilthis *QStyleOptionGraphicsItem
	rv := nilthis.LevelOfDetailFromTransform(worldTransform)
	return rv
}

func DeleteQStyleOptionGraphicsItem(this *QStyleOptionGraphicsItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum is used to hold information about the type of the style option, and is defined for each QStyleOption subclass.

QStyleOption::TypeSO_DefaultThe type of style option provided (SO_Default for this class).


The type is used internally by QStyleOption, its subclasses, and qstyleoption_cast() to determine the type of style option. In general you do not need to worry about this unless you want to create your own QStyleOption subclass and your own styles.

See also StyleOptionVersion.

*/
type QStyleOptionGraphicsItem__StyleOptionType = int

//
const QStyleOptionGraphicsItem__Type QStyleOptionGraphicsItem__StyleOptionType = 15

/*
This enum is used to hold information about the version of the style option, and is defined for each QStyleOption subclass.



The version is used by QStyleOption subclasses to implement extensions without breaking compatibility. If you use qstyleoption_cast(), you normally do not need to check it.

See also StyleOptionType.

*/
type QStyleOptionGraphicsItem__StyleOptionVersion = int

// 1
const QStyleOptionGraphicsItem__Version QStyleOptionGraphicsItem__StyleOptionVersion = 1

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
