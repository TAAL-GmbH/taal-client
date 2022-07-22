<script context="module">
  let firstLoad = true
</script>

<script>
  import { onMount } from 'svelte'

  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageBasic from '../../lib/components/page/template/basic/index.svelte'
  import Row from '../../lib/components/layout/row/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'

  import * as api from '../../lib/api'

  // injected by svelte-navigator
  export let location = null
  export let navigate = null

  function onKey() {
    navigate('/register-key')
  }

  function onNoKey() {
    window.open('https://console.taal.com', '_blank')
  }

  onMount(async () => {
    const result = await api.getApiKeys()

    console.log(result)

    if (firstLoad && result?.keys && result?.keys.length > 0) {
      firstLoad = false
      navigate('/send-data')
    }
  })
</script>

<PageBasic>
  <div class="island">
    <Heading value="Do you have an API key?" />
    <Spacer h={64} />
    <Row gap={16}>
      <Button variant="secondary" size="large" width={236} on:click={onNoKey}>
        I don't have an API key
      </Button>
      <Button width={236} size="large" on:click={onKey}>
        I have an API key
      </Button>
    </Row>
  </div>
</PageBasic>

<style>
  .island {
    text-align: center;
  }
</style>
