package main

import "flag"
import "fmt"
import "log"
import "os"
import "plan9/p"
import "plan9/p/clnt"

var addr = flag.String("addr", "127.0.0.1:5640", "network address")

func main() {
	var n int;
	var user p.User;
	var err *p.Error;
	var c *clnt.Clnt;
	var file *clnt.File;

	flag.Parse();
	user = p.OsUsers.Uid2User(os.Geteuid());
	c, err = clnt.Mount("tcp", *addr, "", user);
	if err != nil {
		goto error
	}

	if flag.NArg() != 1 {
		log.Stderr("invalid arguments");
		return;
	}

	file, err = c.FOpen(flag.Arg(0), p.OREAD);
	if err != nil {
		goto error
	}

	buf := make([]byte, 8192);
	for {
		n, err = file.Read(buf);
		if err != nil {
			goto error
		}

		if n == 0 {
			break
		}

		os.Stdout.Write(buf[0:n]);
	}

	file.Close();
	return;

error:
	log.Stderr(fmt.Sprintf("Error: %s %d", err.Error, err.Nerror));
}