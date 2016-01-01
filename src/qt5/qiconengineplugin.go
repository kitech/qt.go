package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qiconengineplugin.h
// dst-file: /src/gui/qiconengineplugin.go
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
// class sizeof(QIconEnginePlugin)=1
type QIconEnginePlugin struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QIconEnginePlugin) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QIconEnginePlugin10metaObjectEv
  default:
    qtrt.ErrorResolve("QIconEnginePlugin", "metaObject", args)
  }

}


func NewQIconEnginePlugin(args ...interface{}) QIconEnginePlugin {
  return QIconEnginePlugin{}
}


func (this *QIconEnginePlugin) create(args ...interface{}) () {
  // create(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QIconEnginePlugin6createERK7QString
  default:
    qtrt.ErrorResolve("QIconEnginePlugin", "create", args)
  }

}


func (this *QIconEnginePlugin) FreeQIconEnginePlugin(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QIconEnginePlugin", "~QIconEnginePlugin", args)
  }

}

// <= body block end

