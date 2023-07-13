package query

import (
	"net/url"
	"testing"
)

func TestBindQuery(t *testing.T) {
	type tag struct {
		Key   string `query:"Key"`
		Value string `query:"Value"`
	}
	type request struct {
		Action       string   `query:"Action"`
		RegionID     string   `query:"RegionId"`
		ResourceIds  []string `query:"ResourceId"`
		ResourceType string   `query:"ResourceType"`
		Tags         []tag    `query:"Tag"`
	}

	args1, _ := url.ParseQuery("Action=TagResources&RegionId=cn-hangzhou&ResourceId.1=i-bp1j6qtvdm8w0z1o0&ResourceId.2=i-bp1j6qtvdm8w0z1oP&ResourceType=instance&Tag.1.Key=TestKey&Tag.1.Value=TestValue&Tag.2.Key=TestKey&Tag.2.Value=TestValue")
	output1 := request{}
	type args struct {
		values url.Values
		output interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "aliyun", args: args{values: args1, output: &output1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BindQuery(tt.args.values, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("BindQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
