package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QMetaObject * QRasterWindow::metaObject();
extern void _ZNK13QRasterWindow10metaObjectEv(void* qthis); // 4
  // proto:  void QRasterWindow::QRasterWindow(QWindow * parent);
extern void _ZN13QRasterWindowC2EP7QWindow(void* qthis, void* arg0); // 3
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

// metaObject()
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

// QRasterWindow(class QWindow *)
func NewQRasterWindow(args ...interface{}) QRasterWindow {
  // QRasterWindow(class QWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWindow{}) // "QWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QRasterWindowC1EP7QWindow
    // invoke: void QRasterWindow(class QWindow *)
    var arg0 = args[0].(QWindow).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN13QRasterWindowC2EP7QWindow(qthis, arg0)
  default:
    qtrt.ErrorResolve("QRasterWindow", "QRasterWindow", args)
  }

  return QRasterWindow{}
}

// <= body block end

