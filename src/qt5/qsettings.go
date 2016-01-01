package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qsettings.h
// dst-file: /src/core/qsettings.go
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
// class sizeof(QSettings)=1
type QSettings struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQSettings(args ...interface{}) QSettings {
  return QSettings{}
}


func (this *QSettings) isWritable(args ...interface{}) () {
  // isWritable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings10isWritableEv
  default:
    qtrt.ErrorResolve("QSettings", "isWritable", args)
  }

}


func (this *QSettings) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings8fileNameEv
  default:
    qtrt.ErrorResolve("QSettings", "fileName", args)
  }

}


func (this *QSettings) fallbacksEnabled(args ...interface{}) () {
  // fallbacksEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings16fallbacksEnabledEv
  default:
    qtrt.ErrorResolve("QSettings", "fallbacksEnabled", args)
  }

}


func (this *QSettings) applicationName(args ...interface{}) () {
  // applicationName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings15applicationNameEv
  default:
    qtrt.ErrorResolve("QSettings", "applicationName", args)
  }

}


func (this *QSettings) sync(args ...interface{}) () {
  // sync()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings4syncEv
  default:
    qtrt.ErrorResolve("QSettings", "sync", args)
  }

}


func (this *QSettings) setValue(args ...interface{}) () {
  // setValue(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings8setValueERK7QStringRK8QVariant
  default:
    qtrt.ErrorResolve("QSettings", "setValue", args)
  }

}


func (this *QSettings) setArrayIndex(args ...interface{}) () {
  // setArrayIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings13setArrayIndexEi
  default:
    qtrt.ErrorResolve("QSettings", "setArrayIndex", args)
  }

}


func (this *QSettings) setIniCodec(args ...interface{}) () {
  // setIniCodec(class QTextCodec *)
  // setIniCodec(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings11setIniCodecEP10QTextCodec
  case 1:
    // invoke: _ZN9QSettings11setIniCodecEPKc
  default:
    qtrt.ErrorResolve("QSettings", "setIniCodec", args)
  }

}


func (this *QSettings) beginReadArray(args ...interface{}) () {
  // beginReadArray(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings14beginReadArrayERK7QString
  default:
    qtrt.ErrorResolve("QSettings", "beginReadArray", args)
  }

}


func (this *QSettings) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings5clearEv
  default:
    qtrt.ErrorResolve("QSettings", "clear", args)
  }

}


func (this *QSettings) FreeQSettings(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSettings", "~QSettings", args)
  }

}


func (this *QSettings) iniCodec(args ...interface{}) () {
  // iniCodec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings8iniCodecEv
  default:
    qtrt.ErrorResolve("QSettings", "iniCodec", args)
  }

}


func (this *QSettings) setUserIniPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSettings", "setUserIniPath", args)
  }

}


func (this *QSettings) childGroups(args ...interface{}) () {
  // childGroups()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings11childGroupsEv
  default:
    qtrt.ErrorResolve("QSettings", "childGroups", args)
  }

}


func (this *QSettings) value(args ...interface{}) () {
  // value(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings5valueERK7QStringRK8QVariant
  default:
    qtrt.ErrorResolve("QSettings", "value", args)
  }

}


func (this *QSettings) organizationName(args ...interface{}) () {
  // organizationName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings16organizationNameEv
  default:
    qtrt.ErrorResolve("QSettings", "organizationName", args)
  }

}


func (this *QSettings) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings10metaObjectEv
  default:
    qtrt.ErrorResolve("QSettings", "metaObject", args)
  }

}


func (this *QSettings) setFallbacksEnabled(args ...interface{}) () {
  // setFallbacksEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings19setFallbacksEnabledEb
  default:
    qtrt.ErrorResolve("QSettings", "setFallbacksEnabled", args)
  }

}


func (this *QSettings) contains(args ...interface{}) () {
  // contains(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings8containsERK7QString
  default:
    qtrt.ErrorResolve("QSettings", "contains", args)
  }

}


func (this *QSettings) remove(args ...interface{}) () {
  // remove(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings6removeERK7QString
  default:
    qtrt.ErrorResolve("QSettings", "remove", args)
  }

}


func (this *QSettings) endGroup(args ...interface{}) () {
  // endGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings8endGroupEv
  default:
    qtrt.ErrorResolve("QSettings", "endGroup", args)
  }

}


func (this *QSettings) beginWriteArray(args ...interface{}) () {
  // beginWriteArray(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings15beginWriteArrayERK7QStringi
  default:
    qtrt.ErrorResolve("QSettings", "beginWriteArray", args)
  }

}


func (this *QSettings) beginGroup(args ...interface{}) () {
  // beginGroup(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings10beginGroupERK7QString
  default:
    qtrt.ErrorResolve("QSettings", "beginGroup", args)
  }

}


func (this *QSettings) childKeys(args ...interface{}) () {
  // childKeys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings9childKeysEv
  default:
    qtrt.ErrorResolve("QSettings", "childKeys", args)
  }

}


func (this *QSettings) endArray(args ...interface{}) () {
  // endArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings8endArrayEv
  default:
    qtrt.ErrorResolve("QSettings", "endArray", args)
  }

}


func (this *QSettings) setSystemIniPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSettings", "setSystemIniPath", args)
  }

}


func (this *QSettings) allKeys(args ...interface{}) () {
  // allKeys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings7allKeysEv
  default:
    qtrt.ErrorResolve("QSettings", "allKeys", args)
  }

}


func (this *QSettings) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings5groupEv
  default:
    qtrt.ErrorResolve("QSettings", "group", args)
  }

}

// <= body block end

