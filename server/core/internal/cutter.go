package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Cutter 实现 io.Writer 接口
// 用于日志切割, strings.Join([]string{director,layout, formats..., level+".log"}, os.PathSeparator)
type Cutter struct {
	level        string
	layout       string
	formats      []string
	director     string
	retentionDay int
	file         *os.File
	mutex        *sync.RWMutex
}

type CutterOption func(*Cutter)

// CutterWithLayout 时间格式
func CutterWithLayout(layout string) CutterOption {
	return func(c *Cutter) {
		c.layout = layout
	}
}

func CutterWithFormats(formats []string) CutterOption {
	return func(c *Cutter) {
		c.formats = formats
	}
}

func NewCutter(director string, level string, retentionDay int, options ...CutterOption) *Cutter {
	rotate := &Cutter{
		level:        level,
		director:     director,
		retentionDay: retentionDay,
		mutex:        &sync.RWMutex{},
	}
	for i := 0; i < len(options); i++ {
		options[i](rotate)
	}
	fmt.Printf("options: %v, %T\n", rotate, rotate)
	return rotate
}

func (c *Cutter) Write(bytes []byte) (n int, err error) {
	c.mutex.Lock()
	defer func() {
		if c.file == nil {
			_ = c.file.Close()
			c.file = nil
		}
		c.mutex.Unlock()
	}()
	length := len(c.formats)
	values := make([]string, 0, 3+length)
	values = append(values, time.Now().Format(c.layout))
	if c.layout != "" {
		values = append(values, c.formats...)
	}
	for i := 0; i < length; i++ {
		values = append(values, c.formats[i])
	}
	values = append(values, c.level+".log")
	filename := filepath.Join(values...)
	director := filepath.Dir(filename)
	err = os.MkdirAll(director, os.ModePerm)
	if err != nil {
		return 0, err
	}
	err = removeNDaysFolders(c.director, c.retentionDay)
	if err != nil {
		return 0, err
	}
	c.file, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return 0, err
	}
	return c.file.Write(bytes)
}

func (c *Cutter) Sync() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.file == nil {
		return c.file.Sync()
	}
	return nil
}

// 增加日志目录文件清理 小于等于零的值默认忽略不再处理
func removeNDaysFolders(dir string, days int) error {
	if days <= 0 {
		return nil
	}
	cutoff := time.Now().AddDate(0, 0, -days)
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.ModTime().Before(cutoff) && path != dir {
			err := os.RemoveAll(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
