package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qcommandlineparser.h
// dst-file: /src/core/qcommandlineparser.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QCommandLineParser::process(const QStringList & arguments);
extern void _ZN18QCommandLineParser7processERK11QStringList(void* qthis, void* arg0);
  // proto:  QString QCommandLineParser::value(const QString & name);
extern void _ZNK18QCommandLineParser5valueERK7QString(void* qthis, void* arg0);
  // proto:  QString QCommandLineParser::errorText();
extern void _ZNK18QCommandLineParser9errorTextEv(void* qthis);
  // proto:  void QCommandLineParser::clearPositionalArguments();
extern void _ZN18QCommandLineParser24clearPositionalArgumentsEv(void* qthis);
  // proto:  QStringList QCommandLineParser::values(const QCommandLineOption & option);
extern void _ZNK18QCommandLineParser6valuesERK18QCommandLineOption(void* qthis, void* arg0);
  // proto:  bool QCommandLineParser::isSet(const QString & name);
extern void _ZNK18QCommandLineParser5isSetERK7QString(void* qthis, void* arg0);
  // proto:  void QCommandLineParser::showHelp(int exitCode);
extern void _ZN18QCommandLineParser8showHelpEi(void* qthis, int32_t arg0);
  // proto:  bool QCommandLineParser::addOption(const QCommandLineOption & commandLineOption);
extern void _ZN18QCommandLineParser9addOptionERK18QCommandLineOption(void* qthis, void* arg0);
  // proto:  void QCommandLineParser::showVersion();
extern void _ZN18QCommandLineParser11showVersionEv(void* qthis);
  // proto:  void QCommandLineParser::QCommandLineParser(const QCommandLineParser & );
extern void* dector_ZN18QCommandLineParserC1ERKS_(void* arg0);
extern void _ZN18QCommandLineParserC1ERKS_(void* qthis, void* arg0);
  // proto:  QCommandLineOption QCommandLineParser::addHelpOption();
extern void _ZN18QCommandLineParser13addHelpOptionEv(void* qthis);
  // proto:  QStringList QCommandLineParser::optionNames();
extern void _ZNK18QCommandLineParser11optionNamesEv(void* qthis);
  // proto:  bool QCommandLineParser::isSet(const QCommandLineOption & option);
extern void _ZNK18QCommandLineParser5isSetERK18QCommandLineOption(void* qthis, void* arg0);
  // proto:  void QCommandLineParser::addPositionalArgument(const QString & name, const QString & description, const QString & syntax);
extern void _ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QCommandLineParser::~QCommandLineParser();
extern void _ZN18QCommandLineParserD0Ev(void* qthis);
  // proto:  void QCommandLineParser::process(const QCoreApplication & app);
extern void _ZN18QCommandLineParser7processERK16QCoreApplication(void* qthis, void* arg0);
  // proto:  QString QCommandLineParser::helpText();
extern void _ZNK18QCommandLineParser8helpTextEv(void* qthis);
  // proto:  QStringList QCommandLineParser::values(const QString & name);
extern void _ZNK18QCommandLineParser6valuesERK7QString(void* qthis, void* arg0);
  // proto:  QString QCommandLineParser::applicationDescription();
extern void _ZNK18QCommandLineParser22applicationDescriptionEv(void* qthis);
  // proto:  QString QCommandLineParser::value(const QCommandLineOption & option);
extern void _ZNK18QCommandLineParser5valueERK18QCommandLineOption(void* qthis, void* arg0);
  // proto:  QCommandLineOption QCommandLineParser::addVersionOption();
extern void _ZN18QCommandLineParser16addVersionOptionEv(void* qthis);
  // proto:  QStringList QCommandLineParser::positionalArguments();
extern void _ZNK18QCommandLineParser19positionalArgumentsEv(void* qthis);
  // proto:  void QCommandLineParser::setApplicationDescription(const QString & description);
extern void _ZN18QCommandLineParser25setApplicationDescriptionERK7QString(void* qthis, void* arg0);
  // proto:  void QCommandLineParser::QCommandLineParser();
extern void* dector_ZN18QCommandLineParserC1Ev();
extern void _ZN18QCommandLineParserC1Ev(void* qthis);
  // proto:  bool QCommandLineParser::parse(const QStringList & arguments);
