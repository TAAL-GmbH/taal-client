<script>
  import Keys from './keys.svelte'
  import Register from './register.svelte'
  import Notifications from 'svelte-notifications'

  let keys
</script>

<main>
  <Notifications>
    <Register />
  </Notifications>
  <Keys {keys} />
</main>
