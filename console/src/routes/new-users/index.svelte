<script>
  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageBasic from '../../lib/components/page/template/basic/index.svelte'
  import Row from '../../lib/components/layout/row/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'

  function onKey() {
    console.log('onKey')
  }

  function onNoKey() {
    console.log('onNoKey')
  }
</script>

<PageBasic>
  <div class="island">
    <Heading>Do you have an API key?</Heading>
    <Spacer h={64} />
    <Row gap={16}>
      <Button variant="secondary" width={236} on:click={onNoKey}
        >I don't have an API key</Button
      >
      <Button width={236} on:click={onKey}>I have an API key</Button>
    </Row>
  </div>
</PageBasic>

<style>
  .island {
    text-align: center;
    margin-bottom: 120px;
  }
</style>
