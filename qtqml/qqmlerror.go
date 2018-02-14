package qtqml

// /usr/include/qt/QtQml/qqmlerror.h
// #include <qqmlerror.h>
// #include <QtQml>

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
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QQmlError struct {
	*qtrt.CObject
}
type QQmlError_ITF interface {
	QQmlError_PTR() *QQmlError
}

func (ptr *QQmlError) QQmlError_PTR() *QQmlError { return ptr }

func (this *QQmlError) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlError) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlErrorFromPointer(cthis unsafe.Pointer) *QQmlError {
	return &QQmlError{&qtrt.CObject{cthis}}
}
func (*QQmlError) NewFromPointer(cthis unsafe.Pointer) *QQmlError {
	return NewQQmlErrorFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlerror.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlError()
func NewQQmlError() *QQmlError {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QQmlErrorC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlErrorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQQmlError)
	return gothis
}

// /usr/include/qt/QtQml/qqmlerror.h:58
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlError & operator=(const QQmlError &)
func (this *QQmlError) Operator_equal(arg0 QQmlError_ITF) *QQmlError {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQmlError_PTR() != nil {
		convArg0 = arg0.QQmlError_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QQmlErroraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlErrorFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlError)
	return rv2
}

// /usr/include/qt/QtQml/qqmlerror.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QQmlError()
func DeleteQQmlError(this *QQmlError) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QQmlErrorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlerror.h:61
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QQmlError) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QQmlError7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlerror.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url()
func (this *QQmlError) Url() *qtcore.QUrl /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QQmlError3urlEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtQml/qqmlerror.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QUrl &)
func (this *QQmlError) SetUrl(arg0 qtcore.QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QUrl_PTR() != nil {
		convArg0 = arg0.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QQmlError6setUrlERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description()
func (this *QQmlError) Description() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QQmlError11descriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQml/qqmlerror.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescription(const QString &)
func (this *QQmlError) SetDescription(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QQmlError14setDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int line()
func (this *QQmlError) Line() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QQmlError4lineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlerror.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLine(int)
func (this *QQmlError) SetLine(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QQmlError7setLineEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] int column()
func (this *QQmlError) Column() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QQmlError6columnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlerror.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumn(int)
func (this *QQmlError) SetColumn(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QQmlError9setColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * object()
func (this *QQmlError) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QQmlError6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlerror.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObject(QObject *)
func (this *QQmlError) SetObject(arg0 qtcore.QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QQmlError9setObjectEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] QtMsgType messageType()
func (this *QQmlError) MessageType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QQmlError11messageTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMessageType(enum QtMsgType)
func (this *QQmlError) SetMessageType(messageType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QQmlError14setMessageTypeE9QtMsgType", qtrt.FFI_TYPE_POINTER, this.GetCthis(), messageType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString()
func (this *QQmlError) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QQmlError8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
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
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
