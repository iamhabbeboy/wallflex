<script lang="ts">
  import Layout from '../components/Layout.svelte';
  import {
    GetConfig,
    SetConfig,
    OpenDirDialogWindow,
    MessageDialog,
  } from '../../wailsjs/go/main/App.js';

  import { onMount } from 'svelte';
  import type { Configuration } from '../types/Config';

  let imageCategory = '';
  let totalImageCount = 0;
  let imageInterval = '';
  let message = '';
  let defaultPath = '';
  let apikey = '';
  let scheduleImageDownloadInterval = '';
  let hasAutoDownloadEnabled = false;
  //@TODO
  // The goal of this is to disable/enable for auto wallpaper. When enable, it should add a binary file to the /usr/local/bin directory. 
  // When disabled, it should unload the file
  let isAutoWallpaperEnabled = false; 

  async function handleSaveSetting() {
    if (imageCategory === '' || totalImageCount === 0 || imageInterval === '') {
      return;
    }

    const conf = {
      ImageCategory: imageCategory,
      TotalImage: Number(totalImageCount),
      DefaultPath: defaultPath,
      Interval: imageInterval,
      Apikey: apikey,
      ScheduleDownloadInterval: scheduleImageDownloadInterval,
      HasAutoDownloadEnabled: hasAutoDownloadEnabled,
    };
    SetConfig(conf);
    try {
      await MessageDialog('Config updated successfully');
    } catch (e) {
      const error = e instanceof Error ? e.message : 'Unknown error';
      message = error;
    }
  }

  onMount(async () => {
    const conf = (await GetConfig()) as Configuration;
    imageCategory = conf.ImageCategory;
    totalImageCount = conf.TotalImage;
    imageInterval = conf.Interval;
    defaultPath = conf.DefaultPath;
    apikey = conf.Apikey;
    scheduleImageDownloadInterval = conf.ScheduleDownloadInterval;
    hasAutoDownloadEnabled = conf.HasAutoDownloadEnabled;
  });

  const handleSelectFolder = async () => {
    const path = await OpenDirDialogWindow();
    if (path === '') {
      defaultPath = defaultPath;
    } else {
      defaultPath = path;
    }
  };

  const handleRestoreSetting = async () => {
    const conf = {
      ImageCategory: 'nature',
      TotalImage: 10,
      DefaultPath: '.wallflex/images',
      Interval: '30s',
      Apikey: '',
      scheduleImageDownloadInterval: '1w',
      HasAutoDownloadEnabled: false,
    };
    imageCategory = conf.ImageCategory;
    totalImageCount = conf.TotalImage;
    defaultPath = conf.DefaultPath;
    imageInterval = conf.Interval;
    apikey = conf.Apikey;
    hasAutoDownloadEnabled = conf.HasAutoDownloadEnabled;

    SetConfig(conf);
    try {
      await MessageDialog('Config restored successfully');
    } catch (e) {
      const error = e instanceof Error ? e.message : 'Unknown error';
      message = error;
    }
  };

  function setImageCategory(query: string) {
    imageCategory = query;
  }

  function handleRemoveWhiteSpace(event) {
    const value = event.target.value;
    event.target.value = value.replace(/\s/g, '');
  }

  function handleToggleAutoDownload(event) {
    hasAutoDownloadEnabled = event.target.checked;
  }

  function handleToggleAutoWallpaper(event) {
    isAutoWallpaperEnabled = event.target.checked;
  }
</script>

<template>
  <Layout>
    <h1 class="font-bold text-gray-600 dark:text-white">Configuration</h1>
    <div class="layout">
      <div
        class="my-4 text-[#999] border border-gray-200 dark:border-gray-500 rounded-md
