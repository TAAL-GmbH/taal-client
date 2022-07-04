<script>
  import Dropdown from '../../lib/components/dropdown/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Switch from '../../lib/components/switch/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Text from '../../lib/components/text/index.svelte'
  import TextInput from '../../lib/components/textinput/index.svelte'

  let textInputs = {
    apiKey: '',
    taalClientUrl: '',
    mimeType: '',
    tag: '',
  }
  let dropdowns = {
    mode: '',
  }
  let checks = {
    devMode: true,
  }

  function onChange(e) {
    const { name, group, type, value, checked } = e.detail
    switch (type) {
      case 'text':
        textInputs[name] = value || ''
        break
      case 'dropdown':
        dropdowns[name] = value || ''
        break
      case 'checkbox':
        checks[name] = checked || false
        break
    }
  }

  function onInputMount(e) {
    e.detail.inputRef.focus()
  }
</script>

<PageWithMenu>
  <div class="island">
    <Heading value="Send data" />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading value="Parameters" size={2} />
      <Switch
        name="devMode"
        label="Developer mode"
        labelPlacement="right"
        checked={checks['devMode']}
        on:change={onChange}
      />
    </div>
    <Spacer h={24} />
    <TextInput
      name="apiKey"
      label="API key"
      required
      value={textInputs['apiKey']}
      on:change={onChange}
      on:mount={onInputMount}
    />
    <Spacer h={24} />
    <TextInput
      name="taalClientUrl"
      label="TAAL Client URL"
      value={textInputs['taalClientUrl']}
      on:change={onChange}
    />
    <Spacer h={24} />
    <TextInput
      name="mimeType"
      label="MIME type"
      value={textInputs['mimeType']}
      on:change={onChange}
    />
    <Spacer h={24} />
    <TextInput
      name="tag"
      label="Tag"
      value={textInputs['tag']}
      on:change={onChange}
    />
    <Spacer h={32} />
    <div class="sub-row">
      <Heading value="Transaction data" size={2} />
    </div>
    <Spacer h={24} />
    <div class="drop">
      <Dropdown
        class="drop"
        name="mode"
        label="Mode"
        required
        value={dropdowns['mode']}
        items={[
          { value: '1', label: 'Mode 1' },
          { value: '2', label: 'Mode 2' },
        ]}
        on:change={onChange}
      />
    </div>

    <Spacer h={24} />
  </div>
</PageWithMenu>

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

  .drop {
    max-width: 445px;
  }
</style>
