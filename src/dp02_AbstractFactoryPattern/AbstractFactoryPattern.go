package dp02_AbstractFactoryPattern

import "fmt"

func (phone *P40) HCall() string {
	fmt.Println("Huawei Phone: P40 calling!")
	return "Huawei Phone call"
}

func (phone *Mi10) MCall() string {
	fmt.Println("Mi Phone: Mi10 calling!")
	return "Mi Phone call"
}

func (laptop *MagicBook) HPlay() string {
	fmt.Println("Huawei Laptop: MagicBook playing!")
	return "Huawei Laptop play"
}

func (laptop *RedMiBook) MPlay() string {
	fmt.Println("Mi Laptop: RedMiBook playing!")
	return "Mi Laptop play"
}

func (factory *HuaweiFactory) ProvidePhone() HuaweiPhone {
	return new(P40)
}

func (factory *HuaweiFactory) ProvideLaptop() HuaweiLaptop {
	return new(MagicBook)
}

func (factory *MiFactory) ProvidePhone() MiPhone {
	return new(Mi10)
}

func (factory *MiFactory) ProvideLaptop() MiLaptop {
	return new(RedMiBook)
}
