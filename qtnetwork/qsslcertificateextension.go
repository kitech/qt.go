package qtnetwork

// /usr/include/qt/QtNetwork/qsslcertificateextension.h
// #include <qsslcertificateextension.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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

type QSslCertificateExtension struct {
	*qtrt.CObject
}

func (this *QSslCertificateExtension) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslCertificateExtension) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslCertificateExtensionFromPointer(cthis unsafe.Pointer) *QSslCertificateExtension {
	return &QSslCertificateExtension{&qtrt.CObject{cthis}}
}
func (*QSslCertificateExtension) NewFromPointer(cthis unsafe.Pointer) *QSslCertificateExtension {
	return NewQSslCertificateExtensionFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslCertificateExtension()
func NewQSslCertificateExtension() *QSslCertificateExtension {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSslCertificateExtensionC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslCertificateExtensionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCertificateExtension)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslCertificateExtension()
func DeleteQSslCertificateExtension(this *QSslCertificateExtension) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSslCertificateExtensionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslCertificateExtension &)
func (this *QSslCertificateExtension) Swap(other *QSslCertificateExtension) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSslCertificateExtension4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QString oid()
func (this *QSslCertificateExtension) Oid() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSslCertificateExtension3oidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QSslCertificateExtension) Name() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSslCertificateExtension4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:71
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value()
func (this *QSslCertificateExtension) Value() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSslCertificateExtension5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCritical()
func (this *QSslCertificateExtension) IsCritical() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSslCertificateExtension10isCriticalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSupported()
func (this *QSslCertificateExtension) IsSupported() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSslCertificateExtension11isSupportedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
