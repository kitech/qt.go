package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtCore/qsettings.h
// dst-file: /src/core/qsettings.go
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
  // proto:  bool QSettings::contains(const QString & key);
extern bool C_ZNK9QSettings8containsERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QSettings::setIniCodec(QTextCodec * codec);
extern void C_ZN9QSettings11setIniCodecEP10QTextCodec(void* qthis, void* arg0); // 4
  // proto:  void QSettings::setIniCodec(const char * codecName);
extern void C_ZN9QSettings11setIniCodecEPKc(void* qthis, void* arg0); // 4
  // proto:  void QSettings::sync();
extern void C_ZN9QSettings4syncEv(void* qthis); // 4
  // proto:  void QSettings::~QSettings();
extern void C_ZN9QSettingsD2Ev(void* qthis); // 4
  // proto:  void QSettings::beginGroup(const QString & prefix);
extern void C_ZN9QSettings10beginGroupERK7QString(void* qthis, void* arg0); // 4
  // proto: static QSettings::Format QSettings::defaultFormat();
extern void C_ZN9QSettings13defaultFormatEv(); // 4
  // proto:  void QSettings::remove(const QString & key);
extern void C_ZN9QSettings6removeERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QSettings::fallbacksEnabled();
extern bool C_ZNK9QSettings16fallbacksEnabledEv(void* qthis); // 4
  // proto:  QStringList QSettings::allKeys();
extern void C_ZNK9QSettings7allKeysEv(void* qthis); // 4
  // proto:  QString QSettings::applicationName();
extern void* C_ZNK9QSettings15applicationNameEv(void* qthis); // 4
  // proto:  QString QSettings::group();
extern void* C_ZNK9QSettings5groupEv(void* qthis); // 4
  // proto:  bool QSettings::isWritable();
extern bool C_ZNK9QSettings10isWritableEv(void* qthis); // 4
  // proto:  QSettings::Format QSettings::format();
extern void C_ZNK9QSettings6formatEv(void* qthis); // 4
  // proto:  void QSettings::setArrayIndex(int i);
extern void C_ZN9QSettings13setArrayIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  void QSettings::endArray();
extern void C_ZN9QSettings8endArrayEv(void* qthis); // 4
  // proto:  QStringList QSettings::childGroups();
extern void C_ZNK9QSettings11childGroupsEv(void* qthis); // 4
  // proto:  const QMetaObject * QSettings::metaObject();
extern void C_ZNK9QSettings10metaObjectEv(void* qthis); // 4
  // proto:  QSettings::Scope QSettings::scope();
extern void C_ZNK9QSettings5scopeEv(void* qthis); // 4
  // proto: static void QSettings::setUserIniPath(const QString & dir);
extern void C_ZN9QSettings14setUserIniPathERK7QString(void* arg0); // 4
  // proto: static void QSettings::setSystemIniPath(const QString & dir);
extern void C_ZN9QSettings16setSystemIniPathERK7QString(void* arg0); // 4
  // proto:  void QSettings::setFallbacksEnabled(bool b);
extern void C_ZN9QSettings19setFallbacksEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QSettings::Status QSettings::status();
extern void C_ZNK9QSettings6statusEv(void* qthis); // 4
  // proto:  void QSettings::setValue(const QString & key, const QVariant & value);
extern void C_ZN9QSettings8setValueERK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QSettings::beginWriteArray(const QString & prefix, int size);
extern void C_ZN9QSettings15beginWriteArrayERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QSettings::endGroup();
extern void C_ZN9QSettings8endGroupEv(void* qthis); // 4
  // proto:  QString QSettings::fileName();
extern void* C_ZNK9QSettings8fileNameEv(void* qthis); // 4
  // proto:  QStringList QSettings::childKeys();
extern void C_ZNK9QSettings9childKeysEv(void* qthis); // 4
  // proto:  QString QSettings::organizationName();
