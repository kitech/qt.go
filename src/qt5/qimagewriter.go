package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qimagewriter.h
// dst-file: /src/gui/qimagewriter.go
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
// class sizeof(QImageWriter)=8
type QImageWriter struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QImageWriter) setText(args ...interface{}) () {
  // setText(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter7setTextERK7QStringS2_
  default:
    qtrt.ErrorResolve("QImageWriter", "setText", args)
  }

}


func (this *QImageWriter) setGamma(args ...interface{}) () {
  // setGamma(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter8setGammaEf
  default:
    qtrt.ErrorResolve("QImageWriter", "setGamma", args)
  }

}


func (this *QImageWriter) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter11setFileNameERK7QString
  default:
    qtrt.ErrorResolve("QImageWriter", "setFileName", args)
  }

}


func (this *QImageWriter) optimizedWrite(args ...interface{}) () {
  // optimizedWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter14optimizedWriteEv
  default:
    qtrt.ErrorResolve("QImageWriter", "optimizedWrite", args)
  }

}


func (this *QImageWriter) FreeQImageWriter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImageWriter", "~QImageWriter", args)
  }

}


func (this *QImageWriter) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter6deviceEv
  default:
    qtrt.ErrorResolve("QImageWriter", "device", args)
  }

}


func (this *QImageWriter) subType(args ...interface{}) () {
  // subType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter7subTypeEv
  default:
    qtrt.ErrorResolve("QImageWriter", "subType", args)
  }

}


func (this *QImageWriter) supportedMimeTypes_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImageWriter", "supportedMimeTypes", args)
  }

}


func (this *QImageWriter) quality(args ...interface{}) () {
  // quality()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter7qualityEv
  default:
    qtrt.ErrorResolve("QImageWriter", "quality", args)
  }

}


func (this *QImageWriter) write(args ...interface{}) () {
  // write(const class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "const QImage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter5writeERK6QImage
  default:
    qtrt.ErrorResolve("QImageWriter", "write", args)
  }

}


func (this *QImageWriter) setCompression(args ...interface{}) () {
  // setCompression(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter14setCompressionEi
  default:
    qtrt.ErrorResolve("QImageWriter", "setCompression", args)
  }

}


func (this *QImageWriter) supportedImageFormats_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImageWriter", "supportedImageFormats", args)
  }

}


func (this *QImageWriter) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter8fileNameEv
  default:
    qtrt.ErrorResolve("QImageWriter", "fileName", args)
  }

}


func (this *QImageWriter) setOptimizedWrite(args ...interface{}) () {
  // setOptimizedWrite(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter17setOptimizedWriteEb
  default:
    qtrt.ErrorResolve("QImageWriter", "setOptimizedWrite", args)
  }

}


func (this *QImageWriter) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter11errorStringEv
  default:
    qtrt.ErrorResolve("QImageWriter", "errorString", args)
  }

}


func (this *QImageWriter) setQuality(args ...interface{}) () {
  // setQuality(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter10setQualityEi
  default:
    qtrt.ErrorResolve("QImageWriter", "setQuality", args)
  }

}


func (this *QImageWriter) gamma(args ...interface{}) () {
  // gamma()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter5gammaEv
  default:
    qtrt.ErrorResolve("QImageWriter", "gamma", args)
  }

}


func (this *QImageWriter) description(args ...interface{}) () {
  // description()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter11descriptionEv
  default:
    qtrt.ErrorResolve("QImageWriter", "description", args)
  }

}


func NewQImageWriter(args ...interface{}) QImageWriter {
  return QImageWriter{}
}


func (this *QImageWriter) setFormat(args ...interface{}) () {
  // setFormat(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter9setFormatERK10QByteArray
  default:
    qtrt.ErrorResolve("QImageWriter", "setFormat", args)
  }

}


func (this *QImageWriter) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter9setDeviceEP9QIODevice
  default:
    qtrt.ErrorResolve("QImageWriter", "setDevice", args)
  }

}


func (this *QImageWriter) setSubType(args ...interface{}) () {
  // setSubType(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter10setSubTypeERK10QByteArray
  default:
    qtrt.ErrorResolve("QImageWriter", "setSubType", args)
  }

}


func (this *QImageWriter) progressiveScanWrite(args ...interface{}) () {
  // progressiveScanWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter20progressiveScanWriteEv
  default:
    qtrt.ErrorResolve("QImageWriter", "progressiveScanWrite", args)
  }

}


func (this *QImageWriter) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter6formatEv
  default:
    qtrt.ErrorResolve("QImageWriter", "format", args)
  }

}


func (this *QImageWriter) supportedSubTypes(args ...interface{}) () {
  // supportedSubTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter17supportedSubTypesEv
  default:
    qtrt.ErrorResolve("QImageWriter", "supportedSubTypes", args)
  }

}


func (this *QImageWriter) canWrite(args ...interface{}) () {
  // canWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter8canWriteEv
  default:
    qtrt.ErrorResolve("QImageWriter", "canWrite", args)
  }

}


func (this *QImageWriter) compression(args ...interface{}) () {
  // compression()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter11compressionEv
  default:
    qtrt.ErrorResolve("QImageWriter", "compression", args)
  }

}


func (this *QImageWriter) setProgressiveScanWrite(args ...interface{}) () {
  // setProgressiveScanWrite(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter23setProgressiveScanWriteEb
  default:
    qtrt.ErrorResolve("QImageWriter", "setProgressiveScanWrite", args)
  }

}


func (this *QImageWriter) setDescription(args ...interface{}) () {
  // setDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter14setDescriptionERK7QString
  default:
    qtrt.ErrorResolve("QImageWriter", "setDescription", args)
  }

}

// <= body block end

