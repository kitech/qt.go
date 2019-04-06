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
// extern C begin: 9
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
type QMetaMethod struct {
	*qtrt.CObject
}
type QMetaMethod_ITF interface {
	QMetaMethod_PTR() *QMetaMethod
}

func (ptr *QMetaMethod) QMetaMethod_PTR() *QMetaMethod { return ptr }

func (this *QMetaMethod) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMetaMethod) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMetaMethodFromPointer(cthis unsafe.Pointer) *QMetaMethod {
	return &QMetaMethod{&qtrt.CObject{cthis}}
}
func (*QMetaMethod) NewFromPointer(cthis unsafe.Pointer) *QMetaMethod {
	return NewQMetaMethodFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmetaobject.h:57
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QMetaMethod()

/*

 */
func (*QMetaMethod) NewForInherit() *QMetaMethod {
	return NewQMetaMethod()
}
func NewQMetaMethod() *QMetaMethod {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaMethodC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMetaMethodFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMetaMethod)
	return gothis
}

// /usr/include/qt/QtCore/qmetaobject.h:59
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray methodSignature() const

/*

 */
func (this *QMetaMethod) MethodSignature() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod15methodSignatureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray name() const

/*

 */
func (this *QMetaMethod) Name() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:61
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * typeName() const

/*

 */
func (this *QMetaMethod) TypeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod8typeNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:62
// index:0
// Public Visibility=Default Availability=Available
// [4] int returnType() const

/*

 */
func (this *QMetaMethod) ReturnType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod10returnTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] int parameterCount() const

/*

 */
func (this *QMetaMethod) ParameterCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14parameterCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] int parameterType(int) const

/*

 */
func (this *QMetaMethod) ParameterType(index int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod13parameterTypeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getParameterTypes(int *) const

/*

 */
func (this *QMetaMethod) GetParameterTypes(types unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod17getParameterTypesEPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), types)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * tag() const

/*

 */
func (this *QMetaMethod) Tag() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod3tagEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] QMetaMethod::Access access() const

/*

 */
func (this *QMetaMethod) Access() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6accessEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:72
// index:0
// Public Visibility=Default Availability=Available
// [4] QMetaMethod::MethodType methodType() const

/*

 */
func (this *QMetaMethod) MethodType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod10methodTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int attributes() const

/*

 */
func (this *QMetaMethod) Attributes() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod10attributesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int methodIndex() const

/*

 */
func (this *QMetaMethod) MethodIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod11methodIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int revision() const

/*

 */
func (this *QMetaMethod) Revision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod8revisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QMetaObject * enclosingMetaObject() const

/*

 */
