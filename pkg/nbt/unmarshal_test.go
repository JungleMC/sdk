package nbt

import (
	"github.com/JungleMC/sdk/internal/nbt"
	"reflect"
	"testing"
)

func TestUnmarshalCompoundMap(t *testing.T) {
	tests := []struct {
		name          string
		input         []byte
		expected      map[string]interface{}
		expectedError bool
	}{
		{
			name:  "unnamed root compound tag",
			input: nbt.UnnamedRootCompoundBytes,
			expected: map[string]interface{}{
				"ByteTag":   byte(0xFF),
				"StringTag": "hello, world",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actualRaw interface{}
			actualRaw = make(map[string]interface{})

			_, err := Unmarshal(tt.input, reflect.ValueOf(actualRaw))
			if (err != nil) != tt.expectedError {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.expectedError)
				return
			}

			if !reflect.DeepEqual(tt.expected, actualRaw) {
				t.Errorf("tags not equal")
				return
			}
		})
	}
}

func TestUnmarshalCompoundStruct(t *testing.T) {
	tests := []struct {
		name        string
		tagBytes    []byte
		wantTagName string
		want        interface{}
		wantErr     bool
	}{
		{
			name:        "unnamed root compound tag",
			tagBytes:    nbt.UnnamedRootCompoundBytes,
			wantTagName: "",
			want: nbt.UnnamedRootCompound{
				ByteTag:   0xFF,
				StringTag: "hello, world",
			},
			wantErr: false,
		},
		{
			name:        "bananrama",
			tagBytes:    nbt.BananramaBytes,
			wantTagName: "",
			want:        nbt.BananramaStruct,
			wantErr:     false,
		},
		{
			name:        "bigtest",
			tagBytes:    nbt.BigTestBytes,
			wantTagName: "Level",
			want: nbt.BigTest{
				LongTest:   9223372036854775807,
				ShortTest:  32767,
				StringTest: "HELLO WORLD THIS IS A TEST STRING \xc3\x85\xc3\x84\xc3\x96!",
				FloatTest:  0.49823147058486938,
				IntTest:    2147483647,
				NCT: nbt.BigTestNCT{
					Egg: nbt.BigTestNameAndFloat32{
						Name:  "Eggbert",
						Value: 0.5,
					},
					Ham: nbt.BigTestNameAndFloat32{
						Name:  "Hampus",
						Value: 0.75,
					},
				},
				ListTest: []int64{11, 12, 13, 14, 15},
				ListTest2: [2]nbt.BigTestCompound{
					{
						Name:      "Compound tag #0",
						CreatedOn: 1264099775885,
					},
					{
						Name:      "Compound tag #1",
						CreatedOn: 1264099775885,
					},
				},
				ByteTest:      127,
				ByteArrayTest: nbt.BigTestByteArray(),
				DoubleTest:    0.49312871321823148,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRaw := reflect.New(reflect.TypeOf(tt.want)).Elem()

			_, err := Unmarshal(tt.tagBytes, actualRaw)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tt.want, actualRaw.Interface()) {
				t.Errorf("tags not equal")
				return
			}
		})
	}
}
