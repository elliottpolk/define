package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"github.com/elliottpolk/define/log"

	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

var (
	phraseFlag = cli.StringFlag{
		Name:  "phrase, p",
		Value: "",
		Usage: "phrase to be defined",
	}
	configFlag = cli.BoolFlag{Name: "config, c"}
)

const (
	ggl = "https://www.googleapis.com/customsearch/v1"

	cfgDir  = ".define"
	cfgFile = "config.yml"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Calls the Google Dictionary API and looks up the provided phrase"
	app.Flags = []cli.Flag{phraseFlag, configFlag}
	app.Action = func(context *cli.Context) {
		context.Command.VisibleFlags()

		if context.Bool(simpl(configFlag.Name)) {
			if err := config(); err != nil {
				log.Fatal(err)
			}

			return
		}

		phrase := context.String(simpl(phraseFlag.Name))
		if len(phrase) < 1 {
			if err := cli.ShowCommandHelp(context, context.Command.FullName()); err != nil {
				log.Fatal(err)
			}
			return
		}

		if err := define(phrase); err != nil {
			log.Fatal(err)
		}

	}

	app.Run(os.Args)
}

func config() error {
	var key, cx string

	fmt.Println("Please provide a valid Google API key: ")
	if _, err := fmt.Scan(&key); err != nil {
		return errors.Wrap(err, "unable to read in Google API key")
	}

	fmt.Println("Please provide a valid custom search engine ID ('cx') value: ")
	if _, err := fmt.Scan(&cx); err != nil {
		return errors.Wrap(err, "unable to read in customer search engine ID")
	}

	usr, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "unable to retrieve current user")
	}

	(&Config{Key: key, Cx: cx}).WriteFile(filepath.Join(usr.HomeDir, cfgDir, cfgFile))

	return nil
}

func define(phrase string) error {
	usr, err := user.Current()
	if err != nil {
		return errors.Wrap(err, "unable to retrieve current user")
	}

	cfg, err := GetConf(filepath.Join(usr.HomeDir, cfgDir, cfgFile))
	if err != nil {
		return errors.Wrap(err, "unable to retrieve config")
	}

	from, err := url.Parse(ggl)
	if err != nil {
		return errors.Wrap(err, "unable to parse API URL")
	}

	from.RawQuery = (&url.Values{
		"q":   []string{fmt.Sprintf("define+%s", phrase)},
		"key": []string{cfg.Key},
		"cx":  []string{cfg.Cx},
	}).Encode()

	res, err := http.Get(from.String())
	if err != nil {
		return errors.Wrap(err, "unable to perform GET on API")
	}
	defer res.Body.Close()

	raw, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "unable to read in API response body")
	}

	if code := res.StatusCode; code != http.StatusOK {
		return errors.Errorf("API responsed with status code %d and mesage %s", code, string(raw))
	}

	results := &Result{}
	if err := json.Unmarshal(raw, &results); err != nil {
		return errors.Wrap(err, "unable to unmarshal API response body")
	}

	if results.Queries == nil || len(results.Queries.Request) < 1 {
		log.Infoln("no results")
		return nil
	}

	req := results.Queries.Request[0]
	if corrections := results.Corrections; corrections != nil && len(corrections.Corrected) > 0 {
		log.Infof("showing top 3 of %s total results for corrected terms %s\n", results.Info.TotalResults, corrections.Corrected)
	} else {
		log.Infof("showing top 3 of %s total results for terms %s\n", results.Info.TotalResults, req.SearchTerms)
	}

	for i, r := range results.Items {
		if i > 2 {
			return nil
		}
		log.Infoln(r.Snippet, "\n")
	}
	return nil
}