extern void* C_ZNK9QSettings16organizationNameEv(void* qthis); // 4
  // proto:  void QSettings::QSettings(QObject * parent);
extern void* C_ZN9QSettingsC2EP7QObject(void* arg0); // 3
  // proto:  void QSettings::QSettings(const QString & organization, const QString & application, QObject * parent);
extern void* C_ZN9QSettingsC2ERK7QStringS2_P7QObject(void* arg0, void* arg1, void* arg2); // 3
  // proto:  QTextCodec * QSettings::iniCodec();
extern void* C_ZNK9QSettings8iniCodecEv(void* qthis); // 4
  // proto:  void QSettings::clear();
extern void C_ZN9QSettings5clearEv(void* qthis); // 4
  // proto:  QVariant QSettings::value(const QString & key, const QVariant & defaultValue);
extern void* C_ZNK9QSettings5valueERK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1); // 4
  // proto:  int QSettings::beginReadArray(const QString & prefix);
extern int32_t C_ZN9QSettings14beginReadArrayERK7QString(void* qthis, void* arg0); // 4
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

// class sizeof(QSettings)=1
type QSettings struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// contains(const class QString &)
func (this *QSettings) Contains(args ...interface{}) (ret interface{}) {
  // contains(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings8containsERK7QString
    // invoke: bool contains(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QSettings8containsERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSettings", "contains", args)
  }

  return
}

// setIniCodec(class QTextCodec *)
func (this *QSettings) Setinicodec(args ...interface{}) () {
  // setIniCodec(class QTextCodec *)
  // setIniCodec(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings11setIniCodecEP10QTextCodec
    // invoke: void setIniCodec(class QTextCodec *)
    var arg0 = args[0].(*QTextCodec).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings11setIniCodecEP10QTextCodec(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN9QSettings11setIniCodecEPKc
    // invoke: void setIniCodec(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN9QSettings11setIniCodecEPKc(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "setIniCodec", args)
  }

  return
}

// sync()
func (this *QSettings) Sync(args ...interface{}) () {
  // sync()
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
    // invoke: _ZN9QSettings4syncEv
    // invoke: void sync()
    C.C_ZN9QSettings4syncEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "sync", args)
  }

  return
}

// ~QSettings()
func (this *QSettings) Freeqsettings(args ...interface{}) () {
  // ~QSettings()
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
    // invoke: _ZN9QSettingsD0Ev
    // invoke: void ~QSettings()
    C.C_ZN9QSettingsD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "~QSettings", args)
  }

  return
}

// beginGroup(const class QString &)
func (this *QSettings) Begingroup(args ...interface{}) () {
  // beginGroup(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings10beginGroupERK7QString
    // invoke: void beginGroup(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings10beginGroupERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "beginGroup", args)
  }

  return
}

// defaultFormat()
func (this *QSettings) Defaultformat_S(args ...interface{}) () {
  // defaultFormat()
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
    // invoke: _ZN9QSettings13defaultFormatEv
    // invoke: QSettings::Format defaultFormat()
    C.C_ZN9QSettings13defaultFormatEv()
  default:
    qtrt.ErrorResolve("QSettings", "defaultFormat", args)
  }

  return
}

// remove(const class QString &)
func (this *QSettings) Remove(args ...interface{}) () {
  // remove(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings6removeERK7QString
    // invoke: void remove(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings6removeERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "remove", args)
  }

  return
}

// fallbacksEnabled()
func (this *QSettings) Fallbacksenabled(args ...interface{}) (ret interface{}) {
  // fallbacksEnabled()
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
    // invoke: _ZNK9QSettings16fallbacksEnabledEv
    // invoke: bool fallbacksEnabled()
    var ret0 = C.C_ZNK9QSettings16fallbacksEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSettings", "fallbacksEnabled", args)
  }

  return
}

// allKeys()
func (this *QSettings) Allkeys(args ...interface{}) () {
  // allKeys()
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
    // invoke: _ZNK9QSettings7allKeysEv
    // invoke: QStringList allKeys()
    C.C_ZNK9QSettings7allKeysEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "allKeys", args)
  }

  return
}

