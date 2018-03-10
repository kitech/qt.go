package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidintent.h
// #include <qandroidintent.h>
// #include <QtAndroidExtras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QAndroidIntent struct {
	*qtrt.CObject
}
type QAndroidIntent_ITF interface {
	QAndroidIntent_PTR() *QAndroidIntent
}

func (ptr *QAndroidIntent) QAndroidIntent_PTR() *QAndroidIntent { return ptr }

func (this *QAndroidIntent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidIntent) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidIntentFromPointer(cthis unsafe.Pointer) *QAndroidIntent {
	return &QAndroidIntent{&qtrt.CObject{cthis}}
}
func (*QAndroidIntent) NewFromPointer(cthis unsafe.Pointer) *QAndroidIntent {
	return NewQAndroidIntentFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/qandroidintent.h:51
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidIntent()

/*
Create a new intent
*/
func NewQAndroidIntent() *QAndroidIntent {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidIntentC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidIntentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidIntent)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidintent.h:53
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAndroidIntent(const QAndroidJniObject &)

/*
Create a new intent
*/
func NewQAndroidIntent_1(intent QAndroidJniObject_ITF) *QAndroidIntent {
	var convArg0 unsafe.Pointer
	if intent != nil && intent.QAndroidJniObject_PTR() != nil {
		convArg0 = intent.QAndroidJniObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidIntentC2ERK17QAndroidJniObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidIntentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidIntent)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidintent.h:54
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QAndroidIntent(const QString &)

/*
Create a new intent
*/
func NewQAndroidIntent_2(action string) *QAndroidIntent {
	var tmpArg0 = qtcore.NewQString_5(action)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidIntentC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidIntentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidIntent)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidintent.h:55
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QAndroidIntent(const QAndroidJniObject &, const char *)

/*
Create a new intent
*/
func NewQAndroidIntent_3(packageContext QAndroidJniObject_ITF, className string) *QAndroidIntent {
	var convArg0 unsafe.Pointer
	if packageContext != nil && packageContext.QAndroidJniObject_PTR() != nil {
		convArg0 = packageContext.QAndroidJniObject_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(className)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidIntentC2ERK17QAndroidJniObjectPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidIntentFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidIntent)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidintent.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAndroidIntent()

/*

 */
func DeleteQAndroidIntent(this *QAndroidIntent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidIntentD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidintent.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void putExtra(const QString &, const QByteArray &)

/*
Sets the key with the data in the Intent extras
*/
func (this *QAndroidIntent) PutExtra(key string, data qtcore.QByteArray_ITF) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg1 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidIntent8putExtraERK7QStringRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidintent.h:60
// index:1
// Public Visibility=Default Availability=Available
// [-2] void putExtra(const QString &, const QVariant &)

/*
Sets the key with the data in the Intent extras
*/
func (this *QAndroidIntent) PutExtra_1(key string, value qtcore.QVariant_ITF) {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidIntent8putExtraERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidintent.h:58
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray extraBytes(const QString &)

/*
Returns the extra key data from the Intent extras
*/
func (this *QAndroidIntent) ExtraBytes(key string) *qtcore.QByteArray /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidIntent10extraBytesERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidintent.h:61
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant extraVariant(const QString &)

/*
Returns the extra key data from the Intent extras as a QVariant
*/
func (this *QAndroidIntent) ExtraVariant(key string) *qtcore.QVariant /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidIntent12extraVariantERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidintent.h:63
// index:0
// Public Visibility=Default Availability=Available
// [16] QAndroidJniObject handle() const

/*
The return value is useful to call other Java API which are not covered by this wrapper
*/
func (this *QAndroidIntent) Handle() *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidIntent6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
