<script>
  // svelte-ignore unused-export-let
  export let location

  import Keys from './keys.svelte'
  import Config from './config.svelte'
  import Notifications from 'svelte-notifications'
</script>

<main>
  <Notifications>
    <Keys />
  </Notifications>
  <Notifications>
    <Config />
  </Notifications>
</main>
