package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtGui/qbitmap.h
// dst-file: /src/gui/qbitmap.go
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
  // proto:  void QBitmap::QBitmap(const QPixmap & );
extern void* dector_ZN7QBitmapC1ERK7QPixmap(void* arg0);
extern void _ZN7QBitmapC1ERK7QPixmap(void* qthis, void* arg0);
  // proto:  void QBitmap::QBitmap(const QSize & );
extern void* dector_ZN7QBitmapC1ERK5QSize(void* arg0);
extern void _ZN7QBitmapC1ERK5QSize(void* qthis, void* arg0);
  // proto:  void QBitmap::QBitmap(int w, int h);
extern void* dector_ZN7QBitmapC1Eii(int32_t arg0, int32_t arg1);
extern void _ZN7QBitmapC1Eii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QBitmap::~QBitmap();
extern void _ZN7QBitmapD0Ev(void* qthis);
  // proto:  void QBitmap::swap(QBitmap & other);
extern void demth_ZN7QBitmap4swapERS_(void* qthis, void* arg0);
  // proto:  QBitmap QBitmap::transformed(const QMatrix & );
extern void _ZNK7QBitmap11transformedERK7QMatrix(void* qthis, void* arg0);
  // proto:  void QBitmap::clear();
extern void demth_ZN7QBitmap5clearEv(void* qthis);
  // proto:  void QBitmap::QBitmap(const QString & fileName, const char * format);
extern void* dector_ZN7QBitmapC1ERK7QStringPKc(void* arg0, unsigned char* arg1);
extern void _ZN7QBitmapC1ERK7QStringPKc(void* qthis, void* arg0, unsigned char* arg1);
  // proto:  void QBitmap::QBitmap();
extern void* dector_ZN7QBitmapC1Ev();
extern void _ZN7QBitmapC1Ev(void* qthis);
  // proto:  QBitmap QBitmap::transformed(const QTransform & matrix);
extern void _ZNK7QBitmap11transformedERK10QTransform(void* qthis, void* arg0);
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

// class sizeof(QBitmap)=1
type QBitmap struct {
  /*qbase*/ QPixmap;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QBitmap::QBitmap(const QPixmap & );
func NewQBitmap(args ...interface{}) QBitmap {
  return QBitmap{}
}

  // proto:  void QBitmap::~QBitmap();
func (this *QBitmap) FreeQBitmap(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QBitmap", "~QBitmap", args)
  }

}

  // proto:  void QBitmap::swap(QBitmap & other);
func (this *QBitmap) swap(args ...interface{}) () {
  // swap(class QBitmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBitmap{}) // "QBitmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBitmap4swapERS_
    // invoke: void swap(class QBitmap &)
    var arg0 = args[0].(QBitmap).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QBitmap4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBitmap", "swap", args)
  }

}

  // proto:  QBitmap QBitmap::transformed(const QMatrix & );
func (this *QBitmap) transformed(args ...interface{}) () {
  // transformed(const class QMatrix &)
  // transformed(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBitmap11transformedERK7QMatrix
    // invoke: QBitmap transformed(const class QMatrix &)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QBitmap11transformedERK7QMatrix(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QBitmap11transformedERK10QTransform
    // invoke: QBitmap transformed(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QBitmap11transformedERK10QTransform(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QBitmap", "transformed", args)
  }

}

  // proto:  void QBitmap::clear();
func (this *QBitmap) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBitmap5clearEv
    // invoke: void clear()
    C.demth_ZN7QBitmap5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QBitmap", "clear", args)
  }

}

// <= body block end

