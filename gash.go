package main

import (
	"context"
	"fmt"
	"github.com/vmware/govmomi/examples"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"log"
	"os"
)

func runsh(command string, args []string) (s string, err error) {
	platform := []rune(command)
	switch string(platform[0:3]) {
	case "esx":
		fmt.Println("vmware")
		ctx := context.Background()

		// Connect and login to ESX or vCenter
		c, err := examples.NewClient(ctx)
		if err != nil {
			log.Fatal(err)
		}

		defer c.Logout(ctx)

		if command == "esxls" {
			m := view.NewManager(c.Client)

			v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
			if err != nil {
				log.Fatal(err)
			}

			defer v.Destroy(ctx)

			// Retrieve summary property for all machines
			// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.VirtualMachine.html
			var vms []mo.VirtualMachine
			err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vms)
			if err != nil {
				log.Fatal(err)
			}

			// Print summary per vm (see also: govc/vm/info.go)

			for _, vm := range vms {
				fmt.Printf("%s: %s\n", vm.Summary.Config.Name, vm.Summary.Config.GuestFullName)
			}
		}
	case "lin":
		fmt.Println("linode")
	case "rac":
		fmt.Println("rackspace")
	case "hel":
		//s := fmt.Sprintf("a %s", "string")
		//    fmt.Println(s)
		s := `All commands begin with a three letter prefix:
		esx - For vmware
		lin - For lindoe
		rac - for Rackspace
		qui or exi - for quit
		hel - for help`
		fmt.Println(s)
	case "exi":
		os.Exit(0)
	case "qui":
		os.Exit(0)
	default:
		fmt.Println("no")
	}
	return command, nil
}
func main() {
	var input string
	var args []string
	for {
		fmt.Scanln(&input)
		//fmt.Println(input)
		runsh(input, args)
	}
}
