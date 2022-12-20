import * as klinecharts from 'klinecharts';

export function createKlineChart(dom: string) {
  const styleOptions = {
    grid: {
      show: true,
      horizontal: {
        show: true,
        size: 1,
        color: '#EDEDED',
        style: 'dash',
        dashValue: [2, 2]
      },
      vertical: {
        show: false,
        size: 1,
        color: '#EDEDED',
        style: 'dash',
        dashValue: [2, 2]
      }
    },
    candle: {
      type: 'candle_solid',
      bar: {
        upColor: '#EF5350',
        downColor: '#26A69A',
        noChangeColor: '#888888'
      },
    },
    priceMark: {
      show: true,
      high: {
        show: true,
        color: '#D9D9D9',
        textMargin: 5,
        textSize: 10,
        textFamily: 'Helvetica Neue',
        textWeight: 'normal'
      },
      low: {
        show: true,
        color: '#D9D9D9',
        textMargin: 5,
        textSize: 10,
        textFamily: 'Helvetica Neue',
        textWeight: 'normal',
      },
      last: {
        show: true,
        upColor: '#26A69A',
        downColor: '#EF5350',
        noChangeColor: '#888888',
        line: {
          show: true,
          style: 'dash',
          dashValue: [4, 4],
          size: 1
        },
        text: {
          show: true,
          size: 12,
          paddingLeft: 2,
          paddingTop: 2,
          paddingRight: 2,
          paddingBottom: 2,
          color: '#FFFFFF',
          family: 'Helvetica Neue',
          weight: 'normal',
          borderRadius: 2
        }
      }
    },
    tooltip: {
      showRule: 'always',
      showType: 'standard',
      labels: ['时间', '开', '收', '高', '低', '成交量'],
      values: null,
      defaultValue: 'n/a',
      rect: {
        paddingLeft: 0,
        paddingRight: 0,
        paddingTop: 0,
        paddingBottom: 6,
        offsetLeft: 8,
        offsetTop: 8,
        offsetRight: 8,
        borderRadius: 4,
        borderSize: 1,
        borderColor: '#3f4254',
        backgroundColor: 'rgba(17, 17, 17, .3)'
      },
      text: {
        size: 12,
        family: 'Helvetica Neue',
        weight: 'normal',
        color: '#D9D9D9',
        marginLeft: 8,
        marginTop: 6,
        marginRight: 8,
        marginBottom: 0
      }
    }
  }

  const chart = klinecharts.init(dom, styleOptions);

  // 创建一个主图技术指标
  chart.createTechnicalIndicator('EMA', false, { id: 'candle_pane' })

  // 创建一个副图技术指标VOL
  chart.createTechnicalIndicator('VOL')

  // 创建一个副图技术指标MACD
  chart.createTechnicalIndicator('MACD');

  return chart;
}