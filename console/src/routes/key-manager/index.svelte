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
  import i18n from '../../lib/i18n'

  $: t = $i18n.t
  const pageKey = 'page.key-manager'

  // injected by svelte-navigator
  export let location = null
  export let navigate = null

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
        failure(t(`${pageKey}.notifications.api-keys-load-failure`, { error }))
    )
  }

  function registerKey(apiKey) {
    api.registerKey(
      apiKey,
      (data) => {
        hidePopup()
        success(t(`${pageKey}.notifications.key-register-success`))
        getApiKeys()
      },
      (error) =>
        failure(t('notifications.failure', { error }), { duration: 5000 })
    )
  }

  function deleteKey(apiKey) {
    api.deleteKey(
      apiKey,
      (data) => {
        success(t(`${pageKey}.notifications.key-delete-success`))
        getApiKeys()
      },
      (error) =>
        failure(t('notifications.failure', { error }), { duration: 5000 })
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
    <Heading value={t(`${pageKey}.key-manager-label`)} />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading value={t(`${pageKey}.registered-keys-label`)} size={2} />
      <Button icon="plus" on:click={showPopup}
        >{t(`${pageKey}.add-new-label`)}</Button
      >
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
