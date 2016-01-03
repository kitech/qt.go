package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtGui/qpictureformatplugin.h
// dst-file: /src/gui/qpictureformatplugin.go
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
  // proto:  bool QPictureFormatPlugin::loadPicture(const QString & format, const QString & filename, QPicture * pic);
extern void _ZN20QPictureFormatPlugin11loadPictureERK7QStringS2_P8QPicture(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  bool QPictureFormatPlugin::savePicture(const QString & format, const QString & filename, const QPicture & pic);
extern void _ZN20QPictureFormatPlugin11savePictureERK7QStringS2_RK8QPicture(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPictureFormatPlugin::~QPictureFormatPlugin();
extern void _ZN20QPictureFormatPluginD0Ev(void* qthis);
  // proto:  void QPictureFormatPlugin::QPictureFormatPlugin(QObject * parent);
extern void* dector_ZN20QPictureFormatPluginC1EP7QObject(void* arg0);
extern void _ZN20QPictureFormatPluginC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QPictureFormatPlugin::metaObject();
extern void _ZNK20QPictureFormatPlugin10metaObjectEv(void* qthis);
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

// class sizeof(QPictureFormatPlugin)=1
type QPictureFormatPlugin struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  bool QPictureFormatPlugin::loadPicture(const QString & format, const QString & filename, QPicture * pic);
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
    // invoke: bool loadPicture(const class QString &, const class QString &, class QPicture *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPicture).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN20QPictureFormatPlugin11loadPictureERK7QStringS2_P8QPicture(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "loadPicture", args)
  }

}

  // proto:  bool QPictureFormatPlugin::savePicture(const QString & format, const QString & filename, const QPicture & pic);
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
    // invoke: bool savePicture(const class QString &, const class QString &, const class QPicture &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPicture).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN20QPictureFormatPlugin11savePictureERK7QStringS2_RK8QPicture(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "savePicture", args)
  }

}

  // proto:  void QPictureFormatPlugin::~QPictureFormatPlugin();
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

  // proto:  void QPictureFormatPlugin::QPictureFormatPlugin(QObject * parent);
func NewQPictureFormatPlugin(args ...interface{}) QPictureFormatPlugin {
  return QPictureFormatPlugin{}
}

  // proto:  const QMetaObject * QPictureFormatPlugin::metaObject();
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK20QPictureFormatPlugin10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "metaObject", args)
  }

}

// <= body block end

