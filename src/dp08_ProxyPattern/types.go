package dp08_ProxyPattern

type Image interface {
	DisplayImage()
}

type RealImage struct {
	filename string
}

type ProxyImage struct {
	Filename string
	Image    Image
}
