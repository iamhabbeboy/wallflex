<script>
  import { link } from 'svelte-routing';

  import IconifyIcon from '@iconify/svelte';
</script>

<template>
  <div>
    <div class="mt-0 p-10 text-left">
      <a
        class="no-underline pl-3 text-gray-600 dark:text-white flex"
        href={'/'}
        use:link
      >
        <IconifyIcon
          icon="bi:arrow-left"
          style="position: relative; top: 3px;padding-right: 5px;"
        />Back
      </a>
    </div>
    <slot />
  </div>
</template>
