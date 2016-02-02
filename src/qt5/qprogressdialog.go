package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qprogressdialog.h
// dst-file: /src/widgets/qprogressdialog.go
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
  // proto:  void QProgressDialog::setLabel(QLabel * label);
extern void C_ZN15QProgressDialog8setLabelEP6QLabel(void* qthis, void* arg0); // 4
  // proto:  int QProgressDialog::minimumDuration();
extern int32_t C_ZNK15QProgressDialog15minimumDurationEv(void* qthis); // 4
  // proto:  QString QProgressDialog::labelText();
extern void* C_ZNK15QProgressDialog9labelTextEv(void* qthis); // 4
  // proto:  int QProgressDialog::minimum();
extern int32_t C_ZNK15QProgressDialog7minimumEv(void* qthis); // 4
  // proto:  void QProgressDialog::cancel();
extern void C_ZN15QProgressDialog6cancelEv(void* qthis); // 4
  // proto:  void QProgressDialog::open(QObject * receiver, const char * member);
extern void C_ZN15QProgressDialog4openEP7QObjectPKc(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QProgressDialog::setLabelText(const QString & text);
extern void C_ZN15QProgressDialog12setLabelTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QProgressDialog::autoReset();
extern bool C_ZNK15QProgressDialog9autoResetEv(void* qthis); // 4
  // proto:  bool QProgressDialog::wasCanceled();
extern bool C_ZNK15QProgressDialog11wasCanceledEv(void* qthis); // 4
  // proto:  void QProgressDialog::setCancelButtonText(const QString & text);
extern void C_ZN15QProgressDialog19setCancelButtonTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QProgressDialog::setMinimumDuration(int ms);
extern void C_ZN15QProgressDialog18setMinimumDurationEi(void* qthis, int32_t arg0); // 4
  // proto:  void QProgressDialog::setMaximum(int maximum);
extern void C_ZN15QProgressDialog10setMaximumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QProgressDialog::setBar(QProgressBar * bar);
extern void C_ZN15QProgressDialog6setBarEP12QProgressBar(void* qthis, void* arg0); // 4
  // proto:  void QProgressDialog::setValue(int progress);
extern void C_ZN15QProgressDialog8setValueEi(void* qthis, int32_t arg0); // 4
  // proto:  void QProgressDialog::setCancelButton(QPushButton * button);
extern void C_ZN15QProgressDialog15setCancelButtonEP11QPushButton(void* qthis, void* arg0); // 4
  // proto:  bool QProgressDialog::autoClose();
extern bool C_ZNK15QProgressDialog9autoCloseEv(void* qthis); // 4
  // proto:  const QMetaObject * QProgressDialog::metaObject();
extern void C_ZNK15QProgressDialog10metaObjectEv(void* qthis); // 4
  // proto:  void QProgressDialog::setRange(int minimum, int maximum);
extern void C_ZN15QProgressDialog8setRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QProgressDialog::reset();
extern void C_ZN15QProgressDialog5resetEv(void* qthis); // 4
  // proto:  QSize QProgressDialog::sizeHint();
extern void* C_ZNK15QProgressDialog8sizeHintEv(void* qthis); // 4
  // proto:  void QProgressDialog::setAutoReset(bool reset);
extern void C_ZN15QProgressDialog12setAutoResetEb(void* qthis, bool arg0); // 4
  // proto:  int QProgressDialog::maximum();
extern int32_t C_ZNK15QProgressDialog7maximumEv(void* qthis); // 4
  // proto:  int QProgressDialog::value();
extern int32_t C_ZNK15QProgressDialog5valueEv(void* qthis); // 4
  // proto:  void QProgressDialog::setAutoClose(bool close);
extern void C_ZN15QProgressDialog12setAutoCloseEb(void* qthis, bool arg0); // 4
  // proto:  void QProgressDialog::~QProgressDialog();
extern void C_ZN15QProgressDialogD2Ev(void* qthis); // 4
  // proto:  void QProgressDialog::setMinimum(int minimum);
extern void C_ZN15QProgressDialog10setMinimumEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QProgressDialog)=1
type QProgressDialog struct {
  /*qbase*/ QDialog;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _canceled QProgressDialog_canceled_signal;
}

// setLabel(class QLabel *)
func (this *QProgressDialog) Setlabel(args ...interface{}) () {
  // setLabel(class QLabel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLabel{}) // "QLabel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog8setLabelEP6QLabel
    // invoke: void setLabel(class QLabel *)
    var arg0 = args[0].(*QLabel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog8setLabelEP6QLabel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setLabel", args)
  }

  return
}

// minimumDuration()
func (this *QProgressDialog) Minimumduration(args ...interface{}) (ret interface{}) {
  // minimumDuration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog15minimumDurationEv
    // invoke: int minimumDuration()
    var ret0 = C.C_ZNK15QProgressDialog15minimumDurationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressDialog", "minimumDuration", args)
  }

  return
}

// labelText()
func (this *QProgressDialog) Labeltext(args ...interface{}) (ret interface{}) {
  // labelText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9labelTextEv
    // invoke: QString labelText()
    var ret0 = C.C_ZNK15QProgressDialog9labelTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressDialog", "labelText", args)
  }

  return
}

// minimum()
func (this *QProgressDialog) Minimum(args ...interface{}) (ret interface{}) {
  // minimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog7minimumEv
    // invoke: int minimum()
    var ret0 = C.C_ZNK15QProgressDialog7minimumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressDialog", "minimum", args)
  }

  return
}

// cancel()
func (this *QProgressDialog) Cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog6cancelEv
    // invoke: void cancel()
    C.C_ZN15QProgressDialog6cancelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "cancel", args)
  }

  return
}

