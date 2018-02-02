package qtgui

// /usr/include/qt/QtGui/qaccessiblebridge.h
// #include <qaccessiblebridge.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QAccessibleBridge struct {
	*qtrt.CObject
}

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
func DeleteQAccessibleBridge(this *QAccessibleBridge) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleBridgeD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setRootObject(QAccessibleInterface *)
func (this *QAccessibleBridge) SetRootObject(arg0 *QAccessibleInterface /*777 QAccessibleInterface **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleBridge13setRootObjectEP20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessiblebridge.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void notifyAccessibilityUpdate(QAccessibleEvent *)
func (this *QAccessibleBridge) NotifyAccessibilityUpdate(event *QAccessibleEvent /*777 QAccessibleEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAccessibleBridge25notifyAccessibilityUpdateEP16QAccessibleEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
