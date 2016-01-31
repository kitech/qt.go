package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZNK9QSettings8containsERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QSettings::setIniCodec(QTextCodec * codec);
extern void C_ZN9QSettings11setIniCodecEP10QTextCodec(void* qthis, void* arg0); // 4
  // proto:  void QSettings::setIniCodec(const char * codecName);
extern void C_ZN9QSettings11setIniCodecEPKc(void* qthis, unsigned char* arg0); // 4
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
extern void C_ZNK9QSettings16fallbacksEnabledEv(void* qthis); // 4
  // proto:  QStringList QSettings::allKeys();
extern void C_ZNK9QSettings7allKeysEv(void* qthis); // 4
  // proto:  QString QSettings::applicationName();
extern void C_ZNK9QSettings15applicationNameEv(void* qthis); // 4
  // proto:  QString QSettings::group();
extern void C_ZNK9QSettings5groupEv(void* qthis); // 4
  // proto:  bool QSettings::isWritable();
extern void C_ZNK9QSettings10isWritableEv(void* qthis); // 4
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
extern void C_ZNK9QSettings8fileNameEv(void* qthis); // 4
  // proto:  QStringList QSettings::childKeys();
extern void C_ZNK9QSettings9childKeysEv(void* qthis); // 4
  // proto:  QString QSettings::organizationName();
extern void C_ZNK9QSettings16organizationNameEv(void* qthis); // 4
  // proto:  void QSettings::QSettings(QObject * parent);
extern void C_ZN9QSettingsC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QSettings::QSettings(const QString & organization, const QString & application, QObject * parent);
extern void C_ZN9QSettingsC2ERK7QStringS2_P7QObject(void* qthis, void* arg0, void* arg1, void* arg2); // 3
  // proto:  QTextCodec * QSettings::iniCodec();
extern void C_ZNK9QSettings8iniCodecEv(void* qthis); // 4
  // proto:  void QSettings::clear();
extern void C_ZN9QSettings5clearEv(void* qthis); // 4
  // proto:  QVariant QSettings::value(const QString & key, const QVariant & defaultValue);
extern void C_ZNK9QSettings5valueERK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1); // 4
  // proto:  int QSettings::beginReadArray(const QString & prefix);
extern void C_ZN9QSettings14beginReadArrayERK7QString(void* qthis, void* arg0); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// contains(const class QString &)
func (this *QSettings) contains(args ...interface{}) () {
  // contains(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings8containsERK7QString
    // invoke: bool contains(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QSettings8containsERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "contains", args)
  }

}

// setIniCodec(class QTextCodec *)
func (this *QSettings) setIniCodec(args ...interface{}) () {
  // setIniCodec(class QTextCodec *)
  // setIniCodec(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings11setIniCodecEP10QTextCodec
    // invoke: void setIniCodec(class QTextCodec *)
    var arg0 = args[0].(QTextCodec).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings11setIniCodecEP10QTextCodec(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN9QSettings11setIniCodecEPKc
    // invoke: void setIniCodec(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings11setIniCodecEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "setIniCodec", args)
  }

}

// sync()
func (this *QSettings) sync(args ...interface{}) () {
  // sync()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings4syncEv
    // invoke: void sync()
    C.C_ZN9QSettings4syncEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "sync", args)
  }

}

// ~QSettings()
func (this *QSettings) FreeQSettings(args ...interface{}) () {
  // ~QSettings()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettingsD0Ev
    // invoke: void ~QSettings()
    C.C_ZN9QSettingsD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "~QSettings", args)
  }

}

// beginGroup(const class QString &)
func (this *QSettings) beginGroup(args ...interface{}) () {
  // beginGroup(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings10beginGroupERK7QString
    // invoke: void beginGroup(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings10beginGroupERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "beginGroup", args)
  }

}

// defaultFormat()
func (this *QSettings) defaultFormat_s(args ...interface{}) () {
  // defaultFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

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

}

// remove(const class QString &)
func (this *QSettings) remove(args ...interface{}) () {
  // remove(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings6removeERK7QString
    // invoke: void remove(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings6removeERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "remove", args)
  }

}

// fallbacksEnabled()
func (this *QSettings) fallbacksEnabled(args ...interface{}) () {
  // fallbacksEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings16fallbacksEnabledEv
    // invoke: bool fallbacksEnabled()
    C.C_ZNK9QSettings16fallbacksEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "fallbacksEnabled", args)
  }

}

// allKeys()
func (this *QSettings) allKeys(args ...interface{}) () {
  // allKeys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings7allKeysEv
    // invoke: QStringList allKeys()
    C.C_ZNK9QSettings7allKeysEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "allKeys", args)
  }

}

// applicationName()
func (this *QSettings) applicationName(args ...interface{}) () {
  // applicationName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings15applicationNameEv
    // invoke: QString applicationName()
    C.C_ZNK9QSettings15applicationNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "applicationName", args)
  }

}

// group()
func (this *QSettings) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings5groupEv
    // invoke: QString group()
    C.C_ZNK9QSettings5groupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "group", args)
  }

}

// isWritable()
func (this *QSettings) isWritable(args ...interface{}) () {
  // isWritable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings10isWritableEv
    // invoke: bool isWritable()
    C.C_ZNK9QSettings10isWritableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "isWritable", args)
  }

}

// format()
func (this *QSettings) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings6formatEv
    // invoke: QSettings::Format format()
    C.C_ZNK9QSettings6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "format", args)
  }

}

