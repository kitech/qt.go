package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qpictureformatplugin.h
// dst-file: /src/gui/qpictureformatplugin.go
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
// class sizeof(QPictureFormatPlugin)=1
type QPictureFormatPlugin struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QPictureFormatPlugin) loadPicture(args ...interface{}) () {
  // loadPicture(const class QString &, const class QString &, class QPicture *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QPicture{}) // "QPicture *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QPictureFormatPlugin11loadPictureERK7QStringS2_P8QPicture
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "loadPicture", args)
 }

}


func (this *QPictureFormatPlugin) savePicture(args ...interface{}) () {
  // savePicture(const class QString &, const class QString &, const class QPicture &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QPicture{}) // "const QPicture &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QPictureFormatPlugin11savePictureERK7QStringS2_RK8QPicture
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "savePicture", args)
 }

}


func (this *QPictureFormatPlugin) FreeQPictureFormatPlugin(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "~QPictureFormatPlugin", args)
 }

}


func NewQPictureFormatPlugin(args ...interface{})() {
}


func (this *QPictureFormatPlugin) installIOHandler(args ...interface{}) () {
  // installIOHandler(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QPictureFormatPlugin16installIOHandlerERK7QString
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "installIOHandler", args)
 }

}


func (this *QPictureFormatPlugin) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK20QPictureFormatPlugin10metaObjectEv
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "metaObject", args)
 }

}

// <= body block end

