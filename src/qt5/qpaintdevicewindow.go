package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qpaintdevicewindow.h
// dst-file: /src/gui/qpaintdevicewindow.go
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
  // proto:  void QPaintDeviceWindow::update(const QRegion & region);
extern void _ZN18QPaintDeviceWindow6updateERK7QRegion(void* qthis, void* arg0);
  // proto:  void QPaintDeviceWindow::update();
extern void _ZN18QPaintDeviceWindow6updateEv(void* qthis);
  // proto:  void QPaintDeviceWindow::QPaintDeviceWindow(const QPaintDeviceWindow & );
extern void* dector_ZN18QPaintDeviceWindowC1ERKS_(void* arg0);
extern void _ZN18QPaintDeviceWindowC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QPaintDeviceWindow::metaObject();
extern void _ZNK18QPaintDeviceWindow10metaObjectEv(void* qthis);
  // proto:  void QPaintDeviceWindow::update(const QRect & rect);
extern void _ZN18QPaintDeviceWindow6updateERK5QRect(void* qthis, void* arg0);
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

// class sizeof(QPaintDeviceWindow)=1
type QPaintDeviceWindow struct {
  /*qbase*/ QWindow;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QPaintDeviceWindow::update(const QRegion & region);
func (this *QPaintDeviceWindow) update(args ...interface{}) () {
  // update(const class QRegion &)
  // update()
  // update(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QPaintDeviceWindow6updateERK7QRegion
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN18QPaintDeviceWindow6updateEv
  case 2:
    // invoke: _ZN18QPaintDeviceWindow6updateERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QPaintDeviceWindow", "update", args)
  }

}

  // proto:  void QPaintDeviceWindow::QPaintDeviceWindow(const QPaintDeviceWindow & );
func NewQPaintDeviceWindow(args ...interface{}) QPaintDeviceWindow {
  return QPaintDeviceWindow{}
}

  // proto:  const QMetaObject * QPaintDeviceWindow::metaObject();
func (this *QPaintDeviceWindow) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QPaintDeviceWindow10metaObjectEv
  default:
    qtrt.ErrorResolve("QPaintDeviceWindow", "metaObject", args)
  }

}

// <= body block end

