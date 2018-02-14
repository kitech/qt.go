package qtqml

// /usr/include/qt/QtQml/qqmlproperty.h
// #include <qqmlproperty.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QQmlProperty struct {
	*qtrt.CObject
}
type QQmlProperty_ITF interface {
	QQmlProperty_PTR() *QQmlProperty
}

func (ptr *QQmlProperty) QQmlProperty_PTR() *QQmlProperty { return ptr }

func (this *QQmlProperty) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlProperty) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlPropertyFromPointer(cthis unsafe.Pointer) *QQmlProperty {
	return &QQmlProperty{&qtrt.CObject{cthis}}
}
func (*QQmlProperty) NewFromPointer(cthis unsafe.Pointer) *QQmlProperty {
	return NewQQmlPropertyFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlproperty.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlProperty()
func NewQQmlProperty() *QQmlProperty {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlPropertyC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlPropertyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlProperty)
	return gothis
}

// /usr/include/qt/QtQml/qqmlproperty.h:74
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQmlProperty(QObject *)
func NewQQmlProperty_1(arg0 qtcore.QObject_ITF /*777 QObject **/) *QQmlProperty {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlPropertyC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlPropertyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlProperty)
	return gothis
}

// /usr/include/qt/QtQml/qqmlproperty.h:75
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QQmlProperty(QObject *, QQmlContext *)
func NewQQmlProperty_2(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 QQmlContext_ITF /*777 QQmlContext **/) *QQmlProperty {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var convArg1 = arg1.QQmlContext_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlPropertyC2EP7QObjectP11QQmlContext", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlPropertyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlProperty)
	return gothis
}

// /usr/include/qt/QtQml/qqmlproperty.h:76
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QQmlProperty(QObject *, QQmlEngine *)
func NewQQmlProperty_3(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 QQmlEngine_ITF /*777 QQmlEngine **/) *QQmlProperty {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var convArg1 = arg1.QQmlEngine_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlPropertyC2EP7QObjectP10QQmlEngine", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlPropertyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlProperty)
	return gothis
}

// /usr/include/qt/QtQml/qqmlproperty.h:78
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QQmlProperty(QObject *, const QString &)
func NewQQmlProperty_4(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string) *QQmlProperty {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlPropertyC2EP7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlPropertyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlProperty)
	return gothis
}

// /usr/include/qt/QtQml/qqmlproperty.h:79
// index:5
// Public Visibility=Default Availability=Available
// [-2] void QQmlProperty(QObject *, const QString &, QQmlContext *)
func NewQQmlProperty_5(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string, arg2 QQmlContext_ITF /*777 QQmlContext **/) *QQmlProperty {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = arg2.QQmlContext_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlPropertyC2EP7QObjectRK7QStringP11QQmlContext", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlPropertyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlProperty)
	return gothis
}

// /usr/include/qt/QtQml/qqmlproperty.h:80
// index:6
// Public Visibility=Default Availability=Available
// [-2] void QQmlProperty(QObject *, const QString &, QQmlEngine *)
func NewQQmlProperty_6(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string, arg2 QQmlEngine_ITF /*777 QQmlEngine **/) *QQmlProperty {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = arg2.QQmlEngine_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlPropertyC2EP7QObjectRK7QStringP10QQmlEngine", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlPropertyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlProperty)
	return gothis
}

