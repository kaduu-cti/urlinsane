// Copyright © 2018 rangertaha rangertaha@gmail.com
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package urlinsane

import (
	"fmt"
	"strings"

	"github.com/bobesa/go-domain-util/domainutil"
	"github.com/kaduu-cti/urlinsane/languages"
	"github.com/spf13/cobra"
)

type BasicConfig struct {
	Domains     []string `json:"domains,omitempty"`
	Keyboards   []string `json:"keyboards,omitempty"`
	Typos       []string `json:"typos,omitempty"`
	Funcs       []string `json:"funcs,omitempty"`
	Concurrency int      `json:"concurrency,omitempty"`
	Format      string   `json:"format,omitempty"`
	File        string   `json:"file,omitempty"`
	Verbose     bool     `json:"verbose,omitempty"`
}

type Config struct {
	domains   []Domain
	keyboards []languages.Keyboard
	// languages   []languages.Language
	typos       []Typo
	funcs       []Extra
	headers     []string
	concurrency int

	format  string
	file    string
	verbose bool
}

// Config creates a Config
func (b *BasicConfig) Config() (c Config) {
	// Basic options
	c.GetDomains(b.Domains)
	c.GetKeyboards(b.Keyboards)

	// Registered functions
	c.GetTypos(b.Typos)
	c.GetFuncs(b.Funcs)

	// Processing option
	c.GetConcurrency(b.Concurrency)

	// Output options
	c.GetFile(b.File)
	c.GetFormat(b.Format)
	c.GetVerbose(b.Verbose)

	// Requires c.GetFuncs()
	c.GetHeaders(c.funcs)

	return
}

// GetDomains
func (c *Config) GetDomains(args []string) {
	dmns := []Domain{}
	for _, str := range args {
		subdomain := domainutil.Subdomain(str)
		domain := domainutil.DomainPrefix(str)
		suffix := domainutil.DomainSuffix(str)
		if domain == "" {
			domain = str
		}
		dmns = append(dmns, Domain{
			Subdomain: subdomain,
			Domain:    domain,
			Suffix:    suffix})
	}
	c.domains = dmns
}

// GetKeyboards retrieves a list of keyboards
func (c *Config) GetKeyboards(keyboards []string) {
	c.keyboards = languages.KEYBOARDS.Keyboards(keyboards...)
}

// GetTypos
func (c *Config) GetTypos(typos []string) {
	c.typos = TRetrieve(typos...)
}

// GetFuncs
func (c *Config) GetFuncs(funcs []string) {
	if funcs := FRetrieve(funcs...); len(funcs) > 0 {
		c.funcs = funcs
	} else {
		c.funcs = FRetrieve("idna")
	}
}

// GetHeaders
func (c *Config) GetHeaders(funcs []Extra) {
	c.headers = []string{"Live", "Type", "Typo", "Suffix"}
	for _, fnc := range funcs {
		for _, h := range fnc.Headers {
			c.headers = appendIfMissing(c.headers, h)
		}
	}
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

// GetConcurrency
func (c *Config) GetConcurrency(concurrency int) {
	c.concurrency = concurrency
}

// GetFile
func (c *Config) GetFile(file string) {
	c.file = file
}

// GetFormat
func (c *Config) GetFormat(format string) {
	c.format = format
}

// GetVerbose
func (c *Config) GetVerbose(verbose bool) {
	c.verbose = verbose
}

// errHandler
func errHandler(err error) {
	// TODO
}

// CobraConfig creates a configuration from a cobra command and arguments
func CobraConfig(cmd *cobra.Command, args []string) (c Config) {

	// Print logo
	fmt.Println(LOGO)

	// Basic options
	c.GetDomains(args)

	keyboards, err := cmd.PersistentFlags().GetStringArray("keyboards")
	errHandler(err)
	c.GetKeyboards(keyboards)

	// Registered functions
	var algorithms []string
	typos, err := cmd.PersistentFlags().GetStringArray("typos")
	for _, typo := range typos {
		algorithms = append(algorithms, strings.ToUpper(typo))
	}
	errHandler(err)
	c.GetTypos(algorithms)

	var funcs []string
	functions, err := cmd.PersistentFlags().GetStringArray("funcs")
	for _, function := range functions {
		funcs = append(funcs, strings.ToUpper(function))
	}
	errHandler(err)
	c.GetFuncs(funcs)

	// Processing option
	concurrency, err := cmd.PersistentFlags().GetInt("concurrency")
	errHandler(err)
	c.GetConcurrency(concurrency)

	// Output options
	file, err := cmd.PersistentFlags().GetString("file")
	errHandler(err)
	c.GetFile(file)

	format, err := cmd.PersistentFlags().GetString("format")
	errHandler(err)
	c.GetFormat(format)

	verbose, err := cmd.PersistentFlags().GetBool("verbose")
	errHandler(err)
	c.GetVerbose(verbose)

	// Requires c.funcs
	c.GetHeaders(c.funcs)

	return
}
