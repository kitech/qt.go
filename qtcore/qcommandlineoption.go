package qtcore

// /usr/include/qt/QtCore/qcommandlineoption.h
// #include <qcommandlineoption.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QCommandLineOption struct {
	*qtrt.CObject
}
type QCommandLineOption_ITF interface {
	QCommandLineOption_PTR() *QCommandLineOption
}

func (ptr *QCommandLineOption) QCommandLineOption_PTR() *QCommandLineOption { return ptr }

func (this *QCommandLineOption) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCommandLineOption) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCommandLineOptionFromPointer(cthis unsafe.Pointer) *QCommandLineOption {
	return &QCommandLineOption{&qtrt.CObject{cthis}}
}
func (*QCommandLineOption) NewFromPointer(cthis unsafe.Pointer) *QCommandLineOption {
	return NewQCommandLineOptionFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QString &)

/*
Constructs a command line option object with the name name.

The name can be either short or long. If the name is one character in length, it is considered a short name. Option names must not be empty, must not start with a dash or a slash character, must not contain a = and cannot be repeated.

See also setDescription(), setValueName(), and setDefaultValues().
*/
func NewQCommandLineOption(name string) *QCommandLineOption {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QStringList &)

/*
Constructs a command line option object with the name name.

The name can be either short or long. If the name is one character in length, it is considered a short name. Option names must not be empty, must not start with a dash or a slash character, must not contain a = and cannot be repeated.

See also setDescription(), setValueName(), and setDefaultValues().
*/
func NewQCommandLineOption_1(names QStringList_ITF) *QCommandLineOption {
	var convArg0 unsafe.Pointer
	if names != nil && names.QStringList_PTR() != nil {
		convArg0 = names.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:63
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QString &, const QString &, const QString &, const QString &)

/*
Constructs a command line option object with the name name.

The name can be either short or long. If the name is one character in length, it is considered a short name. Option names must not be empty, must not start with a dash or a slash character, must not contain a = and cannot be repeated.

See also setDescription(), setValueName(), and setDefaultValues().
*/
func NewQCommandLineOption_2(name string, description string, valueName string, defaultValue string) *QCommandLineOption {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(valueName)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(defaultValue)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK7QStringS2_S2_S2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:63
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QString &, const QString &, const QString &, const QString &)

/*
Constructs a command line option object with the name name.

The name can be either short or long. If the name is one character in length, it is considered a short name. Option names must not be empty, must not start with a dash or a slash character, must not contain a = and cannot be repeated.

See also setDescription(), setValueName(), and setDefaultValues().
*/
func NewQCommandLineOption_2_(name string, description string) *QCommandLineOption {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record,
	var convArg2 = NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record,
	var convArg3 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK7QStringS2_S2_S2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:63
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QString &, const QString &, const QString &, const QString &)

/*
Constructs a command line option object with the name name.

The name can be either short or long. If the name is one character in length, it is considered a short name. Option names must not be empty, must not start with a dash or a slash character, must not contain a = and cannot be repeated.

See also setDescription(), setValueName(), and setDefaultValues().
*/
func NewQCommandLineOption_2_1(name string, description string, valueName string) *QCommandLineOption {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(valueName)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record,
	var convArg3 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK7QStringS2_S2_S2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:66
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QStringList &, const QString &, const QString &, const QString &)

/*
Constructs a command line option object with the name name.

The name can be either short or long. If the name is one character in length, it is considered a short name. Option names must not be empty, must not start with a dash or a slash character, must not contain a = and cannot be repeated.

See also setDescription(), setValueName(), and setDefaultValues().
*/
func NewQCommandLineOption_3(names QStringList_ITF, description string, valueName string, defaultValue string) *QCommandLineOption {
	var convArg0 unsafe.Pointer
	if names != nil && names.QStringList_PTR() != nil {
		convArg0 = names.QStringList_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(valueName)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(defaultValue)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK11QStringListRK7QStringS5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:66
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QStringList &, const QString &, const QString &, const QString &)

/*
Constructs a command line option object with the name name.

The name can be either short or long. If the name is one character in length, it is considered a short name. Option names must not be empty, must not start with a dash or a slash character, must not contain a = and cannot be repeated.

See also setDescription(), setValueName(), and setDefaultValues().
*/
func NewQCommandLineOption_3_(names QStringList_ITF, description string) *QCommandLineOption {
	var convArg0 unsafe.Pointer
	if names != nil && names.QStringList_PTR() != nil {
		convArg0 = names.QStringList_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record,
	var convArg2 = NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record,
	var convArg3 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK11QStringListRK7QStringS5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:66
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineOption(const QStringList &, const QString &, const QString &, const QString &)

/*
Constructs a command line option object with the name name.

The name can be either short or long. If the name is one character in length, it is considered a short name. Option names must not be empty, must not start with a dash or a slash character, must not contain a = and cannot be repeated.

See also setDescription(), setValueName(), and setDefaultValues().
*/
func NewQCommandLineOption_3_1(names QStringList_ITF, description string, valueName string) *QCommandLineOption {
	var convArg0 unsafe.Pointer
	if names != nil && names.QStringList_PTR() != nil {
		convArg0 = names.QStringList_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(valueName)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record,
	var convArg3 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionC2ERK11QStringListRK7QStringS5_S5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineOption)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineoption.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCommandLineOption()

/*

 */
func DeleteQCommandLineOption(this *QCommandLineOption) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QCommandLineOption & operator=(const QCommandLineOption &)

/*

 */
func (this *QCommandLineOption) Operator_equal(other QCommandLineOption_ITF) *QCommandLineOption {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCommandLineOption_PTR() != nil {
		convArg0 = other.QCommandLineOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCommandLineOption)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineoption.h:75
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QCommandLineOption & operator=(QCommandLineOption &&)

/*

 */
func (this *QCommandLineOption) Operator_equal_1(other unsafe.Pointer /*333*/) *QCommandLineOption {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOptionaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCommandLineOption)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineoption.h:78
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QCommandLineOption &)

/*
Swaps option other with this option. This operation is very fast and never fails.
*/
func (this *QCommandLineOption) Swap(other QCommandLineOption_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QCommandLineOption_PTR() != nil {
		convArg0 = other.QCommandLineOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOption4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList names() const

/*
Returns the names set for this option.
*/
func (this *QCommandLineOption) Names() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineOption5namesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineoption.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValueName(const QString &)

/*
Sets the name of the expected value, for the documentation, to valueName.

Options without a value assigned have a boolean-like behavior: either the user specifies --option or they don't.

Options with a value assigned need to set a name for the expected value, for the documentation of the option in the help output. An option with names o and output, and a value name of file will appear as -o, --output <file>.

Call QCommandLineParser::value() if you expect the option to be present only once, and QCommandLineParser::values() if you expect that option to be present multiple times.

See also valueName().
*/
func (this *QCommandLineOption) SetValueName(name string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOption12setValueNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QString valueName() const

/*
Returns the name of the expected value.

If empty, the option doesn't take a value.

See also setValueName().
*/
func (this *QCommandLineOption) ValueName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineOption9valueNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcommandlineoption.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescription(const QString &)

/*
Sets the description used for this option to description.

It is customary to add a "." at the end of the description.

The description is used by QCommandLineParser::showHelp().

See also description().
*/
func (this *QCommandLineOption) SetDescription(description string) {
	var tmpArg0 = NewQString_5(description)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOption14setDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description() const

/*
Returns the description set for this option.

See also setDescription().
*/
func (this *QCommandLineOption) Description() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineOption11descriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcommandlineoption.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultValue(const QString &)

/*
Sets the default value used for this option to defaultValue.

The default value is used if the user of the application does not specify the option on the command line.

If defaultValue is empty, the option has no default values.

See also defaultValues() and setDefaultValues().
*/
func (this *QCommandLineOption) SetDefaultValue(defaultValue string) {
	var tmpArg0 = NewQString_5(defaultValue)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOption15setDefaultValueERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultValues(const QStringList &)

/*
Sets the list of default values used for this option to defaultValues.

The default values are used if the user of the application does not specify the option on the command line.

See also defaultValues() and setDefaultValue().
*/
func (this *QCommandLineOption) SetDefaultValues(defaultValues QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if defaultValues != nil && defaultValues.QStringList_PTR() != nil {
		convArg0 = defaultValues.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOption16setDefaultValuesERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList defaultValues() const

/*
Returns the default values set for this option.

See also setDefaultValues().
*/
func (this *QCommandLineOption) DefaultValues() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineOption13defaultValuesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineoption.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QCommandLineOption::Flags flags() const

/*
Returns a set of flags that affect this command-line option.

This function was introduced in  Qt 5.8.

See also setFlags() and QCommandLineOption::Flags.
*/
func (this *QCommandLineOption) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineOption5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QCommandLineOption::Flags)

/*
Set the set of flags that affect this command-line option to flags.

This function was introduced in  Qt 5.8.

See also flags() and QCommandLineOption::Flags.
*/
func (this *QCommandLineOption) SetFlags(aflags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOption8setFlagsE6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), aflags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHidden(bool)

/*

 */
func (this *QCommandLineOption) SetHidden(hidden bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineOption9setHiddenEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hidden)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineoption.h:100
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isHidden() const

/*

 */
func (this *QCommandLineOption) IsHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineOption8isHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QCommandLineOption__Flag = int

//
const QCommandLineOption__HiddenFromHelp QCommandLineOption__Flag = 1

//
const QCommandLineOption__ShortOptionStyle QCommandLineOption__Flag = 2

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
