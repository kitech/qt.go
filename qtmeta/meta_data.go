package qtmeta

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
)

func init() { log.SetFlags(log.Flags() | log.Lshortfile) }

const Q_REFCOUNT_INITIALIZE_STATIC = int32(-1)

func Q_STATIC_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(size int32, offset int64) QByteArrayData {
	return QByteArrayData{Q_REFCOUNT_INITIALIZE_STATIC, size, 0, offset}
}

func QT_MOC_LITERAL(cnt int32, idx int32, ofs int64, len int32) (QByteArrayData, int64) {
	// offset := qptrdiff(offsetof(qt_meta_stringdata_QtMocRev_t, stringdata0) + ofs - idx * sizeof(QByteArrayData))
	offset := int64(24*cnt) + ofs - int64(idx)*arrdatsz
	bav := Q_STATIC_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, offset)
	return bav, offset
}

var arrdatsz = int64(unsafe.Sizeof(QArrayData{}))

func QT_MOC_LITERAL_OFFSET(cnt int32, idx int32, ofs int64) int64 {
	offset := arrdatsz*int64(cnt) + ofs - int64(idx)*arrdatsz
	return offset
}

/*
QT_MOC_LITERAL(0, 0, 8), // "QtMocRev"
QT_MOC_LITERAL(1, 9, 14), // "firstsignal123"
QT_MOC_LITERAL(2, 24, 0) // ""
*/
// 24B
type QArrayData struct {
	RefCount int32
	Size     int32
	Alloc    uint32
	// CapacityReserved bool
	// pad4   int32 // 此处应有compiler padding
	Offset int64
}
type QByteArrayData QArrayData
type QByteArrayData2 [24]byte

// Data 是 StringData0的索引
// 用于存储元数据中的字符串
// new, append, final, get/count, to*, dump
type QtMetaStringData struct {
	Data        []QByteArrayData
	StringData0 []byte
	nextOffset  int64
	StrMap      map[string]int
	finaled     bool
}

func NewQtMetaStringData() *QtMetaStringData {
	this := &QtMetaStringData{}
	this.Data = make([]QByteArrayData, 0)
	this.StrMap = make(map[string]int)
	return this
}
func (this *QtMetaStringData) Append(s string) int {
	if idx, ok := this.StrMap[s]; ok {
		return idx
	}
	idx := len(this.StrMap)
	this.StrMap[s] = idx
	this.Data = append(this.Data, QByteArrayData{})

	bav, offset := QT_MOC_LITERAL(0, int32(idx), 0, int32(len(s)))
	_ = offset
	this.nextOffset = int64(len(this.StringData0)) + int64(len(s)+1)
	this.Data[idx] = bav
	bav = this.Data[idx]
	// log.Println(idx, this.nextOffset, len(s), bav.Offset, bav.Size, bav.Alloc, bav.RefCount, s)
	this.StringData0 = append(this.StringData0, []byte(s)...)
	this.StringData0 = append(this.StringData0, 0)

	return idx
}

// 计算偏移量
func (this *QtMetaStringData) FinalPass() {
	this.finaled = true
	cnt := len(this.StrMap)
	oldofs := int64(0)
	for i := 0; i < cnt; i++ {
		curofs := QT_MOC_LITERAL_OFFSET(int32(cnt), int32(i), oldofs)
		this.Data[i].Offset = curofs
		oldofs = oldofs + int64(this.Data[i].Size) + 1
	}
}

func (this *QtMetaStringData) Count() int { return len(this.StrMap) }
func (this *QtMetaStringData) Get(s string) (int, bool) {
	idx, ok := this.StrMap[s]
	return idx, ok
}

func memdup(src unsafe.Pointer, len, cap int) unsafe.Pointer {
	if cap <= 0 {
		cap = len
	}
	cptr := C.calloc(1, (C.size_t)(cap))
	C.memcpy(cptr, src, (C.size_t)(len))
	return cptr
}

