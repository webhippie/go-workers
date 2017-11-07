package workers

import (
	"errors"
	"os"
	"sync"

	"github.com/go-kit/kit/log"
)

const (
	RETRY_KEY          = "goretry"
	SCHEDULED_JOBS_KEY = "schedule"
)

var (
	Logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))

	Middleware = NewMiddleware(
		&MiddlewareLogging{},
		&MiddlewareRetry{},
		&MiddlewareStats{},
	)

	managers = make(map[string]*manager)
	control  = make(map[string]chan string)

	schedule *scheduled
	access   sync.Mutex
	started  bool
)

func Process(queue string, job JobFunc, concurrency int, mids ...Action) {
	access.Lock()
	defer access.Unlock()

	managers[queue] = newManager(queue, job, concurrency, mids...)
}

func Run() {
	Start()
	Wait()
}

func Wait() {
	for _, manager := range managers {
		manager.Wait()
	}
}

func Start() {
	access.Lock()
	defer access.Unlock()

	if started {
		return
	}

	runHooks(beforeStart)
	startSchedule()
	startManagers()
	runHooks(afterStart)

	started = true
}

func Quit() {
	access.Lock()
	defer access.Unlock()

	if !started {
		return
	}

	runHooks(beforeQuit)
	quitManagers()
	quitSchedule()
	runHooks(afterQuit)

	Wait()

	started = false
}

func Reset() error {
	access.Lock()
	defer access.Unlock()

	if started {
		return errors.New("Cannot reset worker managers while workers are running")
	}

	managers = make(map[string]*manager)

	return nil
}

func startSchedule() {
	if schedule == nil {
		schedule = newScheduled(RETRY_KEY, SCHEDULED_JOBS_KEY)
	}

	schedule.start()
}

func quitSchedule() {
	if schedule != nil {
		schedule.quit()
		schedule = nil
	}
}

func startManagers() {
	for _, manager := range managers {
		manager.start()
	}
}

func quitManagers() {
	for _, m := range managers {
		go (func(m *manager) { m.quit() })(m)
	}
}
