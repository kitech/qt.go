package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qmessagebox.h
// dst-file: /src/widgets/qmessagebox.go
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
  // proto: static int QMessageBox::critical(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern void _ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int arg3, int arg4, int arg5);
  // proto:  void QMessageBox::setButtonText(int button, const QString & text);
extern void _ZN11QMessageBox13setButtonTextEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  void QMessageBox::~QMessageBox();
extern void _ZN11QMessageBoxD0Ev(void* qthis);
  // proto:  void QMessageBox::setText(const QString & text);
extern void _ZN11QMessageBox7setTextERK7QString(void* qthis, void* arg0);
  // proto:  void QMessageBox::setIconPixmap(const QPixmap & pixmap);
extern void _ZN11QMessageBox13setIconPixmapERK7QPixmap(void* qthis, void* arg0);
  // proto: static void QMessageBox::about(QWidget * parent, const QString & title, const QString & text);
extern void _ZN11QMessageBox5aboutEP7QWidgetRK7QStringS4_(void* arg0, void* arg1, void* arg2);
  // proto:  QString QMessageBox::text();
extern void _ZNK11QMessageBox4textEv(void* qthis);
  // proto: static int QMessageBox::question(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern void _ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int arg6, int arg7);
  // proto: static int QMessageBox::warning(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern void _ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int arg3, int arg4, int arg5);
  // proto:  void QMessageBox::QMessageBox(const QMessageBox & );
extern void* dector_ZN11QMessageBoxC1ERKS_(void* arg0);
extern void _ZN11QMessageBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QMessageBox::metaObject();
extern void _ZNK11QMessageBox10metaObjectEv(void* qthis);
  // proto: static int QMessageBox::question(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern void _ZN11QMessageBox8questionEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int arg3, int arg4, int arg5);
  // proto:  QPushButton * QMessageBox::defaultButton();
extern void _ZNK11QMessageBox13defaultButtonEv(void* qthis);
  // proto:  void QMessageBox::open(QObject * receiver, const char * member);
extern void _ZN11QMessageBox4openEP7QObjectPKc(void* qthis, void* arg0, char* arg1);
  // proto:  QList<QAbstractButton *> QMessageBox::buttons();
extern void _ZNK11QMessageBox7buttonsEv(void* qthis);
  // proto: static void QMessageBox::aboutQt(QWidget * parent, const QString & title);
extern void _ZN11QMessageBox7aboutQtEP7QWidgetRK7QString(void* arg0, void* arg1);
  // proto:  QString QMessageBox::informativeText();
extern void _ZNK11QMessageBox15informativeTextEv(void* qthis);
  // proto:  void QMessageBox::setInformativeText(const QString & text);
extern void _ZN11QMessageBox18setInformativeTextERK7QString(void* qthis, void* arg0);
  // proto:  void QMessageBox::setDetailedText(const QString & text);
extern void _ZN11QMessageBox15setDetailedTextERK7QString(void* qthis, void* arg0);
  // proto:  void QMessageBox::QMessageBox(QWidget * parent);
extern void* dector_ZN11QMessageBoxC1EP7QWidget(void* arg0);
extern void _ZN11QMessageBoxC1EP7QWidget(void* qthis, void* arg0);
  // proto: static int QMessageBox::critical(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern void _ZN11QMessageBox8criticalEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int arg6, int arg7);
  // proto:  QAbstractButton * QMessageBox::clickedButton();
extern void _ZNK11QMessageBox13clickedButtonEv(void* qthis);
  // proto:  void QMessageBox::setDefaultButton(QPushButton * button);
extern void _ZN11QMessageBox16setDefaultButtonEP11QPushButton(void* qthis, void* arg0);
  // proto: static int QMessageBox::warning(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern void _ZN11QMessageBox7warningEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int arg6, int arg7);
  // proto:  void QMessageBox::setEscapeButton(QAbstractButton * button);
extern void _ZN11QMessageBox15setEscapeButtonEP15QAbstractButton(void* qthis, void* arg0);
  // proto: static int QMessageBox::information(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
extern void _ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_iii(void* arg0, void* arg1, void* arg2, int arg3, int arg4, int arg5);
  // proto:  void QMessageBox::setCheckBox(QCheckBox * cb);
extern void _ZN11QMessageBox11setCheckBoxEP9QCheckBox(void* qthis, void* arg0);
  // proto:  void QMessageBox::setWindowTitle(const QString & title);
extern void _ZN11QMessageBox14setWindowTitleERK7QString(void* qthis, void* arg0);
  // proto:  QAbstractButton * QMessageBox::escapeButton();
extern void _ZNK11QMessageBox12escapeButtonEv(void* qthis);
  // proto:  QPixmap QMessageBox::iconPixmap();
