<script lang="ts">
  import Navigation from '../components/Navigation.svelte';
  import { GetGradientImage } from '../../wailsjs/go/main/App.js';

  let style: string = 'linear';
  let colors: string[] = [];

  let commonColors = [
    '#0D4715',
    '#003092',
    '#666666',
    '#FFFFFF',
    '#FF8000',
    '#5F264A',
    '#AC1754',
    '#441752',
    '#493D9E',
    '#66D2CE',
    '#000000',
    '#2589bd',
    '#a3b4a2',
    '#38686a',
    '#023047',
    '#ffb703',
  ];

  const handleGenerateColor = async () => {
    const c1 = hexToRGB(colors[0], 255);
    const c2 = hexToRGB(colors[1], 255);

    const res = await GetGradientImage([c1, c2] as any);
  };

  function hexToRGB(hex: string, alpha: number) {
    const obj = { R: 0, G: 0, B: 0, A: 0 };

    hex = hex.trim();

    hex = hex.replace(/^#/, '');

    if (hex.length === 3) {
      hex = hex
        .split('')
        .map((x) => x + x)
        .join('');
    }

    const bigint = parseInt(hex, 16);
    const r = (bigint >> 16) & 255;
    const g = (bigint >> 8) & 255;
    const b = bigint & 255;

    obj.R = r || 0;
    obj.G = g || 0;
    obj.B = b || 0;

    if (alpha) {
      obj.A = alpha;
    }
    return obj;
  }

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
