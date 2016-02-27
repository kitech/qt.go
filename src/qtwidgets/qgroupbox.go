package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qgroupbox.h
// dst-file: /src/widgets/qgroupbox.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
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
  // proto:  void QGroupBox::~QGroupBox();
extern void C_ZN9QGroupBoxD2Ev(void* qthis); // 4
  // proto:  void QGroupBox::setTitle(const QString & title);
extern void C_ZN9QGroupBox8setTitleERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QGroupBox::isCheckable();
extern bool C_ZNK9QGroupBox11isCheckableEv(void* qthis); // 4
  // proto:  Qt::Alignment QGroupBox::alignment();
extern void C_ZNK9QGroupBox9alignmentEv(void* qthis); // 4
  // proto:  void QGroupBox::setFlat(bool flat);
extern void C_ZN9QGroupBox7setFlatEb(void* qthis, bool arg0); // 4
  // proto:  void QGroupBox::QGroupBox(QWidget * parent);
extern void* C_ZN9QGroupBoxC2EP7QWidget(void* arg0); // 3
  // proto:  void QGroupBox::QGroupBox(const QString & title, QWidget * parent);
extern void* C_ZN9QGroupBoxC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
  // proto:  QString QGroupBox::title();
extern void* C_ZNK9QGroupBox5titleEv(void* qthis); // 4
  // proto:  bool QGroupBox::isChecked();
extern bool C_ZNK9QGroupBox9isCheckedEv(void* qthis); // 4
  // proto:  void QGroupBox::setAlignment(int alignment);
extern void C_ZN9QGroupBox12setAlignmentEi(void* qthis, int32_t arg0); // 4
  // proto:  void QGroupBox::setChecked(bool checked);
extern void C_ZN9QGroupBox10setCheckedEb(void* qthis, bool arg0); // 4
  // proto:  void QGroupBox::setCheckable(bool checkable);
extern void C_ZN9QGroupBox12setCheckableEb(void* qthis, bool arg0); // 4
  // proto:  bool QGroupBox::isFlat();
extern bool C_ZNK9QGroupBox6isFlatEv(void* qthis); // 4
  // proto:  const QMetaObject * QGroupBox::metaObject();
extern void C_ZNK9QGroupBox10metaObjectEv(void* qthis); // 4
  // proto:  QSize QGroupBox::minimumSizeHint();
extern void* C_ZNK9QGroupBox15minimumSizeHintEv(void* qthis); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QGroupBox)=1
type QGroupBox struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _clicked QGroupBox_clicked_signal;
//  _toggled QGroupBox_toggled_signal;
}

// ~QGroupBox()
func (this *QGroupBox) Free(args ...interface{}) () {
  // ~QGroupBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBoxD0Ev
    // invoke: void ~QGroupBox()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN9QGroupBoxD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QGroupBox", "~QGroupBox", args)
  }

  return
}

// setTitle(const class QString &)
func (this *QGroupBox) SetTitle(args ...interface{}) () {
  // setTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBox8setTitleERK7QString
    // invoke: void setTitle(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QGroupBox8setTitleERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setTitle", args)
  }

  return
}

// isCheckable()
func (this *QGroupBox) IsCheckable(args ...interface{}) (ret interface{}) {
  // isCheckable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox11isCheckableEv
    // invoke: bool isCheckable()
    var ret0 = C.C_ZNK9QGroupBox11isCheckableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGroupBox", "isCheckable", args)
  }

  return
}

// alignment()
func (this *QGroupBox) Alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK9QGroupBox9alignmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "alignment", args)
  }

  return
}

// setFlat(_Bool)
func (this *QGroupBox) SetFlat(args ...interface{}) () {
  // setFlat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBox7setFlatEb
    // invoke: void setFlat(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QGroupBox7setFlatEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setFlat", args)
  }

  return
}

// QGroupBox(class QWidget *)
func GcfreeQGroupBox(this *QGroupBox) {
  qtrt.UniverseFree(this)
}
func NewQGroupBox(args ...interface{}) *QGroupBox {
  // QGroupBox(class QWidget *)
  // QGroupBox(const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBoxC1EP7QWidget
    // invoke: void QGroupBox(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QGroupBoxC2EP7QWidget(arg0)
    this := &QGroupBox{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQGroupBox)
    return this
  case 1:
    // invoke: _ZN9QGroupBoxC1ERK7QStringP7QWidget
    // invoke: void QGroupBox(const class QString &, class QWidget *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QGroupBoxC2ERK7QStringP7QWidget(arg0, arg1)
    this := &QGroupBox{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQGroupBox)
    return this
  default:
    qtrt.ErrorResolve("QGroupBox", "QGroupBox", args)
  }

  return nil // QGroupBox{Qclsinst:qthis}
}

// title()
func (this *QGroupBox) Title(args ...interface{}) (ret interface{}) {
  // title()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox5titleEv
    // invoke: QString title()
    var ret0 = C.C_ZNK9QGroupBox5titleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGroupBox", "title", args)
  }

  return
}

// isChecked()
func (this *QGroupBox) IsChecked(args ...interface{}) (ret interface{}) {
  // isChecked()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox9isCheckedEv
    // invoke: bool isChecked()
    var ret0 = C.C_ZNK9QGroupBox9isCheckedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGroupBox", "isChecked", args)
  }

  return
}

// setAlignment(int)
func (this *QGroupBox) SetAlignment(args ...interface{}) () {
  // setAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBox12setAlignmentEi
    // invoke: void setAlignment(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QGroupBox12setAlignmentEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setAlignment", args)
  }

  return
}

// setChecked(_Bool)
func (this *QGroupBox) SetChecked(args ...interface{}) () {
  // setChecked(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBox10setCheckedEb
    // invoke: void setChecked(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QGroupBox10setCheckedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setChecked", args)
  }

  return
}

// setCheckable(_Bool)
func (this *QGroupBox) SetCheckable(args ...interface{}) () {
  // setCheckable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGroupBox12setCheckableEb
    // invoke: void setCheckable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QGroupBox12setCheckableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QGroupBox", "setCheckable", args)
  }

  return
}

// isFlat()
func (this *QGroupBox) IsFlat(args ...interface{}) (ret interface{}) {
  // isFlat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox6isFlatEv
    // invoke: bool isFlat()
    var ret0 = C.C_ZNK9QGroupBox6isFlatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGroupBox", "isFlat", args)
  }

  return
}

// metaObject()
func (this *QGroupBox) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QGroupBox10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGroupBox", "metaObject", args)
  }

  return
}

// minimumSizeHint()
func (this *QGroupBox) MinimumSizeHint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGroupBox15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK9QGroupBox15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGroupBox", "minimumSizeHint", args)
  }

  return
}

// <= body block end

