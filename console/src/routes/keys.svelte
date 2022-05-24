<script>
  import { onMount } from 'svelte'
  import Key from './key.svelte'
  import Notifications from 'svelte-notifications'
  import Register from './register.svelte'

  let keys

  onMount(async () => {
    const r = await fetch(`${BASE_URL}/api/v1/apikeys`)
    const data = await r.json()
    keys = data.keys
  })
</script>

<div class="panel">
  <p class="panel-heading">API keys</p>
  <div class="panel-body pad">
    <Notifications>
      <Register class="pad" />
    </Notifications>
    {#if keys}
      <table class="table">
        <tr>
          <th>Created</th>
          <th>API Key</th>
          <th>Address</th>
          <th />
        </tr>
        {#each keys as key}
          <Notifications>
            <Key {key} />
          </Notifications>
        {/each}
      </table>
    {:else}
      <p class="loading">loading...</p>
    {/if}
  </div>
</div>

<style>
  .pad {
    padding: 20px;
  }
</style>
