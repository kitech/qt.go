package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtWidgets/qabstractspinbox.h
// dst-file: /src/widgets/qabstractspinbox.go
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
  // proto:  void QAbstractSpinBox::selectAll();
extern void C_ZN16QAbstractSpinBox9selectAllEv(void* qthis); // 4
  // proto:  QString QAbstractSpinBox::text();
extern void* C_ZNK16QAbstractSpinBox4textEv(void* qthis); // 4
  // proto:  void QAbstractSpinBox::setAccelerated(bool on);
extern void C_ZN16QAbstractSpinBox14setAcceleratedEb(void* qthis, bool arg0); // 4
  // proto:  bool QAbstractSpinBox::keyboardTracking();
extern bool C_ZNK16QAbstractSpinBox16keyboardTrackingEv(void* qthis); // 4
  // proto:  bool QAbstractSpinBox::hasAcceptableInput();
extern bool C_ZNK16QAbstractSpinBox18hasAcceptableInputEv(void* qthis); // 4
  // proto:  QAbstractSpinBox::ButtonSymbols QAbstractSpinBox::buttonSymbols();
extern void C_ZNK16QAbstractSpinBox13buttonSymbolsEv(void* qthis); // 4
  // proto:  bool QAbstractSpinBox::hasFrame();
extern bool C_ZNK16QAbstractSpinBox8hasFrameEv(void* qthis); // 4
  // proto:  void QAbstractSpinBox::setWrapping(bool w);
extern void C_ZN16QAbstractSpinBox11setWrappingEb(void* qthis, bool arg0); // 4
  // proto:  bool QAbstractSpinBox::event(QEvent * event);
extern bool C_ZN16QAbstractSpinBox5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  Qt::Alignment QAbstractSpinBox::alignment();
extern void C_ZNK16QAbstractSpinBox9alignmentEv(void* qthis); // 4
  // proto:  void QAbstractSpinBox::fixup(QString & input);
extern void C_ZNK16QAbstractSpinBox5fixupER7QString(void* qthis, void* arg0); // 4
  // proto:  void QAbstractSpinBox::setSpecialValueText(const QString & txt);
extern void C_ZN16QAbstractSpinBox19setSpecialValueTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QAbstractSpinBox::isReadOnly();
extern bool C_ZNK16QAbstractSpinBox10isReadOnlyEv(void* qthis); // 4
  // proto:  void QAbstractSpinBox::setGroupSeparatorShown(bool shown);
extern void C_ZN16QAbstractSpinBox22setGroupSeparatorShownEb(void* qthis, bool arg0); // 4
  // proto:  void QAbstractSpinBox::~QAbstractSpinBox();
extern void C_ZN16QAbstractSpinBoxD2Ev(void* qthis); // 4
  // proto:  void QAbstractSpinBox::stepUp();
extern void C_ZN16QAbstractSpinBox6stepUpEv(void* qthis); // 4
  // proto:  void QAbstractSpinBox::stepBy(int steps);
extern void C_ZN16QAbstractSpinBox6stepByEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QAbstractSpinBox::specialValueText();
extern void* C_ZNK16QAbstractSpinBox16specialValueTextEv(void* qthis); // 4
  // proto:  bool QAbstractSpinBox::isAccelerated();
extern bool C_ZNK16QAbstractSpinBox13isAcceleratedEv(void* qthis); // 4
  // proto:  bool QAbstractSpinBox::wrapping();
extern bool C_ZNK16QAbstractSpinBox8wrappingEv(void* qthis); // 4
  // proto:  QSize QAbstractSpinBox::sizeHint();
extern void* C_ZNK16QAbstractSpinBox8sizeHintEv(void* qthis); // 4
  // proto:  void QAbstractSpinBox::interpretText();
extern void C_ZN16QAbstractSpinBox13interpretTextEv(void* qthis); // 4
  // proto:  QValidator::State QAbstractSpinBox::validate(QString & input, int & pos);
