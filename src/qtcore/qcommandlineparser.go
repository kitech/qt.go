//  header block begin
// /usr/include/qt/QtCore/qcommandlineparser.h
// #include <qcommandlineparser.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 49
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qcommandlineparser.h:59
// index:0
// void QCommandLineParser()
func NewQCommandLineParser() *QCommandLineParser {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParserC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QCommandLineParser{cthis}
}

// /usr/include/qt/QtCore/qcommandlineparser.h:60
// index:0
// void ~QCommandLineParser()
func DeleteQCommandLineParser(*QCommandLineParser) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParserD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:66
// index:0
// void setSingleDashWordOptionMode(enum QCommandLineParser::SingleDashWordOptionMode)
func (this *QCommandLineParser) SetSingleDashWordOptionMode(parsingMode int) {
	// 0: (, QCommandLineParser::SingleDashWordOptionMode parsingMode), (&parsingMode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser27setSingleDashWordOptionModeENS_24SingleDashWordOptionModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &parsingMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:72
// index:0
// void setOptionsAfterPositionalArgumentsMode(enum QCommandLineParser::OptionsAfterPositionalArgumentsMode)
func (this *QCommandLineParser) SetOptionsAfterPositionalArgumentsMode(mode int) {
	// 0: (, QCommandLineParser::OptionsAfterPositionalArgumentsMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser38setOptionsAfterPositionalArgumentsModeENS_35OptionsAfterPositionalArgumentsModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:74
// index:0
// bool addOption(const class QCommandLineOption &)
func (this *QCommandLineParser) AddOption(commandLineOption unsafe.Pointer) {
	// 0: (, const QCommandLineOption & commandLineOption), (commandLineOption)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser9addOptionERK18QCommandLineOption", ffiqt.FFI_TYPE_VOID, this.cthis, commandLineOption)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:77
// index:0
// QCommandLineOption addVersionOption()
func (this *QCommandLineParser) AddVersionOption() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser16addVersionOptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:78
// index:0
// QCommandLineOption addHelpOption()
func (this *QCommandLineParser) AddHelpOption() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser13addHelpOptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:79
// index:0
// void setApplicationDescription(const class QString &)
func (this *QCommandLineParser) SetApplicationDescription(description unsafe.Pointer) {
	// 0: (, const QString & description), (description)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser25setApplicationDescriptionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, description)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:80
// index:0
// QString applicationDescription()
func (this *QCommandLineParser) ApplicationDescription() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser22applicationDescriptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:81
// index:0
// void addPositionalArgument(const class QString &, const class QString &, const class QString &)
func (this *QCommandLineParser) AddPositionalArgument(name unsafe.Pointer, description unsafe.Pointer, syntax unsafe.Pointer) {
	// 0: (, const QString & name, const QString & description, const QString & syntax), (name, description, syntax)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_", ffiqt.FFI_TYPE_VOID, this.cthis, name, description, syntax)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:82
// index:0
// void clearPositionalArguments()
func (this *QCommandLineParser) ClearPositionalArguments() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser24clearPositionalArgumentsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:84
// index:0
// void process(const class QStringList &)
func (this *QCommandLineParser) Process(arguments unsafe.Pointer) {
	// 0: (, const QStringList & arguments), (arguments)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser7processERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, arguments)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:85
// index:1
// void process(const class QCoreApplication &)
func (this *QCommandLineParser) Process_1(app unsafe.Pointer) {
	// 1: (, const QCoreApplication & app), (app)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser7processERK16QCoreApplication", ffiqt.FFI_TYPE_VOID, this.cthis, app)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:87
// index:0
// bool parse(const class QStringList &)
func (this *QCommandLineParser) Parse(arguments unsafe.Pointer) {
	// 0: (, const QStringList & arguments), (arguments)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser5parseERK11QStringList", ffiqt.FFI_TYPE_VOID, this.cthis, arguments)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:88
// index:0
// QString errorText()
func (this *QCommandLineParser) ErrorText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser9errorTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:90
// index:0
// bool isSet(const class QString &)
func (this *QCommandLineParser) IsSet(name unsafe.Pointer) {
	// 0: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5isSetERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:94
// index:1
// bool isSet(const class QCommandLineOption &)
func (this *QCommandLineParser) IsSet_1(option unsafe.Pointer) {
	// 1: (, const QCommandLineOption & option), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5isSetERK18QCommandLineOption", ffiqt.FFI_TYPE_VOID, this.cthis, option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:91
// index:0
// QString value(const class QString &)
func (this *QCommandLineParser) Value(name unsafe.Pointer) {
	// 0: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5valueERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:95
// index:1
// QString value(const class QCommandLineOption &)
func (this *QCommandLineParser) Value_1(option unsafe.Pointer) {
	// 1: (, const QCommandLineOption & option), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser5valueERK18QCommandLineOption", ffiqt.FFI_TYPE_VOID, this.cthis, option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:92
// index:0
// QStringList values(const class QString &)
func (this *QCommandLineParser) Values(name unsafe.Pointer) {
	// 0: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser6valuesERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:96
// index:1
// QStringList values(const class QCommandLineOption &)
func (this *QCommandLineParser) Values_1(option unsafe.Pointer) {
	// 1: (, const QCommandLineOption & option), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser6valuesERK18QCommandLineOption", ffiqt.FFI_TYPE_VOID, this.cthis, option)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:98
// index:0
// QStringList positionalArguments()
func (this *QCommandLineParser) PositionalArguments() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser19positionalArgumentsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:99
// index:0
// QStringList optionNames()
func (this *QCommandLineParser) OptionNames() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser11optionNamesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:100
// index:0
// QStringList unknownOptionNames()
func (this *QCommandLineParser) UnknownOptionNames() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser18unknownOptionNamesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:102
// index:0
// void showVersion()
func (this *QCommandLineParser) ShowVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser11showVersionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:103
// index:0
// void showHelp(int)
func (this *QCommandLineParser) ShowHelp(exitCode int) {
	// 0: (, int exitCode), (&exitCode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLineParser8showHelpEi", ffiqt.FFI_TYPE_VOID, this.cthis, &exitCode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:104
// index:0
// QString helpText()
func (this *QCommandLineParser) HelpText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLineParser8helpTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
