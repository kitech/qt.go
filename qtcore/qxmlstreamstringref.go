package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
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
type QXmlStreamStringRef struct {
	*qtrt.CObject
}
type QXmlStreamStringRef_ITF interface {
	QXmlStreamStringRef_PTR() *QXmlStreamStringRef
}

func (ptr *QXmlStreamStringRef) QXmlStreamStringRef_PTR() *QXmlStreamStringRef { return ptr }

func (this *QXmlStreamStringRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamStringRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQXmlStreamStringRefFromPointer(cthis unsafe.Pointer) *QXmlStreamStringRef {
	return &QXmlStreamStringRef{&qtrt.CObject{cthis}}
}
func (*QXmlStreamStringRef) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamStringRef {
	return NewQXmlStreamStringRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QXmlStreamStringRef()

/*

 */
func (*QXmlStreamStringRef) NewForInherit() *QXmlStreamStringRef {
	return NewQXmlStreamStringRef()
}
func NewQXmlStreamStringRef() *QXmlStreamStringRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:59
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QXmlStreamStringRef(const QStringRef &)

/*

 */
func (*QXmlStreamStringRef) NewForInherit1(aString QStringRef_ITF) *QXmlStreamStringRef {
	return NewQXmlStreamStringRef1(aString)
}
func NewQXmlStreamStringRef1(aString QStringRef_ITF) *QXmlStreamStringRef {
	var convArg0 unsafe.Pointer
	if aString != nil && aString.QStringRef_PTR() != nil {
		convArg0 = aString.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2ERK10QStringRef", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:61
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QXmlStreamStringRef(const QString &)

/*

 */
func (*QXmlStreamStringRef) NewForInherit2(aString string) *QXmlStreamStringRef {
	return NewQXmlStreamStringRef2(aString)
}
func NewQXmlStreamStringRef2(aString string) *QXmlStreamStringRef {
	var tmpArg0 = NewQString5(aString)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:63
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QXmlStreamStringRef(QString &&)

/*

 */
func (*QXmlStreamStringRef) NewForInherit3(aString unsafe.Pointer /*333*/) *QXmlStreamStringRef {
	return NewQXmlStreamStringRef3(aString)
}
func NewQXmlStreamStringRef3(aString unsafe.Pointer /*333*/) *QXmlStreamStringRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2EO7QString", qtrt.FFI_TYPE_POINTER, aString)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QXmlStreamStringRef & operator=(QXmlStreamStringRef &&)

/*

 */
func (this *QXmlStreamStringRef) Operator_equal(other unsafe.Pointer /*333*/) *QXmlStreamStringRef {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamStringRefaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQXmlStreamStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQXmlStreamStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:75
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QXmlStreamStringRef & operator=(const QXmlStreamStringRef &)

/*

 */
func (this *QXmlStreamStringRef) Operator_equal1(other QXmlStreamStringRef_ITF) *QXmlStreamStringRef {
	var convArg0 unsafe.Pointer
	if other != nil && other.QXmlStreamStringRef_PTR() != nil {
		convArg0 = other.QXmlStreamStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamStringRefaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQXmlStreamStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQXmlStreamStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QXmlStreamStringRef()

/*

 */
func DeleteQXmlStreamStringRef(this *QXmlStreamStringRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamStringRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QXmlStreamStringRef &)

/*

 */
func (this *QXmlStreamStringRef) Swap(other QXmlStreamStringRef_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QXmlStreamStringRef_PTR() != nil {
		convArg0 = other.QXmlStreamStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamStringRef4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clear()

/*

 */
func (this *QXmlStreamStringRef) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamStringRef5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QString * string() const

/*

 */
func (this *QXmlStreamStringRef) String() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QXmlStreamStringRef6stringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qxmlstream.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int position() const

/*

 */
func (this *QXmlStreamStringRef) Position() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QXmlStreamStringRef8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qxmlstream.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size() const

/*

 */
func (this *QXmlStreamStringRef) Size() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QXmlStreamStringRef4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end

//  keep block begin

func init_unused_10613() {
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
