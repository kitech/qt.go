package qtmeta

import "unsafe"

// syntax parser and mark relate types
type Q_SIGNALS_BEGIN struct{}
type Q_SIGNALS_END struct{}
type Q_META_SECTION_END struct{}

/////
// Derive(this)
// pargs, args for parent1, parent2...
func Derive(obj interface{}, pargs ...[]interface{}) {
	// fill obj's inherit struct
	// parse obj's signal and fill function body
	// 所有的public方法都是slot
}
func Underive(obj interface{}) {}

type QMetaObjectData struct {
	// per class members
	Superdata          unsafe.Pointer
	Stringdata         unsafe.Pointer
	Data               unsafe.Pointer
	Static_metacall    unsafe.Pointer
	relatedMetaObjects unsafe.Pointer
	extradata          unsafe.Pointer

	// per instance members
	// Metacall unsafe.Pointer
	// Metacast unsafe.Pointer
}

func NewQMetaObjectData() {

}

func (this *QMetaObjectData) GetCthis() unsafe.Pointer {
	return unsafe.Pointer(this)
}

func (this *QMetaObjectData) SetCthis(unsafe.Pointer) {}
