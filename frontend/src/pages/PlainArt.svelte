<script lang="ts">
  import Navigation from '../components/Navigation.svelte';
  import { GetGradientImage, GetHexToRGBA } from '../../wailsjs/go/main/App.js';

  let style: string = 'linear';
  let colors: string[] = [];

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

  const handleGenerateColor = async () => {
    console.log(colors);
    //const res = await GetGradientImage(colors);
    const res = await GetHexToRGBA(colors[0]);
    console.log(res);
  };

  const handleSelectGradientStyle = (gradient: string) => {
    style = gradient;
  };

  const handleSelectColor = (color: string) => {
    if (colors.includes(color)) {
      colors = colors.filter((c) => c !== color);
      return;
    }
    colors = [...colors, color];
  };
</script>

<section>
  <Navigation />

  <div class="mt-20 text-gray-900">
    <div class="flex flex-wrap w-full p-2">
      {#each commonColors as c, key}
        <div
          style="background-color: {c}"
          on:click={() => handleSelectColor(c)}
          on:keydown={() => handleSelectColor(c)}
          class="mt-5 cursor-default transition-all hover:opacity-70 rounded-full pt-9 text-center w-[100px] h-[100px] ml-5"
        ></div>
      {/each}
    </div>

    <div class="mt-10 flex mx-auto w-1/2">
      <button
        on:click={() => handleSelectGradientStyle('linear')}
        class="{style === 'linear'
          ? 'bg-gray-500'
          : 'bg-gray-900'} hover:bg-gray-800 transition-all cursor-default text-white py-2 px-10 mb-3 rounded-l-md"
      >
        Linear
      </button>

      <button
        on:click={() => handleSelectGradientStyle('radical')}
        class="{style === 'radical'
          ? 'bg-gray-500'
          : 'bg-gray-900'}  hover:bg-gray-800 transition-all cursor-default text-white py-2 px-10 mb-3"
      >
        Radical
      </button>
      <button
        on:click={() => handleSelectGradientStyle('spiral')}
        class="{style === 'spiral' ? 'bg-gray-500' : 'bg-gray-900'}
            hover:bg-gray-800 cursor-default text-white transition-all py-2 px-10 mb-3 rounded-r-md"
      >
        Spiral
      </button>
    </div>

    <div class="flex flex-wrap w-full p-2">
      {#each colors as color}
        <div
          on:click={() => (colors = colors.filter((c) => c !== color))}
          on:keydown={() => (colors = colors.filter((c) => c !== color))}
          style="background-color: {color}"
          class="mt-5 cursor-default transition-all hover:opacity-70 rounded-full pt-9 text-center w-[100px] h-[100px] ml-5"
        ></div>
      {/each}
    </div>

    <div class="mt-10">
      <button
        on:click={handleGenerateColor}
        class="bg-gray-900 text-white py-2 px-10 mb-3 rounded-md"
        >Generate
      </button>
    </div>
  </div>
</section>
