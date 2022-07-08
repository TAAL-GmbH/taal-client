<script>
  import { useNavigate } from 'svelte-navigator'

  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageBasic from '../../lib/components/page/template/basic/index.svelte'
  import Row from '../../lib/components/layout/row/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Spinner from '../../lib/components/spinner/index.svelte'

  const navigate = useNavigate()
  let showSpinner = false

  function onKey() {
    navigate('/register-key')
  }

  function onNoKey() {
    window.open('https://console.taal.com', '_blank')
  }
</script>

<PageBasic>
  <div class="island">
    <Heading value="Do you have an API key?" />
    <Spacer h={64} />
    <Row gap={16}>
      <Button variant="secondary" width={236} on:click={onNoKey}>
        I don't have an API key
      </Button>
      <Button width={236} on:click={onKey}>I have an API key</Button>
    </Row>
  </div>
</PageBasic>

{#if showSpinner}
  <Spinner />
{/if}

<style>
  .island {
    text-align: center;
    margin-bottom: 120px;
  }
</style>
