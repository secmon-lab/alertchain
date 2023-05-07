package cli

import (
	"io"
	"os"
	"path/filepath"

	"github.com/m-mizutani/alertchain/pkg/chain"
	"github.com/m-mizutani/alertchain/pkg/domain/model"
	"github.com/m-mizutani/alertchain/pkg/domain/types"
	"github.com/m-mizutani/alertchain/pkg/infra/logger"
	"github.com/m-mizutani/goerr"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slog"
)

func cmdPlay(cfg *model.Config) *cli.Command {
	var (
		playbookPath string
		enablePrint  bool
		outDir       string
	)

	return &cli.Command{
		Name:    "play",
		Aliases: []string{"p"},
		Usage:   "Simulate alertchain policy",
		Flags: []cli.Flag{

			&cli.StringFlag{
				Name:        "playbook",
				Aliases:     []string{"b"},
				Usage:       "playbook file",
				EnvVars:     []string{"ALERTCHAIN_PLAYBOOK"},
				Required:    true,
				Destination: &playbookPath,
			},
			&cli.BoolFlag{
				Name:        "enable-print",
				Aliases:     []string{"p"},
				Usage:       "enable print feature in Rego",
				EnvVars:     []string{"ALERTCHAIN_ENABLE_PRINT"},
				Destination: &enablePrint,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "output directory",
				EnvVars:     []string{"ALERTCHAIN_OUTPUT"},
				Destination: &outDir,
				Value:       "./output",
			},
		},

		Action: func(c *cli.Context) error {
			var baseOptions []chain.Option

			// Load playbook
			var playbook model.Playbook
			if err := model.ParsePlaybook(playbookPath, os.ReadFile, &playbook); err != nil {
				return goerr.Wrap(err, "failed to parse playbook")
			}

			ctx := model.NewContext(model.WithBase(c.Context))
			ctx.Logger().Info("starting alertchain with play mode", slog.Any("playbook", playbookPath))

			for _, s := range playbook.Scenarios {
				ctx.Logger().Debug("Start scenario",
					slog.String("id", string(s.ID)),
					slog.String("title", string(s.Title)),
					slog.Any("event", s.Event),
				)

				w, err := openLogFile(outDir, string(s.ID))
				if err != nil {
					return err
				}
				defer func() {
					if err := w.Close(); err != nil {
						ctx.Logger().Warn("Failed to close log file", slog.String("err", err.Error()))
					}
				}()

				lg := logger.NewJSONLogger(w, s)
				options := append(baseOptions,
					chain.WithScenarioLogger(lg),
					chain.WithActionMock(s),
				)

				if playbook.Env != nil {
					options = append(options, chain.WithEnv(func() types.EnvVars {
						return playbook.Env
					}))
				}

				chain, err := buildChain(*cfg, options...)
				if err != nil {
					return err
				}

				if err := chain.HandleAlert(ctx, s.Schema, s.Event); err != nil {
					lg.LogError(err)
				}
			}

			return nil
		},
	}
}

func openLogFile(dir, name string) (io.WriteCloser, error) {
	dirName := filepath.Clean(filepath.Join(dir, name))
	// #nosec G301
	if err := os.MkdirAll(dirName, 0755); err != nil {
		return nil, goerr.Wrap(err, "Failed to create scenario logging directory")
	}

	path := filepath.Join(dirName, "data.json")
	fd, err := os.Create(filepath.Clean(path))
	if err != nil {
		return nil, goerr.Wrap(err, "Failed to create scenario logging file")
	}

	return fd, nil
}
