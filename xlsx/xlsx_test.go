package xlsx_test

type (
	testOb struct {
		String string
		Int    int
		//Int32   int32
		//Int64   int64
		//Uint    uint
		//Uint32  uint32
		//Uint64  uint64
		//Float32 float32
		//Float64 float64
		//Bool    bool
		Nested nested
	}

	nested struct {
		String string
	}
)
