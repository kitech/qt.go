package qtcore

// /usr/include/qt/QtCore/qscopedpointer.h
// #include <qscopedpointer.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
type QScopedPointerPodDeleter struct {
	*qtrt.CObject
}

func (this *QScopedPointerPodDeleter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QScopedPointerPodDeleter) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQScopedPointerPodDeleterFromPointer(cthis unsafe.Pointer) *QScopedPointerPodDeleter {
	return &QScopedPointerPodDeleter{&qtrt.CObject{cthis}}
}
func (*QScopedPointerPodDeleter) NewFromPointer(cthis unsafe.Pointer) *QScopedPointerPodDeleter {
	return NewQScopedPointerPodDeleterFromPointer(cthis)
}

// /usr/include/qt/QtCore/qscopedpointer.h:81
// index:0
// Public static inline
// void cleanup(void *)
func (this *QScopedPointerPodDeleter) Cleanup(pointer unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QScopedPointerPodDeleter7cleanupEPv", ffiqt.FFI_TYPE_POINTER, pointer)
	gopp.ErrPrint(err, rv)
}
func QScopedPointerPodDeleter_Cleanup(pointer unsafe.Pointer /*666*/) {
	var nilthis *QScopedPointerPodDeleter
	nilthis.Cleanup(pointer)
}

//  body block end
