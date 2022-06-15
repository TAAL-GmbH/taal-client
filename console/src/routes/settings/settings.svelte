<script>
  import Keys from './keys.svelte'
  import Config from './config.svelte'
  import Notifications from 'svelte-notifications'
  import { createEventDispatcher } from 'svelte'
  import { onMount } from 'svelte'

  const dispatch = createEventDispatcher()

  onMount(() => {
    dispatch('mounted', {
      message: 'settings',
    })
  })
</script>

<main>
  <Notifications>
    <Keys />
  </Notifications>
  <Notifications>
    <Config />
  </Notifications>
</main>
