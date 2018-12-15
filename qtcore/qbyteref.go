package qtcore

// /usr/include/qt/QtCore/qbytearray.h
// #include <qbytearray.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 141
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
type QByteRef struct {
	*qtrt.CObject
}
type QByteRef_ITF interface {
	QByteRef_PTR() *QByteRef
}

func (ptr *QByteRef) QByteRef_PTR() *QByteRef { return ptr }

func (this *QByteRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QByteRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQByteRefFromPointer(cthis unsafe.Pointer) *QByteRef {
	return &QByteRef{&qtrt.CObject{cthis}}
}
func (*QByteRef) NewFromPointer(cthis unsafe.Pointer) *QByteRef {
	return NewQByteRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qbytearray.h:528
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QByteRef & operator=(char)

/*

 */
func (this *QByteRef) Operator_equal(c byte) *QByteRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QByteRefaSEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteRef)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:531
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QByteRef & operator=(const QByteRef &)

/*

 */
func (this *QByteRef) Operator_equal1(c QByteRef_ITF) *QByteRef {
	var convArg0 unsafe.Pointer
	if c != nil && c.QByteRef_PTR() != nil {
		convArg0 = c.QByteRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QByteRefaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteRef)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:534
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(char) const

/*

 */
func (this *QByteRef) Operator_equal_equal(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QByteRefeqEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:536
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(char) const

/*

 */
func (this *QByteRef) Operator_not_equal(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QByteRefneEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:538
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>(char) const

/*

 */
func (this *QByteRef) Operator_greater_than(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QByteRefgtEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:540
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator>=(char) const

/*

 */
func (this *QByteRef) Operator_greater_than_equal(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QByteRefgeEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:542
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(char) const

/*

 */
func (this *QByteRef) Operator_less_than(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QByteRefltEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qbytearray.h:544
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<=(char) const

/*

 */
func (this *QByteRef) Operator_less_than_equal(c byte) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QByteRefleEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), c)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQByteRef(this *QByteRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QByteRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
