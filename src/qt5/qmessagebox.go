package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qmessagebox.h
// dst-file: /src/widgets/qmessagebox.go
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
// class sizeof(QMessageBox)=1
type QMessageBox struct {
  /*qbase*/ QDialog;
  qclsinst uint64 /* *mut c_void*/;
//  _buttonClicked QMessageBox_buttonClicked_signal;
}


func (this *QMessageBox) critical_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMessageBox", "critical", args)
  }

}


func (this *QMessageBox) setButtonText(args ...interface{}) () {
  // setButtonText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox13setButtonTextEiRK7QString
  default:
    qtrt.ErrorResolve("QMessageBox", "setButtonText", args)
  }

}


func (this *QMessageBox) FreeQMessageBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMessageBox", "~QMessageBox", args)
  }

}


func (this *QMessageBox) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox7setTextERK7QString
  default:
    qtrt.ErrorResolve("QMessageBox", "setText", args)
  }

}


func (this *QMessageBox) setIconPixmap(args ...interface{}) () {
  // setIconPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox13setIconPixmapERK7QPixmap
  default:
    qtrt.ErrorResolve("QMessageBox", "setIconPixmap", args)
  }

}


func (this *QMessageBox) about_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMessageBox", "about", args)
  }

}


func (this *QMessageBox) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox4textEv
  default:
    qtrt.ErrorResolve("QMessageBox", "text", args)
  }

}


func (this *QMessageBox) question_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMessageBox", "question", args)
  }

}


func (this *QMessageBox) warning_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMessageBox", "warning", args)
  }

}


func NewQMessageBox(args ...interface{}) QMessageBox {
  return QMessageBox{}
}


func (this *QMessageBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox10metaObjectEv
  default:
    qtrt.ErrorResolve("QMessageBox", "metaObject", args)
  }

}


func (this *QMessageBox) defaultButton(args ...interface{}) () {
  // defaultButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox13defaultButtonEv
  default:
    qtrt.ErrorResolve("QMessageBox", "defaultButton", args)
  }

}


func (this *QMessageBox) open(args ...interface{}) () {
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
    // invoke: _ZN11QMessageBox4openEP7QObjectPKc
  default:
    qtrt.ErrorResolve("QMessageBox", "open", args)
  }

}


func (this *QMessageBox) buttons(args ...interface{}) () {
  // buttons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox7buttonsEv
  default:
    qtrt.ErrorResolve("QMessageBox", "buttons", args)
  }

}


func (this *QMessageBox) aboutQt_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMessageBox", "aboutQt", args)
  }

}


func (this *QMessageBox) informativeText(args ...interface{}) () {
  // informativeText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox15informativeTextEv
  default:
    qtrt.ErrorResolve("QMessageBox", "informativeText", args)
  }

}


func (this *QMessageBox) setInformativeText(args ...interface{}) () {
  // setInformativeText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox18setInformativeTextERK7QString
  default:
    qtrt.ErrorResolve("QMessageBox", "setInformativeText", args)
  }

}


func (this *QMessageBox) setDetailedText(args ...interface{}) () {
  // setDetailedText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox15setDetailedTextERK7QString
  default:
    qtrt.ErrorResolve("QMessageBox", "setDetailedText", args)
  }

}


func (this *QMessageBox) clickedButton(args ...interface{}) () {
  // clickedButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox13clickedButtonEv
  default:
    qtrt.ErrorResolve("QMessageBox", "clickedButton", args)
  }

}


func (this *QMessageBox) setDefaultButton(args ...interface{}) () {
  // setDefaultButton(enum QMessageBox::StandardButton)
  // setDefaultButton(class QPushButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QMessageBox::StandardButton"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPushButton{}) // "QPushButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox16setDefaultButtonENS_14StandardButtonE
  case 1:
    // invoke: _ZN11QMessageBox16setDefaultButtonEP11QPushButton
  default:
    qtrt.ErrorResolve("QMessageBox", "setDefaultButton", args)
  }

}


func (this *QMessageBox) setEscapeButton(args ...interface{}) () {
  // setEscapeButton(class QAbstractButton *)
  // setEscapeButton(enum QMessageBox::StandardButton)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QMessageBox::StandardButton"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox15setEscapeButtonEP15QAbstractButton
  case 1:
    // invoke: _ZN11QMessageBox15setEscapeButtonENS_14StandardButtonE
  default:
    qtrt.ErrorResolve("QMessageBox", "setEscapeButton", args)
  }

}


func (this *QMessageBox) information_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMessageBox", "information", args)
  }

}


func (this *QMessageBox) setCheckBox(args ...interface{}) () {
  // setCheckBox(class QCheckBox *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCheckBox{}) // "QCheckBox *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox11setCheckBoxEP9QCheckBox
  default:
    qtrt.ErrorResolve("QMessageBox", "setCheckBox", args)
  }

}


func (this *QMessageBox) setWindowTitle(args ...interface{}) () {
  // setWindowTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox14setWindowTitleERK7QString
  default:
    qtrt.ErrorResolve("QMessageBox", "setWindowTitle", args)
  }

}


func (this *QMessageBox) escapeButton(args ...interface{}) () {
  // escapeButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox12escapeButtonEv
  default:
    qtrt.ErrorResolve("QMessageBox", "escapeButton", args)
  }

}


func (this *QMessageBox) iconPixmap(args ...interface{}) () {
  // iconPixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox10iconPixmapEv
  default:
    qtrt.ErrorResolve("QMessageBox", "iconPixmap", args)
  }

}


func (this *QMessageBox) removeButton(args ...interface{}) () {
  // removeButton(class QAbstractButton *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractButton{}) // "QAbstractButton *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QMessageBox12removeButtonEP15QAbstractButton
  default:
    qtrt.ErrorResolve("QMessageBox", "removeButton", args)
  }

}


func (this *QMessageBox) detailedText(args ...interface{}) () {
  // detailedText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox12detailedTextEv
  default:
    qtrt.ErrorResolve("QMessageBox", "detailedText", args)
  }

}


func (this *QMessageBox) checkBox(args ...interface{}) () {
  // checkBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox8checkBoxEv
  default:
    qtrt.ErrorResolve("QMessageBox", "checkBox", args)
  }

}


func (this *QMessageBox) buttonText(args ...interface{}) () {
  // buttonText(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QMessageBox10buttonTextEi
  default:
    qtrt.ErrorResolve("QMessageBox", "buttonText", args)
  }

}

// <= body block end