func (this *QtMetaStringData) ToC() unsafe.Pointer {
	dptr := unsafe.Pointer(&this.Data[0])
	f1sz := len(this.StrMap) * int(unsafe.Sizeof(QByteArrayData{}))
	f2sz := len(this.StringData0)
	tsz := f1sz + f2sz

	cptr := C.calloc(1, (C.size_t)(tsz))
	C.memcpy(cptr, dptr, (C.size_t)(f1sz))

	f2ptr := unsafe.Pointer(uintptr(cptr) + uintptr(f1sz))
	C.memcpy(f2ptr, unsafe.Pointer(&this.StringData0[0]), C.size_t(f2sz))
	return cptr
}
func (this *QtMetaStringData) ToCs() string {
	f1sz := len(this.StrMap) * int(unsafe.Sizeof(QByteArrayData{}))
	f2sz := len(this.StringData0)
	tsz := f1sz + f2sz

	bin2int32 := func(bin []byte) (rv int32) {
		binary.Read(bytes.NewBuffer(bin), binary.LittleEndian, &rv)
		return
	}
	bin2int64 := func(bin []byte) (rv int64) {
		binary.Read(bytes.NewBuffer(bin), binary.LittleEndian, &rv)
		return
	}

	p := this.ToC()
	retstr := "\nidx, refcnt, size, alloc, offset, //\n"
	cbv := C.GoBytes((p), (C.int)(tsz))
	for idx := 0; idx < len(this.Data); idx++ {
		ofs0 := idx * int(arrdatsz)
		refcntb := cbv[ofs0 : ofs0+4]
		sizeb := cbv[ofs0+4 : ofs0+8]
		allocb := cbv[ofs0+8 : ofs0+12]   // 这个位置应该有pad
		offsetb := cbv[ofs0+16 : ofs0+24] // ??? 这内存排列。。。

		retstr += fmt.Sprintf("%d:   %d,   %d,   %d,   %v, //\n", idx,
			bin2int32(refcntb), bin2int32(sizeb), bin2int32(allocb), bin2int64(offsetb))
	}
	return retstr
}

// name => index
func (this *QtMetaStringData) ToMap() map[string]int {
	ret := make(map[string]int)
	ret = this.StrMap
	return ret
}

// like qt generated format
func (this *QtMetaStringData) Dump() string {
	sp1 := "\n"
	sp1 += fmt.Sprintf("cnt: %d\n", len(this.Data))
	instrofs := int32(0)
	for idx, da := range this.Data {
		es := this.StringData0[int(instrofs):int(instrofs+da.Size)]
		sp1 += fmt.Sprintf("QT_MOC_LITERAL(%d, %d, %d), // \"%s\"\n",
			idx, instrofs, da.Size, string(es))
		instrofs += da.Size + 1
	}

	splen := fmt.Sprintf("len: %d\n", len(this.StringData0))
	sp2 := ""
	for _, c := range string(this.StringData0) {
		if c == 0 {
			sp2 += "\\0"
		} else {
			sp2 += string(c)
		}
	}
	s := sp1 + "\n" + splen + `"` + sp2 + `"`
	return s
}

const Q_MOC_OUTPUT_REVISION = 67
const Q_MOC_METADATA_REVEISION = 7

// 可以考虑三遍处理，1,添加项，2,转换成qt所需要的数据格式，3,计算偏移
type QtMetaData struct {
	// header part
	Revision          uint32 // always 7
	ClassNameIdx      uint32 // always 0
	ClassInfoCount    uint32
	ClassInfoOffset   uint32
	MethodCount       uint32
	MethodOffset      uint32
	PropertyCount     uint32
	PropertyOffset    uint32
	EnumCount         uint32
	EnumOffset        uint32
	ConstructorCount  uint32 // always 0
	ConstructorOffset uint32 // always 0
	Flags             uint32 // always 0
	SignalCount       uint32

	ClassInfos []ClassInfo
	Signals    []Method
	Slots      []Method
	Properties []Property
	Enums      []Enum

	Eod uint32

	ClassName  string
	StrMap     map[string]int // 来源于QtMetaStringData
	MetaStrDat *QtMetaStringData
	finaled    bool
	RawData    []uint32 // 全量数据，以上数据的复制，但是是qt所需无结构模式
}

func NewQtMetaData() *QtMetaData {
	this := &QtMetaData{}
	this.Revision = 7

	this.MetaStrDat = NewQtMetaStringData()
	return this
}