extern void _ZN18QCommandLineParser5parseERK11QStringList(void* qthis, void* arg0);
  // proto:  QStringList QCommandLineParser::unknownOptionNames();
extern void _ZNK18QCommandLineParser18unknownOptionNamesEv(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QCommandLineParser)=8
type QCommandLineParser struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QCommandLineParser::process(const QStringList & arguments);
func (this *QCommandLineParser) process(args ...interface{}) () {
  // process(const class QStringList &)
  // process(const class QCoreApplication &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QCoreApplication{}) // "const QCoreApplication &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser7processERK11QStringList
    // invoke: void process(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLineParser7processERK11QStringList(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN18QCommandLineParser7processERK16QCoreApplication
    // invoke: void process(const class QCoreApplication &)
    var arg0 = args[0].(QCoreApplication).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLineParser7processERK16QCoreApplication(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "process", args)
  }

}

  // proto:  QString QCommandLineParser::value(const QString & name);
func (this *QCommandLineParser) value(args ...interface{}) () {
  // value(const class QString &)
  // value(const class QCommandLineOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QCommandLineOption{}) // "const QCommandLineOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser5valueERK7QString
    // invoke: QString value(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QCommandLineParser5valueERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK18QCommandLineParser5valueERK18QCommandLineOption
    // invoke: QString value(const class QCommandLineOption &)
    var arg0 = args[0].(QCommandLineOption).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QCommandLineParser5valueERK18QCommandLineOption(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "value", args)
  }

}

  // proto:  QString QCommandLineParser::errorText();
func (this *QCommandLineParser) errorText(args ...interface{}) () {
  // errorText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser9errorTextEv
    // invoke: QString errorText()
    C._ZNK18QCommandLineParser9errorTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "errorText", args)
  }

}

  // proto:  void QCommandLineParser::clearPositionalArguments();
func (this *QCommandLineParser) clearPositionalArguments(args ...interface{}) () {
  // clearPositionalArguments()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser24clearPositionalArgumentsEv
    // invoke: void clearPositionalArguments()
    C._ZN18QCommandLineParser24clearPositionalArgumentsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "clearPositionalArguments", args)
  }

}

  // proto:  QStringList QCommandLineParser::values(const QCommandLineOption & option);
func (this *QCommandLineParser) values(args ...interface{}) () {
  // values(const class QCommandLineOption &)
  // values(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCommandLineOption{}) // "const QCommandLineOption &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser6valuesERK18QCommandLineOption
    // invoke: QStringList values(const class QCommandLineOption &)
    var arg0 = args[0].(QCommandLineOption).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QCommandLineParser6valuesERK18QCommandLineOption(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK18QCommandLineParser6valuesERK7QString
    // invoke: QStringList values(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QCommandLineParser6valuesERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "values", args)
  }

}

  // proto:  bool QCommandLineParser::isSet(const QString & name);
func (this *QCommandLineParser) isSet(args ...interface{}) () {
  // isSet(const class QString &)
  // isSet(const class QCommandLineOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QCommandLineOption{}) // "const QCommandLineOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser5isSetERK7QString
    // invoke: bool isSet(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QCommandLineParser5isSetERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK18QCommandLineParser5isSetERK18QCommandLineOption
    // invoke: bool isSet(const class QCommandLineOption &)
    var arg0 = args[0].(QCommandLineOption).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QCommandLineParser5isSetERK18QCommandLineOption(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "isSet", args)
  }

}

  // proto:  void QCommandLineParser::showHelp(int exitCode);
func (this *QCommandLineParser) showHelp(args ...interface{}) () {
  // showHelp(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser8showHelpEi
    // invoke: void showHelp(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLineParser8showHelpEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "showHelp", args)
  }

}

  // proto:  bool QCommandLineParser::addOption(const QCommandLineOption & commandLineOption);
func (this *QCommandLineParser) addOption(args ...interface{}) () {
  // addOption(const class QCommandLineOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCommandLineOption{}) // "const QCommandLineOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser9addOptionERK18QCommandLineOption
    // invoke: bool addOption(const class QCommandLineOption &)
    var arg0 = args[0].(QCommandLineOption).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLineParser9addOptionERK18QCommandLineOption(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addOption", args)
  }

}

  // proto:  void QCommandLineParser::showVersion();
