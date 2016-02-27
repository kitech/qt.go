package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QPictureFormatPlugin::savePicture(const QString & format, const QString & filename, const QPicture & pic);
extern bool C_ZN20QPictureFormatPlugin11savePictureERK7QStringS2_RK8QPicture(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  const QMetaObject * QPictureFormatPlugin::metaObject();
extern void C_ZNK20QPictureFormatPlugin10metaObjectEv(void* qthis); // 4
  // proto:  bool QPictureFormatPlugin::loadPicture(const QString & format, const QString & filename, QPicture * pic);
extern bool C_ZN20QPictureFormatPlugin11loadPictureERK7QStringS2_P8QPicture(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  void QPictureFormatPlugin::QPictureFormatPlugin(QObject * parent);
extern void* C_ZN20QPictureFormatPluginC2EP7QObject(void* arg0); // 3
  // proto:  void QPictureFormatPlugin::~QPictureFormatPlugin();
extern void C_ZN20QPictureFormatPluginD2Ev(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QPictureFormatPlugin)=1
type QPictureFormatPlugin struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// savePicture(const class QString &, const class QString &, const class QPicture &)
func (this *QPictureFormatPlugin) SavePicture(args ...interface{}) (ret interface{}) {
  // savePicture(const class QString &, const class QString &, const class QPicture &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QPicture{}) // "const QPicture &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QPictureFormatPlugin11savePictureERK7QStringS2_RK8QPicture
    // invoke: bool savePicture(const class QString &, const class QString &, const class QPicture &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QPicture).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN20QPictureFormatPlugin11savePictureERK7QStringS2_RK8QPicture(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "savePicture", args)
  }

  return
}

// metaObject()
func (this *QPictureFormatPlugin) MetaObject(args ...interface{}) () {
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
    C.C_ZNK20QPictureFormatPlugin10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "metaObject", args)
  }

  return
}

// loadPicture(const class QString &, const class QString &, class QPicture *)
func (this *QPictureFormatPlugin) LoadPicture(args ...interface{}) (ret interface{}) {
  // loadPicture(const class QString &, const class QString &, class QPicture *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QPicture{}) // "QPicture *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QPictureFormatPlugin11loadPictureERK7QStringS2_P8QPicture
    // invoke: bool loadPicture(const class QString &, const class QString &, class QPicture *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QPicture).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN20QPictureFormatPlugin11loadPictureERK7QStringS2_P8QPicture(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "loadPicture", args)
  }

  return
}

// QPictureFormatPlugin(class QObject *)
func GcfreeQPictureFormatPlugin(this *QPictureFormatPlugin) {
  qtrt.UniverseFree(this)
}
func NewQPictureFormatPlugin(args ...interface{}) *QPictureFormatPlugin {
  // QPictureFormatPlugin(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QPictureFormatPluginC1EP7QObject
    // invoke: void QPictureFormatPlugin(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QPictureFormatPluginC2EP7QObject(arg0)
    this := &QPictureFormatPlugin{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQPictureFormatPlugin)
    return this
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "QPictureFormatPlugin", args)
  }

  return nil // QPictureFormatPlugin{Qclsinst:qthis}
}

// ~QPictureFormatPlugin()
func (this *QPictureFormatPlugin) Free(args ...interface{}) () {
  // ~QPictureFormatPlugin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QPictureFormatPluginD0Ev
    // invoke: void ~QPictureFormatPlugin()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN20QPictureFormatPluginD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QPictureFormatPlugin", "~QPictureFormatPlugin", args)
  }

  return
}

// <= body block end

