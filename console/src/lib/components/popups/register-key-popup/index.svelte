<script>
  import { createEventDispatcher } from 'svelte'
  import Button from '../../../../lib/components/button/index.svelte'
  import Modal from '../../../../lib/components/modal/index.svelte'
  import Popup from '../../../../lib/components/popup/index.svelte'
  import Spacer from '../../../../lib/components/layout/spacer/index.svelte'
  import Text from '../../../../lib/components/text/index.svelte'
  import TextInput from '../../../../lib/components/textinput/index.svelte'
  import { link } from '../../../../lib/utils/format'

  const dispatch = createEventDispatcher()

  let apiKey = ''

  function onInputChange(e) {
    apiKey = e.detail.value
  }

  function onInputMount(e) {
    e.detail.inputRef.focus()
  }

  function onClose() {
    dispatch('close')
  }
  function onRegisterClick() {
    dispatch('register', { apiKey })
  }
</script>

<Modal>
  <Popup maxW={480} title="Register API key" on:close={onClose}>
    <svelte:fragment slot="body">
      <Spacer h={20} />
      <TextInput
        name="api-key"
        label="API key"
        value={apiKey}
        required
        on:change={onInputChange}
        on:mount={onInputMount}
      />
      <Spacer h={8} />
      <Text
        size={5}
        html
        value="API keys can be obtained from {link('https://console.taal.com')}"
      />
      <Spacer h={80} />
    </svelte:fragment>
    <svelte:fragment slot="footer">
      <Button variant="ghost" size="small" on:click={onClose}>Cancel</Button>
      <Button size="small" on:click={onRegisterClick}>Register</Button>
    </svelte:fragment>
  </Popup>
</Modal>
