package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QScrollerProperties::QScrollerProperties(const QScrollerProperties & sp);
extern void* dector_ZN19QScrollerPropertiesC1ERKS_(void* arg0);
extern void _ZN19QScrollerPropertiesC1ERKS_(void* qthis, void* arg0);
  // proto: static void QScrollerProperties::setDefaultScrollerProperties(const QScrollerProperties & sp);
extern void _ZN19QScrollerProperties28setDefaultScrollerPropertiesERKS_(void* arg0);
  // proto:  void QScrollerProperties::~QScrollerProperties();
extern void _ZN19QScrollerPropertiesD0Ev(void* qthis);
  // proto: static void QScrollerProperties::unsetDefaultScrollerProperties();
extern void _ZN19QScrollerProperties30unsetDefaultScrollerPropertiesEv();
  // proto:  void QScrollerProperties::QScrollerProperties();
extern void* dector_ZN19QScrollerPropertiesC1Ev();
extern void _ZN19QScrollerPropertiesC1Ev(void* qthis);
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

// class sizeof(QScrollerProperties)=1
type QScrollerProperties struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QScrollerProperties::QScrollerProperties(const QScrollerProperties & sp);
func NewQScrollerProperties(args ...interface{}) QScrollerProperties {
  return QScrollerProperties{}
}

  // proto: static void QScrollerProperties::setDefaultScrollerProperties(const QScrollerProperties & sp);
func (this *QScrollerProperties) setDefaultScrollerProperties_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScrollerProperties", "setDefaultScrollerProperties", args)
  }

}

  // proto:  void QScrollerProperties::~QScrollerProperties();
func (this *QScrollerProperties) FreeQScrollerProperties(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScrollerProperties", "~QScrollerProperties", args)
  }

}

  // proto: static void QScrollerProperties::unsetDefaultScrollerProperties();
func (this *QScrollerProperties) unsetDefaultScrollerProperties_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScrollerProperties", "unsetDefaultScrollerProperties", args)
  }

}

// <= body block end

