package config

import (
	"crypto/sha1"
	"encoding/hex"
	"flag"
	. "github.com/daniel-fanjul-alcuten/libmjolnir"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

var (
	v   = flag.Bool("v", false, "show the commands that are run")
	vv  = flag.Bool("vv", false, "like -v, but including the commands that are not run becauseof the cache")
	r   = flag.Bool("r", false, "do not remove temporary files")
	cc  = flag.Bool("cc", false, "use cc compiler")
	gcc = flag.Bool("gcc", false, "use gcc compiler")
	dc  = flag.Bool("dc", true, "use disk cache")
	lc  = flag.Bool("lc", true, "use local cache")
)

func Configure(m *Mjölnir, basepath string) {

	if *v {
		m.Verbose = 1
	} else if *vv {
		m.Verbose = 2
	}

	m.Unclean = *r

	if *cc {
		m.Backend = "cc"
	} else if *gcc {
		m.Backend = "gcc"
	}

	var u *user.User
	if *dc || *lc {
		var err error
		u, err = user.Current()
		if err != nil {
			log.Fatalln("error:", err)
		}
	}

	if *dc || *lc {
		path := filepath.Join(os.TempDir(), "disk-cache-"+u.Uid+".mjolnir")
		m.DataCache = NewDiskDataCache(path)
	}

	if *lc {
		hasher := sha1.New()
		hasher.Write([]byte(basepath))
		hash := hex.EncodeToString(hasher.Sum(nil))
		path := filepath.Join(os.TempDir(), "local-cache-"+u.Uid+"-"+hash+".mjolnir")
		m.LocalCache = NewLocalCache(path)
	}
}
