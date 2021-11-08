//Repeatedly fetch, parse and store the feeds defined in $CSPACE_FEED_DIR
package main

import (
	"github.com/derbexuk/wurzel/harvester/feeds/scheduler"
)

func main() {
  var s scheduler.Scheduler
  s.Setup()
  s.Loop()
}
