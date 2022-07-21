<script>
  import { onMount } from 'svelte'

  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import KeyCard from '../../lib/components/cards/key-card/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import RegisterKeyPopup from '../../lib/components/popups/register-key-popup/index.svelte'

  import * as api from '../../lib/api'
  import { success, failure } from '../../lib/utils/notifications'

  let keys

  let popupVisible = false
  function showPopup() {
    popupVisible = true
  }
  function hidePopup() {
    popupVisible = false
  }

  function getApiKeys() {
    api.getApiKeysUsage(
      (data) => (keys = data.key_usages),
      (error) =>
        failure(`Failed to load api keys: ${error}`, { duration: 2000 })
    )
  }

  function registerKey(apiKey) {
    api.registerKey(
      apiKey,
      (data) => {
        hidePopup()
        success('API key registered successfully', { duration: 1000 })
        getApiKeys()
      },
      (error) => failure(`Error: ${error}`, { duration: 5000 })
    )
  }

  function deleteKey(apiKey) {
    api.deleteKey(
      apiKey,
      (data) => {
        success('API key deleted successfully', { duration: 1000 })
        getApiKeys()
      },
      (error) => failure(`Error: ${error}`, { duration: 5000 })
    )
  }

  function onRegister(e) {
    registerKey(e.detail.apiKey)
  }

  onMount(() => {
    getApiKeys()
  })
</script>

<PageWithMenu>
  <div class="island">
    <Heading value="Key manager" />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading value="Registered API keys" size={2} />
      <Button icon="plus" on:click={showPopup}>Add new</Button>
    </div>
    <Spacer h={24} />
    {#if keys}
      <div class="grid">
        {#each keys as key (key.api_key)}
          <KeyCard {key} />
        {/each}
      </div>
    {/if}
  </div>
</PageWithMenu>

{#if popupVisible}
  <RegisterKeyPopup on:close={hidePopup} on:register={onRegister} />
{/if}

<style>
  .island {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 920px;
  }

  .sub-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));
    column-gap: 16px;
    row-gap: 16px;
  }
</style>
