<script type="ts">
  import { fade, slide } from 'svelte/transition';
  import { imagePathStore } from '../store/app';
  import CloseIcon from '../assets/images/close.svg';
  import { SetWallpaper } from '../../wailsjs/go/main/App.js';
  import { BrowserOpenURL } from '../../wailsjs/runtime';

  export let path: string;
  let image: string = 'dsfsdf';

  function getUsername(imageFileName: string) {
    const [_, user] = imageFileName.split('_@');
    if (!user) return '';

    const [u] = user.split('.');
    if (!u) return '';

    return `@${u}`;
  }

  /*imagePathStore.subscribe((value) => {
    path = value;
  });*/

  const closeModal = () => {
    path = '';
    return;
  };

  function openBrowser(uri: string) {
    BrowserOpenURL(`https://${uri}`);
  }

  const handleKeydown = (event: KeyboardEvent) => {
    if (event.key === 'Escape') {
      closeModal();
    }
  };

  const setAsWallpaper = (imgPath: string) => {
    SetWallpaper(imgPath).then((res) => {
      closeModal();
    });
  };
</script>

<template>
  <div>
    {#if path !== ''}
      <div
        class="modal-background"
        on:click={closeModal}
        on:keydown={handleKeydown}
      >
        <div class="modal" on:click|stopPropagation on:keydown|stopPropagation>
          <div class="layout">
            <button
              class=" p-2 text-sm underline bg-transparent text-gray-900 hover:text-gray-600"
              on:click={() => setAsWallpaper(path)}>Set as wallpaper</button
            >
            <a href="#/" class="close-btn" on:click|preventDefault={closeModal}>
              <img src={CloseIcon} width="20" alt="close icon" />
            </a>
          </div>
          {#if getUsername(path) !== ''}
            <div class="align-left text-gray-700 text-xs py-2">
              <span class="mr-3"
                >Credit: <a
                  href="#/"
                  on:click={() => openBrowser('unsplash.com')}>Unsplash</a
                ></span
              >
              /
              <span class="ml-3"
                >Author: <a
                  on:click={() =>
                    openBrowser(`unsplash.com/${getUsername(path)}`)}
                  href="#/"
                  class="underline hover:no-underline"
                  >{getUsername(path)}
                </a></span
              >
            </div>
          {/if}
          <div class="overflow-hidden w-[800px] h-[600px] mx-auto">
            <img
              src={path}
              alt=""
              width="100%"
              height="100%"
              class="w-full h-full object-cover"
            />
          </div>
          <p></p>
        </div>
      </div>
    {/if}
  </div>
</template>

<style>
  .layout {
    display: flex;
    justify-content: space-between;
    margin: 5px 0px;
  }

  .modal-background {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
    color: #999;
  }

  .modal {
    background-color: rgba(255, 255, 255, 0.7);
    padding: 0.9rem;
    border-radius: 8px;
    width: 850px;
    max-width: 100%;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
  }

  .close-btn {
    background-color: #fff;
    border: none;
    border-radius: 50%;
    color: white;
    width: 25px;
    height: 25px;
    padding: 4px 3px 0px;
    cursor: pointer;
    color: #999;
  }
</style>
