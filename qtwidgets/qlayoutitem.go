package qtwidgets

// /usr/include/qt/QtWidgets/qlayoutitem.h
// #include <qlayoutitem.h>
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
// size 16
type QLayoutItem struct {
	*qtrt.CObject
}
type QLayoutItem_ITF interface {
	QLayoutItem_PTR() *QLayoutItem
}

func (ptr *QLayoutItem) QLayoutItem_PTR() *QLayoutItem { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QLayoutItemFromptr(cthis Voidptr) *QLayoutItem {
	return &QLayoutItem{&qtrt.CObject{cthis}}
}
func (*QLayoutItem) Fromptr(cthis Voidptr) *QLayoutItem {
	return QLayoutItemFromptr(cthis)
}

func DeleteQLayoutItem(this *QLayoutItem) {
	rv, err := qtrt.Qtcc3(2677748963, "_ZN11QLayoutItemD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10199() {
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
