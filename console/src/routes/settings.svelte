<script>
  import Keys from './keys.svelte'
  import Register from './register.svelte'
  import Config from './config.svelte'
  import Notifications from 'svelte-notifications'

  let keys
</script>

<main>
  <Notifications>
    <Config />
    <Register />
  </Notifications>
  <Keys {keys} />
</main>
