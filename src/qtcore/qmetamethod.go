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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qmetaobject.h:57
// index:0
// inline
// void QMetaMethod()
func NewQMetaMethod() *QMetaMethod {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMetaMethodC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QMetaMethod{cthis}
}

// /usr/include/qt/QtCore/qmetaobject.h:59
// index:0
// QByteArray methodSignature()
func (this *QMetaMethod) MethodSignature() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod15methodSignatureEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:60
// index:0
// QByteArray name()
func (this *QMetaMethod) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:61
// index:0
// const char * typeName()
func (this *QMetaMethod) TypeName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod8typeNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:62
// index:0
// int returnType()
func (this *QMetaMethod) ReturnType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod10returnTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:63
// index:0
// int parameterCount()
func (this *QMetaMethod) ParameterCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14parameterCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:64
// index:0
// int parameterType(int)
func (this *QMetaMethod) ParameterType(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod13parameterTypeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:65
// index:0
// void getParameterTypes(int *)
func (this *QMetaMethod) GetParameterTypes(types unsafe.Pointer) {
	// 0: (, int * types), (types)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod17getParameterTypesEPi", ffiqt.FFI_TYPE_VOID, this.cthis, types)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:66
// index:0
// QList<QByteArray> parameterTypes()
func (this *QMetaMethod) ParameterTypes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14parameterTypesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:67
// index:0
// QList<QByteArray> parameterNames()
func (this *QMetaMethod) ParameterNames() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14parameterNamesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:68
// index:0
// const char * tag()
func (this *QMetaMethod) Tag() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod3tagEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:70
// index:0
// QMetaMethod::Access access()
func (this *QMetaMethod) Access() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6accessEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:72
// index:0
// QMetaMethod::MethodType methodType()
func (this *QMetaMethod) MethodType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod10methodTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:74
// index:0
// int attributes()
func (this *QMetaMethod) Attributes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod10attributesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:75
// index:0
// int methodIndex()
func (this *QMetaMethod) MethodIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod11methodIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:76
// index:0
// int revision()
func (this *QMetaMethod) Revision() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod8revisionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:78
// index:0
// inline
// const QMetaObject * enclosingMetaObject()
func (this *QMetaMethod) EnclosingMetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod19enclosingMetaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// bool invoke(class QObject *, Qt::ConnectionType, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke(object unsafe.Pointer, connectionType int, returnValue unsafe.Pointer, val0 unsafe.Pointer, val1 unsafe.Pointer, val2 unsafe.Pointer, val3 unsafe.Pointer, val4 unsafe.Pointer, val5 unsafe.Pointer, val6 unsafe.Pointer, val7 unsafe.Pointer, val8 unsafe.Pointer, val9 unsafe.Pointer) {
	// 0: (, QObject * object, Qt::ConnectionType connectionType, QGenericReturnArgument returnValue, QGenericArgument val0, QGenericArgument val1, QGenericArgument val2, QGenericArgument val3, QGenericArgument val4, QGenericArgument val5, QGenericArgument val6, QGenericArgument val7, QGenericArgument val8, QGenericArgument val9), (object, &connectionType, returnValue, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", ffiqt.FFI_TYPE_VOID, this.cthis, object, &connectionType, returnValue, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// inline
// bool invoke(class QObject *, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke_1(object unsafe.Pointer, returnValue unsafe.Pointer, val0 unsafe.Pointer, val1 unsafe.Pointer, val2 unsafe.Pointer, val3 unsafe.Pointer, val4 unsafe.Pointer, val5 unsafe.Pointer, val6 unsafe.Pointer, val7 unsafe.Pointer, val8 unsafe.Pointer, val9 unsafe.Pointer) {
	// 1: (, QObject * object, QGenericReturnArgument returnValue, QGenericArgument val0, QGenericArgument val1, QGenericArgument val2, QGenericArgument val3, QGenericArgument val4, QGenericArgument val5, QGenericArgument val6, QGenericArgument val7, QGenericArgument val8, QGenericArgument val9), (object, returnValue, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", ffiqt.FFI_TYPE_VOID, this.cthis, object, returnValue, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// inline
// bool invoke(class QObject *, Qt::ConnectionType, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke_2(object unsafe.Pointer, connectionType int, val0 unsafe.Pointer, val1 unsafe.Pointer, val2 unsafe.Pointer, val3 unsafe.Pointer, val4 unsafe.Pointer, val5 unsafe.Pointer, val6 unsafe.Pointer, val7 unsafe.Pointer, val8 unsafe.Pointer, val9 unsafe.Pointer) {
	// 2: (, QObject * object, Qt::ConnectionType connectionType, QGenericArgument val0, QGenericArgument val1, QGenericArgument val2, QGenericArgument val3, QGenericArgument val4, QGenericArgument val5, QGenericArgument val6, QGenericArgument val7, QGenericArgument val8, QGenericArgument val9), (object, &connectionType, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", ffiqt.FFI_TYPE_VOID, this.cthis, object, &connectionType, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// inline
// bool invoke(class QObject *, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) Invoke_3(object unsafe.Pointer, val0 unsafe.Pointer, val1 unsafe.Pointer, val2 unsafe.Pointer, val3 unsafe.Pointer, val4 unsafe.Pointer, val5 unsafe.Pointer, val6 unsafe.Pointer, val7 unsafe.Pointer, val8 unsafe.Pointer, val9 unsafe.Pointer) {
	// 3: (, QObject * object, QGenericArgument val0, QGenericArgument val1, QGenericArgument val2, QGenericArgument val3, QGenericArgument val4, QGenericArgument val5, QGenericArgument val6, QGenericArgument val7, QGenericArgument val8, QGenericArgument val9), (object, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", ffiqt.FFI_TYPE_VOID, this.cthis, object, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// bool invokeOnGadget(void *, class QGenericReturnArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) InvokeOnGadget(gadget unsafe.Pointer, returnValue unsafe.Pointer, val0 unsafe.Pointer, val1 unsafe.Pointer, val2 unsafe.Pointer, val3 unsafe.Pointer, val4 unsafe.Pointer, val5 unsafe.Pointer, val6 unsafe.Pointer, val7 unsafe.Pointer, val8 unsafe.Pointer, val9 unsafe.Pointer) {
	// 0: (, void * gadget, QGenericReturnArgument returnValue, QGenericArgument val0, QGenericArgument val1, QGenericArgument val2, QGenericArgument val3, QGenericArgument val4, QGenericArgument val5, QGenericArgument val6, QGenericArgument val7, QGenericArgument val8, QGenericArgument val9), (gadget, returnValue, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", ffiqt.FFI_TYPE_VOID, this.cthis, gadget, returnValue, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// inline
// bool invokeOnGadget(void *, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument, class QGenericArgument)
func (this *QMetaMethod) InvokeOnGadget_1(gadget unsafe.Pointer, val0 unsafe.Pointer, val1 unsafe.Pointer, val2 unsafe.Pointer, val3 unsafe.Pointer, val4 unsafe.Pointer, val5 unsafe.Pointer, val6 unsafe.Pointer, val7 unsafe.Pointer, val8 unsafe.Pointer, val9 unsafe.Pointer) {
	// 1: (, void * gadget, QGenericArgument val0, QGenericArgument val1, QGenericArgument val2, QGenericArgument val3, QGenericArgument val4, QGenericArgument val5, QGenericArgument val6, QGenericArgument val7, QGenericArgument val8, QGenericArgument val9), (gadget, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", ffiqt.FFI_TYPE_VOID, this.cthis, gadget, val0, val1, val2, val3, val4, val5, val6, val7, val8, val9)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:169
// index:0
// inline
// bool isValid()
func (this *QMetaMethod) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMetaMethod7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