// /usr/include/qt/QtQml/qqmlproperty.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QQmlProperty()
func DeleteQQmlProperty(this *QQmlProperty) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlPropertyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlproperty.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] QQmlProperty::Type type()
func (this *QQmlProperty) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlproperty.h:88
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QQmlProperty) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isProperty()
func (this *QQmlProperty) IsProperty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty10isPropertyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSignalProperty()
func (this *QQmlProperty) IsSignalProperty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty16isSignalPropertyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int propertyType()
func (this *QQmlProperty) PropertyType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty12propertyTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlproperty.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QQmlProperty::PropertyTypeCategory propertyTypeCategory()
func (this *QQmlProperty) PropertyTypeCategory() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty20propertyTypeCategoryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlproperty.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * propertyTypeName()
func (this *QQmlProperty) PropertyTypeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty16propertyTypeNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtQml/qqmlproperty.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QQmlProperty) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlproperty.h:98
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant read()
func (this *QQmlProperty) Read() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty4readEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQml/qqmlproperty.h:99
// index:1
// Public static Visibility=Default Availability=Available
// [16] QVariant read(const QObject *, const QString &)
func (this *QQmlProperty) Read_1(arg0 qtcore.QObject_ITF /*777 const QObject **/, arg1 string) *qtcore.QVariant /*123*/ {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlProperty4readEPK7QObjectRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}
func QQmlProperty_Read_1(arg0 qtcore.QObject_ITF /*777 const QObject **/, arg1 string) *qtcore.QVariant /*123*/ {
	var nilthis *QQmlProperty
	rv := nilthis.Read_1(arg0, arg1)
	return rv
}

