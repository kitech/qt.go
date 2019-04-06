//  header block begin

// +build !minimal

package qtgui

// /usr/include/qt/QtGui/qstandarditemmodel.h
// #include <qstandarditemmodel.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// /usr/include/qt/QtGui/qstandarditemmodel.h:99
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString whatsThis() const

/*

 */
func (this *QStandardItem) WhatsThis() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem9whatsThisEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:102
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)

/*

 */
func (this *QStandardItem) SetWhatsThis(whatsThis string) {
	var tmpArg0 = qtcore.NewQString5(whatsThis)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem12setWhatsThisERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:184
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDragEnabled() const

/*

 */
func (this *QStandardItem) IsDragEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem13isDragEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDragEnabled(bool)

/*

 */
func (this *QStandardItem) SetDragEnabled(dragEnabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem14setDragEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dragEnabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:189
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDropEnabled() const

/*

 */
func (this *QStandardItem) IsDropEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QStandardItem13isDropEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDropEnabled(bool)

/*

 */
func (this *QStandardItem) SetDropEnabled(dropEnabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QStandardItem14setDropEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dropEnabled)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10936() {
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
}

//  keep block end