func (this *QtMetaData) SetClassName(name string) {
	idx := this.MetaStrDat.Append(name)
	this.ClassNameIdx = uint32(idx)
	this.ClassName = name
}

func (this *QtMetaData) AddClassInfo(key, value string) {
	//TODO check is classinfo with key exist
	keyIdx := this.MetaStrDat.Append(key)
	valIdx := this.MetaStrDat.Append(value)
	this.ClassInfoCount += 1
	// this.ClassInfoOffset
	this.ClassInfos = append(this.ClassInfos, NewClassInfo(key, value, keyIdx, valIdx))
}

// 虽然参数接受reflect.Type, 最好调用端先做 check
func (this *QtMetaData) AddSignal(name string, rettys, argtys []reflect.Type) {
	this.MethodCount += 1
	this.SignalCount += 1
	mth := this._AddMethod(name, rettys, argtys, true)
	this.Signals = append(this.Signals, mth)
}
func (this *QtMetaData) AddSlot(name string, rettys, argtys []reflect.Type) {
	this.MethodCount += 1
	mth := this._AddMethod(name, rettys, argtys, false)
	this.Slots = append(this.Slots, mth)
}
func (this *QtMetaData) _AddMethod(name string, rettys, argtys []reflect.Type, issignal bool) Method {
	nameIdx := this.MetaStrDat.Append(name)
	tagIdx := this.MetaStrDat.Append("")

	mth := Method{}
	mth.Name = name
	mth.NameIdx = uint32(nameIdx)
	mth.Argc = uint32(len(argtys))
	mth.IsSignal = issignal
	mth.Flags = uint32(qtrt.IfElseInt(issignal, 0x06, 0x0a))
	mth.Tag = uint32(tagIdx)

	var prm Parameter
	/// return params
	if len(rettys) == 0 {
		prm = this._AddParameter(&mth, nil, -1)
	} else {
		prm = this._AddParameter(&mth, rettys[0], -1)
	}
	mth.Params = append(mth.Params, prm)

	/// in params
	for idx, argty := range argtys {
		prm = this._AddParameter(&mth, argty, idx)
		prm.TagIdx = mth.Tag
		mth.Params = append(mth.Params, prm)
	}

	return mth
}

func (this *QtMetaData) _AddParameter(mth *Method, argty reflect.Type, idx int) Parameter {
	prm := Parameter{}
	if argty == nil {
		prm.TypeId = _QMetaType__Void
		prm.TypeShowName = qmtynames[int(prm.TypeId)]
		return prm
	}

	basetypemap := map[reflect.Kind]int{
		reflect.Bool: _QMetaType__Bool,

		reflect.Int:   _QMetaType__Int,
		reflect.Int64: _QMetaType__LongLong,
		reflect.Int32: _QMetaType__Int,
		reflect.Int16: _QMetaType__Short,
		reflect.Int8:  _QMetaType__Char,

		reflect.Uint:   _QMetaType__UInt,
		reflect.Uint64: _QMetaType__ULongLong,
		reflect.Uint32: _QMetaType__UInt,
		reflect.Uint16: _QMetaType__UShort,
		reflect.Uint8:  _QMetaType__UChar,

		reflect.Float64: _QMetaType__Double,
		reflect.Float32: _QMetaType__Float,

		reflect.String:        _QMetaType__QString,
		reflect.UnsafePointer: _QMetaType__VoidStar,
	}

	if qtyid, ok := basetypemap[argty.Kind()]; ok {
		prm.TypeId = uint32(qtyid)
		prm.TypeShowName = qmtynames[qtyid]
		return prm
	}

	// qt pod types: typeid, found
	isqtpodty := func(tyname string) (typeid int, found bool) {
		for tyid, showname := range qmtynames {
			if strings.HasSuffix(showname, fmt.Sprintf("::%s", tyname)) {
				found = true
				typeid = tyid
				break
			}
		}
		return
	}
	// TODO see also qtrt.GetQtclassName
	reg := regexp.MustCompile(`^(qt[a-z]+\.)?(Q[A-Z](.+))`)
	if argty.Kind() == reflect.Ptr && reg.MatchString(argty.Elem().String()) {
		// qt classes
		mats := reg.FindAllStringSubmatch(argty.Elem().String(), -1)
		typeid, found := isqtpodty(mats[0][2])
		if found {
			prm.TypeId = uint32(typeid)
			prm.TypeShowName = qmtynames[typeid]
			return prm
		}
	}

	switch argty.Kind() {
	default:
		log.Println("Unsupported type:", mth.Name, argty, idx)
	}

	return prm
}

