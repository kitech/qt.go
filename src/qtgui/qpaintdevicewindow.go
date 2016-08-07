package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QPaintDeviceWindow::update(const QRect & rect);
extern void C_ZN18QPaintDeviceWindow6updateERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QPaintDeviceWindow::update(const QRegion & region);
extern void C_ZN18QPaintDeviceWindow6updateERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QPaintDeviceWindow::update();
extern void C_ZN18QPaintDeviceWindow6updateEv(void* qthis); // 4
  // proto:  const QMetaObject * QPaintDeviceWindow::metaObject();
extern void C_ZNK18QPaintDeviceWindow10metaObjectEv(void* qthis); // 4
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
}

// class sizeof(QPaintDeviceWindow)=1
type QPaintDeviceWindow struct {
  /*qbase*/ QWindow;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// update(const class QRect &)
func (this *QPaintDeviceWindow) Update(args ...interface{}) () {
  // update(const class QRect &)
  // update(const class QRegion &)
  // update()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[2] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QPaintDeviceWindow6updateERK5QRect
    // invoke: void update(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QPaintDeviceWindow6updateERK5QRect(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN18QPaintDeviceWindow6updateERK7QRegion
    // invoke: void update(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QPaintDeviceWindow6updateERK7QRegion(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN18QPaintDeviceWindow6updateEv
    // invoke: void update()
    C.C_ZN18QPaintDeviceWindow6updateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintDeviceWindow", "update", args)
  }

  return
}

// metaObject()
func (this *QPaintDeviceWindow) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QPaintDeviceWindow10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QPaintDeviceWindow10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPaintDeviceWindow", "metaObject", args)
  }

  return
}

// <= body block end

