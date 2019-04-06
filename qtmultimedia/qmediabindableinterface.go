package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediabindableinterface.h
// #include <qmediabindableinterface.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

// bool setMediaObject(QMediaObject *)
func (this *QMediaBindableInterface) InheritSetMediaObject(f func(object *QMediaObject /*777 QMediaObject **/) bool) {
	qtrt.SetAllInheritCallback(this, "setMediaObject", f)
}

/*

 */
type QMediaBindableInterface struct {
	*qtrt.CObject
}
type QMediaBindableInterface_ITF interface {
	QMediaBindableInterface_PTR() *QMediaBindableInterface
}

func (ptr *QMediaBindableInterface) QMediaBindableInterface_PTR() *QMediaBindableInterface { return ptr }

func (this *QMediaBindableInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaBindableInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaBindableInterfaceFromPointer(cthis unsafe.Pointer) *QMediaBindableInterface {
	return &QMediaBindableInterface{&qtrt.CObject{cthis}}
}
func (*QMediaBindableInterface) NewFromPointer(cthis unsafe.Pointer) *QMediaBindableInterface {
	return NewQMediaBindableInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediabindableinterface.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaBindableInterface()

/*

 */
func DeleteQMediaBindableInterface(this *QMediaBindableInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaBindableInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediabindableinterface.h:55
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QMediaObject * mediaObject() const

/*
Return the currently attached media object.

See also setMediaObject().
*/
func (this *QMediaBindableInterface) MediaObject() *QMediaObject /*777 QMediaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QMediaBindableInterface11mediaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediabindableinterface.h:59
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [1] bool setMediaObject(QMediaObject *)

/*
Attaches to the media object. Returns true if attached successfully, otherwise returns false.

See also mediaObject().
*/
func (this *QMediaBindableInterface) SetMediaObject(object QMediaObject_ITF /*777 QMediaObject **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QMediaObject_PTR() != nil {
		convArg0 = object.QMediaObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QMediaBindableInterface14setMediaObjectEP12QMediaObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init_unused_11767() {
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
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
