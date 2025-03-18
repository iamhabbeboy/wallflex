<script type="ts">
  import Main from './pages/Main.svelte';
  import Setting from './pages/Setting.svelte';
  import PlainArt from './pages/PlainArt.svelte';

  import type { SvelteComponentTyped } from 'svelte';
  import {
    Router as RouterComponent,
    Route as RouteComponent,
  } from 'svelte-routing';

  const Router: typeof SvelteComponentTyped = RouterComponent;
  const Route: typeof SvelteComponentTyped = RouteComponent;
  import AbstractArt from './pages/AbstractArt.svelte';
  import Ai from './pages/AI.svelte';
  export let url = '';
</script>

<main>
  <Router {url}>
    <Route path="/setting" component={Setting} />
    <Route path="/ai" component={Ai} />
    <!-- <Route path="/abstract-art" component={AbstractArt} /> -->
    <Route path="/" component={Main} />
    <Route path="/plain-art" component={PlainArt} />
  </Router>
</main>
