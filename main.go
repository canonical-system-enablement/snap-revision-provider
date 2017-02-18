/*
 * Copyright (C) 2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"text/tabwriter"
)

var (
	channel  = flag.String("channel", "candidate", "Channel from which the snap version is queried")
	snapName = flag.String("snap", "bluez", "Snap to query")
)

var architectures = []string{"armhf", "arm64", "i386", "amd64"}

// Shows the results in a rmadison fashion, that is:
// snap name | channel | revision | archiitecture
func show(storeResponses []StoreResponse) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, sd := range storeResponses {
		fmt.Fprintf(w, " %s\t| %s\t| %d\t| %s\n", *snapName, *channel, sd.Revision, sd.Architecture[0])
	}
	w.Flush()
}

func main() {
	flag.Parse()

	url := fmt.Sprintf("https://search.apps.ubuntu.com/api/v1/snaps/details/%s?channel=%s", *snapName, *channel)

	jsonResponses := make(chan io.ReadCloser)
	var storeResponses []StoreResponse

	var wg sync.WaitGroup
	wg.Add(len(architectures) * 2)

	// Go and fetch for each architecture asynchronuously
	for _, arch := range architectures {
		go func(arch string, url string) {
			defer wg.Done()

			// Create HTTP request & Client
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatal("NewRequest: ", err)
				return
			}
			req.Header.Set("X-Ubuntu-Series", "16")
			client := &http.Client{}
			req.Header.Set("X-Ubuntu-Architecture", arch)

			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			} else {
				jsonResponses <- resp.Body
			}
		}(arch, url)
	}

	// Collect and process the results
	go func() {
		for response := range jsonResponses {
			defer response.Close()
			var tmp StoreResponse
			err := json.NewDecoder(response).Decode(&tmp)
			if err != nil {
				log.Fatal(err)
			} else {
				storeResponses = append(storeResponses, tmp)
			}
			wg.Done()
		}
	}()

	wg.Wait()

	// Report back
	show(storeResponses)
}
