//  header block begin
// /usr/include/qt/QtCore/qloggingcategory.h
// #include <qloggingcategory.h>
// #include <QtCore>
package qtcore

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
type QLoggingCategory struct {
	*qtrt.CObject
}

func (this *QLoggingCategory) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qloggingcategory.h:53
// index:0
// void QLoggingCategory(const char *)
func NewQLoggingCategory(category unsafe.Pointer) *QLoggingCategory {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategoryC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, category)
	gopp.ErrPrint(err, rv)
	gothis := NewQLoggingCategoryFromPointer(cthis)
	return gothis
}
func NewQLoggingCategoryFromPointer(cthis unsafe.Pointer) *QLoggingCategory {
	return &QLoggingCategory{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qloggingcategory.h:54
// index:1
// void QLoggingCategory(const char *, enum QtMsgType)
func NewQLoggingCategory_1(category unsafe.Pointer, severityLevel int) *QLoggingCategory {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategoryC2EPKc9QtMsgType", ffiqt.FFI_TYPE_VOID, cthis, category, &severityLevel)
	gopp.ErrPrint(err, rv)
	gothis := NewQLoggingCategoryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qloggingcategory.h:55
// index:0
// void ~QLoggingCategory()
func DeleteQLoggingCategory(*QLoggingCategory) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategoryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:57
// index:0
// bool isEnabled(enum QtMsgType)
func (this *QLoggingCategory) IsEnabled(type_ int) {
	// 0: (, type QtMsgType), (&type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory9isEnabledE9QtMsgType", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:58
// index:0
// void setEnabled(enum QtMsgType, _Bool)
func (this *QLoggingCategory) SetEnabled(type_ int, enable bool) {
	// 0: (, type QtMsgType, enable bool), (&type_, &enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategory10setEnabledE9QtMsgTypeb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &type_, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:61
// index:0
// inline
// bool isDebugEnabled()
func (this *QLoggingCategory) IsDebugEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory14isDebugEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:62
// index:0
// inline
// bool isInfoEnabled()
func (this *QLoggingCategory) IsInfoEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory13isInfoEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:63
// index:0
// inline
// bool isWarningEnabled()
func (this *QLoggingCategory) IsWarningEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory16isWarningEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:64
// index:0
// inline
// bool isCriticalEnabled()
func (this *QLoggingCategory) IsCriticalEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory17isCriticalEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:71
// index:0
// inline
// const char * categoryName()
func (this *QLoggingCategory) CategoryName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QLoggingCategory12categoryNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qloggingcategory.h:77
// index:0
// static
// QLoggingCategory * defaultCategory()
func (this *QLoggingCategory) DefaultCategory() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategory15defaultCategoryEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLoggingCategory_DefaultCategory() {
	// 0: (), ()
	var nilthis *QLoggingCategory
	nilthis.DefaultCategory()
}

// /usr/include/qt/QtCore/qloggingcategory.h:82
// index:0
// static
// void setFilterRules(const class QString &)
func (this *QLoggingCategory) SetFilterRules(rules unsafe.Pointer) {
	// 0: (rules const QString &), (rules)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QLoggingCategory14setFilterRulesERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QLoggingCategory_SetFilterRules(rules unsafe.Pointer) {
	// 0: (rules const QString &), (rules)
	var nilthis *QLoggingCategory
	nilthis.SetFilterRules(rules)
}

//  body block end
