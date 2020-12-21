package qtcore

// /usr/include/qt/QtCore/qobjectdefs.h
// #include <qobjectdefs.h>
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
// size 48
type QMetaObject struct {
	*qtrt.CObject
}
type QMetaObject_ITF interface {
	QMetaObject_PTR() *QMetaObject
}

func (ptr *QMetaObject) QMetaObject_PTR() *QMetaObject { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QMetaObjectFromptr(cthis Voidptr) *QMetaObject {
	return &QMetaObject{&qtrt.CObject{cthis}}
}
func (*QMetaObject) Fromptr(cthis Voidptr) *QMetaObject {
	return QMetaObjectFromptr(cthis)
}

// /usr/include/qt/QtCore/qobjectdefs.h:340
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] const char * className() const

/*
 */
func (this *QMetaObject) ClassName() string {
	rv, err := qtrt.Qtcc3(2567088323, "_ZNK11QMetaObject9classNameEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qobjectdefs.h:341
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] const QMetaObject * superClass() const

/*
 */
func (this *QMetaObject) SuperClass() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.Qtcc3(4180358434, "_ZNK11QMetaObject10superClassEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QMetaObjectFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:344
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] QObject * cast(QObject *) const

/*
 */
func (this *QMetaObject) Cast(obj QObject_ITF /*777 QObject **/) *QObject /*777 QObject **/ {
	var convArg0 Voidptr
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(4086969823, "_ZNK11QMetaObject4castEP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QObjectFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:345
// index:1
// Public Direct Visibility=Default Availability=Available
// [8] const QObject * cast(const QObject *) const

/*
 */
func (this *QMetaObject) Cast1(obj QObject_ITF /*777 const QObject **/) *QObject /*777 const QObject **/ {
	var convArg0 Voidptr
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1982383341, "_ZNK11QMetaObject4castEPK7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return /*==*/ QObjectFromptr(Voidptr(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:351
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int methodOffset() const

/*
 */
func (this *QMetaObject) MethodOffset() int {
	rv, err := qtrt.Qtcc3(2344179760, "_ZNK11QMetaObject12methodOffsetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:352
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int enumeratorOffset() const

/*
 */
func (this *QMetaObject) EnumeratorOffset() int {
	rv, err := qtrt.Qtcc3(3130508702, "_ZNK11QMetaObject16enumeratorOffsetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:353
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int propertyOffset() const

/*
 */
func (this *QMetaObject) PropertyOffset() int {
	rv, err := qtrt.Qtcc3(2486189546, "_ZNK11QMetaObject14propertyOffsetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:354
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int classInfoOffset() const

/*
 */
func (this *QMetaObject) ClassInfoOffset() int {
	rv, err := qtrt.Qtcc3(3719363778, "_ZNK11QMetaObject15classInfoOffsetEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:356
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int constructorCount() const

/*
 */
func (this *QMetaObject) ConstructorCount() int {
	rv, err := qtrt.Qtcc3(722228823, "_ZNK11QMetaObject16constructorCountEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:357
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int methodCount() const

/*
 */
func (this *QMetaObject) MethodCount() int {
	rv, err := qtrt.Qtcc3(1964931486, "_ZNK11QMetaObject11methodCountEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:358
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int enumeratorCount() const

/*
 */
func (this *QMetaObject) EnumeratorCount() int {
	rv, err := qtrt.Qtcc3(2750070921, "_ZNK11QMetaObject15enumeratorCountEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:359
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int propertyCount() const

/*
 */
func (this *QMetaObject) PropertyCount() int {
	rv, err := qtrt.Qtcc3(2808090851, "_ZNK11QMetaObject13propertyCountEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:360
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int classInfoCount() const

/*
 */
func (this *QMetaObject) ClassInfoCount() int {
	rv, err := qtrt.Qtcc3(884557363, "_ZNK11QMetaObject14classInfoCountEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:362
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int indexOfConstructor(const char *) const

/*
 */
func (this *QMetaObject) IndexOfConstructor(constructor string) int {
	var convArg0 = qtrt.CStringRef(&constructor)
	rv, err := qtrt.Qtcc3(541226815, "_ZNK11QMetaObject18indexOfConstructorEPKc", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:363
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int indexOfMethod(const char *) const

/*
 */
func (this *QMetaObject) IndexOfMethod(method string) int {
	var convArg0 = qtrt.CStringRef(&method)
	rv, err := qtrt.Qtcc3(1248615640, "_ZNK11QMetaObject13indexOfMethodEPKc", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:364
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int indexOfSignal(const char *) const

/*
 */
func (this *QMetaObject) IndexOfSignal(signal string) int {
	var convArg0 = qtrt.CStringRef(&signal)
	rv, err := qtrt.Qtcc3(245446612, "_ZNK11QMetaObject13indexOfSignalEPKc", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:365
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int indexOfSlot(const char *) const

/*
 */
func (this *QMetaObject) IndexOfSlot(slot string) int {
	var convArg0 = qtrt.CStringRef(&slot)
	rv, err := qtrt.Qtcc3(2159463968, "_ZNK11QMetaObject11indexOfSlotEPKc", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:366
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int indexOfEnumerator(const char *) const

/*
 */
func (this *QMetaObject) IndexOfEnumerator(name string) int {
	var convArg0 = qtrt.CStringRef(&name)
	rv, err := qtrt.Qtcc3(3528180915, "_ZNK11QMetaObject17indexOfEnumeratorEPKc", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:367
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int indexOfProperty(const char *) const

/*
 */
func (this *QMetaObject) IndexOfProperty(name string) int {
	var convArg0 = qtrt.CStringRef(&name)
	rv, err := qtrt.Qtcc3(2753009205, "_ZNK11QMetaObject15indexOfPropertyEPKc", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:368
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int indexOfClassInfo(const char *) const

/*
 */
func (this *QMetaObject) IndexOfClassInfo(name string) int {
	var convArg0 = qtrt.CStringRef(&name)
	rv, err := qtrt.Qtcc3(3923681255, "_ZNK11QMetaObject16indexOfClassInfoEPKc", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:370
// index:0
// Public Direct Visibility=Default Availability=Available
// [16] QMetaMethod constructor(int) const

/*
 */
func (this *QMetaObject) Constructor(index int) *QMetaMethod /*123*/ {
	rv, err := qtrt.Qtcc3(3534876830, "_ZNK11QMetaObject11constructorEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
	rv2 := /*==*/ QMetaMethodFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaMethod)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:371
// index:0
// Public Direct Visibility=Default Availability=Available
// [16] QMetaMethod method(int) const

/*
 */
func (this *QMetaObject) Method(index int) *QMetaMethod /*123*/ {
	rv, err := qtrt.Qtcc3(981682568, "_ZNK11QMetaObject6methodEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
	rv2 := /*==*/ QMetaMethodFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaMethod)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:372
// index:0
// Public Direct Visibility=Default Availability=Available
// [16] QMetaEnum enumerator(int) const

/*
 */
func (this *QMetaObject) Enumerator(index int) *QMetaEnum /*123*/ {
	rv, err := qtrt.Qtcc3(819545584, "_ZNK11QMetaObject10enumeratorEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
	rv2 := /*==*/ QMetaEnumFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaEnum)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:373
// index:0
// Public Indirect Visibility=Default Availability=Available
// [32] QMetaProperty property(int) const

/*
 */
func (this *QMetaObject) Property(index int) *QMetaProperty /*123*/ {
	sretobj := qtrt.Malloc(32) // QMetaProperty
	rv, err := qtrt.Qtcc3(588131701, "_ZNK11QMetaObject8propertyEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QMetaPropertyFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaProperty)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:374
// index:0
// Public Direct Visibility=Default Availability=Available
// [16] QMetaClassInfo classInfo(int) const

/*
 */
func (this *QMetaObject) ClassInfo(index int) *QMetaClassInfo /*123*/ {
	rv, err := qtrt.Qtcc3(1372127368, "_ZNK11QMetaObject9classInfoEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&index))
	qtrt.ErrPrint2(err, rv)
	rv2 := /*==*/ QMetaClassInfoFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaClassInfo)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:375
// index:0
// Public Indirect Visibility=Default Availability=Available
// [32] QMetaProperty userProperty() const

/*
 */
func (this *QMetaObject) UserProperty() *QMetaProperty /*123*/ {
	sretobj := qtrt.Malloc(32) // QMetaProperty
	rv, err := qtrt.Qtcc3(4052217834, "_ZNK11QMetaObject12userPropertyEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QMetaPropertyFromptr(Voidptr(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaProperty)
	return rv2
}

func DeleteQMetaObject(this *QMetaObject) {
	rv, err := qtrt.Qtcc1(0, "_ZN11QMetaObjectD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QMetaObject__Call = int

//
const QMetaObject__InvokeMetaMethod QMetaObject__Call = 0

//
const QMetaObject__ReadProperty QMetaObject__Call = 1

//
const QMetaObject__WriteProperty QMetaObject__Call = 2

//
const QMetaObject__ResetProperty QMetaObject__Call = 3

//
const QMetaObject__QueryPropertyDesignable QMetaObject__Call = 4

//
const QMetaObject__QueryPropertyScriptable QMetaObject__Call = 5

//
const QMetaObject__QueryPropertyStored QMetaObject__Call = 6

//
const QMetaObject__QueryPropertyEditable QMetaObject__Call = 7

//
const QMetaObject__QueryPropertyUser QMetaObject__Call = 8

//
const QMetaObject__CreateInstance QMetaObject__Call = 9

//
const QMetaObject__IndexOfMethod QMetaObject__Call = 10

//
const QMetaObject__RegisterPropertyMetaType QMetaObject__Call = 11

//
const QMetaObject__RegisterMethodArgumentMetaType QMetaObject__Call = 12

func (this *QMetaObject) CallItemName(val int) string {
	switch val {
	case QMetaObject__InvokeMetaMethod: // 0
		return "InvokeMetaMethod"
	case QMetaObject__ReadProperty: // 1
		return "ReadProperty"
	case QMetaObject__WriteProperty: // 2
		return "WriteProperty"
	case QMetaObject__ResetProperty: // 3
		return "ResetProperty"
	case QMetaObject__QueryPropertyDesignable: // 4
		return "QueryPropertyDesignable"
	case QMetaObject__QueryPropertyScriptable: // 5
		return "QueryPropertyScriptable"
	case QMetaObject__QueryPropertyStored: // 6
		return "QueryPropertyStored"
	case QMetaObject__QueryPropertyEditable: // 7
		return "QueryPropertyEditable"
	case QMetaObject__QueryPropertyUser: // 8
		return "QueryPropertyUser"
	case QMetaObject__CreateInstance: // 9
		return "CreateInstance"
	case QMetaObject__IndexOfMethod: // 10
		return "IndexOfMethod"
	case QMetaObject__RegisterPropertyMetaType: // 11
		return "RegisterPropertyMetaType"
	case QMetaObject__RegisterMethodArgumentMetaType: // 12
		return "RegisterMethodArgumentMetaType"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QMetaObject_CallItemName(val int) string {
	var nilthis *QMetaObject
	return nilthis.CallItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10005() {
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
