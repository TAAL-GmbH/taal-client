<script context="module">
  let firstLoad = true
</script>

<script>
  import { onMount } from 'svelte'
  import { useNavigate } from 'svelte-navigator'

  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageBasic from '../../lib/components/page/template/basic/index.svelte'
  import Row from '../../lib/components/layout/row/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Spinner from '../../lib/components/spinner/index.svelte'

  import { spinCount } from '../../lib/stores'
  import * as api from '../../lib/api'

  const navigate = useNavigate()

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
      console.log(
        'on first load, we already have keys and redirect to send data page'
      )
      // navigate('/send-data')
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

{#if $spinCount > 0}
  <Spinner />
{/if}

<style>
  .island {
    text-align: center;
  }
</style>
