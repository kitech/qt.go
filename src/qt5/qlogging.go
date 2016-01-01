package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qlogging.h
// dst-file: /src/core/qlogging.go
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
// class sizeof(QMessageLogContext)=32
type QMessageLogContext struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QMessageLogger)=32
type QMessageLogger struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQMessageLogContext(args ...interface{}) QMessageLogContext {
  return QMessageLogContext{}
}


func (this *QMessageLogContext) copy(args ...interface{}) () {
  // copy(const class QMessageLogContext &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMessageLogContext{}) // "const QMessageLogContext &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QMessageLogContext4copyERKS_
  default:
    qtrt.ErrorResolve("QMessageLogContext", "copy", args)
  }

}


func (this *QMessageLogger) info(args ...interface{}) () {
  // info(const class QLoggingCategory &, const char *, ...)
  // info(CategoryFunction)
  // info(const class QLoggingCategory &)
  // info(CategoryFunction, const char *, ...)
  // info()
  // info(const char *, ...)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[3][1] = qtrt.ByteTy(true) // "const char *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger4infoERK16QLoggingCategoryPKcz
  case 1:
    // invoke: _ZNK14QMessageLogger4infoEPFRK16QLoggingCategoryvE
  case 2:
    // invoke: _ZNK14QMessageLogger4infoERK16QLoggingCategory
  case 3:
    // invoke: _ZNK14QMessageLogger4infoEPFRK16QLoggingCategoryvEPKcz
  case 4:
    // invoke: _ZNK14QMessageLogger4infoEv
  case 5:
    // invoke: _ZNK14QMessageLogger4infoEPKcz
  default:
    qtrt.ErrorResolve("QMessageLogger", "info", args)
  }

}


func (this *QMessageLogger) debug(args ...interface{}) () {
  // debug(const char *, ...)
  // debug(const class QLoggingCategory &, const char *, ...)
  // debug(CategoryFunction, const char *, ...)
  // debug()
  // debug(const class QLoggingCategory &)
  // debug(CategoryFunction)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[2][1] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.VoidpTy() // "CategoryFunction"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger5debugEPKcz
  case 1:
    // invoke: _ZNK14QMessageLogger5debugERK16QLoggingCategoryPKcz
  case 2:
    // invoke: _ZNK14QMessageLogger5debugEPFRK16QLoggingCategoryvEPKcz
  case 3:
    // invoke: _ZNK14QMessageLogger5debugEv
  case 4:
    // invoke: _ZNK14QMessageLogger5debugERK16QLoggingCategory
  case 5:
    // invoke: _ZNK14QMessageLogger5debugEPFRK16QLoggingCategoryvE
  default:
    qtrt.ErrorResolve("QMessageLogger", "debug", args)
  }

}


func (this *QMessageLogger) warning(args ...interface{}) () {
  // warning(const class QLoggingCategory &, const char *, ...)
  // warning(const char *, ...)
  // warning(CategoryFunction)
  // warning(const class QLoggingCategory &)
  // warning(CategoryFunction, const char *, ...)
  // warning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[4][1] = qtrt.ByteTy(true) // "const char *"
  vtys[5] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger7warningERK16QLoggingCategoryPKcz
  case 1:
    // invoke: _ZNK14QMessageLogger7warningEPKcz
  case 2:
    // invoke: _ZNK14QMessageLogger7warningEPFRK16QLoggingCategoryvE
  case 3:
    // invoke: _ZNK14QMessageLogger7warningERK16QLoggingCategory
  case 4:
    // invoke: _ZNK14QMessageLogger7warningEPFRK16QLoggingCategoryvEPKcz
  case 5:
    // invoke: _ZNK14QMessageLogger7warningEv
  default:
    qtrt.ErrorResolve("QMessageLogger", "warning", args)
  }

}


func (this *QMessageLogger) fatal(args ...interface{}) () {
  // fatal(const char *, ...)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger5fatalEPKcz
  default:
    qtrt.ErrorResolve("QMessageLogger", "fatal", args)
  }

}


func (this *QMessageLogger) critical(args ...interface{}) () {
  // critical(CategoryFunction)
  // critical(const char *, ...)
  // critical(const class QLoggingCategory &)
  // critical()
  // critical(CategoryFunction, const char *, ...)
  // critical(const class QLoggingCategory &, const char *, ...)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[4][1] = qtrt.ByteTy(true) // "const char *"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[5][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger8criticalEPFRK16QLoggingCategoryvE
  case 1:
    // invoke: _ZNK14QMessageLogger8criticalEPKcz
  case 2:
    // invoke: _ZNK14QMessageLogger8criticalERK16QLoggingCategory
  case 3:
    // invoke: _ZNK14QMessageLogger8criticalEv
  case 4:
    // invoke: _ZNK14QMessageLogger8criticalEPFRK16QLoggingCategoryvEPKcz
  case 5:
    // invoke: _ZNK14QMessageLogger8criticalERK16QLoggingCategoryPKcz
  default:
    qtrt.ErrorResolve("QMessageLogger", "critical", args)
  }

}


func NewQMessageLogger(args ...interface{}) QMessageLogger {
  return QMessageLogger{}
}


func (this *QMessageLogger) noDebug(args ...interface{}) () {
  // noDebug(const char *, ...)
  // noDebug()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger7noDebugEPKcz
  case 1:
    // invoke: _ZNK14QMessageLogger7noDebugEv
  default:
    qtrt.ErrorResolve("QMessageLogger", "noDebug", args)
  }

}

// <= body block end

