package common

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
)

func PrometheusBoot(port int) {
	// 上报数据的http接口
	http.Handle("/metrics", promhttp.Handler())
	// 启动web 服务
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
		if err != nil {
			log.Fatal("启动失败")
		}
		log.Info("监控启动，端口为：" + strconv.Itoa(port))
	}()
}
