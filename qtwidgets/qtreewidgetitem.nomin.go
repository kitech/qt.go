//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qtreewidget.h
// #include <qtreewidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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

// /usr/include/qt/QtWidgets/qtreewidget.h:123
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString whatsThis(int) const

/*

 */
func (this *QTreeWidgetItem) WhatsThis(column int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTreeWidgetItem9whatsThisEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qtreewidget.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWhatsThis(int, const QString &)

/*

 */
func (this *QTreeWidgetItem) SetWhatsThis(column int, whatsThis string) {
	var tmpArg1 = qtcore.NewQString5(whatsThis)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTreeWidgetItem12setWhatsThisEiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11360() {
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
