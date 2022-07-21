<script>
  import { useNavigate } from 'svelte-navigator'

  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageBasic from '../../lib/components/page/template/basic/index.svelte'
  import Row from '../../lib/components/layout/row/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Text from '../../lib/components/text/index.svelte'
  import TextInput from '../../lib/components/textinput/index.svelte'

  import { link } from '../../lib/utils/format'
  import { success, failure } from '../../lib/utils/notifications'
  import * as api from '../../lib/api'

  const navigate = useNavigate()

  let key = ''

  function onInputChange(e) {
    key = e.detail.value
  }

  async function registerKey(apiKey) {
    let ok = false
    await api.registerKey(
      apiKey,
      (data) => {
        success('API key registered successfully', { duration: 1000 })
        ok = true
      },
      (error) => failure(`Error: ${error}`, { duration: 5000 })
    )
    if (ok) {
      navigate('/key-manager')
    }
  }

  function onReset() {
    key = ''
  }

  function onRegister() {
    registerKey(key)
  }
</script>

<PageBasic>
  <div class="island">
    <Heading value="Register API key" />
    <Spacer h={48} />
    <TextInput
      name="api-key"
      label="API key"
      value={key}
      required
      on:change={onInputChange}
    />
    <Spacer h={8} />
    <Text
      size={5}
      html
      value="API keys can be obtained from {link('https://console.taal.com')}"
    />
    <Spacer h={120} />
    <div class="right">
      <Row gap={16}>
        <Button variant="ghost" size="large" on:click={onReset}>Reset</Button>
        <Button size="large" on:click={onRegister}>Register</Button>
      </Row>
    </div>
  </div>
</PageBasic>

<style>
  .island {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 920px;
  }

  .right {
    text-align: right;
  }
</style>
