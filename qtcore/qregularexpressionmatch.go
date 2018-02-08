package qtcore

// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type QRegularExpressionMatch struct {
	*qtrt.CObject
}

func (this *QRegularExpressionMatch) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRegularExpressionMatch) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRegularExpressionMatchFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatch {
	return &QRegularExpressionMatch{&qtrt.CObject{cthis}}
}
func (*QRegularExpressionMatch) NewFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatch {
	return NewQRegularExpressionMatchFromPointer(cthis)
}

// /usr/include/qt/QtCore/qregularexpression.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRegularExpressionMatch()
func NewQRegularExpressionMatch() *QRegularExpressionMatch {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QRegularExpressionMatchC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionMatchFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRegularExpressionMatch)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRegularExpressionMatch()
func DeleteQRegularExpressionMatch(this *QRegularExpressionMatch) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QRegularExpressionMatchD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qregularexpression.h:187
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRegularExpressionMatch &)
func (this *QRegularExpressionMatch) Swap(other *QRegularExpressionMatch) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QRegularExpressionMatch4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:189
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegularExpression regularExpression()
func (this *QRegularExpressionMatch) RegularExpression() *QRegularExpression /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch17regularExpressionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegularExpression)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:190
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegularExpression::MatchType matchType()
func (this *QRegularExpressionMatch) MatchType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch9matchTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:191
// index:0
// Public Visibility=Default Availability=Available
// [4] QRegularExpression::MatchOptions matchOptions()
func (this *QRegularExpressionMatch) MatchOptions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12matchOptionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:193
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasMatch()
func (this *QRegularExpressionMatch) HasMatch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8hasMatchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:194
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasPartialMatch()
func (this *QRegularExpressionMatch) HasPartialMatch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch15hasPartialMatchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:196
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QRegularExpressionMatch) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:198
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastCapturedIndex()
func (this *QRegularExpressionMatch) LastCapturedIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch17lastCapturedIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QString captured(int)
func (this *QRegularExpressionMatch) Captured(nth int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregularexpression.h:205
// index:1
// Public Visibility=Default Availability=Available
// [8] QString captured(const QString &)
func (this *QRegularExpressionMatch) Captured_1(name string) string {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregularexpression.h:209
// index:2
// Public Visibility=Default Availability=Available
// [8] QString captured(QStringView)
func (this *QRegularExpressionMatch) Captured_2(name *QStringView /*123*/) string {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qregularexpression.h:201
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringRef capturedRef(int)
func (this *QRegularExpressionMatch) CapturedRef(nth int) *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:206
// index:1
// Public Visibility=Default Availability=Available
// [16] QStringRef capturedRef(const QString &)
func (this *QRegularExpressionMatch) CapturedRef_1(name string) *QStringRef /*123*/ {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:210
// index:2
// Public Visibility=Default Availability=Available
// [16] QStringRef capturedRef(QStringView)
func (this *QRegularExpressionMatch) CapturedRef_2(name *QStringView /*123*/) *QStringRef /*123*/ {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:202
// index:0
// Public Visibility=Default Availability=Available
// [16] QStringView capturedView(int)
func (this *QRegularExpressionMatch) CapturedView(nth int) *QStringView /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12capturedViewEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:211
// index:1
// Public Visibility=Default Availability=Available
// [16] QStringView capturedView(QStringView)
func (this *QRegularExpressionMatch) CapturedView_1(name *QStringView /*123*/) *QStringView /*123*/ {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12capturedViewE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringView)
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:215
// index:0
// Public Visibility=Default Availability=Available
// [4] int capturedStart(int)
func (this *QRegularExpressionMatch) CapturedStart(nth int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:220
// index:1
// Public Visibility=Default Availability=Available
// [4] int capturedStart(const QString &)
func (this *QRegularExpressionMatch) CapturedStart_1(name string) int {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:225
// index:2
// Public Visibility=Default Availability=Available
// [4] int capturedStart(QStringView)
func (this *QRegularExpressionMatch) CapturedStart_2(name *QStringView /*123*/) int {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:216
// index:0
// Public Visibility=Default Availability=Available
// [4] int capturedLength(int)
func (this *QRegularExpressionMatch) CapturedLength(nth int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:221
// index:1
// Public Visibility=Default Availability=Available
// [4] int capturedLength(const QString &)
func (this *QRegularExpressionMatch) CapturedLength_1(name string) int {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:226
// index:2
// Public Visibility=Default Availability=Available
// [4] int capturedLength(QStringView)
func (this *QRegularExpressionMatch) CapturedLength_2(name *QStringView /*123*/) int {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:217
// index:0
// Public Visibility=Default Availability=Available
// [4] int capturedEnd(int)
func (this *QRegularExpressionMatch) CapturedEnd(nth int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:222
// index:1
// Public Visibility=Default Availability=Available
// [4] int capturedEnd(const QString &)
func (this *QRegularExpressionMatch) CapturedEnd_1(name string) int {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qregularexpression.h:227
// index:2
// Public Visibility=Default Availability=Available
// [4] int capturedEnd(QStringView)
func (this *QRegularExpressionMatch) CapturedEnd_2(name *QStringView /*123*/) int {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