extern void C_ZNK16QAbstractSpinBox8validateER7QStringRi(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QAbstractSpinBox::isGroupSeparatorShown();
extern bool C_ZNK16QAbstractSpinBox21isGroupSeparatorShownEv(void* qthis); // 4
  // proto:  QAbstractSpinBox::CorrectionMode QAbstractSpinBox::correctionMode();
extern void C_ZNK16QAbstractSpinBox14correctionModeEv(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractSpinBox::metaObject();
extern void C_ZNK16QAbstractSpinBox10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractSpinBox::QAbstractSpinBox(QWidget * parent);
extern void* C_ZN16QAbstractSpinBoxC2EP7QWidget(void* arg0); // 3
  // proto:  void QAbstractSpinBox::stepDown();
extern void C_ZN16QAbstractSpinBox8stepDownEv(void* qthis); // 4
  // proto:  void QAbstractSpinBox::setFrame(bool );
extern void C_ZN16QAbstractSpinBox8setFrameEb(void* qthis, bool arg0); // 4
  // proto:  void QAbstractSpinBox::clear();
extern void C_ZN16QAbstractSpinBox5clearEv(void* qthis); // 4
  // proto:  QSize QAbstractSpinBox::minimumSizeHint();
extern void* C_ZNK16QAbstractSpinBox15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QAbstractSpinBox::setKeyboardTracking(bool kt);
extern void C_ZN16QAbstractSpinBox19setKeyboardTrackingEb(void* qthis, bool arg0); // 4
  // proto:  void QAbstractSpinBox::setReadOnly(bool r);
extern void C_ZN16QAbstractSpinBox11setReadOnlyEb(void* qthis, bool arg0); // 4
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

// class sizeof(QAbstractSpinBox)=1
type QAbstractSpinBox struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _editingFinished QAbstractSpinBox_editingFinished_signal;
}

// selectAll()
func (this *QAbstractSpinBox) Selectall(args ...interface{}) () {
  // selectAll()
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
    // invoke: _ZN16QAbstractSpinBox9selectAllEv
    // invoke: void selectAll()
    C.C_ZN16QAbstractSpinBox9selectAllEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "selectAll", args)
  }

  return
}

// text()
func (this *QAbstractSpinBox) Text(args ...interface{}) (ret interface{}) {
  // text()
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
    // invoke: _ZNK16QAbstractSpinBox4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK16QAbstractSpinBox4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "text", args)
  }

  return
}

// setAccelerated(_Bool)
func (this *QAbstractSpinBox) Setaccelerated(args ...interface{}) () {
  // setAccelerated(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox14setAcceleratedEb
    // invoke: void setAccelerated(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QAbstractSpinBox14setAcceleratedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setAccelerated", args)
  }

  return
}

// keyboardTracking()
func (this *QAbstractSpinBox) Keyboardtracking(args ...interface{}) (ret interface{}) {
  // keyboardTracking()
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
    // invoke: _ZNK16QAbstractSpinBox16keyboardTrackingEv
    // invoke: bool keyboardTracking()
    var ret0 = C.C_ZNK16QAbstractSpinBox16keyboardTrackingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "keyboardTracking", args)
  }

  return
}

// hasAcceptableInput()
func (this *QAbstractSpinBox) Hasacceptableinput(args ...interface{}) (ret interface{}) {
  // hasAcceptableInput()
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
    // invoke: _ZNK16QAbstractSpinBox18hasAcceptableInputEv
    // invoke: bool hasAcceptableInput()
    var ret0 = C.C_ZNK16QAbstractSpinBox18hasAcceptableInputEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "hasAcceptableInput", args)
  }

  return
}

// buttonSymbols()
func (this *QAbstractSpinBox) Buttonsymbols(args ...interface{}) () {
  // buttonSymbols()
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
    // invoke: _ZNK16QAbstractSpinBox13buttonSymbolsEv
    // invoke: QAbstractSpinBox::ButtonSymbols buttonSymbols()
    C.C_ZNK16QAbstractSpinBox13buttonSymbolsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "buttonSymbols", args)
  }

  return
}

// hasFrame()
func (this *QAbstractSpinBox) Hasframe(args ...interface{}) (ret interface{}) {
  // hasFrame()
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
    // invoke: _ZNK16QAbstractSpinBox8hasFrameEv
    // invoke: bool hasFrame()
    var ret0 = C.C_ZNK16QAbstractSpinBox8hasFrameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "hasFrame", args)
  }

  return
}

