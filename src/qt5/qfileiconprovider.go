package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qfileiconprovider.h
// dst-file: /src/widgets/qfileiconprovider.go
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
// class sizeof(QFileIconProvider)=1
type QFileIconProvider struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QFileIconProvider) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileIconProvider", "type", args)
  }

}


func (this *QFileIconProvider) icon(args ...interface{}) () {
  // icon(const class QFileInfo &)
  // icon(enum QFileIconProvider::IconType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileInfo{}) // "const QFileInfo &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QFileIconProvider::IconType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QFileIconProvider4iconERK9QFileInfo
  case 1:
    // invoke: _ZNK17QFileIconProvider4iconENS_8IconTypeE
  default:
    qtrt.ErrorResolve("QFileIconProvider", "icon", args)
  }

}


func NewQFileIconProvider(args ...interface{}) QFileIconProvider {
  return QFileIconProvider{}
}


func (this *QFileIconProvider) FreeQFileIconProvider(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileIconProvider", "~QFileIconProvider", args)
  }

}

// <= body block end

