package qtgui

// /usr/include/qt/QtGui/qaccessiblebridge.h
// #include <qaccessiblebridge.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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

/*

 */
type QAccessibleBridge struct {
	*qtrt.CObject
}
type QAccessibleBridge_ITF interface {
	QAccessibleBridge_PTR() *QAccessibleBridge
}

func (ptr *QAccessibleBridge) QAccessibleBridge_PTR() *QAccessibleBridge { return ptr }

func (this *QAccessibleBridge) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleBridge) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleBridgeFromPointer(cthis unsafe.Pointer) *QAccessibleBridge {
	return &QAccessibleBridge{&qtrt.CObject{cthis}}
}
func (*QAccessibleBridge) NewFromPointer(cthis unsafe.Pointer) *QAccessibleBridge {
	return NewQAccessibleBridgeFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:58
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleBridge()

/*

 */
func DeleteQAccessibleBridge(this *QAccessibleBridge) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessibleBridgeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setRootObject(QAccessibleInterface *)

/*

 */
func (this *QAccessibleBridge) SetRootObject(arg0 QAccessibleInterface_ITF /*777 QAccessibleInterface **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QAccessibleInterface_PTR() != nil {
		convArg0 = arg0.QAccessibleInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessibleBridge13setRootObjectEP20QAccessibleInterface", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void notifyAccessibilityUpdate(QAccessibleEvent *)

/*

 */
func (this *QAccessibleBridge) NotifyAccessibilityUpdate(event QAccessibleEvent_ITF /*777 QAccessibleEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QAccessibleEvent_PTR() != nil {
		convArg0 = event.QAccessibleEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QAccessibleBridge25notifyAccessibilityUpdateEP16QAccessibleEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10835() {
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
