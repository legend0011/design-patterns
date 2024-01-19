package main

import (
	"fmt"
)

type ThirdPartyYouTubeLib interface {
	DownloadVideo(id string) string
}

// A concret implementation of youtube video downloader, focusing on connection and content download.
type ThirdPartyYouTubeClass struct{}

func (d *ThirdPartyYouTubeClass) DownloadVideo(id string) string {
	fmt.Printf("Download video #%s from Youtube\n", id)
	return "VideoContent of " + id
}

// 代理模式最常用的一个应用场景就是，在业务系统中开发一些非功能性需求，比如：监控、统计、鉴权、限流、事务、幂等、日志。
// 我们将这些附加功能与业务功能解耦，放到代理类中统一处理，让程序员只需要关注业务方面的开发。
type CachedYouTubeClass struct {
	thirdPartyYouTube ThirdPartyYouTubeLib // Here should be interface
	cache             map[string]string
}

func (c *CachedYouTubeClass) DownloadVideo(id string) string {
	if video, ok := c.cache[id]; ok {
		return video
	}
	c.cache[id] = c.thirdPartyYouTube.DownloadVideo(id)
	return c.cache[id]
}

// 代理类 UserControllerProxy 和原始类 UserController 实现相同的接口 IUserController。
// UserController 类只负责业务功能。代理类 UserControllerProxy 负责在业务代码执行前后附加其他逻辑代码，并通过委托的方式调用原始类来执行业务代码。
func main() {
	downloader := &ThirdPartyYouTubeClass{}
	cachedDownloader := &CachedYouTubeClass{thirdPartyYouTube: downloader, cache: map[string]string{}}

	fmt.Println("content:" + cachedDownloader.DownloadVideo("123"))
}
