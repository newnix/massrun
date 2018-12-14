//
// Copyright (c) 2018, Exile Heavy Industries
// All rights reserved.
// 
// Redistribution and use in source and binary forms, with or without
// modification, are permitted (subject to the limitations in the disclaimer
// below) provided that the following conditions are met:
// 
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
// 
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
// 
// * Neither the name of the copyright holder nor the names of its contributors may be used
//   to endorse or promote products derived from this software without specific
//   prior written permission.
// 
// NO EXPRESS OR IMPLIED LICENSES TO ANY PARTY'S PATENT RIGHTS ARE GRANTED BY THIS
// LICENSE. THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
// THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE
// GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT
// OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH
// DAMAGE.
//

// This is a F/LOSS reimplementation of a similar tool I wrote for work
// it essentially acts as a means to concurrently run a large number of utilities
// typically something like expect scripts to update or fetch device configurations
// The version written for work was capable of updating 1700+ devices in under a minute,
// as well as generate backups of their configurations.
package main;

import ( 
	"flag";
	"fmt";
	"io/ioutil";
	"strings";
)

var authkey string
var custArgs string
var dbg bool
var conffile string
var help bool
var internal bool
var logfile string
var password string
var readlist string
var script string
var user string

// Read config file key/value style
func readconfig(file string) error {
	conf, err := ioutil.ReadFile(file); if err != nil {
		return(err)
	}

	// Split up the config data into a bunch of strings
	confstr := strings.Split(string(conf), "\n")

	if (dbg) {
		fmt.Println("Read:")
		fmt.Println(confstr)
	}

	for i := 0; i < len(confstr) - 1; i++ {
		option := strings.Split(confstr[i], "=")
		switch option[0] {
		case "keyfile":
			authkey = option[1]
		case "argv_override":
			custArgs = option[1]
		case "debug":
			if option[1] == "true" {
				dbg = true
			} else {
				dbg = false
			}
		case "internal":
			if option[1] == "true" {
				internal = true
			} else {
				internal = false
			}
		case "logfile":
			logfile = option[1]
		case "password":
			password = option[1]
		case "list":
			readlist = option[1]
		case "script":
			script = option[1]
		case "user":
			script = option[1]
		default:
			// ignore invalid entries
			;
		}
	}
	// Default, unknown values are ignored
	return(nil)
}

// Flag setting/usage generation
func init() {
	flag.StringVar(&authkey, "k", "", "`keyfile` to use for ssh authentication")
	flag.StringVar(&custArgs, "A", "", "`arguments` to pass for the script being run, will override other flags")
	flag.BoolVar(&dbg, "d", false, "Print out some debugging information")
	flag.StringVar(&conffile,"f", "", "`file` to use for authentication data")
	flag.BoolVar(&help, "h", false, "This help message")
	flag.BoolVar(&internal, "i", false, "Toggle the use of internal command handling [not yet supported]")
	flag.StringVar(&logfile, "l", "", "`file` to write logging data to")
	flag.StringVar(&password, "sp", "", "`password` for connecting to the remote hosts/devices")
	flag.StringVar(&readlist, "L", "", "`list` of hosts to operate on")
	flag.StringVar(&script, "S", "", "`script` to run against the host list")
	flag.StringVar(&user, "su", "", "`user` account to use when connecting to remote hosts/devices")
}

func main() {
	flag.Parse()
	if flag.NFlag() == 0 || help {
		flag.Usage()
	}
	if conffile != "" {
		readconfig(conffile)
	}
}