func (this *QCommandLineParser) showVersion(args ...interface{}) () {
  // showVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser11showVersionEv
    // invoke: void showVersion()
    C._ZN18QCommandLineParser11showVersionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "showVersion", args)
  }

}

  // proto:  void QCommandLineParser::QCommandLineParser(const QCommandLineParser & );
func NewQCommandLineParser(args ...interface{}) QCommandLineParser {
  return QCommandLineParser{}
}

  // proto:  QCommandLineOption QCommandLineParser::addHelpOption();
func (this *QCommandLineParser) addHelpOption(args ...interface{}) () {
  // addHelpOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser13addHelpOptionEv
    // invoke: QCommandLineOption addHelpOption()
    C._ZN18QCommandLineParser13addHelpOptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addHelpOption", args)
  }

}

  // proto:  QStringList QCommandLineParser::optionNames();
func (this *QCommandLineParser) optionNames(args ...interface{}) () {
  // optionNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser11optionNamesEv
    // invoke: QStringList optionNames()
    C._ZNK18QCommandLineParser11optionNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "optionNames", args)
  }

}

  // proto:  void QCommandLineParser::addPositionalArgument(const QString & name, const QString & description, const QString & syntax);
func (this *QCommandLineParser) addPositionalArgument(args ...interface{}) () {
  // addPositionalArgument(const class QString &, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_
    // invoke: void addPositionalArgument(const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN18QCommandLineParser21addPositionalArgumentERK7QStringS2_S2_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addPositionalArgument", args)
  }

}

  // proto:  void QCommandLineParser::~QCommandLineParser();
func (this *QCommandLineParser) FreeQCommandLineParser(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCommandLineParser", "~QCommandLineParser", args)
  }

}

  // proto:  QString QCommandLineParser::helpText();
func (this *QCommandLineParser) helpText(args ...interface{}) () {
  // helpText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser8helpTextEv
    // invoke: QString helpText()
    C._ZNK18QCommandLineParser8helpTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "helpText", args)
  }

}

  // proto:  QString QCommandLineParser::applicationDescription();
func (this *QCommandLineParser) applicationDescription(args ...interface{}) () {
  // applicationDescription()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser22applicationDescriptionEv
    // invoke: QString applicationDescription()
    C._ZNK18QCommandLineParser22applicationDescriptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "applicationDescription", args)
  }

}

  // proto:  QCommandLineOption QCommandLineParser::addVersionOption();
func (this *QCommandLineParser) addVersionOption(args ...interface{}) () {
  // addVersionOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser16addVersionOptionEv
    // invoke: QCommandLineOption addVersionOption()
    C._ZN18QCommandLineParser16addVersionOptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addVersionOption", args)
  }

}

  // proto:  QStringList QCommandLineParser::positionalArguments();
func (this *QCommandLineParser) positionalArguments(args ...interface{}) () {
  // positionalArguments()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser19positionalArgumentsEv
    // invoke: QStringList positionalArguments()
    C._ZNK18QCommandLineParser19positionalArgumentsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "positionalArguments", args)
  }

}

  // proto:  void QCommandLineParser::setApplicationDescription(const QString & description);
func (this *QCommandLineParser) setApplicationDescription(args ...interface{}) () {
  // setApplicationDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser25setApplicationDescriptionERK7QString
    // invoke: void setApplicationDescription(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLineParser25setApplicationDescriptionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "setApplicationDescription", args)
  }

}

  // proto:  bool QCommandLineParser::parse(const QStringList & arguments);
func (this *QCommandLineParser) parse(args ...interface{}) () {
  // parse(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineParser5parseERK11QStringList
    // invoke: bool parse(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLineParser5parseERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "parse", args)
  }

}

  // proto:  QStringList QCommandLineParser::unknownOptionNames();
func (this *QCommandLineParser) unknownOptionNames(args ...interface{}) () {
  // unknownOptionNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineParser18unknownOptionNamesEv
    // invoke: QStringList unknownOptionNames()
    C._ZNK18QCommandLineParser18unknownOptionNamesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineParser", "unknownOptionNames", args)
  }

}

// <= body block end

