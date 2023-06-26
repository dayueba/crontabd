package master

import (
	"github.com/dayueba/crontabd/internal"
	"github.com/dayueba/crontabd/internal/master/config"
	"github.com/dayueba/crontabd/internal/master/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ApiServer struct {
	jobService    *service.JobService
	workerService *service.WorkerService
	logService    *service.LogService
	httpServer    *gin.Engine
	conf          *config.Config
}

type DeleteJobReq struct {
	Name string
}

func NewApiServer(conf *config.Config, jobService *service.JobService, workerService *service.WorkerService, logService *service.LogService) *ApiServer {
	r := gin.Default()

	server := ApiServer{
		jobService:    jobService,
		workerService: workerService,
		logService:    logService,
		httpServer:    r,
		conf:          conf,
	}

	r.POST("/jobs", func(c *gin.Context) {
		job := internal.Job{}
		err := c.BindJSON(&job)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		// 4, 保存到etcd
		if oldJob, err := jobService.SaveJob(&job); err != nil {
			c.AbortWithError(400, err)
		} else {
			c.JSON(200, oldJob)
		}
	})
	r.DELETE("/jobs", func(c *gin.Context) {
		var (
			oldJob *internal.Job
			req    = DeleteJobReq{}
			err    error
		)
		c.BindJSON(&req)

		// 去删除任务
		if oldJob, err = jobService.DeleteJob(req.Name); err != nil {
			c.AbortWithError(500, err)
		}

		c.JSON(200, oldJob)
	})
	r.GET("/jobs", func(c *gin.Context) {
		var (
			jobList []*internal.Job
			err     error
		)

		// 获取任务列表
		if jobList, err = jobService.ListJobs(); err != nil {
			c.AbortWithError(200, err)
		}
		c.JSON(200, jobList)
	})
	r.POST("/jobs/kill", func(c *gin.Context) {
		var (
			err error
		)
		name := c.PostForm("name")
		if name == "" {
			c.Abort()
		}

		// 去删除任务
		if err = jobService.KillJob(name); err != nil {
			c.AbortWithError(500, err)
		}

		c.JSON(200, nil)
	})
	r.GET("/workers", func(c *gin.Context) {
		var (
			workerList []string
			err        error
		)

		if workerList, err = workerService.ListWorkers(); err != nil {
			c.AbortWithError(200, err)
		}

		c.JSON(200, workerList)
	})
	r.GET("/jobs/logs", func(c *gin.Context) {
		// 获取请求参数 /job/log?name=job10&skip=0&limit=10
		name := c.Query("name")
		limitParam := c.Query("limit")
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			limit = 20
		}

		logList, err := logService.ListLog(name, 0, int64(limit))
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(200, logList)
	})

	return &server
}

func (server *ApiServer) Run() error {
	return server.httpServer.Run(":" + strconv.Itoa(server.conf.ApiPort))
}
