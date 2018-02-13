package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QPointingDeviceUniqueId struct {
	*qtrt.CObject
}
type QPointingDeviceUniqueId_ITF interface {
	QPointingDeviceUniqueId_PTR() *QPointingDeviceUniqueId
}

func (ptr *QPointingDeviceUniqueId) QPointingDeviceUniqueId_PTR() *QPointingDeviceUniqueId { return ptr }

func (this *QPointingDeviceUniqueId) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPointingDeviceUniqueId) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPointingDeviceUniqueIdFromPointer(cthis unsafe.Pointer) *QPointingDeviceUniqueId {
	return &QPointingDeviceUniqueId{&qtrt.CObject{cthis}}
}
func (*QPointingDeviceUniqueId) NewFromPointer(cthis unsafe.Pointer) *QPointingDeviceUniqueId {
	return NewQPointingDeviceUniqueIdFromPointer(cthis)
}

// /usr/include/qt/QtGui/qevent.h:809
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QPointingDeviceUniqueId()
func NewQPointingDeviceUniqueId() *QPointingDeviceUniqueId {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QPointingDeviceUniqueIdC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPointingDeviceUniqueIdFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPointingDeviceUniqueId)
	return gothis
}

// /usr/include/qt/QtGui/qevent.h:813
// index:0
// Public static Visibility=Default Availability=Available
// [8] QPointingDeviceUniqueId fromNumericId(qint64)
func (this *QPointingDeviceUniqueId) FromNumericId(id int64) *QPointingDeviceUniqueId /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QPointingDeviceUniqueId13fromNumericIdEx", qtrt.FFI_TYPE_POINTER, id)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPointingDeviceUniqueIdFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPointingDeviceUniqueId)
	return rv2
}
func QPointingDeviceUniqueId_FromNumericId(id int64) *QPointingDeviceUniqueId /*123*/ {
	var nilthis *QPointingDeviceUniqueId
	rv := nilthis.FromNumericId(id)
	return rv
}

// /usr/include/qt/QtGui/qevent.h:815
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QPointingDeviceUniqueId) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QPointingDeviceUniqueId7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qevent.h:816
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 numericId()
func (this *QPointingDeviceUniqueId) NumericId() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QPointingDeviceUniqueId9numericIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

func DeleteQPointingDeviceUniqueId(this *QPointingDeviceUniqueId) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QPointingDeviceUniqueIdD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

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
		qtcore.KeepMe()
	}
}

//  keep block end
