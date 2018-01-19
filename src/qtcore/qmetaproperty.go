//  header block begin
// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QMetaProperty struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qmetaobject.h:248
// index:0
// void QMetaProperty()
func NewQMetaProperty() *QMetaProperty {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QMetaPropertyC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QMetaProperty{cthis}
}

// /usr/include/qt/QtCore/qmetaobject.h:250
// index:0
// const char * name()
func (this *QMetaProperty) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:251
// index:0
// const char * typeName()
func (this *QMetaProperty) TypeName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty8typeNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:252
// index:0
// QVariant::Type type()
func (this *QMetaProperty) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:253
// index:0
// int userType()
func (this *QMetaProperty) UserType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty8userTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:254
// index:0
// int propertyIndex()
func (this *QMetaProperty) PropertyIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty13propertyIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:256
// index:0
// bool isReadable()
func (this *QMetaProperty) IsReadable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty10isReadableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:257
// index:0
// bool isWritable()
func (this *QMetaProperty) IsWritable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty10isWritableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:258
// index:0
// bool isResettable()
func (this *QMetaProperty) IsResettable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty12isResettableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:259
// index:0
// bool isDesignable(const class QObject *)
func (this *QMetaProperty) IsDesignable(obj unsafe.Pointer) {
	// 0: (, const QObject * obj), (obj)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty12isDesignableEPK7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, obj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:260
// index:0
// bool isScriptable(const class QObject *)
func (this *QMetaProperty) IsScriptable(obj unsafe.Pointer) {
	// 0: (, const QObject * obj), (obj)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty12isScriptableEPK7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, obj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:261
// index:0
// bool isStored(const class QObject *)
func (this *QMetaProperty) IsStored(obj unsafe.Pointer) {
	// 0: (, const QObject * obj), (obj)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty8isStoredEPK7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, obj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:262
// index:0
// bool isEditable(const class QObject *)
func (this *QMetaProperty) IsEditable(obj unsafe.Pointer) {
	// 0: (, const QObject * obj), (obj)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty10isEditableEPK7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, obj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:263
// index:0
// bool isUser(const class QObject *)
func (this *QMetaProperty) IsUser(obj unsafe.Pointer) {
	// 0: (, const QObject * obj), (obj)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty6isUserEPK7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, obj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:264
// index:0
// bool isConstant()
func (this *QMetaProperty) IsConstant() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty10isConstantEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:265
// index:0
// bool isFinal()
func (this *QMetaProperty) IsFinal() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty7isFinalEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:267
// index:0
// bool isFlagType()
func (this *QMetaProperty) IsFlagType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty10isFlagTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:268
// index:0
// bool isEnumType()
func (this *QMetaProperty) IsEnumType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty10isEnumTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:269
// index:0
// QMetaEnum enumerator()
func (this *QMetaProperty) Enumerator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty10enumeratorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:271
// index:0
// bool hasNotifySignal()
func (this *QMetaProperty) HasNotifySignal() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty15hasNotifySignalEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:272
// index:0
// QMetaMethod notifySignal()
func (this *QMetaProperty) NotifySignal() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty12notifySignalEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:273
// index:0
// int notifySignalIndex()
func (this *QMetaProperty) NotifySignalIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty17notifySignalIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:275
// index:0
// int revision()
func (this *QMetaProperty) Revision() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty8revisionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:277
// index:0
// QVariant read(const class QObject *)
func (this *QMetaProperty) Read(obj unsafe.Pointer) {
	// 0: (, const QObject * obj), (obj)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty4readEPK7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, obj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:278
// index:0
// bool write(class QObject *, const class QVariant &)
func (this *QMetaProperty) Write(obj unsafe.Pointer, value unsafe.Pointer) {
	// 0: (, QObject * obj, const QVariant & value), (obj, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty5writeEP7QObjectRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, obj, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:279
// index:0
// bool reset(class QObject *)
func (this *QMetaProperty) Reset(obj unsafe.Pointer) {
	// 0: (, QObject * obj), (obj)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty5resetEP7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, obj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:281
// index:0
// QVariant readOnGadget(const void *)
func (this *QMetaProperty) ReadOnGadget(gadget unsafe.Pointer) {
	// 0: (, const void * gadget), (gadget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty12readOnGadgetEPKv", ffiqt.FFI_TYPE_VOID, this.cthis, gadget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:282
// index:0
// bool writeOnGadget(void *, const class QVariant &)
func (this *QMetaProperty) WriteOnGadget(gadget unsafe.Pointer, value unsafe.Pointer) {
	// 0: (, void * gadget, const QVariant & value), (gadget, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty13writeOnGadgetEPvRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, gadget, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:283
// index:0
// bool resetOnGadget(void *)
func (this *QMetaProperty) ResetOnGadget(gadget unsafe.Pointer) {
	// 0: (, void * gadget), (gadget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty13resetOnGadgetEPv", ffiqt.FFI_TYPE_VOID, this.cthis, gadget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:285
// index:0
// bool hasStdCppSet()
func (this *QMetaProperty) HasStdCppSet() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty12hasStdCppSetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:286
// index:0
// inline
// bool isValid()
func (this *QMetaProperty) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:287
// index:0
// inline
// const QMetaObject * enclosingMetaObject()
func (this *QMetaProperty) EnclosingMetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QMetaProperty19enclosingMetaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
