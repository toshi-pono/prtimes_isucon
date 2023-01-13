// Code generated by qtc from "banned.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line banned.qtpl:1
package templates

//line banned.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line banned.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line banned.qtpl:1
func StreamBanned(qw422016 *qt422016.Writer, users []User, csrfToken string) {
//line banned.qtpl:1
	qw422016.N().S(`
<div>
  <form method="post" action="/admin/banned">
    `)
//line banned.qtpl:4
	for i := range users {
//line banned.qtpl:4
		qw422016.N().S(`
    <div>
      <input type="checkbox" name="uid[]" id="uid_`)
//line banned.qtpl:6
		qw422016.N().D(users[i].ID)
//line banned.qtpl:6
		qw422016.N().S(`" value="`)
//line banned.qtpl:6
		qw422016.N().D(users[i].ID)
//line banned.qtpl:6
		qw422016.N().S(`" data-account-name="`)
//line banned.qtpl:6
		qw422016.E().S(users[i].AccountName)
//line banned.qtpl:6
		qw422016.N().S(`"> <label for="uid_`)
//line banned.qtpl:6
		qw422016.N().D(users[i].ID)
//line banned.qtpl:6
		qw422016.N().S(`">`)
//line banned.qtpl:6
		qw422016.E().S(users[i].AccountName)
//line banned.qtpl:6
		qw422016.N().S(`</label>
    </div>
    `)
//line banned.qtpl:8
	}
//line banned.qtpl:8
	qw422016.N().S(`
    <div class="form-submit">
      <input type="hidden" name="csrf_token" value="`)
//line banned.qtpl:10
	qw422016.E().S(csrfToken)
//line banned.qtpl:10
	qw422016.N().S(`">
      <input type="submit" name="submit" value="submit">
    </div>
  </form>
</div>
`)
//line banned.qtpl:15
}

//line banned.qtpl:15
func WriteBanned(qq422016 qtio422016.Writer, users []User, csrfToken string) {
//line banned.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line banned.qtpl:15
	StreamBanned(qw422016, users, csrfToken)
//line banned.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line banned.qtpl:15
}

//line banned.qtpl:15
func Banned(users []User, csrfToken string) string {
//line banned.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line banned.qtpl:15
	WriteBanned(qb422016, users, csrfToken)
//line banned.qtpl:15
	qs422016 := string(qb422016.B)
//line banned.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line banned.qtpl:15
	return qs422016
//line banned.qtpl:15
}
