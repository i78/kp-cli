package runners

// Runner contains the interface contract for Runners which thenselves represent operation modes of the Koalapaste CLI.
type Runner interface {
	run()
}
