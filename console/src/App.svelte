<script>
  import Notifications from 'svelte-notifications'

  import AppOld from './AppOld.svelte'
  import AppNew from './AppNew.svelte'

  let old = false
</script>

{#if old}
  <AppOld />
{:else}
  <Notifications>
    <AppNew />
  </Notifications>
{/if}
