package preparer

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/square/p2/pkg/intent"
	"github.com/square/p2/pkg/logging"
	"github.com/square/p2/pkg/pods"
	"github.com/square/p2/pkg/reality"
)

type Pod interface {
	Launch(*pods.PodManifest) (bool, error)
	Install(*pods.PodManifest) error
	CurrentManifest() (*pods.PodManifest, error)
	Halt() (bool, error)
}

type Preparer struct {
	node   string
	iStore *intent.Store
	rStore *reality.Store
	hooks  *pods.HookDir
	Logger logging.Logger
}

func New(nodeName string, consulAddress string, hooksDirectory string, logger logging.Logger) (*Preparer, error) {
	iStore, err := intent.LookupStore(intent.Options{
		Token:   nodeName,
		Address: consulAddress,
	})
	if err != nil {
		return nil, err
	}

	rStore, err := reality.LookupStore(reality.Options{
		Token:   nodeName,
		Address: consulAddress,
	})
	if err != nil {
		return nil, err
	}

	return &Preparer{
		node:   nodeName,
		iStore: iStore,
		rStore: rStore,
		hooks:  pods.Hooks(hooksDirectory),
		Logger: logger,
	}, nil
}

func (p *Preparer) WatchForPodManifestsForNode() {
	pods.Log = p.Logger
	path := fmt.Sprintf("%s/%s", intent.INTENT_TREE, p.node)

	// This allows us to signal the goroutine watching consul to quit
	watcherQuit := make(<-chan struct{})
	errChan := make(chan error)
	podChan := make(chan pods.PodManifest)

	go p.iStore.WatchPods(path, watcherQuit, errChan, podChan)

	// we will have one long running goroutine for each app installed on this
	// host. We keep a map of podId => podChan so we can send the new manifests
	// that come in to the appropriate goroutine
	podChanMap := make(map[string]chan pods.PodManifest)
	quitChanMap := make(map[string]chan struct{})

	for {
		select {
		case err := <-errChan:
			p.Logger.WithFields(logrus.Fields{
				"inner_err": err,
			}).Errorln("there was an error reading the manifest")
		case manifest := <-podChan:
			podId := manifest.ID()
			if podChanMap[podId] == nil {
				// No goroutine is servicing this app currently, let's start one
				podChanMap[podId] = make(chan pods.PodManifest)
				quitChanMap[podId] = make(chan struct{})
				go p.handlePods(podChanMap[podId], quitChanMap[podId])
			}
			podChanMap[podId] <- manifest
		}
	}
}

// no return value, no output channels. This should do everything it needs to do
// without outside intervention (other than being signalled to quit)
func (p *Preparer) handlePods(podChan <-chan pods.PodManifest, quit <-chan struct{}) {
	// install new launchables
	var manifestToLaunch pods.PodManifest

	// used to track if we have work to do (i.e. pod manifest came through channel
	// and we have yet to operate on it)
	working := false
	var manifestLogger logging.Logger
	for {
		select {
		case <-quit:
			return
		case manifestToLaunch = <-podChan:
			sha, err := manifestToLaunch.SHA()
			manifestLogger = p.Logger.SubLogger(logrus.Fields{
				"manifest": manifestToLaunch.ID(),
				"sha":      sha,
				"sha_err":  err,
			})
			manifestLogger.NoFields().Infoln("New manifest received")
			working = true
		case <-time.After(1 * time.Second):
			if working {
				pod := pods.PodFromManifestId(manifestToLaunch.ID())
				err := p.hooks.RunBefore(pod, &manifestToLaunch)
				if err != nil {
					manifestLogger.WithFields(logrus.Fields{
						"err":   err,
						"hooks": "before",
					}).Warnln("Could not run before hooks")
				}
				ok := p.installAndLaunchPod(&manifestToLaunch, pod, manifestLogger)
				if ok {
					manifestToLaunch = pods.PodManifest{}
					working = false
				}
				err = p.hooks.RunAfter(pod, &manifestToLaunch)
				if err != nil {
					manifestLogger.WithFields(logrus.Fields{
						"err":   err,
						"hooks": "after",
					}).Warnln("Could not run after hooks")
				}
			}
		}
	}
}

func (p *Preparer) installAndLaunchPod(newManifest *pods.PodManifest, pod Pod, logger logging.Logger) bool {
	// do not remove the logger argument, it's not the same as p.Logger

	err := pod.Install(newManifest)
	if err != nil {
		// install failed, abort and retry
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("Install failed")
		return false
	}

	// get currently running pod to compare with the new pod
	currentManifest, err := pod.CurrentManifest()
	currentSHA, _ := currentManifest.SHA()
	newSHA, _ := newManifest.SHA()

	// if new or the manifest is different, launch
	newOrDifferent := err == pods.NoCurrentManifest || currentSHA != newSHA
	if newOrDifferent {
		logger.WithFields(logrus.Fields{
			"old_sha":  currentSHA,
			"sha":      newSHA,
			"manifest": newManifest.ID(),
		}).Infoln("SHA is new or different from old, will update")
	}

	// if the old manifest is corrupted somehow, re-launch since we don't know if this is an update.
	problemReadingCurrentManifest := (err != nil && err != pods.NoCurrentManifest)
	if problemReadingCurrentManifest {
		logger.WithFields(logrus.Fields{
			"sha":       newSHA,
			"inner_err": err,
		}).Errorln("Current manifest not readable, will relaunch")
	}

	if newOrDifferent || problemReadingCurrentManifest {
		err := p.iStore.RegisterPodService(*newManifest)
		if err != nil {
			logger.WithField("err", err).Errorln("Service registration failed")
			return false
		}
		ok, err := pod.Launch(newManifest)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("Launch failed")
		} else {
			p.rStore.SetPod(p.node, *newManifest)
		}
		return err == nil && ok
	}

	// TODO: shut down removed launchables between pod versions.
	return true
}