// 计算偏移量
func (this *QtMetaData) FinalPass() {
	this.finaled = true
	this.MetaStrDat.FinalPass()

	ivec := make([]uint32, 0)
	ivec_append := func(v uint32) { ivec = append(ivec, v) }
	ivec_set := func(idx int, v uint32) { ivec[idx] = v }

	// calc indexs
	curpos := uint32(14)
	if this.ClassInfoCount > 0 {
		this.ClassInfoOffset = curpos
	}
	curpos += this.ClassInfoCount * 2
	if this.MethodCount > 0 {
		this.MethodOffset = curpos
	}
	curpos += 5 * this.MethodCount
	for _, mth := range this.Signals {
		curpos += 1 /*ret*/ + mth.Argc*2
	}
	for _, mth := range this.Slots {
		curpos += 1 /*ret*/ + mth.Argc*2
	}
	if this.PropertyCount > 0 {
		this.PropertyOffset = curpos
	}
	curpos = curpos + this.PropertyCount*3
	if this.EnumCount > 0 {
		this.EnumOffset = curpos
	}

	// header
	ivec_append(this.Revision)
	ivec_append(this.ClassNameIdx)
	ivec_append(this.ClassInfoCount)
	ivec_append(this.ClassInfoOffset)
	ivec_append(this.MethodCount)
	ivec_append(this.MethodOffset)
	ivec_append(this.PropertyCount)
	ivec_append(this.PropertyOffset)
	ivec_append(this.EnumCount)
	ivec_append(this.EnumOffset)
	ivec_append(this.ConstructorCount)
	ivec_append(this.ConstructorOffset)
	ivec_append(this.Flags)
	ivec_append(this.SignalCount)
	ivec_set(0, this.Revision)

	for _, info := range this.ClassInfos {
		ivec_append(info.KeyIdx)
		ivec_append(info.ValueIdx)
	}
	mths := []Method{}
	mths = append(mths, this.Signals...)
	mths = append(mths, this.Slots...)

	for _, mth := range mths {
		ivec_append(mth.NameIdx) // name
		ivec_append(mth.Argc)    // argc
		ivec_append(0)           // paramidx
		ivec_append(mth.Tag)     // tag
		ivec_append(mth.Flags)   // flags
	}
	for midx, mth := range mths {
		mths[midx].ParamOffset = uint32(len(ivec))
		mthofs := int(this.MethodOffset) + midx*5 + 2
		ivec_set(mthofs, mths[midx].ParamOffset)

		for _, prm := range mth.Params {
			ivec_append(prm.TypeId)
		}
		for _, _ = range mth.Params[1:] {
			ivec_append(12345)
		}
	}
	this.Signals = mths[0:this.SignalCount]
	this.Slots = mths[this.SignalCount:]

	for _, prop := range this.Properties {
		ivec_append(prop.NameIdx) // name
		ivec_append(0)            // type
		ivec_append(0)            // flags
	}
	for _, e := range this.Enums {
		ivec_append(e.NameIdx) // name
		ivec_append(e.Flags)   // flags
		ivec_append(e.Count)   // count
		ivec_append(0)         // data
	}

	ivec_append(0) // eof
	// log.Println("meta data total count:", len(ivec), len(ivec)*4,
	//	this.MetaStrDat.Count(), len(this.MetaStrDat.ToCs()))
	this.RawData = ivec
}

func (this *QtMetaData) ToC() unsafe.Pointer {
	return unsafe.Pointer(&this.RawData[0])
}

func (this *QtMetaData) GetMetaStrDat() unsafe.Pointer { return this.MetaStrDat.ToC() }
func (this *QtMetaData) GetMetaDat() unsafe.Pointer    { return this.ToC() }

