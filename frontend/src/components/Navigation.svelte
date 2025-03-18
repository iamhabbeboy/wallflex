<script lang="ts">
  import {
    GetDownloadedImages,
    SelectImageDir,
    DownloadImages,
  } from '../../wailsjs/go/main/App.js';

  import { link } from 'svelte-routing';

  import DownloadImage from '../../src/assets/images/download.svg';
  import ConfigImage from '../../src/assets/images/config.svg';

  import { dispatcher } from '../../src/utilities/util';

  let isLoading = true;

  let images: string[] = [];

  async function downloadImages() {
    dispatcher('downloading');
  }
</script>

<template>
  <div
    class=" dark:bg-gray-700 dark:text-gray-50 bg-gray-50 p-3 border-b dark:border-gray-600 flex justify-between fixed w-[100%] top-0 left-0"
  >
    <div class="flex">
      <a
        href="/"
        use:link
        class="dark:text-gray-50 text-gray-500 text-xs flex hover:cursor-default"
      >
        <img
          src={DownloadImage}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        /> Images</a
      >
      <!--
      <a
        href="/plain-art"
        use:link
        class="dark:text-gray-50 text-gray-500 text-xs flex ml-10 hover:cursor-default"
      >
        <img
          src={PlainArtImage}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        /> Plain Art</a
      >

      <a
        href="/abstract-art"
        class="dark:text-gray-50 text-gray-500 text-xs flex ml-10 hover:cursor-default"
      >
        <img
          src={AbstractArt}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        /> Abstract Art</a
      >

       <a
        href="/ai"
        class="dark:text-gray-50 text-gray-500 text-xs flex ml-10 hover:cursor-default"
      >
        <img
          src={AIArt}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        /> AI
      </a>

       <a
        href="#"
        class="dark:text-gray-50 text-gray-500 text-xs flex ml-10 hover:cursor-default"
      >
        <img
          src={DownloadImage}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        /> Upload folder</a
      >-->
    </div>
    <div class="flex">
      <a
        href="#"
        on:click|preventDefault={downloadImages}
        class="dark:text-gray-50 text-gray-500 text-xs flex mr-5 hover:cursor-default"
      >
        <img
          src={DownloadImage}
          width="15"
          alt=""
          class="mr-1 dark:brightness-0 dark:invert-[1]"
        />Download
      </a>

      <a
        href="/setting"
        class=" text-xs flex text-gray-500 dark:text-gray-50 hover:cursor-default"
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
