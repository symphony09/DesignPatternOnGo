package dp09_FacadePattern

import "fmt"

func (cpu *CPU) freeze() {
	fmt.Println("CPU freeze ...")
}

func (cpu *CPU) jump() {
	fmt.Println("CPU jump ...")
}

func (cpu *CPU) execute() {
	fmt.Println("CPU execute ...")
}

func (m *Memory) load(data interface{}) {
	fmt.Println("Memory load ", data)
}

func (hd *HardDrive) read() interface{} {
	fmt.Println("HardDrive read ...")
	return "data"
}

func (c *Computer) Start() {
	cpu := CPU{}
	cpu.freeze()
	drive := HardDrive{}
	memory := Memory{}
	memory.load(drive.read())
	cpu.jump()
	cpu.execute()
}