// setWrapping(_Bool)
func (this *QAbstractSpinBox) Setwrapping(args ...interface{}) () {
  // setWrapping(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox11setWrappingEb
    // invoke: void setWrapping(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QAbstractSpinBox11setWrappingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setWrapping", args)
  }

  return
}

// event(class QEvent *)
func (this *QAbstractSpinBox) Event(args ...interface{}) (ret interface{}) {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QEvent{}) // "QEvent *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(*qtcore.QEvent).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN16QAbstractSpinBox5eventEP6QEvent(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "event", args)
  }

  return
}

// alignment()
func (this *QAbstractSpinBox) Alignment(args ...interface{}) () {
  // alignment()
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
    // invoke: _ZNK16QAbstractSpinBox9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK16QAbstractSpinBox9alignmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "alignment", args)
  }

  return
}

// fixup(class QString &)
func (this *QAbstractSpinBox) Fixup(args ...interface{}) () {
  // fixup(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox5fixupER7QString
    // invoke: void fixup(class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK16QAbstractSpinBox5fixupER7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "fixup", args)
  }

  return
}

// setSpecialValueText(const class QString &)
func (this *QAbstractSpinBox) Setspecialvaluetext(args ...interface{}) () {
  // setSpecialValueText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox19setSpecialValueTextERK7QString
    // invoke: void setSpecialValueText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QAbstractSpinBox19setSpecialValueTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setSpecialValueText", args)
  }

  return
}

// isReadOnly()
func (this *QAbstractSpinBox) Isreadonly(args ...interface{}) (ret interface{}) {
  // isReadOnly()
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
    // invoke: _ZNK16QAbstractSpinBox10isReadOnlyEv
    // invoke: bool isReadOnly()
    var ret0 = C.C_ZNK16QAbstractSpinBox10isReadOnlyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "isReadOnly", args)
  }

  return
}

// setGroupSeparatorShown(_Bool)
func (this *QAbstractSpinBox) Setgroupseparatorshown(args ...interface{}) () {
  // setGroupSeparatorShown(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox22setGroupSeparatorShownEb
    // invoke: void setGroupSeparatorShown(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QAbstractSpinBox22setGroupSeparatorShownEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setGroupSeparatorShown", args)
  }

  return
}

// ~QAbstractSpinBox()
func (this *QAbstractSpinBox) Freeqabstractspinbox(args ...interface{}) () {
  // ~QAbstractSpinBox()
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
    // invoke: _ZN16QAbstractSpinBoxD0Ev
    // invoke: void ~QAbstractSpinBox()
    C.C_ZN16QAbstractSpinBoxD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "~QAbstractSpinBox", args)
  }

  return
}

// stepUp()
func (this *QAbstractSpinBox) Stepup(args ...interface{}) () {
  // stepUp()
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
    // invoke: _ZN16QAbstractSpinBox6stepUpEv
    // invoke: void stepUp()
    C.C_ZN16QAbstractSpinBox6stepUpEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "stepUp", args)
  }

  return
}

// stepBy(int)
func (this *QAbstractSpinBox) Stepby(args ...interface{}) () {
  // stepBy(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox6stepByEi
    // invoke: void stepBy(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QAbstractSpinBox6stepByEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "stepBy", args)
  }

  return
}

// specialValueText()
func (this *QAbstractSpinBox) Specialvaluetext(args ...interface{}) (ret interface{}) {
  // specialValueText()
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
    // invoke: _ZNK16QAbstractSpinBox16specialValueTextEv
    // invoke: QString specialValueText()
    var ret0 = C.C_ZNK16QAbstractSpinBox16specialValueTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "specialValueText", args)
  }

  return
}

// isAccelerated()
func (this *QAbstractSpinBox) Isaccelerated(args ...interface{}) (ret interface{}) {
  // isAccelerated()
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
    // invoke: _ZNK16QAbstractSpinBox13isAcceleratedEv
    // invoke: bool isAccelerated()
    var ret0 = C.C_ZNK16QAbstractSpinBox13isAcceleratedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "isAccelerated", args)
  }

  return
}

// wrapping()
func (this *QAbstractSpinBox) Wrapping(args ...interface{}) (ret interface{}) {
  // wrapping()
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
    // invoke: _ZNK16QAbstractSpinBox8wrappingEv
    // invoke: bool wrapping()
    var ret0 = C.C_ZNK16QAbstractSpinBox8wrappingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "wrapping", args)
  }

  return
}

