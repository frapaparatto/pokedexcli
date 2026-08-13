package repl

import (
	"errors"
	"fmt"
)

func commandMap(conf *Config, options ...string) error {
	locationsResp, err := conf.pokeapiClient.ListLocations(conf.nextLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = locationsResp.Next
	conf.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(conf *Config, options ...string) error {
	if conf.prevLocationsURL == nil {
		return errors.New("no previous page")
	}

	locationsResp, err := conf.pokeapiClient.ListLocations(conf.prevLocationsURL)
	if err != nil {
		return err
	}

	conf.nextLocationsURL = locationsResp.Next
	conf.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
