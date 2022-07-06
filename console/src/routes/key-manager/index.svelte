<script>
  import { onMount } from 'svelte'
  import { getNotificationsContext } from 'svelte-notifications'

  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import KeyCard from '../../lib/components/cards/key-card/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import RegisterKeyPopup from '../../lib/components/popups/register-key-popup/index.svelte'
  import Spinner from '../../lib/components/spinner/index.svelte'

  import { spinCount } from '../../lib/stores'
  import * as api from '../../lib/api'

  const { addNotification } = getNotificationsContext()

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
      (data) => {
        keys = data.key_usages
      },
      (error) => {
        addNotification({
          text: `Failed to load api keys: ${error}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 2000,
        })
      }
    )
  }

  function register(apiKey) {
    api.register(
      apiKey,
      (data) => {
        hidePopup()

        addNotification({
          text: `API key registered successfully`,
          position: 'bottom-left',
          type: 'success',
          removeAfter: 1000,
        })

        getApiKeys()
      },
      (error) => {
        addNotification({
          text: `Error: ${error}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 5000,
        })
      }
    )
  }

  function onRegister(e) {
    console.log('onRegister: apiKey = ', e.detail.apiKey)
    register(e.detail.apiKey)
  }

  function onDeactivate(e) {
    console.log('onDeactivate: key = ', e.detail.key)
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
          <KeyCard {key} on:deactivate={onDeactivate} />
        {/each}
      </div>
    {/if}
  </div>
</PageWithMenu>

{#if popupVisible}
  <RegisterKeyPopup on:close={hidePopup} on:register={onRegister} />
{/if}

{#if $spinCount > 0}
  <Spinner />
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
