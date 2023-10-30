//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)

func printServerStatus(servers map[string]int) {
	numServersByStatus := map[int]int{
		Online:      0,
		Offline:     0,
		Maintenance: 0,
		Retired:     0,
	}

	for _, server := range servers {
		numServersByStatus[server]++
	}

	fmt.Println("Number of servers:", len(servers))
	fmt.Println("Number of servers for each status:", numServersByStatus)
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatuses := make(map[string]int)
	for _, server := range servers {
		serverStatuses[server] = Online
		fmt.Println("serverstatuses:", serverStatuses)
	}

	//* After creating the map, perform the following actions:
	//  - call display server info function
	//  - change server status of `darkstar` to `Retired`
	//  - change server status of `aiur` to `Offline`
	//  - call display server info function
	//  - change server status of all servers to `Maintenance`
	//  - call display server info function
	printServerStatus(serverStatuses)

	serverStatuses["darkstar"] = Retired
	serverStatuses["aiur"] = Offline
	printServerStatus(serverStatuses)

	for server := range serverStatuses {
		serverStatuses[server] = Maintenance
	}
	printServerStatus(serverStatuses)

}
