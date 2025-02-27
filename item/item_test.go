package item

import (
	"bytes"
	"testing"
)

func TestRelativeTime(t *testing.T) {
	tests := []struct {
		item Item
		want string
	}{
		{
			item: Item{
				HoursSincePosting: 0.001,
			},
			want: "<1h",
		},
		{
			item: Item{
				HoursSincePosting: 0007,
			},
			want: "7h",
		},
		{
			item: Item{
				HoursSincePosting: 730,
			},
			want: "1m",
		},
		{
			item: Item{
				HoursSincePosting: 370,
			},
			want: "15d",
		},
		{
			item: Item{
				HoursSincePosting: 8760,
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

func TestMarshall(t *testing.T) {
	tests := []struct {
		dataToMarshall Item
		want           []byte
	}{
		{
			dataToMarshall: Item{
				Title:        "A",
				Author:       "L",
				UnixPostTime: 1739890030,
			},
			want: []byte(`{"title":"A","by":"L","time":1739890030,"HoursSincePosting":0,"FormattedTime":"","url":"","score":0}`),
		},
		{
			dataToMarshall: Item{
				Title:  "",
				Author: "M",
			},
			want: []byte(`{"title":"","by":"M","time":0,"HoursSincePosting":0,"FormattedTime":"","url":"","score":0}`),
		},
		{
			dataToMarshall: Item{
				Title:  "Meditations",
				Author: "",
			},
			want: []byte(`{"title":"Meditations","by":"","time":0,"HoursSincePosting":0,"FormattedTime":"","url":"","score":0}`),
		},
		{
			dataToMarshall: Item{
				Title:  "",
				Author: "",
			},
			want: []byte(`{"title":"","by":"","time":0,"HoursSincePosting":0,"FormattedTime":"","url":"","score":0}`),
		},
		{
			dataToMarshall: Item{
				Author: "Marcus",
			},
			want: []byte(`{"title":"","by":"Marcus","time":0,"HoursSincePosting":0,"FormattedTime":"","url":"","score":0}`),
		},
		{
			dataToMarshall: Item{
				Title: "Meditations",
			},
			want: []byte(`{"title":"Meditations","by":"","time":0,"HoursSincePosting":0,"FormattedTime":"","url":"","score":0}`),
		},
		{
			dataToMarshall: Item{},
			want:           []byte(`{"title":"","by":"","time":0,"HoursSincePosting":0,"FormattedTime":"","url":"","score":0}`),
		},
	}

	for _, test := range tests {
		got, err := Marshall(test.dataToMarshall)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(got, test.want) {
			t.Errorf("Got: %v, want: %v", string(got), string(test.want))
		}
	}
}

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
		// 	now := time.Now().Unix()
		// test.want.hoursSincePosting = time.Duration(time.Since(time.Unix(now, 0)).Hours())
		// fmt.Println(test.want.hoursSincePosting)
		got, err := Unmarshal(test.dataToUnmarshall)
		if err != nil {
			t.Fatal(err)
		}

		if got != test.want {
			t.Errorf("Got: %v, want: %v", got, test.want)
		}
	}
}
