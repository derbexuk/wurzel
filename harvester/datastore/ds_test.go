package datastore

import (
	. "github.com/onsi/gomega"
	"gopkg.in/redis.v5"

	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	Init(false)
	code := m.Run()
	os.Exit(code)
}

func TestSet(t *testing.T) {
	RegisterTestingT(t)

	err := Set("testing/t1", "ttt")
	Expect(err).Should(BeNil())
}

func TestGet(t *testing.T) {
	RegisterTestingT(t)

	err := Set("testing/t1", "ttt")
	Expect(err).Should(BeNil())

	val, err := Get("testing/arse")
	Expect(err.Error()).To(Equal("Path Not Found : testing/arse"))

	val, err = Get("testing/t1")
	Expect(err).Should(BeNil())
	Expect(val).To(Equal("ttt"))
}

func TestExp(t *testing.T) {
	Init(true)
	RegisterTestingT(t)

	err := SetExp("testing/exp", "ttt", 1)
	Expect(err).Should(BeNil())

	val, err := Get("testing/exp")
	Expect(err).Should(BeNil())
	Expect(val).To(Equal("ttt"))

	time.Sleep(2 * time.Second)
	val, err = redisClient.Get("testing/exp").Result()
	Expect(err).To(Equal(redis.Nil))

	val, err = Get("testing/exp")
	Expect(err).Should(BeNil())
	Expect(val).To(Equal("ttt"))

	err = Del("testing/exp")
	val, err = Get("testing/exp")
	Expect(err).ShouldNot(BeNil())
	Init(false)
}

func TestGetPaths(t *testing.T) {
	RegisterTestingT(t)

	vals, err := GetPaths("arse/*")
	Expect(err.Error()).To(Equal("Path Not Found : arse/*"))
	Expect(len(vals)).To(Equal(0))

	err = Set("testing/t1", "tt1")
	Expect(err).Should(BeNil())
	err = Set("testing/t2", "tt2")
	Expect(err).Should(BeNil())
	err = Set("testing/t3", "tt3")
	Expect(err).Should(BeNil())
	err = Set("testing/t4", "tt4")
	Expect(err).Should(BeNil())

	vals, err = GetPaths("testing/*")
	Expect(err).Should(BeNil())
	Expect(len(vals)).To(Equal(4))
	//Can't guarantee order
	//Expect(vals[2]).To(Equal("testing/t1"))
}

func TestDel(t *testing.T) {
	RegisterTestingT(t)

	err := Set("testing/t1", "tt1")
	Expect(err).Should(BeNil())

	err = Del("testing/t1")
	Expect(err).Should(BeNil())

	_, err = Get("testing/t1")
	Expect(err.Error()).To(Equal("Path Not Found : testing/t1"))
}

func TestAddToIndex(t *testing.T) {
	RegisterTestingT(t)

	//Clean out index
	Del("/index/members/testing/idx1")
	Del("/index/list/testing/idx1")
	r, err := InIndex("testing/idx1", "tt3")
	Expect(err).Should(BeNil())
	Expect(r).To(Equal(false))

	err = AddToSeriesIndex("testing/idx1", "tt1")
	Expect(err).Should(BeNil())
	r, err = InIndex("testing/idx1", "tt3")
	Expect(err).Should(BeNil())
	Expect(r).To(Equal(false))
	err = AddToSeriesIndex("testing/idx1", "tt3")
	Expect(err).Should(BeNil())
	r, err = InIndex("testing/idx1", "tt3")
	Expect(err).Should(BeNil())
	Expect(r).To(Equal(true))
	err = AddToSeriesIndex("testing/idx1", "tt2")
	Expect(err).Should(BeNil())
	err = AddToSeriesIndex("testing/idx1", "tt9")
	Expect(err).Should(BeNil())
	err = AddToSeriesIndex("testing/idx1", "tt8")
	Expect(err).Should(BeNil())
	err = AddToSeriesIndex("testing/idx1", "tt7")
	Expect(err).Should(BeNil())

	rge, err := GetSeriesRange("testing/idx1", 0, 3)
	Expect(err).Should(BeNil())
	Expect(len(rge)).To(Equal(4))
	Expect(rge[0]).To(Equal("tt7"))
	Expect(rge[3]).To(Equal("tt2"))
	rge, err = GetSeriesRange("testing/idx1", 4, 9)
	Expect(err).Should(BeNil())
	Expect(rge[0]).To(Equal("tt3"))
	Expect(len(rge)).To(Equal(2))
}

func TestMapHash(t *testing.T) {
	RegisterTestingT(t)

	hpath := "/hashtest"
	tdata := map[string]string{"a": "1", "b": "2", "c": "3"}
	//Clean out index
	Del(hpath)

	err := MapSet(hpath, tdata)
	Expect(err).Should(BeNil())

	r, err := MapElGet(hpath, "b")
	Expect(err).Should(BeNil())
	Expect(r).To(Equal("2"))

	r, err = MapElGet("duffpath", "d")
	Expect(r).To(Equal(""))
	Expect(err.Error()).To(Equal("Not Found : duffpath or d"))

	r, err = MapElGet(hpath, "d")
	Expect(r).To(Equal(""))
	Expect(err.Error()).To(Equal("Not Found : /hashtest or d"))

	err = MapElSet(hpath, "d", "4")
	Expect(err).Should(BeNil())

	err = MapElDel(hpath, "d")
	Expect(err).Should(BeNil())

	r, err = MapElGet(hpath, "d")
	Expect(r).To(Equal(""))
	Expect(err.Error()).To(Equal("Not Found : /hashtest or d"))
}

func TestPersist(t *testing.T) {
	RegisterTestingT(t)

	useDb := true
	Init(useDb)

	//Paths are unique
	err := Set("/persist/testing/p1", "tt1")
	err = Set("/persist/testing/p1", "tt2")
	Expect(err).Should(BeNil())
	_, err = redisClient.Del("/persist/testing/p1").Result()
	Expect(err).Should(BeNil())

	val, err := Get("/persist/testing/p1")
	Expect(err).Should(BeNil())
	Expect(val).To(Equal("tt2"))
	_, err = redisClient.Del("/persist/testing/p1").Result()

	err = Set("/persist/testing/p2", "tt2")
	_, err = redisClient.Del("/persist/testing/p2").Result()
	rs, err := GetPaths("/persist/testing/*")
	Expect(err).Should(BeNil())
	log.Println(rs)
	Expect(len(rs)).To(Equal(2))

	err = Del("/persist/testing/p0")
	Expect(err).Should(BeNil())

	err = Del("/persist/testing/p1")
	Expect(err).Should(BeNil())
	val, err = Get("/persist/testing/p1")
	Expect(err).ShouldNot(BeNil())
	Expect(val).To(Equal(""))
	dbh.DelPath(DB, COL, "/persist/testing/p2")
	err = Del("/persist/testing/p2")
	Expect(err).Should(BeNil())
}
