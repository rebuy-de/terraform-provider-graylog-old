package e2e

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/ory/dockertest"
	"github.com/sirupsen/logrus"
)

const (
	EnvEnable        = "ACCEPTANCE_TEST"
	EnvPluginDir     = "TERRAFORM_PLUGIN_DIR"
	DefaultPluginDir = "../../dist"
)

func must(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func hasTrueEnvVar(name string) bool {
	return strings.ToLower(os.Getenv(name)) == "true"
}

func TestExampleTemplate(t *testing.T) {
	if !hasTrueEnvVar(EnvEnable) {
		t.Logf("Acceptance test disabled. Enable it by setting env var `%s` to `true`", EnvEnable)
		t.SkipNow()
	}

	if testing.Verbose() {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	pool, err := dockertest.NewPool("")
	must(t, err)

	graylog, err := pool.BuildAndRunWithBuildOptions(&dockertest.BuildOptions{
		ContextDir: "docker",
		Dockerfile: "Dockerfile",
	}, &dockertest.RunOptions{
		Name:         "terraform-provider-graylog-e2e-" + uuid.New().String(),
		ExposedPorts: []string{"9000/tcp"},
	})
	must(t, err)
	defer pool.Purge(graylog)
	must(t, graylog.Expire(3600))

	logrus.Info("waiting for graylog to be ready")
	err = pool.Retry(func() error {
		url := fmt.Sprintf("http://%s/api/system/lbstatus", graylog.GetHostPort("9000/tcp"))
		resp, err := http.Get(url)
		if err != nil {
			logrus.Debug(err)
			return err
		}

		if resp.StatusCode != http.StatusOK {
			logrus.Debug(err)
			return fmt.Errorf("Got status code %d", resp.StatusCode)
		}

		return nil
	})
	must(t, err)

	pluginDir := os.Getenv(EnvPluginDir)
	if pluginDir == "" {
		pluginDir = DefaultPluginDir
	}

	commands := [][]string{
		[]string{"terraform", "init", "-plugin-dir", pluginDir},
		[]string{"terraform", "plan"},
		[]string{"terraform", "apply", "-auto-approve"},
		[]string{"terraform", "plan", "-detailed-exitcode"},
		[]string{"terraform", "destroy", "-auto-approve"},
		[]string{"terraform", "plan", "-detailed-exitcode", "-destroy"},
	}

	for _, args := range commands {
		logrus.Infof("Running %s", strings.Join(args, " "))

		tf := exec.Command(args[0], args[1:]...)
		tf.Dir = "example"
		tf.Stdout = os.Stdout
		tf.Stderr = os.Stderr
		tf.Env = os.Environ()
		tf.Env = append(tf.Env,
			fmt.Sprintf("GRAYLOG_SERVER_URL=http://%s",
				graylog.GetHostPort("9000/tcp")),
			"GRAYLOG_USERNAME=admin",
			"GRAYLOG_PASSWORD=admin",
		)
		must(t, tf.Run())
	}
}
