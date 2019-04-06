package qtcore

// /usr/include/qt/QtCore/qthreadstorage.h
// #include <qthreadstorage.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QThreadStorageData struct {
	*qtrt.CObject
}
type QThreadStorageData_ITF interface {
	QThreadStorageData_PTR() *QThreadStorageData
}

func (ptr *QThreadStorageData) QThreadStorageData_PTR() *QThreadStorageData { return ptr }

func (this *QThreadStorageData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QThreadStorageData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQThreadStorageDataFromPointer(cthis unsafe.Pointer) *QThreadStorageData {
	return &QThreadStorageData{&qtrt.CObject{cthis}}
}
func (*QThreadStorageData) NewFromPointer(cthis unsafe.Pointer) *QThreadStorageData {
	return NewQThreadStorageDataFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_10599() {
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
