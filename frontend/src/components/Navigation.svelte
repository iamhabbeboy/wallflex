<script lang="ts">
  import {
    GetDownloadedImages,
    SelectImageDir,
    DownloadImages,
  } from '../../wailsjs/go/main/App.js';

  import { link } from 'svelte-routing';

  import DownloadImage from '../../src/assets/images/download.svg';
  import ConfigImage from '../../src/assets/images/config.svg';

  import PlainArtImage from '../../src/assets/images/plain.svg';

  import AbstractArt from '../../src/assets/images/art.svg';

  import AIArt from '../../src/assets/images/ai.svg';

  let isLoading = true;

  let images: string[] = [];

  async function downloadImages() {
    try {
      isLoading = true;
      const res = await DownloadImages();
    } catch (e) {
    } finally {
      const result = await GetDownloadedImages();
      images = result ?? [];
      isLoading = false;
    }
  }
</script>

<template>
  <div
    class="bg-gray-50 p-3 border-b dark:border-gray-600 flex justify-between fixed w-[100%] top-0 left-0"
  >
    <div class="flex">
      <a href="#" class="dark:text-gray-50 text-gray-500 text-xs flex">
        <img
          src={DownloadImage}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        /> Download images</a
      >

      <a href="#/" class="dark:text-gray-50 text-gray-500 text-xs flex ml-10">
        <img
          src={PlainArtImage}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        /> Plain</a
      >

      <a href="#/" class="dark:text-gray-50 text-gray-500 text-xs flex ml-10">
        <img
          src={AbstractArt}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        /> Abstract Art</a
      >

      <a
        href="#/"
        on:click={downloadImages}
        class="dark:text-gray-50 text-gray-500 text-xs flex ml-10"
      >
        <img
          src={AIArt}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        /> AI
      </a>
    </div>
    <div>
      <a
        href="/setting"
        class=" text-xs flex text-gray-500 dark:text-gray-50"
        use:link
      >
        <img
          src={ConfigImage}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        /> Configuration
      </a>
    </div>
  </div>
</template>
