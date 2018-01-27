package qtcore

// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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
	this.CObject = &qtrt.CObject{cthis}
}
func NewQRegularExpressionMatchFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatch {
	return &QRegularExpressionMatch{&qtrt.CObject{cthis}}
}
func (*QRegularExpressionMatch) NewFromPointer(cthis unsafe.Pointer) *QRegularExpressionMatch {
	return NewQRegularExpressionMatchFromPointer(cthis)
}

// /usr/include/qt/QtCore/qregularexpression.h:178
// index:0
// Public
// void QRegularExpressionMatch()
func NewQRegularExpressionMatch() *QRegularExpressionMatch {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QRegularExpressionMatchC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionMatchFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:179
// index:0
// Public
// void ~QRegularExpressionMatch()
func DeleteQRegularExpressionMatch(*QRegularExpressionMatch) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QRegularExpressionMatchD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:187
// index:0
// Public inline
// void swap(QRegularExpressionMatch &)
func (this *QRegularExpressionMatch) Swap(other *QRegularExpressionMatch) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QRegularExpressionMatch4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:189
// index:0
// Public
// QRegularExpression regularExpression()
func (this *QRegularExpressionMatch) RegularExpression() *QRegularExpression /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch17regularExpressionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegularExpressionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:190
// index:0
// Public
// QRegularExpression::MatchType matchType()
func (this *QRegularExpressionMatch) MatchType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch9matchTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:191
// index:0
// Public
// QRegularExpression::MatchOptions matchOptions()
func (this *QRegularExpressionMatch) MatchOptions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12matchOptionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:193
// index:0
// Public
// bool hasMatch()
func (this *QRegularExpressionMatch) HasMatch() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8hasMatchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:194
// index:0
// Public
// bool hasPartialMatch()
func (this *QRegularExpressionMatch) HasPartialMatch() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch15hasPartialMatchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:196
// index:0
// Public
// bool isValid()
func (this *QRegularExpressionMatch) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:198
// index:0
// Public
// int lastCapturedIndex()
func (this *QRegularExpressionMatch) LastCapturedIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch17lastCapturedIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:200
// index:0
// Public
// QString captured(int)
func (this *QRegularExpressionMatch) Captured(nth int) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:205
// index:1
// Public
// QString captured(const QString &)
func (this *QRegularExpressionMatch) Captured_1(name *QString) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:209
// index:2
// Public
// QString captured(QStringView)
func (this *QRegularExpressionMatch) Captured_2(name *QStringView /*123*/) *QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch8capturedE11QStringView", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:201
// index:0
// Public
// QStringRef capturedRef(int)
func (this *QRegularExpressionMatch) CapturedRef(nth int) *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:206
// index:1
// Public
// QStringRef capturedRef(const QString &)
func (this *QRegularExpressionMatch) CapturedRef_1(name *QString) *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefERK7QString", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:210
// index:2
// Public
// QStringRef capturedRef(QStringView)
func (this *QRegularExpressionMatch) CapturedRef_2(name *QStringView /*123*/) *QStringRef /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedRefE11QStringView", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:202
// index:0
// Public
// QStringView capturedView(int)
func (this *QRegularExpressionMatch) CapturedView(nth int) *QStringView /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12capturedViewEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:211
// index:1
// Public
// QStringView capturedView(QStringView)
func (this *QRegularExpressionMatch) CapturedView_1(name *QStringView /*123*/) *QStringView /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch12capturedViewE11QStringView", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQStringViewFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:215
// index:0
// Public
// int capturedStart(int)
func (this *QRegularExpressionMatch) CapturedStart(nth int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:220
// index:1
// Public
// int capturedStart(const QString &)
func (this *QRegularExpressionMatch) CapturedStart_1(name *QString) int {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:225
// index:2
// Public
// int capturedStart(QStringView)
func (this *QRegularExpressionMatch) CapturedStart_2(name *QStringView /*123*/) int {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch13capturedStartE11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:216
// index:0
// Public
// int capturedLength(int)
func (this *QRegularExpressionMatch) CapturedLength(nth int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:221
// index:1
// Public
// int capturedLength(const QString &)
func (this *QRegularExpressionMatch) CapturedLength_1(name *QString) int {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:226
// index:2
// Public
// int capturedLength(QStringView)
func (this *QRegularExpressionMatch) CapturedLength_2(name *QStringView /*123*/) int {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch14capturedLengthE11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:217
// index:0
// Public
// int capturedEnd(int)
func (this *QRegularExpressionMatch) CapturedEnd(nth int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), nth)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:222
// index:1
// Public
// int capturedEnd(const QString &)
func (this *QRegularExpressionMatch) CapturedEnd_1(name *QString) int {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:227
// index:2
// Public
// int capturedEnd(QStringView)
func (this *QRegularExpressionMatch) CapturedEnd_2(name *QStringView /*123*/) int {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QRegularExpressionMatch11capturedEndE11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
