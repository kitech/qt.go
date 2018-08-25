
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
	*qtcore.QThread  `qt:inherit` // 要继承的类
	*qtcore.QProcess `qt:inherit`

	_ struct{} `qt:classinfo,appName=12345s`  //
	_ struct{} `qt:classinfo,appCorp=hehehhe` //

	Prop123 int    `qt:property,123` // 最后字段为默认值
	Prop456 string `qt:property,"testv123"`
	Enum123 AType  `qt:enum` // 就怕无法支持

	_          Q_SIGNALS_BEGIN
	Clicked123 func(bool) `qt:signal`

	_ func(int)     `qt:slot,SlotFunc1` // 如果为小写则为私有slot, 大写为公有slot
	_ func(float32) `qt:slot,SlotFunc2`
}

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

