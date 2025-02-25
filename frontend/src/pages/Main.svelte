<script lang="ts">
  import { link } from 'svelte-routing';
  import Modal from '../components/Modal.svelte';
  import { imagePathStore } from '../store/app';
  import { onMount } from 'svelte';
  import ImageConfig from '../components/ImageConfig.svelte';
  import DownloadImage from '../../src/assets/images/download.svg';
  import ConfigImage from '../../src/assets/images/config.svg';

  import Navigation from '../components/Navigation.svelte';
  import ListImages from '../components/ListImages.svelte';

  let images: string[] = [];
  let path: string;
  let isFolderSelected: boolean = false;
  let isLoading = true;

  imagePathStore.subscribe((value) => {
    path = value;
  });

  function dispatcher(image: string) {
    console.log('Clicked...');
    imagePathStore.set(image);
  }

  const disableRightClick = (event) => {
    event.preventDefault();
  };

  onMount(async (): Promise<any> => {
    //document.addEventListener('contextmenu', disableRightClick);
    /*return () => {
      document.removeEventListener('contextmenu', disableRightClick);
    };*/
  });
</script>

<template>
  <div on:contextmenu={disableRightClick}>
    <Navigation />

    <ListImages />
  </div>
</template>
