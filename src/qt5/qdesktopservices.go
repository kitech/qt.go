package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtGui/qdesktopservices.h
// dst-file: /src/gui/qdesktopservices.go
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static void QDesktopServices::setUrlHandler(const QString & scheme, QObject * receiver, const char * method);
extern void C_ZN16QDesktopServices13setUrlHandlerERK7QStringP7QObjectPKc(void* arg0, void* arg1, void* arg2); // 4
  // proto: static void QDesktopServices::unsetUrlHandler(const QString & scheme);
extern void C_ZN16QDesktopServices15unsetUrlHandlerERK7QString(void* arg0); // 4
  // proto: static bool QDesktopServices::openUrl(const QUrl & url);
extern bool C_ZN16QDesktopServices7openUrlERK4QUrl(void* arg0); // 4
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

// class sizeof(QDesktopServices)=1
type QDesktopServices struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setUrlHandler(const class QString &, class QObject *, const char *)
func (this *QDesktopServices) Seturlhandler_S(args ...interface{}) () {
  // setUrlHandler(const class QString &, class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDesktopServices13setUrlHandlerERK7QStringP7QObjectPKc
    // invoke: void setUrlHandler(const class QString &, class QObject *, const char *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[0][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    C.C_ZN16QDesktopServices13setUrlHandlerERK7QStringP7QObjectPKc(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QDesktopServices", "setUrlHandler", args)
  }

  return
}

// unsetUrlHandler(const class QString &)
func (this *QDesktopServices) Unseturlhandler_S(args ...interface{}) () {
  // unsetUrlHandler(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDesktopServices15unsetUrlHandlerERK7QString
    // invoke: void unsetUrlHandler(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QDesktopServices15unsetUrlHandlerERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QDesktopServices", "unsetUrlHandler", args)
  }

  return
}

// openUrl(const class QUrl &)
func (this *QDesktopServices) Openurl_S(args ...interface{}) (ret interface{}) {
  // openUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QDesktopServices7openUrlERK4QUrl
    // invoke: bool openUrl(const class QUrl &)
    var arg0 = args[0].(*QUrl).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN16QDesktopServices7openUrlERK4QUrl(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDesktopServices", "openUrl", args)
  }

  return
}

// <= body block end

