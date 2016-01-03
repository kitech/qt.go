package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtCore/qcommandlineoption.h
// dst-file: /src/core/qcommandlineoption.go
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
  // proto:  void QCommandLineOption::setValueName(const QString & name);
extern void _ZN18QCommandLineOption12setValueNameERK7QString(void* qthis, void* arg0);
  // proto:  QStringList QCommandLineOption::names();
extern void _ZNK18QCommandLineOption5namesEv(void* qthis);
  // proto:  void QCommandLineOption::QCommandLineOption(const QCommandLineOption & other);
extern void* dector_ZN18QCommandLineOptionC1ERKS_(void* arg0);
extern void _ZN18QCommandLineOptionC1ERKS_(void* qthis, void* arg0);
  // proto:  void QCommandLineOption::setDescription(const QString & description);
extern void _ZN18QCommandLineOption14setDescriptionERK7QString(void* qthis, void* arg0);
  // proto:  void QCommandLineOption::QCommandLineOption(const QString & name, const QString & description, const QString & valueName, const QString & defaultValue);
extern void* dector_ZN18QCommandLineOptionC1ERK7QStringS2_S2_S2_(void* arg0, void* arg1, void* arg2, void* arg3);
extern void _ZN18QCommandLineOptionC1ERK7QStringS2_S2_S2_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3);
  // proto:  QString QCommandLineOption::valueName();
extern void _ZNK18QCommandLineOption9valueNameEv(void* qthis);
  // proto:  void QCommandLineOption::QCommandLineOption(const QStringList & names, const QString & description, const QString & valueName, const QString & defaultValue);
extern void* dector_ZN18QCommandLineOptionC1ERK11QStringListRK7QStringS5_S5_(void* arg0, void* arg1, void* arg2, void* arg3);
extern void _ZN18QCommandLineOptionC1ERK11QStringListRK7QStringS5_S5_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3);
  // proto:  void QCommandLineOption::swap(QCommandLineOption & other);
extern void demth_ZN18QCommandLineOption4swapERS_(void* qthis, void* arg0);
  // proto:  QString QCommandLineOption::description();
extern void _ZNK18QCommandLineOption11descriptionEv(void* qthis);
  // proto:  void QCommandLineOption::~QCommandLineOption();
extern void _ZN18QCommandLineOptionD0Ev(void* qthis);
  // proto:  void QCommandLineOption::QCommandLineOption(const QStringList & names);
extern void* dector_ZN18QCommandLineOptionC1ERK11QStringList(void* arg0);
extern void _ZN18QCommandLineOptionC1ERK11QStringList(void* qthis, void* arg0);
  // proto:  void QCommandLineOption::setDefaultValue(const QString & defaultValue);
extern void _ZN18QCommandLineOption15setDefaultValueERK7QString(void* qthis, void* arg0);
  // proto:  void QCommandLineOption::setDefaultValues(const QStringList & defaultValues);
extern void _ZN18QCommandLineOption16setDefaultValuesERK11QStringList(void* qthis, void* arg0);
  // proto:  void QCommandLineOption::QCommandLineOption(const QString & name);
extern void* dector_ZN18QCommandLineOptionC1ERK7QString(void* arg0);
extern void _ZN18QCommandLineOptionC1ERK7QString(void* qthis, void* arg0);
  // proto:  QStringList QCommandLineOption::defaultValues();
extern void _ZNK18QCommandLineOption13defaultValuesEv(void* qthis);
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

// class sizeof(QCommandLineOption)=1
type QCommandLineOption struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QCommandLineOption::setValueName(const QString & name);
func (this *QCommandLineOption) setValueName(args ...interface{}) () {
  // setValueName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOption12setValueNameERK7QString
    // invoke: void setValueName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLineOption12setValueNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setValueName", args)
  }

}

  // proto:  QStringList QCommandLineOption::names();
func (this *QCommandLineOption) names(args ...interface{}) () {
  // names()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineOption5namesEv
    // invoke: QStringList names()
    C._ZNK18QCommandLineOption5namesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "names", args)
  }

}

  // proto:  void QCommandLineOption::QCommandLineOption(const QCommandLineOption & other);
func NewQCommandLineOption(args ...interface{}) QCommandLineOption {
  return QCommandLineOption{}
}

  // proto:  void QCommandLineOption::setDescription(const QString & description);
func (this *QCommandLineOption) setDescription(args ...interface{}) () {
  // setDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOption14setDescriptionERK7QString
    // invoke: void setDescription(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLineOption14setDescriptionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setDescription", args)
  }

}

  // proto:  QString QCommandLineOption::valueName();
func (this *QCommandLineOption) valueName(args ...interface{}) () {
  // valueName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineOption9valueNameEv
    // invoke: QString valueName()
    C._ZNK18QCommandLineOption9valueNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "valueName", args)
  }

}

  // proto:  void QCommandLineOption::swap(QCommandLineOption & other);
func (this *QCommandLineOption) swap(args ...interface{}) () {
  // swap(class QCommandLineOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCommandLineOption{}) // "QCommandLineOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOption4swapERS_
    // invoke: void swap(class QCommandLineOption &)
    var arg0 = args[0].(QCommandLineOption).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN18QCommandLineOption4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "swap", args)
  }

}

  // proto:  QString QCommandLineOption::description();
func (this *QCommandLineOption) description(args ...interface{}) () {
  // description()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineOption11descriptionEv
    // invoke: QString description()
    C._ZNK18QCommandLineOption11descriptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "description", args)
  }

}

  // proto:  void QCommandLineOption::~QCommandLineOption();
func (this *QCommandLineOption) FreeQCommandLineOption(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCommandLineOption", "~QCommandLineOption", args)
  }

}

  // proto:  void QCommandLineOption::setDefaultValue(const QString & defaultValue);
func (this *QCommandLineOption) setDefaultValue(args ...interface{}) () {
  // setDefaultValue(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOption15setDefaultValueERK7QString
    // invoke: void setDefaultValue(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLineOption15setDefaultValueERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setDefaultValue", args)
  }

}

  // proto:  void QCommandLineOption::setDefaultValues(const QStringList & defaultValues);
func (this *QCommandLineOption) setDefaultValues(args ...interface{}) () {
  // setDefaultValues(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOption16setDefaultValuesERK11QStringList
    // invoke: void setDefaultValues(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLineOption16setDefaultValuesERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setDefaultValues", args)
  }

}

  // proto:  QStringList QCommandLineOption::defaultValues();
func (this *QCommandLineOption) defaultValues(args ...interface{}) () {
  // defaultValues()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineOption13defaultValuesEv
    // invoke: QStringList defaultValues()
    C._ZNK18QCommandLineOption13defaultValuesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "defaultValues", args)
  }

}

// <= body block end

