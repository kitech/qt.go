package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qcommandlineparser.h
// dst-file: /src/core/qcommandlineparser.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QCommandLineParser)=8
type QCommandLineParser struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


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
  case 1:
    // invoke: _ZN18QCommandLineParser7processERK16QCoreApplication
  default:
    qtrt.ErrorResolve("QCommandLineParser", "process", args)
  }

}


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
  case 1:
    // invoke: _ZNK18QCommandLineParser5valueERK18QCommandLineOption
  default:
    qtrt.ErrorResolve("QCommandLineParser", "value", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "errorText", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "clearPositionalArguments", args)
  }

}


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
  case 1:
    // invoke: _ZNK18QCommandLineParser6valuesERK7QString
  default:
    qtrt.ErrorResolve("QCommandLineParser", "values", args)
  }

}


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
  case 1:
    // invoke: _ZNK18QCommandLineParser5isSetERK18QCommandLineOption
  default:
    qtrt.ErrorResolve("QCommandLineParser", "isSet", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "showHelp", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addOption", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "showVersion", args)
  }

}


func NewQCommandLineParser(args ...interface{}) QCommandLineParser {
  return QCommandLineParser{}
}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addHelpOption", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "optionNames", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addPositionalArgument", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "helpText", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "applicationDescription", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "addVersionOption", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "positionalArguments", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "setApplicationDescription", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "parse", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QCommandLineParser", "unknownOptionNames", args)
  }

}

// <= body block end

