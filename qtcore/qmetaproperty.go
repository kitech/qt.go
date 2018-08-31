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
// extern C begin: 14
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
type QMetaProperty struct {
	*qtrt.CObject
}
type QMetaProperty_ITF interface {
	QMetaProperty_PTR() *QMetaProperty
}

func (ptr *QMetaProperty) QMetaProperty_PTR() *QMetaProperty { return ptr }

func (this *QMetaProperty) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMetaProperty) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMetaPropertyFromPointer(cthis unsafe.Pointer) *QMetaProperty {
	return &QMetaProperty{&qtrt.CObject{cthis}}
}
func (*QMetaProperty) NewFromPointer(cthis unsafe.Pointer) *QMetaProperty {
	return NewQMetaPropertyFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmetaobject.h:248
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMetaProperty()

/*

 */
func (*QMetaProperty) NewForInherit() *QMetaProperty {
	return NewQMetaProperty()
}
func NewQMetaProperty() *QMetaProperty {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMetaPropertyC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMetaPropertyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMetaProperty)
	return gothis
}

// /usr/include/qt/QtCore/qmetaobject.h:250
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * name() const

/*

 */
func (this *QMetaProperty) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:251
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * typeName() const

/*

 */
func (this *QMetaProperty) TypeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty8typeNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:252
// index:0
// Public Visibility=Default Availability=Available
// [4] QVariant::Type type() const

/*

 */
func (this *QMetaProperty) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:253
// index:0
// Public Visibility=Default Availability=Available
// [4] int userType() const

/*

 */
func (this *QMetaProperty) UserType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty8userTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:254
// index:0
// Public Visibility=Default Availability=Available
// [4] int propertyIndex() const

/*

 */
func (this *QMetaProperty) PropertyIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty13propertyIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:256
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadable() const

/*

 */
func (this *QMetaProperty) IsReadable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isReadableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:257
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWritable() const

/*

 */
func (this *QMetaProperty) IsWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:258
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isResettable() const

/*

 */
func (this *QMetaProperty) IsResettable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12isResettableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:259
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDesignable(const QObject *) const

/*

 */
func (this *QMetaProperty) IsDesignable(obj QObject_ITF /*777 const QObject **/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12isDesignableEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:259
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDesignable(const QObject *) const

/*

 */
func (this *QMetaProperty) IsDesignable__() bool {
	// arg: 0, const QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12isDesignableEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:260
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isScriptable(const QObject *) const

/*

 */
func (this *QMetaProperty) IsScriptable(obj QObject_ITF /*777 const QObject **/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12isScriptableEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:260
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isScriptable(const QObject *) const

/*

 */
func (this *QMetaProperty) IsScriptable__() bool {
	// arg: 0, const QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12isScriptableEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:261
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStored(const QObject *) const

/*

 */
func (this *QMetaProperty) IsStored(obj QObject_ITF /*777 const QObject **/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty8isStoredEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:261
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStored(const QObject *) const

/*

 */
func (this *QMetaProperty) IsStored__() bool {
	// arg: 0, const QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty8isStoredEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:262
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEditable(const QObject *) const

/*

 */
func (this *QMetaProperty) IsEditable(obj QObject_ITF /*777 const QObject **/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isEditableEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:262
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEditable(const QObject *) const

/*

 */
func (this *QMetaProperty) IsEditable__() bool {
	// arg: 0, const QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isEditableEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:263
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUser(const QObject *) const

/*

 */
func (this *QMetaProperty) IsUser(obj QObject_ITF /*777 const QObject **/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty6isUserEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:263
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUser(const QObject *) const

/*

 */
func (this *QMetaProperty) IsUser__() bool {
	// arg: 0, const QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty6isUserEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:264
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isConstant() const

/*

 */
func (this *QMetaProperty) IsConstant() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isConstantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:265
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFinal() const

/*

 */
func (this *QMetaProperty) IsFinal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty7isFinalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:267
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlagType() const

/*

 */
func (this *QMetaProperty) IsFlagType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isFlagTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:268
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnumType() const

/*

 */
func (this *QMetaProperty) IsEnumType() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10isEnumTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:269
// index:0
// Public Visibility=Default Availability=Available
// [16] QMetaEnum enumerator() const

/*
Returns the meta-data for the enumerator with the given index.

See also enumeratorCount(), enumeratorOffset(), and indexOfEnumerator().
*/
func (this *QMetaProperty) Enumerator() *QMetaEnum /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty10enumeratorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMetaEnumFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaEnum)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:271
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasNotifySignal() const

/*

 */
func (this *QMetaProperty) HasNotifySignal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty15hasNotifySignalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:272
// index:0
// Public Visibility=Default Availability=Available
// [16] QMetaMethod notifySignal() const

/*

 */
func (this *QMetaProperty) NotifySignal() *QMetaMethod /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12notifySignalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMetaMethodFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMetaMethod)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:273
// index:0
// Public Visibility=Default Availability=Available
// [4] int notifySignalIndex() const

/*

 */
func (this *QMetaProperty) NotifySignalIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty17notifySignalIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:275
// index:0
// Public Visibility=Default Availability=Available
// [4] int revision() const

/*

 */
func (this *QMetaProperty) Revision() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty8revisionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmetaobject.h:277
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant read(const QObject *) const

/*

 */
func (this *QMetaProperty) Read(obj QObject_ITF /*777 const QObject **/) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty4readEPK7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:278
// index:0
// Public Visibility=Default Availability=Available
// [1] bool write(QObject *, const QVariant &) const

/*

 */
func (this *QMetaProperty) Write(obj QObject_ITF /*777 QObject **/, value QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty5writeEP7QObjectRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:279
// index:0
// Public Visibility=Default Availability=Available
// [1] bool reset(QObject *) const

/*

 */
func (this *QMetaProperty) Reset(obj QObject_ITF /*777 QObject **/) bool {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty5resetEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:281
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant readOnGadget(const void *) const

/*

 */
func (this *QMetaProperty) ReadOnGadget(gadget unsafe.Pointer /*666*/) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12readOnGadgetEPKv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qmetaobject.h:282
// index:0
// Public Visibility=Default Availability=Available
// [1] bool writeOnGadget(void *, const QVariant &) const

/*

 */
func (this *QMetaProperty) WriteOnGadget(gadget unsafe.Pointer /*666*/, value QVariant_ITF) bool {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty13writeOnGadgetEPvRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:283
// index:0
// Public Visibility=Default Availability=Available
// [1] bool resetOnGadget(void *) const

/*

 */
func (this *QMetaProperty) ResetOnGadget(gadget unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty13resetOnGadgetEPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), gadget)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:285
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasStdCppSet() const

/*

 */
func (this *QMetaProperty) HasStdCppSet() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty12hasStdCppSetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:286
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QMetaProperty) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmetaobject.h:287
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QMetaObject * enclosingMetaObject() const

/*

 */
func (this *QMetaProperty) EnclosingMetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QMetaProperty19enclosingMetaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

func DeleteQMetaProperty(this *QMetaProperty) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QMetaPropertyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

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
