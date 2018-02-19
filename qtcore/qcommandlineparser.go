package qtcore

// /usr/include/qt/QtCore/qcommandlineparser.h
// #include <qcommandlineparser.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 52
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QCommandLineParser struct {
	*qtrt.CObject
}
type QCommandLineParser_ITF interface {
	QCommandLineParser_PTR() *QCommandLineParser
}

func (ptr *QCommandLineParser) QCommandLineParser_PTR() *QCommandLineParser { return ptr }

func (this *QCommandLineParser) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCommandLineParser) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCommandLineParserFromPointer(cthis unsafe.Pointer) *QCommandLineParser {
	return &QCommandLineParser{&qtrt.CObject{cthis}}
}
func (*QCommandLineParser) NewFromPointer(cthis unsafe.Pointer) *QCommandLineParser {
	return NewQCommandLineParserFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCommandLineParser()
func NewQCommandLineParser() *QCommandLineParser {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParserC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCommandLineParserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineParser)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineparser.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCommandLineParser()
func DeleteQCommandLineParser(this *QCommandLineParser) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParserD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSingleDashWordOptionMode(enum QCommandLineParser::SingleDashWordOptionMode)
func (this *QCommandLineParser) SetSingleDashWordOptionMode(parsingMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser27setSingleDashWordOptionModeENS_24SingleDashWordOptionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parsingMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptionsAfterPositionalArgumentsMode(enum QCommandLineParser::OptionsAfterPositionalArgumentsMode)
func (this *QCommandLineParser) SetOptionsAfterPositionalArgumentsMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser38setOptionsAfterPositionalArgumentsModeENS_35OptionsAfterPositionalArgumentsModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addOption(const QCommandLineOption &)
func (this *QCommandLineParser) AddOption(commandLineOption QCommandLineOption_ITF) bool {
	var convArg0 unsafe.Pointer
	if commandLineOption != nil && commandLineOption.QCommandLineOption_PTR() != nil {
		convArg0 = commandLineOption.QCommandLineOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser9addOptionERK18QCommandLineOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QCommandLineOption addVersionOption()
func (this *QCommandLineParser) AddVersionOption() *QCommandLineOption /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser16addVersionOptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCommandLineOption)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QCommandLineOption addHelpOption()
func (this *QCommandLineParser) AddHelpOption() *QCommandLineOption /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser13addHelpOptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCommandLineOption)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setApplicationDescription(const QString &)
func (this *QCommandLineParser) SetApplicationDescription(description string) {
	var tmpArg0 = NewQString_5(description)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser25setApplicationDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString applicationDescription() const
func (this *QCommandLineParser) ApplicationDescription() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser22applicationDescriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcommandlineparser.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPositionalArgument(const QString &, const QString &, const QString &)
func (this *QCommandLineParser) AddPositionalArgument(name string, description string, syntax string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(syntax)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPositionalArgument(const QString &, const QString &, const QString &)
func (this *QCommandLineParser) AddPositionalArgument__(name string, description string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record,
	var convArg2 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearPositionalArguments()
func (this *QCommandLineParser) ClearPositionalArguments() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser24clearPositionalArgumentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void process(const QStringList &)
func (this *QCommandLineParser) Process(arguments QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if arguments != nil && arguments.QStringList_PTR() != nil {
		convArg0 = arguments.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser7processERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:85
// index:1
// Public Visibility=Default Availability=Available
// [-2] void process(const QCoreApplication &)
func (this *QCommandLineParser) Process_1(app QCoreApplication_ITF) {
	var convArg0 unsafe.Pointer
	if app != nil && app.QCoreApplication_PTR() != nil {
		convArg0 = app.QCoreApplication_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser7processERK16QCoreApplication", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool parse(const QStringList &)
func (this *QCommandLineParser) Parse(arguments QStringList_ITF) bool {
	var convArg0 unsafe.Pointer
	if arguments != nil && arguments.QStringList_PTR() != nil {
		convArg0 = arguments.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser5parseERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorText() const
func (this *QCommandLineParser) ErrorText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser9errorTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcommandlineparser.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSet(const QString &) const
func (this *QCommandLineParser) IsSet(name string) bool {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser5isSetERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:94
// index:1
// Public Visibility=Default Availability=Available
// [1] bool isSet(const QCommandLineOption &) const
func (this *QCommandLineParser) IsSet_1(option QCommandLineOption_ITF) bool {
	var convArg0 unsafe.Pointer
	if option != nil && option.QCommandLineOption_PTR() != nil {
		convArg0 = option.QCommandLineOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser5isSetERK18QCommandLineOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QString value(const QString &) const
func (this *QCommandLineParser) Value(name string) string {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser5valueERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcommandlineparser.h:95
// index:1
// Public Visibility=Default Availability=Available
// [8] QString value(const QCommandLineOption &) const
func (this *QCommandLineParser) Value_1(option QCommandLineOption_ITF) string {
	var convArg0 unsafe.Pointer
	if option != nil && option.QCommandLineOption_PTR() != nil {
		convArg0 = option.QCommandLineOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser5valueERK18QCommandLineOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcommandlineparser.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList values(const QString &) const
func (this *QCommandLineParser) Values(name string) *QStringList /*123*/ {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser6valuesERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:96
// index:1
// Public Visibility=Default Availability=Available
// [8] QStringList values(const QCommandLineOption &) const
func (this *QCommandLineParser) Values_1(option QCommandLineOption_ITF) *QStringList /*123*/ {
	var convArg0 unsafe.Pointer
	if option != nil && option.QCommandLineOption_PTR() != nil {
		convArg0 = option.QCommandLineOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser6valuesERK18QCommandLineOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList positionalArguments() const
func (this *QCommandLineParser) PositionalArguments() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser19positionalArgumentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:99
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList optionNames() const
func (this *QCommandLineParser) OptionNames() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser11optionNamesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList unknownOptionNames() const
func (this *QCommandLineParser) UnknownOptionNames() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser18unknownOptionNamesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showVersion()
func (this *QCommandLineParser) ShowVersion() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser11showVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showHelp(int)
func (this *QCommandLineParser) ShowHelp(exitCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser8showHelpEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), exitCode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showHelp(int)
func (this *QCommandLineParser) ShowHelp__() {
	// arg: 0, int=Int, =Invalid,
	exitCode := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser8showHelpEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), exitCode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QString helpText() const
func (this *QCommandLineParser) HelpText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser8helpTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

type QCommandLineParser__SingleDashWordOptionMode = int

const QCommandLineParser__ParseAsCompactedShortOptions QCommandLineParser__SingleDashWordOptionMode = 0
const QCommandLineParser__ParseAsLongOptions QCommandLineParser__SingleDashWordOptionMode = 1

type QCommandLineParser__OptionsAfterPositionalArgumentsMode = int

const QCommandLineParser__ParseAsOptions QCommandLineParser__OptionsAfterPositionalArgumentsMode = 0
const QCommandLineParser__ParseAsPositionalArguments QCommandLineParser__OptionsAfterPositionalArgumentsMode = 1

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