// open(class QObject *, const char *)
func (this *QProgressDialog) Open(args ...interface{}) () {
  // open(class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog4openEP7QObjectPKc
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN15QProgressDialog4openEP7QObjectPKc(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QProgressDialog", "open", args)
  }

  return
}

// setLabelText(const class QString &)
func (this *QProgressDialog) Setlabeltext(args ...interface{}) () {
  // setLabelText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog12setLabelTextERK7QString
    // invoke: void setLabelText(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog12setLabelTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setLabelText", args)
  }

  return
}

// autoReset()
func (this *QProgressDialog) Autoreset(args ...interface{}) (ret interface{}) {
  // autoReset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9autoResetEv
    // invoke: bool autoReset()
    var ret0 = C.C_ZNK15QProgressDialog9autoResetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressDialog", "autoReset", args)
  }

  return
}

// wasCanceled()
func (this *QProgressDialog) Wascanceled(args ...interface{}) (ret interface{}) {
  // wasCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog11wasCanceledEv
    // invoke: bool wasCanceled()
    var ret0 = C.C_ZNK15QProgressDialog11wasCanceledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressDialog", "wasCanceled", args)
  }

  return
}

// setCancelButtonText(const class QString &)
func (this *QProgressDialog) Setcancelbuttontext(args ...interface{}) () {
  // setCancelButtonText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog19setCancelButtonTextERK7QString
    // invoke: void setCancelButtonText(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog19setCancelButtonTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setCancelButtonText", args)
  }

  return
}

// setMinimumDuration(int)
func (this *QProgressDialog) Setminimumduration(args ...interface{}) () {
  // setMinimumDuration(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog18setMinimumDurationEi
    // invoke: void setMinimumDuration(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog18setMinimumDurationEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMinimumDuration", args)
  }

  return
}

// setMaximum(int)
func (this *QProgressDialog) Setmaximum(args ...interface{}) () {
  // setMaximum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog10setMaximumEi
    // invoke: void setMaximum(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog10setMaximumEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMaximum", args)
  }

  return
}

// setBar(class QProgressBar *)
func (this *QProgressDialog) Setbar(args ...interface{}) () {
  // setBar(class QProgressBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QProgressBar{}) // "QProgressBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog6setBarEP12QProgressBar
    // invoke: void setBar(class QProgressBar *)
    var arg0 = args[0].(*QProgressBar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog6setBarEP12QProgressBar(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setBar", args)
  }

  return
}

// setValue(int)
func (this *QProgressDialog) Setvalue(args ...interface{}) () {
  // setValue(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog8setValueEi
    // invoke: void setValue(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog8setValueEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setValue", args)
  }

  return
}

// setCancelButton(class QPushButton *)
func (this *QProgressDialog) Setcancelbutton(args ...interface{}) () {
  // setCancelButton(class QPushButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPushButton{}) // "QPushButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog15setCancelButtonEP11QPushButton
    // invoke: void setCancelButton(class QPushButton *)
    var arg0 = args[0].(*QPushButton).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog15setCancelButtonEP11QPushButton(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setCancelButton", args)
  }

  return
}

// autoClose()
func (this *QProgressDialog) Autoclose(args ...interface{}) (ret interface{}) {
  // autoClose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9autoCloseEv
    // invoke: bool autoClose()
    var ret0 = C.C_ZNK15QProgressDialog9autoCloseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressDialog", "autoClose", args)
  }

  return
}

// metaObject()
func (this *QProgressDialog) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QProgressDialog10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "metaObject", args)
  }

  return
}

// setRange(int, int)
func (this *QProgressDialog) Setrange(args ...interface{}) () {
  // setRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog8setRangeEii
    // invoke: void setRange(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN15QProgressDialog8setRangeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setRange", args)
  }

  return
}

// reset()
func (this *QProgressDialog) Reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog5resetEv
    // invoke: void reset()
    C.C_ZN15QProgressDialog5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "reset", args)
  }

  return
}

// sizeHint()
func (this *QProgressDialog) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK15QProgressDialog8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressDialog", "sizeHint", args)
  }

  return
}

// setAutoReset(_Bool)
func (this *QProgressDialog) Setautoreset(args ...interface{}) () {
  // setAutoReset(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog12setAutoResetEb
    // invoke: void setAutoReset(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog12setAutoResetEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setAutoReset", args)
  }

  return
}

// maximum()
func (this *QProgressDialog) Maximum(args ...interface{}) (ret interface{}) {
  // maximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog7maximumEv
    // invoke: int maximum()
    var ret0 = C.C_ZNK15QProgressDialog7maximumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressDialog", "maximum", args)
  }

  return
}

// value()
func (this *QProgressDialog) Value(args ...interface{}) (ret interface{}) {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog5valueEv
    // invoke: int value()
    var ret0 = C.C_ZNK15QProgressDialog5valueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressDialog", "value", args)
  }

  return
}

// setAutoClose(_Bool)
func (this *QProgressDialog) Setautoclose(args ...interface{}) () {
  // setAutoClose(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog12setAutoCloseEb
    // invoke: void setAutoClose(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog12setAutoCloseEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setAutoClose", args)
  }

  return
}

// ~QProgressDialog()
func (this *QProgressDialog) Freeqprogressdialog(args ...interface{}) () {
  // ~QProgressDialog()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialogD0Ev
    // invoke: void ~QProgressDialog()
    C.C_ZN15QProgressDialogD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressDialog", "~QProgressDialog", args)
  }

  return
}

// setMinimum(int)
func (this *QProgressDialog) Setminimum(args ...interface{}) () {
  // setMinimum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog10setMinimumEi
    // invoke: void setMinimum(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QProgressDialog10setMinimumEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMinimum", args)
  }

  return
}

// <= body block end

