package server

import (
	"awesomeProject/setting"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func Run(cfg *setting.Setting) {

	// Загрузка HTML файлов
	router.LoadHTMLGlob("template/*")

	// Усказать путь к статическим ресурсам
	router.Static("d", "assets/")

	// Подготовить запросы
	prepare()

	// Запустить "слушатель"
	_ = router.Run(cfg.Address + ":" + cfg.Port)

}
