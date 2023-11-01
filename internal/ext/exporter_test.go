package ext

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.flipt.io/flipt/rpc/flipt"
	"gopkg.in/yaml.v3"
)

type mockLister struct {
	namespaces []*flipt.Namespace

	nsToFlags    map[string][]*flipt.Flag
	nsToSegments map[string][]*flipt.Segment
	nsToRules    map[string][]*flipt.Rule
	nsToRollouts map[string][]*flipt.Rollout
}

func (m mockLister) ListNamespaces(_ context.Context, _ *flipt.ListNamespaceRequest) (*flipt.NamespaceList, error) {
	return &flipt.NamespaceList{
		Namespaces: m.namespaces,
	}, nil
}

func (m mockLister) ListFlags(_ context.Context, listRequest *flipt.ListFlagRequest) (*flipt.FlagList, error) {
	flags := m.nsToFlags[listRequest.NamespaceKey]

	return &flipt.FlagList{
		Flags: flags,
	}, nil
}

func (m mockLister) ListRules(_ context.Context, listRequest *flipt.ListRuleRequest) (*flipt.RuleList, error) {
	rules := m.nsToRules[listRequest.NamespaceKey]

	if listRequest.FlagKey == "flag1" {
		return &flipt.RuleList{
			Rules: rules,
		}, nil
	}

	return &flipt.RuleList{}, nil
}

func (m mockLister) ListSegments(_ context.Context, listRequest *flipt.ListSegmentRequest) (*flipt.SegmentList, error) {
	segments := m.nsToSegments[listRequest.NamespaceKey]

	return &flipt.SegmentList{
		Segments: segments,
	}, nil
}

func (m mockLister) ListRollouts(_ context.Context, listRequest *flipt.ListRolloutRequest) (*flipt.RolloutList, error) {
	rollouts := m.nsToRollouts[listRequest.NamespaceKey]

	if listRequest.FlagKey == "flag2" {
		return &flipt.RolloutList{
			Rules: rollouts,
		}, nil
	}

	return &flipt.RolloutList{}, nil
}

