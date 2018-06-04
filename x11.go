package x11

import (
	"github.com/samli88/go-x11-hash/blake"
	"github.com/samli88/go-x11-hash/bmw"
	"github.com/samli88/go-x11-hash/cubehash"
	"github.com/samli88/go-x11-hash/echo"
	"github.com/samli88/go-x11-hash/groestl"
	"github.com/samli88/go-x11-hash/hash"
	"github.com/samli88/go-x11-hash/jh"
	"github.com/samli88/go-x11-hash/keccak"
	"github.com/samli88/go-x11-hash/luffa"
	"github.com/samli88/go-x11-hash/shavite"
	"github.com/samli88/go-x11-hash/simd"
	"github.com/samli88/go-x11-hash/skein"
)

// Hash contains the state objects
// required to perform the x11.Hash.
type Hash struct {
	tha [64]byte
	thb [64]byte

	blake    hash.Digest
	bmw      hash.Digest
	cubehash hash.Digest
	echo     hash.Digest
	groestl  hash.Digest
	jh       hash.Digest
	keccak   hash.Digest
	luffa    hash.Digest
	shavite  hash.Digest
	simd     hash.Digest
	skein    hash.Digest
}

// New returns a new object to compute a x11 hash.
func New() *Hash {
	ref := &Hash{}

	ref.blake = blake.New()
	ref.bmw = bmw.New()
	ref.cubehash = cubehash.New()
	ref.echo = echo.New()
	ref.groestl = groestl.New()
	ref.jh = jh.New()
	ref.keccak = keccak.New()
	ref.luffa = luffa.New()
	ref.shavite = shavite.New()
	ref.simd = simd.New()
	ref.skein = skein.New()

	return ref
}

// Hash computes the hash from the src bytes and stores the result in dst.
func (ref *Hash) Hash(src []byte, dst []byte) {
	ta := ref.tha[:]
	tb := ref.thb[:]

	ref.blake.Write(src)
	ref.blake.Close(tb, 0, 0)

	ref.bmw.Write(tb)
	ref.bmw.Close(ta, 0, 0)

	ref.groestl.Write(ta)
	ref.groestl.Close(tb, 0, 0)

	ref.skein.Write(tb)
	ref.skein.Close(ta, 0, 0)

	ref.jh.Write(ta)
	ref.jh.Close(tb, 0, 0)

	ref.keccak.Write(tb)
	ref.keccak.Close(ta, 0, 0)

	ref.luffa.Write(ta)
	ref.luffa.Close(tb, 0, 0)

	ref.cubehash.Write(tb)
	ref.cubehash.Close(ta, 0, 0)

	ref.shavite.Write(ta)
	ref.shavite.Close(tb, 0, 0)

	ref.simd.Write(tb)
	ref.simd.Close(ta, 0, 0)

	ref.echo.Write(ta)
	ref.echo.Close(tb, 0, 0)

	copy(dst, tb)
}
