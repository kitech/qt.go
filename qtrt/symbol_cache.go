package qtrt

import "sync"

// hope it faster than realtime symbol lookup
// symbol name crc32 =>
var symcache sync.Map // uint32 => unsafe.Pointer
