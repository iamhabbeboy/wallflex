<script type="ts">
  import Main from './pages/Main.svelte';
  import Setting from './pages/Setting.svelte';
  import { Router, Route } from 'svelte-routing';
  export let url = '';
</script>

<main>
  <Router {url}>
    <Route path="/setting" component={Setting} />
    <Route path="/"><Main /></Route>
  </Router>
</main>
