package qtcore

// /usr/include/qt/QtCore/qsharedpointer_impl.h
// #include <qsharedpointer_impl.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
type ExternalRefCountData struct {
	*qtrt.CObject
}

func (this *ExternalRefCountData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *ExternalRefCountData) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewExternalRefCountDataFromPointer(cthis unsafe.Pointer) *ExternalRefCountData {
	return &ExternalRefCountData{&qtrt.CObject{cthis}}
}
func (*ExternalRefCountData) NewFromPointer(cthis unsafe.Pointer) *ExternalRefCountData {
	return NewExternalRefCountDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:154
// index:0
// Public inline
// void ExternalRefCountData(Qt::Initialization)
func NewExternalRefCountData(arg0 int) *ExternalRefCountData {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountDataC2EN2Qt14InitializationE", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewExternalRefCountDataFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:155
// index:0
// Public inline
// void ~ExternalRefCountData()
func DeleteExternalRefCountData(*ExternalRefCountData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountDataD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:157
// index:0
// Public inline
// void destroy()
func (this *ExternalRefCountData) Destroy() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountData7destroyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:160
// index:0
// Public static
// QtSharedPointer::ExternalRefCountData * getAndRef(const QObject *)
func (this *ExternalRefCountData) GetAndRef(arg0 *QObject /*777 const QObject **/) unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountData9getAndRefEPK7QObject", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return unsafe.Pointer(uintptr(rv))
}
func ExternalRefCountData_GetAndRef(arg0 *QObject /*777 const QObject **/) unsafe.Pointer /*666*/ {
	var nilthis *ExternalRefCountData
	rv := nilthis.GetAndRef(arg0)
	return rv
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:161
// index:0
// Public
// void setQObjectShared(const QObject *, bool)
func (this *ExternalRefCountData) SetQObjectShared(arg0 *QObject /*777 const QObject **/, enable bool) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountData16setQObjectSharedEPK7QObjectb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsharedpointer_impl.h:162
// index:0
// Public
// void checkQObjectShared(const QObject *)
func (this *ExternalRefCountData) CheckQObjectShared(arg0 *QObject /*777 const QObject **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QtSharedPointer20ExternalRefCountData18checkQObjectSharedEPK7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