"
      >
        <div class="selection">
          <div>
            <label for="imagepath" class="form-label"
              >Selected directory path <p class="italic">
                {defaultPath}
              </p></label
            >
            <button
              class="text-gray-100 py-2 px-5 mb-3 cursor-default rounded-md bg-gray-500"
              on:click={handleSelectFolder}
            >
              Change folder</button
            >
          </div>
          <div>
            <label for="imagepath" class="form-label"> Image query</label>
            <input
              type="text"
              class="border border-gray-400 w-6/12 outline-none p-2 rounded-md"
              spellcheck="false"
              autocomplete="off"
              id="image-query"
              bind:value={imageCategory}
              on:keydown={handleRemoveWhiteSpace}
            />
            <div class="cursor-default">
              <i class="block">Suggestions:</i>
              <span
                class="underline"
                on:click={() => setImageCategory('nature')}
                on:keydown={() => setImageCategory('nature')}
              >
                nature</span
              >,
              <span
                class="underline"
                on:click={() => setImageCategory('landscape')}
                on:keydown={() => setImageCategory('landscape')}>landscape</span
              >,
              <span
                class="underline"
                on:click={() => setImageCategory('lamborghini')}
                on:keydown={() => setImageCategory('lamborghini')}
                >lamborghini</span
              >,
              <span
                class="underline"
                on:click={() => setImageCategory('tokyo')}
                on:keydown={() => setImageCategory('tokyo')}
              >
                tokyo</span
              >,
              <span
                class="underline"
                on:click={() => setImageCategory('abstract')}
                on:keydown={() => setImageCategory('abstract')}>abstract</span
              >
            </div>
          </div>
          <div class="my-2">
            <label for="imagepath" class="form-label"> Total image </label>
            <select
              bind:value={totalImageCount}
              class="border border-gray-400 p-2 h-11 w-6/12 rounded-md outline-none"
            >
              <option value={3} selected={totalImageCount === 3}>3</option>
              <option value={5} selected={totalImageCount === 5}>5</option>
              <option value={10} selected={totalImageCount === 10}>10</option>
              <option value={15} selected={totalImageCount === 15}>15</option>

              <option value={20} selected={totalImageCount === 20}>20</option>
            </select>
          </div>
          <div class="bg-gray-200 p-2 rounded-md w-6/12 mx-auto my-5">
            <label class="inline-flex items-center cursor-pointer">
              <input type="checkbox" value="" class="sr-only peer" on:change={handleToggleAutoDownload} checked={hasAutoDownloadEnabled}/>
              <div
                class="relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-green-600 dark:peer-checked:bg-green-600"
              ></div>
              <span class="ms-3 text-sm font-medium">Auto download images</span>
            </label>
          </div>
          <div>
            <label for="imagepath" class="form-label">
              Image change interval </label
            >
            <select
              bind:value={imageInterval}
              class="border border-gray-400 p-2 h-11 w-6/12 rounded-md outline-none"
            >
              <option value="30s" selected={imageInterval === '30s'}>30s</option
              >
              <option value="5m" selected={imageInterval === '5m'}>5m</option>

              <option value="10m" selected={imageInterval === '10m'}>10m</option
              >
              <option value="30m" selected={imageInterval === '30m'}>30m</option
              >
              <option value="1h" selected={imageInterval === '1h'}>1h</option>
            </select>
          </div>

          <div class="my-2">
            <label for="imagepath" class="form-label"> Schedule Image Download Interval </label>
            <select
              bind:value={scheduleImageDownloadInterval}
              class=" border border-gray-400 p-2 h-11 w-6/12 rounded-md outline-none"
            >
              <option value="30s" selected={scheduleImageDownloadInterval === '30s'}>3s</option
              >
              <option value="5m" selected={scheduleImageDownloadInterval === '5m'}>5m</option>

              <option value="10m" selected={scheduleImageDownloadInterval === '10m'}>10m</option
              >
              <option value="30m" selected={scheduleImageDownloadInterval === '30m'}>30m</option
              >
              <option value="1w" selected={scheduleImageDownloadInterval === '1w'}>1w</option>
            </select>
          </div>
          <div class="mt-2">
            <label for="imagepath" class="form-label"> Unsplash API key</label>
            <input
              type="text"
              class="border border-gray-400 p-2 w-6/12 rounded-md outline-none"
              spellcheck="false"
              autocomplete="off"
              id="image-query"
              bind:value={apikey}
            />
          </div>

          <div class="bg-gray-200 p-2 rounded-md w-6/12 mx-auto my-5">
            <label class="inline-flex items-center cursor-pointer">
              <input type="checkbox" value="" class="sr-only peer" on:change={handleToggleAutoWallpaper} checked={isAutoWallpaperEnabled}/>
              <div
                class="relative w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-green-600 dark:peer-checked:bg-green-600"
              ></div>
              <span class="ms-3 text-sm font-medium">Auto change wallpaper</span>
            </label>
          </div>
          <div class="mt-5">
            <button
              class="text-gray-100 py-2 px-10 cursor-default mb-3 rounded-md bg-gray-500"
              on:click={handleSaveSetting}
            >
              Save
            </button>
            <button
              class="text-gray-100 py-2 px-10 mb-3 cursor-default rounded-md bg-gray-500"
              on:click={handleRestoreSetting}
            >
              Restore config
            </button>
          </div>
        </div>
      </div>
      <div
        class="my-4 text-[#999] block h-5 dark:border-gray-500 rounded-md"
      ></div>
    </div>
  </Layout>
</template>

<style>
  .layout {
    width: 40rem;
    margin: auto;
    font-size: 14px;
  }

  .selection {
    padding: 30px 20px 20px 20px;
  }

  .selection .form-label {
    display: block;
    margin-top: 10px;
  }

  select,
  input {
    appearance: none;
    -webkit-appearance: none;
  }
</style>
