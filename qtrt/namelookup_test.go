package qtrt

/*
func OverloadInvoke(args []interface{}, fs []interface{}) []interface{} {
	var vtys = make(map[int32]map[int32]reflect.Type, 0)
	for idx := 0; idx < len(fs); idx++ {
		vtys[int32(idx)] = getFuncTypes(fs[idx])
	}
	var midx = SymbolResolve(args, vtys)
	if midx == -1 {
		log.Println("Can not resolve function:", args)
		return nil
	}
	// log.Println("matched_index:", midx)

	ifunc := reflect.ValueOf(fs[midx])
	argv := Iface2Value(args)

	// log.Println(len(args), len(argv), argv)
	vrets := ifunc.Call(argv)
	if vrets == nil {
		return nil
	}
	irets := Value2Iface(vrets)
	return irets
}

func DemoOverload(args ...interface{}) []interface{} {
	f0 := func(a int) {
		log.Println("int")
	}
	f1 := func(a int64) {
		log.Println("int64")
	}
	f2 := func(a string) {
		log.Println("string")
	}
	fs := []interface{}{f0, f1, f2}

	irets := OverloadInvoke(args, fs)
	return irets
}

func Test_do1(t *testing.T) {
	DemoOverload(int(1))
	DemoOverload(int64(1))
	DemoOverload("1")
	t.Log("do1 ok")
}

*/
