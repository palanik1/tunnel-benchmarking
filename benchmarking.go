package main

import "fmt"
import "os/exec"
import "os"
import "log"
import "strconv"


//TODO: Move to netlink
func setupTunnels(){
	num_tunnels := 5
	REMOTE_ADDR := "10.162.185.158"
	count := 100
	os.Setenv("PATH", "/usr/bin:/sbin")
	for idx := 0; idx < num_tunnels; idx++ {
		count++
		cmd := "ip link add name irlgeneve"+strconv.Itoa(idx)+ " type geneve id " + strconv.Itoa(count) + " remote " + REMOTE_ADDR
		fmt.Println("cmd: ",cmd)
		name := "irlgeneve" + strconv.Itoa(idx)
		execCmd := exec.Command("ip","link", "add","name",name,"type", "geneve","id",strconv.Itoa(count),"remote",REMOTE_ADDR)
		err := execCmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func runServer(){
	cmd := exec.Command("netserver")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
 
func runClient(server string, duration int){
	cmd := exec.Command("netperf ", "- H", server,"-l",strconv.Itoa(duration))
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main(){

	const num_tunnels int = 5
	REMOTE_ADDR := "10.162.185.158"
	count := 100
	os.Setenv("PATH", "/usr/bin:/sbin")
	for idx := 0; idx < num_tunnels; idx++ {
		count++
		//cmd := "ip link add name irlgeneve"+strconv.Itoa(idx)+ " type geneve id " + strconv.Itoa(count) + " remote " + REMOTE_ADDR
		//fmt.Println("cmd: ",cmd)
		name := "irlgeneve" + strconv.Itoa(idx)
		execCmd := exec.Command("ip","link", "add","name",name,"type", "geneve","id",strconv.Itoa(count),"remote",REMOTE_ADDR)
		err := execCmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		cmd := exec.Command("ip","link","set","dev",name,"up")
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

