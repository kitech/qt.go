package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtWidgets/qfontcombobox.h
// dst-file: /src/widgets/qfontcombobox.go
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
  // proto:  FontFilters QFontComboBox::fontFilters();
extern void C_ZNK13QFontComboBox11fontFiltersEv(void* qthis); // 4
  // proto:  void QFontComboBox::setCurrentFont(const QFont & f);
extern void C_ZN13QFontComboBox14setCurrentFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  void QFontComboBox::QFontComboBox(QWidget * parent);
extern void* C_ZN13QFontComboBoxC2EP7QWidget(void* arg0); // 3
  // proto:  QSize QFontComboBox::sizeHint();
extern void* C_ZNK13QFontComboBox8sizeHintEv(void* qthis); // 4
  // proto:  const QMetaObject * QFontComboBox::metaObject();
extern void C_ZNK13QFontComboBox10metaObjectEv(void* qthis); // 4
  // proto:  QFont QFontComboBox::currentFont();
extern void* C_ZNK13QFontComboBox11currentFontEv(void* qthis); // 4
  // proto:  void QFontComboBox::~QFontComboBox();
extern void C_ZN13QFontComboBoxD2Ev(void* qthis); // 4
  // proto:  QFontDatabase::WritingSystem QFontComboBox::writingSystem();
extern void C_ZNK13QFontComboBox13writingSystemEv(void* qthis); // 4
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

// class sizeof(QFontComboBox)=1
type QFontComboBox struct {
  /*qbase*/ QComboBox;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _currentFontChanged QFontComboBox_currentFontChanged_signal;
}

// fontFilters()
func (this *QFontComboBox) Fontfilters(args ...interface{}) () {
  // fontFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox11fontFiltersEv
    // invoke: FontFilters fontFilters()
    C.C_ZNK13QFontComboBox11fontFiltersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontComboBox", "fontFilters", args)
  }

  return
}

// setCurrentFont(const class QFont &)
func (this *QFontComboBox) Setcurrentfont(args ...interface{}) () {
  // setCurrentFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontComboBox14setCurrentFontERK5QFont
    // invoke: void setCurrentFont(const class QFont &)
    var arg0 = args[0].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QFontComboBox14setCurrentFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontComboBox", "setCurrentFont", args)
  }

  return
}

// QFontComboBox(class QWidget *)
func NewQFontComboBox(args ...interface{}) *QFontComboBox {
  // QFontComboBox(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontComboBoxC1EP7QWidget
    // invoke: void QFontComboBox(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QFontComboBoxC2EP7QWidget(arg0)
    return &QFontComboBox{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFontComboBox", "QFontComboBox", args)
  }

  return nil // QFontComboBox{Qclsinst:qthis}
}

// sizeHint()
func (this *QFontComboBox) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK13QFontComboBox8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontComboBox", "sizeHint", args)
  }

  return
}

// metaObject()
func (this *QFontComboBox) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QFontComboBox10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontComboBox", "metaObject", args)
  }

  return
}

// currentFont()
func (this *QFontComboBox) Currentfont(args ...interface{}) (ret interface{}) {
  // currentFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox11currentFontEv
    // invoke: QFont currentFont()
    var ret0 = C.C_ZNK13QFontComboBox11currentFontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontComboBox", "currentFont", args)
  }

  return
}

// ~QFontComboBox()
func (this *QFontComboBox) Freeqfontcombobox(args ...interface{}) () {
  // ~QFontComboBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontComboBoxD0Ev
    // invoke: void ~QFontComboBox()
    C.C_ZN13QFontComboBoxD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontComboBox", "~QFontComboBox", args)
  }

  return
}

// writingSystem()
func (this *QFontComboBox) Writingsystem(args ...interface{}) () {
  // writingSystem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox13writingSystemEv
    // invoke: QFontDatabase::WritingSystem writingSystem()
    C.C_ZNK13QFontComboBox13writingSystemEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontComboBox", "writingSystem", args)
  }

  return
}

// <= body block end

