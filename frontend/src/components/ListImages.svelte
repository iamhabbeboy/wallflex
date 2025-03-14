<script lang="ts">
  import Modal from '../components/Modal.svelte';

  export let images: string[];
  let path: string = '';

  function setImagePath(image: string) {
    path = image;
  }
</script>

<template>
  <section class="w-[95%] mx-auto mt-14">
    <div class="flex flex-wrap justify-between">
      {#each images as image}
        <div
          class="bg-gray-800 w-[300px] h-[300px] mb-3 cursor-default"
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
    <Modal {path} />
  </section>
</template>
