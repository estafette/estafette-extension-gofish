package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"text/template"

	"github.com/alecthomas/kingpin"
	foundation "github.com/estafette/estafette-foundation"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

var (
	appgroup  string
	app       string
	version   string
	branch    string
	revision  string
	buildDate string
	goVersion = runtime.Version()
)

var (
	paramsYAML = kingpin.Flag("params-yaml", "Extension parameters, created from custom properties.").Envar("ESTAFETTE_EXTENSION_CUSTOM_PROPERTIES_YAML").Required().String()

	buildVersion = kingpin.Flag("build-version", "Version number, used if not passed explicitly.").Envar("ESTAFETTE_BUILD_VERSION").String()
	repoSource   = kingpin.Flag("repo-source", "Hostname for vsc.").Envar("ESTAFETTE_GIT_SOURCE").String()
	repoOwner    = kingpin.Flag("repo-owner", "Owner for vsc.").Envar("ESTAFETTE_GIT_OWNER").String()
	repoName     = kingpin.Flag("repo-name", "Repository for vsc.").Envar("ESTAFETTE_GIT_NAME").String()
)

func main() {

	// parse command line parameters
	kingpin.Parse()

	// init log format from envvar ESTAFETTE_LOG_FORMAT
	foundation.InitLoggingFromEnv(appgroup, app, version, branch, revision, buildDate)

	log.Info().Msg("Unmarshalling parameters / custom properties...")
	var params Params
	err := yaml.Unmarshal([]byte(*paramsYAML), &params)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed unmarshalling parameters")
	}

	// set defaults
	params.SetDefaults(*buildVersion, *repoSource, *repoOwner, *repoName)

	// validate parameters
	valid, warnings := params.Validate()
	if !valid {
		log.Fatal().Msgf("Some parameters are not valid: %v", warnings)
	}

	// create target file to render template to
	targetFilePath := fmt.Sprintf("%v/%v/%v.lua", params.RigReposityDirectory, params.FoodDirectory, params.Name)
	targetFile, err := os.Create(targetFilePath)
	defer targetFile.Close()

	// read and parse template
	foodTemplate, err := template.ParseFiles("/templates/food.lua")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed parsing template")
	}

	params.DarwinSha256 = getSha256(params.DarwinDownloadUrl)
	params.LinuxSha256 = getSha256(params.LinuxDownloadUrl)
	params.WindowsSha256 = getSha256(params.WindowsDownloadUrl)

	// write rendered template to file
	err = foodTemplate.Execute(targetFile, params)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed writing template to file")
	}
}

func getSha256(downloadUrl string) string {
	// download binary
	resp, err := http.Get(downloadUrl)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed retrieving binary from url %v", downloadUrl)
	}
	defer resp.Body.Close()

	// calculate sha256 checksum
	hasher := sha256.New()
	if _, err := io.Copy(hasher, resp.Body); err != nil {
		log.Fatal().Err(err).Msgf("Failed calculating sha256 checksum for binary from url %v", downloadUrl)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}
