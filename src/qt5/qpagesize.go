package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qpagesize.h
// dst-file: /src/gui/qpagesize.go
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
  // proto:  void QPageSize::QPageSize();
extern void* dector_ZN9QPageSizeC1Ev();
extern void _ZN9QPageSizeC1Ev(void* qthis);
  // proto:  void QPageSize::QPageSize(const QString & key, const QSize & pointSize, const QString & name);
extern void* dector_ZN9QPageSizeC1ERK7QStringRK5QSizeS2_(void* arg0, void* arg1, void* arg2);
extern void _ZN9QPageSizeC1ERK7QStringRK5QSizeS2_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QPageSize::~QPageSize();
extern void _ZN9QPageSizeD0Ev(void* qthis);
  // proto:  QString QPageSize::key();
extern void _ZNK9QPageSize3keyEv(void* qthis);
  // proto:  QString QPageSize::name();
extern void _ZNK9QPageSize4nameEv(void* qthis);
  // proto:  QSizeF QPageSize::definitionSize();
extern void _ZNK9QPageSize14definitionSizeEv(void* qthis);
  // proto:  void QPageSize::swap(QPageSize & other);
extern void _ZN9QPageSize4swapERS_(void* qthis, void* arg0);
  // proto:  int QPageSize::windowsId();
extern void _ZNK9QPageSize9windowsIdEv(void* qthis);
  // proto:  QSize QPageSize::sizePixels(int resolution);
extern void _ZNK9QPageSize10sizePixelsEi(void* qthis, int arg0);
  // proto:  void QPageSize::QPageSize(const QPageSize & other);
extern void* dector_ZN9QPageSizeC1ERKS_(void* arg0);
extern void _ZN9QPageSizeC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QPageSize::isValid();
extern void _ZNK9QPageSize7isValidEv(void* qthis);
  // proto:  QRect QPageSize::rectPixels(int resolution);
extern void _ZNK9QPageSize10rectPixelsEi(void* qthis, int arg0);
  // proto:  QRect QPageSize::rectPoints();
extern void _ZNK9QPageSize10rectPointsEv(void* qthis);
  // proto:  void QPageSize::QPageSize(int windowsId, const QSize & pointSize, const QString & name);
extern void* dector_ZN9QPageSizeC1EiRK5QSizeRK7QString(int arg0, void* arg1, void* arg2);
extern void _ZN9QPageSizeC1EiRK5QSizeRK7QString(void* qthis, int arg0, void* arg1, void* arg2);
  // proto:  bool QPageSize::isEquivalentTo(const QPageSize & other);
extern void _ZNK9QPageSize14isEquivalentToERKS_(void* qthis, void* arg0);
  // proto:  QSize QPageSize::sizePoints();
extern void _ZNK9QPageSize10sizePointsEv(void* qthis);
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

// class sizeof(QPageSize)=1
type QPageSize struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QPageSize::QPageSize();
func NewQPageSize(args ...interface{}) QPageSize {
  return QPageSize{}
}

  // proto:  void QPageSize::~QPageSize();
func (this *QPageSize) FreeQPageSize(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPageSize", "~QPageSize", args)
  }

}

  // proto:  QString QPageSize::key();
func (this *QPageSize) key(args ...interface{}) () {
  // key()
  // key(enum QPageSize::PageSizeId)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize3keyEv
  case 1:
    // invoke: _ZN9QPageSize3keyENS_10PageSizeIdE
  default:
    qtrt.ErrorResolve("QPageSize", "key", args)
  }

}

  // proto:  QString QPageSize::name();
func (this *QPageSize) name(args ...interface{}) () {
  // name()
  // name(enum QPageSize::PageSizeId)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize4nameEv
  case 1:
    // invoke: _ZN9QPageSize4nameENS_10PageSizeIdE
  default:
    qtrt.ErrorResolve("QPageSize", "name", args)
  }

}

  // proto:  QSizeF QPageSize::definitionSize();
func (this *QPageSize) definitionSize(args ...interface{}) () {
  // definitionSize()
  // definitionSize(enum QPageSize::PageSizeId)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize14definitionSizeEv
  case 1:
    // invoke: _ZN9QPageSize14definitionSizeENS_10PageSizeIdE
  default:
    qtrt.ErrorResolve("QPageSize", "definitionSize", args)
  }

}

  // proto:  void QPageSize::swap(QPageSize & other);
func (this *QPageSize) swap(args ...interface{}) () {
  // swap(class QPageSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageSize{}) // "QPageSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPageSize4swapERS_
  default:
    qtrt.ErrorResolve("QPageSize", "swap", args)
  }

}

  // proto:  int QPageSize::windowsId();
func (this *QPageSize) windowsId(args ...interface{}) () {
  // windowsId(enum QPageSize::PageSizeId)
  // windowsId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPageSize9windowsIdENS_10PageSizeIdE
  case 1:
    // invoke: _ZNK9QPageSize9windowsIdEv
  default:
    qtrt.ErrorResolve("QPageSize", "windowsId", args)
  }

}

  // proto:  QSize QPageSize::sizePixels(int resolution);
func (this *QPageSize) sizePixels(args ...interface{}) () {
  // sizePixels(int)
  // sizePixels(enum QPageSize::PageSizeId, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize10sizePixelsEi
  case 1:
    // invoke: _ZN9QPageSize10sizePixelsENS_10PageSizeIdEi
  default:
    qtrt.ErrorResolve("QPageSize", "sizePixels", args)
  }

}

  // proto:  bool QPageSize::isValid();
func (this *QPageSize) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize7isValidEv
  default:
    qtrt.ErrorResolve("QPageSize", "isValid", args)
  }

}

  // proto:  QRect QPageSize::rectPixels(int resolution);
func (this *QPageSize) rectPixels(args ...interface{}) () {
  // rectPixels(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize10rectPixelsEi
  default:
    qtrt.ErrorResolve("QPageSize", "rectPixels", args)
  }

}

  // proto:  QRect QPageSize::rectPoints();
func (this *QPageSize) rectPoints(args ...interface{}) () {
  // rectPoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize10rectPointsEv
  default:
    qtrt.ErrorResolve("QPageSize", "rectPoints", args)
  }

}

  // proto:  bool QPageSize::isEquivalentTo(const QPageSize & other);
func (this *QPageSize) isEquivalentTo(args ...interface{}) () {
  // isEquivalentTo(const class QPageSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPageSize{}) // "const QPageSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QPageSize14isEquivalentToERKS_
  default:
    qtrt.ErrorResolve("QPageSize", "isEquivalentTo", args)
  }

}

  // proto:  QSize QPageSize::sizePoints();
func (this *QPageSize) sizePoints(args ...interface{}) () {
  // sizePoints(enum QPageSize::PageSizeId)
  // sizePoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QPageSize::PageSizeId"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QPageSize10sizePointsENS_10PageSizeIdE
  case 1:
    // invoke: _ZNK9QPageSize10sizePointsEv
  default:
    qtrt.ErrorResolve("QPageSize", "sizePoints", args)
  }

}

// <= body block end

