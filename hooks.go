package workers

var (
	beforeStart []HookFunc
	afterStart  []HookFunc
	beforeQuit  []HookFunc
	afterQuit   []HookFunc
)

type HookFunc func()

// BeforeStart appends a hook before start.
func BeforeStart(f HookFunc) {
	access.Lock()
	defer access.Unlock()
	beforeStart = append(beforeStart, f)
}

// AfterStart appends a hook after start.
func AfterStart(f HookFunc) {
	access.Lock()
	defer access.Unlock()
	afterStart = append(afterStart, f)
}

// BeforeQuit appends a hook before quit.
func BeforeQuit(f HookFunc) {
	access.Lock()
	defer access.Unlock()
	beforeQuit = append(beforeQuit, f)
}

// AfterQuit appends a hook after quit.
func AfterQuit(f HookFunc) {
	access.Lock()
	defer access.Unlock()
	afterQuit = append(afterQuit, f)
}

// runHooks gets called internally to run the hooks.
func runHooks(hooks []HookFunc) {
	for _, f := range hooks {
		f()
	}
}
