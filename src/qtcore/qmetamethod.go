//  header block begin
// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>
package qtcore

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
type QMetaMethod struct {
	*qtrt.CObject
}

func (this *QMetaMethod) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQMetaMethodFromPointer(cthis unsafe.Pointer) *QMetaMethod {
	return &QMetaMethod{&qtrt.CObject{cthis}}
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
func (this *QMetaMethod) MethodSignature() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod15methodSignatureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:60
// index:0
// Public
// QByteArray name()
func (this *QMetaMethod) Name() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:61
// index:0
// Public
// const char * typeName()
func (this *QMetaMethod) TypeName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod8typeNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:62
// index:0
// Public
// int returnType()
func (this *QMetaMethod) ReturnType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod10returnTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:63
// index:0
// Public
// int parameterCount()
func (this *QMetaMethod) ParameterCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14parameterCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:64
// index:0
// Public
// int parameterType(int)
func (this *QMetaMethod) ParameterType(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod13parameterTypeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:65
// index:0
// Public
// void getParameterTypes(int *)
func (this *QMetaMethod) GetParameterTypes(types unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod17getParameterTypesEPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), types)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:66
// index:0
// Public
// QList<QByteArray> parameterTypes()
func (this *QMetaMethod) ParameterTypes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14parameterTypesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:67
// index:0
// Public
// QList<QByteArray> parameterNames()
func (this *QMetaMethod) ParameterNames() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14parameterNamesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:68
// index:0
// Public
// const char * tag()
func (this *QMetaMethod) Tag() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod3tagEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:70
// index:0
// Public
// QMetaMethod::Access access()
func (this *QMetaMethod) Access() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6accessEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:72
// index:0
// Public
// QMetaMethod::MethodType methodType()
func (this *QMetaMethod) MethodType() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod10methodTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:74
// index:0
// Public
// int attributes()
func (this *QMetaMethod) Attributes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod10attributesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:75
// index:0
// Public
// int methodIndex()
func (this *QMetaMethod) MethodIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod11methodIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:76
// index:0
// Public
// int revision()
func (this *QMetaMethod) Revision() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod8revisionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:78
// index:0
// Public inline
// const QMetaObject * enclosingMetaObject()
func (this *QMetaMethod) EnclosingMetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod19enclosingMetaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public
// bool invoke(class QObject *, Qt::ConnectionType, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke(object unsafe.Pointer, connectionType int, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) interface{} {
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object, &connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline
// bool invoke(class QObject *, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke_1(object unsafe.Pointer, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) interface{} {
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline
// bool invoke(class QObject *, Qt::ConnectionType, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke_2(object unsafe.Pointer, connectionType int, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) interface{} {
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object, &connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline
// bool invoke(class QObject *, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke_3(object unsafe.Pointer, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) interface{} {
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public
// bool invokeOnGadget(void *, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) InvokeOnGadget(gadget unsafe.Pointer, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) interface{} {
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
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline
// bool invokeOnGadget(void *, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) InvokeOnGadget_1(gadget unsafe.Pointer, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) interface{} {
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
	return rv
}

// /usr/include/qt/QtCore/qmetaobject.h:169
// index:0
// Public inline
// bool isValid()
func (this *QMetaMethod) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
