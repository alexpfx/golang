package merge

import "testing"

func TestFormatOutput(t *testing.T) {
	type args struct {
		jsonInput string
		output    string
	}
	const input = `[{
      "merge": {
         "iid": 7005,
         "title": "Revert \"Merge branch 'feature/217447/inclusao-bi-com-atestado-adicao-cpf-como-entrada' i…\"",
         "description": "This reverts merge request !7004",
         "state": "merged",
         "created_at": "2020-04-06T18:47:12.116-03:00",
         "updated_at": "2020-04-08T12:44:55.427-03:00",
         "target_branch": "desenvolvimento",
         "source_branch": "revert-d658f5b8",
         "web_url": "https://www-scm.prevnet/sibe-pu/sibe-pu-repo/merge_requests/7005",
         "author": {
            "username": "alexandre.alessi"
         },
         "merge_commit_sha": "28235f2feebf2519236c1b5b39524298e3244d94",
         "merge_commit": {
			 "id": "28235f2feebf2519236c1b5b39524298e3244d94",
			 "author_email": "alexandre.alessi@dataprev.gov.br",
			 "created_at": "2020-04-06T19:00:56.000-03:00",
			 "username": "alexandre.alessi"
      		}
      }
      
   },
   {
      "merge": {
         "iid": 7006,
         "title": "Revert \"Merge branch 'revert-d658f5b8' into 'desenvolvimento'\"",
         "description": "This reverts merge request !7005",
         "state": "merged",
         "created_at": "2020-04-06T19:04:53.984-03:00",
         "updated_at": "2020-04-08T12:44:55.409-03:00",
         "target_branch": "desenvolvimento",
         "source_branch": "revert-28235f2f",
         "web_url": "https://www-scm.prevnet/sibe-pu/sibe-pu-repo/merge_requests/7006",
         "author": {
            "username": "alexandre.alessi"
         },
         "merge_commit_sha": "15fbeae7a342b90b7ceda33b949b44997563b41c",
		"merge_commit": {
				 "id": "15fbeae7a342b90b7ceda33b949b44997563b41c",
				 "author_email": "thiago.zimmermann@dataprev.gov.br",
				 "created_at": "2020-04-06T19:08:18.000-03:00",
				 "username": "thiago.zimmermann"
			  }

      	}
      ]`

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				output:    ".author.username",
				jsonInput: input,
			},
			want: "\nalexandre.alessi\t\nalexandre.alessi\t\n",
		},
		{
			name: "t2",
			args: args{
				output:    ".web_url .author.username .merge_commit.username .merge_commit.created_at",
				jsonInput: input,
			},
			want: `
https://www-scm.prevnet/sibe-pu/sibe-pu-repo/merge_requests/7005	alexandre.alessi	alexandre.alessi	2020-04-06T19:00:56.000-03:00	
https://www-scm.prevnet/sibe-pu/sibe-pu-repo/merge_requests/7006	alexandre.alessi	thiago.zimmermann	2020-04-06T19:08:18.000-03:00	
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatString(tt.args.jsonInput, NewFormatter(tt.args.output)); got != tt.want {
				t.Errorf("FormatString():[%v], want:[%v]", got, tt.want)
			}
		})
	}
}
