package
// versions:
//
//	protoc-gen-go v1.34.2
//	protoc        v5.27.2
//
// source: api/cpu/cpu.proto
// Verify that this generated code is sufficiently up-to-date.
// Verify that runtime/protoimpl is sufficiently up-to-date.
// Информация о CPU.
// Бренд процессора.
// Наименование процессора.
// Кол-во ядер процессора.
// Чипсет  процессора.
// Минимальная частота процессора.
// Максимальная частота процессора.
// Deprecated: Use CPU.ProtoReflect.Descriptor instead.
// 0: cpu.CPU
// [0:0] is the sub-list for method output_type
// [0:0] is the sub-list for method input_type
// [0:0] is the sub-list for extension type_name
// [0:0] is the sub-list for extension extendee
// [0:0] is the sub-list for field type_name
out

type CPU struct {
	Brand       string
	Name        string
	NumberCores uint32
	Chipset     string
	MinGhz      float64
	MaxGhz      float64
}
