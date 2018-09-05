
```
/////
// Derive(this)
// pargs, args for parent1, parent2...
func Derive(obj interface{}, pargs ...[]interface{}) {
	// fill obj's inherit struct
	// parse obj's signal and fill function body
	// 所有的public方法都是slot
}
func Underive(obj interface{}) {}

type WantAsQClass struct {
	*qtcore.QThread  `qt:"inherit"` // 要继承的类
	*qtcore.QProcess `qt:"inherit"`

	_ struct{} `qt:"classinfo" key:"12345s" value:"hehehhe"` //

	Prop123 int    `qt:"property" value:"123"` // 最后字段为默认值
	Prop456 string `qt:"property" value:"testv123"`
	// Enum123 AType  `qt:enum` // 就怕无法支持

	_          qtmeta.Q_SIGNALS_BEGIN
	Clicked123 func(bool) `qt:"signal"`

	_SlotFunc1 func(int)     `qt:"slot"` // Prefix _ of real method name
	_SlotFunc2 func(float32) `qt:"slot"`
	_SlotFunc3 func()        `qt:"slot"`
}

func (this *WantAsQClass) SlotFunc1(int) {}
func (this *WantAsQClass) SlotFunc2(float32) {}
func (this *WantAsQClass) SlotFunc3() {}


func NewWantAsQClass() *WantAsQClass {
	this := &WantAsQClass{}

	Derive(this)
	return this
}

func (this *WantAsQClass) SlotFunc1(int) {

}

type Q_SIGNALS_BEGIN struct{}
type Q_SIGNALS_END struct{}
type Q_META_SECTION_END struct{}

type AType int

const (
	One   AType = 1 << iota // 1
	Two                     // 2 (i.e. 1 << 1)
	Three                   // 4 (i.e 1 << 2)
)

```