func (this *QMetaMethod) EnclosingMetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod19enclosingMetaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invokep(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invokep1(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invokep2(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invokep3(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invokep4(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invokep5(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invokep6(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invokep7(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invokep8(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invokep9(object QObject_ITF /*777 QObject **/, connectionType int, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg2 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE22QGenericReturnArgument16QGenericArgumentS5_S5_S5_S5_S5_S5_S5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11, convArg12)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1p(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1p1(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1p2(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1p3(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1p4(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1p5(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1p6(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1p7(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1p8(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:93
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke1p9(object QObject_ITF /*777 QObject **/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject22QGenericReturnArgument16QGenericArgumentS3_S3_S3_S3_S3_S3_S3_S3_S3_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2(object QObject_ITF /*777 QObject **/, connectionType int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2p(object QObject_ITF /*777 QObject **/, connectionType int) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
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
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2p1(object QObject_ITF /*777 QObject **/, connectionType int, val0 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2p2(object QObject_ITF /*777 QObject **/, connectionType int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2p3(object QObject_ITF /*777 QObject **/, connectionType int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2p4(object QObject_ITF /*777 QObject **/, connectionType int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2p5(object QObject_ITF /*777 QObject **/, connectionType int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2p6(object QObject_ITF /*777 QObject **/, connectionType int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2p7(object QObject_ITF /*777 QObject **/, connectionType int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2p8(object QObject_ITF /*777 QObject **/, connectionType int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:109
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, Qt::ConnectionType, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke2p9(object QObject_ITF /*777 QObject **/, connectionType int, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObjectN2Qt14ConnectionTypeE16QGenericArgumentS4_S4_S4_S4_S4_S4_S4_S4_S4_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, connectionType, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3(object QObject_ITF /*777 QObject **/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg6 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg7 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg8 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg9 = val8.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val9 != nil && val9.QGenericArgument_PTR() != nil {
		convArg10 = val9.QGenericArgument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3p(object QObject_ITF /*777 QObject **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
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
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3p1(object QObject_ITF /*777 QObject **/, val0 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
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
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3p2(object QObject_ITF /*777 QObject **/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3p3(object QObject_ITF /*777 QObject **/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3p4(object QObject_ITF /*777 QObject **/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3p5(object QObject_ITF /*777 QObject **/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3p6(object QObject_ITF /*777 QObject **/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg6 = val5.QGenericArgument_PTR().GetCthis()
	}
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3p7(object QObject_ITF /*777 QObject **/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg6 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg7 = val6.QGenericArgument_PTR().GetCthis()
	}
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3p8(object QObject_ITF /*777 QObject **/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg6 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg7 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg8 = val7.QGenericArgument_PTR().GetCthis()
	}
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:125
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool invoke(QObject *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) Invoke3p9(object QObject_ITF /*777 QObject **/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg6 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg7 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg8 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg9 = val8.QGenericArgument_PTR().GetCthis()
	}
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod6invokeEP7QObject16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadgetp(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
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
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	// arg: 11, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg11 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadgetp1(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadgetp2(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadgetp3(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadgetp4(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadgetp5(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadgetp6(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadgetp7(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadgetp8(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericReturnArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadgetp9(gadget unsafe.Pointer /*666*/, returnValue QGenericReturnArgument_ITF /*123*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if returnValue != nil && returnValue.QGenericReturnArgument_PTR() != nil {
		convArg1 = returnValue.QGenericReturnArgument_PTR().GetCthis()
	}
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv22QGenericReturnArgument16QGenericArgumentS2_S2_S2_S2_S2_S2_S2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10, convArg11)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1(gadget unsafe.Pointer /*666*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/, val9 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg6 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg7 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg8 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg9 = val8.QGenericArgument_PTR().GetCthis()
	}
	var convArg10 unsafe.Pointer
	if val9 != nil && val9.QGenericArgument_PTR() != nil {
		convArg10 = val9.QGenericArgument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1p(gadget unsafe.Pointer /*666*/) bool {
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
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1p1(gadget unsafe.Pointer /*666*/, val0 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
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
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1p2(gadget unsafe.Pointer /*666*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1p3(gadget unsafe.Pointer /*666*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1p4(gadget unsafe.Pointer /*666*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1p5(gadget unsafe.Pointer /*666*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
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
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1p6(gadget unsafe.Pointer /*666*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg6 = val5.QGenericArgument_PTR().GetCthis()
	}
	// arg: 7, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg7 unsafe.Pointer
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1p7(gadget unsafe.Pointer /*666*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg6 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg7 = val6.QGenericArgument_PTR().GetCthis()
	}
	// arg: 8, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg8 unsafe.Pointer
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1p8(gadget unsafe.Pointer /*666*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg6 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg7 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg8 = val7.QGenericArgument_PTR().GetCthis()
	}
	// arg: 9, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg9 unsafe.Pointer
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:153
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool invokeOnGadget(void *, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument, QGenericArgument) const

/*

 */
func (this *QMetaMethod) InvokeOnGadget1p9(gadget unsafe.Pointer /*666*/, val0 QGenericArgument_ITF /*123*/, val1 QGenericArgument_ITF /*123*/, val2 QGenericArgument_ITF /*123*/, val3 QGenericArgument_ITF /*123*/, val4 QGenericArgument_ITF /*123*/, val5 QGenericArgument_ITF /*123*/, val6 QGenericArgument_ITF /*123*/, val7 QGenericArgument_ITF /*123*/, val8 QGenericArgument_ITF /*123*/) bool {
	var convArg1 unsafe.Pointer
	if val0 != nil && val0.QGenericArgument_PTR() != nil {
		convArg1 = val0.QGenericArgument_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if val1 != nil && val1.QGenericArgument_PTR() != nil {
		convArg2 = val1.QGenericArgument_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if val2 != nil && val2.QGenericArgument_PTR() != nil {
		convArg3 = val2.QGenericArgument_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if val3 != nil && val3.QGenericArgument_PTR() != nil {
		convArg4 = val3.QGenericArgument_PTR().GetCthis()
	}
	var convArg5 unsafe.Pointer
	if val4 != nil && val4.QGenericArgument_PTR() != nil {
		convArg5 = val4.QGenericArgument_PTR().GetCthis()
	}
	var convArg6 unsafe.Pointer
	if val5 != nil && val5.QGenericArgument_PTR() != nil {
		convArg6 = val5.QGenericArgument_PTR().GetCthis()
	}
	var convArg7 unsafe.Pointer
	if val6 != nil && val6.QGenericArgument_PTR() != nil {
		convArg7 = val6.QGenericArgument_PTR().GetCthis()
	}
	var convArg8 unsafe.Pointer
	if val7 != nil && val7.QGenericArgument_PTR() != nil {
		convArg8 = val7.QGenericArgument_PTR().GetCthis()
	}
	var convArg9 unsafe.Pointer
	if val8 != nil && val8.QGenericArgument_PTR() != nil {
		convArg9 = val8.QGenericArgument_PTR().GetCthis()
	}
	// arg: 10, QGenericArgument=Record, QGenericArgument=Record, , Invalid
	var convArg10 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod14invokeOnGadgetEPv16QGenericArgumentS1_S1_S1_S1_S1_S1_S1_S1_S1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1, convArg2, convArg3, convArg4, convArg5, convArg6, convArg7, convArg8, convArg9, convArg10)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:169
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QMetaMethod) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMetaMethod7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQMetaMethod(this *QMetaMethod) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMetaMethodD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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

func init_unused_10475() {
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
