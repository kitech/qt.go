package qtqml

// /usr/include/qt/QtQml/qqmlinfo.h
// #include <qqmlinfo.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

/*

 */
type QQmlInfo struct {
	*qtrt.CObject
}
type QQmlInfo_ITF interface {
	QQmlInfo_PTR() *QQmlInfo
}

func (ptr *QQmlInfo) QQmlInfo_PTR() *QQmlInfo { return ptr }

func (this *QQmlInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlInfo) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlInfoFromPointer(cthis unsafe.Pointer) *QQmlInfo {
	return &QQmlInfo{&qtrt.CObject{cthis}}
}
func (*QQmlInfo) NewFromPointer(cthis unsafe.Pointer) *QQmlInfo {
	return NewQQmlInfoFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlinfo.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QQmlInfo()

/*

 */
func DeleteQQmlInfo(this *QQmlInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlinfo.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(QChar)

/*

 */
func (this *QQmlInfo) Operator_left_shift(t qtcore.QChar_ITF /*123*/) *QQmlInfo {
	var convArg0 unsafe.Pointer
	if t != nil && t.QChar_PTR() != nil {
		convArg0 = t.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:79
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(bool)

/*

 */
func (this *QQmlInfo) Operator_left_shift1(t bool) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:80
// index:2
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(char)

/*

 */
func (this *QQmlInfo) Operator_left_shift2(t byte) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:81
// index:3
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(short)

/*

 */
func (this *QQmlInfo) Operator_left_shift3(t int16) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEs", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:82
// index:4
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(unsigned short)

/*

 */
func (this *QQmlInfo) Operator_left_shift4(t uint16) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEt", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:83
// index:5
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(int)

/*

 */
func (this *QQmlInfo) Operator_left_shift5(t int) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:84
// index:6
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(unsigned int)

/*

 */
func (this *QQmlInfo) Operator_left_shift6(t uint) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:85
// index:7
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(long)

/*

 */
func (this *QQmlInfo) Operator_left_shift7(t int) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:86
// index:8
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(unsigned long)

/*

 */
func (this *QQmlInfo) Operator_left_shift8(t uint) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEm", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:87
// index:9
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(qint64)

/*

 */
func (this *QQmlInfo) Operator_left_shift9(t int64) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:88
// index:10
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(quint64)

/*

 */
func (this *QQmlInfo) Operator_left_shift10(t uint64) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:89
// index:11
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(float)

/*

 */
func (this *QQmlInfo) Operator_left_shift11(t float32) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:90
// index:12
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(double)

/*

 */
func (this *QQmlInfo) Operator_left_shift12(t float64) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:91
// index:13
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(const char *)

/*

 */
func (this *QQmlInfo) Operator_left_shift13(t string) *QQmlInfo {
	var convArg0 = qtrt.CString(t)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:92
// index:14
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(const QString &)

/*

 */
func (this *QQmlInfo) Operator_left_shift14(t string) *QQmlInfo {
	var tmpArg0 = qtcore.NewQString5(t)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:93
// index:15
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(const QStringRef &)

/*

 */
func (this *QQmlInfo) Operator_left_shift15(t qtcore.QStringRef_ITF) *QQmlInfo {
	var convArg0 unsafe.Pointer
	if t != nil && t.QStringRef_PTR() != nil {
		convArg0 = t.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsERK10QStringRef", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:94
// index:16
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(const QLatin1String &)

/*

 */
func (this *QQmlInfo) Operator_left_shift16(t qtcore.QLatin1String_ITF) *QQmlInfo {
	var convArg0 unsafe.Pointer
	if t != nil && t.QLatin1String_PTR() != nil {
		convArg0 = t.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsERK13QLatin1String", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:95
// index:17
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(const QByteArray &)

/*

 */
func (this *QQmlInfo) Operator_left_shift17(t qtcore.QByteArray_ITF) *QQmlInfo {
	var convArg0 unsafe.Pointer
	if t != nil && t.QByteArray_PTR() != nil {
		convArg0 = t.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:96
// index:18
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(const void *)

/*

 */
func (this *QQmlInfo) Operator_left_shift18(t unsafe.Pointer /*666*/) *QQmlInfo {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsEPKv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), t)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:98
// index:19
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(QTextStreamManipulator)

/*

 */
func (this *QQmlInfo) Operator_left_shift19(m qtcore.QTextStreamManipulator_ITF /*123*/) *QQmlInfo {
	var convArg0 unsafe.Pointer
	if m != nil && m.QTextStreamManipulator_PTR() != nil {
		convArg0 = m.QTextStreamManipulator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsE22QTextStreamManipulator", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

// /usr/include/qt/QtQml/qqmlinfo.h:100
// index:20
// Public inline Visibility=Default Availability=Available
// [16] QQmlInfo & operator<<(const QUrl &)

/*

 */
func (this *QQmlInfo) Operator_left_shift20(t qtcore.QUrl_ITF) *QQmlInfo {
	var convArg0 unsafe.Pointer
	if t != nil && t.QUrl_PTR() != nil {
		convArg0 = t.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QQmlInfolsERK4QUrl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQmlInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQmlInfo)
	return rv2
}

//  body block end

//  keep block begin

func init_unused_11537() {
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
