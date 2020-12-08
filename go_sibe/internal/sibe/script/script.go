package script

type Script struct {
	Name string
	Id   int
}

var ClientScripts = []Script{
	{
		Name: "All",
		Id:   0,
	},
	{
		Name: "SIBE_COMMONS",
		Id:   1,
	},
	{
		Name: "mcb",
		Id:   16,
	},
	{
		Name: "sisben",
		Id:   21,
	},
	{
		Name: "sibe hiscre",
		Id:   23,
	},
	{
		Name: "tcb",
		Id:   26,
	},
	{
		Name: "negocio",
		Id:   27,
	},
	{
		Name: "broker",
		Id:   28,
	},
	{
		Name: "PortalSibe",
		Id:   33,
	},
}

var DeployScripts = []Script{
	{
		Name: "All",
		Id:   0,
	},
	{
		Name: "LIBS-SIBE",
		Id:   1,
	}, {
		Name: "sibe cache",
		Id:   2,
	},
	{
		Name: "atuben",
		Id:   4,
	},
	{
		Name: "camvri",
		Id:   5,
	},
	{
		Name: "comsub",
		Id:   7,
	},
	{
		Name: "MCB",
		Id:   16,
	},
	{
		Name: "migra",
		Id:   18,
	},
	{
		Name: "reavdir",
		Id:   20,
	},
	{
		Name: "sisben",
		Id:   21,
	},
	{
		Name: "hiscre",
		Id:   23,
	},
	{
		Name: "tcb",
		Id:   26,
	},
	{
		Name: "negocio",
		Id:   27,
	},
	{
		Name: "broker",
		Id:   28,
	},
	{
		Name: "conev processamento",
		Id:   29,
	},
	{
		Name: "PortalSibe",
		Id:   33,
	},
	{
		Name: "PortalPericia",
		Id:   34,
	},
	{
		Name: "PortalSibeExploded",
		Id:   333,
	},
	{
		Name: "webservices",
		Id:   40,
	},
}
