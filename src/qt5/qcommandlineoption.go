package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QCommandLineOption::setDescription(const QString & description);
extern void C_ZN18QCommandLineOption14setDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringList QCommandLineOption::defaultValues();
extern void C_ZNK18QCommandLineOption13defaultValuesEv(void* qthis); // 4
  // proto:  QString QCommandLineOption::description();
extern void* C_ZNK18QCommandLineOption11descriptionEv(void* qthis); // 4
  // proto:  void QCommandLineOption::QCommandLineOption(const QStringList & names, const QString & description, const QString & valueName, const QString & defaultValue);
extern void* C_ZN18QCommandLineOptionC2ERK11QStringListRK7QStringS5_S5_(void* arg0, void* arg1, void* arg2, void* arg3); // 3
  // proto:  void QCommandLineOption::QCommandLineOption(const QStringList & names);
extern void* C_ZN18QCommandLineOptionC2ERK11QStringList(void* arg0); // 3
  // proto:  void QCommandLineOption::QCommandLineOption(const QCommandLineOption & other);
extern void* C_ZN18QCommandLineOptionC2ERKS_(void* arg0); // 3
  // proto:  void QCommandLineOption::QCommandLineOption(const QString & name);
extern void* C_ZN18QCommandLineOptionC2ERK7QString(void* arg0); // 3
  // proto:  void QCommandLineOption::QCommandLineOption(const QString & name, const QString & description, const QString & valueName, const QString & defaultValue);
extern void* C_ZN18QCommandLineOptionC2ERK7QStringS2_S2_S2_(void* arg0, void* arg1, void* arg2, void* arg3); // 3
  // proto:  void QCommandLineOption::setDefaultValue(const QString & defaultValue);
extern void C_ZN18QCommandLineOption15setDefaultValueERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QCommandLineOption::~QCommandLineOption();
extern void C_ZN18QCommandLineOptionD2Ev(void* qthis); // 4
  // proto:  void QCommandLineOption::setDefaultValues(const QStringList & defaultValues);
extern void C_ZN18QCommandLineOption16setDefaultValuesERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QCommandLineOption::setValueName(const QString & name);
extern void C_ZN18QCommandLineOption12setValueNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringList QCommandLineOption::names();
extern void C_ZNK18QCommandLineOption5namesEv(void* qthis); // 4
  // proto:  QString QCommandLineOption::valueName();
extern void* C_ZNK18QCommandLineOption9valueNameEv(void* qthis); // 4
  // proto:  void QCommandLineOption::swap(QCommandLineOption & other);
extern void C_ZN18QCommandLineOption4swapERS_(void* qthis, void* arg0); // 2
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setDescription(const class QString &)
func (this *QCommandLineOption) Setdescription(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QCommandLineOption14setDescriptionERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setDescription", args)
  }

  return
}

// defaultValues()
func (this *QCommandLineOption) Defaultvalues(args ...interface{}) () {
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
    C.C_ZNK18QCommandLineOption13defaultValuesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "defaultValues", args)
  }

  return
}

// description()
func (this *QCommandLineOption) Description(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QCommandLineOption11descriptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineOption", "description", args)
  }

  return
}

// QCommandLineOption(const class QStringList &, const class QString &, const class QString &, const class QString &)
func NewQCommandLineOption(args ...interface{}) *QCommandLineOption {
  // QCommandLineOption(const class QStringList &, const class QString &, const class QString &, const class QString &)
  // QCommandLineOption(const class QStringList &)
  // QCommandLineOption(const class QCommandLineOption &)
  // QCommandLineOption(const class QString &)
  // QCommandLineOption(const class QString &, const class QString &, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QCommandLineOption{}) // "const QCommandLineOption &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][3] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOptionC1ERK11QStringListRK7QStringS5_S5_
    // invoke: void QCommandLineOption(const class QStringList &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QString).Qclsinst
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QCommandLineOptionC2ERK11QStringListRK7QStringS5_S5_(arg0, arg1, arg2, arg3)
    return &QCommandLineOption{Qclsinst:qthis}
  case 1:
    // invoke: _ZN18QCommandLineOptionC1ERK11QStringList
    // invoke: void QCommandLineOption(const class QStringList &)
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QCommandLineOptionC2ERK11QStringList(arg0)
    return &QCommandLineOption{Qclsinst:qthis}
  case 2:
    // invoke: _ZN18QCommandLineOptionC1ERKS_
    // invoke: void QCommandLineOption(const class QCommandLineOption &)
    var arg0 = args[0].(*QCommandLineOption).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QCommandLineOptionC2ERKS_(arg0)
    return &QCommandLineOption{Qclsinst:qthis}
  case 3:
    // invoke: _ZN18QCommandLineOptionC1ERK7QString
    // invoke: void QCommandLineOption(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QCommandLineOptionC2ERK7QString(arg0)
    return &QCommandLineOption{Qclsinst:qthis}
  case 4:
    // invoke: _ZN18QCommandLineOptionC1ERK7QStringS2_S2_S2_
    // invoke: void QCommandLineOption(const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QString).Qclsinst
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QCommandLineOptionC2ERK7QStringS2_S2_S2_(arg0, arg1, arg2, arg3)
    return &QCommandLineOption{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QCommandLineOption", "QCommandLineOption", args)
  }

  return nil // QCommandLineOption{Qclsinst:qthis}
}

// setDefaultValue(const class QString &)
func (this *QCommandLineOption) Setdefaultvalue(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QCommandLineOption15setDefaultValueERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setDefaultValue", args)
  }

  return
}

// ~QCommandLineOption()
func (this *QCommandLineOption) Freeqcommandlineoption(args ...interface{}) () {
  // ~QCommandLineOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOptionD0Ev
    // invoke: void ~QCommandLineOption()
    C.C_ZN18QCommandLineOptionD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "~QCommandLineOption", args)
  }

  return
}

// setDefaultValues(const class QStringList &)
func (this *QCommandLineOption) Setdefaultvalues(args ...interface{}) () {
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
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QCommandLineOption16setDefaultValuesERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setDefaultValues", args)
  }

  return
}

// setValueName(const class QString &)
func (this *QCommandLineOption) Setvaluename(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QCommandLineOption12setValueNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setValueName", args)
  }

  return
}

// names()
func (this *QCommandLineOption) Names(args ...interface{}) () {
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
    C.C_ZNK18QCommandLineOption5namesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "names", args)
  }

  return
}

// valueName()
func (this *QCommandLineOption) Valuename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QCommandLineOption9valueNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLineOption", "valueName", args)
  }

  return
}

// swap(class QCommandLineOption &)
func (this *QCommandLineOption) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QCommandLineOption).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QCommandLineOption4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLineOption", "swap", args)
  }

  return
}

// <= body block end

