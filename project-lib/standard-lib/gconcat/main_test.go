package gconcat

import (
	"encoding/json"
	"testing"
)

//go test -v -run ^Test_Build_OneStr
func Test_Build_OneStr(t *testing.T) {
	var many1String, many1Int, many1Int8, many1Int16, many1Int32, many1Int64 interface{}
	var many1Uint, many1Uint8, many1Uint16, many1Uint32, many1Uint64 interface{}
	var many1Float32, many1Float64 interface{}
	var many1bool, many2bool interface{}

	many1bool = false
	many2bool = true
	many1String = "21"
	many1Int = 100
	many1Int8 = int8(100)
	many1Int16 = int16(100)
	many1Int32 = int32(100)
	many1Int64 = int64(100)
	many1Uint = uint(100)
	many1Uint8 = uint8(100)
	many1Uint16 = uint16(100)
	many1Uint32 = uint32(100)
	many1Uint64 = uint64(100)
	//many1Uintptr = uintptr(100)

	many1Float32 = float32(100.01)
	many1Float64 = float64(100.01)
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test_buildStr_", args{str: many1String}, "21"},
		{"test_buildStr_", args{str: many1Int}, "100"},
		{"test_buildStr_", args{str: many1Int8}, "100"},
		{"test_buildStr_", args{str: many1Int16}, "100"},
		{"test_buildStr_", args{str: many1Int32}, "100"},
		{"test_buildStr_", args{str: many1Int64}, "100"},
		{"test_buildStr_", args{str: many1Uint}, "100"},
		{"test_buildStr_", args{str: many1Uint8}, "100"},
		{"test_buildStr_", args{str: many1Uint16}, "100"},
		{"test_buildStr_", args{str: many1Uint32}, "100"},
		{"test_buildStr_", args{str: many1Uint64}, "100"},
		//{"test_buildStr_12", args{str: many1Uintptr}, "d"}, //endereço do ponteiro
		{"test_buildStr_", args{str: many1Float32}, "100.010002"},
		{"test_buildStr_", args{str: many1Float64}, "100.010000"},
		{"test_buildStr_", args{str: many1bool}, "false"},
		{"test_buildStr_", args{str: many2bool}, "true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt1 := tt
			got := Build(tt1.args.str)
			switch {
			case got == tt1.want:
				break
			default:
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
			print("\n", got)
		})
	}
}

//go test -v -run ^Test_Build_Many
func Test_Build_Many(t *testing.T) {

	var many1String, many1Int, many1Int8, many1Int16, many1Int32, many1Int64 []interface{}
	var many1Uint, many1Uint8, many1Uint16, many1Uint32, many1Uint64, many1Uintptr []interface{}
	var many1Float32, many1Float64 []interface{}
	var many1bool []interface{}
	many1bool = append(many1bool, []bool{true, false})
	many1String = append(many1String, []string{"21", "joao", "beta"})
	many1Int = append(many1Int, []int{100, 98, 23})
	many1Int8 = append(many1Int8, []int8{100, 98, 23})
	many1Int16 = append(many1Int16, []int16{100, 98, 23})
	many1Int32 = append(many1Int32, []int32{100, 98, 23})
	many1Int64 = append(many1Int64, []int64{100, 98, 23})
	many1Uint = append(many1Uint, []uint{100, 98, 23})
	many1Uint8 = append(many1Uint8, []uint8{100, 98, 23})
	many1Uint16 = append(many1Uint16, []uint16{100, 98, 23})
	many1Uint32 = append(many1Uint32, []uint32{100, 98, 23})
	many1Uint64 = append(many1Uint64, []uint64{100, 98, 23})
	many1Uintptr = append(many1Uintptr, []uintptr{100})

	many1Float32 = append(many1Float32, []float32{100.01, 100.01})
	many1Float64 = append(many1Float64, []float64{100.01, 100.01})
	type args struct {
		strs []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		//{"test_concat_12", args{strs: many1Uintptr}, "d"}, //endereço do ponteiro
		{"test_concat_", args{strs: many1bool}, "truefalse"},
		{"test_concat_", args{strs: many1String}, "21joaobeta"},
		{"test_concat_", args{strs: many1Int}, "1009823"},
		{"test_concat_", args{strs: many1Int8}, "1009823"},
		{"test_concat_", args{strs: many1Int16}, "1009823"},
		{"test_concat_", args{strs: many1Int32}, "1009823"},
		{"test_concat_", args{strs: many1Int64}, "1009823"},
		{"test_concat_", args{strs: many1Uint}, "1009823"},
		{"test_concat_", args{strs: many1Uint8}, "1009823"},
		{"test_concat_", args{strs: many1Uint16}, "1009823"},
		{"test_concat_", args{strs: many1Uint32}, "1009823"},
		{"test_concat_", args{strs: many1Uint64}, "1009823"},
		{"test_concat_", args{strs: many1Float32}, "100.010002100.010002"},
		{"test_concat_", args{strs: many1Float64}, "100.010000100.010000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()
			tt1 := tt
			got := Build(tt1.args.strs...)
			switch {
			case got == tt1.want:
				break
			default:
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
			print("\n", got)
		})
	}
}

func BenchmarkBuild(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Build(
			[]int{1, 2, 3, 4, 5, 56, 6, 7, 7, 778, 8, 88, 8, 8, 8, 8, 8, 8, 9, 9, 123, 4, 4, 5, 6, 7, 77, 8, 8, 99, 9, 93, 3, 3, 3, 3, 45, 5, 6, 6, 7},
			[]string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
			[]string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
			[]string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
			[]string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
			[]string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
			[]string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
			[]string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
			[]string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
			[]string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
			" ",
			"jeff", "somente", "string heree........",
		)
	}
}

func BenchmarkMarshal(b *testing.B) {
	type X struct {
		Num   []int
		Str1  []string
		Str2  []string
		Str3  []string
		Str4  []string
		Str5  []string
		Str6  []string
		Str7  []string
		Str8  []string
		Str9  []string
		Str10 []string
	}

	var obj = X{
		Num:   []int{2, 3, 4, 5, 56, 6, 7, 7, 778, 8, 88, 8, 8, 8, 8, 8, 8, 9, 9, 123, 4, 4, 5, 6, 7, 77, 8, 8, 99, 9, 93, 3, 3, 3, 3, 45, 5, 6, 6, 7},
		Str1:  []string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
		Str2:  []string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
		Str3:  []string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
		Str4:  []string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
		Str5:  []string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
		Str6:  []string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
		Str7:  []string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
		Str8:  []string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
		Str9:  []string{"jeff", "otoni", "lima", " ", " vamos testar esta cria "},
		Str10: []string{"jeff", "somente", "string heree........"},
	}

	for n := 0; n < b.N; n++ {
		_, err := json.Marshal(&obj)
		if err != nil {
			panic(err)
		}
	}
}