extern void _ZNK11QMessageBox10iconPixmapEv(void* qthis);
  // proto:  void QMessageBox::removeButton(QAbstractButton * button);
extern void _ZN11QMessageBox12removeButtonEP15QAbstractButton(void* qthis, void* arg0);
  // proto:  QString QMessageBox::detailedText();
extern void _ZNK11QMessageBox12detailedTextEv(void* qthis);
  // proto:  QCheckBox * QMessageBox::checkBox();
extern void _ZNK11QMessageBox8checkBoxEv(void* qthis);
  // proto:  QString QMessageBox::buttonText(int button);
extern void _ZNK11QMessageBox10buttonTextEi(void* qthis, int arg0);
  // proto: static int QMessageBox::information(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
extern void _ZN11QMessageBox11informationEP7QWidgetRK7QStringS4_S4_S4_S4_ii(void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, int arg6, int arg7);
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

// class sizeof(QMessageBox)=1
type QMessageBox struct {
  /*qbase*/ QDialog;
  qclsinst uint64 /* *mut c_void*/;
//  _buttonClicked QMessageBox_buttonClicked_signal;
}

  // proto: static int QMessageBox::critical(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
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

  // proto:  void QMessageBox::setButtonText(int button, const QString & text);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QMessageBox", "setButtonText", args)
  }

}

  // proto:  void QMessageBox::~QMessageBox();
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

  // proto:  void QMessageBox::setText(const QString & text);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageBox", "setText", args)
  }

}

  // proto:  void QMessageBox::setIconPixmap(const QPixmap & pixmap);
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
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageBox", "setIconPixmap", args)
  }

}

  // proto: static void QMessageBox::about(QWidget * parent, const QString & title, const QString & text);
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

  // proto:  QString QMessageBox::text();
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

  // proto: static int QMessageBox::question(QWidget * parent, const QString & title, const QString & text, const QString & button0Text, const QString & button1Text, const QString & button2Text, int defaultButtonNumber, int escapeButtonNumber);
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

  // proto: static int QMessageBox::warning(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
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

  // proto:  void QMessageBox::QMessageBox(const QMessageBox & );
func NewQMessageBox(args ...interface{}) QMessageBox {
  return QMessageBox{}
}

  // proto:  const QMetaObject * QMessageBox::metaObject();
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

  // proto:  QPushButton * QMessageBox::defaultButton();
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

  // proto:  void QMessageBox::open(QObject * receiver, const char * member);
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
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QMessageBox", "open", args)
  }

}

  // proto:  QList<QAbstractButton *> QMessageBox::buttons();
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

  // proto: static void QMessageBox::aboutQt(QWidget * parent, const QString & title);
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

  // proto:  QString QMessageBox::informativeText();
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

  // proto:  void QMessageBox::setInformativeText(const QString & text);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageBox", "setInformativeText", args)
  }

}

  // proto:  void QMessageBox::setDetailedText(const QString & text);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageBox", "setDetailedText", args)
  }

}

  // proto:  QAbstractButton * QMessageBox::clickedButton();
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

  // proto:  void QMessageBox::setDefaultButton(QPushButton * button);
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
    var arg0 = C.int(args[0].(int32))
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN11QMessageBox16setDefaultButtonEP11QPushButton
    var arg0 = args[0].(QPushButton).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageBox", "setDefaultButton", args)
  }

}

  // proto:  void QMessageBox::setEscapeButton(QAbstractButton * button);
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
    var arg0 = args[0].(QAbstractButton).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN11QMessageBox15setEscapeButtonENS_14StandardButtonE
    var arg0 = C.int(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageBox", "setEscapeButton", args)
  }

}

  // proto: static int QMessageBox::information(QWidget * parent, const QString & title, const QString & text, int button0, int button1, int button2);
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

  // proto:  void QMessageBox::setCheckBox(QCheckBox * cb);
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
    var arg0 = args[0].(QCheckBox).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageBox", "setCheckBox", args)
  }

}

  // proto:  void QMessageBox::setWindowTitle(const QString & title);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageBox", "setWindowTitle", args)
  }

}

  // proto:  QAbstractButton * QMessageBox::escapeButton();
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

  // proto:  QPixmap QMessageBox::iconPixmap();
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

  // proto:  void QMessageBox::removeButton(QAbstractButton * button);
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
    var arg0 = args[0].(QAbstractButton).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageBox", "removeButton", args)
  }

}

  // proto:  QString QMessageBox::detailedText();
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

  // proto:  QCheckBox * QMessageBox::checkBox();
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

  // proto:  QString QMessageBox::buttonText(int button);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageBox", "buttonText", args)
  }

}

// <= body block end

