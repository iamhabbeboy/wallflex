<script lang="ts">
  import Navigation from '../components/Navigation.svelte';
  import { GetGradient } from '../../wailsjs/go/main/App.js';

  let style: string = 'linear';

  let commonColors = [
    'green',
    'blue',
    'gray',
    'pink',
    'orange',
    'brown',
    'red',
    'violet',
    'yellow',
    'purple',
    'black',
    '#2589bd',
    '#a3b4a2',
    '#38686a',
    '#023047',
    '#ffb703',
  ];

  const handleGenerateColor = () => {
    console.log('Clicked');
    GetGradient().then(() => {});
  };

  const handleSelectGradientStyle = (gradient: string) => {
    style = gradient;
    console.log(gradient);
  };
</script>

<template>
  <section>
    <Navigation />

    <div class="mt-20 text-gray-900">
      <div class="flex flex-wrap w-full p-2">
        {#each commonColors as c, key}
          <div
            style="background-color: {c}"
            class={`mt-5 cursor-default transition-all hover:opacity-70 rounded-full pt-9 text-center w-[100px] h-[100px] ml-5`}
          ></div>
        {/each}
      </div>

      <div class="mt-10 flex mx-auto w-1/2">
        <button
          on:click={() => handleSelectGradientStyle('linear')}
          class={`bg-gray-500 cursor-default text-white py-2 px-10 mb-3 rounded-l-md`}
        >
          Linear {style === 'linear' ? 'selected' : 'n.a'}
        </button>
        <button
          on:click={() => handleSelectGradientStyle('radial')}
          class="hover:bg-gray-500 cursor-default bg-gray-900 text-white py-2 px-10 mb-3"
        >
          Radical
        </button>
        <button
          on:click={() => handleSelectGradientStyle('spiral')}
          class=" hover:bg-gray-500 cursor-default bg-gray-900 text-white py-2 px-10 mb-3 rounded-r-md"
        >
          Spiral
        </button>
      </div>

      <p class="text-white">{style} is here</p>
      <div class="mt-10">
        <button
          on:click={handleGenerateColor}
          class="bg-gray-900 text-white py-2 px-10 mb-3 rounded-md"
          >Generate
        </button>
      </div>
    </div>
  </section>
</template>
