# go-structure-to-php-array
Go-structure-to-php-array is a form of encoding a Go data structure into a PHP array.
# Use
```go
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

target := StructTOPhpArray(source)
```
Output string structure:
```php
[
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
]
```