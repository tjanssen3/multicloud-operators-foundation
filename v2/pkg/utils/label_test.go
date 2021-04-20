package utils

import (
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCloneAndAddLabel(t *testing.T) {
	type args struct {
		labels     map[string]string
		labelKey   string
		labelValue string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{"case1:", args{labels: map[string]string{"label1": "va1", "label2": "va2"}, labelKey: "key", labelValue: "value"},
			map[string]string{"label1": "va1", "label2": "va2", "key": "value"}},
		{"case2:", args{labels: nil, labelKey: "", labelValue: ""}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CloneAndAddLabel(tt.args.labels, tt.args.labelKey, tt.args.labelValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CloneAndAddLabel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatchLabelForLabelSelector(t *testing.T) {
	type args struct {
		targetLabels  map[string]string
		labelSelector *metav1.LabelSelector
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1:", args{targetLabels: map[string]string{"label1": "va1", "label2": "va2"}}, true},
		{"case2:", args{targetLabels: map[string]string{"label1": "va1", "label2": "va2"},
			labelSelector: &metav1.LabelSelector{MatchLabels: map[string]string{"label1": "va1", "label2": "va2"}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatchLabelForLabelSelector(tt.args.targetLabels, tt.args.labelSelector); got != tt.want {
				t.Errorf("MatchLabelForLabelSelector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeMap(t *testing.T) {
	modified := false
	type args struct {
		modified *bool
		existing map[string]string
		required map[string]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{modified: &modified, required: map[string]string{"label1": "va1"}, existing: map[string]string{"label1": "va1", "label2": "va2"}}, false},
		{"case2", args{modified: &modified, required: map[string]string{"label1": "va1"}, existing: map[string]string{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeMap(tt.args.modified, tt.args.existing, tt.args.required)
			if tt.want != modified {
				t.Errorf("failed to merge map")
			}
		})
	}
}
func TestStringToMap(t *testing.T) {
	testCases := []struct {
		name string
		str  string
		rst  map[string]string
	}{
		{
			name: "case1",
			str:  "",
			rst:  nil,
		},
		{
			name: "case2",
			str:  "app=opt,zone=east-1",
			rst: map[string]string{
				"app":  "opt",
				"zone": "east-1",
			},
		},
	}

	for _, testCase := range testCases {
		rst := StringToMap(testCase.str)
		if len(rst) != len(testCase.rst) {
			t.Errorf("test case %s fail", testCase.name)
		}
	}
}

func TestStrings(t *testing.T) {
	str := []string{"label1", "label2", "label3"}
	outStr := RemoveString(str, "label2")
	if ContainsString(outStr, "label2") {
		t.Errorf("failed to remove string from slice")
	}
}
