
export type Level = '1m' | '5m' | '15m' | '30m' | '1D' | '1W' | '1M';

export interface Kline {
  timestamp: number
  open: number
  close: number
  high: number
  low: number
  volume: number
}

export interface Stock {
  code: string
  name: string
  symbol: string
  area: string
  industry: string
  market: string
  maxDate: string
  minDate: string
  gameBeginAt: number
  gameEndAt: number
}