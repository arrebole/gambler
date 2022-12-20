
<template>
  <section class="border w-full">
    <div class="flex text-gray-300 border-b">
      <button class="p-2.5 text-gray-700 hover:bg-gray-100" @click="handleChageLevel('1m')">1m</button>
      <button class="p-2.5 text-gray-700 hover:bg-gray-100" @click="handleChageLevel('5m')">5m</button>
      <button class="p-2.5 text-gray-700 hover:bg-gray-100" @click="handleChageLevel('15m')">15m</button>
      <button class="p-2.5 text-gray-700 hover:bg-gray-100" @click="handleChageLevel('30m')">30m</button>
      <button class="p-2.5 text-gray-700 hover:bg-gray-100" @click="handleChageLevel('1D')">1D</button>
      <button class="p-2.5 text-gray-700 hover:bg-gray-100" @click="handleChageLevel('1W')">1W</button>
      <button class="p-2.5 text-gray-700 hover:bg-gray-100" @click="handleChageLevel('1M')">1M</button>
    </div>
    <div id="charts" class="w-full h-full"></div>
  </section>

  <section class="border p-3">
    <button
      class="rounded-lg w-24 px-4 py-1.5 text-base font-semibold leading-7 text-gray-900 ring-1 ring-gray-900/10 hover:ring-gray-900/20"
      @click="handleNextKline">等待</button>
  </section>
</template>

<script lang="ts" setup>
import klinecharts from 'klinecharts';
import { onMounted } from 'vue';
import { Level } from './api/types';
import { createKlineChart } from './chart';
import { DataSource } from './datasource';

let chart: klinecharts.Chart;
let currentLevel: Level = '1D';
const dataSource = new DataSource();

// 获取数据
onMounted(async () => {
  // 渲染图表
  await dataSource.init();
  chart = createKlineChart('charts');
  chart.applyNewData(
    await dataSource.klines(currentLevel),
  );
});

// 切换级别
async function handleChageLevel(level: Level) {
  if (!chart) {
    return;
  }
  currentLevel = level
  chart.applyNewData(
    await dataSource.klines(currentLevel),
  );
}

// 下一根K线
function handleNextKline() {
  if (!chart) {
    return;
  }
  chart.updateData(
    dataSource.nextKline(currentLevel),
  )
}

</script>
