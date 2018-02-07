package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h
// #include <qandroidparcel.h>
// #include <QtAndroidExtras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QAndroidParcel struct {
	*qtrt.CObject
}

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
func NewQAndroidParcel() *QAndroidParcel {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidParcelC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidParcelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidParcel)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:55
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QAndroidParcel(const QAndroidJniObject &)
func NewQAndroidParcel_1(parcel *QAndroidJniObject) *QAndroidParcel {
	var convArg0 = parcel.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidParcelC2ERK17QAndroidJniObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidParcelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAndroidParcel)
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAndroidParcel()
func DeleteQAndroidParcel(this *QAndroidParcel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QAndroidParcelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 1)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeData(const QByteArray &)
func (this *QAndroidParcel) WriteData(data *qtcore.QByteArray) {
	var convArg0 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel9writeDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeVariant(const QVariant &)
func (this *QAndroidParcel) WriteVariant(value *qtcore.QVariant) {
	var convArg0 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel12writeVariantERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeBinder(const QAndroidBinder &)
func (this *QAndroidParcel) WriteBinder(binder *QAndroidBinder) {
	var convArg0 = binder.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel11writeBinderERK14QAndroidBinder", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void writeFileDescriptor(int)
func (this *QAndroidParcel) WriteFileDescriptor(fd int) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel19writeFileDescriptorEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fd)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray readData()
func (this *QAndroidParcel) ReadData() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel8readDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:64
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant readVariant()
func (this *QAndroidParcel) ReadVariant() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel11readVariantEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] QAndroidBinder readBinder()
func (this *QAndroidParcel) ReadBinder() *QAndroidBinder /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel10readBinderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAndroidBinderFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidBinder)
	return rv2
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] int readFileDescriptor()
func (this *QAndroidParcel) ReadFileDescriptor() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel18readFileDescriptorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtAndroidExtras/qandroidparcel.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] QAndroidJniObject handle()
func (this *QAndroidParcel) Handle() *QAndroidJniObject /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QAndroidParcel6handleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAndroidJniObject)
	return rv2
}

//  body block end
