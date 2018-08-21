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

/*

 */
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

/*
Constructs a command line parser object.
*/
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

/*

 */
func DeleteQCommandLineParser(this *QCommandLineParser) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParserD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSingleDashWordOptionMode(QCommandLineParser::SingleDashWordOptionMode)

/*
Sets the parsing mode to singleDashWordOptionMode. This must be called before process() or parse().
*/
func (this *QCommandLineParser) SetSingleDashWordOptionMode(parsingMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser27setSingleDashWordOptionModeENS_24SingleDashWordOptionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), parsingMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOptionsAfterPositionalArgumentsMode(QCommandLineParser::OptionsAfterPositionalArgumentsMode)

/*
Sets the parsing mode to parsingMode. This must be called before process() or parse().

This function was introduced in  Qt 5.6.
*/
func (this *QCommandLineParser) SetOptionsAfterPositionalArgumentsMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser38setOptionsAfterPositionalArgumentsModeENS_35OptionsAfterPositionalArgumentsModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool addOption(const QCommandLineOption &)

/*
Adds the option option to look for while parsing.

Returns true if adding the option was successful; otherwise returns false.

Adding the option fails if there is no name attached to the option, or the option has a name that clashes with an option name added before.
*/
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

/*
Adds the -v / --version option, which displays the version string of the application.

This option is handled automatically by QCommandLineParser.

You can set the actual version string by using QCoreApplication::setApplicationVersion().

Returns the option instance, which can be used to call isSet().
*/
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

/*
Adds the help option (-h, --help and -? on Windows) This option is handled automatically by QCommandLineParser.

Remember to use setApplicationDescription to set the application description, which will be displayed when this option is used.

Example:


  int main(int argc, char *argv[])
  {
      QCoreApplication app(argc, argv);
      QCoreApplication::setApplicationName("my-copy-program");
      QCoreApplication::setApplicationVersion("1.0");

      QCommandLineParser parser;
      parser.setApplicationDescription("Test helper");
      parser.addHelpOption();
      parser.addVersionOption();
      parser.addPositionalArgument("source", QCoreApplication::translate("main", "Source file to copy."));
      parser.addPositionalArgument("destination", QCoreApplication::translate("main", "Destination directory."));

      // A boolean option with a single name (-p)
      QCommandLineOption showProgressOption("p", QCoreApplication::translate("main", "Show progress during copy"));
      parser.addOption(showProgressOption);

      // A boolean option with multiple names (-f, --force)
      QCommandLineOption forceOption(QStringList() << "f" << "force",
              QCoreApplication::translate("main", "Overwrite existing files."));
      parser.addOption(forceOption);

      // An option with a value
      QCommandLineOption targetDirectoryOption(QStringList() << "t" << "target-directory",
              QCoreApplication::translate("main", "Copy all source files into <directory>."),
              QCoreApplication::translate("main", "directory"));
      parser.addOption(targetDirectoryOption);

      // Process the actual command line arguments given by the user
      parser.process(app);

      const QStringList args = parser.positionalArguments();
      // source is args.at(0), destination is args.at(1)

      bool showProgress = parser.isSet(showProgressOption);
      bool force = parser.isSet(forceOption);
      QString targetDir = parser.value(targetDirectoryOption);
      // ...
  }



Returns the option instance, which can be used to call isSet().
*/
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

/*
Sets the application description shown by helpText().

See also applicationDescription().
*/
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

/*
Returns the application description set in setApplicationDescription().

See also setApplicationDescription().
*/
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

/*
Defines an additional argument to the application, for the benefit of the help text.

The argument name and description will appear under the Arguments: section of the help. If syntax is specified, it will be appended to the Usage line, otherwise the name will be appended.

Example:


  // Usage: image-editor file
  //
  // Arguments:
  //   file                  The file to open.
  parser.addPositionalArgument("file", QCoreApplication::translate("main", "The file to open."));

  // Usage: web-browser [urls...]
  //
  // Arguments:
  //   urls                URLs to open, optionally.
  parser.addPositionalArgument("urls", QCoreApplication::translate("main", "URLs to open, optionally."), "[urls...]");

  // Usage: cp source destination
  //
  // Arguments:
  //   source                Source file to copy.
  //   destination           Destination directory.
  parser.addPositionalArgument("source", QCoreApplication::translate("main", "Source file to copy."));
  parser.addPositionalArgument("destination", QCoreApplication::translate("main", "Destination directory."));



See also addHelpOption() and helpText().
*/
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

