package qtcore

// /usr/include/qt/QtCore/qobjectdefs.h
// #include <qobjectdefs.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
type QMetaObject struct {
	*qtrt.CObject
}

func (this *QMetaObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMetaObject) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQMetaObjectFromPointer(cthis unsafe.Pointer) *QMetaObject {
	return &QMetaObject{&qtrt.CObject{cthis}}
}
func (*QMetaObject) NewFromPointer(cthis unsafe.Pointer) *QMetaObject {
	return NewQMetaObjectFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobjectdefs.h:345
// index:0
// Public
// const char * className()
func (this *QMetaObject) ClassName() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject9classNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qobjectdefs.h:346
// index:0
// Public
// const QMetaObject * superClass()
func (this *QMetaObject) SuperClass() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject10superClassEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:349
// index:0
// Public
// QObject * cast(class QObject *)
func (this *QMetaObject) Cast(obj *QObject /*777 QObject **/) *QObject /*777 QObject **/ {
	var convArg0 = obj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject4castEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:350
// index:1
// Public
// const QObject * cast(const class QObject *)
func (this *QMetaObject) Cast_1(obj *QObject /*777 const QObject **/) *QObject /*777 const QObject **/ {
	var convArg0 = obj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject4castEPK7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:356
// index:0
// Public
// int methodOffset()
func (this *QMetaObject) MethodOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject12methodOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:357
// index:0
// Public
// int enumeratorOffset()
func (this *QMetaObject) EnumeratorOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject16enumeratorOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:358
// index:0
// Public
// int propertyOffset()
func (this *QMetaObject) PropertyOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject14propertyOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:359
// index:0
// Public
// int classInfoOffset()
func (this *QMetaObject) ClassInfoOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject15classInfoOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:361
// index:0
// Public
// int constructorCount()
func (this *QMetaObject) ConstructorCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject16constructorCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:362
// index:0
// Public
// int methodCount()
func (this *QMetaObject) MethodCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject11methodCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:363
// index:0
// Public
// int enumeratorCount()
func (this *QMetaObject) EnumeratorCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject15enumeratorCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:364
// index:0
// Public
// int propertyCount()
func (this *QMetaObject) PropertyCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject13propertyCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:365
// index:0
// Public
// int classInfoCount()
func (this *QMetaObject) ClassInfoCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject14classInfoCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:367
// index:0
// Public
// int indexOfConstructor(const char *)
func (this *QMetaObject) IndexOfConstructor(constructor string) int {
	var convArg0 = qtrt.CString(constructor)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject18indexOfConstructorEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:368
// index:0
// Public
// int indexOfMethod(const char *)
func (this *QMetaObject) IndexOfMethod(method string) int {
	var convArg0 = qtrt.CString(method)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject13indexOfMethodEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:369
// index:0
// Public
// int indexOfSignal(const char *)
func (this *QMetaObject) IndexOfSignal(signal string) int {
	var convArg0 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject13indexOfSignalEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:370
// index:0
// Public
// int indexOfSlot(const char *)
func (this *QMetaObject) IndexOfSlot(slot string) int {
	var convArg0 = qtrt.CString(slot)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject11indexOfSlotEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:371
// index:0
// Public
// int indexOfEnumerator(const char *)
func (this *QMetaObject) IndexOfEnumerator(name string) int {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject17indexOfEnumeratorEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:372
// index:0
// Public
// int indexOfProperty(const char *)
func (this *QMetaObject) IndexOfProperty(name string) int {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject15indexOfPropertyEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:373
// index:0
// Public
// int indexOfClassInfo(const char *)
func (this *QMetaObject) IndexOfClassInfo(name string) int {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject16indexOfClassInfoEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:375
// index:0
// Public
// QMetaMethod constructor(int)
func (this *QMetaObject) Constructor(index int) *QMetaMethod /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject11constructorEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMetaMethodFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:376
// index:0
// Public
// QMetaMethod method(int)
func (this *QMetaObject) Method(index int) *QMetaMethod /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject6methodEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMetaMethodFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:377
// index:0
// Public
// QMetaEnum enumerator(int)
func (this *QMetaObject) Enumerator(index int) *QMetaEnum /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject10enumeratorEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMetaEnumFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:378
// index:0
// Public
// QMetaProperty property(int)
func (this *QMetaObject) Property(index int) *QMetaProperty /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject8propertyEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMetaPropertyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:379
// index:0
// Public
// QMetaClassInfo classInfo(int)
func (this *QMetaObject) ClassInfo(index int) *QMetaClassInfo /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject9classInfoEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMetaClassInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:380
// index:0
// Public
// QMetaProperty userProperty()
func (this *QMetaObject) UserProperty() *QMetaProperty /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject12userPropertyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQMetaPropertyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:382
// index:0
// Public static
// bool checkConnectArgs(const char *, const char *)
func (this *QMetaObject) CheckConnectArgs(signal string, method string) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject16checkConnectArgsEPKcS1_", ffiqt.FFI_TYPE_POINTER, signal, method)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QMetaObject_CheckConnectArgs(signal string, method string) bool {
	var nilthis *QMetaObject
	rv := nilthis.CheckConnectArgs(signal, method)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:383
// index:1
// Public static
// bool checkConnectArgs(const class QMetaMethod &, const class QMetaMethod &)
func (this *QMetaObject) CheckConnectArgs_1(signal *QMetaMethod, method *QMetaMethod) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject16checkConnectArgsERK11QMetaMethodS2_", ffiqt.FFI_TYPE_POINTER, signal, method)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QMetaObject_CheckConnectArgs_1(signal *QMetaMethod, method *QMetaMethod) bool {
	var nilthis *QMetaObject
	rv := nilthis.CheckConnectArgs_1(signal, method)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:385
// index:0
// Public static
// QByteArray normalizedSignature(const char *)
func (this *QMetaObject) NormalizedSignature(method string) *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject19normalizedSignatureEPKc", ffiqt.FFI_TYPE_POINTER, method)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QMetaObject_NormalizedSignature(method string) *QByteArray /*123*/ {
	var nilthis *QMetaObject
	rv := nilthis.NormalizedSignature(method)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:386
// index:0
// Public static
// QByteArray normalizedType(const char *)
func (this *QMetaObject) NormalizedType(type_ string) *QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject14normalizedTypeEPKc", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QMetaObject_NormalizedType(type_ string) *QByteArray /*123*/ {
	var nilthis *QMetaObject
	rv := nilthis.NormalizedType(type_)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:389
// index:0
// Public static
// QMetaObject::Connection connect(const class QObject *, int, const class QObject *, int, int, int *)
func (this *QMetaObject) Connect(sender *QObject /*777 const QObject **/, signal_index int, receiver *QObject /*777 const QObject **/, method_index int, type_ int, types unsafe.Pointer /*666*/) unsafe.Pointer /*444*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject7connectEPK7QObjectiS2_iiPi", ffiqt.FFI_TYPE_POINTER, sender, signal_index, receiver, method_index, type_, types)
	gopp.ErrPrint(err, rv)
	// return rv
	return unsafe.Pointer(uintptr(rv))
}
func QMetaObject_Connect(sender *QObject /*777 const QObject **/, signal_index int, receiver *QObject /*777 const QObject **/, method_index int, type_ int, types unsafe.Pointer /*666*/) unsafe.Pointer /*444*/ {
	var nilthis *QMetaObject
	rv := nilthis.Connect(sender, signal_index, receiver, method_index, type_, types)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:393
// index:0
// Public static
// bool disconnect(const class QObject *, int, const class QObject *, int)
func (this *QMetaObject) Disconnect(sender *QObject /*777 const QObject **/, signal_index int, receiver *QObject /*777 const QObject **/, method_index int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject10disconnectEPK7QObjectiS2_i", ffiqt.FFI_TYPE_POINTER, sender, signal_index, receiver, method_index)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QMetaObject_Disconnect(sender *QObject /*777 const QObject **/, signal_index int, receiver *QObject /*777 const QObject **/, method_index int) bool {
	var nilthis *QMetaObject
	rv := nilthis.Disconnect(sender, signal_index, receiver, method_index)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:395
// index:0
// Public static
// bool disconnectOne(const class QObject *, int, const class QObject *, int)
func (this *QMetaObject) DisconnectOne(sender *QObject /*777 const QObject **/, signal_index int, receiver *QObject /*777 const QObject **/, method_index int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject13disconnectOneEPK7QObjectiS2_i", ffiqt.FFI_TYPE_POINTER, sender, signal_index, receiver, method_index)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QMetaObject_DisconnectOne(sender *QObject /*777 const QObject **/, signal_index int, receiver *QObject /*777 const QObject **/, method_index int) bool {
	var nilthis *QMetaObject
	rv := nilthis.DisconnectOne(sender, signal_index, receiver, method_index)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:398
// index:0
// Public static
// void connectSlotsByName(class QObject *)
func (this *QMetaObject) ConnectSlotsByName(o *QObject /*777 QObject **/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject18connectSlotsByNameEP7QObject", ffiqt.FFI_TYPE_POINTER, o)
	gopp.ErrPrint(err, rv)
}
func QMetaObject_ConnectSlotsByName(o *QObject /*777 QObject **/) {
	var nilthis *QMetaObject
	nilthis.ConnectSlotsByName(o)
}

// /usr/include/qt/QtCore/qobjectdefs.h:401
// index:0
// Public static
// void activate(class QObject *, int, void **)
func (this *QMetaObject) Activate(sender *QObject /*777 QObject **/, signal_index int, argv unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject8activateEP7QObjectiPPv", ffiqt.FFI_TYPE_POINTER, sender, signal_index, argv)
	gopp.ErrPrint(err, rv)
}
func QMetaObject_Activate(sender *QObject /*777 QObject **/, signal_index int, argv unsafe.Pointer /*666*/) {
	var nilthis *QMetaObject
	nilthis.Activate(sender, signal_index, argv)
}

// /usr/include/qt/QtCore/qobjectdefs.h:403
// index:1
// Public static
// void activate(class QObject *, int, int, void **)
func (this *QMetaObject) Activate_1(sender *QObject /*777 QObject **/, signal_offset int, local_signal_index int, argv unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject8activateEP7QObjectiiPPv", ffiqt.FFI_TYPE_POINTER, sender, signal_offset, local_signal_index, argv)
	gopp.ErrPrint(err, rv)
}
func QMetaObject_Activate_1(sender *QObject /*777 QObject **/, signal_offset int, local_signal_index int, argv unsafe.Pointer /*666*/) {
	var nilthis *QMetaObject
	nilthis.Activate_1(sender, signal_offset, local_signal_index, argv)
}

// /usr/include/qt/QtCore/qobjectdefs.h:405
// index:0
// Public static
// bool invokeMethod(class QObject *, const char *, Qt::ConnectionType, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaObject) InvokeMethod(obj *QObject /*777 QObject **/, member string, arg2 int, ret *QGenericReturnArgument /*123*/, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS7_S7_S7_S7_S7_S7_S7_S7_S7_", ffiqt.FFI_TYPE_POINTER, obj, member, arg2, ret, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QMetaObject_InvokeMethod(obj *QObject /*777 QObject **/, member string, arg2 int, ret *QGenericReturnArgument /*123*/, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	var nilthis *QMetaObject
	rv := nilthis.InvokeMethod(obj, member, arg2, ret, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:419
// index:1
// Public static inline
// bool invokeMethod(class QObject *, const char *, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaObject) InvokeMethod_1(obj *QObject /*777 QObject **/, member string, ret *QGenericReturnArgument /*123*/, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", ffiqt.FFI_TYPE_POINTER, obj, member, ret, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QMetaObject_InvokeMethod_1(obj *QObject /*777 QObject **/, member string, ret *QGenericReturnArgument /*123*/, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	var nilthis *QMetaObject
	rv := nilthis.InvokeMethod_1(obj, member, ret, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:436
// index:2
// Public static inline
// bool invokeMethod(class QObject *, const char *, Qt::ConnectionType, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaObject) InvokeMethod_2(obj *QObject /*777 QObject **/, member string, type_ int, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKcN2Qt14ConnectionTypeE16QGenericArgumentS6_S6_S6_S6_S6_S6_S6_S6_S6_", ffiqt.FFI_TYPE_POINTER, obj, member, type_, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QMetaObject_InvokeMethod_2(obj *QObject /*777 QObject **/, member string, type_ int, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	var nilthis *QMetaObject
	rv := nilthis.InvokeMethod_2(obj, member, type_, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:453
// index:3
// Public static inline
// bool invokeMethod(class QObject *, const char *, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaObject) InvokeMethod_3(obj *QObject /*777 QObject **/, member string, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject12invokeMethodEP7QObjectPKc16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", ffiqt.FFI_TYPE_POINTER, obj, member, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QMetaObject_InvokeMethod_3(obj *QObject /*777 QObject **/, member string, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	var nilthis *QMetaObject
	rv := nilthis.InvokeMethod_3(obj, member, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	return rv
}

// /usr/include/qt/QtCore/qobjectdefs.h:554
// index:0
// Public
// QObject * newInstance(class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaObject) NewInstance(val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) *QObject /*777 QObject **/ {
	var convArg0 = val0.GetCthis()
	var convArg1 = val1.GetCthis()
	var convArg2 = val2.GetCthis()
	var convArg3 = val3.GetCthis()
	var convArg4 = val4.GetCthis()
	var convArg5 = val5.GetCthis()
	var convArg6 = val6.GetCthis()
	var convArg7 = val7.GetCthis()
	var convArg8 = val8.GetCthis()
	var convArg9 = val9.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject11newInstanceE16QGenericArgumentS0_S0_S0_S0_S0_S0_S0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qobjectdefs.h:581
// index:0
// Public
// int static_metacall(enum QMetaObject::Call, int, void **)
func (this *QMetaObject) Static_metacall(arg0 int, arg1 int, arg2 unsafe.Pointer /*666*/) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaObject15static_metacallENS_4CallEiPPv", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1, arg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qobjectdefs.h:582
// index:0
// Public static
// int metacall(class QObject *, enum QMetaObject::Call, int, void **)
func (this *QMetaObject) Metacall(arg0 *QObject /*777 QObject **/, arg1 int, arg2 int, arg3 unsafe.Pointer /*666*/) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaObject8metacallEP7QObjectNS_4CallEiPPv", ffiqt.FFI_TYPE_POINTER, arg0, arg1, arg2, arg3)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QMetaObject_Metacall(arg0 *QObject /*777 QObject **/, arg1 int, arg2 int, arg3 unsafe.Pointer /*666*/) int {
	var nilthis *QMetaObject
	rv := nilthis.Metacall(arg0, arg1, arg2, arg3)
	return rv
}

type QMetaObject__Call = int

const QMetaObject__InvokeMetaMethod QMetaObject__Call = 0
const QMetaObject__ReadProperty QMetaObject__Call = 1
const QMetaObject__WriteProperty QMetaObject__Call = 2
const QMetaObject__ResetProperty QMetaObject__Call = 3
const QMetaObject__QueryPropertyDesignable QMetaObject__Call = 4
const QMetaObject__QueryPropertyScriptable QMetaObject__Call = 5
const QMetaObject__QueryPropertyStored QMetaObject__Call = 6
const QMetaObject__QueryPropertyEditable QMetaObject__Call = 7
const QMetaObject__QueryPropertyUser QMetaObject__Call = 8
const QMetaObject__CreateInstance QMetaObject__Call = 9
const QMetaObject__IndexOfMethod QMetaObject__Call = 10
const QMetaObject__RegisterPropertyMetaType QMetaObject__Call = 11
const QMetaObject__RegisterMethodArgumentMetaType QMetaObject__Call = 12

//  body block end
