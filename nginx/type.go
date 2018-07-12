package nginx

type Nginx struct {
	InstallPath string
	Version string
	Params struct{
		WorkerProcess int64
		WorkerConnections int64



	}
}

const NGINX_OFFICIAL_SITE  = "https://nginx.org/en/download.html"



