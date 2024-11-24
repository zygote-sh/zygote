package git

const (
	bitbucketHost = "bitbucket.org"
	githubHost    = "github.com"
	gitlabHost    = "gitlab.com"
	giteaHost     = "gitea.com"
)

var hostMap = map[string]string{
	"bitbucket.org": bitbucketHost,
	"bitbucket":     bitbucketHost,
	"bitb":          bitbucketHost,
	"bit":           bitbucketHost,
	"bb":            bitbucketHost,
	"":              githubHost,
	"github.com":    githubHost,
	"github":        githubHost,
	"gith":          githubHost,
	"ghub":          githubHost,
	"git":           githubHost,
	"gh":            githubHost,
	"gitlab.com":    gitlabHost,
	"gitlab":        gitlabHost,
	"gitl":          gitlabHost,
	"glab":          gitlabHost,
	"glb":           gitlabHost,
	"gl":            gitlabHost,
	"gitea.com":     giteaHost,
	"gitea":         giteaHost,
	"gite":          giteaHost,
	"gt":            giteaHost,
}