/*
Defines an additional argument to the application, for the benefit of the help text.

The argument name and description will appear under the Arguments: section of the help. If syntax is specified, it will be appended to the Usage line, otherwise the name will be appended.

Example:


  // Usage: image-editor file
  //
  // Arguments:
  //   file                  The file to open.
  parser.addPositionalArgument("file", QCoreApplication::translate("main", "The file to open."));

  // Usage: web-browser [urls...]
  //
  // Arguments:
  //   urls                URLs to open, optionally.
  parser.addPositionalArgument("urls", QCoreApplication::translate("main", "URLs to open, optionally."), "[urls...]");

  // Usage: cp source destination
  //
  // Arguments:
  //   source                Source file to copy.
  //   destination           Destination directory.
  parser.addPositionalArgument("source", QCoreApplication::translate("main", "Source file to copy."));
  parser.addPositionalArgument("destination", QCoreApplication::translate("main", "Destination directory."));



See also addHelpOption() and helpText().
*/
func (this *QCommandLineParser) AddPositionalArgument__(name string, description string) {
	var tmpArg0 = NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearPositionalArguments()

/*
Clears the definitions of additional arguments from the help text.

This is only needed for the special case of tools which support multiple commands with different options. Once the actual command has been identified, the options for this command can be defined, and the help text for the command can be adjusted accordingly.

Example:


  QCoreApplication app(argc, argv);
  QCommandLineParser parser;

  parser.addPositionalArgument("command", "The command to execute.");

  // Call parse() to find out the positional arguments.
  parser.parse(QCoreApplication::arguments());

  const QStringList args = parser.positionalArguments();
  const QString command = args.isEmpty() ? QString() : args.first();
  if (command == "resize") {
      parser.clearPositionalArguments();
      parser.addPositionalArgument("resize", "Resize the object to a new size.", "resize [resize_options]");
      parser.addOption(QCommandLineOption("size", "New size.", "new_size"));
      parser.process(app);
      // ...
  }

  \/*
  This code results in context-dependent help:

  $ tool --help
  Usage: tool command

  Arguments:
    command  The command to execute.

  $ tool resize --help
  Usage: tool resize [resize_options]

  Options:
    --size <size>  New size.

  Arguments:
    resize         Resize the object to a new size.
  *\/
*/
func (this *QCommandLineParser) ClearPositionalArguments() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser24clearPositionalArgumentsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void process(const QStringList &)

/*
Processes the command line arguments.

In addition to parsing the options (like parse()), this function also handles the builtin options and handles errors.

The builtin options are --version if addVersionOption was called and --help if addHelpOption was called.

When invoking one of these options, or when an error happens (for instance an unknown option was passed), the current process will then stop, using the exit() function.

See also QCoreApplication::arguments() and parse().
*/
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

/*
Processes the command line arguments.

In addition to parsing the options (like parse()), this function also handles the builtin options and handles errors.

The builtin options are --version if addVersionOption was called and --help if addHelpOption was called.

When invoking one of these options, or when an error happens (for instance an unknown option was passed), the current process will then stop, using the exit() function.

See also QCoreApplication::arguments() and parse().
*/
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

/*
Parses the command line arguments.

Most programs don't need to call this, a simple call to process() is enough.

parse() is more low-level, and only does the parsing. The application will have to take care of the error handling, using errorText() if parse() returns false. This can be useful for instance to show a graphical error message in graphical programs.

Calling parse() instead of process() can also be useful in order to ignore unknown options temporarily, because more option definitions will be provided later on (depending on one of the arguments), before calling process().

Don't forget that arguments must start with the name of the executable (ignored, though).

Returns false in case of a parse error (unknown option or missing value); returns true otherwise.

See also process().
*/
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

/*
Returns a translated error text for the user. This should only be called when parse() returns false.
*/
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

/*
Checks whether the option name was passed to the application.

Returns true if the option name was set, false otherwise.

The name provided can be any long or short name of any option that was added with addOption(). All the options names are treated as being equivalent. If the name is not recognized or that option was not present, false is returned.

Example:


  bool verbose = parser.isSet("verbose");
*/
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

/*
Checks whether the option name was passed to the application.

Returns true if the option name was set, false otherwise.

The name provided can be any long or short name of any option that was added with addOption(). All the options names are treated as being equivalent. If the name is not recognized or that option was not present, false is returned.

Example:


  bool verbose = parser.isSet("verbose");
*/
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

/*
Returns the option value found for the given option name optionName, or an empty string if not found.

The name provided can be any long or short name of any option that was added with addOption(). All the option names are treated as being equivalent. If the name is not recognized or that option was not present, an empty string is returned.

For options found by the parser, the last value found for that option is returned. If the option wasn't specified on the command line, the default value is returned.

An empty string is returned if the option does not take a value.

See also values(), QCommandLineOption::setDefaultValue(), and QCommandLineOption::setDefaultValues().
*/
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

/*
Returns the option value found for the given option name optionName, or an empty string if not found.

The name provided can be any long or short name of any option that was added with addOption(). All the option names are treated as being equivalent. If the name is not recognized or that option was not present, an empty string is returned.

For options found by the parser, the last value found for that option is returned. If the option wasn't specified on the command line, the default value is returned.

An empty string is returned if the option does not take a value.

See also values(), QCommandLineOption::setDefaultValue(), and QCommandLineOption::setDefaultValues().
*/
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

/*
Returns a list of option values found for the given option name optionName, or an empty list if not found.

The name provided can be any long or short name of any option that was added with addOption(). All the options names are treated as being equivalent. If the name is not recognized or that option was not present, an empty list is returned.

For options found by the parser, the list will contain an entry for each time the option was encountered by the parser. If the option wasn't specified on the command line, the default values are returned.

An empty list is returned if the option does not take a value.

See also value(), QCommandLineOption::setDefaultValue(), and QCommandLineOption::setDefaultValues().
*/
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

/*
Returns a list of option values found for the given option name optionName, or an empty list if not found.

The name provided can be any long or short name of any option that was added with addOption(). All the options names are treated as being equivalent. If the name is not recognized or that option was not present, an empty list is returned.

For options found by the parser, the list will contain an entry for each time the option was encountered by the parser. If the option wasn't specified on the command line, the default values are returned.

An empty list is returned if the option does not take a value.

See also value(), QCommandLineOption::setDefaultValue(), and QCommandLineOption::setDefaultValues().
*/
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

/*
Returns a list of positional arguments.

These are all of the arguments that were not recognized as part of an option.
*/
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

/*
Returns a list of option names that were found.

This returns a list of all the recognized option names found by the parser, in the order in which they were found. For any long options that were in the form {--option=value}, the value part will have been dropped.

The names in this list do not include the preceding dash characters. Names may appear more than once in this list if they were encountered more than once by the parser.

Any entry in the list can be used with value() or with values() to get any relevant option values.
*/
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

/*
Returns a list of unknown option names.

This list will include both long an short name options that were not recognized. For any long options that were in the form {--option=value}, the value part will have been dropped and only the long name is added.

The names in this list do not include the preceding dash characters. Names may appear more than once in this list if they were encountered more than once by the parser.

See also optionNames().
*/
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

/*
Displays the version information from QCoreApplication::applicationVersion(), and exits the application. This is automatically triggered by the --version option, but can also be used to display the version when not using process(). The exit code is set to EXIT_SUCCESS (0).

This function was introduced in  Qt 5.4.

See also addVersionOption().
*/
func (this *QCommandLineParser) ShowVersion() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser11showVersionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showHelp(int)

/*
Displays the help information, and exits the application. This is automatically triggered by the --help option, but can also be used to display the help when the user is not invoking the application correctly. The exit code is set to exitCode. It should be set to 0 if the user requested to see the help, and to any other value in case of an error.

See also helpText().
*/
func (this *QCommandLineParser) ShowHelp(exitCode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser8showHelpEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), exitCode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showHelp(int)

/*
Displays the help information, and exits the application. This is automatically triggered by the --help option, but can also be used to display the help when the user is not invoking the application correctly. The exit code is set to exitCode. It should be set to 0 if the user requested to see the help, and to any other value in case of an error.

See also helpText().
*/
func (this *QCommandLineParser) ShowHelp__() {
	// arg: 0, int=Int, =Invalid, , Invalid
	exitCode := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLineParser8showHelpEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), exitCode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcommandlineparser.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QString helpText() const

/*
Returns a string containing the complete help information.

See also showHelp().
*/
func (this *QCommandLineParser) HelpText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLineParser8helpTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

/*
This enum describes the way the parser interprets command-line options that use a single dash followed by multiple letters, as as -abc.



See also setSingleDashWordOptionMode().

*/
type QCommandLineParser__SingleDashWordOptionMode = int

// -abc is interpreted as -a -b -c, i.e. as three short options that have been compacted on the command-line, if none of the options take a value. If a takes a value, then it is interpreted as -a bc, i.e. the short option a followed by the value bc. This is typically used in tools that behave like compilers, in order to handle options such as -DDEFINE=VALUE or -I/include/path. This is the default parsing mode. New applications are recommended to use this mode.
const QCommandLineParser__ParseAsCompactedShortOptions QCommandLineParser__SingleDashWordOptionMode = 0

// -abc is interpreted as --abc, i.e. as the long option named abc. This is how Qt's own tools (uic, rcc...) have always been parsing arguments. This mode should be used for preserving compatibility in applications that were parsing arguments in such a way. There is an exception if the a option has the QCommandLineOption::ShortOptionStyle flag set, in which case it is still interpreted as -a bc.
const QCommandLineParser__ParseAsLongOptions QCommandLineParser__SingleDashWordOptionMode = 1

func (this *QCommandLineParser) SingleDashWordOptionModeItemName(val int) string {
	switch val {
	case QCommandLineParser__ParseAsCompactedShortOptions: // 0
		return "ParseAsCompactedShortOptions"
	case QCommandLineParser__ParseAsLongOptions: // 1
		return "ParseAsLongOptions"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QCommandLineParser_SingleDashWordOptionModeItemName(val int) string {
	var nilthis *QCommandLineParser
	return nilthis.SingleDashWordOptionModeItemName(val)
}

/*
This enum describes the way the parser interprets options that occur after positional arguments.



This enum was introduced or modified in  Qt 5.6.

See also setOptionsAfterPositionalArgumentsMode().

*/
type QCommandLineParser__OptionsAfterPositionalArgumentsMode = int

// application argument --opt -t is interpreted as setting the options opt and t, just like application --opt -t argument would do. This is the default parsing mode. In order to specify that --opt and -t are positional arguments instead, the user can use --, as in application argument -- --opt -t.
const QCommandLineParser__ParseAsOptions QCommandLineParser__OptionsAfterPositionalArgumentsMode = 0

// application argument --opt is interpreted as having two positional arguments, argument and --opt. This mode is useful for executables that aim to launch other executables (e.g. wrappers, debugging tools, etc.) or that support internal commands followed by options for the command. argument is the name of the command, and all options occurring after it can be collected and parsed by another command line parser, possibly in another executable.
const QCommandLineParser__ParseAsPositionalArguments QCommandLineParser__OptionsAfterPositionalArgumentsMode = 1

func (this *QCommandLineParser) OptionsAfterPositionalArgumentsModeItemName(val int) string {
	switch val {
	case QCommandLineParser__ParseAsOptions: // 0
		return "ParseAsOptions"
	case QCommandLineParser__ParseAsPositionalArguments: // 1
		return "ParseAsPositionalArguments"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QCommandLineParser_OptionsAfterPositionalArgumentsModeItemName(val int) string {
	var nilthis *QCommandLineParser
	return nilthis.OptionsAfterPositionalArgumentsModeItemName(val)
}

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
