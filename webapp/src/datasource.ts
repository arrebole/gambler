import { randomStock, fetchKlines } from './api';
import { Kline, Level, Stock } from './api/types';

export class DataSource {
  private currentTimestamp: number;
  private klineData: Map<Level, Kline[]> = new Map();
  private stock: Stock;

  async init() {
    this.stock = await randomStock();
    this.klineData.set('1D', await fetchKlines({
      level: '1D',
      code: this.stock.code,
      begin: this.stock.gameBeginAt,
      latest: this.stock.gameEndAt,
    }))
    // 默认从第100天开始
    // @ts-ignore
    this.currentTimestamp = this.klineData.get('1D')[100].timestamp;
  }

  async klines(level: Level) {
    if (!this.klineData.has(level)) {
      this.klineData.set(level, await fetchKlines({
        level: level,
        code: this.stock.code,
        begin: this.stock.gameBeginAt,
        latest: this.stock.gameEndAt,
      }));
    }
    const result = [];
    for (const item of this.klineData.get(level)!) {
      const kline = item as Kline;
      if (kline.timestamp <= this.currentTimestamp) {
        result.push(kline);
      }
    }
    return result;
  }

  nextKline(level: Level) {
    for (const item of this.klineData.get(level)!) {
      const kline = item as Kline;
      if (kline.timestamp > this.currentTimestamp) {
        this.currentTimestamp = kline.timestamp;
        return kline;
      }
    }
    return this.klineData.get(level)![0];
  }
}