// /usr/include/qt/QtQml/qqmlproperty.h:100
// index:2
// Public static Visibility=Default Availability=Available
// [16] QVariant read(const QObject *, const QString &, QQmlContext *)
func (this *QQmlProperty) Read_2(arg0 qtcore.QObject_ITF /*777 const QObject **/, arg1 string, arg2 QQmlContext_ITF /*777 QQmlContext **/) *qtcore.QVariant /*123*/ {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = arg2.QQmlContext_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlProperty4readEPK7QObjectRK7QStringP11QQmlContext", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}
func QQmlProperty_Read_2(arg0 qtcore.QObject_ITF /*777 const QObject **/, arg1 string, arg2 QQmlContext_ITF /*777 QQmlContext **/) *qtcore.QVariant /*123*/ {
	var nilthis *QQmlProperty
	rv := nilthis.Read_2(arg0, arg1, arg2)
	return rv
}

// /usr/include/qt/QtQml/qqmlproperty.h:101
// index:3
// Public static Visibility=Default Availability=Available
// [16] QVariant read(const QObject *, const QString &, QQmlEngine *)
func (this *QQmlProperty) Read_3(arg0 qtcore.QObject_ITF /*777 const QObject **/, arg1 string, arg2 QQmlEngine_ITF /*777 QQmlEngine **/) *qtcore.QVariant /*123*/ {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = arg2.QQmlEngine_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlProperty4readEPK7QObjectRK7QStringP10QQmlEngine", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}
func QQmlProperty_Read_3(arg0 qtcore.QObject_ITF /*777 const QObject **/, arg1 string, arg2 QQmlEngine_ITF /*777 QQmlEngine **/) *qtcore.QVariant /*123*/ {
	var nilthis *QQmlProperty
	rv := nilthis.Read_3(arg0, arg1, arg2)
	return rv
}

// /usr/include/qt/QtQml/qqmlproperty.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool write(const QVariant &)
func (this *QQmlProperty) Write(arg0 qtcore.QVariant_ITF) bool {
	var convArg0 = arg0.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty5writeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:104
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool write(QObject *, const QString &, const QVariant &)
func (this *QQmlProperty) Write_1(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string, arg2 qtcore.QVariant_ITF) bool {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = arg2.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlProperty5writeEP7QObjectRK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQmlProperty_Write_1(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string, arg2 qtcore.QVariant_ITF) bool {
	var nilthis *QQmlProperty
	rv := nilthis.Write_1(arg0, arg1, arg2)
	return rv
}

// /usr/include/qt/QtQml/qqmlproperty.h:105
// index:2
// Public static Visibility=Default Availability=Available
// [1] bool write(QObject *, const QString &, const QVariant &, QQmlContext *)
func (this *QQmlProperty) Write_2(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string, arg2 qtcore.QVariant_ITF, arg3 QQmlContext_ITF /*777 QQmlContext **/) bool {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = arg2.QVariant_PTR().GetCthis()
	var convArg3 = arg3.QQmlContext_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlProperty5writeEP7QObjectRK7QStringRK8QVariantP11QQmlContext", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQmlProperty_Write_2(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string, arg2 qtcore.QVariant_ITF, arg3 QQmlContext_ITF /*777 QQmlContext **/) bool {
	var nilthis *QQmlProperty
	rv := nilthis.Write_2(arg0, arg1, arg2, arg3)
	return rv
}

// /usr/include/qt/QtQml/qqmlproperty.h:106
// index:3
// Public static Visibility=Default Availability=Available
// [1] bool write(QObject *, const QString &, const QVariant &, QQmlEngine *)
func (this *QQmlProperty) Write_3(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string, arg2 qtcore.QVariant_ITF, arg3 QQmlEngine_ITF /*777 QQmlEngine **/) bool {
	var convArg0 = arg0.QObject_PTR().GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = arg2.QVariant_PTR().GetCthis()
	var convArg3 = arg3.QQmlEngine_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQmlProperty5writeEP7QObjectRK7QStringRK8QVariantP10QQmlEngine", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQmlProperty_Write_3(arg0 qtcore.QObject_ITF /*777 QObject **/, arg1 string, arg2 qtcore.QVariant_ITF, arg3 QQmlEngine_ITF /*777 QQmlEngine **/) bool {
	var nilthis *QQmlProperty
	rv := nilthis.Write_3(arg0, arg1, arg2, arg3)
	return rv
}

// /usr/include/qt/QtQml/qqmlproperty.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool reset()
func (this *QQmlProperty) Reset() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty5resetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:110
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasNotifySignal()
func (this *QQmlProperty) HasNotifySignal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty15hasNotifySignalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool needsNotifySignal()
func (this *QQmlProperty) NeedsNotifySignal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty17needsNotifySignalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool connectNotifySignal(QObject *, const char *)
func (this *QQmlProperty) ConnectNotifySignal(dest qtcore.QObject_ITF /*777 QObject **/, slot string) bool {
	var convArg0 = dest.QObject_PTR().GetCthis()
	var convArg1 = qtrt.CString(slot)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty19connectNotifySignalEP7QObjectPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:113
// index:1
// Public Visibility=Default Availability=Available
// [1] bool connectNotifySignal(QObject *, int)
func (this *QQmlProperty) ConnectNotifySignal_1(dest qtcore.QObject_ITF /*777 QObject **/, method int) bool {
	var convArg0 = dest.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty19connectNotifySignalEP7QObjecti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, method)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWritable()
func (this *QQmlProperty) IsWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty10isWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDesignable()
func (this *QQmlProperty) IsDesignable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty12isDesignableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isResettable()
func (this *QQmlProperty) IsResettable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty12isResettableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlproperty.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * object()
func (this *QQmlProperty) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlproperty.h:120
// index:0
// Public Visibility=Default Availability=Available
// [4] int index()
func (this *QQmlProperty) Index() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty5indexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlproperty.h:121
// index:0
// Public Visibility=Default Availability=Available
// [32] QMetaProperty property()
func (this *QQmlProperty) Property() *qtcore.QMetaProperty /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty8propertyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMetaPropertyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMetaProperty)
	return rv2
}

// /usr/include/qt/QtQml/qqmlproperty.h:122
// index:0
// Public Visibility=Default Availability=Available
// [16] QMetaMethod method()
func (this *QQmlProperty) Method() *qtcore.QMetaMethod /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQmlProperty6methodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMetaMethodFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMetaMethod)
	return rv2
}

type QQmlProperty__PropertyTypeCategory = int

const QQmlProperty__InvalidCategory QQmlProperty__PropertyTypeCategory = 0
const QQmlProperty__List QQmlProperty__PropertyTypeCategory = 1
const QQmlProperty__Object QQmlProperty__PropertyTypeCategory = 2
const QQmlProperty__Normal QQmlProperty__PropertyTypeCategory = 3

type QQmlProperty__Type = int

const QQmlProperty__Invalid QQmlProperty__Type = 0
const QQmlProperty__Property QQmlProperty__Type = 1
const QQmlProperty__SignalProperty QQmlProperty__Type = 2

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
