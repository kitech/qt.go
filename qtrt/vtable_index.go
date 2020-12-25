package qtrt

type getVTIFunc = func(string, string) int

var get_vtable_index_funcs = []getVTIFunc{}

func RegisterVtableIndexFunc(fn getVTIFunc) {
	get_vtable_index_funcs = append(get_vtable_index_funcs, fn)
}

func GetVtableIndex(clsname, mtname string) int {
	cnt := len(get_vtable_index_funcs)
	for i := cnt - 1; i >= 0; i-- {
		idx := get_vtable_index_funcs[i](clsname, mtname)
		if idx != -1 {
			return idx
		}
	}
	return -1
}

func Inherit(qtobj interface{}, mtname string, fn interface{}) {

}
