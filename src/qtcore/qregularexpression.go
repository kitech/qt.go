package qtcore

// /usr/include/qt/QtCore/qregularexpression.h
// #include <qregularexpression.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 68
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
type QRegularExpression struct {
	*qtrt.CObject
}

func (this *QRegularExpression) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQRegularExpressionFromPointer(cthis unsafe.Pointer) *QRegularExpression {
	return &QRegularExpression{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qregularexpression.h:81
// index:0
// Public
// QRegularExpression::PatternOptions patternOptions()
func (this *QRegularExpression) PatternOptions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression14patternOptionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:84
// index:0
// Public
// void QRegularExpression()
func NewQRegularExpression() *QRegularExpression {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpressionC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRegularExpressionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qregularexpression.h:87
// index:0
// Public
// void ~QRegularExpression()
func DeleteQRegularExpression(*QRegularExpression) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpressionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:95
// index:0
// Public inline
// void swap(class QRegularExpression &)
func (this *QRegularExpression) Swap(other *QRegularExpression) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:97
// index:0
// Public
// QString pattern()
func (this *QRegularExpression) Pattern() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression7patternEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:98
// index:0
// Public
// void setPattern(const class QString &)
func (this *QRegularExpression) SetPattern(pattern *QString) {
	var convArg0 = pattern.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression10setPatternERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:100
// index:0
// Public
// bool isValid()
func (this *QRegularExpression) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qregularexpression.h:101
// index:0
// Public
// int patternErrorOffset()
func (this *QRegularExpression) PatternErrorOffset() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression18patternErrorOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:102
// index:0
// Public
// QString errorString()
func (this *QRegularExpression) ErrorString() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtCore/qregularexpression.h:104
// index:0
// Public
// int captureCount()
func (this *QRegularExpression) CaptureCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression12captureCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qregularexpression.h:141
// index:0
// Public
// void optimize()
func (this *QRegularExpression) Optimize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QRegularExpression8optimizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qregularexpression.h:143
// index:0
// Public static
// QString escape(const class QString &)
func (this *QRegularExpression) Escape(str *QString) *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QRegularExpression6escapeERK7QString", ffiqt.FFI_TYPE_POINTER, str)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QRegularExpression_Escape(str *QString) *QString /*123*/ {
	var nilthis *QRegularExpression
	rv := nilthis.Escape(str)
	return rv
}

//  body block end
