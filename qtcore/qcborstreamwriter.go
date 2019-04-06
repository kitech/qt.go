package qtcore

// /usr/include/qt/QtCore/qcborstream.h
// #include <qcborstream.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 57
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
type QCborStreamWriter struct {
	*qtrt.CObject
}
type QCborStreamWriter_ITF interface {
	QCborStreamWriter_PTR() *QCborStreamWriter
}

func (ptr *QCborStreamWriter) QCborStreamWriter_PTR() *QCborStreamWriter { return ptr }

func (this *QCborStreamWriter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCborStreamWriter) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCborStreamWriterFromPointer(cthis unsafe.Pointer) *QCborStreamWriter {
	return &QCborStreamWriter{&qtrt.CObject{cthis}}
}
func (*QCborStreamWriter) NewFromPointer(cthis unsafe.Pointer) *QCborStreamWriter {
	return NewQCborStreamWriterFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcborstream.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCborStreamWriter(QIODevice *)

/*

 */
func (*QCborStreamWriter) NewForInherit(device QIODevice_ITF /*777 QIODevice **/) *QCborStreamWriter {
	return NewQCborStreamWriter(device)
}
func NewQCborStreamWriter(device QIODevice_ITF /*777 QIODevice **/) *QCborStreamWriter {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriterC2EP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborStreamWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborStreamWriter)
	return gothis
}

// /usr/include/qt/QtCore/qcborstream.h:67
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCborStreamWriter(QByteArray *)

/*

 */
func (*QCborStreamWriter) NewForInherit1(data QByteArray_ITF /*777 QByteArray **/) *QCborStreamWriter {
	return NewQCborStreamWriter1(data)
}
func NewQCborStreamWriter1(data QByteArray_ITF /*777 QByteArray **/) *QCborStreamWriter {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriterC2EP10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCborStreamWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCborStreamWriter)
	return gothis
}

// /usr/include/qt/QtCore/qcborstream.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCborStreamWriter()

/*

 */
func DeleteQCborStreamWriter(this *QCborStreamWriter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcborstream.h:69
// index:0
// Public inline Visibility=Default Availability=NotAvailable
// [8] QCborStreamWriter & operator=(const QCborStreamWriter &)

/*

 */
func (this *QCborStreamWriter) Operator_equal(arg0 QCborStreamWriter_ITF) *QCborStreamWriter {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QCborStreamWriter_PTR() != nil {
		convArg0 = arg0.QCborStreamWriter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriteraSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCborStreamWriterFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCborStreamWriter)
	return rv2
}

// /usr/include/qt/QtCore/qcborstream.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)

/*

 */
func (this *QCborStreamWriter) SetDevice(device QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if device != nil && device.QIODevice_PTR() != nil {
		convArg0 = device.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device() const

/*

 */
func (this *QCborStreamWriter) Device() *QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QCborStreamWriter6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qcborstream.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void appendByteString(const char *, qsizetype)

/*

 */
func (this *QCborStreamWriter) AppendByteString(data string, len_ int64) {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter16appendByteStringEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void appendTextString(const char *, qsizetype)

/*

 */
func (this *QCborStreamWriter) AppendTextString(utf8 string, len_ int64) {
	var convArg0 = qtrt.CString(utf8)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter16appendTextStringEPKcx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, len_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void appendNull()

/*

 */
func (this *QCborStreamWriter) AppendNull() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter10appendNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:94
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void appendUndefined()

/*

 */
func (this *QCborStreamWriter) AppendUndefined() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter15appendUndefinedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startArray()

/*

 */
func (this *QCborStreamWriter) StartArray() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter10startArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:107
// index:1
// Public Visibility=Default Availability=Available
// [-2] void startArray(quint64)

/*

 */
func (this *QCborStreamWriter) StartArray1(count uint64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter10startArrayEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:108
// index:0
// Public Visibility=Default Availability=Available
// [1] bool endArray()

/*

 */
func (this *QCborStreamWriter) EndArray() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter8endArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcborstream.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startMap()

/*

 */
func (this *QCborStreamWriter) StartMap() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter8startMapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:110
// index:1
// Public Visibility=Default Availability=Available
// [-2] void startMap(quint64)

/*

 */
func (this *QCborStreamWriter) StartMap1(count uint64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter8startMapEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcborstream.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool endMap()

/*

 */
func (this *QCborStreamWriter) EndMap() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QCborStreamWriter6endMapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init_unused_10359() {
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
