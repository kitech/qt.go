package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtWidgets/qscrollerproperties.h
// dst-file: /src/widgets/qscrollerproperties.go
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
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static void QScrollerProperties::unsetDefaultScrollerProperties();
extern void C_ZN19QScrollerProperties30unsetDefaultScrollerPropertiesEv(); // 4
  // proto:  void QScrollerProperties::QScrollerProperties();
extern void* C_ZN19QScrollerPropertiesC2Ev(); // 3
  // proto:  void QScrollerProperties::QScrollerProperties(const QScrollerProperties & sp);
extern void* C_ZN19QScrollerPropertiesC2ERKS_(void* arg0); // 3
  // proto: static void QScrollerProperties::setDefaultScrollerProperties(const QScrollerProperties & sp);
extern void C_ZN19QScrollerProperties28setDefaultScrollerPropertiesERKS_(void* arg0); // 4
  // proto:  void QScrollerProperties::~QScrollerProperties();
extern void C_ZN19QScrollerPropertiesD2Ev(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QScrollerProperties)=1
type QScrollerProperties struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// unsetDefaultScrollerProperties()
func (this *QScrollerProperties) Unsetdefaultscrollerproperties_S(args ...interface{}) () {
  // unsetDefaultScrollerProperties()
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
    // invoke: _ZN19QScrollerProperties30unsetDefaultScrollerPropertiesEv
    // invoke: void unsetDefaultScrollerProperties()
    C.C_ZN19QScrollerProperties30unsetDefaultScrollerPropertiesEv()
  default:
    qtrt.ErrorResolve("QScrollerProperties", "unsetDefaultScrollerProperties", args)
  }

  return
}

// QScrollerProperties()
func NewQScrollerProperties(args ...interface{}) *QScrollerProperties {
  // QScrollerProperties()
  // QScrollerProperties(const class QScrollerProperties &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QScrollerProperties{}) // "const QScrollerProperties &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollerPropertiesC1Ev
    // invoke: void QScrollerProperties()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QScrollerPropertiesC2Ev()
    return &QScrollerProperties{Qclsinst:qthis}
  case 1:
    // invoke: _ZN19QScrollerPropertiesC1ERKS_
    // invoke: void QScrollerProperties(const class QScrollerProperties &)
    var arg0 = args[0].(*QScrollerProperties).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QScrollerPropertiesC2ERKS_(arg0)
    return &QScrollerProperties{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QScrollerProperties", "QScrollerProperties", args)
  }

  return nil // QScrollerProperties{Qclsinst:qthis}
}

// setDefaultScrollerProperties(const class QScrollerProperties &)
func (this *QScrollerProperties) Setdefaultscrollerproperties_S(args ...interface{}) () {
  // setDefaultScrollerProperties(const class QScrollerProperties &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScrollerProperties{}) // "const QScrollerProperties &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QScrollerProperties28setDefaultScrollerPropertiesERKS_
    // invoke: void setDefaultScrollerProperties(const class QScrollerProperties &)
    var arg0 = args[0].(*QScrollerProperties).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QScrollerProperties28setDefaultScrollerPropertiesERKS_(arg0)
  default:
    qtrt.ErrorResolve("QScrollerProperties", "setDefaultScrollerProperties", args)
  }

  return
}

// ~QScrollerProperties()
func (this *QScrollerProperties) Freeqscrollerproperties(args ...interface{}) () {
  // ~QScrollerProperties()
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
    // invoke: _ZN19QScrollerPropertiesD0Ev
    // invoke: void ~QScrollerProperties()
    C.C_ZN19QScrollerPropertiesD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScrollerProperties", "~QScrollerProperties", args)
  }

  return
}

// <= body block end

