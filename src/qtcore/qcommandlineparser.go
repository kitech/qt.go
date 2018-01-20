//  header block begin
// /usr/include/qt/QtCore/qcommandlineparser.h
// #include <qcommandlineparser.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 51
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
type QCommandLineParser struct {
	*qtrt.CObject
}

func (this *QCommandLineParser) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQCommandLineParserFromPointer(cthis unsafe.Pointer) *QCommandLineParser {
	return &QCommandLineParser{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qcommandlineparser.h:59
// index:0
// Public
// void QCommandLineParser()
func NewQCommandLineParser() *QCommandLineParser {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParserC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLineParserFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcommandlineparser.h:60
// index:0
// Public
// void ~QCommandLineParser()
func DeleteQCommandLineParser(*QCommandLineParser) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParserD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:66
// index:0
// Public
// void setSingleDashWordOptionMode(enum QCommandLineParser::SingleDashWordOptionMode)
func (this *QCommandLineParser) SetSingleDashWordOptionMode(parsingMode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser27setSingleDashWordOptionModeENS_24SingleDashWordOptionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &parsingMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:72
// index:0
// Public
// void setOptionsAfterPositionalArgumentsMode(enum QCommandLineParser::OptionsAfterPositionalArgumentsMode)
func (this *QCommandLineParser) SetOptionsAfterPositionalArgumentsMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser38setOptionsAfterPositionalArgumentsModeENS_35OptionsAfterPositionalArgumentsModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:74
// index:0
// Public
// bool addOption(const class QCommandLineOption &)
func (this *QCommandLineParser) AddOption(commandLineOption *QCommandLineOption) interface{} {
	var convArg0 = commandLineOption.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser9addOptionERK18QCommandLineOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:77
// index:0
// Public
// QCommandLineOption addVersionOption()
func (this *QCommandLineParser) AddVersionOption() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser16addVersionOptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:78
// index:0
// Public
// QCommandLineOption addHelpOption()
func (this *QCommandLineParser) AddHelpOption() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser13addHelpOptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:79
// index:0
// Public
// void setApplicationDescription(const class QString &)
func (this *QCommandLineParser) SetApplicationDescription(description *QString) {
	var convArg0 = description.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser25setApplicationDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:80
// index:0
// Public
// QString applicationDescription()
func (this *QCommandLineParser) ApplicationDescription() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser22applicationDescriptionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:81
// index:0
// Public
// void addPositionalArgument(const class QString &, const class QString &, const class QString &)
func (this *QCommandLineParser) AddPositionalArgument(name *QString, description *QString, syntax *QString) {
	var convArg0 = name.GetCthis()
	var convArg1 = description.GetCthis()
	var convArg2 = syntax.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:82
// index:0
// Public
// void clearPositionalArguments()
func (this *QCommandLineParser) ClearPositionalArguments() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser24clearPositionalArgumentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:84
// index:0
// Public
// void process(const class QStringList &)
func (this *QCommandLineParser) Process(arguments *QStringList) {
	var convArg0 = arguments.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser7processERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:85
// index:1
// Public
// void process(const class QCoreApplication &)
func (this *QCommandLineParser) Process_1(app *QCoreApplication) {
	var convArg0 = app.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser7processERK16QCoreApplication", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:87
// index:0
// Public
// bool parse(const class QStringList &)
func (this *QCommandLineParser) Parse(arguments *QStringList) interface{} {
	var convArg0 = arguments.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser5parseERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:88
// index:0
// Public
// QString errorText()
func (this *QCommandLineParser) ErrorText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser9errorTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:90
// index:0
// Public
// bool isSet(const class QString &)
func (this *QCommandLineParser) IsSet(name *QString) interface{} {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5isSetERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:94
// index:1
// Public
// bool isSet(const class QCommandLineOption &)
func (this *QCommandLineParser) IsSet_1(option *QCommandLineOption) interface{} {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5isSetERK18QCommandLineOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:91
// index:0
// Public
// QString value(const class QString &)
func (this *QCommandLineParser) Value(name *QString) interface{} {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5valueERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:95
// index:1
// Public
// QString value(const class QCommandLineOption &)
func (this *QCommandLineParser) Value_1(option *QCommandLineOption) interface{} {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5valueERK18QCommandLineOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:92
// index:0
// Public
// QStringList values(const class QString &)
func (this *QCommandLineParser) Values(name *QString) interface{} {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser6valuesERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:96
// index:1
// Public
// QStringList values(const class QCommandLineOption &)
func (this *QCommandLineParser) Values_1(option *QCommandLineOption) interface{} {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser6valuesERK18QCommandLineOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:98
// index:0
// Public
// QStringList positionalArguments()
func (this *QCommandLineParser) PositionalArguments() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser19positionalArgumentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:99
// index:0
// Public
// QStringList optionNames()
func (this *QCommandLineParser) OptionNames() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser11optionNamesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:100
// index:0
// Public
// QStringList unknownOptionNames()
func (this *QCommandLineParser) UnknownOptionNames() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser18unknownOptionNamesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcommandlineparser.h:102
// index:0
// Public
// void showVersion()
func (this *QCommandLineParser) ShowVersion() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser11showVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:103
// index:0
// Public
// void showHelp(int)
func (this *QCommandLineParser) ShowHelp(exitCode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser8showHelpEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &exitCode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:104
// index:0
// Public
// QString helpText()
func (this *QCommandLineParser) HelpText() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser8helpTextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
