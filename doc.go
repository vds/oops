/*

Overview

Package oops, inspired by the OOPS tool of launchpad.net (https://launchpad.net/python-oops), is simpler, less bloated, go native implementation of oops reports.
An OOPS is an error report that a developer or a system administrator can access easily and quickly.
Oopses are stored locally as files. Once recorded they can be moved to a central service that indexes them and provide a convenient interface to search and browse them, refer to the oopsio package for a simple implementation of such service.

Example

	err := errors.New("this is an error")
	tempFolder, err := ioutil.TempDir("/tmp", "oops")
	p := oops.DiskPublisher{tempFolder}
	of := oops.Factory{p}
	requestData := map[string]string{}
	id := of.New(err, requestData)
	// Log that an Oops was recorded, together with the Oops id
	log.Printf("Oops! Oops id: %s\n", id)

*/

package oops