// setArrayIndex(int)
func (this *QSettings) setArrayIndex(args ...interface{}) () {
  // setArrayIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings13setArrayIndexEi
    // invoke: void setArrayIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings13setArrayIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "setArrayIndex", args)
  }

}

// endArray()
func (this *QSettings) endArray(args ...interface{}) () {
  // endArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings8endArrayEv
    // invoke: void endArray()
    C.C_ZN9QSettings8endArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "endArray", args)
  }

}

// childGroups()
func (this *QSettings) childGroups(args ...interface{}) () {
  // childGroups()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings11childGroupsEv
    // invoke: QStringList childGroups()
    C.C_ZNK9QSettings11childGroupsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "childGroups", args)
  }

}

// metaObject()
func (this *QSettings) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QSettings10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "metaObject", args)
  }

}

// scope()
func (this *QSettings) scope(args ...interface{}) () {
  // scope()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings5scopeEv
    // invoke: QSettings::Scope scope()
    C.C_ZNK9QSettings5scopeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "scope", args)
  }

}

// setUserIniPath(const class QString &)
func (this *QSettings) setUserIniPath_s(args ...interface{}) () {
  // setUserIniPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings14setUserIniPathERK7QString
    // invoke: void setUserIniPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings14setUserIniPathERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QSettings", "setUserIniPath", args)
  }

}

// setSystemIniPath(const class QString &)
func (this *QSettings) setSystemIniPath_s(args ...interface{}) () {
  // setSystemIniPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings16setSystemIniPathERK7QString
    // invoke: void setSystemIniPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings16setSystemIniPathERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QSettings", "setSystemIniPath", args)
  }

}

// setFallbacksEnabled(_Bool)
func (this *QSettings) setFallbacksEnabled(args ...interface{}) () {
  // setFallbacksEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings19setFallbacksEnabledEb
    // invoke: void setFallbacksEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings19setFallbacksEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "setFallbacksEnabled", args)
  }

}

// status()
func (this *QSettings) status(args ...interface{}) () {
  // status()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings6statusEv
    // invoke: QSettings::Status status()
    C.C_ZNK9QSettings6statusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "status", args)
  }

}

// setValue(const class QString &, const class QVariant &)
func (this *QSettings) setValue(args ...interface{}) () {
  // setValue(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings8setValueERK7QStringRK8QVariant
    // invoke: void setValue(const class QString &, const class QVariant &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QSettings8setValueERK7QStringRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSettings", "setValue", args)
  }

}

// beginWriteArray(const class QString &, int)
func (this *QSettings) beginWriteArray(args ...interface{}) () {
  // beginWriteArray(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings15beginWriteArrayERK7QStringi
    // invoke: void beginWriteArray(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN9QSettings15beginWriteArrayERK7QStringi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSettings", "beginWriteArray", args)
  }

}

// endGroup()
func (this *QSettings) endGroup(args ...interface{}) () {
  // endGroup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings8endGroupEv
    // invoke: void endGroup()
    C.C_ZN9QSettings8endGroupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "endGroup", args)
  }

}

// fileName()
func (this *QSettings) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings8fileNameEv
    // invoke: QString fileName()
    C.C_ZNK9QSettings8fileNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "fileName", args)
  }

}

// childKeys()
func (this *QSettings) childKeys(args ...interface{}) () {
  // childKeys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings9childKeysEv
    // invoke: QStringList childKeys()
    C.C_ZNK9QSettings9childKeysEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "childKeys", args)
  }

}

// organizationName()
func (this *QSettings) organizationName(args ...interface{}) () {
  // organizationName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings16organizationNameEv
    // invoke: QString organizationName()
    C.C_ZNK9QSettings16organizationNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "organizationName", args)
  }

}

// QSettings(class QObject *)
func NewQSettings(args ...interface{}) QSettings {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettingsC1EP7QObject
    // invoke: void QSettings(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QSettingsC2EP7QObject(qthis, arg0)
  case 1:
    // invoke: _ZN9QSettingsC1ERK7QStringS2_P7QObject
    // invoke: void QSettings(const class QString &, const class QString &, class QObject *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QObject).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QSettingsC2ERK7QStringS2_P7QObject(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QSettings", "QSettings", args)
  }

  return QSettings{}
}

// iniCodec()
func (this *QSettings) iniCodec(args ...interface{}) () {
  // iniCodec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings8iniCodecEv
    // invoke: QTextCodec * iniCodec()
    C.C_ZNK9QSettings8iniCodecEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "iniCodec", args)
  }

}

// clear()
func (this *QSettings) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings5clearEv
    // invoke: void clear()
    C.C_ZN9QSettings5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSettings", "clear", args)
  }

}

// value(const class QString &, const class QVariant &)
func (this *QSettings) value(args ...interface{}) () {
  // value(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QSettings5valueERK7QStringRK8QVariant
    // invoke: QVariant value(const class QString &, const class QVariant &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZNK9QSettings5valueERK7QStringRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QSettings", "value", args)
  }

}

// beginReadArray(const class QString &)
func (this *QSettings) beginReadArray(args ...interface{}) () {
  // beginReadArray(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QSettings14beginReadArrayERK7QString
    // invoke: int beginReadArray(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QSettings14beginReadArrayERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSettings", "beginReadArray", args)
  }

}

// <= body block end

