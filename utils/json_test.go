package utils

import (
	"reflect"
	"testing"
)

func TestAppendWrite(t *testing.T) {
	// testPath := "./test.json"
	// d := make(map[string]interface{})
	// d["A"] = "a"
	// d["One"] = 1

	type args struct {
		path string
		data map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// {
		// 	name: "１",
		// 	args: args{
		// 		path: testPath,
		// 		data: d,
		// 	},
		// 	wantErr: false,
		// },
		// {
		// 	name: "2",
		// 	args: args{
		// 		path: testPath,
		// 		data: d,
		// 	},
		// 	wantErr: false,
		// },
		// {
		// 	name: "3",
		// 	args: args{
		// 		path: "./test.txt",
		// 		data: d,
		// 	},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AppendWrite(tt.args.path, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("AppendWrite() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRead(t *testing.T) {
	// pre-process
	// testPath := "./test.json"
	// d := make(map[string]interface{})
	// d["A"] = "a"
	// d["One"] = 1
	// if err := AppendWrite(testPath, d); err != nil {
	// 	panic(err.Error())
	// }

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		// {
		// 	name: "1",
		// 	args: args{
		// 		path: testPath,
		// 	},
		// 	want:    d,
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Read(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