func TestExport(t *testing.T) {
	tests := []struct {
		name          string
		lister        mockLister
		path          string
		namespaces    string
		allNamespaces bool
	}{
		{
			name: "single default namespace",
			lister: mockLister{
				nsToFlags: map[string][]*flipt.Flag{
					"default": {
						{
							Key:         "flag1",
							Name:        "flag1",
							Type:        flipt.FlagType_VARIANT_FLAG_TYPE,
							Description: "description",
							Enabled:     true,
							Variants: []*flipt.Variant{
								{
									Id:   "1",
									Key:  "variant1",
									Name: "variant1",
									Attachment: `{
										"pi": 3.141,
										"happy": true,
										"name": "Niels",
										"nothing": null,
										"answer": {
										  "everything": 42
										},
										"list": [1, 0, 2],
										"object": {
										  "currency": "USD",
										  "value": 42.99
										}
									  }`,
								},
								{
									Id:  "2",
									Key: "foo",
								},
							},
						},
						{
							Key:         "flag2",
							Name:        "flag2",
							Type:        flipt.FlagType_BOOLEAN_FLAG_TYPE,
							Description: "a boolean flag",
							Enabled:     false,
						},
					},
				},
				nsToSegments: map[string][]*flipt.Segment{
					"default": {
						{
							Key:         "segment1",
							Name:        "segment1",
							Description: "description",
							MatchType:   flipt.MatchType_ANY_MATCH_TYPE,
							Constraints: []*flipt.Constraint{
								{
									Id:          "1",
									Type:        flipt.ComparisonType_STRING_COMPARISON_TYPE,
									Property:    "foo",
									Operator:    "eq",
									Value:       "baz",
									Description: "desc",
								},
								{
									Id:          "2",
									Type:        flipt.ComparisonType_STRING_COMPARISON_TYPE,
									Property:    "fizz",
									Operator:    "neq",
									Value:       "buzz",
									Description: "desc",
								},
							},
						},
						{
							Key:         "segment2",
							Name:        "segment2",
							Description: "description",
							MatchType:   flipt.MatchType_ANY_MATCH_TYPE,
						},
					},
				},
				nsToRules: map[string][]*flipt.Rule{
					"default": {
						{
							Id:         "1",
							SegmentKey: "segment1",
							Rank:       1,
							Distributions: []*flipt.Distribution{
								{
									Id:        "1",
									VariantId: "1",
									RuleId:    "1",
									Rollout:   100,
								},
							},
						},
						{
							Id:              "2",
							SegmentKeys:     []string{"segment1", "segment2"},
							SegmentOperator: flipt.SegmentOperator_AND_SEGMENT_OPERATOR,
							Rank:            2,
						},
					},
				},

				nsToRollouts: map[string][]*flipt.Rollout{
					"default": {
						{
							Id:          "1",
							FlagKey:     "flag2",
							Type:        flipt.RolloutType_SEGMENT_ROLLOUT_TYPE,
							Description: "enabled for internal users",
							Rank:        int32(1),
							Rule: &flipt.Rollout_Segment{
								Segment: &flipt.RolloutSegment{
									SegmentKey: "internal_users",
									Value:      true,
								},
							},
						},
						{
							Id:          "2",
							FlagKey:     "flag2",
							Type:        flipt.RolloutType_THRESHOLD_ROLLOUT_TYPE,
							Description: "enabled for 50%",
							Rank:        int32(2),
							Rule: &flipt.Rollout_Threshold{
								Threshold: &flipt.RolloutThreshold{
									Percentage: float32(50.0),
									Value:      true,
								},
							},
						},
					},
				},
			},
			path:          "testdata/export.yml",
			namespaces:    "default",
			allNamespaces: false,
		},
		{
			name: "multiple namespaces",
			lister: mockLister{
				nsToFlags: map[string][]*flipt.Flag{
					"default": {
						{
							Key:         "flag1",
							Name:        "flag1",
							Type:        flipt.FlagType_VARIANT_FLAG_TYPE,
							Description: "description",
							Enabled:     true,
							Variants: []*flipt.Variant{
								{
									Id:   "1",
									Key:  "variant1",
									Name: "variant1",
									Attachment: `{
										"pi": 3.141,
										"happy": true,
										"name": "Niels",
										"nothing": null,
										"answer": {
										  "everything": 42
										},
										"list": [1, 0, 2],
										"object": {
										  "currency": "USD",
										  "value": 42.99
										}
									  }`,
								},
								{
									Id:  "2",
									Key: "foo",
								},
							},
						},
						{
							Key:         "flag2",
							Name:        "flag2",
							Type:        flipt.FlagType_BOOLEAN_FLAG_TYPE,
							Description: "a boolean flag",
							Enabled:     false,
						},
					},
					"foo": {
						{
							Key:         "flag1",
							Name:        "flag1",
							Type:        flipt.FlagType_VARIANT_FLAG_TYPE,
							Description: "description",
							Enabled:     true,
							Variants: []*flipt.Variant{
								{
									Id:   "1",
									Key:  "variant1",
									Name: "variant1",
									Attachment: `{
										"pi": 3.141,
										"happy": true,
										"name": "Niels",
										"nothing": null,
										"answer": {
										  "everything": 42
										},
										"list": [1, 0, 2],
										"object": {
										  "currency": "USD",
										  "value": 42.99
										}
									  }`,
								},
								{
									Id:  "2",
									Key: "foo",
								},
							},
						},
						{
							Key:         "flag2",
							Name:        "flag2",
							Type:        flipt.FlagType_BOOLEAN_FLAG_TYPE,
							Description: "a boolean flag",
							Enabled:     false,
						},
					},
				},
				nsToSegments: map[string][]*flipt.Segment{
					"default": {
						{
							Key:         "segment1",
							Name:        "segment1",
							Description: "description",
							MatchType:   flipt.MatchType_ANY_MATCH_TYPE,
							Constraints: []*flipt.Constraint{
								{
									Id:          "1",
									Type:        flipt.ComparisonType_STRING_COMPARISON_TYPE,
									Property:    "foo",
									Operator:    "eq",
									Value:       "baz",
									Description: "desc",
								},
								{
									Id:          "2",
									Type:        flipt.ComparisonType_STRING_COMPARISON_TYPE,
									Property:    "fizz",
									Operator:    "neq",
									Value:       "buzz",
									Description: "desc",
								},
							},
						},
						{
							Key:         "segment2",
							Name:        "segment2",
							Description: "description",
							MatchType:   flipt.MatchType_ANY_MATCH_TYPE,
						},
					},
					"foo": {
						{
							Key:         "segment1",
							Name:        "segment1",
							Description: "description",
							MatchType:   flipt.MatchType_ANY_MATCH_TYPE,
							Constraints: []*flipt.Constraint{
								{
									Id:          "1",
									Type:        flipt.ComparisonType_STRING_COMPARISON_TYPE,
									Property:    "foo",
									Operator:    "eq",
									Value:       "baz",
									Description: "desc",
								},
								{
									Id:          "2",
									Type:        flipt.ComparisonType_STRING_COMPARISON_TYPE,
									Property:    "fizz",
									Operator:    "neq",
									Value:       "buzz",
									Description: "desc",
								},
							},
						},
						{
							Key:         "segment2",
							Name:        "segment2",
							Description: "description",
							MatchType:   flipt.MatchType_ANY_MATCH_TYPE,
						},
					},
				},
				nsToRules: map[string][]*flipt.Rule{
					"default": {
						{
							Id:         "1",
							SegmentKey: "segment1",
							Rank:       1,
							Distributions: []*flipt.Distribution{
								{
									Id:        "1",
									VariantId: "1",
									RuleId:    "1",
									Rollout:   100,
								},
							},
						},
						{
							Id:              "2",
							SegmentKeys:     []string{"segment1", "segment2"},
							SegmentOperator: flipt.SegmentOperator_AND_SEGMENT_OPERATOR,
							Rank:            2,
						},
					},
					"foo": {
						{
							Id:         "1",
							SegmentKey: "segment1",
							Rank:       1,
							Distributions: []*flipt.Distribution{
								{
									Id:        "1",
									VariantId: "1",
									RuleId:    "1",
									Rollout:   100,
								},
							},
						},
						{
							Id:              "2",
							SegmentKeys:     []string{"segment1", "segment2"},
							SegmentOperator: flipt.SegmentOperator_AND_SEGMENT_OPERATOR,
							Rank:            2,
						},
					},
				},

				nsToRollouts: map[string][]*flipt.Rollout{
					"default": {
						{
							Id:          "1",
							FlagKey:     "flag2",
							Type:        flipt.RolloutType_SEGMENT_ROLLOUT_TYPE,
							Description: "enabled for internal users",
							Rank:        int32(1),
							Rule: &flipt.Rollout_Segment{
								Segment: &flipt.RolloutSegment{
									SegmentKey: "internal_users",
									Value:      true,
								},
							},
						},
						{
							Id:          "2",
							FlagKey:     "flag2",
							Type:        flipt.RolloutType_THRESHOLD_ROLLOUT_TYPE,
							Description: "enabled for 50%",
							Rank:        int32(2),
							Rule: &flipt.Rollout_Threshold{
								Threshold: &flipt.RolloutThreshold{
									Percentage: float32(50.0),
									Value:      true,
								},
							},
						},
					},
					"foo": {
						{
							Id:          "1",
							FlagKey:     "flag2",
							Type:        flipt.RolloutType_SEGMENT_ROLLOUT_TYPE,
							Description: "enabled for internal users",
							Rank:        int32(1),
							Rule: &flipt.Rollout_Segment{
								Segment: &flipt.RolloutSegment{
									SegmentKey: "internal_users",
									Value:      true,
								},
							},
						},
						{
							Id:          "2",
							FlagKey:     "flag2",
							Type:        flipt.RolloutType_THRESHOLD_ROLLOUT_TYPE,
							Description: "enabled for 50%",
							Rank:        int32(2),
							Rule: &flipt.Rollout_Threshold{
								Threshold: &flipt.RolloutThreshold{
									Percentage: float32(50.0),
									Value:      true,
								},
							},
						},
					},
				},
			},
			path:          "testdata/export_default_and_foo.yml",
			namespaces:    "default,foo",
			allNamespaces: false,
		},
		{
			name: "all namespaces",
			lister: mockLister{
				namespaces: []*flipt.Namespace{
					{
						Key: "default",
					},
					{
						Key: "foo",
					},
					{
						Key: "bar",
					},
				},
				nsToFlags: map[string][]*flipt.Flag{
					"foo": {
						{
							Key:         "flag1",
							Name:        "flag1",
							Type:        flipt.FlagType_VARIANT_FLAG_TYPE,
							Description: "description",
							Enabled:     true,
							Variants: []*flipt.Variant{
								{
									Id:   "1",
									Key:  "variant1",
									Name: "variant1",
									Attachment: `{
										"pi": 3.141,
										"happy": true,
										"name": "Niels",
										"nothing": null,
										"answer": {
										  "everything": 42
										},
										"list": [1, 0, 2],
										"object": {
										  "currency": "USD",
										  "value": 42.99
										}
									  }`,
								},
								{
									Id:  "2",
									Key: "foo",
								},
							},
						},
						{
							Key:         "flag2",
							Name:        "flag2",
							Type:        flipt.FlagType_BOOLEAN_FLAG_TYPE,
							Description: "a boolean flag",
							Enabled:     false,
						},
					},
					"bar": {
						{
							Key:         "flag1",
							Name:        "flag1",
							Type:        flipt.FlagType_VARIANT_FLAG_TYPE,
							Description: "description",
							Enabled:     true,
							Variants: []*flipt.Variant{
								{
									Id:   "1",
									Key:  "variant1",
									Name: "variant1",
									Attachment: `{
										"pi": 3.141,
										"happy": true,
										"name": "Niels",
										"nothing": null,
										"answer": {
										  "everything": 42
										},
										"list": [1, 0, 2],
										"object": {
										  "currency": "USD",
										  "value": 42.99
										}
									  }`,
								},
								{
									Id:  "2",
									Key: "foo",
								},
							},
						},
						{
							Key:         "flag2",
							Name:        "flag2",
							Type:        flipt.FlagType_BOOLEAN_FLAG_TYPE,
							Description: "a boolean flag",
							Enabled:     false,
						},
					},
				},
				nsToSegments: map[string][]*flipt.Segment{
					"foo": {
						{
							Key:         "segment1",
							Name:        "segment1",
							Description: "description",
							MatchType:   flipt.MatchType_ANY_MATCH_TYPE,
							Constraints: []*flipt.Constraint{
								{
									Id:          "1",
									Type:        flipt.ComparisonType_STRING_COMPARISON_TYPE,
									Property:    "foo",
									Operator:    "eq",
									Value:       "baz",
									Description: "desc",
								},
								{
									Id:          "2",
									Type:        flipt.ComparisonType_STRING_COMPARISON_TYPE,
									Property:    "fizz",
									Operator:    "neq",
									Value:       "buzz",
									Description: "desc",
								},
							},
						},
						{
							Key:         "segment2",
							Name:        "segment2",
							Description: "description",
							MatchType:   flipt.MatchType_ANY_MATCH_TYPE,
						},
					},
					"bar": {
						{
							Key:         "segment1",
							Name:        "segment1",
							Description: "description",
							MatchType:   flipt.MatchType_ANY_MATCH_TYPE,
							Constraints: []*flipt.Constraint{
								{
									Id:          "1",
									Type:        flipt.ComparisonType_STRING_COMPARISON_TYPE,
									Property:    "foo",
									Operator:    "eq",
									Value:       "baz",
									Description: "desc",
								},
								{
									Id:          "2",
									Type:        flipt.ComparisonType_STRING_COMPARISON_TYPE,
									Property:    "fizz",
									Operator:    "neq",
									Value:       "buzz",
									Description: "desc",
								},
							},
						},
						{
							Key:         "segment2",
							Name:        "segment2",
							Description: "description",
							MatchType:   flipt.MatchType_ANY_MATCH_TYPE,
						},
					},
				},
				nsToRules: map[string][]*flipt.Rule{
					"foo": {
						{
							Id:         "1",
							SegmentKey: "segment1",
							Rank:       1,
							Distributions: []*flipt.Distribution{
								{
									Id:        "1",
									VariantId: "1",
									RuleId:    "1",
									Rollout:   100,
								},
							},
						},
						{
							Id:              "2",
							SegmentKeys:     []string{"segment1", "segment2"},
							SegmentOperator: flipt.SegmentOperator_AND_SEGMENT_OPERATOR,
							Rank:            2,
						},
					},
					"bar": {
						{
							Id:         "1",
							SegmentKey: "segment1",
							Rank:       1,
							Distributions: []*flipt.Distribution{
								{
									Id:        "1",
									VariantId: "1",
									RuleId:    "1",
									Rollout:   100,
								},
							},
						},
						{
							Id:              "2",
							SegmentKeys:     []string{"segment1", "segment2"},
							SegmentOperator: flipt.SegmentOperator_AND_SEGMENT_OPERATOR,
							Rank:            2,
						},
					},
				},

				nsToRollouts: map[string][]*flipt.Rollout{
					"foo": {
						{
							Id:          "1",
							FlagKey:     "flag2",
							Type:        flipt.RolloutType_SEGMENT_ROLLOUT_TYPE,
							Description: "enabled for internal users",
							Rank:        int32(1),
							Rule: &flipt.Rollout_Segment{
								Segment: &flipt.RolloutSegment{
									SegmentKey: "internal_users",
									Value:      true,
								},
							},
						},
						{
							Id:          "2",
							FlagKey:     "flag2",
							Type:        flipt.RolloutType_THRESHOLD_ROLLOUT_TYPE,
							Description: "enabled for 50%",
							Rank:        int32(2),
							Rule: &flipt.Rollout_Threshold{
								Threshold: &flipt.RolloutThreshold{
									Percentage: float32(50.0),
									Value:      true,
								},
							},
						},
					},
					"bar": {
						{
							Id:          "1",
							FlagKey:     "flag2",
							Type:        flipt.RolloutType_SEGMENT_ROLLOUT_TYPE,
							Description: "enabled for internal users",
							Rank:        int32(1),
							Rule: &flipt.Rollout_Segment{
								Segment: &flipt.RolloutSegment{
									SegmentKey: "internal_users",
									Value:      true,
								},
							},
						},
						{
							Id:          "2",
							FlagKey:     "flag2",
							Type:        flipt.RolloutType_THRESHOLD_ROLLOUT_TYPE,
							Description: "enabled for 50%",
							Rank:        int32(2),
							Rule: &flipt.Rollout_Threshold{
								Threshold: &flipt.RolloutThreshold{
									Percentage: float32(50.0),
									Value:      true,
								},
							},
						},
					},
				},
			},
			path:          "testdata/export_all_namespaces.yml",
			namespaces:    "",
			allNamespaces: true,
		},
	}

	for _, tc := range tests {
		opts := []ExporterOption{WithNamespaces(strings.Split(tc.namespaces, ","))}
		if tc.allNamespaces {
			opts = append(opts, WithAllNamespaces())
		}

		var (
			exporter = NewExporter(tc.lister, opts...)
			b        = new(bytes.Buffer)
			enc      = yaml.NewEncoder(b)
		)

		err := exporter.Export(context.Background(), enc)
		assert.NoError(t, err)

		in, err := os.ReadFile(tc.path)
		assert.NoError(t, err)

		assert.YAMLEq(t, string(in), b.String())
	}
}
