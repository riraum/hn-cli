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
			want: []byte(`{"title":"A","by":"L","time":1739890030,` +
				`"HoursSincePosting":0,"FormattedTime":"","url":"","CommentURL":"","score":0}`),
		},
		{
			dataToMarshall: Item{
				Title:  "",
				Author: "M",
			},
			want: []byte(`{"title":"","by":"M","time":0,` + `"HoursSincePosting":0,` +
				`"FormattedTime":"",` +
				`"url":"","CommentURL":"","score":0}`),
		},
		{
			dataToMarshall: Item{
				Title:  "Meditations",
				Author: "",
			},
			want: []byte(`{"title":"Meditations","by":"",` +
				`"time":0,"HoursSincePosting":0,"FormattedTime":"",` +
				`"url":"","CommentURL":"","score":0}`),
		},
		{
			dataToMarshall: Item{
				Title:  "",
				Author: "",
			},
			want: []byte(`{"title":"","by":"","time":0,` +
				`"HoursSincePosting":0,"FormattedTime":"","url":"",` +
				`"CommentURL":"","score":0}`),
		},
		{
			dataToMarshall: Item{
				Author: "Marcus",
			},
			want: []byte(`{"title":"","by":"Marcus","time":0,` +
				`"HoursSincePosting":0,"FormattedTime":"","url":"",` +
				`"CommentURL":"","score":0}`),
		},
		{
			dataToMarshall: Item{
				Title: "Meditations",
			},
			want: []byte(`{"title":"Meditations","by":"",` +
				`"time":0,"HoursSincePosting":0,"FormattedTime":"",` +
				`"url":"","CommentURL":"","score":0}`),
		},
		{
			dataToMarshall: Item{},
			want: []byte(`{"title":"","by":"","time":0,` +
				`"HoursSincePosting":0,"FormattedTime":"","url":"",` +
				`"CommentURL":"","score":0}`),
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
