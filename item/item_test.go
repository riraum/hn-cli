package item

import (
	"testing"
)

func TestRelativeTime(t *testing.T) {
	tests := []struct {
		item Item
		want string
	}{
		{
			item: Item{
				UnixPostTime: 1725210000,
			},
			want: "6m",
		},
		{
			item: Item{
				UnixPostTime: 1737478800,
			},
			want: "1m",
		},
		{
			item: Item{
				UnixPostTime: 1739145600,
			},
			want: "13d",
		},
		{
			item: Item{
				UnixPostTime: 1708560000,
			},
			want: "1y",
		},
	}

	for _, test := range tests {
		got := test.item.RelativeTime()
		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}

// func TestMarshall(t *testing.T) {
// 	tests := []struct {
// 		dataToMarshall Item
// 		want           []byte
// 	}{
// 		{
// 			dataToMarshall: Item{
// 				Title:        "Alice in Wonderland",
// 				Author:       "Lewis Carroll",
// 				UnixPostTime: 1739890030,
// 			},
// 			want: []byte(`{"title":"Alice in Wonderland","by":"Lewis Carroll","time":1739890030,"url":"","score":0}`),
// 		},
// 		{
// 			dataToMarshall: Item{
// 				Title:  "",
// 				Author: "Marcus Aurelius",
// 			},
// 			want: []byte(`{"title":"","by":"Marcus Aurelius","time":0,"url":"","score":0}`),
// 		},
// 		{
// 			dataToMarshall: Item{
// 				Title:  "Meditations",
// 				Author: "",
// 			},
// 			want: []byte(`{"title":"Meditations","by":"","time":0,"url":"","score":0}`),
// 		},
// 		{
// 			dataToMarshall: Item{
// 				Title:  "",
// 				Author: "",
// 			},
// 			want: []byte(`{"title":"","by":"","time":0,"url":"","score":0}`),
// 		},
// 		{
// 			dataToMarshall: Item{
// 				Author: "Marcus Aurelius",
// 			},
// 			want: []byte(`{"title":"","by":"Marcus Aurelius","time":0,"url":"","score":0}`),
// 		},
// 		{
// 			dataToMarshall: Item{
// 				Title: "Meditations",
// 			},
// 			want: []byte(`{"title":"Meditations","by":"","time":0,"url":"","score":0}`),
// 		},
// 		{
// 			dataToMarshall: Item{},
// 			want:           []byte(`{"title":"","by":"","time":0,"url":"","score":0}`),
// 		},
// 	}

// 	for _, test := range tests {
// 		got, err := Marshall(test.dataToMarshall)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if !bytes.Equal(got, test.want) {
// 			t.Errorf("Got: %v, want: %v", string(got), string(test.want))
// 		}
// 	}
// }

func TestUnmarshall(t *testing.T) {
	tests := []struct {
		dataToUnmarshall []byte
		want             Item
	}{
		{
			dataToUnmarshall: []byte(`{"title":"Alice in Wonderland","by":"Lewis Carroll","url":"","score":0}`),
			want: Item{
				Title:  "Alice in Wonderland",
				Author: "Lewis Carroll",
			},
		},
		{
			dataToUnmarshall: []byte(`{"title":"Meditations"}`),
			want: Item{
				Title: "Meditations",
			},
		},
		{
			dataToUnmarshall: []byte(`{"by":"Marcus Aurelius"}`),
			want: Item{
				Author: "Marcus Aurelius",
			},
		},
		{
			dataToUnmarshall: []byte(`{"title":"Meditations","by":""}`),
			want: Item{
				Author: "",
				Title:  "Meditations",
			},
		},
		{
			dataToUnmarshall: []byte(`{"title":"","by":"Marcus Aurelius"}`),
			want: Item{
				Author: "Marcus Aurelius",
				Title:  "",
			},
		},
		{
			dataToUnmarshall: []byte(`{"title":"","by":""}`),
			want: Item{
				Author: "",
				Title:  "",
			},
		},
		{
			dataToUnmarshall: []byte(`{}`),
			want:             Item{},
		},
	}

	for _, test := range tests {
		got, err := Unmarshal(test.dataToUnmarshall)
		if err != nil {
			t.Fatal(err)
		}

		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}
