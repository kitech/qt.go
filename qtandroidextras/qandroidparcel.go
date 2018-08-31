package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h
// #include <qandroidparcel.h>
// #include <QtAndroidExtras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QAndroidParcel struct {
	*qtrt.CObject
}
type QAndroidParcel_ITF interface {
	QAndroidParcel_PTR() *QAndroidParcel
}

func (ptr *QAndroidParcel) QAndroidParcel_PTR() *QAndroidParcel { return ptr }

func (this *QAndroidParcel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAndroidParcel) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAndroidParcelFromPointer(cthis unsafe.Pointer) *QAndroidParcel {
	return &QAndroidParcel{&qtrt.CObject{cthis}}
}
func (*QAndroidParcel) NewFromPointer(cthis unsafe.Pointer) *QAndroidParcel {
	return NewQAndroidParcelFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidParcel()

/*
Creates a new object.
*/
func (*QAndroidParcel) NewForInherit() *QAndroidParcel {
	return NewQAndroidParcel()
}
func NewQAndroidParcel() *QAndroidParcel {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidParcelC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidParcelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidParcel)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:55
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAndroidParcel(const QAndroidJniObject &)

/*
Creates a new object.
*/
func (*QAndroidParcel) NewForInherit_1(parcel QAndroidJniObject_ITF) *QAndroidParcel {
	return NewQAndroidParcel_1(parcel)
}
func NewQAndroidParcel_1(parcel QAndroidJniObject_ITF) *QAndroidParcel {
	var convArg0 unsafe.Pointer
	if parcel != nil && parcel.QAndroidJniObject_PTR() != nil {
		convArg0 = parcel.QAndroidJniObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidParcelC2ERK17QAndroidJniObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidParcelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidParcel)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAndroidParcel()

/*

 */
func DeleteQAndroidParcel(this *QAndroidParcel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidParcelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeData(const QByteArray &) const

/*
Writes the provided data as a byte array
*/
func (this *QAndroidParcel) WriteData(data qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel9writeDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeVariant(const QVariant &) const

/*
Writes the provided value. The value is converted into a QByteArray before is written.
*/
func (this *QAndroidParcel) WriteVariant(value qtcore.QVariant_ITF) {
	var convArg0 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg0 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel12writeVariantERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeBinder(const QAndroidBinder &) const

/*
Writes a binder object. This is useful for a client to send to a server a binder which can be used by the server callback the client.
*/
func (this *QAndroidParcel) WriteBinder(binder QAndroidBinder_ITF) {
	var convArg0 unsafe.Pointer
	if binder != nil && binder.QAndroidBinder_PTR() != nil {
		convArg0 = binder.QAndroidBinder_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel11writeBinderERK14QAndroidBinder", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeFileDescriptor(int) const

/*
Writes the provided fd.
*/
func (this *QAndroidParcel) WriteFileDescriptor(fd int) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel19writeFileDescriptorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fd)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray readData() const

/*
Returns the data as a QByteArray
*/
func (this *QAndroidParcel) ReadData() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel8readDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:64
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant readVariant() const

/*
Returns the data as a QVariant
*/
func (this *QAndroidParcel) ReadVariant() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel11readVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:65
// index:0
// Public Visibility=Default Availability=Available
// [24] QAndroidBinder readBinder() const

/*
Returns the binder as a QAndroidBinder
*/
func (this *QAndroidParcel) ReadBinder() *QAndroidBinder /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel10readBinderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAndroidBinderFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidBinder)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] int readFileDescriptor() const

/*
Returns the file descriptor
*/
func (this *QAndroidParcel) ReadFileDescriptor() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel18readFileDescriptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:68
// index:0
// Public Visibility=Default Availability=Available
// [16] QAndroidJniObject handle() const

/*
The return value is useful to call other Java API which are not covered by this wrapper
*/
func (this *QAndroidParcel) Handle() *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
