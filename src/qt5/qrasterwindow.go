package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtGui/qrasterwindow.h
// dst-file: /src/gui/qrasterwindow.go
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
  // proto:  void QRasterWindow::QRasterWindow(QWindow * parent);
extern void* dector_ZN13QRasterWindowC1EP7QWindow(void* arg0);
extern void _ZN13QRasterWindowC1EP7QWindow(void* qthis, void* arg0);
  // proto:  const QMetaObject * QRasterWindow::metaObject();
extern void _ZNK13QRasterWindow10metaObjectEv(void* qthis);
  // proto:  void QRasterWindow::QRasterWindow(const QRasterWindow & );
extern void* dector_ZN13QRasterWindowC1ERKS_(void* arg0);
extern void _ZN13QRasterWindowC1ERKS_(void* qthis, void* arg0);
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

// class sizeof(QRasterWindow)=1
type QRasterWindow struct {
  /*qbase*/ QPaintDeviceWindow;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QRasterWindow::QRasterWindow(QWindow * parent);
func NewQRasterWindow(args ...interface{}) QRasterWindow {
  return QRasterWindow{}
}

  // proto:  const QMetaObject * QRasterWindow::metaObject();
func (this *QRasterWindow) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QRasterWindow10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QRasterWindow10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QRasterWindow", "metaObject", args)
  }

}

// <= body block end

