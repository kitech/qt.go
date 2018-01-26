package qtcore

// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QMetaMethod struct {
	*qtrt.CObject
}

func (this *QMetaMethod) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMetaMethod) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQMetaMethodFromPointer(cthis unsafe.Pointer) *QMetaMethod {
	return &QMetaMethod{&qtrt.CObject{cthis}}
}
func (*QMetaMethod) NewFromPointer(cthis unsafe.Pointer) *QMetaMethod {
	return NewQMetaMethodFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmetaobject.h:57
// index:0
// Public inline
// void QMetaMethod()
func NewQMetaMethod() *QMetaMethod {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaMethodC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMetaMethodFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmetaobject.h:59
// index:0
// Public
// QByteArray methodSignature()
func (this *QMetaMethod) MethodSignature() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod15methodSignatureEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:60
// index:0
// Public
// QByteArray name()
func (this *QMetaMethod) Name() *QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:61
// index:0
// Public
// const char * typeName()
func (this *QMetaMethod) TypeName() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod8typeNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:62
// index:0
// Public
// int returnType()
func (this *QMetaMethod) ReturnType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod10returnTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmetaobject.h:63
// index:0
// Public
// int parameterCount()
func (this *QMetaMethod) ParameterCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14parameterCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmetaobject.h:64
// index:0
// Public
// int parameterType(int)
func (this *QMetaMethod) ParameterType(index int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod13parameterTypeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmetaobject.h:65
// index:0
// Public
// void getParameterTypes(int *)
func (this *QMetaMethod) GetParameterTypes(types unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod17getParameterTypesEPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &types)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:68
// index:0
// Public
// const char * tag()
func (this *QMetaMethod) Tag() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod3tagEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:70
// index:0
// Public
// QMetaMethod::Access access()
func (this *QMetaMethod) Access() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6accessEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:72
// index:0
// Public
// QMetaMethod::MethodType methodType()
func (this *QMetaMethod) MethodType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod10methodTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:74
// index:0
// Public
// int attributes()
func (this *QMetaMethod) Attributes() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod10attributesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmetaobject.h:75
// index:0
// Public
// int methodIndex()
func (this *QMetaMethod) MethodIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod11methodIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmetaobject.h:76
// index:0
// Public
// int revision()
func (this *QMetaMethod) Revision() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod8revisionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qmetaobject.h:78
// index:0
// Public inline
// const QMetaObject * enclosingMetaObject()
func (this *QMetaMethod) EnclosingMetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod19enclosingMetaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public
// bool invoke(class QObject *, Qt::ConnectionType, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke(object *QObject /*777 QObject **/, connectionType int, returnValue *QGenericReturnArgument /*123*/, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	var convArg0 = object.GetCthis()
	var convArg2 = returnValue.GetCthis()
	var convArg3 = val0.GetCthis()
	var convArg4 = val1.GetCthis()
	var convArg5 = val2.GetCthis()
	var convArg6 = val3.GetCthis()
	var convArg7 = val4.GetCthis()
	var convArg8 = val5.GetCthis()
	var convArg9 = val6.GetCthis()
	var convArg10 = val7.GetCthis()
	var convArg11 = val8.GetCthis()
	var convArg12 = val9.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline
// bool invoke(class QObject *, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke_1(object *QObject /*777 QObject **/, returnValue *QGenericReturnArgument /*123*/, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	var convArg0 = object.GetCthis()
	var convArg1 = returnValue.GetCthis()
	var convArg2 = val0.GetCthis()
	var convArg3 = val1.GetCthis()
	var convArg4 = val2.GetCthis()
	var convArg5 = val3.GetCthis()
	var convArg6 = val4.GetCthis()
	var convArg7 = val5.GetCthis()
	var convArg8 = val6.GetCthis()
	var convArg9 = val7.GetCthis()
	var convArg10 = val8.GetCthis()
	var convArg11 = val9.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline
// bool invoke(class QObject *, Qt::ConnectionType, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke_2(object *QObject /*777 QObject **/, connectionType int, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	var convArg0 = object.GetCthis()
	var convArg2 = val0.GetCthis()
	var convArg3 = val1.GetCthis()
	var convArg4 = val2.GetCthis()
	var convArg5 = val3.GetCthis()
	var convArg6 = val4.GetCthis()
	var convArg7 = val5.GetCthis()
	var convArg8 = val6.GetCthis()
	var convArg9 = val7.GetCthis()
	var convArg10 = val8.GetCthis()
	var convArg11 = val9.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline
// bool invoke(class QObject *, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke_3(object *QObject /*777 QObject **/, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	var convArg0 = object.GetCthis()
	var convArg1 = val0.GetCthis()
	var convArg2 = val1.GetCthis()
	var convArg3 = val2.GetCthis()
	var convArg4 = val3.GetCthis()
	var convArg5 = val4.GetCthis()
	var convArg6 = val5.GetCthis()
	var convArg7 = val6.GetCthis()
	var convArg8 = val7.GetCthis()
	var convArg9 = val8.GetCthis()
	var convArg10 = val9.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public
// bool invokeOnGadget(void *, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) InvokeOnGadget(gadget unsafe.Pointer /*666*/, returnValue *QGenericReturnArgument /*123*/, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	var convArg1 = returnValue.GetCthis()
	var convArg2 = val0.GetCthis()
	var convArg3 = val1.GetCthis()
	var convArg4 = val2.GetCthis()
	var convArg5 = val3.GetCthis()
	var convArg6 = val4.GetCthis()
	var convArg7 = val5.GetCthis()
	var convArg8 = val6.GetCthis()
	var convArg9 = val7.GetCthis()
	var convArg10 = val8.GetCthis()
	var convArg11 = val9.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline
// bool invokeOnGadget(void *, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) InvokeOnGadget_1(gadget unsafe.Pointer /*666*/, val0 *QGenericArgument /*123*/, val1 *QGenericArgument /*123*/, val2 *QGenericArgument /*123*/, val3 *QGenericArgument /*123*/, val4 *QGenericArgument /*123*/, val5 *QGenericArgument /*123*/, val6 *QGenericArgument /*123*/, val7 *QGenericArgument /*123*/, val8 *QGenericArgument /*123*/, val9 *QGenericArgument /*123*/) bool {
	var convArg1 = val0.GetCthis()
	var convArg2 = val1.GetCthis()
	var convArg3 = val2.GetCthis()
	var convArg4 = val3.GetCthis()
	var convArg5 = val4.GetCthis()
	var convArg6 = val5.GetCthis()
	var convArg7 = val6.GetCthis()
	var convArg8 = val7.GetCthis()
	var convArg9 = val8.GetCthis()
	var convArg10 = val9.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:169
// index:0
// Public inline
// bool isValid()
func (this *QMetaMethod) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QMetaMethod__Access = int

const QMetaMethod__Private QMetaMethod__Access = 0
const QMetaMethod__Protected QMetaMethod__Access = 1
const QMetaMethod__Public QMetaMethod__Access = 2

type QMetaMethod__MethodType = int

const QMetaMethod__Method QMetaMethod__MethodType = 0
const QMetaMethod__Signal QMetaMethod__MethodType = 1
const QMetaMethod__Slot QMetaMethod__MethodType = 2
const QMetaMethod__Constructor QMetaMethod__MethodType = 3

type QMetaMethod__Attributes = int

const QMetaMethod__Compatibility QMetaMethod__Attributes = 1
const QMetaMethod__Cloned QMetaMethod__Attributes = 2
const QMetaMethod__Scriptable QMetaMethod__Attributes = 4

//  body block end
