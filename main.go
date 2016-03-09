package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	dry := flag.Bool("dry", false, "Just output parsed words")
	uri := flag.String("uri", "", "MySQL URI")
	flag.Parse()

	v := *uri
	if v == "" {
		buf, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalln(err)
			return
		}
		v = string(buf)
	}

	cmd := &Command{URI: v}
	cmd.Options.Dry = *dry

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

// Command ...
type Command struct {
	output  io.Writer
	Options struct {
		Dry bool
	}
	URI string
}

// MySQL ...
type MySQL struct {
	URL *url.URL
}

func (mysql MySQL) options() []string {
	options := []string{
		"--user=" + mysql.URL.User.Username(),
		"--host=" + mysql.URL.Host,
		"--database=" + path.Base(mysql.URL.Path),
	}
	if pass, ok := mysql.URL.User.Password(); ok {
		options = append(options, "--password="+pass)
	}
	return options
}

// Command ...
func (mysql MySQL) Command() *exec.Cmd {
	return exec.Command(
		mysql.URL.Scheme,
		mysql.options()...,
	)
}

// String ...
func (mysql MySQL) String() string {
	options := mysql.options()
	return strings.Join(append([]string{mysql.URL.Scheme}, options...), " ")
}

// Run ...
func (cmd *Command) Run() error {
	u, err := url.Parse(cmd.URI)
	if err != nil {
		return err
	}
	mysql := MySQL{URL: u}

	if cmd.Options.Dry {
		fmt.Println(mysql.String())
		return nil
	}
	// return mysql.Command().Run()
	// return mysql.Command().Start()
	fmt.Println(mysql.String())
	return nil
}
