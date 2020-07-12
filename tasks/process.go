package tasks

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

// handleEmailDeliveryTask 处理发送邮件任务
func handleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	userID, err := t.Payload.GetInt("user_id")
	if err != nil {
		return err
	}
	tmplID, err := t.Payload.GetString("template_id")
	if err != nil {
		return err
	}
	fmt.Printf("Send Email to User: user_id = %d, template_id = %s\n", userID, tmplID)
	// Email delivery logic ...
	return nil
}

// imageProcessor 图片处理任务结构体
type imageProcessor struct {
	// ... fields for struct
}

// ProcessTask 实现图片处理接口
func (p *imageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	src, err := t.Payload.GetString("src")
	if err != nil {
		return err
	}
	dst, err := t.Payload.GetString("dst")
	if err != nil {
		return err
	}
	fmt.Printf("Process image: src = %s, dst = %s\n", src, dst)
	// Image processing logic ...
	return nil
}

// newImageProcessor 创建图片处理实例
func newImageProcessor() *imageProcessor {
	// ... return an instance
	return &imageProcessor{}
}
