package main

import "fmt"

// Params are the parameters passed to this extension via custom stage properties
type Params struct {
	Name                 string `json:"name,omitempty" yaml:"name,omitempty"`
	Version              string `json:"version,omitempty" yaml:"version,omitempty"`
	Description          string `json:"description,omitempty" yaml:"description,omitempty"`
	Homepage             string `json:"homepage,omitempty" yaml:"homepage,omitempty"`
	Repository           string `json:"repository,omitempty" yaml:"repository,omitempty"`
	RigReposityDirectory string `json:"rigRepoDir,omitempty" yaml:"rigRepoDir,omitempty"`
	FoodDirectory        string `json:"foodDir,omitempty" yaml:"foodDir,omitempty"`

	DarwinDownloadUrl string `json:"darwinDownloadUrl,omitempty" yaml:"darwinDownloadUrl,omitempty"`
	DarwinSha256      string `json:"darwinSha256,omitempty" yaml:"darwinSha256,omitempty"`

	LinuxDownloadUrl string `json:"linuxDownloadUrl,omitempty" yaml:"linuxDownloadUrl,omitempty"`
	LinuxSha256      string `json:"linuxSha256,omitempty" yaml:"linuxSha256,omitempty"`

	WindowsDownloadUrl string `json:"windowsDownloadUrl,omitempty" yaml:"windowsDownloadUrl,omitempty"`
	WindowsSha256      string `json:"windowsSha256,omitempty" yaml:"windowsSha256,omitempty"`
}

// SetDefaults sets some sane defaults
func (p *Params) SetDefaults(buildVersion, repoSource, repoOwner, repoName string) {
	if p.Name == "" {
		p.Name = repoName
	}
	if p.Version == "" {
		p.Version = buildVersion
	}
	if p.Homepage == "" {
		p.Homepage = fmt.Sprintf("https://%v/%v/%v", repoSource, repoOwner, repoName)
	}
	if p.Repository == "" {
		p.Repository = fmt.Sprintf("https://%v/%v/%v", repoSource, repoOwner, repoName)
	}
	if p.RigReposityDirectory == "" {
		p.RigReposityDirectory = "fish-food"
	}
	if p.FoodDirectory == "" {
		p.FoodDirectory = "Food"
	}

	if p.DarwinDownloadUrl == "" {
		p.DarwinDownloadUrl = fmt.Sprintf("%v/releases/download/v%v/%v-v%v-darwin-amd64.zip", p.Repository, p.Version, p.Name, p.Version)
	}
	if p.LinuxDownloadUrl == "" {
		p.LinuxDownloadUrl = fmt.Sprintf("%v/releases/download/v%v/%v-v%v-linux-amd64.zip", p.Repository, p.Version, p.Name, p.Version)
	}
	if p.WindowsDownloadUrl == "" {
		p.WindowsDownloadUrl = fmt.Sprintf("%v/releases/download/v%v/%v-v%v-windows-amd64.exe.zip", p.Repository, p.Version, p.Name, p.Version)
	}
}

// Validate checks if the parameters are valid
func (p *Params) Validate() (valid bool, warnings []string) {
	return len(warnings) == 0, warnings
}
