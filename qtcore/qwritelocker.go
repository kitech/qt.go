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
// extern C begin: 5
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
type QWriteLocker struct {
	*qtrt.CObject
}
type QWriteLocker_ITF interface {
	QWriteLocker_PTR() *QWriteLocker
}

func (ptr *QWriteLocker) QWriteLocker_PTR() *QWriteLocker { return ptr }

func (this *QWriteLocker) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWriteLocker) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWriteLockerFromPointer(cthis unsafe.Pointer) *QWriteLocker {
	return &QWriteLocker{&qtrt.CObject{cthis}}
}
func (*QWriteLocker) NewFromPointer(cthis unsafe.Pointer) *QWriteLocker {
	return NewQWriteLockerFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_10517() {
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
