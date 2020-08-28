package handlers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/baking-bad/bcdhub/internal/database"
	"github.com/baking-bad/bcdhub/internal/logger"
	"github.com/baking-bad/bcdhub/internal/mq"
	"github.com/baking-bad/bcdhub/internal/verifier/compilation"
	"github.com/baking-bad/bcdhub/internal/verifier/filesgenerator"
	"github.com/gin-gonic/gin"
)

// VerifyContract -
func (ctx *Context) VerifyContract(c *gin.Context) {
	userID := CurrentUserID(c)
	if userID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user"})
		return
	}

	var req verificationRequest
	if err := c.ShouldBindJSON(&req); handleError(c, err, http.StatusBadRequest) {
		return
	}

	user, err := ctx.DB.GetUser(userID)
	if handleError(c, err, 0) {
		return
	}

	task := database.CompilationTask{
		UserID:  user.ID,
		Address: req.Address,
		Network: req.Network,
		Kind:    compilation.KindVerification,
		Status:  compilation.StatusPending,
	}

	err = ctx.DB.CreateCompilationTask(&task)
	if handleError(c, err, 0) {
		return
	}

	go ctx.runVerification(task.ID, req.SourceURL)

	c.JSON(http.StatusOK, gin.H{"status": compilation.StatusPending})
}

func (ctx *Context) runVerification(taskID uint, sourceURL string) {
	dir := filepath.Join(ctx.SharePath, "/compilations")

	tempDir, err := ioutil.TempDir(dir, "verification")
	if err != nil {
		return
	}

	files, err := filesgenerator.FromRepo(sourceURL, tempDir)
	if ctx.handleCompilationError(taskID, err) {
		return
	}

	data := compilation.Task{
		ID:    taskID,
		Kind:  compilation.KindVerification,
		Files: files,
		Dir:   tempDir,
	}

	err = ctx.MQ.Send(mq.ChannelNew, mq.QueueCompilations, data)
	if ctx.handleCompilationError(taskID, err) {
		return
	}
}

func (ctx *Context) handleCompilationError(taskID uint, err error) bool {
	if err == nil {
		return false
	}

	logger.Error(err)

	if err := ctx.DB.UpdateTaskStatus(taskID, compilation.StatusError); err != nil {
		logger.Error(err)
	}

	return true
}