// sizeHint()
func (this *QAbstractSpinBox) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
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
    // invoke: _ZNK16QAbstractSpinBox8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK16QAbstractSpinBox8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "sizeHint", args)
  }

  return
}

// interpretText()
func (this *QAbstractSpinBox) Interprettext(args ...interface{}) () {
  // interpretText()
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
    // invoke: _ZN16QAbstractSpinBox13interpretTextEv
    // invoke: void interpretText()
    C.C_ZN16QAbstractSpinBox13interpretTextEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "interpretText", args)
  }

  return
}

// validate(class QString &, int &)
func (this *QAbstractSpinBox) Validate(args ...interface{}) () {
  // validate(class QString &, int &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox8validateER7QStringRi
    // invoke: QValidator::State validate(class QString &, int &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK16QAbstractSpinBox8validateER7QStringRi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "validate", args)
  }

  return
}

// isGroupSeparatorShown()
func (this *QAbstractSpinBox) Isgroupseparatorshown(args ...interface{}) (ret interface{}) {
  // isGroupSeparatorShown()
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
    // invoke: _ZNK16QAbstractSpinBox21isGroupSeparatorShownEv
    // invoke: bool isGroupSeparatorShown()
    var ret0 = C.C_ZNK16QAbstractSpinBox21isGroupSeparatorShownEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "isGroupSeparatorShown", args)
  }

  return
}

// correctionMode()
func (this *QAbstractSpinBox) Correctionmode(args ...interface{}) () {
  // correctionMode()
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
    // invoke: _ZNK16QAbstractSpinBox14correctionModeEv
    // invoke: QAbstractSpinBox::CorrectionMode correctionMode()
    C.C_ZNK16QAbstractSpinBox14correctionModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "correctionMode", args)
  }

  return
}

// metaObject()
func (this *QAbstractSpinBox) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK16QAbstractSpinBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK16QAbstractSpinBox10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "metaObject", args)
  }

  return
}

// QAbstractSpinBox(class QWidget *)
func NewQAbstractSpinBox(args ...interface{}) *QAbstractSpinBox {
  // QAbstractSpinBox(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBoxC1EP7QWidget
    // invoke: void QAbstractSpinBox(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QAbstractSpinBoxC2EP7QWidget(arg0)
    return &QAbstractSpinBox{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "QAbstractSpinBox", args)
  }

  return nil // QAbstractSpinBox{Qclsinst:qthis}
}

// stepDown()
func (this *QAbstractSpinBox) Stepdown(args ...interface{}) () {
  // stepDown()
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
    // invoke: _ZN16QAbstractSpinBox8stepDownEv
    // invoke: void stepDown()
    C.C_ZN16QAbstractSpinBox8stepDownEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "stepDown", args)
  }

  return
}

// setFrame(_Bool)
func (this *QAbstractSpinBox) Setframe(args ...interface{}) () {
  // setFrame(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox8setFrameEb
    // invoke: void setFrame(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QAbstractSpinBox8setFrameEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setFrame", args)
  }

  return
}

// clear()
func (this *QAbstractSpinBox) Clear(args ...interface{}) () {
  // clear()
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
    // invoke: _ZN16QAbstractSpinBox5clearEv
    // invoke: void clear()
    C.C_ZN16QAbstractSpinBox5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "clear", args)
  }

  return
}

// minimumSizeHint()
func (this *QAbstractSpinBox) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
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
    // invoke: _ZNK16QAbstractSpinBox15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK16QAbstractSpinBox15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "minimumSizeHint", args)
  }

  return
}

// setKeyboardTracking(_Bool)
func (this *QAbstractSpinBox) Setkeyboardtracking(args ...interface{}) () {
  // setKeyboardTracking(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox19setKeyboardTrackingEb
    // invoke: void setKeyboardTracking(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QAbstractSpinBox19setKeyboardTrackingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setKeyboardTracking", args)
  }

  return
}

// setReadOnly(_Bool)
func (this *QAbstractSpinBox) Setreadonly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox11setReadOnlyEb
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QAbstractSpinBox11setReadOnlyEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setReadOnly", args)
  }

  return
}

// <= body block end

