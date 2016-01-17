package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
// src-file: /QtCore/qplugin.h
// dst-file: /src/core/qplugin.go
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
  // proto:  QJsonObject QStaticPlugin::metaData();
extern void _ZNK13QStaticPlugin8metaDataEv(void* qthis); // 4
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

// class sizeof(QStaticPlugin)=16
type QStaticPlugin struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// metaData()
func (this *QStaticPlugin) metaData(args ...interface{}) () {
  // metaData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStaticPlugin8metaDataEv
    // invoke: QJsonObject metaData()
    C._ZNK13QStaticPlugin8metaDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStaticPlugin", "metaData", args)
  }

}

// <= body block end

