package merge

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestFetch(t *testing.T) {
	type args struct {
		token   string
		baseUrl string
		mrs     []int
		filter  map[string]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []MRResult
		wantErr    bool
	}{
		{
			name: "t1 - fluxo feliz",
			args: args{
				token:   os.Getenv("PRIVATE_TOKEN"),
				baseUrl: "https://www-scm.prevnet/api/v3/projects",
				mrs:     []int{8893},
				filter:  nil,
			},
			wantResult: []MRResult{
				{
					Merge{
						Iid: 8893,
						Author: User{
							Username: "alexandre.alessi",
						},
						TargetBranch: "desenvolvimento",

						WebUrl:         "https://www-scm.prevnet/sibe-pu/sibe-pu-repo/merge_requests/8893",
						MergeCommitSha: "f4c84edd6a3b5736b06d3a31b4f1b7fd36a3f5a5",
					},
					Commit{
						Id:        "f4c84edd6a3b5736b06d3a31b4f1b7fd36a3f5a5",
						Email:     "",
						CreatedAt: "2020-11-14T16:31:47.000-03:00",
						Username:  "isabel.tiburski",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "t2 - array de ids",
			args: args{
				token:   os.Getenv("PRIVATE_TOKEN"),
				baseUrl: "https://www-scm.prevnet/api/v3/projects",
				mrs:     []int{8888, 8890},
				filter:  nil,
			},
			wantResult: []MRResult{
				{
					Merge{
						Iid: 8888,
						Author: User{
							Username: "alexandre.riso",
						},
						TargetBranch: "desenvolvimento",

						WebUrl:         "https://www-scm.prevnet/sibe-pu/sibe-pu-repo/merge_requests/8888",
						MergeCommitSha: "f4c84edd6a3b5736b06d3a31b4f1b7fd36a3f5a5",
					},
					Commit{
						Id:        "f4c84edd6a3b5736b06d3a31b4f1b7fd36a3f5a5",
						Email:     "",
						CreatedAt: "2020-11-13T12:22:35.000-03:00",
						Username:  "isabel.tiburski",
					},
				},
				{
					Merge{
						Iid: 8890,
						Author: User{
							Username: "isabel.tiburski",
						},
						TargetBranch: "homologacao",

						WebUrl:         "https://www-scm.prevnet/sibe-pu/sibe-pu-repo/merge_requests/8890",
						MergeCommitSha: "f4c84edd6a3b5736b06d3a31b4f1b7fd36a3f5a5",
					},
					Commit{
						Id:        "f4c84edd6a3b5736b06d3a31b4f1b7fd36a3f5a5",
						Email:     "",
						CreatedAt: "2020-11-13T13:59:16.000-03:00",
						Username:  "isabel.tiburski",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := fetch(tt.args.token, tt.args.baseUrl, "754", tt.args.mrs, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, r := range gotResult {
				wR := tt.wantResult[i]
				if r.Merge.Author != wR.Merge.Author {
					t.Errorf("fetch() gotResult = %v, want %v\n", r.Merge.Author, wR.Merge.Author)
				}

				if r.Merge.TargetBranch != wR.Merge.TargetBranch {
					t.Errorf("fetch() gotResult = %v, want %v\n", r.Merge.TargetBranch, wR.Merge.TargetBranch)
				}
				if r.MergeCommit.Username != wR.MergeCommit.Username {
					t.Errorf("fetch() gotResult = %v, want %v\n", r.MergeCommit.Username, wR.MergeCommit.Username)
				}
				if r.Merge.WebUrl != wR.Merge.WebUrl {
					t.Errorf("fetch() gotResult = %v, want %v\n", r.Merge.WebUrl, wR.Merge.WebUrl)
				}
				if r.MergeCommit.CreatedAt != wR.MergeCommit.CreatedAt {
					t.Errorf("fetch() gotResult = %v, want %v\n", r.MergeCommit.CreatedAt, wR.MergeCommit.CreatedAt)

				}
				fmt.Println(r)

			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				//t.Errorf("fetch() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_extractIds(t *testing.T) {
	type args struct {
		urls []string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
		wantErr    bool
	}{
		{
			name: "t1",
			args: args{
				urls: []string{
					"https://www-scm.prevnet/sibe-pu/sibe-pu-repo/merge_requests/8893",
					"https://www-scm.prevnet/sibe-pu/sibe-pu-repo/merge_requests/7777",
					"3333https://www-scm.prevnet/sibe-pu/sibe-pu-repo/merge_requests/1000",
					"3333https://www-scm.prevnet/sibe-pu/sibe-pu-repo11111/merge_requests1221/1000_1111",
				},
			},
			wantResult: []int{
				8893, 7777, 1000, 1111,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := extractIds(tt.args.urls)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("extractIds() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestGetFetchMode(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name       string
		args       args
		wantRanges []string
		wantErr    bool
	}{
		{
			name: "t1 interval",
			args: args{
				args: []string{"pname",
					"3333:4444",
				},
			},
			wantRanges: []string{

			},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRanges, err := GetFetchMode(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFetchMode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotRanges != nil && !reflect.DeepEqual(gotRanges, tt.wantRanges) {
				t.Errorf("GetFetchMode() gotRanges = %v, want %v", gotRanges, tt.wantRanges)
			}
		})
	}
}
