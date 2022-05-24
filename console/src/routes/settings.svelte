<script>
  import Keys from './keys.svelte'
  import Config from './config.svelte'
  import Notifications from 'svelte-notifications'

  let keys
</script>

<main>
  <Notifications>
    <Config />
  </Notifications>
  <Keys {keys} />
</main>
