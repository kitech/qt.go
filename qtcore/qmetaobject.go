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
// extern C begin: 1
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
type QMetaObject struct {
	*qtrt.CObject
}
type QMetaObject_ITF interface {
	QMetaObject_PTR() *QMetaObject
}

func (ptr *QMetaObject) QMetaObject_PTR() *QMetaObject { return ptr }

func (this *QMetaObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMetaObject) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMetaObjectFromPointer(cthis unsafe.Pointer) *QMetaObject {
	return &QMetaObject{&qtrt.CObject{cthis}}
}
func (*QMetaObject) NewFromPointer(cthis unsafe.Pointer) *QMetaObject {
	return NewQMetaObjectFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobjectdefs.h:345
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * className() const

/*

 */
func (this *QMetaObject) ClassName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject9classNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qobjectdefs.h:346
// index:0
// Public Visibility=Default Availability=Available
// [8] const QMetaObject * superClass() const

/*

 */
func (this *QMetaObject) SuperClass() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject10superClassEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:349
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * cast(QObject *) const

/*

 */
func (this *QMetaObject) Cast(obj QObject_ITF /*777 QObject **/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject4castEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:350
// index:1
// Public Visibility=Default Availability=Available
// [8] const QObject * cast(const QObject *) const

/*

 */
func (this *QMetaObject) Cast_1(obj QObject_ITF /*777 const QObject **/) *QObject /*777 const QObject **/ {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject4castEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:356
// index:0
// Public Visibility=Default Availability=Available
// [4] int methodOffset() const

/*

 */
func (this *QMetaObject) MethodOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject12methodOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:357
// index:0
// Public Visibility=Default Availability=Available
// [4] int enumeratorOffset() const

/*

 */
func (this *QMetaObject) EnumeratorOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject16enumeratorOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:358
// index:0
// Public Visibility=Default Availability=Available
// [4] int propertyOffset() const

/*

 */
func (this *QMetaObject) PropertyOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject14propertyOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:359
// index:0
// Public Visibility=Default Availability=Available
// [4] int classInfoOffset() const

/*

 */
func (this *QMetaObject) ClassInfoOffset() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject15classInfoOffsetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:361
// index:0
// Public Visibility=Default Availability=Available
// [4] int constructorCount() const

/*

 */
func (this *QMetaObject) ConstructorCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject16constructorCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:362
// index:0
// Public Visibility=Default Availability=Available
// [4] int methodCount() const

/*

 */
func (this *QMetaObject) MethodCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11methodCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:363
// index:0
// Public Visibility=Default Availability=Available
// [4] int enumeratorCount() const

/*

 */
func (this *QMetaObject) EnumeratorCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject15enumeratorCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:364
// index:0
// Public Visibility=Default Availability=Available
// [4] int propertyCount() const

/*

 */
func (this *QMetaObject) PropertyCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject13propertyCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:365
// index:0
// Public Visibility=Default Availability=Available
// [4] int classInfoCount() const

/*

 */
func (this *QMetaObject) ClassInfoCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject14classInfoCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:367
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOfConstructor(const char *) const

/*

 */
func (this *QMetaObject) IndexOfConstructor(constructor string) int {
	var convArg0 = qtrt.CString(constructor)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject18indexOfConstructorEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:368
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOfMethod(const char *) const

/*

 */
func (this *QMetaObject) IndexOfMethod(method string) int {
	var convArg0 = qtrt.CString(method)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject13indexOfMethodEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:369
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOfSignal(const char *) const

/*

 */
func (this *QMetaObject) IndexOfSignal(signal string) int {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject13indexOfSignalEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:370
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOfSlot(const char *) const

/*

 */
func (this *QMetaObject) IndexOfSlot(slot string) int {
	var convArg0 = qtrt.CString(slot)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11indexOfSlotEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:371
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOfEnumerator(const char *) const

/*

 */
func (this *QMetaObject) IndexOfEnumerator(name string) int {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject17indexOfEnumeratorEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:372
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOfProperty(const char *) const

/*

 */
func (this *QMetaObject) IndexOfProperty(name string) int {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject15indexOfPropertyEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:373
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOfClassInfo(const char *) const

/*

 */
func (this *QMetaObject) IndexOfClassInfo(name string) int {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject16indexOfClassInfoEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:375
// index:0
// Public Visibility=Default Availability=Available
// [16] QMetaMethod constructor(int) const

/*

 */
func (this *QMetaObject) Constructor(index int) *QMetaMethod /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11constructorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMetaMethodFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaMethod)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:376
// index:0
// Public Visibility=Default Availability=Available
// [16] QMetaMethod method(int) const

/*

 */
func (this *QMetaObject) Method(index int) *QMetaMethod /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject6methodEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMetaMethodFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaMethod)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:377
// index:0
// Public Visibility=Default Availability=Available
// [16] QMetaEnum enumerator(int) const

/*

 */
func (this *QMetaObject) Enumerator(index int) *QMetaEnum /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject10enumeratorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMetaEnumFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaEnum)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:378
// index:0
// Public Visibility=Default Availability=Available
// [32] QMetaProperty property(int) const

/*

 */
func (this *QMetaObject) Property(index int) *QMetaProperty /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject8propertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMetaPropertyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaProperty)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:379
// index:0
// Public Visibility=Default Availability=Available
// [16] QMetaClassInfo classInfo(int) const

/*

 */
func (this *QMetaObject) ClassInfo(index int) *QMetaClassInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject9classInfoEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMetaClassInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaClassInfo)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:380
// index:0
// Public Visibility=Default Availability=Available
// [32] QMetaProperty userProperty() const

/*

 */
func (this *QMetaObject) UserProperty() *QMetaProperty /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject12userPropertyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMetaPropertyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaProperty)
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:382
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool checkConnectArgs(const char *, const char *)

/*

 */
func (this *QMetaObject) CheckConnectArgs(signal string, method string) bool {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(method)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject16checkConnectArgsEPKcS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QMetaObject_CheckConnectArgs(signal string, method string) bool {
	var nilthis *QMetaObject
	rv := nilthis.CheckConnectArgs(signal, method)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:383
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool checkConnectArgs(const QMetaMethod &, const QMetaMethod &)

/*

 */
func (this *QMetaObject) CheckConnectArgs_1(signal QMetaMethod_ITF, method QMetaMethod_ITF) bool {
	var convArg0 unsafe.Pointer
	if signal != nil && signal.QMetaMethod_PTR() != nil {
		convArg0 = signal.QMetaMethod_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if method != nil && method.QMetaMethod_PTR() != nil {
		convArg1 = method.QMetaMethod_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject16checkConnectArgsERK11QMetaMethodS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QMetaObject_CheckConnectArgs_1(signal QMetaMethod_ITF, method QMetaMethod_ITF) bool {
	var nilthis *QMetaObject
	rv := nilthis.CheckConnectArgs_1(signal, method)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:385
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray normalizedSignature(const char *)

/*

 */
func (this *QMetaObject) NormalizedSignature(method string) *QByteArray /*123*/ {
	var convArg0 = qtrt.CString(method)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject19normalizedSignatureEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QMetaObject_NormalizedSignature(method string) *QByteArray /*123*/ {
	var nilthis *QMetaObject
	rv := nilthis.NormalizedSignature(method)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:386
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray normalizedType(const char *)

/*

 */
func (this *QMetaObject) NormalizedType(type_ string) *QByteArray /*123*/ {
	var convArg0 = qtrt.CString(type_)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject14normalizedTypeEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QMetaObject_NormalizedType(type_ string) *QByteArray /*123*/ {
	var nilthis *QMetaObject
	rv := nilthis.NormalizedType(type_)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:389
// index:0
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, int, const QObject *, int, int, int *)

/*

 */
func (this *QMetaObject) Connect(sender QObject_ITF /*777 const QObject **/, signal_index int, receiver QObject_ITF /*777 const QObject **/, method_index int, type_ int, types unsafe.Pointer /*666*/) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject7connectEPK7QObjectiS2_iiPi", qtrt.FFI_TYPE_POINTER, convArg0, signal_index, convArg2, method_index, type_, types)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QMetaObject_Connect(sender QObject_ITF /*777 const QObject **/, signal_index int, receiver QObject_ITF /*777 const QObject **/, method_index int, type_ int, types unsafe.Pointer /*666*/) unsafe.Pointer /*444*/ {
	var nilthis *QMetaObject
	rv := nilthis.Connect(sender, signal_index, receiver, method_index, type_, types)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:389
// index:0
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, int, const QObject *, int, int, int *)

/*

 */
func (this *QMetaObject) Connect__(sender QObject_ITF /*777 const QObject **/, signal_index int, receiver QObject_ITF /*777 const QObject **/, method_index int) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	// arg: 4, int=Int, =Invalid, , Invalid
	type_ := int(0)
	// arg: 5, int *=Pointer, =Invalid, , Invalid
	var types unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject7connectEPK7QObjectiS2_iiPi", qtrt.FFI_TYPE_POINTER, convArg0, signal_index, convArg2, method_index, type_, types)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qobjectdefs.h:389
// index:0
// Public static Visibility=Default Availability=Available
// [8] QMetaObject::Connection connect(const QObject *, int, const QObject *, int, int, int *)

/*

 */
func (this *QMetaObject) Connect__1(sender QObject_ITF /*777 const QObject **/, signal_index int, receiver QObject_ITF /*777 const QObject **/, method_index int, type_ int) unsafe.Pointer /*444*/ {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	// arg: 5, int *=Pointer, =Invalid, , Invalid
	var types unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject7connectEPK7QObjectiS2_iiPi", qtrt.FFI_TYPE_POINTER, convArg0, signal_index, convArg2, method_index, type_, types)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qobjectdefs.h:393
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool disconnect(const QObject *, int, const QObject *, int)

/*

 */
func (this *QMetaObject) Disconnect(sender QObject_ITF /*777 const QObject **/, signal_index int, receiver QObject_ITF /*777 const QObject **/, method_index int) bool {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject10disconnectEPK7QObjectiS2_i", qtrt.FFI_TYPE_POINTER, convArg0, signal_index, convArg2, method_index)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QMetaObject_Disconnect(sender QObject_ITF /*777 const QObject **/, signal_index int, receiver QObject_ITF /*777 const QObject **/, method_index int) bool {
	var nilthis *QMetaObject
	rv := nilthis.Disconnect(sender, signal_index, receiver, method_index)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:395
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool disconnectOne(const QObject *, int, const QObject *, int)

/*

 */
func (this *QMetaObject) DisconnectOne(sender QObject_ITF /*777 const QObject **/, signal_index int, receiver QObject_ITF /*777 const QObject **/, method_index int) bool {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg2 = receiver.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject13disconnectOneEPK7QObjectiS2_i", qtrt.FFI_TYPE_POINTER, convArg0, signal_index, convArg2, method_index)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QMetaObject_DisconnectOne(sender QObject_ITF /*777 const QObject **/, signal_index int, receiver QObject_ITF /*777 const QObject **/, method_index int) bool {
	var nilthis *QMetaObject
	rv := nilthis.DisconnectOne(sender, signal_index, receiver, method_index)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:398
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void connectSlotsByName(QObject *)

/*

 */
func (this *QMetaObject) ConnectSlotsByName(o QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if o != nil && o.QObject_PTR() != nil {
		convArg0 = o.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject18connectSlotsByNameEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QMetaObject_ConnectSlotsByName(o QObject_ITF /*777 QObject **/) {
	var nilthis *QMetaObject
	nilthis.ConnectSlotsByName(o)
}

// /usr/include/qt/QtCore/qobjectdefs.h:401
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void activate(QObject *, int, void **)

/*

 */
func (this *QMetaObject) Activate(sender QObject_ITF /*777 QObject **/, signal_index int, argv unsafe.Pointer /*666*/) {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject8activateEP7QObjectiPPv", qtrt.FFI_TYPE_POINTER, convArg0, signal_index, argv)
	qtrt.ErrPrint(err, rv)
}
func QMetaObject_Activate(sender QObject_ITF /*777 QObject **/, signal_index int, argv unsafe.Pointer /*666*/) {
	var nilthis *QMetaObject
	nilthis.Activate(sender, signal_index, argv)
}

// /usr/include/qt/QtCore/qobjectdefs.h:403
// index:1
// Public static Visibility=Default Availability=Available
// [-2] void activate(QObject *, int, int, void **)

/*

 */
func (this *QMetaObject) Activate_1(sender QObject_ITF /*777 QObject **/, signal_offset int, local_signal_index int, argv unsafe.Pointer /*666*/) {
	var convArg0 unsafe.Pointer
	if sender != nil && sender.QObject_PTR() != nil {
		convArg0 = sender.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject8activateEP7QObjectiiPPv", qtrt.FFI_TYPE_POINTER, convArg0, signal_offset, local_signal_index, argv)
	qtrt.ErrPrint(err, rv)
}
func QMetaObject_Activate_1(sender QObject_ITF /*777 QObject **/, signal_offset int, local_signal_index int, argv unsafe.Pointer /*666*/) {
	var nilthis *QMetaObject
	nilthis.Activate_1(sender, signal_offset, local_signal_index, argv)
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg4 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg5 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg6 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg7 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg8 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg9 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg10 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg11 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg11 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg12 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg12 = val8.QGenericArgument_PTR().GetCthis()
	}
	var convArg13 unsafe.Pointer
	if val9 != nil && val9.QGenericArgument_PTR() != nil {
		convArg13 = val9.QGenericArgument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QMetaObject_InvokeMethod(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var nilthis *QMetaObject
	rv := nilthis.InvokeMethod(obj, member, arg2, ret, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod__(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	// arg: 13, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg13 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod__1(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg4 = val0.QGenericArgument_PTR().GetCthis()
	}
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	// arg: 13, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg13 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod__2(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg4 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg5 = val1.QGenericArgument_PTR().GetCthis()
	}
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	// arg: 13, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg13 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod__3(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg4 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg5 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg6 = val2.QGenericArgument_PTR().GetCthis()
	}
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	// arg: 13, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg13 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod__4(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg4 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg5 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg6 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg7 = val3.QGenericArgument_PTR().GetCthis()
	}
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	// arg: 13, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg13 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod__5(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg4 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg5 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg6 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg7 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg8 = val4.QGenericArgument_PTR().GetCthis()
	}
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	// arg: 13, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg13 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod__6(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg4 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg5 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg6 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg7 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg8 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg9 = val5.QGenericArgument_PTR().GetCthis()
	}
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	// arg: 13, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg13 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod__7(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg4 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg5 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg6 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg7 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg8 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg9 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg10 = val6.QGenericArgument_PTR().GetCthis()
	}
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	// arg: 13, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg13 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod__8(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg4 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg5 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg6 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg7 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg8 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg9 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg10 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg11 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg11 = val7.QGenericArgument_PTR().GetCthis()
	}
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	// arg: 13, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg13 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod__9(obj QObject_ITF /*777 QObject **/, member string, arg2 int, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg3 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg4 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg5 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg6 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg7 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg8 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg9 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg10 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg11 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg11 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg12 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg12 = val8.QGenericArgument_PTR().GetCthis()
	}
	// arg: 13, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg13 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, arg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12, convArg13)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg8 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg9 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg10 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg11 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg11 = val8.QGenericArgument_PTR().GetCthis()
	}
	var convArg12 unsafe.Pointer
	if val9 != nil && val9.QGenericArgument_PTR() != nil {
		convArg12 = val9.QGenericArgument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QMetaObject_InvokeMethod_1(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var nilthis *QMetaObject
	rv := nilthis.InvokeMethod_1(obj, member, ret, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1_(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	// arg: 3, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg3 unsafe.Pointer
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1_1(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1_2(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1_3(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1_4(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1_5(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1_6(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg8 = val5.QGenericArgument_PTR().GetCthis()
	}
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1_7(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg8 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg9 = val6.QGenericArgument_PTR().GetCthis()
	}
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1_8(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg8 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg9 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg10 = val7.QGenericArgument_PTR().GetCthis()
	}
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_1_9(obj QObject_ITF /*777 QObject **/, member string, ret QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if ret != nil && ret.QGenericReturnArgument_PTR() != nil {
		convArg2 = ret.QGenericReturnArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg8 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg9 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg10 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg11 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg11 = val8.QGenericArgument_PTR().GetCthis()
	}
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg8 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg9 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg10 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg11 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg11 = val8.QGenericArgument_PTR().GetCthis()
	}
	var convArg12 unsafe.Pointer
	if val9 != nil && val9.QGenericArgument_PTR() != nil {
		convArg12 = val9.QGenericArgument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QMetaObject_InvokeMethod_2(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var nilthis *QMetaObject
	rv := nilthis.InvokeMethod_2(obj, member, type_, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2_(obj QObject_ITF /*777 QObject **/, member string, type_ int) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	// arg: 3, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg3 unsafe.Pointer
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2_1(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2_2(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2_3(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2_4(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2_5(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2_6(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg8 = val5.QGenericArgument_PTR().GetCthis()
	}
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2_7(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg8 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg9 = val6.QGenericArgument_PTR().GetCthis()
	}
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2_8(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg8 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg9 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg10 = val7.QGenericArgument_PTR().GetCthis()
	}
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_2_9(obj QObject_ITF /*777 QObject **/, member string, type_ int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg3 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg3 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg4 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg5 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg6 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg7 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg8 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg9 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg10 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg11 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg11 = val8.QGenericArgument_PTR().GetCthis()
	}
	// arg: 12, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg12 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, type_, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg2 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg3 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg4 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg5 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg6 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg7 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg8 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg9 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg10 = val8.QGenericArgument_PTR().GetCthis()
	}
	var convArg11 unsafe.Pointer
	if val9 != nil && val9.QGenericArgument_PTR() != nil {
		convArg11 = val9.QGenericArgument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QMetaObject_InvokeMethod_3(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var nilthis *QMetaObject
	rv := nilthis.InvokeMethod_3(obj, member, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3_(obj QObject_ITF /*777 QObject **/, member string) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg3 unsafe.Pointer
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3_1(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg2 = val0.QGenericArgument_PTR().GetCthis()
	}
	// arg: 3, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg3 unsafe.Pointer
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3_2(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg2 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg3 = val1.QGenericArgument_PTR().GetCthis()
	}
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3_3(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg2 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg3 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg4 = val2.QGenericArgument_PTR().GetCthis()
	}
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3_4(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg2 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg3 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg4 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg5 = val3.QGenericArgument_PTR().GetCthis()
	}
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3_5(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg2 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg3 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg4 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg5 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg6 = val4.QGenericArgument_PTR().GetCthis()
	}
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3_6(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg2 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg3 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg4 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg5 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg6 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg7 = val5.QGenericArgument_PTR().GetCthis()
	}
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3_7(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg2 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg3 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg4 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg5 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg6 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg7 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg8 = val6.QGenericArgument_PTR().GetCthis()
	}
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3_8(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg2 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg3 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg4 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg5 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg6 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg7 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg8 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg9 = val7.QGenericArgument_PTR().GetCthis()
	}
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline Visibility=Default Availability=Available
// [1] bool invokeMethod(QObject *, const char *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument)

/*

 */
func (this *QMetaObject) InvokeMethod_3_9(obj QObject_ITF /*777 QObject **/, member string, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(member)
	defer qtrt.FreeMem(convArg1)
	var convArg2 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg2 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg3 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg4 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg5 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg6 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg7 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg8 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg9 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg10 = val8.QGenericArgument_PTR().GetCthis()
	}
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance(val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg0 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg1 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg2 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg3 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg4 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg5 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg6 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg7 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg8 = val8.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val9 != nil && val9.QGenericArgument_PTR() != nil {
		convArg9 = val9.QGenericArgument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance__() *QObject /*777 QObject **/ {
	// arg: 0, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg3 unsafe.Pointer
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance__1(val0 QGenericArgument_ITF /*123*/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg0 = val0.QGenericArgument_PTR().GetCthis()
	}
	// arg: 1, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg1 unsafe.Pointer
	// arg: 2, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg3 unsafe.Pointer
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance__2(val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg0 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg1 = val1.QGenericArgument_PTR().GetCthis()
	}
	// arg: 2, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg3 unsafe.Pointer
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance__3(val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg0 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg1 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg2 = val2.QGenericArgument_PTR().GetCthis()
	}
	// arg: 3, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg3 unsafe.Pointer
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance__4(val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg0 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg1 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg2 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg3 = val3.QGenericArgument_PTR().GetCthis()
	}
	// arg: 4, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg4 unsafe.Pointer
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance__5(val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg0 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg1 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg2 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg3 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg4 = val4.QGenericArgument_PTR().GetCthis()
	}
	// arg: 5, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg5 unsafe.Pointer
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance__6(val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg0 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg1 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg2 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg3 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg4 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg5 = val5.QGenericArgument_PTR().GetCthis()
	}
	// arg: 6, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg6 unsafe.Pointer
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance__7(val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg0 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg1 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg2 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg3 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg4 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg5 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg6 = val6.QGenericArgument_PTR().GetCthis()
	}
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance__8(val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg0 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg1 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg2 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg3 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg4 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg5 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg6 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg7 = val7.QGenericArgument_PTR().GetCthis()
	}
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * newInstance(QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaObject) NewInstance__9(val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg0 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg1 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg2 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg3 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg4 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg5 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg6 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg7 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg8 = val8.QGenericArgument_PTR().GetCthis()
	}
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobjectdefs.h:581
// index:0
// Public Visibility=Default Availability=Available
// [4] int static_metacall(QMetaObject::Call, int, void **) const

/*

 */
func (this *QMetaObject) Static_metacall(arg0 int, arg1 int, arg2 unsafe.Pointer /*666*/) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaObject15static_metacallENS_4CallEiPPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1, arg2)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qobjectdefs.h:582
// index:0
// Public static Visibility=Default Availability=Available
// [4] int metacall(QObject *, QMetaObject::Call, int, void **)

/*

 */
func (this *QMetaObject) Metacall(arg0 QObject_ITF /*777 QObject **/, arg1 int, arg2 int, arg3 unsafe.Pointer /*666*/) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObject8metacallEP7QObjectNS_4CallEiPPv", qtrt.FFI_TYPE_POINTER, convArg0, arg1, arg2, arg3)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QMetaObject_Metacall(arg0 QObject_ITF /*777 QObject **/, arg1 int, arg2 int, arg3 unsafe.Pointer /*666*/) int {
	var nilthis *QMetaObject
	rv := nilthis.Metacall(arg0, arg1, arg2, arg3)
	return rv
}

func DeleteQMetaObject(this *QMetaObject) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaObjectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
