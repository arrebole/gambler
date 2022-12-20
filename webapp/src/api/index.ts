import dayjs from "dayjs";
import { Stock, Level, Kline } from './types';

function toQueryString(query: any) {
  const result = []
  for (const i of Object.keys(query)) {
    result.push(`${i}=${query[i]}`);
  }
  return result.join('&')
}

export async function randomStock(): Promise<Stock> {

  // 随机返回一支股票
  const stock: Stock = await fetch('/random')
    .then((response) => response.json())

  // 随机选择一天为游戏开始时间
  const count = dayjs(stock.maxDate)
    .subtract(450, 'days')
    .diff(stock.minDate, 'day');

  const randomInt = Math.floor(Math.random() * count);
  const gameBeginAt = dayjs(stock.minDate).add(randomInt, 'day');
  const gameEndAt = gameBeginAt.add(400, 'day');

  return {
    ...stock,
    gameBeginAt: gameBeginAt.unix(),
    gameEndAt: gameEndAt.unix(),
  }
}

export async function fetchKlines(findOptions: {
  level: Level,
  code: string,
  begin: string | number,
  latest: string | number,
}): Promise<Kline[]> {
  // 随机返回一支股票
  const klineData: number[][] = await fetch('/klines?' + toQueryString(findOptions))
    .then((response) => response.json())
  return klineData.map(v => ({
    timestamp: v[0] * 1000,
    open: v[1],
    close: v[2],
    high: v[3],
    low: v[4],
    volume: v[5],
  }))
}