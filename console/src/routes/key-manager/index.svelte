<script>
  import { onMount } from 'svelte'
  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import KeyCard from '../../lib/components/cards/key-card/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import RegisterKeyPopup from '../../lib/components/popups/register-key-popup/index.svelte'
  import Spinner from '../../lib/components/spinner/index.svelte'
  import { getNotificationsContext } from 'svelte-notifications'

  const { addNotification } = getNotificationsContext()

  let keys

  let popupVisible = false
  function showPopup() {
    popupVisible = true
  }
  function hidePopup() {
    popupVisible = false
  }

  let spinCount = 0
  function showSpinner() {
    spinCount++
  }
  function hideSpinner() {
    spinCount--
  }

  function getApiKeys() {
    showSpinner()

    fetch(`${BASE_URL}/api/v1/apikeys`)
      .then((res) => {
        if (!res.ok) {
          return res.text().then((text) => {
            throw new Error(text)
          })
        }

        return res.json()
      })
      .then((data) => {
        keys = data.keys
      })
      .catch((err) => {
        var errMessage = ''
        try {
          const errJson = JSON.parse(err.message)
          errMessage = errJson.error
        } catch (error) {
          errMessage = err
        }

        addNotification({
          text: `Failed to load api keys: ${errMessage}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 2000,
        })

        console.log(errMessage)
      })
      .finally(hideSpinner)
  }

  function register(apiKey) {
    showSpinner()

    fetch(`${BASE_URL}/api/v1/apikeys/${apiKey}`, {
      method: 'POST',
    })
      .then((res) => {
        hidePopup()

        if (!res.ok) {
          return res.text().then((text) => {
            throw new Error(text)
          })
        } else {
          addNotification({
            text: `API key registered successfully`,
            position: 'bottom-left',
            type: 'success',
            removeAfter: 1000,
          })

          getApiKeys()

          return res.json()
        }
      })
      .catch((err) => {
        const errJson = JSON.parse(err.message)
        addNotification({
          text: `Error: ${errJson.error}`,
          position: 'bottom-left',
          type: 'danger',
          removeAfter: 5000,
        })

        console.log(err)
      })
      .finally(hideSpinner)
  }

  function onRegister(e) {
    console.log('onRegister: key = ', e.detail.key)
    register(e.detail.key)
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

{#if spinCount > 0}
  <Spinner />
{/if}

<style>
  .island {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 920px;
    padding-top: 40px;
    margin-bottom: 100px;
  }

  .sub-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(400px, 1fr));
    column-gap: 16px;
    row-gap: 16px;
  }
</style>
