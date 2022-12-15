package main

import (
    "io/ioutil"
    "log"
    "os/exec"
    "time"
)

const (
    dir = "/path/to/directory"
    dest = "user@host:/path/to/destination"
)

func main() {
    for {
        files, err := ioutil.ReadDir(dir)
        if err != nil {
            log.Fatal(err)
        }

        for _, file := range files {
            if !file.IsDir() {
                // Use the `scp` command to copy the file to the destination
                cmd := exec.Command("scp", file.Name(), dest)
                err = cmd.Run()
                if err != nil {
                    log.Printf("Error copying file: %v", err)
                    continue
                }
            }
        }

        time.Sleep(1 * time.Second)
    }
}
//config for scp can be formatted like this:

//cmd := exec.Command("scp", dir+"/*", "user@foo.bar:/path/to/destination")

//that arguement will grab every file in teh directory and scp it to a server called foo.bar (assuming the user has a public key on the server)
// This exec.Command call constructs the scp command with the following arguments:

//"scp": the name of the scp command
//dir+"/*": the path to the directory that contains the files to be copied, followed by /* to include all files in the directory
//"user@foo.bar:/path/to/destination": the destination for the copied files,
// in the form user@host:/path/to/destination where user is the username on the destination machine,
// host is the hostname or IP address of the destination machine, 
//and /path/to/destination is the path to the directory where the files should be copied

//if you need to specify a password
//cmd := exec.Command("scp", "-p", "PASSWORD", dir+"/*", "user@foo.bar:/path/to/destination")
//this application essentially performs
//scp -p PASSWORD /path/to/directory/* user@foo.bar:/path/to/destination
//in this configuration
