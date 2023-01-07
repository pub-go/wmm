package tmdb

import (
	"context"
	"os"

	"code.gopub.tech/logs"
	"code.gopub.tech/logs/pkg/arg"
	"github.com/cockroachdb/errors"
	tmdb "github.com/cyruzin/golang-tmdb"
)

func SearchTV(ctx context.Context, name string) error {
	client, err := tmdb.Init(GetTmdbKey())
	if err != nil {
		return errors.Wrapf(err, "Init TMDB Api Key Error")
	}
	opt := map[string]string{
		"language": "zh-CN",
	}
	shows, err := client.GetSearchTVShow(name, opt)
	if err != nil {
		return errors.Wrapf(err, "Search TV Show error|%v", name)
	}
	logs.Info(ctx, "Shows: %v", arg.JSON(shows))
	if len(shows.Results) == 1 {
		show := shows.Results[0]
		logs.Info(ctx, "id=%v", show.ID)
		detail, err := client.GetTVDetails(int(show.ID), opt)
		if err != nil {
			return errors.Wrapf(err, "Get TV Detail error")
		}
		logs.Info(ctx, "detail: %v", arg.JSON(detail))
		season, err := client.GetTVSeasonDetails(int(show.ID), 1, opt)
		if err != nil {
			return errors.Wrapf(err, "Get TV season Detail error")
		}
		logs.Info(ctx, "Season detail: %v", arg.JSON(season))
	}
	return nil
}

func GetTmdbKey() string {
	return os.Getenv("TMDB_API_KEY")
}
