package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/Orfeo42/sqlclass"
)

func init() {
	register.Plugin("sqlclass", New)
}

// Settings mirrors sqlclass.Config for golangci-lint's settings decoding.
type Settings struct {
	ExcludePaths      string `json:"exclude-paths"`
	SQLUtilityPaths   string `json:"sql-utility-paths"`
	SQLUtilityPackage string `json:"sql-utility-package"`
	DomainPackage     string `json:"domain-package"`
	DaoPackage        string `json:"dao-package"`
	AssertPackage     string `json:"assert-package"`
}

// New decodes settings and builds the sqlclass linter plugin.
func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[Settings](settings)
	if err != nil {
		return nil, err
	}

	return &plugin{settings: s}, nil
}

type plugin struct {
	settings Settings
}

// BuildAnalyzers returns the sqlclass analyzer configured from the plugin
// settings. An unset package selector stays empty here and New fills it with
// the analyzer's own default.
func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		sqlclass.New(sqlclass.Config{
			ExcludePaths:      p.settings.ExcludePaths,
			SQLUtilityPaths:   p.settings.SQLUtilityPaths,
			SQLUtilityPackage: p.settings.SQLUtilityPackage,
			DomainPackage:     p.settings.DomainPackage,
			DaoPackage:        p.settings.DaoPackage,
			AssertPackage:     p.settings.AssertPackage,
		}),
	}, nil
}

// GetLoadMode reports that sqlclass reads syntax only: it matches call
// shapes and identifier names, never resolved types.
func (p *plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
