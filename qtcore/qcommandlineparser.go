package qtcore

// /usr/include/qt/QtCore/qcommandlineparser.h
// #include <qcommandlineparser.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 50
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

type QCommandLineParser struct {
	*qtrt.CObject
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParserC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineParserFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQCommandLineParser)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineparser.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QCommandLineParser()
func DeleteQCommandLineParser(this *QCommandLineParser) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParserD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSingleDashWordOptionMode(enum QCommandLineParser::SingleDashWordOptionMode)
func (this *QCommandLineParser) SetSingleDashWordOptionMode(parsingMode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser27setSingleDashWordOptionModeENS_24SingleDashWordOptionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), parsingMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptionsAfterPositionalArgumentsMode(enum QCommandLineParser::OptionsAfterPositionalArgumentsMode)
func (this *QCommandLineParser) SetOptionsAfterPositionalArgumentsMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser38setOptionsAfterPositionalArgumentsModeENS_35OptionsAfterPositionalArgumentsModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addOption(const QCommandLineOption &)
func (this *QCommandLineParser) AddOption(commandLineOption *QCommandLineOption) bool {
	var convArg0 = commandLineOption.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser9addOptionERK18QCommandLineOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QCommandLineOption addVersionOption()
func (this *QCommandLineParser) AddVersionOption() *QCommandLineOption /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser16addVersionOptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCommandLineOption)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QCommandLineOption addHelpOption()
func (this *QCommandLineParser) AddHelpOption() *QCommandLineOption /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser13addHelpOptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQCommandLineOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCommandLineOption)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setApplicationDescription(const QString &)
func (this *QCommandLineParser) SetApplicationDescription(description *QString) {
	var convArg0 = description.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser25setApplicationDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString applicationDescription()
func (this *QCommandLineParser) ApplicationDescription() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser22applicationDescriptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPositionalArgument(const QString &, const QString &, const QString &)
func (this *QCommandLineParser) AddPositionalArgument(name *QString, description *QString, syntax *QString) {
	var convArg0 = name.GetCthis()
	var convArg1 = description.GetCthis()
	var convArg2 = syntax.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearPositionalArguments()
func (this *QCommandLineParser) ClearPositionalArguments() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser24clearPositionalArgumentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void process(const QStringList &)
func (this *QCommandLineParser) Process(arguments *QStringList) {
	var convArg0 = arguments.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser7processERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:85
// index:1
// Public Visibility=Default Availability=Available
// [-2] void process(const QCoreApplication &)
func (this *QCommandLineParser) Process_1(app *QCoreApplication) {
	var convArg0 = app.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser7processERK16QCoreApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool parse(const QStringList &)
func (this *QCommandLineParser) Parse(arguments *QStringList) bool {
	var convArg0 = arguments.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser5parseERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorText()
func (this *QCommandLineParser) ErrorText() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser9errorTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSet(const QString &)
func (this *QCommandLineParser) IsSet(name *QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5isSetERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:94
// index:1
// Public Visibility=Default Availability=Available
// [1] bool isSet(const QCommandLineOption &)
func (this *QCommandLineParser) IsSet_1(option *QCommandLineOption) bool {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5isSetERK18QCommandLineOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcommandlineparser.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QString value(const QString &)
func (this *QCommandLineParser) Value(name *QString) *QString /*123*/ {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5valueERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:95
// index:1
// Public Visibility=Default Availability=Available
// [8] QString value(const QCommandLineOption &)
func (this *QCommandLineParser) Value_1(option *QCommandLineOption) *QString /*123*/ {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5valueERK18QCommandLineOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qcommandlineparser.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showVersion()
func (this *QCommandLineParser) ShowVersion() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser11showVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showHelp(int)
func (this *QCommandLineParser) ShowHelp(exitCode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser8showHelpEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), exitCode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QString helpText()
func (this *QCommandLineParser) HelpText() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser8helpTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

type QCommandLineParser__SingleDashWordOptionMode = int

const QCommandLineParser__ParseAsCompactedShortOptions QCommandLineParser__SingleDashWordOptionMode = 0
const QCommandLineParser__ParseAsLongOptions QCommandLineParser__SingleDashWordOptionMode = 1

type QCommandLineParser__OptionsAfterPositionalArgumentsMode = int

const QCommandLineParser__ParseAsOptions QCommandLineParser__OptionsAfterPositionalArgumentsMode = 0
const QCommandLineParser__ParseAsPositionalArguments QCommandLineParser__OptionsAfterPositionalArgumentsMode = 1

//  body block end
