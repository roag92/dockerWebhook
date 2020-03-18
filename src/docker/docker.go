package docker

type Docker interface {
	Start()
}

type docker struct {
	runner   Runner
}

func NewDocker(runner Runner) Docker {
	return docker{runner}
}

func (d docker) pull() {
	d.runner.run("docker-compose", "pull")
}

func (d docker) up() {
	args := []string{"up", "-d"}

	d.runner.run("docker-compose", args...)
}

func (d docker) stop() {
	d.runner.run("docker-compose", "stop")
}

func (d docker) Start()  {
	d.stop()
	d.pull()
	d.up()
}
