package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qprogressdialog.h
// dst-file: /src/widgets/qprogressdialog.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QProgressDialog)=1
type QProgressDialog struct {
  /*qbase*/ QDialog;
  qclsinst uint64 /* *mut c_void*/;
//  _canceled QProgressDialog_canceled_signal;
}


func (this *QProgressDialog) setAutoClose(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setAutoClose", args)
 }

}


func (this *QProgressDialog) open(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "open", args)
 }

}


func (this *QProgressDialog) setMaximum(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMaximum", args)
 }

}


func (this *QProgressDialog) setMinimum(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMinimum", args)
 }

}


func (this *QProgressDialog) setLabelText(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setLabelText", args)
 }

}


func (this *QProgressDialog) wasCanceled(args ...interface{}) () {
  // wasCanceled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog11wasCanceledEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "wasCanceled", args)
 }

}


func (this *QProgressDialog) FreeQProgressDialog(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QProgressDialog", "~QProgressDialog", args)
 }

}


func (this *QProgressDialog) minimumDuration(args ...interface{}) () {
  // minimumDuration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog15minimumDurationEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "minimumDuration", args)
 }

}


func (this *QProgressDialog) setMinimumDuration(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setMinimumDuration", args)
 }

}


func (this *QProgressDialog) maximum(args ...interface{}) () {
  // maximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog7maximumEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "maximum", args)
 }

}


func (this *QProgressDialog) setBar(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setBar", args)
 }

}


func (this *QProgressDialog) cancel(args ...interface{}) () {
  // cancel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog6cancelEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "cancel", args)
 }

}


func (this *QProgressDialog) autoClose(args ...interface{}) () {
  // autoClose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9autoCloseEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "autoClose", args)
 }

}


func (this *QProgressDialog) minimum(args ...interface{}) () {
  // minimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog7minimumEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "minimum", args)
 }

}


func (this *QProgressDialog) autoReset(args ...interface{}) () {
  // autoReset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9autoResetEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "autoReset", args)
 }

}


func (this *QProgressDialog) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QProgressDialog5resetEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "reset", args)
 }

}


func NewQProgressDialog(args ...interface{})() {
}


func (this *QProgressDialog) setRange(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setRange", args)
 }

}


func (this *QProgressDialog) setCancelButtonText(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setCancelButtonText", args)
 }

}


func (this *QProgressDialog) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog8sizeHintEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "sizeHint", args)
 }

}


func (this *QProgressDialog) labelText(args ...interface{}) () {
  // labelText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog9labelTextEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "labelText", args)
 }

}


func (this *QProgressDialog) setLabel(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setLabel", args)
 }

}


func (this *QProgressDialog) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog10metaObjectEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "metaObject", args)
 }

}


func (this *QProgressDialog) setAutoReset(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setAutoReset", args)
 }

}


func (this *QProgressDialog) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QProgressDialog5valueEv
  default:
    qtrt.ErrorResolve("QProgressDialog", "value", args)
 }

}


func (this *QProgressDialog) setCancelButton(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setCancelButton", args)
 }

}


func (this *QProgressDialog) setValue(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QProgressDialog", "setValue", args)
 }

}

// <= body block end

