package semver

import (
	"reflect"
	"testing"
)

func TestGetIncludedChannels(t *testing.T) {
	type args struct {
		ch string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "testCandidate",
			args: args{"candidate"},
			want: []string{"candidate"},
		},
		{
			name: "testFast",
			args: args{"fast"},
			want: []string{"fast", "candidate"},
		},
		{
			name: "testStable",
			args: args{"stable"},
			want: []string{"stable", "fast", "candidate"},
		},
		{
			name: "testUnknown",
			args: args{"unknown"},
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIncludedChannels(tt.args.ch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIncludedChannels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bundleSlice_add(t *testing.T) {
	type fields struct {
		Bundles []semverVeneerBundleEntry
	}

	type args struct {
		i string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantLen int
	}{
		{
			name: "testEmptyAdd",
			fields: fields{
				Bundles: []semverVeneerBundleEntry{},
			},
			args:    args{"testimage"},
			wantErr: false,
			wantLen: 1,
		},
		{
			name: "testAddUnique",
			fields: fields{
				Bundles: []semverVeneerBundleEntry{
					{
						Image: "testunique",
					},
				},
			},
			args:    args{"testimage"},
			wantErr: false,
			wantLen: 2,
		},
		{
			name: "testAddDuplicate",
			fields: fields{
				Bundles: []semverVeneerBundleEntry{
					{
						Image: "testimage",
					},
				},
			},
			args:    args{"testimage"},
			wantErr: false,
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bundleSlice{
				Bundles: tt.fields.Bundles,
			}
			if err := b.add(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("bundleSlice.add() error = %v, wantErr %v", err, tt.wantErr)
			}
			if l := len(b.Bundles); l != tt.wantLen {
				t.Errorf("bundleSlice.add() length = %d, wantLen %d", l, tt.wantLen)
			}
		})
	}
}

func TestSemverVeneer_AddBundleToChannel(t *testing.T) {
	type fields struct {
		Candidate bundleSlice
		Fast      bundleSlice
		Stable    bundleSlice
	}

	type args struct {
		bundle string
		ch     string
	}

	emptyFields := func() fields {
		return fields{
			bundleSlice{},
			bundleSlice{},
			bundleSlice{},
		}
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantFields fields
	}{
		{
			name:   "testCandidate",
			fields: emptyFields(),
			args: args{
				bundle: "testimage",
				ch:     "candidate",
			},
			wantErr:    false,
			wantFields: emptyFields().Candidate.Bundles[0]{Image: "testimage"},
			wantLen:    1,
			wantLenCh:  channel("candidate"),
		},
		{
			name:   "testFast",
			fields: emptyFields(),
			args: args{
				bundle: "testimage",
				ch:     "fast",
			},
			wantErr:   false,
			wantLen:   1,
			wantLenCh: channel("fast"),
		},
		{
			name:   "testStable",
			fields: emptyFields(),
			args: args{
				bundle: "testimage",
				ch:     "stable",
			},
			wantErr:   false,
			wantLen:   1,
			wantLenCh: channel("stable"),
		},
		{
			name:   "testUnknown",
			fields: emptyFields(),
			args: args{
				bundle: "testimage",
				ch:     "unknown",
			},
			wantErr:   true,
			wantLen:   0,
			wantLenCh: channel("candidate"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sv := &SemverVeneer{
				Candidate: tt.fields.Candidate,
				Fast:      tt.fields.Fast,
				Stable:    tt.fields.Stable,
			}
			if err := sv.AddBundleToChannel(tt.args.bundle, tt.args.ch); (err != nil) != tt.wantErr {
				t.Errorf("SemverVeneer.AddBundleToChannel() error = %v, wantErr %v", err, tt.wantErr)
			}
			if c, _ := sv.getChannel(tt.wantLenCh); len(c.Bundles) != tt.wantLen {
				t.Errorf("SemverVeneer.AddBundleToChannel() len = %d on channel %s, wantLen %d", len(c.Bundles), tt.wantLenCh, tt.wantLen)
			}
		})
	}
}