// applicationName()
func (this *QSettings) Applicationname(args ...interface{}) (ret interface{}) {
  // applicationName()
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
    // invoke: _ZNK9QSettings15applicationNameEv
    // invoke: QString applicationName()
    var ret0 = C.C_ZNK9QSettings15applicationNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSettings", "applicationName", args)
  }

  return
}

// group()
func (this *QSettings) Group(args ...interface{}) (ret interface{}) {
  // group()
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
    // invoke: _ZNK9QSettings5groupEv
    // invoke: QString group()
    var ret0 = C.C_ZNK9QSettings5groupEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSettings", "group", args)
  }

  return
}

// isWritable()
func (this *QSettings) Iswritable(args ...interface{}) (ret interface{}) {
  // isWritable()
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
    // invoke: _ZNK9QSettings10isWritableEv
    // invoke: bool isWritable()
    var ret0 = C.C_ZNK9QSettings10isWritableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSettings", "isWritable", args)
  }

  return
}

// format()
func (this *QSettings) Format(args ...interface{}) () {
  // format()
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
    // invoke: _ZNK9QSettings6formatEv
    // invoke: QSettings::Format format()
    C.C_ZNK9QSettings6formatEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "format", args)
  }

  return
}

// setArrayIndex(int)
func (this *QSettings) Setarrayindex(args ...interface{}) () {
  // setArrayIndex(int)
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
    // invoke: _ZN9QSettings13setArrayIndexEi
    // invoke: void setArrayIndex(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings13setArrayIndexEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "setArrayIndex", args)
  }

  return
}

// endArray()
func (this *QSettings) Endarray(args ...interface{}) () {
  // endArray()
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
    // invoke: _ZN9QSettings8endArrayEv
    // invoke: void endArray()
    C.C_ZN9QSettings8endArrayEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "endArray", args)
  }

  return
}

// childGroups()
func (this *QSettings) Childgroups(args ...interface{}) () {
  // childGroups()
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
    // invoke: _ZNK9QSettings11childGroupsEv
    // invoke: QStringList childGroups()
    C.C_ZNK9QSettings11childGroupsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "childGroups", args)
  }

  return
}

// metaObject()
func (this *QSettings) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK9QSettings10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QSettings10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "metaObject", args)
  }

  return
}

// scope()
func (this *QSettings) Scope(args ...interface{}) () {
  // scope()
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
    // invoke: _ZNK9QSettings5scopeEv
    // invoke: QSettings::Scope scope()
    C.C_ZNK9QSettings5scopeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "scope", args)
  }

  return
}

// setUserIniPath(const class QString &)
func (this *QSettings) Setuserinipath_S(args ...interface{}) () {
  // setUserIniPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings14setUserIniPathERK7QString
    // invoke: void setUserIniPath(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings14setUserIniPathERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QSettings", "setUserIniPath", args)
  }

  return
}

// setSystemIniPath(const class QString &)
func (this *QSettings) Setsysteminipath_S(args ...interface{}) () {
  // setSystemIniPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings16setSystemIniPathERK7QString
    // invoke: void setSystemIniPath(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings16setSystemIniPathERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QSettings", "setSystemIniPath", args)
  }

  return
}

// setFallbacksEnabled(_Bool)
func (this *QSettings) Setfallbacksenabled(args ...interface{}) () {
  // setFallbacksEnabled(_Bool)
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
    // invoke: _ZN9QSettings19setFallbacksEnabledEb
    // invoke: void setFallbacksEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings19setFallbacksEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "setFallbacksEnabled", args)
  }

  return
}

// status()
func (this *QSettings) Status(args ...interface{}) () {
  // status()
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
    // invoke: _ZNK9QSettings6statusEv
    // invoke: QSettings::Status status()
    C.C_ZNK9QSettings6statusEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "status", args)
  }

  return
}

