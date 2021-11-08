//Many of tests use channel state rather than testing that a goroutine is actually running or not.
//I think this is reasonable for a first pass at least.
package scheduler

import (
	"io"
	"log"
	"os"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func copyFile(src, tgt string) {

	srcFile, err := os.Open(src)
	if err != nil {
		log.Printf("Error : %s\n", err.Error())
		os.Exit(1)
	}
	defer srcFile.Close()

	destFile, err := os.Create(tgt) // creates if file doesn't exist
	if err != nil {
		log.Printf("Error : %s\n", err.Error())
		os.Exit(1)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	if err != nil {
		log.Printf("Error : %s\n", err.Error())
		os.Exit(1)
	}

	err = destFile.Sync()
	if err != nil {
		log.Printf("Error : %s\n", err.Error())
		os.Exit(1)
	}
}

func TestDefaultSetup(t *testing.T) {
	RegisterTestingT(t)

	s := Scheduler{}
	s.Setup()
	Expect(s.feedDir).To(Equal("/etc/cspace/feeds.d/"))
	Expect(s.freq).To(Equal(30 * time.Second))
}

func TestEnvSetup(t *testing.T) {
	RegisterTestingT(t)

	os.Setenv("CSPACE_FEED_DIR", "/tmp/feeds.d/")
	os.Setenv("CSPACE_FEED_FREQ", "5s")

	s := Scheduler{}
	s.Setup()
	Expect(s.feedDir).To(Equal("/tmp/feeds.d/"))
	Expect(s.freq).To(Equal(5 * time.Second))
}

func TestRunProcess(t *testing.T) {
	RegisterTestingT(t)

	os.Setenv("CSPACE_FEED_DIR", "../feeds.d/")
	s := Scheduler{}
	s.Setup()

	s.runFile("london_air_quality.yaml")
	Expect(s.processChans).Should(HaveKey("london_air_quality.yaml"))
	Expect(s.processChans["london_air_quality.yaml"]).ShouldNot(BeClosed())
}

func TestKillProcess(t *testing.T) {
	RegisterTestingT(t)

	os.Setenv("CSPACE_FEED_DIR", "../feeds.d/")
	s := Scheduler{}
	s.Setup()

	//Can't kill it if it's not running
	s.runFile("london_air_quality.yaml")

	s.killFile("london_air_quality.yaml")
	Expect(s.processChans).ShouldNot(HaveKey("london_air_quality.yaml"))
}

func TestMissingFeedDir(t *testing.T) {
	os.Setenv("CSPACE_FEED_DIR", "/tmp/not_there/")
	s := Scheduler{}
	s.Setup()

	Expect(func() {
		s.Loop()
	}).Should(Panic())
}

func TestLoop(t *testing.T) {
	RegisterTestingT(t)

	os.Setenv("CSPACE_FEED_DIR", "../feeds.d/")
	os.Setenv("CSPACE_FEED_FREQ", "100ms")
	s := Scheduler{}
	s.Setup()

	go s.Loop()

	time.Sleep(1 * time.Second)
	Expect(s.processChans).Should(HaveLen(2))
	Expect(s.processChans).Should(HaveKey("london_air_quality.yaml"))

	copyFile("../feeds.d/london_air_quality.yaml", "../feeds.d/copy.yaml")
	time.Sleep(1 * time.Second)
	Expect(s.processChans).Should(HaveLen(3))
	Expect(s.processChans).Should(HaveKey("copy.yaml"))
	Expect(s.processChans["copy.yaml"]).ShouldNot(BeClosed())

	os.Remove("../feeds.d/copy.yaml")
	time.Sleep(1 * time.Second)
	Expect(s.processChans).Should(HaveLen(2))
	Expect(s.processChans).ShouldNot(HaveKey("copy.yaml"))
}

func TestConfigUpdate(t *testing.T) {
	RegisterTestingT(t)

	os.Setenv("CSPACE_FEED_DIR", "../feeds.d/")
	os.Setenv("CSPACE_FEED_FREQ", "100ms")
	s := Scheduler{}
	s.Setup()

	go s.Loop()

	time.Sleep(1 * time.Second)
	Expect(s.processChans).Should(HaveLen(2))
	Expect(s.processChans).Should(HaveKey("london_air_quality.yaml"))

	channel := s.processChans["london_air_quality.yaml"]

	if err := os.Chtimes("../feeds.d/london_air_quality.yaml", time.Now(), time.Now()); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
	Expect(s.processChans["london_air_quality.yaml"]).ToNot(Equal(channel))
	Expect(s.processChans["london_air_quality.yaml"]).ShouldNot(BeClosed())
}
