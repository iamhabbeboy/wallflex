<script lang="ts">
  import {
    GetDownloadedImages,
    SelectImageDir,
    DownloadImages,
  } from '../../wailsjs/go/main/App.js';

  import { navigate } from 'svelte-routing';

  import LoaderImage from '../../src/assets/images/loader.svg';
  import { imagePathStore } from '../store/app';

  import Modal from '../components/Modal.svelte';

  import { onMount } from 'svelte';

  import rpc from '../rpc';
  let images: string[] = [];
  let path: string = '';
  let isFolderSelected: boolean = false;
  let isLoading = true;

  imagePathStore.subscribe((value) => {
    path = value;
  });

  function dispatcher(image: string) {
    console.log('Clicked...');
    imagePathStore.set(image);
  }

  function setImagePath(image: string) {
    path = image;
  }

  onMount(async (): Promise<any> => {
    const result = await GetDownloadedImages();
    images = result ?? [];
    isLoading = false;

    rpc.on('shortcut.page.setting', () => {
      navigate('/setting', { replace: true });
    });
  });
</script>

<template>
  <section class="w-[95%] mx-auto mt-14">
    {#if isLoading}
      <div class="mx-auto w-48 h-screen flex justify-center items-center">
        <img src={LoaderImage} alt="" />
        <h4 class="font-bold">Processing...</h4>
      </div>
    {:else}
      <div class="flex flex-wrap justify-between">
        {#each images as image}
          <div
            class="bg-gray-800 w-[300px] h-[300px] mb-3 cursor-pointer"
            on:click={() => setImagePath(image)}
            on:keydown={() => setImagePath(image)}
          >
            <img
              src={image.toString()}
              alt=""
              class="object-cover h-[100%] w-[100%]"
            />
          </div>
        {/each}
      </div>
    {/if}
    <Modal {path} />
  </section>
</template>
