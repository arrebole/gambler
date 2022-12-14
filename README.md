## 交易回测系统

### 后端接口


#### 获取随机的股票
> Get /random

返回结果

```typescript
type Response = {
  code: string
  symbol: string
  name: string
  area: string
  industry: string
  market: string
  minDate: string
  maxDate: string
}
```

#### 获取股票的 k 线图
> Get /klines

| 参数名  | 描述 |
| ---    | ---  |
| code    | 股票代码 |
| begin  | k线开始的时间 |
| latest  | k线结束的时间 |
| level  | k线的级别    |

返回结果

```typescript
interface KLineData {
  timestamp: number
  open: number
  close: number
  high: number
  low: number
  volume: number
}

type Response = KLineData[]

```