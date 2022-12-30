package main

import (
	adapter2 "github.com/assimon/svbot/videos/adapter"
	"testing"
)

func TestDouyinAdapter_GetShortVideoInfo(t *testing.T) {
	adapter := &adapter2.DouyinAdapter{}
	resp, err := adapter.GetShortVideoInfo("https://v.douyin.com/MLBnNgY/")
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("Douyin get no watermark download url fail")
	}
}

func TestPipixiaAdapter_GetShortVideoInfo(t *testing.T) {
	adapter := &adapter2.PipixiaAdapter{}
	resp, err := adapter.GetShortVideoInfo("https://h5.pipix.com/s/MLYcVL7/")
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("Pipixia get no watermark download url fail")
	}
}

func TestHuoshanAdapter_GetShortVideoInfo(t *testing.T) {
	adapter := &adapter2.HuoshanAdapter{}
	resp, err := adapter.GetShortVideoInfo("https://share.huoshan.com/hotsoon/s/6LtxUndlBy8/")
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("Huoshan get no watermark download url fail")
	}
}

func TestWeishiAdapter_GetShortVideoInfo(t *testing.T) {
	adapter := &adapter2.WeishiAdapter{}
	resp, err := adapter.GetShortVideoInfo("https://isee.weishi.qq.com/ws/app-pages/share/index.html?wxplay=1&id=7wkzwZSDL1On6v1A2&spid=3712057846823065229&qua=v1_iph_weishi_8.81.0_266_app_a&chid=100081014&pkg=3670&attach=cp_reserves3_1000370011")
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("Weishi get no watermark download url fail")
	}
}

func TestWeiboAdapter_GetShortVideoInfo(t *testing.T) {
	adapter := &adapter2.WeiboAdapter{}
	resp, err := adapter.GetShortVideoInfo("https://video.weibo.com/show?fid=1034:4284200640308284")
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("Weibo get no watermark download url fail")
	}
}

func TestZuiyouAdapter_GetShortVideoInfo(t *testing.T) {
	adapter := &adapter2.ZuiyouAdapter{}
	resp, err := adapter.GetShortVideoInfo("https://share.xiaochuankeji.cn/hybrid/share/post?pid=293152528&zy_to=applink&share_count=1&m=2dcd4cdcc14c53c091cdaba91128dcea&d=491114c4fa959852b33646f0c14bf5c9a96519302501213465277c825f1c48d2481eb4298838cc516821e843a90252a2&app=zuiyou&recommend=r0&name=n0&title_type=t0")
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("Zuiyou get no watermark download url fail")
	}
}

func TestLvZhouAdapter_GetShortVideoInfo(t *testing.T) {
	adapter := &adapter2.LvZhouAdapter{}
	resp, err := adapter.GetShortVideoInfo("https://m.oasis.weibo.cn/v1/h5/share?sid=4822391043132619")
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("LvZhou get no watermark download url fail")
	}
}

func TestKuaiShouAdapter_GetShortVideoInfo(t *testing.T) {
	adapter := &adapter2.KuaiShouAdapter{}
	resp, err := adapter.GetShortVideoInfo("https://v.kuaishou.com/uH5N5k")
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("KuaiShou get no watermark download url fail")
	}
}

func TestXiGuaAdapter_GetShortVideoInfo(t *testing.T) {
	adapter := &adapter2.XiGuaAdapter{}
	resp, err := adapter.GetShortVideoInfo("https://v.ixigua.com/MNuytsN/")
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("Zuiyou get no watermark download url fail")
	}
}
