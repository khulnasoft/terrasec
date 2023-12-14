package policy

import (
	"reflect"
	"sort"
	"testing"
)

func TestSupportedPolicyTypes(t *testing.T) {
	t.Run("supported policy types", func(t *testing.T) {
		var want []string
		for k := range supportedCloudProvider {
			want = append(want, string(k))
		}
		sort.Strings(want)
		got := SupportedPolicyTypes(true)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: '%v', want: '%v'", got, want)
		}
	})
}

func TestSupportedNotIndirectPolicyTypes(t *testing.T) {
	t.Run("supported policy types", func(t *testing.T) {
		var want []string
		for k, v := range supportedCloudProvider {
			if !v.isIndirect {
				want = append(want, string(k))
			}
		}
		sort.Strings(want)
		got := SupportedPolicyTypes(false)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: '%v', want: '%v'", got, want)
		}
	})
}
