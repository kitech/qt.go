package qtcore

// /usr/include/qt/QtCore/qstringview.h
// #include <qstringview.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QStringView struct {
	*qtrt.CObject
}
type QStringView_ITF interface {
	QStringView_PTR() *QStringView
}

func (ptr *QStringView) QStringView_PTR() *QStringView { return ptr }

func (this *QStringView) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStringView) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStringViewFromPointer(cthis unsafe.Pointer) *QStringView {
	return &QStringView{&qtrt.CObject{cthis}}
}
func (*QStringView) NewFromPointer(cthis unsafe.Pointer) *QStringView {
	return NewQStringViewFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstringview.h:171
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QStringView()
func NewQStringView() *QStringView {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringViewC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStringView)
	return gothis
}

// /usr/include/qt/QtCore/qstringview.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toString()
func (this *QStringView) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qstringview.h:216
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qsizetype size()
func (this *QStringView) Size() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qstringview.h:217
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringView::const_pointer data()
func (this *QStringView) Data() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:218
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QStringView::storage_type * utf16()
func (this *QStringView) Utf16() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5utf16Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qstringview.h:220
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar operator[](qsizetype)
func (this *QStringView) Operator_get_index(n int64) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringViewixEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:227
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLatin1()
func (this *QStringView) ToLatin1() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8toLatin1Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:228
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toUtf8()
func (this *QStringView) ToUtf8() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView6toUtf8Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:229
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QByteArray toLocal8Bit()
func (this *QStringView) ToLocal8Bit() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView11toLocal8BitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar at(qsizetype)
func (this *QStringView) At(n int64) *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView2atEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:234
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringView mid(qsizetype)
func (this *QStringView) Mid(pos int64) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView3midEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:236
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QStringView mid(qsizetype, qsizetype)
func (this *QStringView) Mid_1(pos int64, n int64) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView3midExx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringView left(qsizetype)
func (this *QStringView) Left(n int64) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4leftEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:240
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringView right(qsizetype)
func (this *QStringView) Right(n int64) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5rightEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:242
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringView chopped(qsizetype)
func (this *QStringView) Chopped(n int64) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView7choppedEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:245
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void truncate(qsizetype)
func (this *QStringView) Truncate(n int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringView8truncateEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:247
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void chop(qsizetype)
func (this *QStringView) Chop(n int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringView4chopEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), n)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringview.h:250
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringView trimmed()
func (this *QStringView) Trimmed() *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView7trimmedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:252
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QStringView, Qt::CaseSensitivity)
func (this *QStringView) StartsWith(s QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView10startsWithES_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:254
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QLatin1String, Qt::CaseSensitivity)
func (this *QStringView) StartsWith_1(s QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView10startsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:255
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QChar)
func (this *QStringView) StartsWith_2(c QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView10startsWithE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:257
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool startsWith(QChar, Qt::CaseSensitivity)
func (this *QStringView) StartsWith_3(c QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView10startsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:260
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QStringView, Qt::CaseSensitivity)
func (this *QStringView) EndsWith(s QStringView_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QStringView_PTR() != nil {
		convArg0 = s.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8endsWithES_N2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:262
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QLatin1String, Qt::CaseSensitivity)
func (this *QStringView) EndsWith_1(s QLatin1String_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QLatin1String_PTR() != nil {
		convArg0 = s.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8endsWithE13QLatin1StringN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:263
// index:2
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QChar)
func (this *QStringView) EndsWith_2(c QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8endsWithE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:265
// index:3
// Public inline Visibility=Default Availability=Available
// [1] bool endsWith(QChar, Qt::CaseSensitivity)
func (this *QStringView) EndsWith_3(c QChar_ITF /*123*/, cs int) bool {
	var convArg0 unsafe.Pointer
	if c != nil && c.QChar_PTR() != nil {
		convArg0 = c.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView8endsWithE5QCharN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, cs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:271
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringView::const_iterator begin()
func (this *QStringView) Begin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:272
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringView::const_iterator end()
func (this *QStringView) End() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:273
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringView::const_iterator cbegin()
func (this *QStringView) Cbegin() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView6cbeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:274
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStringView::const_iterator cend()
func (this *QStringView) Cend() *QChar /*777 const QChar **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4cendEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) //555
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:280
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool empty()
func (this *QStringView) Empty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5emptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:281
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar front()
func (this *QStringView) Front() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5frontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:282
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar back()
func (this *QStringView) Back() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4backEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:287
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QStringView) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:288
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QStringView) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringview.h:289
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int length()
func (this *QStringView) Length() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView6lengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringview.h:291
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar first()
func (this *QStringView) First() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView5firstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

// /usr/include/qt/QtCore/qstringview.h:292
// index:0
// Public inline Visibility=Default Availability=Available
// [2] QChar last()
func (this *QStringView) Last() *QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStringView4lastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQChar)
	return rv2
}

func DeleteQStringView(this *QStringView) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStringViewD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
}

//  keep block end
