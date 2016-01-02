package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto: static void QDesktopServices::unsetUrlHandler(const QString & scheme);
extern void _ZN16QDesktopServices15unsetUrlHandlerERK7QString(void* arg0);
  // proto: static bool QDesktopServices::openUrl(const QUrl & url);
extern void _ZN16QDesktopServices7openUrlERK4QUrl(void* arg0);
  // proto: static void QDesktopServices::setUrlHandler(const QString & scheme, QObject * receiver, const char * method);
extern void _ZN16QDesktopServices13setUrlHandlerERK7QStringP7QObjectPKc(void* arg0, void* arg1, char* arg2);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto: static void QDesktopServices::unsetUrlHandler(const QString & scheme);
func (this *QDesktopServices) unsetUrlHandler_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDesktopServices", "unsetUrlHandler", args)
  }

}

  // proto: static bool QDesktopServices::openUrl(const QUrl & url);
func (this *QDesktopServices) openUrl_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDesktopServices", "openUrl", args)
  }

}

  // proto: static void QDesktopServices::setUrlHandler(const QString & scheme, QObject * receiver, const char * method);
func (this *QDesktopServices) setUrlHandler_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDesktopServices", "setUrlHandler", args)
  }

}

// <= body block end

