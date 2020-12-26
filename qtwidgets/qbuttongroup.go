// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qbuttongroup.h
// #include <qbuttongroup.h>
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
// size 16
type QButtonGroup struct {
	*qtcore.QObject
}
type QButtonGroup_ITF interface {
	qtcore.QObject_ITF
	QButtonGroup_PTR() *QButtonGroup
}

func (ptr *QButtonGroup) QButtonGroup_PTR() *QButtonGroup { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QButtonGroupFromptr(cthis Voidptr) *QButtonGroup {
	bcthis0 := qtcore.QObjectFromptr(cthis)
	return &QButtonGroup{bcthis0}
}
func (*QButtonGroup) Fromptr(cthis Voidptr) *QButtonGroup {
	return QButtonGroupFromptr(cthis)
}

func DeleteQButtonGroup(this *QButtonGroup) {
	rv, err := qtrt.Qtcc3(788836440, "_ZN12QButtonGroupD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
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
