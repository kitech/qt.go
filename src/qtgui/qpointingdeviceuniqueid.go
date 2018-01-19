//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
type QPointingDeviceUniqueId struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:809
// index:0
// inline
// void QPointingDeviceUniqueId()
func NewQPointingDeviceUniqueId() *QPointingDeviceUniqueId {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QPointingDeviceUniqueIdC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QPointingDeviceUniqueId{cthis}
}

// /usr/include/qt/QtGui/qevent.h:813
// index:0
// static
// QPointingDeviceUniqueId fromNumericId(qint64)
func (this *QPointingDeviceUniqueId) FromNumericId(id int64) {
	// 0: (qint64 id), (id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QPointingDeviceUniqueId13fromNumericIdEx", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPointingDeviceUniqueId_FromNumericId(id int64) {
	// 0: (qint64 id), (id)
	var nilthis *QPointingDeviceUniqueId
	nilthis.FromNumericId(id)
}

// /usr/include/qt/QtGui/qevent.h:815
// index:0
// inline
// bool isValid()
func (this *QPointingDeviceUniqueId) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QPointingDeviceUniqueId7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:816
// index:0
// qint64 numericId()
func (this *QPointingDeviceUniqueId) NumericId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QPointingDeviceUniqueId9numericIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
