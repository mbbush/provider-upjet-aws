package clients

import "testing"

func TestPartitions(t *testing.T) {
	cases := map[string]struct {
		partition        string
		servicesToRegion map[string]string
		wantOk           bool
	}{
		"AWSDefaultPartition": {
			partition: "aws",
			servicesToRegion: map[string]string{
				"iam": "us-east-1",
			},
			wantOk: true,
		},
		"AWSChina": {
			partition: "aws-cn",
			servicesToRegion: map[string]string{
				"iam": "cn-north-1",
			},
			wantOk: true,
		},
		"AWSISO": {
			partition: "aws-iso",
			servicesToRegion: map[string]string{
				"iam": "us-iso-east-1",
			},
			wantOk: true},
		"AWSISOB": {
			partition:        "aws-iso-b",
			servicesToRegion: map[string]string{},
			wantOk:           true,
		},
		"AWSISOE": {
			partition:        "aws-iso-e",
			servicesToRegion: map[string]string{},
			wantOk:           true,
		},
		"AWSISOF": {
			partition:        "aws-iso-f",
			servicesToRegion: map[string]string{},
			wantOk:           true,
		},
		"AWSUSGov": {
			partition: "aws-us-gov",
			servicesToRegion: map[string]string{
				"iam": "us-gov-west-1",
			},
			wantOk: true,
		},
		"NonExistentPartition": {
			partition:        "aws-foo",
			servicesToRegion: map[string]string{},
			wantOk:           false,
		},
	}
	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			partition, ok := partitions[tc.partition]
			if ok != tc.wantOk {
				t.Errorf("expected partition existence: got %v, want %v", ok, tc.wantOk)
			}
			for wantSvc, wantRegion := range tc.servicesToRegion {
				reg, okSvc := partition.serviceToDefaultRegions[wantSvc]
				if !okSvc {
					t.Errorf("expected service %q to exist in partition %q", wantSvc, tc.partition)
				}
				if reg != wantRegion {
					t.Errorf("expected default region %q for service %q in partition %q, got region %q", wantRegion, wantSvc, tc.partition, reg)
				}
			}
		})
	}
	var defaultRegions = map[string]string{}
	for name, partition := range partitions {
		reg, ok := partition.serviceToDefaultRegions["iam"]
		if !ok {
			continue
		}
		defaultRegions[name] = reg
	}
}

func TestGetIAMDefaultSigningRegions(t *testing.T) {
	cases := map[string]struct {
		partition string
		region    string
		wantOk    bool
	}{
		"AWSDefaultPartition":  {"aws", "us-east-1", true},
		"AWSChina":             {"aws-cn", "cn-north-1", true},
		"AWSISO":               {"aws-iso", "us-iso-east-1", true},
		"AWSISOB":              {"aws-iso-b", "us-isob-east-1", true},
		"AWSISOE":              {"aws-iso-e", "", false},
		"AWSISOF":              {"aws-iso-f", "us-isof-south-1", true},
		"AWSUSGov":             {"aws-us-gov", "us-gov-west-1", true},
		"NonExistentPartition": {"aws-foo", "", false},
	}
	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			defaultRegions := getIAMDefaultSigningRegions()
			region, ok := defaultRegions[tc.partition]
			if ok != tc.wantOk {
				t.Errorf("expected partition existence: got %v, want %v", ok, tc.wantOk)
			}
			if region != tc.region {
				t.Errorf("expected default region for partition %v: got %v, want %v", tc.partition, region, tc.region)
			}
		})
	}
}
