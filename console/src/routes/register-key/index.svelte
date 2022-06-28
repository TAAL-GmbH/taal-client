<script>
  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageBasic from '../../lib/components/page/template/basic/index.svelte'
  import Row from '../../lib/components/layout/row/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'

  function onCancel() {
    console.log('onCancel')
  }

  function onRegister() {
    console.log('onRegister')
  }
</script>

<PageBasic>
  <div class="island">
    <Heading>Register API key</Heading>
    <Spacer h={64} />
    TODO: input here
    <Spacer h={64} />
    <div class="right">
      <Row gap={16}>
        <Button variant="ghost" on:click={onCancel}>Cancel</Button>
        <Button on:click={onRegister}>Register</Button>
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

    margin-bottom: 120px;
  }

  .right {
    text-align: right;
  }
</style>
