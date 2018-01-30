package qtqml

// /usr/include/qt/QtQml/qqmlerror.h
// #include <qqmlerror.h>
// #include <QtQml>

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
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQmlError struct {
	*qtrt.CObject
}

func (this *QQmlError) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlError) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QQmlErrorC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQQmlErrorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlerror.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QQmlError()
func DeleteQQmlError(*QQmlError) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QQmlErrorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:61
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QQmlError) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QQmlError7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQml/qqmlerror.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl url()
func (this *QQmlError) Url() *qtcore.QUrl /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QQmlError3urlEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlerror.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QUrl &)
func (this *QQmlError) SetUrl(arg0 *qtcore.QUrl) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QQmlError6setUrlERK4QUrl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description()
func (this *QQmlError) Description() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QQmlError11descriptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQml/qqmlerror.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescription(const QString &)
func (this *QQmlError) SetDescription(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QQmlError14setDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int line()
func (this *QQmlError) Line() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QQmlError4lineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlerror.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLine(int)
func (this *QQmlError) SetLine(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QQmlError7setLineEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] int column()
func (this *QQmlError) Column() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QQmlError6columnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQml/qqmlerror.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumn(int)
func (this *QQmlError) SetColumn(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QQmlError9setColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QObject * object()
func (this *QQmlError) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QQmlError6objectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQml/qqmlerror.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObject(QObject *)
func (this *QQmlError) SetObject(arg0 *qtcore.QObject /*777 QObject **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QQmlError9setObjectEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:73
// index:0
// Public Visibility=Default Availability=Available
// [4] QtMsgType messageType()
func (this *QQmlError) MessageType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QQmlError11messageTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMessageType(enum QtMsgType)
func (this *QQmlError) SetMessageType(messageType int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QQmlError14setMessageTypeE9QtMsgType", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), messageType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlerror.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString()
func (this *QQmlError) ToString() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QQmlError8toStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
