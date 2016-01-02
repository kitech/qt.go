package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtCore/qstringlist.h
// dst-file: /src/core/qstringlist.go
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
  // proto:  int QStringList::lastIndexOf(const QRegularExpression & re, int from);
extern void demth_ZNK11QStringList11lastIndexOfERK18QRegularExpressioni(void* qthis, void* arg0, int arg1);
  // proto:  void QStringList::QStringList();
extern void* dector_ZN11QStringListC1Ev();
extern void demth_ZN11QStringListC1Ev(void* qthis);
  // proto:  int QStringList::indexOf(const QRegExp & rx, int from);
extern void demth_ZNK11QStringList7indexOfERK7QRegExpi(void* qthis, void* arg0, int arg1);
  // proto:  int QStringList::indexOf(QRegExp & rx, int from);
extern void demth_ZNK11QStringList7indexOfER7QRegExpi(void* qthis, void* arg0, int arg1);
  // proto:  int QStringList::indexOf(const QRegularExpression & re, int from);
extern void demth_ZNK11QStringList7indexOfERK18QRegularExpressioni(void* qthis, void* arg0, int arg1);
  // proto:  int QStringList::lastIndexOf(const QRegExp & rx, int from);
extern void demth_ZNK11QStringList11lastIndexOfERK7QRegExpi(void* qthis, void* arg0, int arg1);
  // proto:  int QStringList::lastIndexOf(QRegExp & rx, int from);
extern void demth_ZNK11QStringList11lastIndexOfER7QRegExpi(void* qthis, void* arg0, int arg1);
  // proto:  void QStringList::QStringList(const QString & i);
extern void* dector_ZN11QStringListC1ERK7QString(void* arg0);
extern void demth_ZN11QStringListC1ERK7QString(void* qthis, void* arg0);
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

// class sizeof(QStringList)=1
type QStringList struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int QStringList::lastIndexOf(const QRegularExpression & re, int from);
func (this *QStringList) lastIndexOf(args ...interface{}) () {
  // lastIndexOf(const class QRegularExpression &, int)
  // lastIndexOf(const class QRegExp &, int)
  // lastIndexOf(class QRegExp &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStringList11lastIndexOfERK18QRegularExpressioni
    // invoke: int lastIndexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK11QStringList11lastIndexOfERK18QRegularExpressioni(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK11QStringList11lastIndexOfERK7QRegExpi
    // invoke: int lastIndexOf(const class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK11QStringList11lastIndexOfERK7QRegExpi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK11QStringList11lastIndexOfER7QRegExpi
    // invoke: int lastIndexOf(class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK11QStringList11lastIndexOfER7QRegExpi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringList", "lastIndexOf", args)
  }

}

  // proto:  void QStringList::QStringList();
func NewQStringList(args ...interface{}) QStringList {
  return QStringList{}
}

  // proto:  int QStringList::indexOf(const QRegExp & rx, int from);
func (this *QStringList) indexOf(args ...interface{}) () {
  // indexOf(const class QRegExp &, int)
  // indexOf(class QRegExp &, int)
  // indexOf(const class QRegularExpression &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStringList7indexOfERK7QRegExpi
    // invoke: int indexOf(const class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK11QStringList7indexOfERK7QRegExpi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK11QStringList7indexOfER7QRegExpi
    // invoke: int indexOf(class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK11QStringList7indexOfER7QRegExpi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK11QStringList7indexOfERK18QRegularExpressioni
    // invoke: int indexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK11QStringList7indexOfERK18QRegularExpressioni(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringList", "indexOf", args)
  }

}

// <= body block end