func (this *QtMetaData) Dump() string {
	s := "" + this.ClassName + "'s meta data:\n"
	hdrfmt := ` // content:
       %d,       // revision
       %d,       // classname
       %d,   %d, // classinfo
       %d,   %d, // methods
       %d,   %d, // properties
       %d,  %d, // enums/sets
       %d,    %d, // constructors
       %d,       // flags
       %d,       // signalCount` + "\n"

	mdo := this
	s += fmt.Sprintf(hdrfmt, mdo.Revision, mdo.ClassNameIdx, mdo.ClassInfoCount, mdo.ClassInfoOffset,
		mdo.MethodCount, mdo.MethodOffset, mdo.PropertyCount, mdo.PropertyOffset,
		mdo.EnumCount, mdo.EnumOffset, mdo.ConstructorCount, mdo.ConstructorOffset,
		mdo.Flags, mdo.SignalCount)

	// classinfo
	s += "\n // classinfo: key, value\n"
	for _, info := range this.ClassInfos {
		s += fmt.Sprintf("       %d,    %d,\n", info.KeyIdx, info.ValueIdx)
	}

	// signals
	s += "\n // signals: name, argc, parameters, tag, flags\n"
	for _, sig := range this.Signals {
		s += fmt.Sprintf("       %d,    %d,   %d,    %d, 0x%x /* Public */,\n",
			sig.NameIdx, sig.Argc, sig.ParamOffset, sig.Tag, sig.Flags)
	}

	// slots
	s += "\n // slots: name, argc, parameters, tag, flags\n"
	for _, slot := range this.Slots {
		s += fmt.Sprintf("       %d,    %d,   %d,    %d, 0x%x /* Public */,\n",
			slot.NameIdx, slot.Argc, slot.ParamOffset, slot.Tag, slot.Flags)
	}

	dumpParams := func(prms []Parameter) string {
		ps := ""
		for _, prm := range prms {
			if prm.TypeShowName == "" {
				ps += fmt.Sprintf("%d, ", prm.TypeId)
			} else {
				ps += fmt.Sprintf("%s, ", prm.TypeShowName)
			}
		}
		for _, prm := range prms[1:] {
			ps += fmt.Sprintf("%d, ", prm.TagIdx) // TODO
		}
		return ps
	}

	// signals: parameters
	s += "\n // signals: parameters\n"
	for _, sig := range this.Signals {
		s += "\t" + dumpParams(sig.Params) + "\n"
	}

	// slots: parameters
	s += "\n // slots: parameters\n"
	for _, slot := range this.Slots {
		s += "\t" + dumpParams(slot.Params) + "\n"
	}

	// properties
	s += "\n // properties: name, type, flags\n"

	// enums
	s += "\n // enums: name, flags, count, data\n"

	// enum data
	s += "\n // enum data: key, value\n"

	// eof
	s += "\n\t 0  //eof\n"

	return s
}

/////
type ClassInfo struct {
	Key      string
	Value    string
	KeyIdx   uint32
	ValueIdx uint32
}

func NewClassInfo(key, value string, keyIdx, valIdx int) ClassInfo {
	return ClassInfo{key, value, uint32(keyIdx), uint32(valIdx)}
}

type Method struct {
	IsSignal    bool
	Name        string
	NameIdx     uint32
	Argc        uint32
	ParamOffset uint32
	Tag         uint32
	Flags       uint32

	Params []Parameter // return info is the first element
}

type Parameter struct {
	TypeName     string // "int"
	TypeShowName string // "QMetaType::Int"
	TypeId       uint32 // 基本类型：直接为id，复合类型，0x80000000 | TypeNameIdx
	TagIdx       uint32
}

func NewMethod(name string, nameIdx int, rettys, argtys []reflect.Type, issignal bool) Method {
	mth := Method{}
	mth.IsSignal = issignal
	mth.Flags = uint32(qtrt.IfElseInt(issignal, 0x06, 0x0a))

	return mth
}

type Constructor struct{}

type Property struct {
	Name    string
	NameIdx uint32
	TypeId  uint32
	Flags   uint32
}

type Enum struct {
	Name       string
	NameIdx    uint32
	Flags      uint32
	Count      uint32
	DataOffset uint32

	Elems []EnumElem
}

type EnumElem struct {
	Name    string
	NameIdx uint32
	Value   uint32
}
