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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qcommandlineoption.h:61
// index:0
// void QCommandLineOption(const class QString &)
func NewQCommandLineOption(name unsafe.Pointer) *QCommandLineOption {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, name)
	gopp.ErrPrint(err, rv)
	return &QCommandLineOption{cthis}
}

// /usr/include/qt/QtCore/qcommandlineoption.h:62
// index:1
// void QCommandLineOption(const class QStringList &)
func NewQCommandLineOption_1(names unsafe.Pointer) *QCommandLineOption {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK11QStringList", ffiqt.FFI_TYPE_VOID, cthis, names)
	gopp.ErrPrint(err, rv)
	return &QCommandLineOption{cthis}
}

// /usr/include/qt/QtCore/qcommandlineoption.h:63
// index:2
// void QCommandLineOption(const class QString &, const class QString &, const class QString &, const class QString &)
func NewQCommandLineOption_2(name unsafe.Pointer, description unsafe.Pointer, valueName unsafe.Pointer, defaultValue unsafe.Pointer) *QCommandLineOption {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK7QStringS2_S2_S2_", ffiqt.FFI_TYPE_VOID, cthis, name, description, valueName, defaultValue)
	gopp.ErrPrint(err, rv)
	return &QCommandLineOption{cthis}
}

// /usr/include/qt/QtCore/qcommandlineoption.h:66
// index:3
// void QCommandLineOption(const class QStringList &, const class QString &, const class QString &, const class QString &)
func NewQCommandLineOption_3(names unsafe.Pointer, description unsafe.Pointer, valueName unsafe.Pointer, defaultValue unsafe.Pointer) *QCommandLineOption {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK11QStringListRK7QStringS5_S5_", ffiqt.FFI_TYPE_VOID, cthis, names, description, valueName, defaultValue)
	gopp.ErrPrint(err, rv)
	return &QCommandLineOption{cthis}
}

// /usr/include/qt/QtCore/qcommandlineoption.h:71
// index:0
// void ~QCommandLineOption()
func DeleteQCommandLineOption(*QCommandLineOption) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOptionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:78
// index:0
// inline
// void swap(class QCommandLineOption &)
func (this *QCommandLineOption) Swap(other unsafe.Pointer) {
	// 0: (, QCommandLineOption & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:81
// index:0
// QStringList names()
func (this *QCommandLineOption) Names() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption5namesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:83
// index:0
// void setValueName(const class QString &)
func (this *QCommandLineOption) SetValueName(name unsafe.Pointer) {
	// 0: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption12setValueNameERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:84
// index:0
// QString valueName()
func (this *QCommandLineOption) ValueName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption9valueNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:86
// index:0
// void setDescription(const class QString &)
func (this *QCommandLineOption) SetDescription(description unsafe.Pointer) {
	// 0: (, const QString & description), (description)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption14setDescriptionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, description)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:87
// index:0
// QString description()
func (this *QCommandLineOption) Description() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption11descriptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:89
// index:0
// void setDefaultValue(const class QString &)
func (this *QCommandLineOption) SetDefaultValue(defaultValue unsafe.Pointer) {
	// 0: (, const QString & defaultValue), (defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption15setDefaultValueERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:90
// index:0
// void setDefaultValues(const class QStringList &)
func (this *QCommandLineOption) SetDefaultValues(defaultValues unsafe.Pointer) {
	// 0: (, const QStringList & defaultValues), (defaultValues)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption16setDefaultValuesERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, defaultValues)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:91
// index:0
// QStringList defaultValues()
func (this *QCommandLineOption) DefaultValues() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption13defaultValuesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:93
// index:0
// QCommandLineOption::Flags flags()
func (this *QCommandLineOption) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption5flagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:98
// index:0
// void setHidden(_Bool)
func (this *QCommandLineOption) SetHidden(hidden bool) {
	// 0: (, bool hidden), (&hidden)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineOption9setHiddenEb", ffiqt.FFI_TYPE_VOID, this.cthis, &hidden)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:100
// index:0
// bool isHidden()
func (this *QCommandLineOption) IsHidden() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineOption8isHiddenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
