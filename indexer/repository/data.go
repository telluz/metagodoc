package repository

// This was all copied from https://github.com/golang/gddo/blob/master/gosrc/data.go

const (
	goRepoPath  = 1
	packagePath = 2
)

var pathFlags = map[string]int{
	"C":                                                       2,
	"archive":                                                 1,
	"archive/tar":                                             3,
	"archive/zip":                                             3,
	"bufio":                                                   3,
	"builtin":                                                 3,
	"bytes":                                                   3,
	"cmd":                                                     1,
	"cmd/addr2line":                                           3,
	"cmd/api":                                                 3,
	"cmd/asm":                                                 3,
	"cmd/asm/internal":                                        1,
	"cmd/asm/internal/arch":                                   3,
	"cmd/asm/internal/asm":                                    3,
	"cmd/asm/internal/flags":                                  3,
	"cmd/asm/internal/lex":                                    3,
	"cmd/cgo":                                                 3,
	"cmd/compile":                                             3,
	"cmd/compile/internal":                                    1,
	"cmd/compile/internal/amd64":                              3,
	"cmd/compile/internal/arm":                                3,
	"cmd/compile/internal/arm64":                              3,
	"cmd/compile/internal/gc":                                 3,
	"cmd/compile/internal/mips":                               3,
	"cmd/compile/internal/mips64":                             3,
	"cmd/compile/internal/ppc64":                              3,
	"cmd/compile/internal/s390x":                              3,
	"cmd/compile/internal/ssa":                                3,
	"cmd/compile/internal/syntax":                             3,
	"cmd/compile/internal/test":                               3,
	"cmd/compile/internal/types":                              3,
	"cmd/compile/internal/x86":                                3,
	"cmd/cover":                                               3,
	"cmd/dist":                                                3,
	"cmd/doc":                                                 3,
	"cmd/fix":                                                 3,
	"cmd/go":                                                  3,
	"cmd/go/internal":                                         1,
	"cmd/go/internal/base":                                    3,
	"cmd/go/internal/bug":                                     3,
	"cmd/go/internal/buildid":                                 3,
	"cmd/go/internal/cfg":                                     3,
	"cmd/go/internal/clean":                                   3,
	"cmd/go/internal/cmdflag":                                 3,
	"cmd/go/internal/doc":                                     3,
	"cmd/go/internal/envcmd":                                  3,
	"cmd/go/internal/fix":                                     3,
	"cmd/go/internal/fmtcmd":                                  3,
	"cmd/go/internal/generate":                                3,
	"cmd/go/internal/get":                                     3,
	"cmd/go/internal/help":                                    3,
	"cmd/go/internal/list":                                    3,
	"cmd/go/internal/load":                                    3,
	"cmd/go/internal/run":                                     3,
	"cmd/go/internal/str":                                     3,
	"cmd/go/internal/test":                                    3,
	"cmd/go/internal/tool":                                    3,
	"cmd/go/internal/version":                                 3,
	"cmd/go/internal/vet":                                     3,
	"cmd/go/internal/web":                                     3,
	"cmd/go/internal/work":                                    3,
	"cmd/gofmt":                                               3,
	"cmd/internal":                                            1,
	"cmd/internal/bio":                                        3,
	"cmd/internal/browser":                                    3,
	"cmd/internal/dwarf":                                      3,
	"cmd/internal/gcprog":                                     3,
	"cmd/internal/goobj":                                      3,
	"cmd/internal/obj":                                        3,
	"cmd/internal/obj/arm":                                    3,
	"cmd/internal/obj/arm64":                                  3,
	"cmd/internal/obj/mips":                                   3,
	"cmd/internal/obj/ppc64":                                  3,
	"cmd/internal/obj/s390x":                                  3,
	"cmd/internal/obj/x86":                                    3,
	"cmd/internal/objabi":                                     3,
	"cmd/internal/objfile":                                    3,
	"cmd/internal/src":                                        3,
	"cmd/internal/sys":                                        3,
	"cmd/link":                                                3,
	"cmd/link/internal":                                       1,
	"cmd/link/internal/amd64":                                 3,
	"cmd/link/internal/arm":                                   3,
	"cmd/link/internal/arm64":                                 3,
	"cmd/link/internal/ld":                                    3,
	"cmd/link/internal/mips":                                  3,
	"cmd/link/internal/mips64":                                3,
	"cmd/link/internal/ppc64":                                 3,
	"cmd/link/internal/s390x":                                 3,
	"cmd/link/internal/x86":                                   3,
	"cmd/nm":                                                  3,
	"cmd/objdump":                                             3,
	"cmd/pack":                                                3,
	"cmd/pprof":                                               3,
	"cmd/trace":                                               3,
	"cmd/vendor":                                              1,
	"cmd/vendor/github.com":                                   1,
	"cmd/vendor/github.com/google":                            1,
	"cmd/vendor/github.com/google/pprof":                      1,
	"cmd/vendor/github.com/google/pprof/driver":               3,
	"cmd/vendor/github.com/google/pprof/internal":             1,
	"cmd/vendor/github.com/google/pprof/internal/binutils":    3,
	"cmd/vendor/github.com/google/pprof/internal/driver":      3,
	"cmd/vendor/github.com/google/pprof/internal/elfexec":     3,
	"cmd/vendor/github.com/google/pprof/internal/graph":       3,
	"cmd/vendor/github.com/google/pprof/internal/measurement": 3,
	"cmd/vendor/github.com/google/pprof/internal/plugin":      3,
	"cmd/vendor/github.com/google/pprof/internal/proftest":    3,
	"cmd/vendor/github.com/google/pprof/internal/report":      3,
	"cmd/vendor/github.com/google/pprof/internal/symbolizer":  3,
	"cmd/vendor/github.com/google/pprof/internal/symbolz":     3,
	"cmd/vendor/github.com/google/pprof/profile":              3,
	"cmd/vendor/github.com/google/pprof/third_party":          1,
	"cmd/vendor/github.com/google/pprof/third_party/svg":      3,
	"cmd/vendor/github.com/ianlancetaylor":                    1,
	"cmd/vendor/github.com/ianlancetaylor/demangle":           3,
	"cmd/vendor/golang.org":                                   1,
	"cmd/vendor/golang.org/x":                                 1,
	"cmd/vendor/golang.org/x/arch":                            1,
	"cmd/vendor/golang.org/x/arch/arm":                        1,
	"cmd/vendor/golang.org/x/arch/arm/armasm":                 3,
	"cmd/vendor/golang.org/x/arch/ppc64":                      1,
	"cmd/vendor/golang.org/x/arch/ppc64/ppc64asm":             3,
	"cmd/vendor/golang.org/x/arch/x86":                        1,
	"cmd/vendor/golang.org/x/arch/x86/x86asm":                 3,
	"cmd/vet":                           3,
	"cmd/vet/internal":                  1,
	"cmd/vet/internal/cfg":              3,
	"cmd/vet/internal/whitelist":        3,
	"compress":                          1,
	"compress/bzip2":                    3,
	"compress/flate":                    3,
	"compress/gzip":                     3,
	"compress/lzw":                      3,
	"compress/zlib":                     3,
	"container":                         1,
	"container/heap":                    3,
	"container/list":                    3,
	"container/ring":                    3,
	"context":                           3,
	"crypto":                            3,
	"crypto/aes":                        3,
	"crypto/cipher":                     3,
	"crypto/des":                        3,
	"crypto/dsa":                        3,
	"crypto/ecdsa":                      3,
	"crypto/elliptic":                   3,
	"crypto/hmac":                       3,
	"crypto/internal":                   1,
	"crypto/internal/cipherhw":          3,
	"crypto/md5":                        3,
	"crypto/rand":                       3,
	"crypto/rc4":                        3,
	"crypto/rsa":                        3,
	"crypto/sha1":                       3,
	"crypto/sha256":                     3,
	"crypto/sha512":                     3,
	"crypto/subtle":                     3,
	"crypto/tls":                        3,
	"crypto/x509":                       3,
	"crypto/x509/pkix":                  3,
	"database":                          1,
	"database/sql":                      3,
	"database/sql/driver":               3,
	"debug":                             1,
	"debug/dwarf":                       3,
	"debug/elf":                         3,
	"debug/gosym":                       3,
	"debug/macho":                       3,
	"debug/pe":                          3,
	"debug/plan9obj":                    3,
	"encoding":                          3,
	"encoding/ascii85":                  3,
	"encoding/asn1":                     3,
	"encoding/base32":                   3,
	"encoding/base64":                   3,
	"encoding/binary":                   3,
	"encoding/csv":                      3,
	"encoding/gob":                      3,
	"encoding/hex":                      3,
	"encoding/json":                     3,
	"encoding/pem":                      3,
	"encoding/xml":                      3,
	"errors":                            3,
	"expvar":                            3,
	"flag":                              3,
	"fmt":                               3,
	"go":                                1,
	"go/ast":                            3,
	"go/build":                          3,
	"go/constant":                       3,
	"go/doc":                            3,
	"go/format":                         3,
	"go/importer":                       3,
	"go/internal":                       1,
	"go/internal/gccgoimporter":         3,
	"go/internal/gcimporter":            3,
	"go/internal/srcimporter":           3,
	"go/parser":                         3,
	"go/printer":                        3,
	"go/scanner":                        3,
	"go/token":                          3,
	"go/types":                          3,
	"hash":                              3,
	"hash/adler32":                      3,
	"hash/crc32":                        3,
	"hash/crc64":                        3,
	"hash/fnv":                          3,
	"html":                              3,
	"html/template":                     3,
	"image":                             3,
	"image/color":                       3,
	"image/color/palette":               3,
	"image/draw":                        3,
	"image/gif":                         3,
	"image/internal":                    1,
	"image/internal/imageutil":          3,
	"image/jpeg":                        3,
	"image/png":                         3,
	"index":                             1,
	"index/suffixarray":                 3,
	"internal":                          1,
	"internal/cpu":                      3,
	"internal/nettrace":                 3,
	"internal/poll":                     3,
	"internal/race":                     3,
	"internal/singleflight":             3,
	"internal/syscall":                  1,
	"internal/syscall/unix":             3,
	"internal/syscall/windows":          3,
	"internal/syscall/windows/registry": 3,
	"internal/syscall/windows/sysdll":   3,
	"internal/testenv":                  3,
	"internal/trace":                    3,
	"io":                                3,
	"io/ioutil":                         3,
	"log":                               3,
	"log/syslog":                        3,
	"math":                              3,
	"math/big":                          3,
	"math/bits":                         3,
	"math/cmplx":                        3,
	"math/rand":                         3,
	"mime":                              3,
	"mime/multipart":                    3,
	"mime/quotedprintable":              3,
	"net":                                                           3,
	"net/http":                                                      3,
	"net/http/cgi":                                                  3,
	"net/http/cookiejar":                                            3,
	"net/http/fcgi":                                                 3,
	"net/http/httptest":                                             3,
	"net/http/httptrace":                                            3,
	"net/http/httputil":                                             3,
	"net/http/internal":                                             3,
	"net/http/pprof":                                                3,
	"net/internal":                                                  1,
	"net/internal/socktest":                                         3,
	"net/mail":                                                      3,
	"net/rpc":                                                       3,
	"net/rpc/jsonrpc":                                               3,
	"net/smtp":                                                      3,
	"net/textproto":                                                 3,
	"net/url":                                                       3,
	"os":                                                            3,
	"os/exec":                                                       3,
	"os/signal":                                                     3,
	"os/user":                                                       3,
	"path":                                                          3,
	"path/filepath":                                                 3,
	"plugin":                                                        3,
	"reflect":                                                       3,
	"regexp":                                                        3,
	"regexp/syntax":                                                 3,
	"runtime":                                                       3,
	"runtime/cgo":                                                   3,
	"runtime/debug":                                                 3,
	"runtime/internal":                                              1,
	"runtime/internal/atomic":                                       3,
	"runtime/internal/sys":                                          3,
	"runtime/pprof":                                                 3,
	"runtime/pprof/internal":                                        1,
	"runtime/pprof/internal/profile":                                3,
	"runtime/race":                                                  3,
	"runtime/trace":                                                 3,
	"sort":                                                          3,
	"strconv":                                                       3,
	"strings":                                                       3,
	"sync":                                                          3,
	"sync/atomic":                                                   3,
	"syscall":                                                       3,
	"testing":                                                       3,
	"testing/internal":                                              1,
	"testing/internal/testdeps":                                     3,
	"testing/iotest":                                                3,
	"testing/quick":                                                 3,
	"text":                                                          1,
	"text/scanner":                                                  3,
	"text/tabwriter":                                                3,
	"text/template":                                                 3,
	"text/template/parse":                                           3,
	"time":                                                          3,
	"unicode":                                                       3,
	"unicode/utf16":                                                 3,
	"unicode/utf8":                                                  3,
	"unsafe":                                                        3,
	"vendor":                                                        1,
	"vendor/golang_org":                                             1,
	"vendor/golang_org/x":                                           1,
	"vendor/golang_org/x/crypto":                                    1,
	"vendor/golang_org/x/crypto/chacha20poly1305":                   3,
	"vendor/golang_org/x/crypto/chacha20poly1305/internal":          1,
	"vendor/golang_org/x/crypto/chacha20poly1305/internal/chacha20": 3,
	"vendor/golang_org/x/crypto/curve25519":                         3,
	"vendor/golang_org/x/crypto/poly1305":                           3,
	"vendor/golang_org/x/net":                                       1,
	"vendor/golang_org/x/net/http2":                                 1,
	"vendor/golang_org/x/net/http2/hpack":                           3,
	"vendor/golang_org/x/net/idna":                                  3,
	"vendor/golang_org/x/net/lex":                                   1,
	"vendor/golang_org/x/net/lex/httplex":                           3,
	"vendor/golang_org/x/net/lif":                                   3,
	"vendor/golang_org/x/net/nettest":                               3,
	"vendor/golang_org/x/net/proxy":                                 3,
	"vendor/golang_org/x/net/route":                                 3,
	"vendor/golang_org/x/text":                                      1,
	"vendor/golang_org/x/text/secure":                               3,
	"vendor/golang_org/x/text/secure/bidirule":                      3,
	"vendor/golang_org/x/text/transform":                            3,
	"vendor/golang_org/x/text/unicode":                              3,
	"vendor/golang_org/x/text/unicode/bidi":                         3,
	"vendor/golang_org/x/text/unicode/norm":                         3,
}
