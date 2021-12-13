package ch2

import "image"

type Drone struct {
}

type Vec3 struct {
}

func (d *Drone) NagivateTo(dst Vec3) {

}

func (d *Drone) CaptureImage() (*image.RGBA, error) {
	return nil, nil
}

//
type Target struct {
}

type MobileNet struct{
}

func (d *MobileNet) DetectTargets() ([]*Target, error) {
	return nil, nil
}
