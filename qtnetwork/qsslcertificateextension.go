package qtnetwork

// /usr/include/qt/QtNetwork/qsslcertificateextension.h
// #include <qsslcertificateextension.h>
// #include <QtNetwork>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QSslCertificateExtension struct {
	*qtrt.CObject
}
type QSslCertificateExtension_ITF interface {
	QSslCertificateExtension_PTR() *QSslCertificateExtension
}

func (ptr *QSslCertificateExtension) QSslCertificateExtension_PTR() *QSslCertificateExtension {
	return ptr
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

/*
Constructs a QSslCertificateExtension.
*/
func (*QSslCertificateExtension) NewForInherit() *QSslCertificateExtension {
	return NewQSslCertificateExtension()
}
func NewQSslCertificateExtension() *QSslCertificateExtension {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSslCertificateExtensionC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSslCertificateExtensionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCertificateExtension)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSslCertificateExtension & operator=(QSslCertificateExtension &&)

/*

 */
func (this *QSslCertificateExtension) Operator_equal(other unsafe.Pointer /*333*/) *QSslCertificateExtension {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSslCertificateExtensionaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateExtensionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificateExtension)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:64
// index:1
// Public Visibility=Default Availability=Available
// [8] QSslCertificateExtension & operator=(const QSslCertificateExtension &)

/*

 */
func (this *QSslCertificateExtension) Operator_equal_1(other QSslCertificateExtension_ITF) *QSslCertificateExtension {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslCertificateExtension_PTR() != nil {
		convArg0 = other.QSslCertificateExtension_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSslCertificateExtensionaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSslCertificateExtensionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSslCertificateExtension)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslCertificateExtension()

/*

 */
func DeleteQSslCertificateExtension(this *QSslCertificateExtension) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSslCertificateExtensionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslCertificateExtension &)

/*
Swaps this certificate extension instance with other. This function is very fast and never fails.
*/
func (this *QSslCertificateExtension) Swap(other QSslCertificateExtension_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSslCertificateExtension_PTR() != nil {
		convArg0 = other.QSslCertificateExtension_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN24QSslCertificateExtension4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QString oid() const

/*
Returns the ASN.1 OID of this extension.
*/
func (this *QSslCertificateExtension) Oid() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSslCertificateExtension3oidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns the name of the extension. If no name is known for the extension then the OID will be returned.
*/
func (this *QSslCertificateExtension) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSslCertificateExtension4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:71
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value() const

/*
Returns the value of the extension. The structure of the value returned depends on the extension type.
*/
func (this *QSslCertificateExtension) Value() *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSslCertificateExtension5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCritical() const

/*
Returns the criticality of the extension.
*/
func (this *QSslCertificateExtension) IsCritical() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSslCertificateExtension10isCriticalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcertificateextension.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSupported() const

/*
Returns the true if this extension is supported. In this case, supported simply means that the structure of the QVariant returned by the value() accessor will remain unchanged between versions. Unsupported extensions can be freely used, however there is no guarantee that the returned data will have the same structure between versions.
*/
func (this *QSslCertificateExtension) IsSupported() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QSslCertificateExtension11isSupportedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
