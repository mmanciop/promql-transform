package transform_test

import (
	"testing"

	"github.com/canonical/promql-transform/pkg/transform"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Input    string
	Matchers map[string]string
	Expected string
}

func TestShouldApplyLabelMatcherToVectorSelector(t *testing.T) {
	cases := []TestCase{
		{
			Input:    "rate(metric[5m]) > 0.5",
			Matchers: map[string]string{"bar": "baz"},
			Expected: `rate(metric{bar="baz"}[5m]) > 0.5`,
		},
		{
			Input:    "metric",
			Matchers: map[string]string{"bar": "baz"},
			Expected: `metric{bar="baz"}`,
		},
		{
			Input:    "up == 0",
			Matchers: map[string]string{"cool": "breeze", "hot": "sunrays"},
			Expected: `up{cool="breeze",hot="sunrays"} == 0`,
		},
		{
			Input:    "absent(up{job=\"prometheus\"})",
			Matchers: map[string]string{"model": "lma"},
			Expected: `absent(up{job="prometheus",model="lma"})`,
		},
	}
	for _, c := range cases {
		out, err := transform.Transform(c.Input, &c.Matchers)
		assert.NoError(t, err)
		assert.Equal(t, c.Expected, out)
	}
}