// setValue(const class QString &, const class QVariant &)
func (this *QSettings) Setvalue(args ...interface{}) () {
  // setValue(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings8setValueERK7QStringRK8QVariant
    // invoke: void setValue(const class QString &, const class QVariant &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QSettings8setValueERK7QStringRK8QVariant(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSettings", "setValue", args)
  }

  return
}

// beginWriteArray(const class QString &, int)
func (this *QSettings) Beginwritearray(args ...interface{}) () {
  // beginWriteArray(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings15beginWriteArrayERK7QStringi
    // invoke: void beginWriteArray(const class QString &, int)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QSettings15beginWriteArrayERK7QStringi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSettings", "beginWriteArray", args)
  }

  return
}

// endGroup()
func (this *QSettings) Endgroup(args ...interface{}) () {
  // endGroup()
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
    // invoke: _ZN9QSettings8endGroupEv
    // invoke: void endGroup()
    C.C_ZN9QSettings8endGroupEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "endGroup", args)
  }

  return
}

// fileName()
func (this *QSettings) Filename(args ...interface{}) (ret interface{}) {
  // fileName()
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
    // invoke: _ZNK9QSettings8fileNameEv
    // invoke: QString fileName()
    var ret0 = C.C_ZNK9QSettings8fileNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSettings", "fileName", args)
  }

  return
}

// childKeys()
func (this *QSettings) Childkeys(args ...interface{}) () {
  // childKeys()
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
    // invoke: _ZNK9QSettings9childKeysEv
    // invoke: QStringList childKeys()
    C.C_ZNK9QSettings9childKeysEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "childKeys", args)
  }

  return
}

// organizationName()
func (this *QSettings) Organizationname(args ...interface{}) (ret interface{}) {
  // organizationName()
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
    // invoke: _ZNK9QSettings16organizationNameEv
    // invoke: QString organizationName()
    var ret0 = C.C_ZNK9QSettings16organizationNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSettings", "organizationName", args)
  }

  return
}

// QSettings(class QObject *)
func NewQSettings(args ...interface{}) *QSettings {
  // QSettings(class QObject *)
  // QSettings(const class QString &, const class QString &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettingsC1EP7QObject
    // invoke: void QSettings(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QSettingsC2EP7QObject(arg0)
    return &QSettings{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QSettingsC1ERK7QStringS2_P7QObject
    // invoke: void QSettings(const class QString &, const class QString &, class QObject *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QSettingsC2ERK7QStringS2_P7QObject(arg0, arg1, arg2)
    return &QSettings{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSettings", "QSettings", args)
  }

  return nil // QSettings{Qclsinst:qthis}
}

// iniCodec()
func (this *QSettings) Inicodec(args ...interface{}) (ret interface{}) {
  // iniCodec()
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
    // invoke: _ZNK9QSettings8iniCodecEv
    // invoke: QTextCodec * iniCodec()
    var ret0 = C.C_ZNK9QSettings8iniCodecEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSettings", "iniCodec", args)
  }

  return
}

// clear()
func (this *QSettings) Clear(args ...interface{}) () {
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
    // invoke: _ZN9QSettings5clearEv
    // invoke: void clear()
    C.C_ZN9QSettings5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "clear", args)
  }

  return
}

// value(const class QString &, const class QVariant &)
func (this *QSettings) Value(args ...interface{}) (ret interface{}) {
  // value(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings5valueERK7QStringRK8QVariant
    // invoke: QVariant value(const class QString &, const class QVariant &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QSettings5valueERK7QStringRK8QVariant(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSettings", "value", args)
  }

  return
}

// beginReadArray(const class QString &)
func (this *QSettings) Beginreadarray(args ...interface{}) (ret interface{}) {
  // beginReadArray(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings14beginReadArrayERK7QString
    // invoke: int beginReadArray(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QSettings14beginReadArrayERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSettings", "beginReadArray", args)
  }

  return
}

// <= body block end

