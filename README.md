# decklog-waifu2x-downloader

## 前言

為了讓我能夠快樂的用桌遊模擬器測試新的牌

我寫了這個東東

最近自學GO 寫得很醜但是能用 傷眼請見諒

## 這是甚麼?

這是一個搭配瀏覽器腳本 把Decklog上的牌組下載下來弄成Tabletop Simulator可以讀的牌組的小工具

我覺得官網的圖有點糊 所以用了waifu2x放大解析度兩倍 這樣在遊戲中看了比較清楚

當然 你不想放大也是可以的 下面會有詳細說明

## 如何開始使用

請先安裝TamperMonkey等瀏覽器腳本外掛

然後安裝以下的腳本

我不會javascript 寫的很可怕請見諒

https://greasyfork.org/zh-TW/scripts/423058-decklog-ws-deck-output

## 下載牌組資訊

啟用這個腳本 到Decklog之後 找到目標牌組

等到頁面跑完之後滾到最左下角點選Create TSDB按鈕

![](https://i.imgur.com/mKYV2Sk.jpg)

這樣會下載一個txt 以供後續使用

![](https://i.imgur.com/pdILlSP.png)


## 下載decklog-waifu2x-downloader

你可以自己編譯 當然 你要有裝Go

```
git clone https://github.com/tail9951/decklog-waifu2x-downloader.git
cd decklog-waifu2x-downloader
go build .
```

或是用我build好的

在Releases中應該可以找到 選那個zip就可以下載了

https://github.com/tail9951/decklog-waifu2x-downloader/releases

下載好後解壓縮 你會看到有一個decklog-waifu2x-downloader.exe跟一個config.ini

## 進行設定

使用前可以先打開config.ini檔

裡面可以設置三種東西

1. tsdb

我建議大家不要動 這就是等等輸出的總圖解析度之類的設定

太大模擬器也不好跑 想玩的可以自己測試 我這邊留一個我平常用的設定

2. waifu2x

我的waifu2x是call別人的api 裡面有免費額度

註冊之後可以在Dashboard看到自己的api-key

![](https://i.imgur.com/ltMVkE5.png)

貼上之後就可以了

或是你嫌麻煩不想放大2倍 那就留空就好了

下載器會省去這一步

3. output

字面意思 幫你等等的輸出資料夾取個名字


## 下載牌組

設定好之後把剛剛的txt拖進去這個程式裡就好了

![](https://i.imgur.com/oYURzJR.png)

![](https://i.imgur.com/fZ4GR46.png)

會多出一個剛剛設定的輸出資料夾 裡面有所有的圖跟tsdb檔

## 怎麼用tsdb檔

如果你有桌遊模擬器 你應該可以在

```
SteamLibrary\steamapps\common\Tabletop Simulator\Modding\Deck Builder
```

以上的路徑找到牌組建構器

![](https://i.imgur.com/7SYvRoz.png)


用Deck Builder打開 點選左上角的File/Export

![](https://i.imgur.com/R6xStSx.jpg)


把圖片輸出 不用做甚麼調整

## 準備你的卡背

就... 準備一張要拿來當卡背的圖

好像能用gif 我沒試過就是

## 訂閱工作坊的WS牌桌

我是用這個

https://steamcommunity.com/sharedfiles/filedetails/?id=2054847285

用的順手 開心就好

## 打開桌遊模擬器

創房間 開牌桌

點選上面的Object 選擇Components

選擇Cards裡的Custom Deck 然後手上應該會有一副半透明的牌 點牌桌會出現詳細資訊

Face就選擇剛剛Export出來的newdeck.png

記得選cloud 這樣才能跟朋友玩 沒朋友的話隨便

Back這邊選卡背

10 * 5  50張卡是一定要的

最下面的Back is hidden記得打勾 這樣別人看你的卡背才能正確辨識

就降 黑放