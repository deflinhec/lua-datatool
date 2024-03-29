# Lua data tool

Lua 工具集

## 圖形化數據編輯工具 (datatool)

> 透過Imgui圖形化介面編輯數據．

- [語法檢查](#語法檢查)

- [數據校驗和](#數據校驗和)

- [通用工具指令](#通用工具指令)

- [檔案規範](#檔案規範)

|長參數|短參數|選填|說明|預設值|範例|
|-|-|-|-|-|-|
|--help|-h|✔️| 幫助說明|-|-|
|--version|-v|✔️| 檢視程序建置版號|-|-|
|--pprof|-p|✔️| pprof profiling|-|--pprof|
|--launch|-l|✔️| 啟動時開啟檔案|-|--pprof test.lua|

## 格式化工具 (formater)

> 對執行目錄下所有lua腳本進行格式化．

- [語法檢查](#語法檢查)

- [數據校驗和](#數據校驗和)

- [索引鍵排序](#索引鍵排序)

    排除重複索引鍵

- [自動縮排、對齊](#自動縮排、對齊)

- [格式化差異比對](#格式化差異比對)

|長參數|短參數|選填|說明|預設值|範例|
|-|-|-|-|-|-|
|--help|-h|✔️| 幫助說明|-|-|
|--version|-v|✔️| 檢視程序建置版號|-|-|
|--path|-p|✔️| 目標檔案、目錄|.|-|
|--ignore|-i|✔️| 忽略特定檔案名稱|-|--ignore=init.lua|
|--dry-run|-|✔️| 將原始檔案另存為 **tmp.lua** 並執行 [格式化差異比對](#格式化差異比對)|-|-|

---

### 主要功能

## 語法檢查

經由內嵌Lua直譯器對腳本進行直譯、載入，並從中提取數據．

## 數據校驗和

```lua
module("Module")
md5sum = md5sum or {}
md5sum.Field="97a15070f5f8c3bfe47678c5409471f6"
Field=
{
    -- 格式化內容
}
```

根據 ***Field*** 中的數據內容產生 數據校驗和，以便於執行期驗證。

## 自動縮排、對齊

```lua
-- $id$
module("Module")
md5sum = md5sum or {}
md5sum.Field="d0796fb16aed2249755dc974b18650fd"
Field=
{
    { -- Module.Field[0]
        number=1,
        table={ -- Data.RoomData[0].table
            key=1,
        },
        array={ -- total: 5
            1,2,3,4,5,
        },
        float=0,
    },
}
```

## 索引鍵排序

```lua
-- $id$
module("Module")
md5sum = md5sum or {}
md5sum.Field="cfcae4802f729e9fc167843df87c6a28"
Field=
{ -- Module.Field
    [0]="string",
    [3]="string",
    [99]={ -- Module.Field[99]
        key4=4,
        key1=1,
        key3=3,
    },
    [-1]="string",
}
```

更多格式化範例請參考 [edge.lua](./doc/tests/edge.lua)。

### 檔案規範

```lua
module("Module")
Field=
{
    -- (未)格式化內容
}
```

- 檔案必須為有效lua語法。

- 數據必須為全域變數。

    ```lua
    print(assert(Module.Field))
    ```

- 數據最少為兩個層級。

    ```lua
    print(assert(assert(Module).Field))
    ```

### 格式化差異比對

該專案使用 ***gopher-lua*** 模塊，使用 **lua虛擬機** 加載 **原始檔案(.lua)** ＆ **輸出檔案(.tmp.lua)** 
至記憶體中，對於完整的 **樹狀表(table)** 結構比對差異．




