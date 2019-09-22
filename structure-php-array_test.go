package go_structure_to_php_array

import "testing"

func TestStructTOPhpArray(t *testing.T) {
	source := struct {
		A string           `php:"a"`
		B map[string][]int `php:"b"`
		C struct {
			D map[string][]int `php:"d"`
		} `php:"c"`
		E []struct {
			F map[string][]int `php:"f"`
		} `php:"e"`
	}{
		"2",
		map[string][]int{
			"ddd": {
				2, 3,
			},
		},
		struct {
			D map[string][]int `php:"d"`
		}{D: map[string][]int{
			"dssdd": {
				2, 2,
			},
		},
		},
		[]struct {
			F map[string][]int `php:"f"`
		}{
			{F: map[string][]int{"222": {2, 3}}},
		},
	}
	target := `[
   'a' => '2',
   'b' => [
      'ddd' => [
         2,
         3,
      ],
   ],
   'c' => [
      'd' => [
         'dssdd' => [
            2,
            2,
         ],
      ],
   ],
   'e' => [
      [
         'f' => [
            '222' => [
               2,
               3,
            ],
         ],
      ],
   ],
]`
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "structTOPhpArray", args: args{v: source}, want: target},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructTOPhpArray(tt.args.v); got != tt.want {
				t.Errorf("StructTOPhpArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
