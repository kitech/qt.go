package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaservice.h
// #include <qmediaservice.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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

/*

 */
type QMediaService struct {
	*qtcore.QObject
}
type QMediaService_ITF interface {
	qtcore.QObject_ITF
	QMediaService_PTR() *QMediaService
}

func (ptr *QMediaService) QMediaService_PTR() *QMediaService { return ptr }

func (this *QMediaService) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMediaService) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQMediaServiceFromPointer(cthis unsafe.Pointer) *QMediaService {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QMediaService{bcthis0}
}
func (*QMediaService) NewFromPointer(cthis unsafe.Pointer) *QMediaService {
	return NewQMediaServiceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaservice.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QMediaService) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMediaService10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaservice.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMediaService()

/*

 */
func DeleteQMediaService(this *QMediaService) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaServiceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaservice.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QMediaControl * requestControl(const char *)

/*
Returns a pointer to the media control implementing interface.

If the service does not implement the control, or if it is unavailable a null pointer is returned instead.

Controls must be returned to the service when no longer needed using the releaseControl() function.
*/
func (this *QMediaService) RequestControl(name string) *QMediaControl /*777 QMediaControl **/ {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaService14requestControlEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaControlFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qmediaservice.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void releaseControl(QMediaControl *)

/*
Releases a control back to the service.
*/
func (this *QMediaService) ReleaseControl(control QMediaControl_ITF /*777 QMediaControl **/) {
	var convArg0 unsafe.Pointer
	if control != nil && control.QMediaControl_PTR() != nil {
		convArg0 = control.QMediaControl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaService14releaseControlEP13QMediaControl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediaservice.h:74
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QMediaService(QObject *)

/*
Construct a media service with the given parent. This class is meant as a base class for Multimedia services so this constructor is protected.
*/
func (*QMediaService) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaService {
	return NewQMediaService(parent)
}
func NewQMediaService(parent qtcore.QObject_ITF /*777 QObject **/) *QMediaService {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMediaServiceC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaServiceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QMediaService")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11797() {
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
