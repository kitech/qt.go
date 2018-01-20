//  header block begin
// /usr/include/qt/QtCore/qcommandlineoption.h
// #include <qcommandlineoption.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QCommandLineOption struct {
	*qtrt.CObject
}

func (this *QCommandLineOption) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQCommandLineOptionFromPointer(cthis unsafe.Pointer) *QCommandLineOption {
	return &QCommandLineOption{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qcommandlineoption.h:61
// index:0
// Public
// void QCommandLineOption(const class QString &)
func NewQCommandLineOption(name *QString) *QCommandLineOption {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:62
// index:1
// Public
// void QCommandLineOption(const class QStringList &)
func NewQCommandLineOption_1(names *QStringList) *QCommandLineOption {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = names.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK11QStringList", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:63
// index:2
// Public
// void QCommandLineOption(const class QString &, const class QString &, const class QString &, const class QString &)
func NewQCommandLineOption_2(name *QString, description *QString, valueName *QString, defaultValue *QString) *QCommandLineOption {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = name.GetCthis()
	var convArg1 = description.GetCthis()
	var convArg2 = valueName.GetCthis()
	var convArg3 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK7QStringS2_S2_S2_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:66
// index:3
// Public
// void QCommandLineOption(const class QStringList &, const class QString &, const class QString &, const class QString &)
func NewQCommandLineOption_3(names *QStringList, description *QString, valueName *QString, defaultValue *QString) *QCommandLineOption {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = names.GetCthis()
	var convArg1 = description.GetCthis()
	var convArg2 = valueName.GetCthis()
	var convArg3 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK11QStringListRK7QStringS5_S5_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:71
// index:0
// Public
// void ~QCommandLineOption()
func DeleteQCommandLineOption(*QCommandLineOption) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:78
// index:0
// Public inline
// void swap(class QCommandLineOption &)
func (this *QCommandLineOption) Swap(other *QCommandLineOption) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:81
// index:0
// Public
// QStringList names()
func (this *QCommandLineOption) Names() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption5namesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineoption.h:83
// index:0
// Public
// void setValueName(const class QString &)
func (this *QCommandLineOption) SetValueName(name *QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption12setValueNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:84
// index:0
// Public
// QString valueName()
func (this *QCommandLineOption) ValueName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption9valueNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineoption.h:86
// index:0
// Public
// void setDescription(const class QString &)
func (this *QCommandLineOption) SetDescription(description *QString) {
	var convArg0 = description.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption14setDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:87
// index:0
// Public
// QString description()
func (this *QCommandLineOption) Description() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption11descriptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineoption.h:89
// index:0
// Public
// void setDefaultValue(const class QString &)
func (this *QCommandLineOption) SetDefaultValue(defaultValue *QString) {
	var convArg0 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption15setDefaultValueERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:90
// index:0
// Public
// void setDefaultValues(const class QStringList &)
func (this *QCommandLineOption) SetDefaultValues(defaultValues *QStringList) {
	var convArg0 = defaultValues.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption16setDefaultValuesERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:91
// index:0
// Public
// QStringList defaultValues()
func (this *QCommandLineOption) DefaultValues() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption13defaultValuesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineoption.h:93
// index:0
// Public
// QCommandLineOption::Flags flags()
func (this *QCommandLineOption) Flags() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineoption.h:98
// index:0
// Public
// void setHidden(_Bool)
func (this *QCommandLineOption) SetHidden(hidden bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption9setHiddenEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &hidden)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:100
// index:0
// Public
// bool isHidden()
func (this *QCommandLineOption) IsHidden() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption8isHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
