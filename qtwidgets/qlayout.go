package qtwidgets

// /usr/include/qt/QtWidgets/qlayout.h
// #include <qlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
// size 32
type QLayout struct {
	*qtcore.QObject
	*QLayoutItem
}
type QLayout_ITF interface {
	qtcore.QObject_ITF
	QLayoutItem_ITF
	QLayout_PTR() *QLayout
}

func (ptr *QLayout) QLayout_PTR() *QLayout { return ptr }

func (this *QLayout) GetCthis() Voidptr {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QLayout) SetCthis(cthis Voidptr) {
	this.QObject = qtcore.QObjectFromptr(cthis)
	this.QLayoutItem = QLayoutItemFromptr(cthis)
}
func (this *QLayout) Addr() Voidptr {
	if this == nil {
		return nil
	} else {
		return this.QObject.Addr()
	}
}
func QLayoutFromptr(cthis Voidptr) *QLayout {
	bcthis0 := qtcore.QObjectFromptr(cthis)
	bcthis1 := QLayoutItemFromptr(cthis)
	return &QLayout{bcthis0, bcthis1}
}
func (*QLayout) Fromptr(cthis Voidptr) *QLayout {
	return QLayoutFromptr(cthis)
}

func DeleteQLayout(this *QLayout) {
	rv, err := qtrt.Qtcc3(925367012, "_ZN7QLayoutD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QLayout__SizeConstraint = int

//
const QLayout__SetDefaultConstraint QLayout__SizeConstraint = 0

//
const QLayout__SetNoConstraint QLayout__SizeConstraint = 1

//
const QLayout__SetMinimumSize QLayout__SizeConstraint = 2

//
const QLayout__SetFixedSize QLayout__SizeConstraint = 3

//
const QLayout__SetMaximumSize QLayout__SizeConstraint = 4

//
const QLayout__SetMinAndMaxSize QLayout__SizeConstraint = 5

func (this *QLayout) SizeConstraintItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QLayout_SizeConstraintItemName(val int) string {
	var nilthis *QLayout
	return nilthis.SizeConstraintItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10205() {
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
