//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractbutton.h
// #include <qabstractbutton.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 288
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

// /usr/include/qt/QtWidgets/qabstractbutton.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QButtonGroup * group() const

/*
Returns the group that this button belongs to.

If the button is not a member of any QButtonGroup, this function returns 0.

See also QButtonGroup.
*/
func (this *QAbstractButton) Group() *QButtonGroup /*777 QButtonGroup **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QAbstractButton5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQButtonGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_10986() {
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
