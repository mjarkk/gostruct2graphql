package structs

// TestStruct is a struct that shows the options
type TestStruct struct {
	Item1String string
	Item2Int    int
	Item3Uint   uint
	SomeStruct  struct {
		Item1 string
		Item2 string
	}
	// StructArr []struct {
	// 	Item1 string
	// 	Item2 string
	// }
	StringArr []string
}

// TestSlice is just an slice with the TestStruct
type TestSlice []TestStruct
