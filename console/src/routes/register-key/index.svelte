<script>
  import Button from '../../lib/components/button/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageBasic from '../../lib/components/page/template/basic/index.svelte'
  import Row from '../../lib/components/layout/row/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Text from '../../lib/components/text/index.svelte'
  import TextInput from '../../lib/components/text-input/index.svelte'
  import { link } from '../../lib/utils/format'

  function onCancel() {
    console.log('onCancel')
  }

  function onRegister() {
    console.log('onRegister')
  }

  let key = ''

  function onInputChange(e) {
    const target = e.explicitOriginalTarget
    key = target.value
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
