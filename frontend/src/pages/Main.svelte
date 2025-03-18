<script lang="ts">
  import Navigation from '../components/Navigation.svelte';
  import ListImages from '../components/ListImages.svelte';
  import {
    GetDownloadedImages,
    DownloadImages,
    MessageDialog,
  } from '../../wailsjs/go/main/App.js';
  import { navigate } from 'svelte-routing';
  import LoaderImage from '../../src/assets/images/loader.svg';
  import NoImagesPlaceHolder from '../../src/assets/images/color-art.svg';
  import { onMount, onDestroy } from 'svelte';

  import rpc from '../rpc';
  import { imagePathStore } from '../../src/store/app';
  let images: string[] = [];
  let isLoading = true;
  let isVisible = true;

  const disableRightClick = (event) => {
    event.preventDefault();
  };

  const handleVisibilityChange = () => {
    isVisible = document.visibilityState === 'visible';
    if (isVisible) {
      location.reload();
    }
  };

  async function downloadImages() {
    try {
      await DownloadImages();
    } catch (e) {
      console.log(e);
      MessageDialog(e);
    } finally {
      const result = await GetDownloadedImages();
      images = result ?? [];
      isLoading = false;
    }
  }

  onMount(async (): Promise<any> => {
    try {
      const result = await GetDownloadedImages();
      images = result ?? [];
      isLoading = false;
    } catch (e) {
      MessageDialog(e);
    } finally {
      isLoading = false;
    }

    rpc.on('shortcut.page.setting', () => {
      navigate('/setting', { replace: true });
    });

    imagePathStore.subscribe(async (value) => {
      if (value === 'downloading') {
        isLoading = true;
        await downloadImages();
      }
    });
    isLoading = false;

    document.addEventListener('visibilitychange', handleVisibilityChange);
    //document.addEventListener('contextmenu', disableRightClick);
    return () => {
      // document.removeEventListener('contextmenu', disableRightClick);
    };
  });

  onDestroy(() => {
    document.removeEventListener('visibilitychange', handleVisibilityChange);
  });
</script>

<div class="text-gray-400">
  <Navigation />
  {#if !isLoading && images.length === 0}
    <div class="mt-32 w-4/12 mx-auto text-center flex flex-col">
      <img src={NoImagesPlaceHolder} alt="" width="300" />
      <h1 class="text-2xl font-bold text-gray-400">
        No Images found in your library
      </h1>
      <p class="mt-3">To get started click on the download button.</p>
    </div>
  {/if}
  {#if isLoading}
    <div class="mx-auto w-48 h-screen flex justify-center items-center">
      <img src={LoaderImage} alt="" />
      <h4 class="font-bold">Processing...</h4>
    </div>
  {:else}
    <ListImages {images} />
  {/if}
</div>
