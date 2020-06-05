package dp08_ProxyPattern

import "fmt"

func (i *RealImage) DisplayImage() {
	i.loadImageFromDisk()
	fmt.Println("displaying:" + i.filename)
}

func (i *RealImage) loadImageFromDisk() {
	fmt.Println("loading:" + i.filename)
}

func (pi *ProxyImage) DisplayImage() {
	if pi.Image == nil {
		pi.Image = &RealImage{filename: pi.Filename}
	}
	pi.Image.DisplayImage()
}
