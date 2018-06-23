package dvbscan

import (
	"github.com/ziutek/dvb/linuxdvb/frontend"
	"fmt"
)

type channel struct {
	name string
}



func Scan(device int) (channels []channel, err error) {

	d, err := frontend.OpenRO("/dev/dvb/adapter" + string(device) + "/frontend0")

	if err != nil {
		fmt.Errorf("Could not open device", err)	
		return nil, err
	}

	err2 := d.Tune()
	if err2 != nil {
		fmt.Errorf("Failed to tune device", err2)
		return nil, err2
	}

	return nil, nil
}
