package qtcore

// /usr/include/qt/QtCore/qreadwritelock.h
// #include <qreadwritelock.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QReadLocker struct {
	*qtrt.CObject
}
type QReadLocker_ITF interface {
	QReadLocker_PTR() *QReadLocker
}

func (ptr *QReadLocker) QReadLocker_PTR() *QReadLocker { return ptr }

func (this *QReadLocker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QReadLocker) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQReadLockerFromPointer(cthis unsafe.Pointer) *QReadLocker {
	return &QReadLocker{&qtrt.CObject{cthis}}
}
func (*QReadLocker) NewFromPointer(cthis unsafe.Pointer) *QReadLocker {
	return NewQReadLockerFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_10515() {
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
}

//  keep block end
