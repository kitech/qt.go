package qtcore

// /usr/include/qt/QtCore/qloggingcategory.h
// #include <qloggingcategory.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QLoggingCategory struct {
	*qtrt.CObject
}

func (this *QLoggingCategory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLoggingCategory) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQLoggingCategoryFromPointer(cthis unsafe.Pointer) *QLoggingCategory {
	return &QLoggingCategory{&qtrt.CObject{cthis}}
}
func (*QLoggingCategory) NewFromPointer(cthis unsafe.Pointer) *QLoggingCategory {
	return NewQLoggingCategoryFromPointer(cthis)
}

// /usr/include/qt/QtCore/qloggingcategory.h:53
// index:0
// Public
// void QLoggingCategory(const char *)
func NewQLoggingCategory(category string) *QLoggingCategory {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = qtrt.CString(category)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategoryC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLoggingCategoryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qloggingcategory.h:54
// index:1
// Public
// void QLoggingCategory(const char *, QtMsgType)
func NewQLoggingCategory_1(category string, severityLevel int) *QLoggingCategory {
	cthis := qtrt.Calloc(1, 256) // 24
	var convArg0 = qtrt.CString(category)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategoryC2EPKc9QtMsgType", ffiqt.FFI_TYPE_VOID, cthis, convArg0, severityLevel)
	gopp.ErrPrint(err, rv)
	gothis := NewQLoggingCategoryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qloggingcategory.h:55
// index:0
// Public
// void ~QLoggingCategory()
func DeleteQLoggingCategory(*QLoggingCategory) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategoryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:57
// index:0
// Public
// bool isEnabled(QtMsgType)
func (this *QLoggingCategory) IsEnabled(type_ int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory9isEnabledE9QtMsgType", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qloggingcategory.h:58
// index:0
// Public
// void setEnabled(QtMsgType, bool)
func (this *QLoggingCategory) SetEnabled(type_ int, enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategory10setEnabledE9QtMsgTypeb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_, enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:61
// index:0
// Public inline
// bool isDebugEnabled()
func (this *QLoggingCategory) IsDebugEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory14isDebugEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qloggingcategory.h:62
// index:0
// Public inline
// bool isInfoEnabled()
func (this *QLoggingCategory) IsInfoEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory13isInfoEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qloggingcategory.h:63
// index:0
// Public inline
// bool isWarningEnabled()
func (this *QLoggingCategory) IsWarningEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory16isWarningEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qloggingcategory.h:64
// index:0
// Public inline
// bool isCriticalEnabled()
func (this *QLoggingCategory) IsCriticalEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory17isCriticalEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qloggingcategory.h:71
// index:0
// Public inline
// const char * categoryName()
func (this *QLoggingCategory) CategoryName() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory12categoryNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:77
// index:0
// Public static
// QLoggingCategory * defaultCategory()
func (this *QLoggingCategory) DefaultCategory() *QLoggingCategory /*777 QLoggingCategory **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategory15defaultCategoryEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQLoggingCategoryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QLoggingCategory_DefaultCategory() *QLoggingCategory /*777 QLoggingCategory **/ {
	var nilthis *QLoggingCategory
	rv := nilthis.DefaultCategory()
	return rv
}

// /usr/include/qt/QtCore/qloggingcategory.h:82
// index:0
// Public static
// void setFilterRules(const QString &)
func (this *QLoggingCategory) SetFilterRules(rules *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategory14setFilterRulesERK7QString", ffiqt.FFI_TYPE_POINTER, rules)
	gopp.ErrPrint(err, rv)
}
func QLoggingCategory_SetFilterRules(rules *QString) {
	var nilthis *QLoggingCategory
	nilthis.SetFilterRules(rules)
}

type QLoggingCategory__ = int

const QLoggingCategory__DebugShift QLoggingCategory__ = 0
const QLoggingCategory__WarningShift QLoggingCategory__ = 8
const QLoggingCategory__CriticalShift QLoggingCategory__ = 16
const QLoggingCategory__InfoShift QLoggingCategory__ = 24

//  body block end
