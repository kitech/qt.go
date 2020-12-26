package qtcore

// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
// size 16
type QMetaMethod struct {
	*qtrt.CObject
}
type QMetaMethod_ITF interface {
	QMetaMethod_PTR() *QMetaMethod
}

func (ptr *QMetaMethod) QMetaMethod_PTR() *QMetaMethod { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QMetaMethodFromptr(cthis Voidptr) *QMetaMethod {
	return &QMetaMethod{&qtrt.CObject{cthis}}
}
func (*QMetaMethod) Fromptr(cthis Voidptr) *QMetaMethod {
	return QMetaMethodFromptr(cthis)
}

// /usr/include/qt/QtCore/qmetaobject.h:59
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QByteArray methodSignature() const

/*
 */
func (this *QMetaMethod) MethodSignature() *QByteArray /*123*/ {
	sretobj := qtrt.Malloc(8) // QByteArray
	rv, err := qtrt.Qtcc3(2549705587, "_ZNK11QMetaMethod15methodSignatureEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QByteArrayFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:60
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QByteArray name() const

/*
 */
func (this *QMetaMethod) Name() *QByteArray /*123*/ {
	sretobj := qtrt.Malloc(8) // QByteArray
	rv, err := qtrt.Qtcc3(2756351335, "_ZNK11QMetaMethod4nameEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QByteArrayFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:61
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] const char * typeName() const

/*
 */
func (this *QMetaMethod) TypeName() string {
	rv, err := qtrt.Qtcc3(1716461499, "_ZNK11QMetaMethod8typeNameEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:62
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int returnType() const

/*
 */
func (this *QMetaMethod) ReturnType() int {
	rv, err := qtrt.Qtcc3(857633469, "_ZNK11QMetaMethod10returnTypeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:63
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int parameterCount() const

/*
 */
func (this *QMetaMethod) ParameterCount() int {
	rv, err := qtrt.Qtcc3(3953292363, "_ZNK11QMetaMethod14parameterCountEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:64
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int parameterType(int) const

/*
 */
func (this *QMetaMethod) ParameterType(index int) int {
	rv, err := qtrt.Qtcc3(1621940323, "_ZNK11QMetaMethod13parameterTypeEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:65
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void getParameterTypes(int *) const

/*
 */
func (this *QMetaMethod) GetParameterTypes(types unsafe.Pointer /*666*/) {
	rv, err := qtrt.Qtcc3(4219375380, "_ZNK11QMetaMethod17getParameterTypesEPi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&types))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:68
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] const char * tag() const

/*
 */
func (this *QMetaMethod) Tag() string {
	rv, err := qtrt.Qtcc3(2225931727, "_ZNK11QMetaMethod3tagEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:70
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] QMetaMethod::Access access() const

/*
 */
func (this *QMetaMethod) Access() int {
	rv, err := qtrt.Qtcc3(1453354162, "_ZNK11QMetaMethod6accessEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:72
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] QMetaMethod::MethodType methodType() const

/*
 */
func (this *QMetaMethod) MethodType() int {
	rv, err := qtrt.Qtcc3(1406394018, "_ZNK11QMetaMethod10methodTypeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:74
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int attributes() const

/*
 */
func (this *QMetaMethod) Attributes() int {
	rv, err := qtrt.Qtcc3(4009498167, "_ZNK11QMetaMethod10attributesEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:75
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int methodIndex() const

/*
 */
func (this *QMetaMethod) MethodIndex() int {
	rv, err := qtrt.Qtcc3(2608998004, "_ZNK11QMetaMethod11methodIndexEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:76
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int revision() const

/*
 */
func (this *QMetaMethod) Revision() int {
	rv, err := qtrt.Qtcc3(245548466, "_ZNK11QMetaMethod8revisionEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:78
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] const QMetaObject * enclosingMetaObject() const

/*
 */
func (this *QMetaMethod) EnclosingMetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.Qtcc3(2832786517, "_ZNK11QMetaMethod19enclosingMetaObjectEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QMetaObjectFromptr(Voidptr(uintptr(rv))) // 444
}

func DeleteQMetaMethod(this *QMetaMethod) {
	rv, err := qtrt.Qtcc3(3696685076, "_ZN11QMetaMethodD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QMetaMethod__Access = int

//
const QMetaMethod__Private QMetaMethod__Access = 0

//
const QMetaMethod__Protected QMetaMethod__Access = 1

//
const QMetaMethod__Public QMetaMethod__Access = 2

func (this *QMetaMethod) AccessItemName(val int) string {
	switch val {
	case QMetaMethod__Private: // 0
		return "Private"
	case QMetaMethod__Protected: // 1
		return "Protected"
	case QMetaMethod__Public: // 2
		return "Public"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QMetaMethod_AccessItemName(val int) string {
	var nilthis *QMetaMethod
	return nilthis.AccessItemName(val)
}

/*


 */
type QMetaMethod__MethodType = int

//
const QMetaMethod__Method QMetaMethod__MethodType = 0

//
const QMetaMethod__Signal QMetaMethod__MethodType = 1

//
const QMetaMethod__Slot QMetaMethod__MethodType = 2

//
const QMetaMethod__Constructor QMetaMethod__MethodType = 3

func (this *QMetaMethod) MethodTypeItemName(val int) string {
	switch val {
	case QMetaMethod__Method: // 0
		return "Method"
	case QMetaMethod__Signal: // 1
		return "Signal"
	case QMetaMethod__Slot: // 2
		return "Slot"
	case QMetaMethod__Constructor: // 3
		return "Constructor"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QMetaMethod_MethodTypeItemName(val int) string {
	var nilthis *QMetaMethod
	return nilthis.MethodTypeItemName(val)
}

/*


 */
type QMetaMethod__Attributes = int

//
const QMetaMethod__Compatibility QMetaMethod__Attributes = 1

//
const QMetaMethod__Cloned QMetaMethod__Attributes = 2

//
const QMetaMethod__Scriptable QMetaMethod__Attributes = 4

func (this *QMetaMethod) AttributesItemName(val int) string {
	switch val {
	case QMetaMethod__Compatibility: // 1
		return "Compatibility"
	case QMetaMethod__Cloned: // 2
		return "Cloned"
	case QMetaMethod__Scriptable: // 4
		return "Scriptable"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QMetaMethod_AttributesItemName(val int) string {
	var nilthis *QMetaMethod
	return nilthis.AttributesItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10041() {
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
